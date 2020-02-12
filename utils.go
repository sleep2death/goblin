package goblin

import (
	"github.com/golang/protobuf/proto"
	"github.com/sleep2death/gotham"
)

// Abort context, write the final message,
// and close the connection.
func AbortWithMessage(c *gotham.Context, msg proto.Message) {
	c.Write(msg)
	c.Abort()
}
