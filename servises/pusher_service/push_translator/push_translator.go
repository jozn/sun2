package push_translator

import "ms/sun/shared/x"

var m = 1

func ChatPushToPbChat(pc *x.PushChat) x.PB_Push {
	pb := x.PB_Push{
		LastPushId:     int64(m),
		LastChatPushId: int64(pc.ToUserId),
	}
	m++
	return pb
}
