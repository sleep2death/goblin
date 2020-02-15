package handler

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
		resp.Error = MsgInternalServerError
		abortWithMessage(c, resp)
		logger.Error(err.Error())
		return
	}

	// TODO: username/password/email validation
	// Hash password from request.
	hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		resp.Error = MsgInternalServerError
		abortWithMessage(c, resp)
		logger.Error(err.Error())
		return
	}

	token, err := dbRegister(req.GetUsername(), req.GetEmail(), string(hash))
	if err != nil {
		logger.Error(err.Error())
		if err == ErrUserNameAlreadyExisted {
			resp.Error = MsgUserAlreadyExisted
		} else {
			resp.Error = MsgInternalServerError
		}
		abortWithMessage(c, resp)
	}

	resp.Token = token
	c.Write(resp)
}

func loginHandler(c *gotham.Context) {
	var req pbs.Login
	var resp *pbs.LoginAck = &pbs.LoginAck{}
	// Unmarshal the register request message.
	if err := proto.Unmarshal(c.Request.Data, &req); err != nil {
		resp.Error = MsgInternalServerError
		abortWithMessage(c, resp)
		logger.Error(err.Error())
		return
	}

	token, err := dbLogin(req.GetUsername(), req.GetPassword())
	if err != nil {
		if err == mongo.ErrNoDocuments ||
			err == bcrypt.ErrMismatchedHashAndPassword {
			resp.Error = MsgLoginFailed
		} else {
			resp.Error = MsgInternalServerError
		}
		abortWithMessage(c, resp)
		logger.Error(err.Error())
	}

	resp.Token = token
	c.Write(resp)
}
