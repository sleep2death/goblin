package goblin

import (
	"bufio"
	"net"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/lithammer/shortuuid"
	"github.com/sleep2death/goblin/pbs"
	"github.com/sleep2death/gotham"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	go Serve()

	time.Sleep(time.Millisecond * 50)
	defer server.Close()

	conn, err := net.Dial("tcp", ":9001")
	if err != nil {
		t.Fatal(err)
	}

	bufw := bufio.NewWriter(conn)
	bufr := bufio.NewReader(conn)

	// generate an uuid for username
	u := shortuuid.New()
	msg := &pbs.Register{
		Username: u,
		Email:    u + "@goblin.com",
		Password: "password!",
	}

	gotham.WriteFrame(bufw, msg)
	bufw.Flush()

	res, _ := gotham.ReadFrame(bufr)
	var resp pbs.RegisterAck
	proto.Unmarshal(res.Data(), &resp)
	assert.Nil(t, resp.GetError())

	time.Sleep(time.Millisecond * 5)

	// register again with the same username
	gotham.WriteFrame(bufw, msg)
	bufw.Flush()

	res, _ = gotham.ReadFrame(bufr)
	proto.Unmarshal(res.Data(), &resp)
	assert.Equal(t, MsgUserAlreadyExisted.GetDescription(), resp.GetError().GetDescription())

	// login with username and password we just registed.
	lmsg := &pbs.Login{
		Username: u,
		Password: "password!",
	}
	gotham.WriteFrame(bufw, lmsg)
	bufw.Flush()

	res, _ = gotham.ReadFrame(bufr)
	var lresp pbs.LoginAck
	proto.Unmarshal(res.Data(), &lresp)
	assert.Nil(t, lresp.GetError())
}
