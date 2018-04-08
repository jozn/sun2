package main

import (
	"fmt"
	"ms/sun/shared/x"
	"time"
)

type pt int

func (pt) OnInsert(ins []int) {
	fmt.Println("OnInsert", ins)
}

func (pt) OnUpdate(ins []int) {
	fmt.Println("OnUpdate", ins)
}

func (pt) OnDelete(ins []int) {
	fmt.Println("OnDelete", ins)
}

var s = x.TriggerModelListener{
	Post: pt(1),
}

func main() {
	x.ActivateTrigger = true

	x.ArrTriggerListeners = append(x.ArrTriggerListeners, s)

	time.Sleep(time.Minute)
}
