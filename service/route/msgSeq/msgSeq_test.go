package msgSeq

import (
	"testing"
)

func Test(t *testing.T) {
	GetMsgSeq("a4b94969-2ba7-420f-851a-03dbf1f3f765", 15)
	//sessionIdInt := hash.Hash(transform.StrTobytes("a4b94969-2ba7-420f-851a-03dbf1f3f765"))
	//
	//lowBits := (sessionIdInt & 0xffff)
	//fmt.Printf("%063b\n", lowBits)
	//fmt.Printf("%b\n", sessionIdInt)
}

func Benchmark_Hello(b *testing.B) {

	for i := 0; i < 4096; i++ {
		GetMsgSeq("a4b94969-2ba7-420f-851a-03dbf1f3f765", 15)
	}
}

func BenchmarkTest(b *testing.B) {

}
