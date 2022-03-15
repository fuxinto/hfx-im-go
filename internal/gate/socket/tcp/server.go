package tcp

import (
	"HFXIM/service/gate/him"
	"bufio"
	"net"
)

// Server is a websocket implement of the Server
type Upgrader struct {
}

// NewServer NewServer
func NewServer(listen string, options ...him.ServerOption) him.Server {
	return him.NewServer(listen, new(Upgrader), options...)
}

func (u *Upgrader) Name() string {
	return "tcp.Server"
}

func (u *Upgrader) Upgrade(rawconn net.Conn, rd *bufio.Reader, wr *bufio.Writer) (him.Conn, error) {
	conn := NewConnWithRW(rawconn, rd, wr)
	return conn, nil
}
