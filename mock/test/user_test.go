package test

import (
	"HIMGo/pkg/transform"
	"crypto/sha256"
	"fmt"
	"testing"
)

func Test_user(t *testing.T) {
	data := []byte("password12333")
	h := sha256.New()
	h.Write(data)
	psd := transform.BytesTostr(h.Sum(nil))
	fmt.Println(psd)
}
