package post_counter_service

import (
    "time"
)

var newSeen = make(chan int) // PostId once seen
var seenMap = make(map[int]int, 100000)

func AddSeen(PostId int) {
    //fmt.Println("AddSeen: ", PostId)
	newSeen <- PostId
}

func newSeenMap() map[int]int {
	old := seenMap
	seenMap = make(map[int]int, 100000)
	return old
}

func proccedSeenMap() {
    defer func() {
        err := recover()
        if err!= nil{
            time.Sleep(time.Second)
            go proccedSeenMap()
        }
    }()

    tick := time.Tick(time.Second * 5)
    for {
        select {
        case pid := <-newSeen:
            seenMap[pid] += 1
            //fmt.Println("added: ", pid)
        case <-tick:
            oldMap := newSeenMap()
            stepperProceedMap(oldMap)
        }
    }
}
