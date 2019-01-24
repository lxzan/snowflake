package snowflake

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"
)

type SnowFlake struct {
	Timestamp uint64
	MachineID uint64
	Index     uint64
}

var snow *SnowFlake

func Initialize(machineID uint64) {
	snow = &SnowFlake{
		Timestamp: uint64(time.Now().Unix() << 30),
		MachineID: machineID << 20,
		Index:     0,
	}
	go timer()
}

func timer() {
	var ticker = time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		var t = uint64(time.Now().In(time.UTC).Unix() << 30)
		atomic.StoreUint64(&snow.Timestamp, t)
		atomic.StoreUint64(&snow.Index, 0)
	}
}

func NextID() uint64 {
	var i = atomic.AddUint64(&snow.Index, 1)
	return snow.Timestamp + snow.MachineID + i
}

func Decode(id uint64) *SnowFlake {
	var b = fmt.Sprintf("%064b", id)
	var t, _ = strconv.ParseInt(b[0:34], 2, 64)
	var m, _ = strconv.ParseInt(b[34:44], 2, 64)
	var i, _ = strconv.ParseInt(b[44:64], 2, 64)
	return &SnowFlake{
		Timestamp: uint64(t),
		MachineID: uint64(m),
		Index:     uint64(i),
	}
}
