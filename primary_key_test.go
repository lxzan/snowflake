package snowflake

import (
	"fmt"
	"sync"
	"testing"
)

func TestNextID(t *testing.T) {
	s := NewSnowFlake(1)
	var m = sync.Map{}

	var wg = sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				var id = s.NextID()
				if _, ok := m.Load(id); ok {
					panic(fmt.Sprintf("%v", id))
				}
				m.Store(id, true)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestDecode(t *testing.T) {
	var o = Decode(1662476606067703809)
	fmt.Println(fmt.Sprintf("%v\n", o))
}

func BenchmarkNextID(b *testing.B) {
	s := NewSnowFlake(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.NextID()
	}
}
