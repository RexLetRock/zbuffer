package bench

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/sirupsen/logrus"

	"github.com/RexLetRock/zbuffer/buffer"
	"github.com/RexLetRock/zlib/zbench"
	"github.com/RexLetRock/zlib/zcount"
)

const cRun = 100_000_000
const cCpu = 10
const cMsg = "How are you today ?"
const cSplit = "|||"

var countAll zcount.Counter

func Bench() {
	handle := func(data []byte, cellName int64) {
		a := strings.Split(string(data), cSplit)
		countAll.Add(int64(len(a) - 1))
	}

	zbuffer := buffer.ZBufferCreate(handle)
	warn("==== ZBUFFER ===\n")
	warn("WRITE ---msg---> BUFFER ---msg---> READER < " + cMsg + " >")
	warnf("Buffer size: %T, %d\n", zbuffer, unsafe.Sizeof(*zbuffer))

	zbench.Run(cRun, cCpu, func(i, j int) {
		zbuffer.Write([]byte(cMsg + cSplit))
	})

	zbench.Run(cRun, cCpu, func(i, j int) {
		zbuffer.Write([]byte(cMsg + cSplit))
	})

	zbench.Run(cRun, cCpu, func(i, j int) {
		zbuffer.Write([]byte(cMsg + cSplit))
	})

	zbench.Run(cRun, cCpu, func(i, j int) {
		zbuffer.Write([]byte(cMsg + cSplit))
	})

	zbench.Run(cRun, cCpu, func(i, j int) {
		zbuffer.Write([]byte(cMsg + cSplit))
	})

	time.Sleep(time.Second)
	warnf("CountAll %v \n", Commaize(int64(countAll.Value())))
}

var warn = logrus.Warn
var warnf = logrus.Warnf
var skip = func() {}

func Commaize(n int64) string {
	s1, s2 := fmt.Sprintf("%d", n), ""
	for i, j := len(s1)-1, 0; i >= 0; i, j = i-1, j+1 {
		if j%3 == 0 && j != 0 {
			s2 = "," + s2
		}
		s2 = string(s1[i]) + s2
	}
	return s2
}

// Count32 fast count - fast get
type Count32 int32

func Count32Create() *Count32 {
	return new(Count32)
}

func (c *Count32) IncMaxInt(i int32) int {
	return int(c.IncMax(i))
}

func (c *Count32) IncMax(i int32) int32 {
	a := atomic.AddInt32((*int32)(c), 1)
	if a < i-1 {
		return a
	} else {
		atomic.StoreInt32((*int32)(c), 0)
		return 0
	}
}

func (c *Count32) Inc() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *Count32) Get() int32 {
	return atomic.LoadInt32((*int32)(c))
}
