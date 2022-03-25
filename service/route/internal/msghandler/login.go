package msghandler

import (
	"HIMGo/pkg/pb"
	"HIMGo/service/route/routeClient"
	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginHandler struct{}

func (LoginHandler) Handler(body []byte) (*routeClient.MessagePushReply, error) {
	signin := pb.LoginReq{}
	err := proto.Unmarshal(body, &signin)
	if err != nil {
		logx.Error("登录请求解析失败")
	}
	return &routeClient.MessagePushReply{}, nil
	//ack := pb.LoginAck{}
	//ack.Code = int32(200)
	//ack.UserId = "str"
	//ack.Msg = "登录成功"

}
