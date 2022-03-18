package handler

import (
	"HIMGo/pkg/response"
	"HIMGo/service/user/api/internal/logic"
	"HIMGo/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp)
		}
	}
}
