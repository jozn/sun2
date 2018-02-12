package chat_service

type CHAT_SYNCType int

const (
	CHAT_SYNC_NEW_CHAT              = 1 //dep : no need chat alwyse must be loaded, to be fault tolerent
	CHAT_SYNC_NEW_MESSAGE           = 2
	CHAT_SYNC_MSG_RECIVED_TO_SERVER = 3
	CHAT_SYNC_MSG_RECIVED_TO_PEERT  = 4
	CHAT_SYNC_MSG_SEEN_BY_PEER      = 5
	CHAT_SYNC_MSG_CHAT_SYNC         = 6
	CHAT_SYNC_MSG_DELETE            = 7
)
