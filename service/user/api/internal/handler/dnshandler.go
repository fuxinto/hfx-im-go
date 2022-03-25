package handler

import (
	"HIMGo/pkg/response"
	"net/http"

	"HIMGo/service/user/api/internal/logic"
	"HIMGo/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DnsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDnsLogic(r.Context(), svcCtx)
		resp, err := l.Dns()
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp)
		}
	}
}
