package event_service

import "ms/sun2/shared/x"

type SubParam struct {
	Added_Post_Event       bool
	Deleted_Post_Event     bool
	Liked_Post_Event       bool
	UnLiked_Post_Event     bool
	Commented_Post_Event   bool
	UnCommented_Post_Event bool
	Followed_User_Event    bool
	UnFollwed_User_Event   bool
	Blocked_User_Event     bool
	UnBlocked_User_Event   bool
	Added_User_Action      bool
	Deleted_User_Action    bool
}

type Sub struct {
	Added_Post_Event       chan GeneralEvent
	Deleted_Post_Event     chan GeneralEvent
	Liked_Post_Event       chan GeneralEvent
	UnLiked_Post_Event     chan GeneralEvent
	Commented_Post_Event   chan GeneralEvent
	UnCommented_Post_Event chan GeneralEvent
	Followed_User_Event    chan GeneralEvent
	UnFollwed_User_Event   chan GeneralEvent
	Blocked_User_Event     chan GeneralEvent
	UnBlocked_User_Event   chan GeneralEvent
	Added_User_Action      chan GeneralEvent
	Deleted_User_Action    chan GeneralEvent
}

type GeneralEvent struct {
	//Event    x.Event
	x.Event
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
		Added_Post_Event:       make(chan GeneralEvent, NUM),
		Deleted_Post_Event:     make(chan GeneralEvent, NUM),
		Liked_Post_Event:       make(chan GeneralEvent, NUM),
		UnLiked_Post_Event:     make(chan GeneralEvent, NUM),
		Commented_Post_Event:   make(chan GeneralEvent, NUM),
		UnCommented_Post_Event: make(chan GeneralEvent, NUM),
		Followed_User_Event:    make(chan GeneralEvent, NUM),
		UnFollwed_User_Event:   make(chan GeneralEvent, NUM),
		Blocked_User_Event:     make(chan GeneralEvent, NUM),
		UnBlocked_User_Event:   make(chan GeneralEvent, NUM),
		Added_User_Action:      make(chan GeneralEvent, NUM),
		Deleted_User_Action:    make(chan GeneralEvent, NUM),
	}
	return sub
}
