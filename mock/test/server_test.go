package test

import (
	"HFXIM/mock"
	"testing"
)

func Test(t *testing.T) {
	server := &mock.ServerDemo{}
	server.Start("tcp", ":8999")
}
