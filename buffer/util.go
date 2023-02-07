package buffer

import (
	"time"

	"github.com/RexLetRock/zlib/zgoid"
	"github.com/sirupsen/logrus"
)

const c1024 = 1024            // For fashion 1024
const cBuffSize = 100 * c1024 // Size of buffer 100K
const cCellSize = 500         // Number of cell for cpu use, per cell for cpu

const cTimeLockSleep = 10 * time.Millisecond // Time to sleep before recheck flush
const cTimeToFlush = 300 * time.Millisecond  // 300ms is good for not racing - Time to flush when there is not new data in long time
const cTimeToFlushExit = 1000 * time.Millisecond

var warn = logrus.Warn
var warnf = logrus.Warnf
var skip = func() {}

// Get ID of goroutine
func getGID() int64 {
	return zgoid.Get() % cCellSize
}
