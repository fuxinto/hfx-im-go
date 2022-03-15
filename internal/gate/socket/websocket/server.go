package websocket

import (
	him2 "HFXIM/service/gate/him"
	"bufio"
	"net"

	"github.com/gobwas/ws"
)

// Server is a websocket implement of the Server
type Upgrader struct {
}

// NewServer NewServer
func NewServer(listen string, options ...him2.ServerOption) him2.Server {
	return him2.NewServer(listen, new(Upgrader), options...)
}

func (u *Upgrader) Name() string {
	return "websocket.Server"
}

func (u *Upgrader) Upgrade(rawconn net.Conn, rd *bufio.Reader, wr *bufio.Writer) (him2.Conn, error) {
	_, err := ws.Upgrade(rawconn)
	if err != nil {
		return nil, err
	}
	conn := NewConnWithRW(rawconn, rd, wr)
	return conn, nil
}
