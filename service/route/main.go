package main

import (
	"log"
	routeRpc "route/kitex_gen/routeRpc/routerpc"
)

func main() {
	svr := routeRpc.NewServer(new(RouteRpcImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
