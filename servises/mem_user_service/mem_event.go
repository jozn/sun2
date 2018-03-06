package mem_user_service

import "ms/sun2/servises/event_service"

func listen() {
	sub := event_service.SubParam{
		ADDED_POST_EVENT:      true,
		DELETED_POST_EVENT:    true,
		FOLLOWED_USER_EVENT:   true,
		UNFOLLOWED_USER_EVENT: true,
	}
	listener := event_service.NewSub(sub)

}
