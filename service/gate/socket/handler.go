package socket

import (
	"HIMGo/service/route/route"
	"context"

	"fmt"

	"time"

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
func (h *Handler) Accept(conn Conn, timeout time.Duration) (string, []byte, error) {

	// 1. 读取登录包
	_ = conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	payload, err := conn.ReadFrame()
	if err != nil {
		return "", nil, err
	}

	id := generateChannelID(h.ServiceID)
	data, err2 := h.rpcPush(id, payload)
	return id, data, err2
}

func generateChannelID(serviceID string) string {
	return fmt.Sprintf("%s:%s", serviceID, ksuid.New().String())
}

func (h *Handler) rpcPush(id string, data []byte) ([]byte, error) {
	req := &route.GateReq{
		ChannelId: id,
		Message:   data,
	}
	reply, err := h.RouteRpc.GatePushMsg(context.TODO(), req)
	return reply.Body, err
}

// Receive default listener
func (h *Handler) Receive(ag Agent, payload []byte) {

	data, err := h.rpcPush(ag.ID(), payload)
	if err != nil {
		logx.Errorf("rpc push Route消息失败", err)
	} else if data != nil {
		ag.Push(data)
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
