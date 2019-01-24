package snowflake

import (
	"fmt"
	"testing"
)

func TestNextID(t *testing.T) {
	Initialize(1)
	println(NextID(), fmt.Sprintf("%v",snow))
}

func TestDecode(t *testing.T) {
	var o = Decode(1662468685074268160)
	println(fmt.Sprintf("%v",o))
}

func BenchmarkKey(b *testing.B) {
	Initialize(1)
	for i := 0; i < b.N; i++ {
		NextID()
	}
}
