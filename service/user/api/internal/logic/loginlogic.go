package logic

import (
	"HIMGo/pkg/fxerror"
	"HIMGo/pkg/jwtx"
	"HIMGo/service/user/api/internal/svc"
	"HIMGo/service/user/api/internal/types"
	"HIMGo/service/user/model"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	model := &model.UserModel{Db: l.svcCtx.Db}
	user, err := model.PasswordLogin(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, user.Uid, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		logx.Error(err)
		return nil, fxerror.NewDefaultError("token生成失败")
	}
	resp = &types.LoginReply{
		Uid:       user.Uid,
		NickName:  user.NickName,
		UserToken: token,
	}
	return resp, nil
}
