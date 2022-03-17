package logic

import (
	"context"
	"fmt"

	"HIMGo/service/route/internal/svc"
	"HIMGo/service/route/routeClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *routeClient.LoginReq) (*routeClient.LoginReply, error) {
	// todo: add your logic here and delete this line
	if err := l.svcCtx.Cache.Set(in.Uid, in.ChannelId); err != nil {
		fmt.Println("redisset失败")
	}
	fmt.Println(in.ChannelId, in.Uid)
	return &routeClient.LoginReply{}, nil
}
