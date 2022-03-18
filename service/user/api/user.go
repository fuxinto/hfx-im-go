package main

import (
	"HIMGo/pkg/fxerror"
	"HIMGo/pkg/response"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"HIMGo/service/user/api/internal/config"
	"HIMGo/service/user/api/internal/handler"
	"HIMGo/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf,
		rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
			if err != nil {
				fmt.Printf(err.Error())
			}
			var body response.Body
			body.Code = 2000
			body.Msg = "认证失败请重新登录"
			httpx.OkJson(w, body)
		}))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *fxerror.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
