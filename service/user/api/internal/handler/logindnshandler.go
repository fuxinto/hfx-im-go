package handler

import (
	"net/http"

	"HIMGo/service/user/api/internal/logic"
	"HIMGo/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginDnsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLoginDnsLogic(r.Context(), svcCtx)
		resp, err := l.LoginDns()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
