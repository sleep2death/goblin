package handlers

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/sleep2death/goblin/pbs"
	"github.com/sleep2death/gotham"
	"github.com/stretchr/testify/assert"
)

func TestRecovery(t *testing.T) {
	doOnce.Do(initTest)

	msgC := &pbs.Login{
		Username: "aspirin",
		Password: "Password!",
	}
	typeurl := proto.MessageName(msgC)
	data, _ := proto.Marshal(msgC)
	req := &gotham.Request{
		TypeURL: typeurl,
		Data:    data[:3], //invalid data
	}

	rw := &ResponseRecorder{keepAlive: true}
	router.ServeProto(rw, req)
	msgD, _ := rw.message.(*pbs.Error)
	assert.Equal(t, MsgInternalServerError.GetDescription(), msgD.GetDescription())
	assert.False(t, rw.KeepAlive())
}
