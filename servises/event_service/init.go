package event_service

import (
	_ "ms/sun2/servises/archiver_service"
)

type EventType int

const (
	ADDED_POST_EVENT       = EventType(1)
	DELETED_POST_EVENT     = EventType(2)
	LIKED_POST_EVENT       = EventType(10)
	UNLIKED_POST_EVENT     = EventType(11)
	COMMENTED_POST_EVENT   = EventType(20)
	UNCOMMENTED_POST_EVENT = EventType(21)
	FOLLOWED_USER_EVENT    = EventType(100)
	UNFOLLOWED_USER_EVENT  = EventType(101)
	BLOCKED_USER_EVENT     = EventType(110)
	UNBLOCKED_USER_EVENT   = EventType(111)
	ADDED_USER_ACTION      = EventType(1000)//for action table
	DELETED_USER_ACTION    = EventType(1001)
)

func init() {
	go delaySaveBuffer()
}
