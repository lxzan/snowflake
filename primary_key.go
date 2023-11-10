package snowflake

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type SnowFlake struct {
	timestamp uint64
	id        uint64
	index     uint64
	ctx       context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
	once      sync.Once
}

func NewSnowFlake(machineID uint64) *SnowFlake {
	snow := SnowFlake{
		timestamp: uint64(time.Now().Unix() << 30),
		id:        machineID << 20,
		index:     0,
	}

	snow.wg.Add(1)
	snow.ctx, snow.cancel = context.WithCancel(context.Background())

	go snow.timer()

	return &snow
}

func (s *SnowFlake) Stop() {
	s.once.Do(func() {
		s.cancel()
		s.wg.Wait()
	})
}

func (s *SnowFlake) timer() {
	var ticker = time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-s.ctx.Done():
			s.wg.Done()
			return
		case <-ticker.C:
			var t = uint64(time.Now().In(time.UTC).Unix() << 30)
			atomic.StoreUint64(&s.timestamp, t)
			atomic.StoreUint64(&s.index, 0)
		}
	}
}

func (s *SnowFlake) NextID() uint64 {
	return s.timestamp + s.id + atomic.AddUint64(&s.index, 1)
}

// 这里解析可能不在返回 SnowFlake 对象，可能是领  SnowFlake 的内部对对象
func Decode(id uint64) *SnowFlake {
	var b = fmt.Sprintf("%064b", id)
	var t, _ = strconv.ParseInt(b[0:34], 2, 64)
	var m, _ = strconv.ParseInt(b[34:44], 2, 64)
	var i, _ = strconv.ParseInt(b[44:64], 2, 64)
	return &SnowFlake{
		timestamp: uint64(t),
		id:        uint64(m),
		index:     uint64(i),
	}
}
