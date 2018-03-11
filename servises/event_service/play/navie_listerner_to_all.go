package main

import "ms/sun2/servises/event_service"

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

func on_Deleted_Post(event event_service.GeneralEvent) {

}

func on_Liked_Post(event event_service.GeneralEvent) {

}

func on_UnLiked_Post(event event_service.GeneralEvent) {

}

func on_Commented_Post(event event_service.GeneralEvent) {

}

func on_UnCommented(event event_service.GeneralEvent) {

}

func on_Followed(event event_service.GeneralEvent) {

}

func on_UnFollwed(event event_service.GeneralEvent) {

}

func on_Blocked(event event_service.GeneralEvent) {

}

func on_UnBlocked(event event_service.GeneralEvent) {

}

/*func on_Deleted_Post(event event_service.GeneralEvent) {

}
*/
