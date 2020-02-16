package handlers

import (
	"github.com/golang/protobuf/proto"
	"github.com/sleep2death/goblin/pbs"
	"github.com/sleep2death/gotham"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// register
func registerHandler(c *gotham.Context) {
	var req pbs.Register
	var resp *pbs.RegisterAck = &pbs.RegisterAck{}
	// Unmarshal the register request message.
	if err := proto.Unmarshal(c.Request.Data, &req); err != nil {
		panic(err)
	}

	// TODO: username/password/email validation
	// Hash password from request.
	hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	token, err := dbRegister(req.GetUsername(), req.GetEmail(), string(hash))
	if err != nil {
		if err == ErrUserNameAlreadyExisted {
			resp.Error = MsgUserAlreadyExisted
			abortWithMessage(c, resp)
			logger.Warn(err.Error())
		} else {
			panic(err)
		}
	}

	resp.Token = token
	c.Write(resp)
}

func loginHandler(c *gotham.Context) {
	var req pbs.Login
	var resp *pbs.LoginAck = &pbs.LoginAck{}
	// Unmarshal the register request message.
	if err := proto.Unmarshal(c.Request.Data, &req); err != nil {
		panic(err)
	}

	token, err := dbLogin(req.GetUsername(), req.GetPassword())
	if err != nil {
		if err == mongo.ErrNoDocuments ||
			err == bcrypt.ErrMismatchedHashAndPassword {
			resp.Error = MsgLoginFailed
			abortWithMessage(c, resp)
			logger.Warn(err.Error())
		} else {
			panic(err)
		}
	}

	resp.Token = token
	c.Write(resp)
}
