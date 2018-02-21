package chat_service

type CHAT_SYNCType int
//todo move to some shared places
const (
	CHAT_SYNC_NEW_MESSAGE           = 2
	CHAT_SYNC_MSG_RECIVED_TO_SERVER = 3
	CHAT_SYNC_MSG_RECIVED_TO_PEER   = 4
	CHAT_SYNC_MSG_SEEN_BY_PEER      = 5
	CHAT_SYNC_MSG_DELETE            = 7
)
