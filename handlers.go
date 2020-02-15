package goblin

import (
	"context"
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
	UserCollection  = "users"
	DBReadTimeout   = time.Second * 5
	DBWriteTimeout  = time.Second * 5
	TokenExpireTime = time.Hour * 72
)

var (
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
)

// RegisterHandler
func getRegisterHandler(db *mongo.Database) gotham.HandlerFunc {
	jwtKey := viper.GetString("tokenjwtkey")
	if len(jwtKey) == 0 {
		logger.Fatal("jwtkey not found.")
	}
	expire := viper.GetDuration("tokenexpiretime")
	return func(c *gotham.Context) {
		var req pbs.Register
		var resp *pbs.RegisterAck = &pbs.RegisterAck{}
		// Unmarshal the register request message.
		if err := proto.Unmarshal(c.Request.Data, &req); err != nil {
			resp.Error = MsgInternalServerError
			AbortWithMessage(c, resp)
			logger.Error(err.Error())
			return
		}

		// TODO: username/password/email validation
		// Hash password from request.
		hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
		if err != nil {
			resp.Error = MsgInternalServerError
			AbortWithMessage(c, resp)
			logger.Error(err.Error())
			return
		}

		// FindOne, if not exists, register one.
		opts := options.FindOneAndUpdate().SetUpsert(true)
		filter := bson.M{"username": req.GetUsername()}
		update := bson.M{"$setOnInsert": bson.M{"username": req.GetUsername(), "email": req.GetEmail(), "password": string(hash)}}
		ctx, cancel := context.WithTimeout(context.Background(), DBReadTimeout)
		defer cancel()
		res := db.Collection(UserCollection).FindOneAndUpdate(ctx, filter, update, opts)

		if err := res.Err(); err != nil {
			// Username not existed, upsert will create one.
			// then, server will generate the token for callback
			if err == mongo.ErrNoDocuments {
				claims := &jwt.StandardClaims{
					Id:        req.GetUsername(),
					ExpiresAt: time.Now().Add(expire).Unix(),
				}
				// create jwt token
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenStr, err := token.SignedString([]byte(jwtKey))
				if err != nil {
					resp.Error = MsgInternalServerError
					AbortWithMessage(c, resp)
					logger.Error(err.Error())
					return
				}
				resp.Token = tokenStr
				c.Write(resp)
				return
			}

			resp.Error = MsgInternalServerError
			AbortWithMessage(c, resp)
			logger.Error(err.Error())
			return
		}

		resp.Error = MsgUserAlreadyExisted
		AbortWithMessage(c, resp)
		logger.Info("username already exists.", zap.String("name", req.GetUsername()))
	}
}

// login binding
type login struct {
	Username string `bson:"username"  binding:"required"`
	Password string `bson:"password" binding:"required"`
}

// LoginHandler
func getLoginHandler(db *mongo.Database) gotham.HandlerFunc {
	jwtKey := viper.GetString("tokenjwtkey")
	if len(jwtKey) == 0 {
		logger.Fatal("jwtkey not found.")
	}

	expire := viper.GetDuration("tokenexpiretime")
	return func(c *gotham.Context) {
		var req pbs.Login
		var resp *pbs.LoginAck = &pbs.LoginAck{}
		// Unmarshal the register request message.
		if err := proto.Unmarshal(c.Request.Data, &req); err != nil {
			resp.Error = MsgInternalServerError
			AbortWithMessage(c, resp)
			logger.Error(err.Error())
			return
		}
		opts := options.FindOne()
		filter := bson.M{"username": req.GetUsername()}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		res := db.Collection(UserCollection).FindOne(ctx, filter, opts)
		if err := res.Err(); err != nil {
			// Username not existed.
			if err == mongo.ErrNoDocuments {
				resp.Error = MsgLoginFailed
			} else {
				resp.Error = MsgInternalServerError
			}

			AbortWithMessage(c, resp)
			logger.Error(err.Error())
			return
		}

		// decode findone result
		r := &login{}
		if err := res.Decode(r); err != nil {
			resp.Error = MsgInternalServerError
			AbortWithMessage(c, resp)
			logger.Error(err.Error())
			return
		}

		// compare password
		if cErr := bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(req.Password)); cErr != nil {
			resp.Error = MsgLoginFailed
			AbortWithMessage(c, resp)
			logger.Error(cErr.Error())
			return
		}

		claims := &jwt.StandardClaims{
			Id:        r.Username,
			ExpiresAt: time.Now().Add(expire).Unix(),
		}

		// create jwt token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString([]byte(jwtKey))
		if err != nil {
			resp.Error = MsgInternalServerError
			AbortWithMessage(c, resp)
			logger.Error(err.Error())
			return
		}
		resp.Token = tokenStr
		c.Write(resp)
	}
}
