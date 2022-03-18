package logic

import (
	"context"

	"HIMGo/service/user/api/internal/svc"
	"HIMGo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginDnsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginDnsLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginDnsLogic {
	return LoginDnsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginDnsLogic) LoginDns() (resp *types.UserDnsResponse, err error) {
	// todo: add your logic here and delete this line

	return &types.UserDnsResponse{Dns: []string{"124.71.100.133:8081"}}, nil
}
