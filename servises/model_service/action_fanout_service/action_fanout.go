package action_fanout_service

import (
	"ms/sun_old/base"
	"ms/sun/shared/helper"
	"ms/sun/servises/event_service"
	"ms/sun/servises/mem_user_service"
	"ms/sun/shared/x"
	"time"
)

var saveActionFans = make(chan x.ActionFanout, 100000)

func saveActionFansLooper() {
	tick := time.Tick(time.Millisecond * 200)
	var arr []x.ActionFanout
	for {
		select {
		case fan := <-saveActionFans:
			arr = append(arr, fan)
		case <-tick:
			x.MassReplace_ActionFanout(arr, base.DB)
			arr = []x.ActionFanout{}
		}
	}
}

func listernAndSaverEvents() {
	for {
		subParam := event_service.SubParam{
			Added_User_Action:   true,
			Deleted_User_Action: true,
		}

		sub := event_service.NewSub(subParam)

		for {
			var e event_service.GeneralEvent
			select {
			case e = <-sub.Added_Post_Event:
				on_Added_Action(e)
			case e = <-sub.Deleted_Post_Event:
				on_Deleted_Action(e)
			}
		}
	}
}

func on_Added_Action(e event_service.GeneralEvent) {
	mu, ok := mem_user_service.GetForUser(e.ByUserId)
	if !ok {
		return
	}

	followers := mu.GetFollowers()
	for _, uid := range followers {
		fan := x.ActionFanout{
			OrderId:     helper.NanoRowIdAtSec(e.Post.CreatedTime),
			ForUserId:   uid,
			ActionId:    e.ActionId,
			ActorUserId: e.ByUserId,
		}
		saveActionFans <- fan
	}
}

func on_Deleted_Action(e event_service.GeneralEvent) {

}
