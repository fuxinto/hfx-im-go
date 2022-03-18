package handler

import (
	"HIMGo/pkg/response"
	"HIMGo/service/user/api/internal/logic"
	"HIMGo/service/user/api/internal/svc"
	"HIMGo/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp)
		}
	}
}
