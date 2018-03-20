package home_fanout_service

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun2/servises/event_service"
	"ms/sun2/servises/mem_user_service"
	"ms/sun2/shared/x"
	"time"
)

var saveHomeFans = make(chan x.HomeFanout, 100000)

func saveHomeFansLooper() {
	tick := time.Tick(time.Millisecond * 200)
	var arr []x.HomeFanout
	for {
		select {
		case fan := <-saveHomeFans:
			arr = append(arr, fan)
		case <-tick:
			x.MassReplace_HomeFanout(arr, base.DB)
			arr = []x.HomeFanout{}
		}
	}
}

func listernAndSaverNotifys() {
	for {
		subParam := event_service.SubParam{
			Added_Post_Event:     true,
			Deleted_Post_Event:   true,
			Followed_User_Event:  true,
			UnFollwed_User_Event: true,
		}

		sub := event_service.NewSub(subParam)

		for {
			var e event_service.GeneralEvent
			select {
			case e = <-sub.Added_Post_Event:
				on_Added_Post(e)
			case e = <-sub.Deleted_Post_Event:
				on_Deleted_Post(e)
			case e = <-sub.Followed_User_Event:
				on_Followed(e)
			case e = <-sub.UnFollwed_User_Event:
				on_UnFollwed(e)
			}
		}
	}
}

func on_Added_Post(e event_service.GeneralEvent) {
	if e.Post == nil {
		return
	}
	mu, ok := mem_user_service.GetForUser(e.ByUserId)
	if !ok {
		return
	}

	followers := mu.GetFollowers()
	for _, uid := range followers {
		fan := x.HomeFanout{
			OrderId:    helper.NanoRowIdAtSec(e.Post.CreatedTime),
			ForUserId:  uid,
			PostId:     e.PostId,
			PostUserId: e.Post.UserId,
		}
		saveHomeFans <- fan
	}

}

func on_Deleted_Post(e event_service.GeneralEvent) {
	if e.PostId < 1 {
		return
	}
	x.NewHomeFanout_Deleter().PostId_Eq(e.PostId).Delete(base.DB)
}

func on_Followed(e event_service.GeneralEvent) {
	mu, ok := mem_user_service.GetForUser(e.ByUserId)
	if !ok {
		return
	}
	posts := mu.GetLastPosts()
	for _, p := range posts {
		if p == nil {
			continue
		}
		fan := x.HomeFanout{
			OrderId:    helper.NanoRowIdAtSec(p.CreatedTime),
			ForUserId:  e.ByUserId,
			PostId:     p.PostId,
			PostUserId: p.UserId,
		}
		saveHomeFans <- fan
	}

}

func on_UnFollwed(e event_service.GeneralEvent) {
	x.NewHomeFanout_Deleter().
		ForUserId_Eq(e.ByUserId).
		PostUserId_Eq(e.PeerUserId).
		Delete(base.DB)
}
