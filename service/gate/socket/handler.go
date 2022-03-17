package socket

import (
	"HIMGo/pkg/fxlog"
	"HIMGo/pkg/pb"
	"HIMGo/service/gate/write"

	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/segmentio/ksuid"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

const (
	MetaKeyApp     = "app"
	MetaKeyAccount = "account"
)

// Handler Handler
type Handler struct {
	ServiceID string
	AppSecret string
}

// Accept this connection
func (h *Handler) Accept(conn Conn, timeout time.Duration) (string, Meta, error) {

	// 1. 读取登录包
	_ = conn.SetReadDeadline(time.Now().Add(timeout))
	frame, err := conn.ReadFrame()
	if err != nil {
		return "", nil, err
	}
	//logx.Infof("%s", frame.GetPayload())
	pack := pb.Pack{}
	err1 := proto.Unmarshal(frame.GetPayload(), &pack)
	if err1 != nil {
		logx.Error("proto解析失败")
	}
	if pack.Type != pb.Pack_loginReq {
		return "", nil, fmt.Errorf("read Not LoginReq")
	}
	//data, err := HIM.NewFrom(HIM.Pack_loginAck, &HIM.LoginAck{Code: 200, Msg: "登录成功"})
	//if err == nil {
	//	conn.Push(data)
	//}
	return ksuid.New().String(), nil, nil
}

// Receive default listener
func (h *Handler) Receive(ag Agent, payload []byte) {

	//buf := bytes.NewBuffer(payload)
	//packet, err := pkt.Read(buf)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//if basicPkt, ok := packet.(*pkt.BasicPkt); ok {
	//	if basicPkt.Code == pkt.CodePing {
	//		_ = ag.Push(pkt.Marshal(&pkt.BasicPkt{Code: pkt.CodePong}))
	//	}
	//	return
	//}
	//if logicPkt, ok := packet.(*pkt.LogicPkt); ok {
	//	logicPkt.ChannelId = ag.ID()
	//
	//	messageInTotal.WithLabelValues(h.ServiceID, wire.SNTGateway, logicPkt.Command).Inc()
	//	messageInFlowBytes.WithLabelValues(h.ServiceID, wire.SNTGateway, logicPkt.Command).Add(float64(len(payload)))
	//
	//	// 把meta注入到header中
	//	if ag.GetMeta() != nil {
	//		logicPkt.AddStringMeta(MetaKeyApp, ag.GetMeta()[MetaKeyApp])
	//		logicPkt.AddStringMeta(MetaKeyAccount, ag.GetMeta()[MetaKeyAccount])
	//	}
	//
	//	err = container.Forward(logicPkt.ServiceName(), logicPkt)
	//	if err != nil {
	//		logger.WithFields(logger.Fields{
	//			"module": "handler",
	//			"id":     ag.ID(),
	//			"cmd":    logicPkt.Command,
	//			"dest":   logicPkt.Dest,
	//		}).Error(err)
	//	}
	//}
}

// Disconnect default listener
func (h *Handler) Disconnect(id string) error {
	fxlog.Infof("disconnect %s", id)
	//logout := pkt.New(wire.CommandLoginSignOut, pkt.WithChannel(id))
	//err := container.Forward(wire.SNLogin, logout)
	//if err != nil {
	//	logx.Errorf("module:handler,id:",id)
	//}

	return nil
}

func generateChannelID(serviceID, account string) string {
	return fmt.Sprintf("%s_%s_%d", serviceID, account, wire.Seq.Next())
}
