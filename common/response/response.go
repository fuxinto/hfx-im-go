package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, data interface{}, code int) {
	var body Body
	body.Code = code
	if code == 200 {
		body.Msg = "ok"
		body.Data = data
	} else {
		body.Msg = "fail"
		body.Data = nil
	}
	httpx.OkJson(w, body)
}
