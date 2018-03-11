package notify_saver_service

import (
	"log"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun2/servises/event_service"
	"ms/sun2/shared/config"
	"ms/sun2/shared/x"
)

var newNotify = make(chan x.Notify, 1000)

func saveNewNotifyes() {
	for act := range newNotify {
		if act.Murmur64Hash == 0 || act.NotifyId == 0 || act.ActorUserId == 0 {
			if config.IS_DEBUG {
				log.Panic("x.Notify is not valid to save: ", act)
			}
		}
		act.Save(base.DB)
	}
}

func listernAndSaverNotifys() {
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
	x.NewNotify_Deleter().
		PostId_Eq(e.PostId).
		Delete(base.DB)
}

func on_Liked_Post(e event_service.GeneralEvent) {
	if e.Post == nil {
		return
	}
	act := x.Notify{
		NotifyId:       e.EventId,
		ForUserId:      e.Post.PostId,
		ActorUserId:    e.ByUserId,
		NotifyTypeEnum: int(x.NotifyEnum_NOTIFY_POST_LIKED),
		PostId:         e.PostId,
		PeerUserId:     0,
		CommentId:      0,
		Murmur64Hash:   e.Murmur64Hash,
		CreatedTime:    helper.TimeNow(),
	}
	newNotify <- act
}

func on_UnLiked_Post(e event_service.GeneralEvent) {
	x.NewNotify_Deleter().
		Murmur64Hash_Eq(e.Murmur64Hash).
		Delete(base.DB)
}

func on_Commented_Post(e event_service.GeneralEvent) {
	if e.Comment == nil || e.Post == nil {
		return
	}
	not := x.Notify{
		NotifyId:       e.EventId,
		ForUserId:      e.Post.PostId,
		ActorUserId:    e.ByUserId,
		NotifyTypeEnum: int(x.NotifyEnum_NOTIFY_POST_COMMENTED),
		PostId:         e.Post.PostId,
		PeerUserId:     0,
		CommentId:      e.Comment.CommentId,
		Murmur64Hash:   e.Event.Murmur64Hash,
		CreatedTime:    helper.TimeNow(),
	}
	newNotify <- not
}

func on_UnCommented(e event_service.GeneralEvent) {
	x.NewNotify_Deleter().
		Murmur64Hash_Eq(e.Murmur64Hash).
		Delete(base.DB)
}

func on_Followed(e event_service.GeneralEvent) {
	not := x.Notify{
		NotifyId:       e.EventId,
		ForUserId:      e.PeerUserId,
		ActorUserId:    e.ByUserId,
		NotifyTypeEnum: int(x.NotifyEnum_NOTIFY_FOLLOWED_YOU),
		PostId:         0,
		PeerUserId:     0,
		CommentId:      0,
		Murmur64Hash:   e.Event.Murmur64Hash,
		CreatedTime:    helper.TimeNow(),
	}
	newNotify <- not
}

func on_UnFollwed(e event_service.GeneralEvent) {
	x.NewNotify_Deleter().
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
