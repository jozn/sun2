package post_counter_service

import (
    "time"
    "ms/sun/helper"
)

var newSeen = make(chan int) // PostId once seen
var seenMap = make(map[int]int, 10000)

func AddSeen(PostId int) {
    //fmt.Println("AddSeen: ", PostId)
	newSeen <- PostId
}

func newSeenMap() map[int]int {
	old := seenMap
	seenMap = make(map[int]int, 10000)
	return old
}

func proccedSeenMap() {
    helper.JustRecover()
    tick := time.Tick(time.Second * 5)
    for {
        select {
        case pid := <-newSeen:
            seenMap[pid] += 1
            //fmt.Println("added: ", pid)
        case <-tick:
            oldMap := newSeenMap
            stepperProceedMap(oldMap())
        }
    }
}
