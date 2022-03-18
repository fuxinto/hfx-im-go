package test

import (
	"HIMGo/mock"
	"testing"
)

func Test(t *testing.T) {
	server := &mock.ServerDemo{}
	server.Start("tcp", ":8999")
}
