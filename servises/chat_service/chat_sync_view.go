package chat_service

import "ms/sun2/shared/x"

func ToUpdates(UserId int, syncs []*x.ChatSync) *x.PB_Updates {
	up := &x.PB_Updates{}
    roomKeys := []string{}

	for _, sync := range syncs {
		switch sync.ChatSyncTypeId {
		case CHAT_SYNC_NEW_MESSAGE:
            roomKeys = append(roomKeys, sync.RoomKey)
		case CHAT_SYNC_MSG_RECIVED_TO_SERVER:
			pb := &x.PB_UpdateMessageMeta{
				RoomKey:   sync.RoomKey,
				MessageId: int64(sync.MessageId),
			}
			up.MessagesReachedServer = append(up.MessagesReachedServer, pb)

		case CHAT_SYNC_MSG_RECIVED_TO_PEER:
			pb := &x.PB_UpdateMessageMeta{
				RoomKey:   sync.RoomKey,
				MessageId: int64(sync.MessageId),
			}
			up.MessagesDeliveredToUser = append(up.MessagesDeliveredToUser, pb)

		case CHAT_SYNC_MSG_SEEN_BY_PEER:
			pb := &x.PB_UpdateMessageMeta{
				RoomKey:   sync.RoomKey,
				MessageId: int64(sync.MessageId),
			}
			up.MessagesSeenByPeer = append(up.MessagesSeenByPeer, pb)
		}
	}
	if len(roomKeys) > 0 {
	    up.Chats = ChatViewsForChatList(UserId,roomKeys)
    }

	return up
}
