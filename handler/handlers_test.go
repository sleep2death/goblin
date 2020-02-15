package handler

import (
	"sync"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/lithammer/shortuuid/v3"
	"github.com/sleep2death/goblin/pbs"
	"github.com/sleep2death/goblin/utils"
	"github.com/sleep2death/gotham"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	doOnce sync.Once
	router *gotham.Router
)

func initTest() {
	// create logger
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	l, _ := config.Build()
	defer l.Sync()

	if err := utils.InitConfig(); err != nil {
		l.Fatal(err.Error())
	}

	db, err := utils.InitDB(viper.GetString("dbaddr"), viper.GetString("dbname"))
	if err != nil {
		l.Fatal(err.Error())
	}

	l.Info("connected to the db")

	// create the router
	router = gotham.New()
	InitHandlers(router, db, l)
}

type ResponseRecorder struct {
	message   proto.Message
	status    int
	keepAlive bool
}

func (rw *ResponseRecorder) SetStatus(code int) {
	rw.status = code
}

func (rw *ResponseRecorder) Status() int {
	return rw.status
}

func (rw *ResponseRecorder) KeepAlive() bool {
	return rw.keepAlive
}

func (rw *ResponseRecorder) SetKeepAlive(value bool) {
	rw.keepAlive = value
}

func (rw *ResponseRecorder) Buffered() int {
	return 0
}

func (rw *ResponseRecorder) Flush() error {
	return nil
}

func (rw *ResponseRecorder) Write(message proto.Message) error {
	rw.message = message
	return nil
}

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
