package handlers

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/lithammer/shortuuid/v3"
	"github.com/sleep2death/goblin/pbs"
	"github.com/sleep2death/gotham"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	doOnce.Do(initTest)

	uuid := shortuuid.New()
	msgA := &pbs.Register{
		Username: uuid,
		Email:    uuid + "@goblin.com",
		Password: "Password!",
	}

	tu := proto.MessageName(msgA)
	data, _ := proto.Marshal(msgA)
	req := &gotham.Request{
		TypeURL: tu,
		Data:    data,
	}

	rw := &ResponseRecorder{keepAlive: true}
	router.ServeProto(rw, req)

	msgB, _ := rw.message.(*pbs.RegisterAck)
	assert.Nil(t, msgB.GetError())
	assert.Greater(t, len(msgB.GetToken()), 0)
	assert.True(t, rw.KeepAlive())

	// register again with same username
	router.ServeProto(rw, req)
	msgB, _ = rw.message.(*pbs.RegisterAck)
	assert.Equal(t, MsgUserAlreadyExisted, msgB.GetError())
	assert.False(t, rw.KeepAlive())

}
func TestLoginHandler(t *testing.T) {
	doOnce.Do(initTest)

	msgA := &pbs.Register{
		Username: "aspirin",
		Email:    "aspirin@goblin.com",
		Password: "Password!",
	}

	tu := proto.MessageName(msgA)
	data, _ := proto.Marshal(msgA)
	req := &gotham.Request{
		TypeURL: tu,
		Data:    data,
	}

	rw := &ResponseRecorder{keepAlive: true}
	router.ServeProto(rw, req)

	msgC := &pbs.Login{
		Username: "aspirin",
		Password: "Password!",
	}
	tu = proto.MessageName(msgC)
	data, _ = proto.Marshal(msgC)
	req = &gotham.Request{
		TypeURL: tu,
		Data:    data,
	}
	rw = &ResponseRecorder{keepAlive: true}
	router.ServeProto(rw, req)
	msgD, _ := rw.message.(*pbs.LoginAck)
	assert.Nil(t, msgD.GetError())
	assert.Greater(t, len(msgD.GetToken()), 0)
	assert.True(t, rw.KeepAlive())

	// invalid password
	msgC.Password = "password!!"
	tu = proto.MessageName(msgC)
	data, _ = proto.Marshal(msgC)
	req = &gotham.Request{
		TypeURL: tu,
		Data:    data,
	}
	router.ServeProto(rw, req)
	msgD, _ = rw.message.(*pbs.LoginAck)
	assert.Equal(t, MsgLoginFailed.GetCode(), msgD.GetError().GetCode())
	assert.False(t, rw.KeepAlive())
}
