package transform

import (
	"strconv"
	"unsafe"
)

func StrTobytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func BytesTostr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//string 转int
func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}
