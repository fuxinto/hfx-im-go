package tcp

import (
	"HIMGo/service/gate/socket"
	"HIMGo/service/gate/socket/endian"
	"bufio"
	"io"
	"net"
)

// TcpConn Conn
type TcpConn struct {
	net.Conn
	rd *bufio.Reader
	wr *bufio.Writer
}

// NewConn NewConn
func NewConn(conn net.Conn) socket.Conn {
	return &TcpConn{
		Conn: conn,
		rd:   bufio.NewReaderSize(conn, 4096),
		wr:   bufio.NewWriterSize(conn, 1024),
	}
}

func NewConnWithRW(conn net.Conn, rd *bufio.Reader, wr *bufio.Writer) *TcpConn {
	return &TcpConn{
		Conn: conn,
		rd:   rd,
		wr:   wr,
	}
}

// ReadFrame ReadFrame
func (c *TcpConn) ReadFrame() ([]byte, error) {
	payload, err := endian.ReadBytes(c.rd)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// WriteFrame WriteFrame
func (c *TcpConn) WriteFrame(payload []byte) error {
	return WriteFrame(c.wr, payload)
}

// Flush Flush
func (c *TcpConn) Flush() error {
	return c.wr.Flush()
}

// WriteFrame write a frame to w
func WriteFrame(w io.Writer, payload []byte) error {
	if err := endian.WriteBytes(w, payload); err != nil {
		return err
	}
	return nil
}
