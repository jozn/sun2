package uuid_play

import (
	"fmt"
	"math/rand"
	"ms/sun/helper"
	"time"
)

func main() {
	collistion()
	o := helper.NanoRowIdAtSec(1400000000)
	os := helper.IntToStr(o)
	fmt.Println(o, len(os))

	for {
		//o := helper.NanoRowIdAtSec(helper.TimeNow())
		o := helper.NanoRowIdAtSec(1400000000)
		os := helper.IntToStr(o)
		fmt.Println(o, len(os))
		time.Sleep(time.Millisecond * 200)
	}
}

func collistion() {
	mp := make(map[int]bool)
	i := 0
	for {
		o := helper.NanoRowIdAtSec(1400000000)
		if mp[o] {
            fmt.Println(i,o)
			panic(o)
		}
		mp[o] = true
		i++
		if i%1000000 == 0 {
			fmt.Println(i,o)
		}
	}
}

//sec 10digit + 9 rand
func NanoRowIdAtSec2(sec int) int {
	const DIGIT = "0123456789"
	rndStr := ""
	for i := 0; i < 9; i++ {
		rndStr += string(DIGIT[rand.Intn(len(DIGIT))])
	}
	out := fmt.Sprintf("%d%s", sec, rndStr)
	o := helper.StrToInt(out, helper.TimeNowNano())
	return o
}
