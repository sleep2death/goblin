package handlers

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/sleep2death/gotham"
	"go.uber.org/zap"
)

func Recovery(logger *zap.Logger) gotham.HandlerFunc {
	return func(c *gotham.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				if logger != nil {
					logger.Warn(fmt.Sprintf("%s\n%s", err, c.Request.RemoteAddr()))
				}
				if brokenPipe {
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
				} else {
					c.Write(MsgInternalServerError)
					c.Abort()
				}
			}
		}()
		c.Next()

	}
}

