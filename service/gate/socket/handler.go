package socket

import (
	"HIMGo/pkg/pb"
	"HIMGo/service/route/route"
	"context"

	"fmt"

	"time"

	"github.com/golang/protobuf/proto"
	"github.com/segmentio/ksuid"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	MetaKeyApp     = "app"
	MetaKeyAccount = "account"
)

// Handler Handler
type Handler struct {
	ServiceID string
	RouteRpc  route.Route
}

// Accept this connection
func (h *Handler) Accept(conn Conn, timeout time.Duration) (string, Meta, error) {

	// 1. 读取登录包
	_ = conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	payload, err := conn.ReadFrame()
	if err != nil {
		return "", nil, err
	}

	id := generateChannelID(h.ServiceID)

	if err2 := h.rpcPush(id, payload); err2 != nil {
		return "", nil, fmt.Errorf("read Not LoginReq")
	}

	return id, nil, nil
}

func generateChannelID(serviceID string) string {
	return fmt.Sprintf("%s:%s", serviceID, ksuid.New().String())
}

func (h *Handler) rpcPush(id string, data []byte) error {
	req := &route.MessagePushReq{
		ChannelId: id,
		Message:   data,
	}
	if _, err2 := h.RouteRpc.GatePushMsg(context.TODO(), req); err2 != nil {
		return err2
	}
	return nil
}

// Receive default listener
func (h *Handler) Receive(ag Agent, payload []byte) {

	pack := pb.Pack{}
	err := proto.Unmarshal(payload, &pack)
	if err != nil {
		logx.Error("proto解析失败")
	}

	switch pack.Type {
	case pb.PackType_heartbeat:
		logx.Info("收到心跳")
		pkt := &pb.Pack{Type: pb.PackType_heartbeatAck}
		data, err := proto.Marshal(pkt)
		if err != nil {
			logx.Errorf("心跳消息编码", err)
			return
		} else {

		}
		err = ag.Push(data)
		if err != nil {
			logx.Errorf("心跳消息发送", err)
		}
	default:
		err1 := h.rpcPush(ag.ID(), payload)
		if err1 != nil {
			logx.Errorf("rpc push Route消息失败", err)
		}
	}

}

// Disconnect default listener
func (h *Handler) Disconnect(id string) error {
	logx.Infof("disconnect %s", id)
	//logout := pkt.New(wire.CommandLoginSignOut, pkt.WithChannel(id))
	//err := container.Forward(wire.SNLogin, logout)
	//if err != nil {
	//	logx.Errorf("module:handler,id:",id)
	//}

	return nil
}
