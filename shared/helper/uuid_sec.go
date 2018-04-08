package helper

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//sec 10digit + 9 rand
func nanoRowIdAtSec_dep(sec int) int {
	const DIGIT = "0123456789"
	//const DIGIT = "01"
	rndStr := ""
	for i := 0; i < 9; i++ {
		rndStr += string(DIGIT[rand.Intn(len(DIGIT))])
	}
	out := fmt.Sprintf("%d%s", sec, rndStr)
	o := StrToInt(out, TimeNowNano())
	return o
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var seqForRow = rand.Intn(1e9)
var seqForRowMx = sync.RWMutex{}

func NanoRowIdAtSec(timeSec int) int {
	seqForRowMx.Lock()
	defer seqForRowMx.Unlock()
	if seqForRow >= 1e9 {
		seqForRow = rand.Intn(1e8) //100million on resets
	}
	seqForRow += 1

	if timeSec < 1 {
		timeSec = TimeNow()
	}
	timeSec = timeSec * 1e9
	timeSec = timeSec + seqForRow

	return int(timeSec)
}

func NanoRowIdSeq() int {
	timeSec := time.Now().Unix()
	timeSec = timeSec * 1e9
	timeSec = timeSec + int64(rand.Intn(1e9))

	return int(timeSec)
}
