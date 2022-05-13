package msgSeq

import (
	"HIMGo/pkg/transform"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/hash"
)

var mutex sync.Mutex
var currentSeq = 0
var maxMessageSeq = 4095

func getMessageSeq() int {
	mutex.Lock()
	defer mutex.Unlock()
	currentSeq++
	if currentSeq > maxMessageSeq {
		currentSeq = 1
	}
	return currentSeq
}

//sessionId会话id
//sessionType会话类型最大15。
func GetMsgSeq(sessionId string, sessionType int) (string, int64) {
	timestamp := time.Now().UnixNano() / 1e6
	higBits := int(timestamp)
	higBits <<= 12
	higBits |= getMessageSeq()
	higBits <<= 4
	higBits |= sessionType & 0xf
	sessionIdInt := int(hash.Hash(transform.StrTobytes(sessionId)) & 0x3fffff)
	higBits <<= 6
	higBits |= sessionIdInt >> 16
	lowBits := sessionIdInt & 0xffff
	higStr := fmt.Sprintf("%064b", higBits)
	lowStr := fmt.Sprintf("%016b", lowBits)
	var builder strings.Builder
	builder.WriteString(higStr)
	builder.WriteString(lowStr)
	str := builder.String()
	return seqCode(str), timestamp
}

func seqCode(str string) string {
	var balance [16]string
	var index = 0
	var i = 0
	for {
		if i >= 80 {
			break
		}
		balance[index] = str[i : i+5]
		i += 5
		index++
	}
	arry := [32]string{"2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C",
		"D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var builder strings.Builder

	for x, value := range balance {
		builder.WriteString(arry[str2DEC(value)])
		if (x+1)%4 == 0 && x != len(balance)-1 {
			builder.WriteString("-")
		}
	}
	return builder.String()
}

func str2DEC(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return num
}
