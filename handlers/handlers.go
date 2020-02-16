package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/proto"
	"github.com/sleep2death/goblin/pbs"
	"github.com/sleep2death/gotham"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserCollection = "users"
)

var (
	logger   *zap.Logger
	database *mongo.Database

	MsgRouterNotFound = &pbs.Error{
		Code:        http.StatusNotFound,
		Description: "router not found.",
	}

	MsgUserAlreadyExisted = &pbs.Error{
		Code:        http.StatusConflict,
		Description: "username already exists.",
	}
	MsgLoginFailed = &pbs.Error{
		Code:        http.StatusUnauthorized,
		Description: "username or password is mismatch.",
	}
	MsgInternalServerError = &pbs.Error{
		Code:        http.StatusInternalServerError,
		Description: "oops...something went wrong.",
	}
	MsgBadUserRequest = &pbs.Error{
		Code:        http.StatusBadRequest,
		Description: "invalid user input.",
	}

	ErrUserNameAlreadyExisted = errors.New("username already existed.")

	jwtKey        string
	tokenExpire   time.Duration
	dbReadTimeout time.Duration
)

func InitHandlers(router *gotham.Router, db *mongo.Database, log *zap.Logger) {
	logger = log
	database = db

	jwtKey = viper.GetString("tokenjwtkey")
	if len(jwtKey) == 0 {
		logger.Fatal("jwtkey not found.")
	}
	tokenExpire = viper.GetDuration("tokenexpiretime")
	if tokenExpire == 0 {
		logger.Fatal("tokenexpiretime not found.")
	}
	dbReadTimeout = viper.GetDuration("dbreadtimeout")
	if dbReadTimeout == 0 {
		logger.Fatal("dbreadtimeout not found.")
	}

	router.Handle("pbs.Login", loginHandler)
	router.Handle("pbs.Register", registerHandler)
	router.NoRoute(func(c *gotham.Context) {
		logger.Warn("router not found.", zap.String("typeurl", c.Request.TypeURL))
		abortWithMessage(c, MsgRouterNotFound)
	})
}

// Abort context, write the final message,
// and close the connection.
func abortWithMessage(c *gotham.Context, msg proto.Message) {
	c.Write(msg)
	c.Abort()

}

func dbRegister(username, email, password string) (string, error) {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.M{"username": username}
	update := bson.M{"$setOnInsert": bson.M{"username": username, "email": email, "password": password}}
	ctx, cancel := context.WithTimeout(context.Background(), dbReadTimeout)
	defer cancel()
	res := database.Collection(UserCollection).FindOneAndUpdate(ctx, filter, update, opts)

	if err := res.Err(); err != nil {
		// Username not existed, upsert will create one.
		// then, server will generate the token for callback
		if err == mongo.ErrNoDocuments {
			claims := &jwt.StandardClaims{
				Id:        username,
				ExpiresAt: time.Now().Add(tokenExpire).Unix(),
			}
			// create jwt token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenStr, err := token.SignedString([]byte(jwtKey))
			if err != nil {
				return "", err
			}
			return tokenStr, nil
		}
		return "", err
	}

	return "", ErrUserNameAlreadyExisted
}

func dbLogin(username, password string) (string, error) {
	opts := options.FindOne().SetProjection(bson.M{"password": 1})
	filter := bson.M{"username": username}
	ctx, cancel := context.WithTimeout(context.Background(), dbReadTimeout)
	defer cancel()
	res := database.Collection(UserCollection).FindOne(ctx, filter, opts)
	if err := res.Err(); err != nil {
		return "", err
	}

	// decode findone result
	var r bson.M
	if err := res.Decode(&r); err != nil {
		return "", err
	}

	hash, _ := r["password"].(string)

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return "", err
	}

	claims := &jwt.StandardClaims{
		Id:        username,
		ExpiresAt: time.Now().Add(tokenExpire).Unix(),
	}

	// create jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
