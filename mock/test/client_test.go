package test

import (
	"HIMGo/mock"
	"fmt"
	"testing"
	"time"
)

func Test_client(t *testing.T) {
	cout := 10000
	for i := 0; i < cout; i++ {
		server := &mock.ClientDemo{}
		str := fmt.Sprintf("text:%v", i)
		go server.Start(str, "tcp", "124.71.100.133:8090")
		//go server.Start(str, "ws", "ws://127.0.0.1:8090/ws")
		time.Sleep(time.Millisecond * 2)
	}
	select {}
}
