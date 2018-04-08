package action_saver_service

import (
	"log"
	"ms/sun_old/base"
	"ms/sun_old/helper"
	"ms/sun/servises/event_service"
	"ms/sun/shared/config"
	"ms/sun/shared/x"
)

var newActions = make(chan x.Action, 1000)

func saveNewActions() {
	for act := range newActions {
		if act.Murmur64Hash == 0 || act.ActionId == 0 || act.ActorUserId == 0 {
			if config.IS_DEBUG {
				log.Panic("x.Action is not valid to save: ", act)
			}
		}
		ev := x.Event{
			EventId:      helper.NanoRowIdSeq(),
			EventType:    int(event_service.ADDED_USER_ACTION),
			ByUserId:     act.ActorUserId,
			PeerUserId:   act.PeerUserId,
			PostId:       act.PostId,
			CommentId:    act.CommentId,
			ActionId:     act.ActionId,
			Murmur64Hash: act.Murmur64Hash,
		}
		event_service.SaveEvent(event_service.ADDED_USER_ACTION, ev)
		act.Save(base.DB)
	}
}

func listernAndSaverActions() {
	for {
		subParam := event_service.SubParam{
			Added_Post_Event:       true,
			Deleted_Post_Event:     true,
			Liked_Post_Event:       true,
			UnLiked_Post_Event:     true,
			Commented_Post_Event:   true,
			UnCommented_Post_Event: true,
			Followed_User_Event:    true,
			UnFollwed_User_Event:   true,
			Blocked_User_Event:     true,
			UnBlocked_User_Event:   true,
		}

		sub := event_service.NewSub(subParam)

		for {
			var e event_service.GeneralEvent
			select {
			case e = <-sub.Deleted_Post_Event:

			case e = <-sub.Deleted_Post_Event:
				on_Deleted_Post(e)
			case e = <-sub.Liked_Post_Event:
				on_Liked_Post(e)
			case e = <-sub.UnLiked_Post_Event:
				on_UnLiked_Post(e)
			case e = <-sub.Commented_Post_Event:
				on_Commented_Post(e)
			case e = <-sub.UnCommented_Post_Event:
				on_UnCommented(e)
			case e = <-sub.Followed_User_Event:
				on_Followed(e)
			case e = <-sub.UnFollwed_User_Event:
				on_UnFollwed(e)
			case e = <-sub.Blocked_User_Event:
				on_Blocked(e)
			case e = <-sub.UnBlocked_User_Event:
				on_UnBlocked(e)
			}
		}

	}
}

func on_Deleted_Post(e event_service.GeneralEvent) {
	x.NewAction_Deleter().
		ActorUserId_Eq(e.ByUserId).
		PostId_Eq(e.PostId).
		Delete(base.DB)
}

func on_Liked_Post(e event_service.GeneralEvent) {
	if e.Post == nil {
		//return
	}
	act := x.Action{
		ActionId:       e.EventId,
		ActorUserId:    e.ByUserId,
		ActionTypeEnum: int(x.ActionEnum_ACTION_POST_LIKED),
		PostId:         e.PostId,
		PeerUserId:     e.PeerUserId,
		CommentId:      0,
		Murmur64Hash:   e.Murmur64Hash,
		CreatedTime:    helper.TimeNow(),
	}
	newActions <- act
}

func on_UnLiked_Post(e event_service.GeneralEvent) {
	x.NewAction_Deleter().
		ActorUserId_Eq(e.ByUserId).
		Murmur64Hash_Eq(e.Murmur64Hash).
		Delete(base.DB)
}

func on_Commented_Post(e event_service.GeneralEvent) {
	if e.Comment == nil || e.Post == nil {
		return
	}
	act := x.Action{
		ActionId:       e.EventId,
		ActorUserId:    e.Comment.UserId,
		ActionTypeEnum: int(x.ActionEnum_ACTION_POST_COMMENTED),
		PostId:         e.Post.PostId,
		PeerUserId:     e.Post.UserId,
		CommentId:      e.Comment.CommentId,
		Murmur64Hash:   e.Event.Murmur64Hash,
		CreatedTime:    helper.TimeNow(),
	}
	newActions <- act
}

func on_UnCommented(e event_service.GeneralEvent) {
	x.NewAction_Deleter().
		ActorUserId_Eq(e.ByUserId).
		Murmur64Hash_Eq(e.Murmur64Hash).
		Delete(base.DB)
}

func on_Followed(e event_service.GeneralEvent) {
	act := x.Action{
		ActionId:       e.EventId,
		ActorUserId:    e.ByUserId,
		ActionTypeEnum: int(x.ActionEnum_ACTION_FOLLOWED_USER),
		PostId:         0,
		PeerUserId:     e.PeerUserId,
		CommentId:      0,
		Murmur64Hash:   e.Event.Murmur64Hash,
		CreatedTime:    helper.TimeNow(),
	}
	newActions <- act
}

func on_UnFollwed(e event_service.GeneralEvent) {
	x.NewAction_Deleter().
		ActorUserId_Eq(e.ByUserId).
		Murmur64Hash_Eq(e.Murmur64Hash).
		Delete(base.DB)
}

func on_Blocked(e event_service.GeneralEvent) {

}

func on_UnBlocked(e event_service.GeneralEvent) {

}

/*func on_Deleted_Post(e event_service.GeneralEvent) {

}
*/
