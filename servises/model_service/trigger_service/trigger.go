package trigger_service

import (
	"fmt"
	"ms/sun/shared/x"
	"ms/sun_old/base"
)

type postTrig int

func (postTrig) OnInsert(ins []int) {
	fmt.Println("OnInsert postTrig", ins)
}

func (postTrig) OnUpdate(ins []int) {
	fmt.Println("OnUpdate postTrig", ins)
}

func (postTrig) OnDelete(ins []int) {
	fmt.Println("OnDelete postTrig", ins)
	x.NewHomeFanout_Deleter().PostId_In(ins).Delete(base.DB)
	x.NewAction_Deleter().PostId_In(ins).Delete(base.DB)
}

//////
type actionTrig int

func (actionTrig) OnInsert(ins []int) {
	fmt.Println("OnInsert actionTrig", ins)
}

func (actionTrig) OnUpdate(ins []int) {
	fmt.Println("OnUpdate actionTrig", ins)
}

func (actionTrig) OnDelete(ins []int) {
	fmt.Println("OnDelete actionTrig", ins)
	x.NewActionFanout_Deleter().ActionId_In(ins).Delete(base.DB)
}

var s = x.TriggerModelListener{
	Action: actionTrig(1),
	Post:   postTrig(1),
}

func init() {
	listen()
}
func listen() {
	x.ActivateTrigger = true

	x.ArrTriggerListeners = append(x.ArrTriggerListeners, s)
}
