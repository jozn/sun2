package event_service

import (
	"ms/sun/base"
	"ms/sun2/servises/log_service"
	"ms/sun2/shared/config"
	"ms/sun2/shared/x"
	"time"
)

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
					if param.ADDED_POST_EVENT {
						sub.ADDED_POST_EVENT <- gEvent
					}
				case DELETED_POST_EVENT:
					if param.DELETED_POST_EVENT {
						sub.DELETED_POST_EVENT <- gEvent
					}
				case LIKED_POST_EVENT:
					if param.LIKED_POST_EVENT {
						sub.LIKED_POST_EVENT <- gEvent
					}
				case UNLIKED_POST_EVENT:
					if param.UNLIKED_POST_EVENT {
						sub.UNLIKED_POST_EVENT <- gEvent
					}
				case COMMENTED_POST_EVENT:
					if param.COMMENTED_POST_EVENT {
						sub.COMMENTED_POST_EVENT <- gEvent
					}
				case UNCOMMENTED_POST_EVENT:
					if param.UNCOMMENTED_POST_EVENT {
						sub.UNCOMMENTED_POST_EVENT <- gEvent
					}
				case FOLLOWED_USER_EVENT:
					if param.FOLLOWED_USER_EVENT {
						sub.FOLLOWED_USER_EVENT <- gEvent
					}
				case UNFOLLOWED_USER_EVENT:
					if param.UNFOLLOWED_USER_EVENT {
						sub.UNFOLLOWED_USER_EVENT <- gEvent
					}
				case BLOCKED_USER_EVENT:
					if param.BLOCKED_USER_EVENT {
						sub.BLOCKED_USER_EVENT <- gEvent
					}
				case UNBLOCKED_USER_EVENT:
					if param.UNBLOCKED_USER_EVENT {
						sub.UNBLOCKED_USER_EVENT <- gEvent
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
