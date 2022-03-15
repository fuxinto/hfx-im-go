package main

import (
	"context"
	"gate/kitex_gen/rpc"
	"google.golang.org/protobuf/proto"
)

// GateImpl implements the last service interface defined in the IDL.
type GateImpl struct{}

// RoutePushMsg implements the GateImpl interface.
func (s *GateImpl) RoutePushMsg(ctx context.Context, req *rpc.MessagePushReq) (resp *rpc.MessagePushReply, err error) {
	data := req.Message
	login := &hi
	if err := proto.Unmarshal(data, login), err != nil {

	}
	// TODO: Your code here...
	return
}
