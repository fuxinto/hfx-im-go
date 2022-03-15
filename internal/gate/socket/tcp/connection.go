package tcp

import (
	"HFXIM/common/helper/endian"
	"HFXIM/service/gate/him"
	"bufio"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net"
)

// Frame Frame
type Frame struct {
	OpCode  him.OpCode
	Payload []byte
}

// SetOpCode SetOpCode
func (f *Frame) SetOpCode(code him.OpCode) {
	f.OpCode = code
}

// GetOpCode GetOpCode
func (f *Frame) GetOpCode() him.OpCode {
	return f.OpCode
}

// SetPayload SetPayload
func (f *Frame) SetPayload(payload []byte) {
	f.Payload = payload
}

// GetPayload GetPayload
func (f *Frame) GetPayload() []byte {
	return f.Payload
}

// TcpConn Conn
type TcpConn struct {
	net.Conn
	rd *bufio.Reader
	wr *bufio.Writer
}

// NewConn NewConn

func NewConn(conn net.Conn) him.Conn {
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
func (c *TcpConn) ReadFrame() (him.Frame, error) {

	opcode, err := endian.ReadUint16(c.rd)
	if err != nil {
		return nil, err
	}
	payload, err := endian.ReadBytes(c.rd)
	if err != nil {
		return nil, err
	}
	code := him.OpCode(opcode)
	//if code == him.OpPong {
	//	logx.Info("收到心跳回应")
	//}
	if code == him.OpPing {
		logx.Info("收到心跳链接")
	}
	return &Frame{
		OpCode:  code,
		Payload: payload,
	}, nil
}

// WriteFrame WriteFrame
func (c *TcpConn) WriteFrame(code him.OpCode, payload []byte) error {
	return WriteFrame(c.wr, code, payload)
}

// Flush Flush
func (c *TcpConn) Flush() error {
	return c.wr.Flush()
}

// WriteFrame write a frame to w
func WriteFrame(w io.Writer, code him.OpCode, payload []byte) error {
	if err := endian.WriteUint16(w, uint16(code)); err != nil {
		return err
	}
	if err := endian.WriteBytes(w, payload); err != nil {
		return err
	}
	return nil
}
