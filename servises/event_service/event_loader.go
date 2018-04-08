package event_service

import (
	"ms/sun_old/base"
	"ms/sun/servises/log_service"
	"ms/sun/shared/config"
	"ms/sun/shared/x"
	"time"
)

//todo: change sub.Deleted_Post_Event <- gEvent # to a func with an select timeout in case event not procceed
func NewSub(param SubParam) Sub {
	last := 0
	sub := newSub()
	go func() {
		for {
			time.Sleep(time.Millisecond * 1000)
			selector := x.NewEvent_Selector().OrderBy_EventId_Asc()
			if last > 0 {
				selector.EventId_GT(last)
			}
			events, err := selector.GetRows2(base.DB) //older Events are in zero index to ...
			if err != nil || len(events) == 0 {
				continue
			}
			last = events[len(events)-1].EventId
			genEvents := preloadevents(events)

			for _, gEvent := range genEvents {
				switch EventType(gEvent.Event.EventType) {
				case ADDED_POST_EVENT:
					if param.Added_Post_Event {
						sub.Added_Post_Event <- gEvent
					}
				case DELETED_POST_EVENT:
					if param.Deleted_Post_Event {
						sub.Deleted_Post_Event <- gEvent
					}
				case LIKED_POST_EVENT:
					if param.Liked_Post_Event {
						sub.Liked_Post_Event <- gEvent
					}
				case UNLIKED_POST_EVENT:
					if param.UnLiked_Post_Event {
						sub.UnLiked_Post_Event <- gEvent
					}
				case COMMENTED_POST_EVENT:
					if param.Commented_Post_Event {
						sub.Commented_Post_Event <- gEvent
					}
				case UNCOMMENTED_POST_EVENT:
					if param.UnCommented_Post_Event {
						sub.UnCommented_Post_Event <- gEvent
					}
				case FOLLOWED_USER_EVENT:
					if param.Followed_User_Event {
						sub.Followed_User_Event <- gEvent
					}
				case UNFOLLOWED_USER_EVENT:
					if param.UnFollwed_User_Event {
						sub.UnFollwed_User_Event <- gEvent
					}
				case BLOCKED_USER_EVENT:
					if param.Blocked_User_Event {
						sub.Blocked_User_Event <- gEvent
					}
				case UNBLOCKED_USER_EVENT:
					if param.UnBlocked_User_Event {
						sub.UnBlocked_User_Event <- gEvent
					}
				case ADDED_USER_ACTION:
					if param.Added_User_Action {
						sub.Added_User_Action <- gEvent
					}
				case DELETED_USER_ACTION:
					if param.Deleted_User_Action {
						sub.Deleted_User_Action <- gEvent
					}
				default:
					if config.IS_DEBUG {
						log_service.BadErrorf("event_service NewSub unknown EventId: %d", gEvent.Event.EventType)
					}
				}
			}
		}
	}()
	return sub
}
