package goblin

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/proto"
	"github.com/sleep2death/goblin/pb"
	"github.com/sleep2death/gotham"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserCollection  = "users"
	DBReadTimeout   = time.Second * 5
	DBWriteTimeout  = time.Second * 5
	TokenExpireTime = time.Hour * 72
)

var (
	MsgUserAlreadyExisted = &pb.Error{
		Code:       http.StatusConflict,
		Desciption: "Username already existed.",
	}
	MsgInternalServerError = &pb.Error{
		Code:       http.StatusInternalServerError,
		Desciption: "Oops...some wrong",
	}
)

func getRegisterHandler(db *mongo.Database) gotham.HandlerFunc {
	return func(c *gotham.Context) {
		var req pb.Register
		var resp *pb.RegisterAck = &pb.RegisterAck{}

		jwtKey := viper.GetString("tokenjwtkey")
		if len(jwtKey) == 0 {
			// c.AbortWithStatus(http.StatusInternalServerError)
			resp.Error = MsgInternalServerError
			c.Write(resp)
			return
		}
		// Unmarshal the register request message.
		if err := proto.Unmarshal(c.Data(), &req); err != nil {
			resp.Error = MsgInternalServerError
			c.Write(resp)
			return
		}

		// Hash password from request.
		hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
		if err != nil {
			resp.Error = MsgInternalServerError
			c.Write(resp)
			return
		}

		// FindOne, if not exist register one.
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
					ExpiresAt: time.Now().Add(time.Second * 60).Unix(),
				}
				// create jwt token
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenStr, err := token.SignedString(jwtKey)
				if err != nil {
					resp.Error = MsgInternalServerError
					c.Write(resp)
					return
				}
				// log.Println(tokenStr)
				c.Write(&pb.RegisterAck{Token: tokenStr})
			} else {
				resp.Error = MsgInternalServerError
				c.Write(resp)
			}
		} else {
			resp.Error = MsgUserAlreadyExisted
			c.Write(resp)
		}
	}
}

func getLoginHandler(db *mongo.Database) gotham.HandlerFunc {
	return func(c *gotham.Context) {
	}
}
