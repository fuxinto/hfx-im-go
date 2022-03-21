package logic

import (
	"HIMGo/service/user/model"
	"context"

	"HIMGo/service/user/api/internal/svc"
	"HIMGo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	model := &model.UserModel{Db: l.svcCtx.Db}
	uid := l.ctx.Value("uid").(string)

	user, err := model.GetUserInfo(uid)
	resp = &types.UserInfoResponse{
		Uid:      uid,
		NickName: user.NickName,
		Sex:      0,
	}
	return resp, nil
}
