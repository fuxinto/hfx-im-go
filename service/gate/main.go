package main

import (
	"gate/kitex_gen/gatek"
	"log"
)

func main() {
	svr := gatek.NewServer(new(GateImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
