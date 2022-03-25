package tcp

import (
	"HIMGo/service/gate/socket"
	"bufio"
	"net"
)

// Server is a websocket implement of the Server
type Upgrader struct {
}

// NewServer NewServer
func NewServer(listen string, options ...socket.ServerOption) socket.Server {
	return socket.NewServer(listen, new(Upgrader), options...)
}

func (u *Upgrader) Name() string {
	return "tcp.Server"
}

func (u *Upgrader) Upgrade(rawconn net.Conn, rd *bufio.Reader, wr *bufio.Writer) (socket.Conn, error) {
	conn := NewConnWithRW(rawconn, rd, wr)
	return conn, nil
}
