package uuid_play

import (
	//"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"sync"
	"time"
    "ms/sun_old/helper"
)

func main() {
	ch := make(chan int, 1000)
	mk := make(map[int]bool)
	go func() {
		for c := range ch {
			if mk[c] {
				panic(c)
			}
			mk[c] = true
		}
	}()
	go func() {
		for {
			ch <- helper.NanoRowIdAtSec(1400000000)
		}
	}()
	go func() {
		for {
			ch <- helper.NanoRowIdAtSec(1400000000)
		}
	}()
	go func() {
		for {
			ch <- helper.NanoRowIdAtSec(1400000000)
		}
	}()
	for {
		//fmt.Println(genu8Int())
		//fmt.Println(genu8IntSimple())
		//fmt.Println(genu8IntTiming())
		fmt.Println(helper.NanoRowIdAtSec(1400000000), len(mk))
		//genu84()
		time.Sleep(time.Millisecond * 200)
	}
}

func genu8() string {
	timeSec := uint32(time.Now().Unix())
	uid := int(rand.Intn(4 * 2e9))
	var bts [8]byte
	binary.BigEndian.PutUint32(bts[:4], timeSec)
	binary.BigEndian.PutUint32(bts[4:], uint32(uid))
	b := bts
	uuidStr := fmt.Sprintf("%x%x", b[0:4], b[4:8])
	//uuidStr := fmt.Sprintf("%d%d", b[0:4], b[4:8])
	return uuidStr
}

func genu8Int() string {
	timeSec := uint32(time.Now().Unix())
	uid := uint32(rand.Intn(4 * 1e15))
	o := uint64(timeSec)
	o = o << 32
	o = o | (uint64(uid))
	fmt.Printf("%b \n", o)
	fmt.Printf("%b \n", timeSec)
	fmt.Printf("out int: %d \n", o)
	uuidStr := fmt.Sprintf("%d0%08d", timeSec, uid)
	return uuidStr
}

func genu8IntSimple() int {
	timeSec := uint32(time.Now().Unix())
	uid := uint32(rand.Intn(4 * 1e15))
	o := int64(timeSec)
	o = o << 32
	o = o | (int64(uid))
	return int(o)
}

func genu8IntTiming() int {
	timeSec := time.Now().Unix()
	timeSec = timeSec * 1e9
	timeSec = timeSec + int64(rand.Intn(1e3))

	return int(timeSec)
}

func genu84() int64 {
	timeSec := uint32(time.Now().Unix())
	uid := int(rand.Intn(4 * 2e9))
	var bts [8]byte
	binary.BigEndian.PutUint32(bts[:4], timeSec)
	binary.BigEndian.PutUint32(bts[4:], uint32(uid))

	fmt.Println(binary.Varint(bts[:]))
	return int64(0)
}

func init()  {
    rand.Seed(time.Now().UnixNano())
}
var seqForRow = rand.Intn(1e9)
var seqForRowMx = sync.RWMutex{}

func NanoRowIdAtSec(timeSec int) int {
	defer seqForRowMx.Unlock()
	seqForRowMx.Lock()
	if seqForRow >= 1e9 {
		seqForRow = rand.Intn(1e8)
	}
	seqForRow += 1

	timeSec = timeSec * 1e9
	timeSec = timeSec + seqForRow

	return int(timeSec)
}
