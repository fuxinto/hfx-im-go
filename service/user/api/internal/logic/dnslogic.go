package logic

import (
	"context"

	"HIMGo/service/user/api/internal/svc"
	"HIMGo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DnsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDnsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DnsLogic {
	return &DnsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DnsLogic) Dns() (resp *types.UserDnsResponse, err error) {
	// todo: add your logic here and delete this line
	dns := types.Dns{
		Host: "127.0.0.1",
		Port: 8090,
	}
	ary := []types.Dns{dns}
	resp = &types.UserDnsResponse{Dns: ary}
	return resp, nil
}
