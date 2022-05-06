package socket

import (
	"context"
	"net"
	"time"
)

const (
	DefaultReadWait  = time.Minute * 4
	DefaultWriteWait = time.Second * 10
	DefaultLoginWait = time.Second * 10
)

const (
	// 定义读取消息的默认goroutine池大小
	DefaultMessageReadPool = 50
	DefaultConnectionPool  = 50
)

// Server 定义了一个tcp/websocket不同协议通用的服务端的接口
type Server interface {

	// SetAcceptor 设置Acceptor
	SetAcceptor(Acceptor)
	//SetMessageListener 设置上行消息监听器
	SetMessageListener(MessageListener)
	//SetStateListener 设置连接状态监听服务
	SetStateListener(StateListener)
	// SetReadWait 设置读超时
	SetReadWait(time.Duration)
	// SetChannelMap 设置Channel管理服务
	SetChannelMap(ChannelMap)

	// Start 用于在内部实现网络端口的监听和接收连接，
	// 并完成一个Channel的初始化过程。
	Start() error
	// Push 消息到指定的Channel中
	//  string channelID
	//  []byte 序列化之后的消息数据
	Push(string, []byte) error
	// Shutdown 服务下线，关闭连接
	Shutdown(context.Context) error
}

// Acceptor 连接接收器
type Acceptor interface {
	// Accept 返回一个握手完成的Channel对象或者一个error。
	// 业务层需要处理不同协议和网络环境下的连接握手协议
	Accept(Conn, time.Duration) (string, []byte, error)
}

// MessageListener 监听消息
type MessageListener interface {
	// 收到消息回调
	Receive(Agent, []byte)
}

// StateListener 状态监听器
type StateListener interface {
	// 连接断开回调
	Disconnect(string) error
}

// type Meta map[string]string

// Agent is interface of client side
type Agent interface {
	ID() string
	Push([]byte) error
	// GetMeta() Meta
}

// Conn Connection
type Conn interface {
	net.Conn
	ReadFrame() ([]byte, error)
	WriteFrame([]byte) error
	Flush() error
}

// Channel is interface of client side
type Channel interface {
	Conn
	Agent
	// Close 关闭连接
	Close() error
	Readloop(lst MessageListener) error
	// SetWriteWait 设置写超时
	SetWriteWait(time.Duration)
	SetReadWait(time.Duration)
	Login() error
}

// Client is interface of client side
type Client interface {
	// connect to server
	Connect(string) error
	// SetDialer 设置拨号处理器
	SetDialer(Dialer)
	Send([]byte) error
	Read() ([]byte, error)
	// Close 关闭
	Close()
}

// Dialer Dialer
type Dialer interface {
	DialAndHandshake(DialerContext) (net.Conn, error)
}

type DialerContext struct {
	Id      string
	Name    string
	Address string
	Timeout time.Duration
}
