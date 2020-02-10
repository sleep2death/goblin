package goblin

import (
	"bufio"
	"log"
	"net"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/sleep2death/goblin/pb"
	"github.com/sleep2death/gotham"
)

func TestRegisterHandler(t *testing.T) {
	go Serve()

	time.Sleep(time.Millisecond * 5)
	conn, err := net.Dial("tcp", ":9000")
	if err != nil {
		t.Fatal(err)
	}
	bufw := bufio.NewWriter(conn)
	bufr := bufio.NewReader(conn)

	msg := &pb.Register{
		Username: "username",
		Email:    "username@goblin.com",
		Password: "password!",
	}

	gotham.WriteFrame(bufw, msg)
	bufw.Flush()

	res, _ := gotham.ReadFrame(bufr)

	var resp pb.RegisterAck
	proto.Unmarshal(res.Data, &resp)
	log.Println(resp.GetError().GetMessage())

	time.Sleep(time.Millisecond * 5)
}
