package mock

import (
	"HIMGo/service/gate/socket"
	"HIMGo/service/gate/socket/tcp"
	"HIMGo/service/gate/socket/websocket"
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/zeromicro/go-zero/core/logx"
	_ "net/http/pprof"
	"time"
)

type ServerDemo struct{}

func (s *ServerDemo) Start(protocol, addr string) {
	//go func() {
	//	_ = http.ListenAndServe("0.0.0.0:6060", nil)
	//}()

	var srv socket.Server
	if protocol == "ws" {
		srv = websocket.NewServer(addr)
	} else if protocol == "tcp" {
		srv = tcp.NewServer(addr)
	}

	handler := &ServerHandler{}

	srv.SetReadWait(time.Minute)
	srv.SetAcceptor(handler)
	srv.SetMessageListener(handler)
	srv.SetStateListener(handler)

	err := srv.Start()
	if err != nil {
		panic(err)
	}
}

// ServerHandler ServerHandler
type ServerHandler struct {
}

// Accept this connection
func (h *ServerHandler) Accept(conn socket.Conn, timeout time.Duration) (string, socket.Meta, error) {
	// 1. 读取登录包
	_ = conn.SetReadDeadline(time.Now().Add(timeout))
	frame, err := conn.ReadFrame()
	if err != nil {
		return "", nil, err
	}
	//pack := HIM.Pack{}
	//err1 := proto.Unmarshal(frame.GetPayload(), &pack)
	//if err1 != nil {
	//	logx.Error("proto解析失败")
	//
	//}
	//if pack.Type != HIM.Pack_loginReq {
	//	return "", fmt.Errorf("read Not LoginReq")
	//}
	logx.Infof(string(frame.GetPayload()))
	conn.WriteFrame(socket.OpBinary, []byte("登录成功"))
	return ksuid.New().String(), nil, nil
}

// Receive default listener
func (h *ServerHandler) Receive(ag socket.Agent, payload []byte) {
	str := fmt.Sprintf("收到消息：", string(payload))
	_ = ag.Push([]byte(str))

	logx.Info(str)
}

// Disconnect default listener
func (h *ServerHandler) Disconnect(id string) error {
	//logger.Warnf("disconnect %s", id)
	logx.Infof("disconnect %s", id)
	return nil
}
