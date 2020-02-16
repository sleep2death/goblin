package handlers

import (
	"sync"
	"testing"

	"github.com/golang/protobuf/proto"
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
	router.Use(Recovery(l))
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

func TestInitHandlers(t *testing.T) {
	doOnce.Do(initTest)
	assert.Len(t, router.Routes(), 2)
}

func TestNoRouter(t *testing.T) {
	doOnce.Do(initTest)

	msgC := &pbs.Login{
		Username: "aspirin",
		Password: "Password!",
	}
	// typeurl := proto.MessageName(msgC)
	data, _ := proto.Marshal(msgC)
	req := &gotham.Request{
		TypeURL: "abc",
		Data:    data, //invalid data
	}

	rw := &ResponseRecorder{keepAlive: true}
	router.ServeProto(rw, req)
	msgD, _ := rw.message.(*pbs.Error)
	assert.Equal(t, MsgRouterNotFound.GetCode(), msgD.GetCode())
	assert.False(t, rw.KeepAlive())
}
