package event_service

import "ms/sun2/shared/x"

type SubParam struct {
	ADDED_POST_EVENT       bool
	DELETED_POST_EVENT     bool
	LIKED_POST_EVENT       bool
	UNLIKED_POST_EVENT     bool
	COMMENTED_POST_EVENT   bool
	UNCOMMENTED_POST_EVENT bool
	FOLLOWED_USER_EVENT    bool
	UNFOLLOWED_USER_EVENT  bool
	BLOCKED_USER_EVENT     bool
	UNBLOCKED_USER_EVENT   bool
}

type Sub struct {
	ADDED_POST_EVENT       chan GeneralEvent
	DELETED_POST_EVENT     chan GeneralEvent
	LIKED_POST_EVENT       chan GeneralEvent
	UNLIKED_POST_EVENT     chan GeneralEvent
	COMMENTED_POST_EVENT   chan GeneralEvent
	UNCOMMENTED_POST_EVENT chan GeneralEvent
	FOLLOWED_USER_EVENT    chan GeneralEvent
	UNFOLLOWED_USER_EVENT  chan GeneralEvent
	BLOCKED_USER_EVENT     chan GeneralEvent
	UNBLOCKED_USER_EVENT   chan GeneralEvent
}

type GeneralEvent struct {
	Event    x.Event
	Post     *x.Post
	Comment  *x.Comment
	Action   *x.Action
	ByUser   *x.User
	PeerUser *x.User
}

func newGeneralEvent(event x.Event) GeneralEvent {
	ev := GeneralEvent{
		Event: event,
	}

	return ev
}

func newSub() Sub {
	NUM := 100
	sub := Sub{
		ADDED_POST_EVENT:       make(chan GeneralEvent, NUM),
		DELETED_POST_EVENT:     make(chan GeneralEvent, NUM),
		LIKED_POST_EVENT:       make(chan GeneralEvent, NUM),
		UNLIKED_POST_EVENT:     make(chan GeneralEvent, NUM),
		COMMENTED_POST_EVENT:   make(chan GeneralEvent, NUM),
		UNCOMMENTED_POST_EVENT: make(chan GeneralEvent, NUM),
		FOLLOWED_USER_EVENT:    make(chan GeneralEvent, NUM),
		UNFOLLOWED_USER_EVENT:  make(chan GeneralEvent, NUM),
		BLOCKED_USER_EVENT:     make(chan GeneralEvent, NUM),
		UNBLOCKED_USER_EVENT:   make(chan GeneralEvent, NUM),
	}
	return sub
}
