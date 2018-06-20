package x

type PB_RoomsChanges_Flat struct {
	VersionTime int
	Chats       []Chat
	rooms       []PB_ChatView
}

//ToPB
func (m *PB_RoomsChanges) ToFlat() *PB_RoomsChanges_Flat {
	r := &PB_RoomsChanges_Flat{
		VersionTime: int(m.VersionTime),
	}
	return r
}

//ToPB
func (m *PB_RoomsChanges_Flat) ToPB() *PB_RoomsChanges {
	r := &PB_RoomsChanges{
		VersionTime: uint64(m.VersionTime),
	}
	return r
}

//folding
var PB_RoomsChanges__FOlD = &PB_RoomsChanges{
	VersionTime: 0,
}

type PB_PushChanges_Flat struct {
	RoomsChanges               []PB_RoomsChanges
	ChatView                   []PB_ChatView
	InvalidateCacheForRoomKeys []string
	InvalidateAllChatCache     int
	InvalidateAllSocialCache   int
}

//ToPB
func (m *PB_PushChanges) ToFlat() *PB_PushChanges_Flat {
	r := &PB_PushChanges_Flat{

		InvalidateCacheForRoomKeys: m.InvalidateCacheForRoomKeys,
		InvalidateAllChatCache:     int(m.InvalidateAllChatCache),
		InvalidateAllSocialCache:   int(m.InvalidateAllSocialCache),
	}
	return r
}

//ToPB
func (m *PB_PushChanges_Flat) ToPB() *PB_PushChanges {
	r := &PB_PushChanges{

		InvalidateCacheForRoomKeys: m.InvalidateCacheForRoomKeys,
		InvalidateAllChatCache:     int32(m.InvalidateAllChatCache),
		InvalidateAllSocialCache:   int32(m.InvalidateAllSocialCache),
	}
	return r
}

//folding
var PB_PushChanges__FOlD = &PB_PushChanges{

	InvalidateAllChatCache:   0,
	InvalidateAllSocialCache: 0,
}

type PB_CommandToServer_Flat struct {
	ClientCallId   int
	Command        string
	RespondReached bool
	Data           []byte
}

//ToPB
func (m *PB_CommandToServer) ToFlat() *PB_CommandToServer_Flat {
	r := &PB_CommandToServer_Flat{
		ClientCallId:   int(m.ClientCallId),
		Command:        m.Command,
		RespondReached: m.RespondReached,
		Data:           []byte(m.Data),
	}
	return r
}

//ToPB
func (m *PB_CommandToServer_Flat) ToPB() *PB_CommandToServer {
	r := &PB_CommandToServer{
		ClientCallId:   int64(m.ClientCallId),
		Command:        m.Command,
		RespondReached: m.RespondReached,
		Data:           m.Data,
	}
	return r
}

//folding
var PB_CommandToServer__FOlD = &PB_CommandToServer{
	ClientCallId:   0,
	Command:        "",
	RespondReached: false,
	Data:           []byte{},
}

type PB_CommandToClient_Flat struct {
	ServerCallId   int
	Command        string
	RespondReached bool
	Data           []byte
}

//ToPB
func (m *PB_CommandToClient) ToFlat() *PB_CommandToClient_Flat {
	r := &PB_CommandToClient_Flat{
		ServerCallId:   int(m.ServerCallId),
		Command:        m.Command,
		RespondReached: m.RespondReached,
		Data:           []byte(m.Data),
	}
	return r
}

//ToPB
func (m *PB_CommandToClient_Flat) ToPB() *PB_CommandToClient {
	r := &PB_CommandToClient{
		ServerCallId:   int64(m.ServerCallId),
		Command:        m.Command,
		RespondReached: m.RespondReached,
		Data:           m.Data,
	}
	return r
}

//folding
var PB_CommandToClient__FOlD = &PB_CommandToClient{
	ServerCallId:   0,
	Command:        "",
	RespondReached: false,
	Data:           []byte{},
}

type PB_CommandReachedToServer_Flat struct {
	ClientCallId int
}

//ToPB
func (m *PB_CommandReachedToServer) ToFlat() *PB_CommandReachedToServer_Flat {
	r := &PB_CommandReachedToServer_Flat{
		ClientCallId: int(m.ClientCallId),
	}
	return r
}

//ToPB
func (m *PB_CommandReachedToServer_Flat) ToPB() *PB_CommandReachedToServer {
	r := &PB_CommandReachedToServer{
		ClientCallId: int64(m.ClientCallId),
	}
	return r
}

//folding
var PB_CommandReachedToServer__FOlD = &PB_CommandReachedToServer{
	ClientCallId: 0,
}

type PB_CommandReachedToClient_Flat struct {
	ServerCallId int
}

//ToPB
func (m *PB_CommandReachedToClient) ToFlat() *PB_CommandReachedToClient_Flat {
	r := &PB_CommandReachedToClient_Flat{
		ServerCallId: int(m.ServerCallId),
	}
	return r
}

//ToPB
func (m *PB_CommandReachedToClient_Flat) ToPB() *PB_CommandReachedToClient {
	r := &PB_CommandReachedToClient{
		ServerCallId: int64(m.ServerCallId),
	}
	return r
}

//folding
var PB_CommandReachedToClient__FOlD = &PB_CommandReachedToClient{
	ServerCallId: 0,
}

type PB_ResponseToClient_Flat struct {
	ClientCallId int
	PBClass      string
	RpcFullName  string
	Data         []byte
}

//ToPB
func (m *PB_ResponseToClient) ToFlat() *PB_ResponseToClient_Flat {
	r := &PB_ResponseToClient_Flat{
		ClientCallId: int(m.ClientCallId),
		PBClass:      m.PBClass,
		RpcFullName:  m.RpcFullName,
		Data:         []byte(m.Data),
	}
	return r
}

//ToPB
func (m *PB_ResponseToClient_Flat) ToPB() *PB_ResponseToClient {
	r := &PB_ResponseToClient{
		ClientCallId: int64(m.ClientCallId),
		PBClass:      m.PBClass,
		RpcFullName:  m.RpcFullName,
		Data:         m.Data,
	}
	return r
}

//folding
var PB_ResponseToClient__FOlD = &PB_ResponseToClient{
	ClientCallId: 0,
	PBClass:      "",
	RpcFullName:  "",
	Data:         []byte{},
}

type RPC_Auth_Types_Flat struct {
}

//ToPB
func (m *RPC_Auth_Types) ToFlat() *RPC_Auth_Types_Flat {
	r := &RPC_Auth_Types_Flat{}
	return r
}

//ToPB
func (m *RPC_Auth_Types_Flat) ToPB() *RPC_Auth_Types {
	r := &RPC_Auth_Types{}
	return r
}

//folding
var RPC_Auth_Types__FOlD = &RPC_Auth_Types{}

type PB_RPC_Chat_Types_Flat struct {
}

//ToPB
func (m *PB_RPC_Chat_Types) ToFlat() *PB_RPC_Chat_Types_Flat {
	r := &PB_RPC_Chat_Types_Flat{}
	return r
}

//ToPB
func (m *PB_RPC_Chat_Types_Flat) ToPB() *PB_RPC_Chat_Types {
	r := &PB_RPC_Chat_Types{}
	return r
}

//folding
var PB_RPC_Chat_Types__FOlD = &PB_RPC_Chat_Types{}

type RPC_General_Types_Flat struct {
}

//ToPB
func (m *RPC_General_Types) ToFlat() *RPC_General_Types_Flat {
	r := &RPC_General_Types_Flat{}
	return r
}

//ToPB
func (m *RPC_General_Types_Flat) ToPB() *RPC_General_Types {
	r := &RPC_General_Types{}
	return r
}

//folding
var RPC_General_Types__FOlD = &RPC_General_Types{}

type RPC_Page_Types_Flat struct {
}

//ToPB
func (m *RPC_Page_Types) ToFlat() *RPC_Page_Types_Flat {
	r := &RPC_Page_Types_Flat{}
	return r
}

//ToPB
func (m *RPC_Page_Types_Flat) ToPB() *RPC_Page_Types {
	r := &RPC_Page_Types{}
	return r
}

//folding
var RPC_Page_Types__FOlD = &RPC_Page_Types{}

type RPC_Social_Types_Flat struct {
}

//ToPB
func (m *RPC_Social_Types) ToFlat() *RPC_Social_Types_Flat {
	r := &RPC_Social_Types_Flat{}
	return r
}

//ToPB
func (m *RPC_Social_Types_Flat) ToPB() *RPC_Social_Types {
	r := &RPC_Social_Types{}
	return r
}

//folding
var RPC_Social_Types__FOlD = &RPC_Social_Types{}

type RPC_User_Types_Flat struct {
}

//ToPB
func (m *RPC_User_Types) ToFlat() *RPC_User_Types_Flat {
	r := &RPC_User_Types_Flat{}
	return r
}

//ToPB
func (m *RPC_User_Types_Flat) ToPB() *RPC_User_Types {
	r := &RPC_User_Types{}
	return r
}

//folding
var RPC_User_Types__FOlD = &RPC_User_Types{}

type PB_Action_Flat struct {
	ActionId     int
	ActorUserId  int
	ActionType   int
	PeerUserId   int
	PostId       int
	CommentId    int
	Murmur64Hash int
	CreatedTime  int
}

//ToPB
func (m *PB_Action) ToFlat() *PB_Action_Flat {
	r := &PB_Action_Flat{
		ActionId:     int(m.ActionId),
		ActorUserId:  int(m.ActorUserId),
		ActionType:   int(m.ActionType),
		PeerUserId:   int(m.PeerUserId),
		PostId:       int(m.PostId),
		CommentId:    int(m.CommentId),
		Murmur64Hash: int(m.Murmur64Hash),
		CreatedTime:  int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Action_Flat) ToPB() *PB_Action {
	r := &PB_Action{
		ActionId:     int64(m.ActionId),
		ActorUserId:  int32(m.ActorUserId),
		ActionType:   int32(m.ActionType),
		PeerUserId:   int32(m.PeerUserId),
		PostId:       int64(m.PostId),
		CommentId:    int64(m.CommentId),
		Murmur64Hash: int64(m.Murmur64Hash),
		CreatedTime:  int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Action__FOlD = &PB_Action{
	ActionId:     0,
	ActorUserId:  0,
	ActionType:   0,
	PeerUserId:   0,
	PostId:       0,
	CommentId:    0,
	Murmur64Hash: 0,
	CreatedTime:  0,
}

type PB_Blocked_Flat struct {
	Id            int
	UserId        int
	BlockedUserId int
	CreatedTime   int
}

//ToPB
func (m *PB_Blocked) ToFlat() *PB_Blocked_Flat {
	r := &PB_Blocked_Flat{
		Id:            int(m.Id),
		UserId:        int(m.UserId),
		BlockedUserId: int(m.BlockedUserId),
		CreatedTime:   int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Blocked_Flat) ToPB() *PB_Blocked {
	r := &PB_Blocked{
		Id:            int64(m.Id),
		UserId:        int32(m.UserId),
		BlockedUserId: int32(m.BlockedUserId),
		CreatedTime:   int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Blocked__FOlD = &PB_Blocked{
	Id:            0,
	UserId:        0,
	BlockedUserId: 0,
	CreatedTime:   0,
}

type PB_Comment_Flat struct {
	CommentId   int
	UserId      int
	PostId      int
	Text        string
	LikesCount  int
	IsEdited    int
	CreatedTime int
}

//ToPB
func (m *PB_Comment) ToFlat() *PB_Comment_Flat {
	r := &PB_Comment_Flat{
		CommentId:   int(m.CommentId),
		UserId:      int(m.UserId),
		PostId:      int(m.PostId),
		Text:        m.Text,
		LikesCount:  int(m.LikesCount),
		IsEdited:    int(m.IsEdited),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Comment_Flat) ToPB() *PB_Comment {
	r := &PB_Comment{
		CommentId:   int64(m.CommentId),
		UserId:      int32(m.UserId),
		PostId:      int64(m.PostId),
		Text:        m.Text,
		LikesCount:  int32(m.LikesCount),
		IsEdited:    int32(m.IsEdited),
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Comment__FOlD = &PB_Comment{
	CommentId:   0,
	UserId:      0,
	PostId:      0,
	Text:        "",
	LikesCount:  0,
	IsEdited:    0,
	CreatedTime: 0,
}

type PB_CommentDeleted_Flat struct {
	CommentId int
	UserId    int
}

//ToPB
func (m *PB_CommentDeleted) ToFlat() *PB_CommentDeleted_Flat {
	r := &PB_CommentDeleted_Flat{
		CommentId: int(m.CommentId),
		UserId:    int(m.UserId),
	}
	return r
}

//ToPB
func (m *PB_CommentDeleted_Flat) ToPB() *PB_CommentDeleted {
	r := &PB_CommentDeleted{
		CommentId: int64(m.CommentId),
		UserId:    int32(m.UserId),
	}
	return r
}

//folding
var PB_CommentDeleted__FOlD = &PB_CommentDeleted{
	CommentId: 0,
	UserId:    0,
}

type PB_Event_Flat struct {
	EventId      int
	EventType    int
	ByUserId     int
	PeerUserId   int
	PostId       int
	CommentId    int
	HashTagId    int
	GroupId      int
	ActionId     int
	ChatId       int
	ChatKey      string
	MessageId    int
	ReSharedId   int
	Murmur64Hash int
}

//ToPB
func (m *PB_Event) ToFlat() *PB_Event_Flat {
	r := &PB_Event_Flat{
		EventId:      int(m.EventId),
		EventType:    int(m.EventType),
		ByUserId:     int(m.ByUserId),
		PeerUserId:   int(m.PeerUserId),
		PostId:       int(m.PostId),
		CommentId:    int(m.CommentId),
		HashTagId:    int(m.HashTagId),
		GroupId:      int(m.GroupId),
		ActionId:     int(m.ActionId),
		ChatId:       int(m.ChatId),
		ChatKey:      m.ChatKey,
		MessageId:    int(m.MessageId),
		ReSharedId:   int(m.ReSharedId),
		Murmur64Hash: int(m.Murmur64Hash),
	}
	return r
}

//ToPB
func (m *PB_Event_Flat) ToPB() *PB_Event {
	r := &PB_Event{
		EventId:      int64(m.EventId),
		EventType:    int32(m.EventType),
		ByUserId:     int64(m.ByUserId),
		PeerUserId:   int64(m.PeerUserId),
		PostId:       int64(m.PostId),
		CommentId:    int64(m.CommentId),
		HashTagId:    int64(m.HashTagId),
		GroupId:      int64(m.GroupId),
		ActionId:     int64(m.ActionId),
		ChatId:       int64(m.ChatId),
		ChatKey:      m.ChatKey,
		MessageId:    int64(m.MessageId),
		ReSharedId:   int64(m.ReSharedId),
		Murmur64Hash: int64(m.Murmur64Hash),
	}
	return r
}

//folding
var PB_Event__FOlD = &PB_Event{
	EventId:      0,
	EventType:    0,
	ByUserId:     0,
	PeerUserId:   0,
	PostId:       0,
	CommentId:    0,
	HashTagId:    0,
	GroupId:      0,
	ActionId:     0,
	ChatId:       0,
	ChatKey:      "",
	MessageId:    0,
	ReSharedId:   0,
	Murmur64Hash: 0,
}

type PB_Followed_Flat struct {
	Id             int
	UserId         int
	FollowedUserId int
	CreatedTime    int
}

//ToPB
func (m *PB_Followed) ToFlat() *PB_Followed_Flat {
	r := &PB_Followed_Flat{
		Id:             int(m.Id),
		UserId:         int(m.UserId),
		FollowedUserId: int(m.FollowedUserId),
		CreatedTime:    int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Followed_Flat) ToPB() *PB_Followed {
	r := &PB_Followed{
		Id:             int64(m.Id),
		UserId:         int32(m.UserId),
		FollowedUserId: int32(m.FollowedUserId),
		CreatedTime:    int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Followed__FOlD = &PB_Followed{
	Id:             0,
	UserId:         0,
	FollowedUserId: 0,
	CreatedTime:    0,
}

type PB_Likes_Flat struct {
	Id          int
	PostId      int
	UserId      int
	PostType    int
	CreatedTime int
}

//ToPB
func (m *PB_Likes) ToFlat() *PB_Likes_Flat {
	r := &PB_Likes_Flat{
		Id:          int(m.Id),
		PostId:      int(m.PostId),
		UserId:      int(m.UserId),
		PostType:    int(m.PostType),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Likes_Flat) ToPB() *PB_Likes {
	r := &PB_Likes{
		Id:          int64(m.Id),
		PostId:      int64(m.PostId),
		UserId:      int32(m.UserId),
		PostType:    int32(m.PostType),
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Likes__FOlD = &PB_Likes{
	Id:          0,
	PostId:      0,
	UserId:      0,
	PostType:    0,
	CreatedTime: 0,
}

type PB_Notify_Flat struct {
	NotifyId     int
	ForUserId    int
	ActorUserId  int
	NotifyType   int
	PostId       int
	CommentId    int
	PeerUserId   int
	Murmur64Hash int
	SeenStatus   int
	CreatedTime  int
}

//ToPB
func (m *PB_Notify) ToFlat() *PB_Notify_Flat {
	r := &PB_Notify_Flat{
		NotifyId:     int(m.NotifyId),
		ForUserId:    int(m.ForUserId),
		ActorUserId:  int(m.ActorUserId),
		NotifyType:   int(m.NotifyType),
		PostId:       int(m.PostId),
		CommentId:    int(m.CommentId),
		PeerUserId:   int(m.PeerUserId),
		Murmur64Hash: int(m.Murmur64Hash),
		SeenStatus:   int(m.SeenStatus),
		CreatedTime:  int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Notify_Flat) ToPB() *PB_Notify {
	r := &PB_Notify{
		NotifyId:     int64(m.NotifyId),
		ForUserId:    int32(m.ForUserId),
		ActorUserId:  int32(m.ActorUserId),
		NotifyType:   int32(m.NotifyType),
		PostId:       int64(m.PostId),
		CommentId:    int64(m.CommentId),
		PeerUserId:   int32(m.PeerUserId),
		Murmur64Hash: int64(m.Murmur64Hash),
		SeenStatus:   int32(m.SeenStatus),
		CreatedTime:  int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Notify__FOlD = &PB_Notify{
	NotifyId:     0,
	ForUserId:    0,
	ActorUserId:  0,
	NotifyType:   0,
	PostId:       0,
	CommentId:    0,
	PeerUserId:   0,
	Murmur64Hash: 0,
	SeenStatus:   0,
	CreatedTime:  0,
}

type PB_NotifyRemoved_Flat struct {
	Murmur64Hash int
	ForUserId    int
	Id           int
}

//ToPB
func (m *PB_NotifyRemoved) ToFlat() *PB_NotifyRemoved_Flat {
	r := &PB_NotifyRemoved_Flat{
		Murmur64Hash: int(m.Murmur64Hash),
		ForUserId:    int(m.ForUserId),
		Id:           int(m.Id),
	}
	return r
}

//ToPB
func (m *PB_NotifyRemoved_Flat) ToPB() *PB_NotifyRemoved {
	r := &PB_NotifyRemoved{
		Murmur64Hash: int64(m.Murmur64Hash),
		ForUserId:    int32(m.ForUserId),
		Id:           int64(m.Id),
	}
	return r
}

//folding
var PB_NotifyRemoved__FOlD = &PB_NotifyRemoved{
	Murmur64Hash: 0,
	ForUserId:    0,
	Id:           0,
}

type PB_PhoneContacts_Flat struct {
	Id        int
	UserId    int
	ClientId  int
	Phone     string
	FirstName string
	LastName  string
}

//ToPB
func (m *PB_PhoneContacts) ToFlat() *PB_PhoneContacts_Flat {
	r := &PB_PhoneContacts_Flat{
		Id:        int(m.Id),
		UserId:    int(m.UserId),
		ClientId:  int(m.ClientId),
		Phone:     m.Phone,
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}
	return r
}

//ToPB
func (m *PB_PhoneContacts_Flat) ToPB() *PB_PhoneContacts {
	r := &PB_PhoneContacts{
		Id:        int64(m.Id),
		UserId:    int32(m.UserId),
		ClientId:  int64(m.ClientId),
		Phone:     m.Phone,
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}
	return r
}

//folding
var PB_PhoneContacts__FOlD = &PB_PhoneContacts{
	Id:        0,
	UserId:    0,
	ClientId:  0,
	Phone:     "",
	FirstName: "",
	LastName:  "",
}

type PB_Post_Flat struct {
	PostId         int
	UserId         int
	PostType       int
	MediaId        int
	FileRefId      int
	PostKey        string
	Text           string
	RichText       string
	MediaCount     int
	SharedTo       int
	DisableComment int
	Via            int
	Seq            int
	CommentsCount  int
	LikesCount     int
	ViewsCount     int
	EditedTime     int
	CreatedTime    int
}

//ToPB
func (m *PB_Post) ToFlat() *PB_Post_Flat {
	r := &PB_Post_Flat{
		PostId:         int(m.PostId),
		UserId:         int(m.UserId),
		PostType:       int(m.PostType),
		MediaId:        int(m.MediaId),
		FileRefId:      int(m.FileRefId),
		PostKey:        m.PostKey,
		Text:           m.Text,
		RichText:       m.RichText,
		MediaCount:     int(m.MediaCount),
		SharedTo:       int(m.SharedTo),
		DisableComment: int(m.DisableComment),
		Via:            int(m.Via),
		Seq:            int(m.Seq),
		CommentsCount:  int(m.CommentsCount),
		LikesCount:     int(m.LikesCount),
		ViewsCount:     int(m.ViewsCount),
		EditedTime:     int(m.EditedTime),
		CreatedTime:    int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Post_Flat) ToPB() *PB_Post {
	r := &PB_Post{
		PostId:         int64(m.PostId),
		UserId:         int32(m.UserId),
		PostType:       int32(m.PostType),
		MediaId:        int64(m.MediaId),
		FileRefId:      int64(m.FileRefId),
		PostKey:        m.PostKey,
		Text:           m.Text,
		RichText:       m.RichText,
		MediaCount:     int32(m.MediaCount),
		SharedTo:       int32(m.SharedTo),
		DisableComment: int32(m.DisableComment),
		Via:            int32(m.Via),
		Seq:            int32(m.Seq),
		CommentsCount:  int32(m.CommentsCount),
		LikesCount:     int32(m.LikesCount),
		ViewsCount:     int32(m.ViewsCount),
		EditedTime:     int32(m.EditedTime),
		CreatedTime:    int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Post__FOlD = &PB_Post{
	PostId:         0,
	UserId:         0,
	PostType:       0,
	MediaId:        0,
	FileRefId:      0,
	PostKey:        "",
	Text:           "",
	RichText:       "",
	MediaCount:     0,
	SharedTo:       0,
	DisableComment: 0,
	Via:            0,
	Seq:            0,
	CommentsCount:  0,
	LikesCount:     0,
	ViewsCount:     0,
	EditedTime:     0,
	CreatedTime:    0,
}

type PB_PostCount_Flat struct {
	PostId          int
	CommentsCount   int
	LikesCount      int
	ViewsCount      int
	ReSharedCount   int
	ChatSharedCount int
}

//ToPB
func (m *PB_PostCount) ToFlat() *PB_PostCount_Flat {
	r := &PB_PostCount_Flat{
		PostId:          int(m.PostId),
		CommentsCount:   int(m.CommentsCount),
		LikesCount:      int(m.LikesCount),
		ViewsCount:      int(m.ViewsCount),
		ReSharedCount:   int(m.ReSharedCount),
		ChatSharedCount: int(m.ChatSharedCount),
	}
	return r
}

//ToPB
func (m *PB_PostCount_Flat) ToPB() *PB_PostCount {
	r := &PB_PostCount{
		PostId:          int64(m.PostId),
		CommentsCount:   int32(m.CommentsCount),
		LikesCount:      int32(m.LikesCount),
		ViewsCount:      int64(m.ViewsCount),
		ReSharedCount:   int32(m.ReSharedCount),
		ChatSharedCount: int32(m.ChatSharedCount),
	}
	return r
}

//folding
var PB_PostCount__FOlD = &PB_PostCount{
	PostId:          0,
	CommentsCount:   0,
	LikesCount:      0,
	ViewsCount:      0,
	ReSharedCount:   0,
	ChatSharedCount: 0,
}

type PB_PostDeleted_Flat struct {
	PostId int
	UserId int
}

//ToPB
func (m *PB_PostDeleted) ToFlat() *PB_PostDeleted_Flat {
	r := &PB_PostDeleted_Flat{
		PostId: int(m.PostId),
		UserId: int(m.UserId),
	}
	return r
}

//ToPB
func (m *PB_PostDeleted_Flat) ToPB() *PB_PostDeleted {
	r := &PB_PostDeleted{
		PostId: int64(m.PostId),
		UserId: int32(m.UserId),
	}
	return r
}

//folding
var PB_PostDeleted__FOlD = &PB_PostDeleted{
	PostId: 0,
	UserId: 0,
}

type PB_PostKeys_Flat struct {
	Id         int
	PostKeyStr string
	Used       int
	RandShard  int
}

//ToPB
func (m *PB_PostKeys) ToFlat() *PB_PostKeys_Flat {
	r := &PB_PostKeys_Flat{
		Id:         int(m.Id),
		PostKeyStr: m.PostKeyStr,
		Used:       int(m.Used),
		RandShard:  int(m.RandShard),
	}
	return r
}

//ToPB
func (m *PB_PostKeys_Flat) ToPB() *PB_PostKeys {
	r := &PB_PostKeys{
		Id:         int32(m.Id),
		PostKeyStr: m.PostKeyStr,
		Used:       int32(m.Used),
		RandShard:  int32(m.RandShard),
	}
	return r
}

//folding
var PB_PostKeys__FOlD = &PB_PostKeys{
	Id:         0,
	PostKeyStr: "",
	Used:       0,
	RandShard:  0,
}

type PB_PostLink_Flat struct {
	LinkId  int
	LinkUrl string
}

//ToPB
func (m *PB_PostLink) ToFlat() *PB_PostLink_Flat {
	r := &PB_PostLink_Flat{
		LinkId:  int(m.LinkId),
		LinkUrl: m.LinkUrl,
	}
	return r
}

//ToPB
func (m *PB_PostLink_Flat) ToPB() *PB_PostLink {
	r := &PB_PostLink{
		LinkId:  int64(m.LinkId),
		LinkUrl: m.LinkUrl,
	}
	return r
}

//folding
var PB_PostLink__FOlD = &PB_PostLink{
	LinkId:  0,
	LinkUrl: "",
}

type PB_PostMedia_Flat struct {
	MediaId       int
	UserId        int
	PostId        int
	AlbumId       int
	MediaTypeEnum int
	Width         int
	Height        int
	Size          int
	Duration      int
	Extension     string
	Md5Hash       string
	Color         string
	CreatedTime   int
	ViewCount     int
	Extra         string
}

//ToPB
func (m *PB_PostMedia) ToFlat() *PB_PostMedia_Flat {
	r := &PB_PostMedia_Flat{
		MediaId:       int(m.MediaId),
		UserId:        int(m.UserId),
		PostId:        int(m.PostId),
		AlbumId:       int(m.AlbumId),
		MediaTypeEnum: int(m.MediaTypeEnum),
		Width:         int(m.Width),
		Height:        int(m.Height),
		Size:          int(m.Size),
		Duration:      int(m.Duration),
		Extension:     m.Extension,
		Md5Hash:       m.Md5Hash,
		Color:         m.Color,
		CreatedTime:   int(m.CreatedTime),
		ViewCount:     int(m.ViewCount),
		Extra:         m.Extra,
	}
	return r
}

//ToPB
func (m *PB_PostMedia_Flat) ToPB() *PB_PostMedia {
	r := &PB_PostMedia{
		MediaId:       int64(m.MediaId),
		UserId:        int32(m.UserId),
		PostId:        int64(m.PostId),
		AlbumId:       int32(m.AlbumId),
		MediaTypeEnum: int32(m.MediaTypeEnum),
		Width:         int32(m.Width),
		Height:        int32(m.Height),
		Size:          int32(m.Size),
		Duration:      int32(m.Duration),
		Extension:     m.Extension,
		Md5Hash:       m.Md5Hash,
		Color:         m.Color,
		CreatedTime:   int32(m.CreatedTime),
		ViewCount:     int32(m.ViewCount),
		Extra:         m.Extra,
	}
	return r
}

//folding
var PB_PostMedia__FOlD = &PB_PostMedia{
	MediaId:       0,
	UserId:        0,
	PostId:        0,
	AlbumId:       0,
	MediaTypeEnum: 0,
	Width:         0,
	Height:        0,
	Size:          0,
	Duration:      0,
	Extension:     "",
	Md5Hash:       "",
	Color:         "",
	CreatedTime:   0,
	ViewCount:     0,
	Extra:         "",
}

type PB_PostReshared_Flat struct {
	ResharedId  int
	PostId      int
	ByUserId    int
	PostUserId  int
	CreatedTime int
}

//ToPB
func (m *PB_PostReshared) ToFlat() *PB_PostReshared_Flat {
	r := &PB_PostReshared_Flat{
		ResharedId:  int(m.ResharedId),
		PostId:      int(m.PostId),
		ByUserId:    int(m.ByUserId),
		PostUserId:  int(m.PostUserId),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_PostReshared_Flat) ToPB() *PB_PostReshared {
	r := &PB_PostReshared{
		ResharedId:  int64(m.ResharedId),
		PostId:      int64(m.PostId),
		ByUserId:    int32(m.ByUserId),
		PostUserId:  int32(m.PostUserId),
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_PostReshared__FOlD = &PB_PostReshared{
	ResharedId:  0,
	PostId:      0,
	ByUserId:    0,
	PostUserId:  0,
	CreatedTime: 0,
}

type PB_ProfileAll_Flat struct {
	Id         int
	UserId     int
	PostId     int
	IsReShared int
}

//ToPB
func (m *PB_ProfileAll) ToFlat() *PB_ProfileAll_Flat {
	r := &PB_ProfileAll_Flat{
		Id:         int(m.Id),
		UserId:     int(m.UserId),
		PostId:     int(m.PostId),
		IsReShared: int(m.IsReShared),
	}
	return r
}

//ToPB
func (m *PB_ProfileAll_Flat) ToPB() *PB_ProfileAll {
	r := &PB_ProfileAll{
		Id:         int64(m.Id),
		UserId:     int32(m.UserId),
		PostId:     int64(m.PostId),
		IsReShared: int32(m.IsReShared),
	}
	return r
}

//folding
var PB_ProfileAll__FOlD = &PB_ProfileAll{
	Id:         0,
	UserId:     0,
	PostId:     0,
	IsReShared: 0,
}

type PB_ProfileMedia_Flat struct {
	Id       int
	UserId   int
	PostId   int
	PostType int
}

//ToPB
func (m *PB_ProfileMedia) ToFlat() *PB_ProfileMedia_Flat {
	r := &PB_ProfileMedia_Flat{
		Id:       int(m.Id),
		UserId:   int(m.UserId),
		PostId:   int(m.PostId),
		PostType: int(m.PostType),
	}
	return r
}

//ToPB
func (m *PB_ProfileMedia_Flat) ToPB() *PB_ProfileMedia {
	r := &PB_ProfileMedia{
		Id:       int64(m.Id),
		UserId:   int32(m.UserId),
		PostId:   int64(m.PostId),
		PostType: int32(m.PostType),
	}
	return r
}

//folding
var PB_ProfileMedia__FOlD = &PB_ProfileMedia{
	Id:       0,
	UserId:   0,
	PostId:   0,
	PostType: 0,
}

type PB_ProfileMentioned_Flat struct {
	Id          int
	ForUserId   int
	PostId      int
	PostUserId  int
	PostType    int
	CreatedTime int
}

//ToPB
func (m *PB_ProfileMentioned) ToFlat() *PB_ProfileMentioned_Flat {
	r := &PB_ProfileMentioned_Flat{
		Id:          int(m.Id),
		ForUserId:   int(m.ForUserId),
		PostId:      int(m.PostId),
		PostUserId:  int(m.PostUserId),
		PostType:    int(m.PostType),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_ProfileMentioned_Flat) ToPB() *PB_ProfileMentioned {
	r := &PB_ProfileMentioned{
		Id:          int64(m.Id),
		ForUserId:   int32(m.ForUserId),
		PostId:      int64(m.PostId),
		PostUserId:  int32(m.PostUserId),
		PostType:    int32(m.PostType),
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_ProfileMentioned__FOlD = &PB_ProfileMentioned{
	Id:          0,
	ForUserId:   0,
	PostId:      0,
	PostUserId:  0,
	PostType:    0,
	CreatedTime: 0,
}

type PB_Session_Flat struct {
	Id            int
	SessionUuid   string
	UserId        int
	LastIpAddress string
	AppVersion    int
	ActiveTime    int
	CreatedTime   int
}

//ToPB
func (m *PB_Session) ToFlat() *PB_Session_Flat {
	r := &PB_Session_Flat{
		Id:            int(m.Id),
		SessionUuid:   m.SessionUuid,
		UserId:        int(m.UserId),
		LastIpAddress: m.LastIpAddress,
		AppVersion:    int(m.AppVersion),
		ActiveTime:    int(m.ActiveTime),
		CreatedTime:   int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Session_Flat) ToPB() *PB_Session {
	r := &PB_Session{
		Id:            int64(m.Id),
		SessionUuid:   m.SessionUuid,
		UserId:        int32(m.UserId),
		LastIpAddress: m.LastIpAddress,
		AppVersion:    int32(m.AppVersion),
		ActiveTime:    int32(m.ActiveTime),
		CreatedTime:   int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Session__FOlD = &PB_Session{
	Id:            0,
	SessionUuid:   "",
	UserId:        0,
	LastIpAddress: "",
	AppVersion:    0,
	ActiveTime:    0,
	CreatedTime:   0,
}

type PB_SettingNotifications_Flat struct {
	UserId                   int
	SocialLedOn              int
	SocialLedColor           string
	ReqestToFollowYou        int
	FollowedYou              int
	AccptedYourFollowRequest int
	YourPostLiked            int
	YourPostCommented        int
	MenthenedYouInPost       int
	MenthenedYouInComment    int
	YourContactsJoined       int
	DirectMessage            int
	DirectAlert              int
	DirectPerview            int
	DirectLedOn              int
	DirectLedColor           int
	DirectVibrate            int
	DirectPopup              int
	DirectSound              int
	DirectPriority           int
}

//ToPB
func (m *PB_SettingNotifications) ToFlat() *PB_SettingNotifications_Flat {
	r := &PB_SettingNotifications_Flat{
		UserId:                   int(m.UserId),
		SocialLedOn:              int(m.SocialLedOn),
		SocialLedColor:           m.SocialLedColor,
		ReqestToFollowYou:        int(m.ReqestToFollowYou),
		FollowedYou:              int(m.FollowedYou),
		AccptedYourFollowRequest: int(m.AccptedYourFollowRequest),
		YourPostLiked:            int(m.YourPostLiked),
		YourPostCommented:        int(m.YourPostCommented),
		MenthenedYouInPost:       int(m.MenthenedYouInPost),
		MenthenedYouInComment:    int(m.MenthenedYouInComment),
		YourContactsJoined:       int(m.YourContactsJoined),
		DirectMessage:            int(m.DirectMessage),
		DirectAlert:              int(m.DirectAlert),
		DirectPerview:            int(m.DirectPerview),
		DirectLedOn:              int(m.DirectLedOn),
		DirectLedColor:           int(m.DirectLedColor),
		DirectVibrate:            int(m.DirectVibrate),
		DirectPopup:              int(m.DirectPopup),
		DirectSound:              int(m.DirectSound),
		DirectPriority:           int(m.DirectPriority),
	}
	return r
}

//ToPB
func (m *PB_SettingNotifications_Flat) ToPB() *PB_SettingNotifications {
	r := &PB_SettingNotifications{
		UserId:                   int32(m.UserId),
		SocialLedOn:              int32(m.SocialLedOn),
		SocialLedColor:           m.SocialLedColor,
		ReqestToFollowYou:        int32(m.ReqestToFollowYou),
		FollowedYou:              int32(m.FollowedYou),
		AccptedYourFollowRequest: int32(m.AccptedYourFollowRequest),
		YourPostLiked:            int32(m.YourPostLiked),
		YourPostCommented:        int32(m.YourPostCommented),
		MenthenedYouInPost:       int32(m.MenthenedYouInPost),
		MenthenedYouInComment:    int32(m.MenthenedYouInComment),
		YourContactsJoined:       int32(m.YourContactsJoined),
		DirectMessage:            int32(m.DirectMessage),
		DirectAlert:              int32(m.DirectAlert),
		DirectPerview:            int32(m.DirectPerview),
		DirectLedOn:              int32(m.DirectLedOn),
		DirectLedColor:           int32(m.DirectLedColor),
		DirectVibrate:            int32(m.DirectVibrate),
		DirectPopup:              int32(m.DirectPopup),
		DirectSound:              int32(m.DirectSound),
		DirectPriority:           int32(m.DirectPriority),
	}
	return r
}

//folding
var PB_SettingNotifications__FOlD = &PB_SettingNotifications{
	UserId:                   0,
	SocialLedOn:              0,
	SocialLedColor:           "",
	ReqestToFollowYou:        0,
	FollowedYou:              0,
	AccptedYourFollowRequest: 0,
	YourPostLiked:            0,
	YourPostCommented:        0,
	MenthenedYouInPost:       0,
	MenthenedYouInComment:    0,
	YourContactsJoined:       0,
	DirectMessage:            0,
	DirectAlert:              0,
	DirectPerview:            0,
	DirectLedOn:              0,
	DirectLedColor:           0,
	DirectVibrate:            0,
	DirectPopup:              0,
	DirectSound:              0,
	DirectPriority:           0,
}

type PB_Sms_Flat struct {
	Id              int
	Hash            string
	AppUuid         string
	ClientPhone     string
	GenratedCode    int
	SmsSenderNumber string
	SmsSendStatues  string
	SmsHttpBody     string
	Err             string
	Carrier         string
	Country         []byte
	IsValidPhone    int
	IsConfirmed     int
	IsLogin         int
	IsRegister      int
	RetriedCount    int
	TTL             int
}

//ToPB
func (m *PB_Sms) ToFlat() *PB_Sms_Flat {
	r := &PB_Sms_Flat{
		Id:              int(m.Id),
		Hash:            m.Hash,
		AppUuid:         m.AppUuid,
		ClientPhone:     m.ClientPhone,
		GenratedCode:    int(m.GenratedCode),
		SmsSenderNumber: m.SmsSenderNumber,
		SmsSendStatues:  m.SmsSendStatues,
		SmsHttpBody:     m.SmsHttpBody,
		Err:             m.Err,
		Carrier:         m.Carrier,
		Country:         []byte(m.Country),
		IsValidPhone:    int(m.IsValidPhone),
		IsConfirmed:     int(m.IsConfirmed),
		IsLogin:         int(m.IsLogin),
		IsRegister:      int(m.IsRegister),
		RetriedCount:    int(m.RetriedCount),
		TTL:             int(m.TTL),
	}
	return r
}

//ToPB
func (m *PB_Sms_Flat) ToPB() *PB_Sms {
	r := &PB_Sms{
		Id:              int32(m.Id),
		Hash:            m.Hash,
		AppUuid:         m.AppUuid,
		ClientPhone:     m.ClientPhone,
		GenratedCode:    int32(m.GenratedCode),
		SmsSenderNumber: m.SmsSenderNumber,
		SmsSendStatues:  m.SmsSendStatues,
		SmsHttpBody:     m.SmsHttpBody,
		Err:             m.Err,
		Carrier:         m.Carrier,
		Country:         m.Country,
		IsValidPhone:    int32(m.IsValidPhone),
		IsConfirmed:     int32(m.IsConfirmed),
		IsLogin:         int32(m.IsLogin),
		IsRegister:      int32(m.IsRegister),
		RetriedCount:    int32(m.RetriedCount),
		TTL:             int32(m.TTL),
	}
	return r
}

//folding
var PB_Sms__FOlD = &PB_Sms{
	Id:              0,
	Hash:            "",
	AppUuid:         "",
	ClientPhone:     "",
	GenratedCode:    0,
	SmsSenderNumber: "",
	SmsSendStatues:  "",
	SmsHttpBody:     "",
	Err:             "",
	Carrier:         "",
	Country:         []byte{},
	IsValidPhone:    0,
	IsConfirmed:     0,
	IsLogin:         0,
	IsRegister:      0,
	RetriedCount:    0,
	TTL:             0,
}

type PB_Tag_Flat struct {
	TagId         int
	Name          string
	Count         int
	TagStatusEnum int
	IsBlocked     int
	GroupId       int
	CreatedTime   int
}

//ToPB
func (m *PB_Tag) ToFlat() *PB_Tag_Flat {
	r := &PB_Tag_Flat{
		TagId:         int(m.TagId),
		Name:          m.Name,
		Count:         int(m.Count),
		TagStatusEnum: int(m.TagStatusEnum),
		IsBlocked:     int(m.IsBlocked),
		GroupId:       int(m.GroupId),
		CreatedTime:   int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Tag_Flat) ToPB() *PB_Tag {
	r := &PB_Tag{
		TagId:         int64(m.TagId),
		Name:          m.Name,
		Count:         int32(m.Count),
		TagStatusEnum: int32(m.TagStatusEnum),
		IsBlocked:     int32(m.IsBlocked),
		GroupId:       int32(m.GroupId),
		CreatedTime:   int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Tag__FOlD = &PB_Tag{
	TagId:         0,
	Name:          "",
	Count:         0,
	TagStatusEnum: 0,
	IsBlocked:     0,
	GroupId:       0,
	CreatedTime:   0,
}

type PB_TagPost_Flat struct {
	Id          int
	TagId       int
	PostId      int
	PostType    int
	CreatedTime int
}

//ToPB
func (m *PB_TagPost) ToFlat() *PB_TagPost_Flat {
	r := &PB_TagPost_Flat{
		Id:          int(m.Id),
		TagId:       int(m.TagId),
		PostId:      int(m.PostId),
		PostType:    int(m.PostType),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_TagPost_Flat) ToPB() *PB_TagPost {
	r := &PB_TagPost{
		Id:          int64(m.Id),
		TagId:       int32(m.TagId),
		PostId:      int32(m.PostId),
		PostType:    int32(m.PostType),
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_TagPost__FOlD = &PB_TagPost{
	Id:          0,
	TagId:       0,
	PostId:      0,
	PostType:    0,
	CreatedTime: 0,
}

type PB_TriggerLog_Flat struct {
	Id         int
	ModelName  string
	ChangeType string
	TargetId   int
	TargetStr  string
	CreatedSe  int
}

//ToPB
func (m *PB_TriggerLog) ToFlat() *PB_TriggerLog_Flat {
	r := &PB_TriggerLog_Flat{
		Id:         int(m.Id),
		ModelName:  m.ModelName,
		ChangeType: m.ChangeType,
		TargetId:   int(m.TargetId),
		TargetStr:  m.TargetStr,
		CreatedSe:  int(m.CreatedSe),
	}
	return r
}

//ToPB
func (m *PB_TriggerLog_Flat) ToPB() *PB_TriggerLog {
	r := &PB_TriggerLog{
		Id:         int64(m.Id),
		ModelName:  m.ModelName,
		ChangeType: m.ChangeType,
		TargetId:   int64(m.TargetId),
		TargetStr:  m.TargetStr,
		CreatedSe:  int32(m.CreatedSe),
	}
	return r
}

//folding
var PB_TriggerLog__FOlD = &PB_TriggerLog{
	Id:         0,
	ModelName:  "",
	ChangeType: "",
	TargetId:   0,
	TargetStr:  "",
	CreatedSe:  0,
}

type PB_User_Flat struct {
	UserId             int
	UserName           string
	UserNameLower      string
	FirstName          string
	LastName           string
	IsVerified         int
	AvatarId           int
	AccessHash         int
	ProfilePrivacy     int
	OnlinePrivacy      int
	CallPrivacy        int
	AddToGroupPrivacy  int
	SeenMessagePrivacy int
	Phone              string
	Email              string
	About              string
	PasswordHash       string
	PasswordSalt       string
	PostSeq            int
	FollowersCount     int
	FollowingCount     int
	PostsCount         int
	MediaCount         int
	PhotoCount         int
	VideoCount         int
	GifCount           int
	AudioCount         int
	VoiceCount         int
	FileCount          int
	LinkCount          int
	BoardCount         int
	PinedCount         int
	LikesCount         int
	ResharedCount      int
	LastPostTime       int
	CreatedTime        int
	VersionTime        int
	IsDeleted          int
	IsBanned           int
}

//ToPB
func (m *PB_User) ToFlat() *PB_User_Flat {
	r := &PB_User_Flat{
		UserId:             int(m.UserId),
		UserName:           m.UserName,
		UserNameLower:      m.UserNameLower,
		FirstName:          m.FirstName,
		LastName:           m.LastName,
		IsVerified:         int(m.IsVerified),
		AvatarId:           int(m.AvatarId),
		AccessHash:         int(m.AccessHash),
		ProfilePrivacy:     int(m.ProfilePrivacy),
		OnlinePrivacy:      int(m.OnlinePrivacy),
		CallPrivacy:        int(m.CallPrivacy),
		AddToGroupPrivacy:  int(m.AddToGroupPrivacy),
		SeenMessagePrivacy: int(m.SeenMessagePrivacy),
		Phone:              m.Phone,
		Email:              m.Email,
		About:              m.About,
		PasswordHash:       m.PasswordHash,
		PasswordSalt:       m.PasswordSalt,
		PostSeq:            int(m.PostSeq),
		FollowersCount:     int(m.FollowersCount),
		FollowingCount:     int(m.FollowingCount),
		PostsCount:         int(m.PostsCount),
		MediaCount:         int(m.MediaCount),
		PhotoCount:         int(m.PhotoCount),
		VideoCount:         int(m.VideoCount),
		GifCount:           int(m.GifCount),
		AudioCount:         int(m.AudioCount),
		VoiceCount:         int(m.VoiceCount),
		FileCount:          int(m.FileCount),
		LinkCount:          int(m.LinkCount),
		BoardCount:         int(m.BoardCount),
		PinedCount:         int(m.PinedCount),
		LikesCount:         int(m.LikesCount),
		ResharedCount:      int(m.ResharedCount),
		LastPostTime:       int(m.LastPostTime),
		CreatedTime:        int(m.CreatedTime),
		VersionTime:        int(m.VersionTime),
		IsDeleted:          int(m.IsDeleted),
		IsBanned:           int(m.IsBanned),
	}
	return r
}

//ToPB
func (m *PB_User_Flat) ToPB() *PB_User {
	r := &PB_User{
		UserId:             int32(m.UserId),
		UserName:           m.UserName,
		UserNameLower:      m.UserNameLower,
		FirstName:          m.FirstName,
		LastName:           m.LastName,
		IsVerified:         int32(m.IsVerified),
		AvatarId:           int64(m.AvatarId),
		AccessHash:         int32(m.AccessHash),
		ProfilePrivacy:     int32(m.ProfilePrivacy),
		OnlinePrivacy:      int32(m.OnlinePrivacy),
		CallPrivacy:        int32(m.CallPrivacy),
		AddToGroupPrivacy:  int32(m.AddToGroupPrivacy),
		SeenMessagePrivacy: int32(m.SeenMessagePrivacy),
		Phone:              m.Phone,
		Email:              m.Email,
		About:              m.About,
		PasswordHash:       m.PasswordHash,
		PasswordSalt:       m.PasswordSalt,
		PostSeq:            int32(m.PostSeq),
		FollowersCount:     int32(m.FollowersCount),
		FollowingCount:     int32(m.FollowingCount),
		PostsCount:         int32(m.PostsCount),
		MediaCount:         int32(m.MediaCount),
		PhotoCount:         int32(m.PhotoCount),
		VideoCount:         int32(m.VideoCount),
		GifCount:           int32(m.GifCount),
		AudioCount:         int32(m.AudioCount),
		VoiceCount:         int32(m.VoiceCount),
		FileCount:          int32(m.FileCount),
		LinkCount:          int32(m.LinkCount),
		BoardCount:         int32(m.BoardCount),
		PinedCount:         int32(m.PinedCount),
		LikesCount:         int32(m.LikesCount),
		ResharedCount:      int32(m.ResharedCount),
		LastPostTime:       int32(m.LastPostTime),
		CreatedTime:        int32(m.CreatedTime),
		VersionTime:        int32(m.VersionTime),
		IsDeleted:          int32(m.IsDeleted),
		IsBanned:           int32(m.IsBanned),
	}
	return r
}

//folding
var PB_User__FOlD = &PB_User{
	UserId:             0,
	UserName:           "",
	UserNameLower:      "",
	FirstName:          "",
	LastName:           "",
	IsVerified:         0,
	AvatarId:           0,
	AccessHash:         0,
	ProfilePrivacy:     0,
	OnlinePrivacy:      0,
	CallPrivacy:        0,
	AddToGroupPrivacy:  0,
	SeenMessagePrivacy: 0,
	Phone:              "",
	Email:              "",
	About:              "",
	PasswordHash:       "",
	PasswordSalt:       "",
	PostSeq:            0,
	FollowersCount:     0,
	FollowingCount:     0,
	PostsCount:         0,
	MediaCount:         0,
	PhotoCount:         0,
	VideoCount:         0,
	GifCount:           0,
	AudioCount:         0,
	VoiceCount:         0,
	FileCount:          0,
	LinkCount:          0,
	BoardCount:         0,
	PinedCount:         0,
	LikesCount:         0,
	ResharedCount:      0,
	LastPostTime:       0,
	CreatedTime:        0,
	VersionTime:        0,
	IsDeleted:          0,
	IsBanned:           0,
}

type PB_UserRelation_Flat struct {
	RelNanoId     int
	UserId        int
	PeerUserId    int
	Follwing      int
	Followed      int
	InContacts    int
	MutualContact int
	IsFavorite    int
	Notify        int
}

//ToPB
func (m *PB_UserRelation) ToFlat() *PB_UserRelation_Flat {
	r := &PB_UserRelation_Flat{
		RelNanoId:     int(m.RelNanoId),
		UserId:        int(m.UserId),
		PeerUserId:    int(m.PeerUserId),
		Follwing:      int(m.Follwing),
		Followed:      int(m.Followed),
		InContacts:    int(m.InContacts),
		MutualContact: int(m.MutualContact),
		IsFavorite:    int(m.IsFavorite),
		Notify:        int(m.Notify),
	}
	return r
}

//ToPB
func (m *PB_UserRelation_Flat) ToPB() *PB_UserRelation {
	r := &PB_UserRelation{
		RelNanoId:     int64(m.RelNanoId),
		UserId:        int32(m.UserId),
		PeerUserId:    int32(m.PeerUserId),
		Follwing:      int32(m.Follwing),
		Followed:      int32(m.Followed),
		InContacts:    int32(m.InContacts),
		MutualContact: int32(m.MutualContact),
		IsFavorite:    int32(m.IsFavorite),
		Notify:        int32(m.Notify),
	}
	return r
}

//folding
var PB_UserRelation__FOlD = &PB_UserRelation{
	RelNanoId:     0,
	UserId:        0,
	PeerUserId:    0,
	Follwing:      0,
	Followed:      0,
	InContacts:    0,
	MutualContact: 0,
	IsFavorite:    0,
	Notify:        0,
}

type PB_Chat_Flat struct {
	ChatId                 int
	ChatKey                string
	RoomKey                string
	RoomType               int
	UserId                 int
	PeerUserId             int
	GroupId                int
	HashTagId              int
	Title                  string
	PinTimeMs              int
	FromMsgId              int
	UnseenCount            int
	Seq                    int
	LastMsgId              int
	LastMyMsgStatus        int
	MyLastSeenSeq          int
	MyLastSeenMsgId        int
	PeerLastSeenMsgId      int
	MyLastDeliveredSeq     int
	MyLastDeliveredMsgId   int
	PeerLastDeliveredMsgId int
	IsActive               int
	IsLeft                 int
	IsCreator              int
	IsKicked               int
	IsAdmin                int
	IsDeactivated          int
	IsMuted                int
	MuteUntil              int
	VersionTimeMs          int
	VersionSeq             int
	OrderTime              int
	CreatedTime            int
	DraftText              string
	DratReplyToMsgId       int
}

//ToPB
func (m *PB_Chat) ToFlat() *PB_Chat_Flat {
	r := &PB_Chat_Flat{
		ChatId:                 int(m.ChatId),
		ChatKey:                m.ChatKey,
		RoomKey:                m.RoomKey,
		RoomType:               int(m.RoomType),
		UserId:                 int(m.UserId),
		PeerUserId:             int(m.PeerUserId),
		GroupId:                int(m.GroupId),
		HashTagId:              int(m.HashTagId),
		Title:                  m.Title,
		PinTimeMs:              int(m.PinTimeMs),
		FromMsgId:              int(m.FromMsgId),
		UnseenCount:            int(m.UnseenCount),
		Seq:                    int(m.Seq),
		LastMsgId:              int(m.LastMsgId),
		LastMyMsgStatus:        int(m.LastMyMsgStatus),
		MyLastSeenSeq:          int(m.MyLastSeenSeq),
		MyLastSeenMsgId:        int(m.MyLastSeenMsgId),
		PeerLastSeenMsgId:      int(m.PeerLastSeenMsgId),
		MyLastDeliveredSeq:     int(m.MyLastDeliveredSeq),
		MyLastDeliveredMsgId:   int(m.MyLastDeliveredMsgId),
		PeerLastDeliveredMsgId: int(m.PeerLastDeliveredMsgId),
		IsActive:               int(m.IsActive),
		IsLeft:                 int(m.IsLeft),
		IsCreator:              int(m.IsCreator),
		IsKicked:               int(m.IsKicked),
		IsAdmin:                int(m.IsAdmin),
		IsDeactivated:          int(m.IsDeactivated),
		IsMuted:                int(m.IsMuted),
		MuteUntil:              int(m.MuteUntil),
		VersionTimeMs:          int(m.VersionTimeMs),
		VersionSeq:             int(m.VersionSeq),
		OrderTime:              int(m.OrderTime),
		CreatedTime:            int(m.CreatedTime),
		DraftText:              m.DraftText,
		DratReplyToMsgId:       int(m.DratReplyToMsgId),
	}
	return r
}

//ToPB
func (m *PB_Chat_Flat) ToPB() *PB_Chat {
	r := &PB_Chat{
		ChatId:                 int64(m.ChatId),
		ChatKey:                m.ChatKey,
		RoomKey:                m.RoomKey,
		RoomType:               int32(m.RoomType),
		UserId:                 int32(m.UserId),
		PeerUserId:             int32(m.PeerUserId),
		GroupId:                int64(m.GroupId),
		HashTagId:              int64(m.HashTagId),
		Title:                  m.Title,
		PinTimeMs:              int64(m.PinTimeMs),
		FromMsgId:              int64(m.FromMsgId),
		UnseenCount:            int32(m.UnseenCount),
		Seq:                    int32(m.Seq),
		LastMsgId:              int64(m.LastMsgId),
		LastMyMsgStatus:        int32(m.LastMyMsgStatus),
		MyLastSeenSeq:          int32(m.MyLastSeenSeq),
		MyLastSeenMsgId:        int64(m.MyLastSeenMsgId),
		PeerLastSeenMsgId:      int64(m.PeerLastSeenMsgId),
		MyLastDeliveredSeq:     int32(m.MyLastDeliveredSeq),
		MyLastDeliveredMsgId:   int64(m.MyLastDeliveredMsgId),
		PeerLastDeliveredMsgId: int64(m.PeerLastDeliveredMsgId),
		IsActive:               int32(m.IsActive),
		IsLeft:                 int32(m.IsLeft),
		IsCreator:              int32(m.IsCreator),
		IsKicked:               int32(m.IsKicked),
		IsAdmin:                int32(m.IsAdmin),
		IsDeactivated:          int32(m.IsDeactivated),
		IsMuted:                int32(m.IsMuted),
		MuteUntil:              int32(m.MuteUntil),
		VersionTimeMs:          int64(m.VersionTimeMs),
		VersionSeq:             int32(m.VersionSeq),
		OrderTime:              int32(m.OrderTime),
		CreatedTime:            int32(m.CreatedTime),
		DraftText:              m.DraftText,
		DratReplyToMsgId:       int64(m.DratReplyToMsgId),
	}
	return r
}

//folding
var PB_Chat__FOlD = &PB_Chat{
	ChatId:                 0,
	ChatKey:                "",
	RoomKey:                "",
	RoomType:               0,
	UserId:                 0,
	PeerUserId:             0,
	GroupId:                0,
	HashTagId:              0,
	Title:                  "",
	PinTimeMs:              0,
	FromMsgId:              0,
	UnseenCount:            0,
	Seq:                    0,
	LastMsgId:              0,
	LastMyMsgStatus:        0,
	MyLastSeenSeq:          0,
	MyLastSeenMsgId:        0,
	PeerLastSeenMsgId:      0,
	MyLastDeliveredSeq:     0,
	MyLastDeliveredMsgId:   0,
	PeerLastDeliveredMsgId: 0,
	IsActive:               0,
	IsLeft:                 0,
	IsCreator:              0,
	IsKicked:               0,
	IsAdmin:                0,
	IsDeactivated:          0,
	IsMuted:                0,
	MuteUntil:              0,
	VersionTimeMs:          0,
	VersionSeq:             0,
	OrderTime:              0,
	CreatedTime:            0,
	DraftText:              "",
	DratReplyToMsgId:       0,
}

type PB_ChatDeleted_Flat struct {
	ChatId  int
	RoomKey string
}

//ToPB
func (m *PB_ChatDeleted) ToFlat() *PB_ChatDeleted_Flat {
	r := &PB_ChatDeleted_Flat{
		ChatId:  int(m.ChatId),
		RoomKey: m.RoomKey,
	}
	return r
}

//ToPB
func (m *PB_ChatDeleted_Flat) ToPB() *PB_ChatDeleted {
	r := &PB_ChatDeleted{
		ChatId:  int64(m.ChatId),
		RoomKey: m.RoomKey,
	}
	return r
}

//folding
var PB_ChatDeleted__FOlD = &PB_ChatDeleted{
	ChatId:  0,
	RoomKey: "",
}

type PB_ChatLastMessage_Flat struct {
	ChatIdGroupId string
	LastMsgPb     []byte
}

//ToPB
func (m *PB_ChatLastMessage) ToFlat() *PB_ChatLastMessage_Flat {
	r := &PB_ChatLastMessage_Flat{
		ChatIdGroupId: m.ChatIdGroupId,
		LastMsgPb:     []byte(m.LastMsgPb),
	}
	return r
}

//ToPB
func (m *PB_ChatLastMessage_Flat) ToPB() *PB_ChatLastMessage {
	r := &PB_ChatLastMessage{
		ChatIdGroupId: m.ChatIdGroupId,
		LastMsgPb:     m.LastMsgPb,
	}
	return r
}

//folding
var PB_ChatLastMessage__FOlD = &PB_ChatLastMessage{
	ChatIdGroupId: "",
	LastMsgPb:     []byte{},
}

type PB_ChatUserVersion_Flat struct {
	UserId          int
	ChatId          int
	VersionTimeNano int
}

//ToPB
func (m *PB_ChatUserVersion) ToFlat() *PB_ChatUserVersion_Flat {
	r := &PB_ChatUserVersion_Flat{
		UserId:          int(m.UserId),
		ChatId:          int(m.ChatId),
		VersionTimeNano: int(m.VersionTimeNano),
	}
	return r
}

//ToPB
func (m *PB_ChatUserVersion_Flat) ToPB() *PB_ChatUserVersion {
	r := &PB_ChatUserVersion{
		UserId:          int32(m.UserId),
		ChatId:          int64(m.ChatId),
		VersionTimeNano: int32(m.VersionTimeNano),
	}
	return r
}

//folding
var PB_ChatUserVersion__FOlD = &PB_ChatUserVersion{
	UserId:          0,
	ChatId:          0,
	VersionTimeNano: 0,
}

type PB_Group_Flat struct {
	GroupId         int
	GroupKey        string
	GroupName       string
	UserName        string
	IsSuperGroup    int
	HashTagId       int
	CreatorUserId   int
	GroupPrivacy    int
	HistoryViewAble int
	Seq             int
	LastMsgId       int
	PinedMsgId      int
	AvatarRefId     int
	AvatarCount     int
	About           string
	InviteLinkHash  string
	MembersCount    int
	AdminsCount     int
	ModeratorCounts int
	SortTime        int
	CreatedTime     int
	IsMute          int
	IsDeleted       int
	IsBanned        int
}

//ToPB
func (m *PB_Group) ToFlat() *PB_Group_Flat {
	r := &PB_Group_Flat{
		GroupId:         int(m.GroupId),
		GroupKey:        m.GroupKey,
		GroupName:       m.GroupName,
		UserName:        m.UserName,
		IsSuperGroup:    int(m.IsSuperGroup),
		HashTagId:       int(m.HashTagId),
		CreatorUserId:   int(m.CreatorUserId),
		GroupPrivacy:    int(m.GroupPrivacy),
		HistoryViewAble: int(m.HistoryViewAble),
		Seq:             int(m.Seq),
		LastMsgId:       int(m.LastMsgId),
		PinedMsgId:      int(m.PinedMsgId),
		AvatarRefId:     int(m.AvatarRefId),
		AvatarCount:     int(m.AvatarCount),
		About:           m.About,
		InviteLinkHash:  m.InviteLinkHash,
		MembersCount:    int(m.MembersCount),
		AdminsCount:     int(m.AdminsCount),
		ModeratorCounts: int(m.ModeratorCounts),
		SortTime:        int(m.SortTime),
		CreatedTime:     int(m.CreatedTime),
		IsMute:          int(m.IsMute),
		IsDeleted:       int(m.IsDeleted),
		IsBanned:        int(m.IsBanned),
	}
	return r
}

//ToPB
func (m *PB_Group_Flat) ToPB() *PB_Group {
	r := &PB_Group{
		GroupId:         int64(m.GroupId),
		GroupKey:        m.GroupKey,
		GroupName:       m.GroupName,
		UserName:        m.UserName,
		IsSuperGroup:    int32(m.IsSuperGroup),
		HashTagId:       int64(m.HashTagId),
		CreatorUserId:   int32(m.CreatorUserId),
		GroupPrivacy:    int32(m.GroupPrivacy),
		HistoryViewAble: int32(m.HistoryViewAble),
		Seq:             int64(m.Seq),
		LastMsgId:       int64(m.LastMsgId),
		PinedMsgId:      int64(m.PinedMsgId),
		AvatarRefId:     int64(m.AvatarRefId),
		AvatarCount:     int32(m.AvatarCount),
		About:           m.About,
		InviteLinkHash:  m.InviteLinkHash,
		MembersCount:    int32(m.MembersCount),
		AdminsCount:     int32(m.AdminsCount),
		ModeratorCounts: int32(m.ModeratorCounts),
		SortTime:        int32(m.SortTime),
		CreatedTime:     int32(m.CreatedTime),
		IsMute:          int32(m.IsMute),
		IsDeleted:       int32(m.IsDeleted),
		IsBanned:        int32(m.IsBanned),
	}
	return r
}

//folding
var PB_Group__FOlD = &PB_Group{
	GroupId:         0,
	GroupKey:        "",
	GroupName:       "",
	UserName:        "",
	IsSuperGroup:    0,
	HashTagId:       0,
	CreatorUserId:   0,
	GroupPrivacy:    0,
	HistoryViewAble: 0,
	Seq:             0,
	LastMsgId:       0,
	PinedMsgId:      0,
	AvatarRefId:     0,
	AvatarCount:     0,
	About:           "",
	InviteLinkHash:  "",
	MembersCount:    0,
	AdminsCount:     0,
	ModeratorCounts: 0,
	SortTime:        0,
	CreatedTime:     0,
	IsMute:          0,
	IsDeleted:       0,
	IsBanned:        0,
}

type PB_GroupMember_Flat struct {
	OrderId     int
	GroupId     int
	UserId      int
	ByUserId    int
	GroupRole   int
	CreatedTime int
}

//ToPB
func (m *PB_GroupMember) ToFlat() *PB_GroupMember_Flat {
	r := &PB_GroupMember_Flat{
		OrderId:     int(m.OrderId),
		GroupId:     int(m.GroupId),
		UserId:      int(m.UserId),
		ByUserId:    int(m.ByUserId),
		GroupRole:   int(m.GroupRole),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_GroupMember_Flat) ToPB() *PB_GroupMember {
	r := &PB_GroupMember{
		OrderId:     int64(m.OrderId),
		GroupId:     int64(m.GroupId),
		UserId:      int32(m.UserId),
		ByUserId:    int32(m.ByUserId),
		GroupRole:   int32(m.GroupRole),
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_GroupMember__FOlD = &PB_GroupMember{
	OrderId:     0,
	GroupId:     0,
	UserId:      0,
	ByUserId:    0,
	GroupRole:   0,
	CreatedTime: 0,
}

type PB_GroupOrderdUser_Flat struct {
	OrderId int
	GroupId int
	UserId  int
}

//ToPB
func (m *PB_GroupOrderdUser) ToFlat() *PB_GroupOrderdUser_Flat {
	r := &PB_GroupOrderdUser_Flat{
		OrderId: int(m.OrderId),
		GroupId: int(m.GroupId),
		UserId:  int(m.UserId),
	}
	return r
}

//ToPB
func (m *PB_GroupOrderdUser_Flat) ToPB() *PB_GroupOrderdUser {
	r := &PB_GroupOrderdUser{
		OrderId: int64(m.OrderId),
		GroupId: int64(m.GroupId),
		UserId:  int32(m.UserId),
	}
	return r
}

//folding
var PB_GroupOrderdUser__FOlD = &PB_GroupOrderdUser{
	OrderId: 0,
	GroupId: 0,
	UserId:  0,
}

type PB_GroupPinedMsg_Flat struct {
	MsgId int
	MsgPb []byte
}

//ToPB
func (m *PB_GroupPinedMsg) ToFlat() *PB_GroupPinedMsg_Flat {
	r := &PB_GroupPinedMsg_Flat{
		MsgId: int(m.MsgId),
		MsgPb: []byte(m.MsgPb),
	}
	return r
}

//ToPB
func (m *PB_GroupPinedMsg_Flat) ToPB() *PB_GroupPinedMsg {
	r := &PB_GroupPinedMsg{
		MsgId: int64(m.MsgId),
		MsgPb: m.MsgPb,
	}
	return r
}

//folding
var PB_GroupPinedMsg__FOlD = &PB_GroupPinedMsg{
	MsgId: 0,
	MsgPb: []byte{},
}

type PB_FileMsg_Flat struct {
	Id         int
	AccessHash int
	FileType   int
	Width      int
	Height     int
	Extension  string
	UserId     int
	DataThumb  []byte
	Data       []byte
}

//ToPB
func (m *PB_FileMsg) ToFlat() *PB_FileMsg_Flat {
	r := &PB_FileMsg_Flat{
		Id:         int(m.Id),
		AccessHash: int(m.AccessHash),
		FileType:   int(m.FileType),
		Width:      int(m.Width),
		Height:     int(m.Height),
		Extension:  m.Extension,
		UserId:     int(m.UserId),
		DataThumb:  []byte(m.DataThumb),
		Data:       []byte(m.Data),
	}
	return r
}

//ToPB
func (m *PB_FileMsg_Flat) ToPB() *PB_FileMsg {
	r := &PB_FileMsg{
		Id:         int64(m.Id),
		AccessHash: int32(m.AccessHash),
		FileType:   int32(m.FileType),
		Width:      int32(m.Width),
		Height:     int32(m.Height),
		Extension:  m.Extension,
		UserId:     int32(m.UserId),
		DataThumb:  m.DataThumb,
		Data:       m.Data,
	}
	return r
}

//folding
var PB_FileMsg__FOlD = &PB_FileMsg{
	Id:         0,
	AccessHash: 0,
	FileType:   0,
	Width:      0,
	Height:     0,
	Extension:  "",
	UserId:     0,
	DataThumb:  []byte{},
	Data:       []byte{},
}

type PB_FilePost_Flat struct {
	Id         int
	AccessHash int
	FileType   int
	Width      int
	Height     int
	Extension  string
	UserId     int
	DataThumb  []byte
	Data       []byte
}

//ToPB
func (m *PB_FilePost) ToFlat() *PB_FilePost_Flat {
	r := &PB_FilePost_Flat{
		Id:         int(m.Id),
		AccessHash: int(m.AccessHash),
		FileType:   int(m.FileType),
		Width:      int(m.Width),
		Height:     int(m.Height),
		Extension:  m.Extension,
		UserId:     int(m.UserId),
		DataThumb:  []byte(m.DataThumb),
		Data:       []byte(m.Data),
	}
	return r
}

//ToPB
func (m *PB_FilePost_Flat) ToPB() *PB_FilePost {
	r := &PB_FilePost{
		Id:         int64(m.Id),
		AccessHash: int32(m.AccessHash),
		FileType:   int32(m.FileType),
		Width:      int32(m.Width),
		Height:     int32(m.Height),
		Extension:  m.Extension,
		UserId:     int32(m.UserId),
		DataThumb:  m.DataThumb,
		Data:       m.Data,
	}
	return r
}

//folding
var PB_FilePost__FOlD = &PB_FilePost{
	Id:         0,
	AccessHash: 0,
	FileType:   0,
	Width:      0,
	Height:     0,
	Extension:  "",
	UserId:     0,
	DataThumb:  []byte{},
	Data:       []byte{},
}

type PB_ActionFanout_Flat struct {
	OrderId     int
	ForUserId   int
	ActionId    int
	ActorUserId int
}

//ToPB
func (m *PB_ActionFanout) ToFlat() *PB_ActionFanout_Flat {
	r := &PB_ActionFanout_Flat{
		OrderId:     int(m.OrderId),
		ForUserId:   int(m.ForUserId),
		ActionId:    int(m.ActionId),
		ActorUserId: int(m.ActorUserId),
	}
	return r
}

//ToPB
func (m *PB_ActionFanout_Flat) ToPB() *PB_ActionFanout {
	r := &PB_ActionFanout{
		OrderId:     int64(m.OrderId),
		ForUserId:   int32(m.ForUserId),
		ActionId:    int64(m.ActionId),
		ActorUserId: int32(m.ActorUserId),
	}
	return r
}

//folding
var PB_ActionFanout__FOlD = &PB_ActionFanout{
	OrderId:     0,
	ForUserId:   0,
	ActionId:    0,
	ActorUserId: 0,
}

type PB_HomeFanout_Flat struct {
	OrderId    int
	ForUserId  int
	PostId     int
	PostUserId int
	ResharedId int
}

//ToPB
func (m *PB_HomeFanout) ToFlat() *PB_HomeFanout_Flat {
	r := &PB_HomeFanout_Flat{
		OrderId:    int(m.OrderId),
		ForUserId:  int(m.ForUserId),
		PostId:     int(m.PostId),
		PostUserId: int(m.PostUserId),
		ResharedId: int(m.ResharedId),
	}
	return r
}

//ToPB
func (m *PB_HomeFanout_Flat) ToPB() *PB_HomeFanout {
	r := &PB_HomeFanout{
		OrderId:    int64(m.OrderId),
		ForUserId:  int64(m.ForUserId),
		PostId:     int64(m.PostId),
		PostUserId: int64(m.PostUserId),
		ResharedId: int64(m.ResharedId),
	}
	return r
}

//folding
var PB_HomeFanout__FOlD = &PB_HomeFanout{
	OrderId:    0,
	ForUserId:  0,
	PostId:     0,
	PostUserId: 0,
	ResharedId: 0,
}

type PB_SuggestedTopPosts_Flat struct {
	Id     int
	PostId int
}

//ToPB
func (m *PB_SuggestedTopPosts) ToFlat() *PB_SuggestedTopPosts_Flat {
	r := &PB_SuggestedTopPosts_Flat{
		Id:     int(m.Id),
		PostId: int(m.PostId),
	}
	return r
}

//ToPB
func (m *PB_SuggestedTopPosts_Flat) ToPB() *PB_SuggestedTopPosts {
	r := &PB_SuggestedTopPosts{
		Id:     int64(m.Id),
		PostId: int64(m.PostId),
	}
	return r
}

//folding
var PB_SuggestedTopPosts__FOlD = &PB_SuggestedTopPosts{
	Id:     0,
	PostId: 0,
}

type PB_SuggestedUser_Flat struct {
	Id          int
	UserId      int
	TargetId    int
	Weight      float32
	CreatedTime int
}

//ToPB
func (m *PB_SuggestedUser) ToFlat() *PB_SuggestedUser_Flat {
	r := &PB_SuggestedUser_Flat{
		Id:          int(m.Id),
		UserId:      int(m.UserId),
		TargetId:    int(m.TargetId),
		Weight:      float32(m.Weight),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_SuggestedUser_Flat) ToPB() *PB_SuggestedUser {
	r := &PB_SuggestedUser{
		Id:          int32(m.Id),
		UserId:      int32(m.UserId),
		TargetId:    int32(m.TargetId),
		Weight:      m.Weight,
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_SuggestedUser__FOlD = &PB_SuggestedUser{
	Id:          0,
	UserId:      0,
	TargetId:    0,
	Weight:      0.0,
	CreatedTime: 0,
}

type PB_PushChat_Flat struct {
	PushId            int
	ToUserId          int
	PushTypeId        int
	RoomKey           string
	ChatKey           string
	Seq               int
	UnseenCount       int
	FromHighMessageId int
	ToLowMessageId    int
	MessageId         int
	MessageFileId     int
	MessagePb         []byte
	MessageJson       string
	CreatedTime       int
}

//ToPB
func (m *PB_PushChat) ToFlat() *PB_PushChat_Flat {
	r := &PB_PushChat_Flat{
		PushId:            int(m.PushId),
		ToUserId:          int(m.ToUserId),
		PushTypeId:        int(m.PushTypeId),
		RoomKey:           m.RoomKey,
		ChatKey:           m.ChatKey,
		Seq:               int(m.Seq),
		UnseenCount:       int(m.UnseenCount),
		FromHighMessageId: int(m.FromHighMessageId),
		ToLowMessageId:    int(m.ToLowMessageId),
		MessageId:         int(m.MessageId),
		MessageFileId:     int(m.MessageFileId),
		MessagePb:         []byte(m.MessagePb),
		MessageJson:       m.MessageJson,
		CreatedTime:       int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_PushChat_Flat) ToPB() *PB_PushChat {
	r := &PB_PushChat{
		PushId:            int64(m.PushId),
		ToUserId:          int32(m.ToUserId),
		PushTypeId:        int32(m.PushTypeId),
		RoomKey:           m.RoomKey,
		ChatKey:           m.ChatKey,
		Seq:               int32(m.Seq),
		UnseenCount:       int32(m.UnseenCount),
		FromHighMessageId: int64(m.FromHighMessageId),
		ToLowMessageId:    int64(m.ToLowMessageId),
		MessageId:         int64(m.MessageId),
		MessageFileId:     int64(m.MessageFileId),
		MessagePb:         m.MessagePb,
		MessageJson:       m.MessageJson,
		CreatedTime:       int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_PushChat__FOlD = &PB_PushChat{
	PushId:            0,
	ToUserId:          0,
	PushTypeId:        0,
	RoomKey:           "",
	ChatKey:           "",
	Seq:               0,
	UnseenCount:       0,
	FromHighMessageId: 0,
	ToLowMessageId:    0,
	MessageId:         0,
	MessageFileId:     0,
	MessagePb:         []byte{},
	MessageJson:       "",
	CreatedTime:       0,
}

type PB_HTTPRPCLog_Flat struct {
	Id              int
	Time            string
	MethodFull      string
	MethodParent    string
	UserId          int
	SessionId       string
	StatusCode      int
	InputSize       int
	OutputSize      int
	ReqestJson      string
	ResponseJson    string
	ReqestParamJson string
	ResponseMsgJson string
}

//ToPB
func (m *PB_HTTPRPCLog) ToFlat() *PB_HTTPRPCLog_Flat {
	r := &PB_HTTPRPCLog_Flat{
		Id:              int(m.Id),
		Time:            m.Time,
		MethodFull:      m.MethodFull,
		MethodParent:    m.MethodParent,
		UserId:          int(m.UserId),
		SessionId:       m.SessionId,
		StatusCode:      int(m.StatusCode),
		InputSize:       int(m.InputSize),
		OutputSize:      int(m.OutputSize),
		ReqestJson:      m.ReqestJson,
		ResponseJson:    m.ResponseJson,
		ReqestParamJson: m.ReqestParamJson,
		ResponseMsgJson: m.ResponseMsgJson,
	}
	return r
}

//ToPB
func (m *PB_HTTPRPCLog_Flat) ToPB() *PB_HTTPRPCLog {
	r := &PB_HTTPRPCLog{
		Id:              int32(m.Id),
		Time:            m.Time,
		MethodFull:      m.MethodFull,
		MethodParent:    m.MethodParent,
		UserId:          int32(m.UserId),
		SessionId:       m.SessionId,
		StatusCode:      int32(m.StatusCode),
		InputSize:       int32(m.InputSize),
		OutputSize:      int32(m.OutputSize),
		ReqestJson:      m.ReqestJson,
		ResponseJson:    m.ResponseJson,
		ReqestParamJson: m.ReqestParamJson,
		ResponseMsgJson: m.ResponseMsgJson,
	}
	return r
}

//folding
var PB_HTTPRPCLog__FOlD = &PB_HTTPRPCLog{
	Id:              0,
	Time:            "",
	MethodFull:      "",
	MethodParent:    "",
	UserId:          0,
	SessionId:       "",
	StatusCode:      0,
	InputSize:       0,
	OutputSize:      0,
	ReqestJson:      "",
	ResponseJson:    "",
	ReqestParamJson: "",
	ResponseMsgJson: "",
}

type PB_MetricLog_Flat struct {
	Id           int
	InstanceId   int
	StartFrom    string
	EndTo        string
	StartTime    int
	Duration     string
	MetericsJson string
}

//ToPB
func (m *PB_MetricLog) ToFlat() *PB_MetricLog_Flat {
	r := &PB_MetricLog_Flat{
		Id:           int(m.Id),
		InstanceId:   int(m.InstanceId),
		StartFrom:    m.StartFrom,
		EndTo:        m.EndTo,
		StartTime:    int(m.StartTime),
		Duration:     m.Duration,
		MetericsJson: m.MetericsJson,
	}
	return r
}

//ToPB
func (m *PB_MetricLog_Flat) ToPB() *PB_MetricLog {
	r := &PB_MetricLog{
		Id:           int32(m.Id),
		InstanceId:   int32(m.InstanceId),
		StartFrom:    m.StartFrom,
		EndTo:        m.EndTo,
		StartTime:    int32(m.StartTime),
		Duration:     m.Duration,
		MetericsJson: m.MetericsJson,
	}
	return r
}

//folding
var PB_MetricLog__FOlD = &PB_MetricLog{
	Id:           0,
	InstanceId:   0,
	StartFrom:    "",
	EndTo:        "",
	StartTime:    0,
	Duration:     "",
	MetericsJson: "",
}

type PB_XfileServiceInfoLog_Flat struct {
	Id          int
	InstanceId  int
	Url         string
	CreatedTime string
}

//ToPB
func (m *PB_XfileServiceInfoLog) ToFlat() *PB_XfileServiceInfoLog_Flat {
	r := &PB_XfileServiceInfoLog_Flat{
		Id:          int(m.Id),
		InstanceId:  int(m.InstanceId),
		Url:         m.Url,
		CreatedTime: m.CreatedTime,
	}
	return r
}

//ToPB
func (m *PB_XfileServiceInfoLog_Flat) ToPB() *PB_XfileServiceInfoLog {
	r := &PB_XfileServiceInfoLog{
		Id:          int64(m.Id),
		InstanceId:  int32(m.InstanceId),
		Url:         m.Url,
		CreatedTime: m.CreatedTime,
	}
	return r
}

//folding
var PB_XfileServiceInfoLog__FOlD = &PB_XfileServiceInfoLog{
	Id:          0,
	InstanceId:  0,
	Url:         "",
	CreatedTime: "",
}

type PB_XfileServiceMetricLog_Flat struct {
	Id         int
	InstanceId int
	MetricJson string
}

//ToPB
func (m *PB_XfileServiceMetricLog) ToFlat() *PB_XfileServiceMetricLog_Flat {
	r := &PB_XfileServiceMetricLog_Flat{
		Id:         int(m.Id),
		InstanceId: int(m.InstanceId),
		MetricJson: m.MetricJson,
	}
	return r
}

//ToPB
func (m *PB_XfileServiceMetricLog_Flat) ToPB() *PB_XfileServiceMetricLog {
	r := &PB_XfileServiceMetricLog{
		Id:         int64(m.Id),
		InstanceId: int32(m.InstanceId),
		MetricJson: m.MetricJson,
	}
	return r
}

//folding
var PB_XfileServiceMetricLog__FOlD = &PB_XfileServiceMetricLog{
	Id:         0,
	InstanceId: 0,
	MetricJson: "",
}

type PB_XfileServiceRequestLog_Flat struct {
	Id          int
	LocalSeq    int
	InstanceId  int
	Url         string
	HttpCode    int
	CreatedTime string
}

//ToPB
func (m *PB_XfileServiceRequestLog) ToFlat() *PB_XfileServiceRequestLog_Flat {
	r := &PB_XfileServiceRequestLog_Flat{
		Id:          int(m.Id),
		LocalSeq:    int(m.LocalSeq),
		InstanceId:  int(m.InstanceId),
		Url:         m.Url,
		HttpCode:    int(m.HttpCode),
		CreatedTime: m.CreatedTime,
	}
	return r
}

//ToPB
func (m *PB_XfileServiceRequestLog_Flat) ToPB() *PB_XfileServiceRequestLog {
	r := &PB_XfileServiceRequestLog{
		Id:          int64(m.Id),
		LocalSeq:    int32(m.LocalSeq),
		InstanceId:  int32(m.InstanceId),
		Url:         m.Url,
		HttpCode:    int32(m.HttpCode),
		CreatedTime: m.CreatedTime,
	}
	return r
}

//folding
var PB_XfileServiceRequestLog__FOlD = &PB_XfileServiceRequestLog{
	Id:          0,
	LocalSeq:    0,
	InstanceId:  0,
	Url:         "",
	HttpCode:    0,
	CreatedTime: "",
}

type PB_InvalidateCache_Flat struct {
	OrderId  int
	CacheKey string
}

//ToPB
func (m *PB_InvalidateCache) ToFlat() *PB_InvalidateCache_Flat {
	r := &PB_InvalidateCache_Flat{
		OrderId:  int(m.OrderId),
		CacheKey: m.CacheKey,
	}
	return r
}

//ToPB
func (m *PB_InvalidateCache_Flat) ToPB() *PB_InvalidateCache {
	r := &PB_InvalidateCache{
		OrderId:  int64(m.OrderId),
		CacheKey: m.CacheKey,
	}
	return r
}

//folding
var PB_InvalidateCache__FOlD = &PB_InvalidateCache{
	OrderId:  0,
	CacheKey: "",
}

type PB_MediaView_Flat struct {
}

//ToPB
func (m *PB_MediaView) ToFlat() *PB_MediaView_Flat {
	r := &PB_MediaView_Flat{}
	return r
}

//ToPB
func (m *PB_MediaView_Flat) ToPB() *PB_MediaView {
	r := &PB_MediaView{}
	return r
}

//folding
var PB_MediaView__FOlD = &PB_MediaView{}

type PB_ActionView_Flat struct {
	ActionId              int
	ActorUserId           int
	ActionTypeEnum        int
	PeerUserId            int
	PostId                int
	CommentId             int
	Murmur64Hash          int
	CreatedTime           int
	ActorUserView         PB_UserView
	PostView              PB_PostView
	CommentView           PB_CommentView
	FollowedUserView      PB_UserView
	ContentOwenerUserView PB_UserView
}

//ToPB
func (m *PB_ActionView) ToFlat() *PB_ActionView_Flat {
	r := &PB_ActionView_Flat{
		ActionId:       int(m.ActionId),
		ActorUserId:    int(m.ActorUserId),
		ActionTypeEnum: int(m.ActionTypeEnum),
		PeerUserId:     int(m.PeerUserId),
		PostId:         int(m.PostId),
		CommentId:      int(m.CommentId),
		Murmur64Hash:   int(m.Murmur64Hash),
		CreatedTime:    int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_ActionView_Flat) ToPB() *PB_ActionView {
	r := &PB_ActionView{
		ActionId:       int64(m.ActionId),
		ActorUserId:    int32(m.ActorUserId),
		ActionTypeEnum: int32(m.ActionTypeEnum),
		PeerUserId:     int32(m.PeerUserId),
		PostId:         int64(m.PostId),
		CommentId:      int64(m.CommentId),
		Murmur64Hash:   int64(m.Murmur64Hash),
		CreatedTime:    int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_ActionView__FOlD = &PB_ActionView{
	ActionId:       0,
	ActorUserId:    0,
	ActionTypeEnum: 0,
	PeerUserId:     0,
	PostId:         0,
	CommentId:      0,
	Murmur64Hash:   0,
	CreatedTime:    0,
}

type PB_NotifyView_Flat struct {
	NotifyId      int
	ForUserId     int
	ActorUserId   int
	NotiyTypeEnum int
	PostId        int
	CommentId     int
	PeerUserId    int
	Murmur64Hash  int
	SeenStatus    int
	CreatedTime   int
	ActorUserView PB_UserView
	PostView      PB_PostView
	CommentView   PB_CommentView
}

//ToPB
func (m *PB_NotifyView) ToFlat() *PB_NotifyView_Flat {
	r := &PB_NotifyView_Flat{
		NotifyId:      int(m.NotifyId),
		ForUserId:     int(m.ForUserId),
		ActorUserId:   int(m.ActorUserId),
		NotiyTypeEnum: int(m.NotiyTypeEnum),
		PostId:        int(m.PostId),
		CommentId:     int(m.CommentId),
		PeerUserId:    int(m.PeerUserId),
		Murmur64Hash:  int(m.Murmur64Hash),
		SeenStatus:    int(m.SeenStatus),
		CreatedTime:   int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_NotifyView_Flat) ToPB() *PB_NotifyView {
	r := &PB_NotifyView{
		NotifyId:      int64(m.NotifyId),
		ForUserId:     int32(m.ForUserId),
		ActorUserId:   int32(m.ActorUserId),
		NotiyTypeEnum: int32(m.NotiyTypeEnum),
		PostId:        int64(m.PostId),
		CommentId:     int64(m.CommentId),
		PeerUserId:    int32(m.PeerUserId),
		Murmur64Hash:  int64(m.Murmur64Hash),
		SeenStatus:    int32(m.SeenStatus),
		CreatedTime:   int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_NotifyView__FOlD = &PB_NotifyView{
	NotifyId:      0,
	ForUserId:     0,
	ActorUserId:   0,
	NotiyTypeEnum: 0,
	PostId:        0,
	CommentId:     0,
	PeerUserId:    0,
	Murmur64Hash:  0,
	SeenStatus:    0,
	CreatedTime:   0,
}

type PB_CommentView_Flat struct {
	CommentId      int
	UserId         int
	PostId         int
	Text           string
	LikesCount     int
	CreatedTime    int
	SenderUserView PB_UserView
}

//ToPB
func (m *PB_CommentView) ToFlat() *PB_CommentView_Flat {
	r := &PB_CommentView_Flat{
		CommentId:   int(m.CommentId),
		UserId:      int(m.UserId),
		PostId:      int(m.PostId),
		Text:        m.Text,
		LikesCount:  int(m.LikesCount),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_CommentView_Flat) ToPB() *PB_CommentView {
	r := &PB_CommentView{
		CommentId:   int64(m.CommentId),
		UserId:      int32(m.UserId),
		PostId:      int64(m.PostId),
		Text:        m.Text,
		LikesCount:  int32(m.LikesCount),
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_CommentView__FOlD = &PB_CommentView{
	CommentId:   0,
	UserId:      0,
	PostId:      0,
	Text:        "",
	LikesCount:  0,
	CreatedTime: 0,
}

type PB_PostView_Flat struct {
	PostId           int
	UserId           int
	PostTypeEnum     PostTypeEnum
	Text             string
	RichText         string
	MediaCount       int
	SharedTo         int
	DisableComment   int
	HasTag           int
	CommentsCount    int
	LikesCount       int
	ViewsCount       int
	EditedTime       int
	CreatedTime      int
	ReSharedPostId   int
	DidILiked        bool
	DidIReShared     bool
	SenderUserView   PB_UserView
	ReSharedUserView PB_UserView
	MediaView        PB_MediaView
	MediaViewList    []PB_MediaView
}

//ToPB
func (m *PB_PostView) ToFlat() *PB_PostView_Flat {
	r := &PB_PostView_Flat{
		PostId: int(m.PostId),
		UserId: int(m.UserId),

		Text:           m.Text,
		RichText:       m.RichText,
		MediaCount:     int(m.MediaCount),
		SharedTo:       int(m.SharedTo),
		DisableComment: int(m.DisableComment),
		HasTag:         int(m.HasTag),
		CommentsCount:  int(m.CommentsCount),
		LikesCount:     int(m.LikesCount),
		ViewsCount:     int(m.ViewsCount),
		EditedTime:     int(m.EditedTime),
		CreatedTime:    int(m.CreatedTime),
		ReSharedPostId: int(m.ReSharedPostId),
		DidILiked:      m.DidILiked,
		DidIReShared:   m.DidIReShared,
	}
	return r
}

//ToPB
func (m *PB_PostView_Flat) ToPB() *PB_PostView {
	r := &PB_PostView{
		PostId: int64(m.PostId),
		UserId: int32(m.UserId),

		Text:           m.Text,
		RichText:       m.RichText,
		MediaCount:     int32(m.MediaCount),
		SharedTo:       int32(m.SharedTo),
		DisableComment: int32(m.DisableComment),
		HasTag:         int32(m.HasTag),
		CommentsCount:  int32(m.CommentsCount),
		LikesCount:     int32(m.LikesCount),
		ViewsCount:     int32(m.ViewsCount),
		EditedTime:     int32(m.EditedTime),
		CreatedTime:    int32(m.CreatedTime),
		ReSharedPostId: int64(m.ReSharedPostId),
		DidILiked:      m.DidILiked,
		DidIReShared:   m.DidIReShared,
	}
	return r
}

//folding
var PB_PostView__FOlD = &PB_PostView{
	PostId: 0,
	UserId: 0,

	Text:           "",
	RichText:       "",
	MediaCount:     0,
	SharedTo:       0,
	DisableComment: 0,
	HasTag:         0,
	CommentsCount:  0,
	LikesCount:     0,
	ViewsCount:     0,
	EditedTime:     0,
	CreatedTime:    0,
	ReSharedPostId: 0,
	DidILiked:      false,
	DidIReShared:   false,
}

type PB_ChatView_Flat struct {
	ChatId             int
	ChatKey            string
	RoomKey            string
	RoomType           int
	UserId             int
	PeerUserId         int
	GroupId            int
	HashTagId          int
	StartedByMe        int
	Title              string
	PinTime            int
	FromMsgId          int
	Seq                int
	LastMsgId          int
	LastMsgStatus      int
	SeenSeq            int
	SeenMsgId          int
	Left               int
	Creator            int
	Kicked             int
	Admin              int
	Deactivated        int
	VersionTime        int
	SortTime           int
	CreatedTime        int
	DraftText          string
	DratReplyToMsgId   int
	IsMute             int
	UserView           PB_UserView
	GroupView          PB_GroupView
	FirstUnreadMessage PB_MessageView
	LastMessage        PB_MessageView
}

//ToPB
func (m *PB_ChatView) ToFlat() *PB_ChatView_Flat {
	r := &PB_ChatView_Flat{
		ChatId:           int(m.ChatId),
		ChatKey:          m.ChatKey,
		RoomKey:          m.RoomKey,
		RoomType:         int(m.RoomType),
		UserId:           int(m.UserId),
		PeerUserId:       int(m.PeerUserId),
		GroupId:          int(m.GroupId),
		HashTagId:        int(m.HashTagId),
		StartedByMe:      int(m.StartedByMe),
		Title:            m.Title,
		PinTime:          int(m.PinTime),
		FromMsgId:        int(m.FromMsgId),
		Seq:              int(m.Seq),
		LastMsgId:        int(m.LastMsgId),
		LastMsgStatus:    int(m.LastMsgStatus),
		SeenSeq:          int(m.SeenSeq),
		SeenMsgId:        int(m.SeenMsgId),
		Left:             int(m.Left),
		Creator:          int(m.Creator),
		Kicked:           int(m.Kicked),
		Admin:            int(m.Admin),
		Deactivated:      int(m.Deactivated),
		VersionTime:      int(m.VersionTime),
		SortTime:         int(m.SortTime),
		CreatedTime:      int(m.CreatedTime),
		DraftText:        m.DraftText,
		DratReplyToMsgId: int(m.DratReplyToMsgId),
		IsMute:           int(m.IsMute),
	}
	return r
}

//ToPB
func (m *PB_ChatView_Flat) ToPB() *PB_ChatView {
	r := &PB_ChatView{
		ChatId:           int64(m.ChatId),
		ChatKey:          m.ChatKey,
		RoomKey:          m.RoomKey,
		RoomType:         int32(m.RoomType),
		UserId:           int32(m.UserId),
		PeerUserId:       int32(m.PeerUserId),
		GroupId:          int64(m.GroupId),
		HashTagId:        int64(m.HashTagId),
		StartedByMe:      int32(m.StartedByMe),
		Title:            m.Title,
		PinTime:          int64(m.PinTime),
		FromMsgId:        int64(m.FromMsgId),
		Seq:              int32(m.Seq),
		LastMsgId:        int64(m.LastMsgId),
		LastMsgStatus:    int32(m.LastMsgStatus),
		SeenSeq:          int32(m.SeenSeq),
		SeenMsgId:        int64(m.SeenMsgId),
		Left:             int32(m.Left),
		Creator:          int32(m.Creator),
		Kicked:           int32(m.Kicked),
		Admin:            int32(m.Admin),
		Deactivated:      int32(m.Deactivated),
		VersionTime:      int32(m.VersionTime),
		SortTime:         int32(m.SortTime),
		CreatedTime:      int32(m.CreatedTime),
		DraftText:        m.DraftText,
		DratReplyToMsgId: int64(m.DratReplyToMsgId),
		IsMute:           int32(m.IsMute),
	}
	return r
}

//folding
var PB_ChatView__FOlD = &PB_ChatView{
	ChatId:           0,
	ChatKey:          "",
	RoomKey:          "",
	RoomType:         0,
	UserId:           0,
	PeerUserId:       0,
	GroupId:          0,
	HashTagId:        0,
	StartedByMe:      0,
	Title:            "",
	PinTime:          0,
	FromMsgId:        0,
	Seq:              0,
	LastMsgId:        0,
	LastMsgStatus:    0,
	SeenSeq:          0,
	SeenMsgId:        0,
	Left:             0,
	Creator:          0,
	Kicked:           0,
	Admin:            0,
	Deactivated:      0,
	VersionTime:      0,
	SortTime:         0,
	CreatedTime:      0,
	DraftText:        "",
	DratReplyToMsgId: 0,
	IsMute:           0,
}

type PB_GroupView_Flat struct {
	GroupId         int
	GroupKey        string
	GroupName       string
	UserName        string
	IsSuperGroup    int
	HashTagId       int
	CreatorUserId   int
	GroupPrivacy    int
	HistoryViewAble int
	Seq             int
	LastMsgId       int
	PinedMsgId      int
	AvatarRefId     int
	AvatarCount     int
	About           string
	InviteLink      string
	MembersCount    int
	SortTime        int
	CreatedTime     int
}

//ToPB
func (m *PB_GroupView) ToFlat() *PB_GroupView_Flat {
	r := &PB_GroupView_Flat{
		GroupId:         int(m.GroupId),
		GroupKey:        m.GroupKey,
		GroupName:       m.GroupName,
		UserName:        m.UserName,
		IsSuperGroup:    int(m.IsSuperGroup),
		HashTagId:       int(m.HashTagId),
		CreatorUserId:   int(m.CreatorUserId),
		GroupPrivacy:    int(m.GroupPrivacy),
		HistoryViewAble: int(m.HistoryViewAble),
		Seq:             int(m.Seq),
		LastMsgId:       int(m.LastMsgId),
		PinedMsgId:      int(m.PinedMsgId),
		AvatarRefId:     int(m.AvatarRefId),
		AvatarCount:     int(m.AvatarCount),
		About:           m.About,
		InviteLink:      m.InviteLink,
		MembersCount:    int(m.MembersCount),
		SortTime:        int(m.SortTime),
		CreatedTime:     int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_GroupView_Flat) ToPB() *PB_GroupView {
	r := &PB_GroupView{
		GroupId:         int64(m.GroupId),
		GroupKey:        m.GroupKey,
		GroupName:       m.GroupName,
		UserName:        m.UserName,
		IsSuperGroup:    int32(m.IsSuperGroup),
		HashTagId:       int64(m.HashTagId),
		CreatorUserId:   int32(m.CreatorUserId),
		GroupPrivacy:    int32(m.GroupPrivacy),
		HistoryViewAble: int32(m.HistoryViewAble),
		Seq:             int64(m.Seq),
		LastMsgId:       int64(m.LastMsgId),
		PinedMsgId:      int64(m.PinedMsgId),
		AvatarRefId:     int64(m.AvatarRefId),
		AvatarCount:     int32(m.AvatarCount),
		About:           m.About,
		InviteLink:      m.InviteLink,
		MembersCount:    int32(m.MembersCount),
		SortTime:        int32(m.SortTime),
		CreatedTime:     int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_GroupView__FOlD = &PB_GroupView{
	GroupId:         0,
	GroupKey:        "",
	GroupName:       "",
	UserName:        "",
	IsSuperGroup:    0,
	HashTagId:       0,
	CreatorUserId:   0,
	GroupPrivacy:    0,
	HistoryViewAble: 0,
	Seq:             0,
	LastMsgId:       0,
	PinedMsgId:      0,
	AvatarRefId:     0,
	AvatarCount:     0,
	About:           "",
	InviteLink:      "",
	MembersCount:    0,
	SortTime:        0,
	CreatedTime:     0,
}

type PB_GroupMemberView_Flat struct {
	OrderId     int
	GroupId     int
	UserId      int
	ByUserId    int
	GroupRole   int
	CreatedTime int
}

//ToPB
func (m *PB_GroupMemberView) ToFlat() *PB_GroupMemberView_Flat {
	r := &PB_GroupMemberView_Flat{
		OrderId:     int(m.OrderId),
		GroupId:     int(m.GroupId),
		UserId:      int(m.UserId),
		ByUserId:    int(m.ByUserId),
		GroupRole:   int(m.GroupRole),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_GroupMemberView_Flat) ToPB() *PB_GroupMemberView {
	r := &PB_GroupMemberView{
		OrderId:     int64(m.OrderId),
		GroupId:     int64(m.GroupId),
		UserId:      int32(m.UserId),
		ByUserId:    int32(m.ByUserId),
		GroupRole:   int32(m.GroupRole),
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_GroupMemberView__FOlD = &PB_GroupMemberView{
	OrderId:     0,
	GroupId:     0,
	UserId:      0,
	ByUserId:    0,
	GroupRole:   0,
	CreatedTime: 0,
}

type PB_MessageView_Flat struct {
	RoomKey          string
	MessageId        int
	UserId           int
	FileRefId        int
	MessageType      int
	Text             string
	Hiden            int
	Seq              int
	ForwardedMsgId   int
	PostId           int
	StickerId        int
	CreatedTime      int
	DeliveredTime    int
	SeenTime         int
	DeliviryStatus   int
	ReplyToMessageId int
	ViewsCount       int
	EditTime         int
	Ttl              int
	FileRedView      PB_FileRedView
}

//ToPB
func (m *PB_MessageView) ToFlat() *PB_MessageView_Flat {
	r := &PB_MessageView_Flat{
		RoomKey:          m.RoomKey,
		MessageId:        int(m.MessageId),
		UserId:           int(m.UserId),
		FileRefId:        int(m.FileRefId),
		MessageType:      int(m.MessageType),
		Text:             m.Text,
		Hiden:            int(m.Hiden),
		Seq:              int(m.Seq),
		ForwardedMsgId:   int(m.ForwardedMsgId),
		PostId:           int(m.PostId),
		StickerId:        int(m.StickerId),
		CreatedTime:      int(m.CreatedTime),
		DeliveredTime:    int(m.DeliveredTime),
		SeenTime:         int(m.SeenTime),
		DeliviryStatus:   int(m.DeliviryStatus),
		ReplyToMessageId: int(m.ReplyToMessageId),
		ViewsCount:       int(m.ViewsCount),
		EditTime:         int(m.EditTime),
		Ttl:              int(m.Ttl),
	}
	return r
}

//ToPB
func (m *PB_MessageView_Flat) ToPB() *PB_MessageView {
	r := &PB_MessageView{
		RoomKey:          m.RoomKey,
		MessageId:        int64(m.MessageId),
		UserId:           int32(m.UserId),
		FileRefId:        int64(m.FileRefId),
		MessageType:      int32(m.MessageType),
		Text:             m.Text,
		Hiden:            int32(m.Hiden),
		Seq:              int32(m.Seq),
		ForwardedMsgId:   int64(m.ForwardedMsgId),
		PostId:           int64(m.PostId),
		StickerId:        int64(m.StickerId),
		CreatedTime:      int32(m.CreatedTime),
		DeliveredTime:    int32(m.DeliveredTime),
		SeenTime:         int32(m.SeenTime),
		DeliviryStatus:   int32(m.DeliviryStatus),
		ReplyToMessageId: int64(m.ReplyToMessageId),
		ViewsCount:       int64(m.ViewsCount),
		EditTime:         int32(m.EditTime),
		Ttl:              int32(m.Ttl),
	}
	return r
}

//folding
var PB_MessageView__FOlD = &PB_MessageView{
	RoomKey:          "",
	MessageId:        0,
	UserId:           0,
	FileRefId:        0,
	MessageType:      0,
	Text:             "",
	Hiden:            0,
	Seq:              0,
	ForwardedMsgId:   0,
	PostId:           0,
	StickerId:        0,
	CreatedTime:      0,
	DeliveredTime:    0,
	SeenTime:         0,
	DeliviryStatus:   0,
	ReplyToMessageId: 0,
	ViewsCount:       0,
	EditTime:         0,
	Ttl:              0,
}

type PB_FileRedView_Flat struct {
	FileRefId int
	UserId    int
	Name      string
	Width     int
	Height    int
	Duration  int
	Extension string
	UrlSource string
}

//ToPB
func (m *PB_FileRedView) ToFlat() *PB_FileRedView_Flat {
	r := &PB_FileRedView_Flat{
		FileRefId: int(m.FileRefId),
		UserId:    int(m.UserId),
		Name:      m.Name,
		Width:     int(m.Width),
		Height:    int(m.Height),
		Duration:  int(m.Duration),
		Extension: m.Extension,
		UrlSource: m.UrlSource,
	}
	return r
}

//ToPB
func (m *PB_FileRedView_Flat) ToPB() *PB_FileRedView {
	r := &PB_FileRedView{
		FileRefId: int64(m.FileRefId),
		UserId:    int64(m.UserId),
		Name:      m.Name,
		Width:     int32(m.Width),
		Height:    int32(m.Height),
		Duration:  int32(m.Duration),
		Extension: m.Extension,
		UrlSource: m.UrlSource,
	}
	return r
}

//folding
var PB_FileRedView__FOlD = &PB_FileRedView{
	FileRefId: 0,
	UserId:    0,
	Name:      "",
	Width:     0,
	Height:    0,
	Duration:  0,
	Extension: "",
	UrlSource: "",
}

type PB_UserView_Flat struct {
	UserId               int
	UserName             string
	FirstName            string
	LastName             string
	AvatarRefId          int
	ProfilePrivacy       int
	Phone                int
	About                string
	FollowersCount       int
	FollowingCount       int
	PostsCount           int
	MediaCount           int
	UserOnlineStatusEnum UserOnlineStatusEnum
	LastActiveTime       int
	LastActiveTimeShow   string
	MyFollwing           FollowingEnum
}

//ToPB
func (m *PB_UserView) ToFlat() *PB_UserView_Flat {
	r := &PB_UserView_Flat{
		UserId:         int(m.UserId),
		UserName:       m.UserName,
		FirstName:      m.FirstName,
		LastName:       m.LastName,
		AvatarRefId:    int(m.AvatarRefId),
		ProfilePrivacy: int(m.ProfilePrivacy),
		Phone:          int(m.Phone),
		About:          m.About,
		FollowersCount: int(m.FollowersCount),
		FollowingCount: int(m.FollowingCount),
		PostsCount:     int(m.PostsCount),
		MediaCount:     int(m.MediaCount),

		LastActiveTime:     int(m.LastActiveTime),
		LastActiveTimeShow: m.LastActiveTimeShow,
	}
	return r
}

//ToPB
func (m *PB_UserView_Flat) ToPB() *PB_UserView {
	r := &PB_UserView{
		UserId:         int32(m.UserId),
		UserName:       m.UserName,
		FirstName:      m.FirstName,
		LastName:       m.LastName,
		AvatarRefId:    int64(m.AvatarRefId),
		ProfilePrivacy: int32(m.ProfilePrivacy),
		Phone:          int64(m.Phone),
		About:          m.About,
		FollowersCount: int32(m.FollowersCount),
		FollowingCount: int32(m.FollowingCount),
		PostsCount:     int32(m.PostsCount),
		MediaCount:     int32(m.MediaCount),

		LastActiveTime:     int32(m.LastActiveTime),
		LastActiveTimeShow: m.LastActiveTimeShow,
	}
	return r
}

//folding
var PB_UserView__FOlD = &PB_UserView{
	UserId:         0,
	UserName:       "",
	FirstName:      "",
	LastName:       "",
	AvatarRefId:    0,
	ProfilePrivacy: 0,
	Phone:          0,
	About:          "",
	FollowersCount: 0,
	FollowingCount: 0,
	PostsCount:     0,
	MediaCount:     0,

	LastActiveTime:     0,
	LastActiveTimeShow: "",
}

type PB_SettingNotificationView_Flat struct {
}

//ToPB
func (m *PB_SettingNotificationView) ToFlat() *PB_SettingNotificationView_Flat {
	r := &PB_SettingNotificationView_Flat{}
	return r
}

//ToPB
func (m *PB_SettingNotificationView_Flat) ToPB() *PB_SettingNotificationView {
	r := &PB_SettingNotificationView{}
	return r
}

//folding
var PB_SettingNotificationView__FOlD = &PB_SettingNotificationView{}

type PB_AppConfig_Flat struct {
	DeprecatedClient bool
	HasNewUpdate     bool
}

//ToPB
func (m *PB_AppConfig) ToFlat() *PB_AppConfig_Flat {
	r := &PB_AppConfig_Flat{
		DeprecatedClient: m.DeprecatedClient,
		HasNewUpdate:     m.HasNewUpdate,
	}
	return r
}

//ToPB
func (m *PB_AppConfig_Flat) ToPB() *PB_AppConfig {
	r := &PB_AppConfig{
		DeprecatedClient: m.DeprecatedClient,
		HasNewUpdate:     m.HasNewUpdate,
	}
	return r
}

//folding
var PB_AppConfig__FOlD = &PB_AppConfig{
	DeprecatedClient: false,
	HasNewUpdate:     false,
}

type PB_UserProfileView_Flat struct {
}

//ToPB
func (m *PB_UserProfileView) ToFlat() *PB_UserProfileView_Flat {
	r := &PB_UserProfileView_Flat{}
	return r
}

//ToPB
func (m *PB_UserProfileView_Flat) ToPB() *PB_UserProfileView {
	r := &PB_UserProfileView{}
	return r
}

//folding
var PB_UserProfileView__FOlD = &PB_UserProfileView{}

type PB_UserViewRowify_Flat struct {
	Id          int
	CreatedTime int
	UserView    PB_UserView
}

//ToPB
func (m *PB_UserViewRowify) ToFlat() *PB_UserViewRowify_Flat {
	r := &PB_UserViewRowify_Flat{
		Id:          int(m.Id),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_UserViewRowify_Flat) ToPB() *PB_UserViewRowify {
	r := &PB_UserViewRowify{
		Id:          int64(m.Id),
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_UserViewRowify__FOlD = &PB_UserViewRowify{
	Id:          0,
	CreatedTime: 0,
}

type PB_PostViewRowify_Flat struct {
	Id       int
	PostView PB_PostView
}

//ToPB
func (m *PB_PostViewRowify) ToFlat() *PB_PostViewRowify_Flat {
	r := &PB_PostViewRowify_Flat{
		Id: int(m.Id),
	}
	return r
}

//ToPB
func (m *PB_PostViewRowify_Flat) ToPB() *PB_PostViewRowify {
	r := &PB_PostViewRowify{
		Id: int64(m.Id),
	}
	return r
}

//folding
var PB_PostViewRowify__FOlD = &PB_PostViewRowify{
	Id: 0,
}

type PB_TagView_Flat struct {
	TagId         int
	Name          string
	Count         int
	TagStatusEnum int
	CreatedTime   int
}

//ToPB
func (m *PB_TagView) ToFlat() *PB_TagView_Flat {
	r := &PB_TagView_Flat{
		TagId:         int(m.TagId),
		Name:          m.Name,
		Count:         int(m.Count),
		TagStatusEnum: int(m.TagStatusEnum),
		CreatedTime:   int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_TagView_Flat) ToPB() *PB_TagView {
	r := &PB_TagView{
		TagId:         int64(m.TagId),
		Name:          m.Name,
		Count:         int32(m.Count),
		TagStatusEnum: int32(m.TagStatusEnum),
		CreatedTime:   int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_TagView__FOlD = &PB_TagView{
	TagId:         0,
	Name:          "",
	Count:         0,
	TagStatusEnum: 0,
	CreatedTime:   0,
}

type PB_TopTagWithSamplePosts_Flat struct {
	TagView      PB_TagView
	PostViewList []PB_PostView
}

//ToPB
func (m *PB_TopTagWithSamplePosts) ToFlat() *PB_TopTagWithSamplePosts_Flat {
	r := &PB_TopTagWithSamplePosts_Flat{}
	return r
}

//ToPB
func (m *PB_TopTagWithSamplePosts_Flat) ToPB() *PB_TopTagWithSamplePosts {
	r := &PB_TopTagWithSamplePosts{}
	return r
}

//folding
var PB_TopTagWithSamplePosts__FOlD = &PB_TopTagWithSamplePosts{}

type PB_SelfUserView_Flat struct {
	UserView            PB_UserView
	ProfilePrivacy      int
	OnlinePrivacy       int
	CallPrivacy         int
	AddToGroupPrivacy   int
	SeenMessagePrivacy  int
	SettingNotification PB_SettingNotificationView
}

//ToPB
func (m *PB_SelfUserView) ToFlat() *PB_SelfUserView_Flat {
	r := &PB_SelfUserView_Flat{

		ProfilePrivacy:     int(m.ProfilePrivacy),
		OnlinePrivacy:      int(m.OnlinePrivacy),
		CallPrivacy:        int(m.CallPrivacy),
		AddToGroupPrivacy:  int(m.AddToGroupPrivacy),
		SeenMessagePrivacy: int(m.SeenMessagePrivacy),
	}
	return r
}

//ToPB
func (m *PB_SelfUserView_Flat) ToPB() *PB_SelfUserView {
	r := &PB_SelfUserView{

		ProfilePrivacy:     int32(m.ProfilePrivacy),
		OnlinePrivacy:      int32(m.OnlinePrivacy),
		CallPrivacy:        int32(m.CallPrivacy),
		AddToGroupPrivacy:  int32(m.AddToGroupPrivacy),
		SeenMessagePrivacy: int32(m.SeenMessagePrivacy),
	}
	return r
}

//folding
var PB_SelfUserView__FOlD = &PB_SelfUserView{

	ProfilePrivacy:     0,
	OnlinePrivacy:      0,
	CallPrivacy:        0,
	AddToGroupPrivacy:  0,
	SeenMessagePrivacy: 0,
}

type PB_Error_Flat struct {
	Error        ServerErrors
	ShowError    bool
	ErrorMessage string
}

//ToPB
func (m *PB_Error) ToFlat() *PB_Error_Flat {
	r := &PB_Error_Flat{

		ShowError:    m.ShowError,
		ErrorMessage: m.ErrorMessage,
	}
	return r
}

//ToPB
func (m *PB_Error_Flat) ToPB() *PB_Error {
	r := &PB_Error{

		ShowError:    m.ShowError,
		ErrorMessage: m.ErrorMessage,
	}
	return r
}

//folding
var PB_Error__FOlD = &PB_Error{

	ShowError:    false,
	ErrorMessage: "",
}

/*
///// to_flat ///

func(m *PB_RoomsChanges)ToFlat() *PB_RoomsChanges_Flat {
r := &PB_RoomsChanges_Flat{
    VersionTime:  int(m.VersionTime) ,


}
return r
}

func(m *PB_PushChanges)ToFlat() *PB_PushChanges_Flat {
r := &PB_PushChanges_Flat{


    InvalidateCacheForRoomKeys:  m.InvalidateCacheForRoomKeys ,
    InvalidateAllChatCache:  int(m.InvalidateAllChatCache) ,
    InvalidateAllSocialCache:  int(m.InvalidateAllSocialCache) ,
}
return r
}

func(m *PB_CommandToServer)ToFlat() *PB_CommandToServer_Flat {
r := &PB_CommandToServer_Flat{
    ClientCallId:  int(m.ClientCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *PB_CommandToClient)ToFlat() *PB_CommandToClient_Flat {
r := &PB_CommandToClient_Flat{
    ServerCallId:  int(m.ServerCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *PB_CommandReachedToServer)ToFlat() *PB_CommandReachedToServer_Flat {
r := &PB_CommandReachedToServer_Flat{
    ClientCallId:  int(m.ClientCallId) ,
}
return r
}

func(m *PB_CommandReachedToClient)ToFlat() *PB_CommandReachedToClient_Flat {
r := &PB_CommandReachedToClient_Flat{
    ServerCallId:  int(m.ServerCallId) ,
}
return r
}

func(m *PB_ResponseToClient)ToFlat() *PB_ResponseToClient_Flat {
r := &PB_ResponseToClient_Flat{
    ClientCallId:  int(m.ClientCallId) ,
    PBClass:  m.PBClass ,
    RpcFullName:  m.RpcFullName ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *RPC_Auth_Types)ToFlat() *RPC_Auth_Types_Flat {
r := &RPC_Auth_Types_Flat{
}
return r
}

func(m *PB_RPC_Chat_Types)ToFlat() *PB_RPC_Chat_Types_Flat {
r := &PB_RPC_Chat_Types_Flat{
}
return r
}

func(m *RPC_General_Types)ToFlat() *RPC_General_Types_Flat {
r := &RPC_General_Types_Flat{
}
return r
}

func(m *RPC_Page_Types)ToFlat() *RPC_Page_Types_Flat {
r := &RPC_Page_Types_Flat{
}
return r
}

func(m *RPC_Social_Types)ToFlat() *RPC_Social_Types_Flat {
r := &RPC_Social_Types_Flat{
}
return r
}

func(m *RPC_User_Types)ToFlat() *RPC_User_Types_Flat {
r := &RPC_User_Types_Flat{
}
return r
}

func(m *PB_Action)ToFlat() *PB_Action_Flat {
r := &PB_Action_Flat{
    ActionId:  int(m.ActionId) ,
    ActorUserId:  int(m.ActorUserId) ,
    ActionType:  int(m.ActionType) ,
    PeerUserId:  int(m.PeerUserId) ,
    PostId:  int(m.PostId) ,
    CommentId:  int(m.CommentId) ,
    Murmur64Hash:  int(m.Murmur64Hash) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_Blocked)ToFlat() *PB_Blocked_Flat {
r := &PB_Blocked_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    BlockedUserId:  int(m.BlockedUserId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_Comment)ToFlat() *PB_Comment_Flat {
r := &PB_Comment_Flat{
    CommentId:  int(m.CommentId) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    Text:  m.Text ,
    LikesCount:  int(m.LikesCount) ,
    IsEdited:  int(m.IsEdited) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_CommentDeleted)ToFlat() *PB_CommentDeleted_Flat {
r := &PB_CommentDeleted_Flat{
    CommentId:  int(m.CommentId) ,
    UserId:  int(m.UserId) ,
}
return r
}

func(m *PB_Event)ToFlat() *PB_Event_Flat {
r := &PB_Event_Flat{
    EventId:  int(m.EventId) ,
    EventType:  int(m.EventType) ,
    ByUserId:  int(m.ByUserId) ,
    PeerUserId:  int(m.PeerUserId) ,
    PostId:  int(m.PostId) ,
    CommentId:  int(m.CommentId) ,
    HashTagId:  int(m.HashTagId) ,
    GroupId:  int(m.GroupId) ,
    ActionId:  int(m.ActionId) ,
    ChatId:  int(m.ChatId) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int(m.MessageId) ,
    ReSharedId:  int(m.ReSharedId) ,
    Murmur64Hash:  int(m.Murmur64Hash) ,
}
return r
}

func(m *PB_Followed)ToFlat() *PB_Followed_Flat {
r := &PB_Followed_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    FollowedUserId:  int(m.FollowedUserId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_Likes)ToFlat() *PB_Likes_Flat {
r := &PB_Likes_Flat{
    Id:  int(m.Id) ,
    PostId:  int(m.PostId) ,
    UserId:  int(m.UserId) ,
    PostType:  int(m.PostType) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_Notify)ToFlat() *PB_Notify_Flat {
r := &PB_Notify_Flat{
    NotifyId:  int(m.NotifyId) ,
    ForUserId:  int(m.ForUserId) ,
    ActorUserId:  int(m.ActorUserId) ,
    NotifyType:  int(m.NotifyType) ,
    PostId:  int(m.PostId) ,
    CommentId:  int(m.CommentId) ,
    PeerUserId:  int(m.PeerUserId) ,
    Murmur64Hash:  int(m.Murmur64Hash) ,
    SeenStatus:  int(m.SeenStatus) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_NotifyRemoved)ToFlat() *PB_NotifyRemoved_Flat {
r := &PB_NotifyRemoved_Flat{
    Murmur64Hash:  int(m.Murmur64Hash) ,
    ForUserId:  int(m.ForUserId) ,
    Id:  int(m.Id) ,
}
return r
}

func(m *PB_PhoneContacts)ToFlat() *PB_PhoneContacts_Flat {
r := &PB_PhoneContacts_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    ClientId:  int(m.ClientId) ,
    Phone:  m.Phone ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
}
return r
}

func(m *PB_Post)ToFlat() *PB_Post_Flat {
r := &PB_Post_Flat{
    PostId:  int(m.PostId) ,
    UserId:  int(m.UserId) ,
    PostType:  int(m.PostType) ,
    MediaId:  int(m.MediaId) ,
    FileRefId:  int(m.FileRefId) ,
    PostKey:  m.PostKey ,
    Text:  m.Text ,
    RichText:  m.RichText ,
    MediaCount:  int(m.MediaCount) ,
    SharedTo:  int(m.SharedTo) ,
    DisableComment:  int(m.DisableComment) ,
    Via:  int(m.Via) ,
    Seq:  int(m.Seq) ,
    CommentsCount:  int(m.CommentsCount) ,
    LikesCount:  int(m.LikesCount) ,
    ViewsCount:  int(m.ViewsCount) ,
    EditedTime:  int(m.EditedTime) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_PostCount)ToFlat() *PB_PostCount_Flat {
r := &PB_PostCount_Flat{
    PostId:  int(m.PostId) ,
    CommentsCount:  int(m.CommentsCount) ,
    LikesCount:  int(m.LikesCount) ,
    ViewsCount:  int(m.ViewsCount) ,
    ReSharedCount:  int(m.ReSharedCount) ,
    ChatSharedCount:  int(m.ChatSharedCount) ,
}
return r
}

func(m *PB_PostDeleted)ToFlat() *PB_PostDeleted_Flat {
r := &PB_PostDeleted_Flat{
    PostId:  int(m.PostId) ,
    UserId:  int(m.UserId) ,
}
return r
}

func(m *PB_PostKeys)ToFlat() *PB_PostKeys_Flat {
r := &PB_PostKeys_Flat{
    Id:  int(m.Id) ,
    PostKeyStr:  m.PostKeyStr ,
    Used:  int(m.Used) ,
    RandShard:  int(m.RandShard) ,
}
return r
}

func(m *PB_PostLink)ToFlat() *PB_PostLink_Flat {
r := &PB_PostLink_Flat{
    LinkId:  int(m.LinkId) ,
    LinkUrl:  m.LinkUrl ,
}
return r
}

func(m *PB_PostMedia)ToFlat() *PB_PostMedia_Flat {
r := &PB_PostMedia_Flat{
    MediaId:  int(m.MediaId) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    AlbumId:  int(m.AlbumId) ,
    MediaTypeEnum:  int(m.MediaTypeEnum) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Size:  int(m.Size) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    Md5Hash:  m.Md5Hash ,
    Color:  m.Color ,
    CreatedTime:  int(m.CreatedTime) ,
    ViewCount:  int(m.ViewCount) ,
    Extra:  m.Extra ,
}
return r
}

func(m *PB_PostReshared)ToFlat() *PB_PostReshared_Flat {
r := &PB_PostReshared_Flat{
    ResharedId:  int(m.ResharedId) ,
    PostId:  int(m.PostId) ,
    ByUserId:  int(m.ByUserId) ,
    PostUserId:  int(m.PostUserId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_ProfileAll)ToFlat() *PB_ProfileAll_Flat {
r := &PB_ProfileAll_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    IsReShared:  int(m.IsReShared) ,
}
return r
}

func(m *PB_ProfileMedia)ToFlat() *PB_ProfileMedia_Flat {
r := &PB_ProfileMedia_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    PostType:  int(m.PostType) ,
}
return r
}

func(m *PB_ProfileMentioned)ToFlat() *PB_ProfileMentioned_Flat {
r := &PB_ProfileMentioned_Flat{
    Id:  int(m.Id) ,
    ForUserId:  int(m.ForUserId) ,
    PostId:  int(m.PostId) ,
    PostUserId:  int(m.PostUserId) ,
    PostType:  int(m.PostType) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_Session)ToFlat() *PB_Session_Flat {
r := &PB_Session_Flat{
    Id:  int(m.Id) ,
    SessionUuid:  m.SessionUuid ,
    UserId:  int(m.UserId) ,
    LastIpAddress:  m.LastIpAddress ,
    AppVersion:  int(m.AppVersion) ,
    ActiveTime:  int(m.ActiveTime) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_SettingNotifications)ToFlat() *PB_SettingNotifications_Flat {
r := &PB_SettingNotifications_Flat{
    UserId:  int(m.UserId) ,
    SocialLedOn:  int(m.SocialLedOn) ,
    SocialLedColor:  m.SocialLedColor ,
    ReqestToFollowYou:  int(m.ReqestToFollowYou) ,
    FollowedYou:  int(m.FollowedYou) ,
    AccptedYourFollowRequest:  int(m.AccptedYourFollowRequest) ,
    YourPostLiked:  int(m.YourPostLiked) ,
    YourPostCommented:  int(m.YourPostCommented) ,
    MenthenedYouInPost:  int(m.MenthenedYouInPost) ,
    MenthenedYouInComment:  int(m.MenthenedYouInComment) ,
    YourContactsJoined:  int(m.YourContactsJoined) ,
    DirectMessage:  int(m.DirectMessage) ,
    DirectAlert:  int(m.DirectAlert) ,
    DirectPerview:  int(m.DirectPerview) ,
    DirectLedOn:  int(m.DirectLedOn) ,
    DirectLedColor:  int(m.DirectLedColor) ,
    DirectVibrate:  int(m.DirectVibrate) ,
    DirectPopup:  int(m.DirectPopup) ,
    DirectSound:  int(m.DirectSound) ,
    DirectPriority:  int(m.DirectPriority) ,
}
return r
}

func(m *PB_Sms)ToFlat() *PB_Sms_Flat {
r := &PB_Sms_Flat{
    Id:  int(m.Id) ,
    Hash:  m.Hash ,
    AppUuid:  m.AppUuid ,
    ClientPhone:  m.ClientPhone ,
    GenratedCode:  int(m.GenratedCode) ,
    SmsSenderNumber:  m.SmsSenderNumber ,
    SmsSendStatues:  m.SmsSendStatues ,
    SmsHttpBody:  m.SmsHttpBody ,
    Err:  m.Err ,
    Carrier:  m.Carrier ,
    Country:  []byte(m.Country) ,
    IsValidPhone:  int(m.IsValidPhone) ,
    IsConfirmed:  int(m.IsConfirmed) ,
    IsLogin:  int(m.IsLogin) ,
    IsRegister:  int(m.IsRegister) ,
    RetriedCount:  int(m.RetriedCount) ,
    TTL:  int(m.TTL) ,
}
return r
}

func(m *PB_Tag)ToFlat() *PB_Tag_Flat {
r := &PB_Tag_Flat{
    TagId:  int(m.TagId) ,
    Name:  m.Name ,
    Count:  int(m.Count) ,
    TagStatusEnum:  int(m.TagStatusEnum) ,
    IsBlocked:  int(m.IsBlocked) ,
    GroupId:  int(m.GroupId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_TagPost)ToFlat() *PB_TagPost_Flat {
r := &PB_TagPost_Flat{
    Id:  int(m.Id) ,
    TagId:  int(m.TagId) ,
    PostId:  int(m.PostId) ,
    PostType:  int(m.PostType) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_TriggerLog)ToFlat() *PB_TriggerLog_Flat {
r := &PB_TriggerLog_Flat{
    Id:  int(m.Id) ,
    ModelName:  m.ModelName ,
    ChangeType:  m.ChangeType ,
    TargetId:  int(m.TargetId) ,
    TargetStr:  m.TargetStr ,
    CreatedSe:  int(m.CreatedSe) ,
}
return r
}

func(m *PB_User)ToFlat() *PB_User_Flat {
r := &PB_User_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
    UserNameLower:  m.UserNameLower ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    IsVerified:  int(m.IsVerified) ,
    AvatarId:  int(m.AvatarId) ,
    AccessHash:  int(m.AccessHash) ,
    ProfilePrivacy:  int(m.ProfilePrivacy) ,
    OnlinePrivacy:  int(m.OnlinePrivacy) ,
    CallPrivacy:  int(m.CallPrivacy) ,
    AddToGroupPrivacy:  int(m.AddToGroupPrivacy) ,
    SeenMessagePrivacy:  int(m.SeenMessagePrivacy) ,
    Phone:  m.Phone ,
    Email:  m.Email ,
    About:  m.About ,
    PasswordHash:  m.PasswordHash ,
    PasswordSalt:  m.PasswordSalt ,
    PostSeq:  int(m.PostSeq) ,
    FollowersCount:  int(m.FollowersCount) ,
    FollowingCount:  int(m.FollowingCount) ,
    PostsCount:  int(m.PostsCount) ,
    MediaCount:  int(m.MediaCount) ,
    PhotoCount:  int(m.PhotoCount) ,
    VideoCount:  int(m.VideoCount) ,
    GifCount:  int(m.GifCount) ,
    AudioCount:  int(m.AudioCount) ,
    VoiceCount:  int(m.VoiceCount) ,
    FileCount:  int(m.FileCount) ,
    LinkCount:  int(m.LinkCount) ,
    BoardCount:  int(m.BoardCount) ,
    PinedCount:  int(m.PinedCount) ,
    LikesCount:  int(m.LikesCount) ,
    ResharedCount:  int(m.ResharedCount) ,
    LastPostTime:  int(m.LastPostTime) ,
    CreatedTime:  int(m.CreatedTime) ,
    VersionTime:  int(m.VersionTime) ,
    IsDeleted:  int(m.IsDeleted) ,
    IsBanned:  int(m.IsBanned) ,
}
return r
}

func(m *PB_UserRelation)ToFlat() *PB_UserRelation_Flat {
r := &PB_UserRelation_Flat{
    RelNanoId:  int(m.RelNanoId) ,
    UserId:  int(m.UserId) ,
    PeerUserId:  int(m.PeerUserId) ,
    Follwing:  int(m.Follwing) ,
    Followed:  int(m.Followed) ,
    InContacts:  int(m.InContacts) ,
    MutualContact:  int(m.MutualContact) ,
    IsFavorite:  int(m.IsFavorite) ,
    Notify:  int(m.Notify) ,
}
return r
}

func(m *PB_Chat)ToFlat() *PB_Chat_Flat {
r := &PB_Chat_Flat{
    ChatId:  int(m.ChatId) ,
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomType:  int(m.RoomType) ,
    UserId:  int(m.UserId) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    HashTagId:  int(m.HashTagId) ,
    Title:  m.Title ,
    PinTimeMs:  int(m.PinTimeMs) ,
    FromMsgId:  int(m.FromMsgId) ,
    UnseenCount:  int(m.UnseenCount) ,
    Seq:  int(m.Seq) ,
    LastMsgId:  int(m.LastMsgId) ,
    LastMyMsgStatus:  int(m.LastMyMsgStatus) ,
    MyLastSeenSeq:  int(m.MyLastSeenSeq) ,
    MyLastSeenMsgId:  int(m.MyLastSeenMsgId) ,
    PeerLastSeenMsgId:  int(m.PeerLastSeenMsgId) ,
    MyLastDeliveredSeq:  int(m.MyLastDeliveredSeq) ,
    MyLastDeliveredMsgId:  int(m.MyLastDeliveredMsgId) ,
    PeerLastDeliveredMsgId:  int(m.PeerLastDeliveredMsgId) ,
    IsActive:  int(m.IsActive) ,
    IsLeft:  int(m.IsLeft) ,
    IsCreator:  int(m.IsCreator) ,
    IsKicked:  int(m.IsKicked) ,
    IsAdmin:  int(m.IsAdmin) ,
    IsDeactivated:  int(m.IsDeactivated) ,
    IsMuted:  int(m.IsMuted) ,
    MuteUntil:  int(m.MuteUntil) ,
    VersionTimeMs:  int(m.VersionTimeMs) ,
    VersionSeq:  int(m.VersionSeq) ,
    OrderTime:  int(m.OrderTime) ,
    CreatedTime:  int(m.CreatedTime) ,
    DraftText:  m.DraftText ,
    DratReplyToMsgId:  int(m.DratReplyToMsgId) ,
}
return r
}

func(m *PB_ChatDeleted)ToFlat() *PB_ChatDeleted_Flat {
r := &PB_ChatDeleted_Flat{
    ChatId:  int(m.ChatId) ,
    RoomKey:  m.RoomKey ,
}
return r
}

func(m *PB_ChatLastMessage)ToFlat() *PB_ChatLastMessage_Flat {
r := &PB_ChatLastMessage_Flat{
    ChatIdGroupId:  m.ChatIdGroupId ,
    LastMsgPb:  []byte(m.LastMsgPb) ,
}
return r
}

func(m *PB_ChatUserVersion)ToFlat() *PB_ChatUserVersion_Flat {
r := &PB_ChatUserVersion_Flat{
    UserId:  int(m.UserId) ,
    ChatId:  int(m.ChatId) ,
    VersionTimeNano:  int(m.VersionTimeNano) ,
}
return r
}

func(m *PB_Group)ToFlat() *PB_Group_Flat {
r := &PB_Group_Flat{
    GroupId:  int(m.GroupId) ,
    GroupKey:  m.GroupKey ,
    GroupName:  m.GroupName ,
    UserName:  m.UserName ,
    IsSuperGroup:  int(m.IsSuperGroup) ,
    HashTagId:  int(m.HashTagId) ,
    CreatorUserId:  int(m.CreatorUserId) ,
    GroupPrivacy:  int(m.GroupPrivacy) ,
    HistoryViewAble:  int(m.HistoryViewAble) ,
    Seq:  int(m.Seq) ,
    LastMsgId:  int(m.LastMsgId) ,
    PinedMsgId:  int(m.PinedMsgId) ,
    AvatarRefId:  int(m.AvatarRefId) ,
    AvatarCount:  int(m.AvatarCount) ,
    About:  m.About ,
    InviteLinkHash:  m.InviteLinkHash ,
    MembersCount:  int(m.MembersCount) ,
    AdminsCount:  int(m.AdminsCount) ,
    ModeratorCounts:  int(m.ModeratorCounts) ,
    SortTime:  int(m.SortTime) ,
    CreatedTime:  int(m.CreatedTime) ,
    IsMute:  int(m.IsMute) ,
    IsDeleted:  int(m.IsDeleted) ,
    IsBanned:  int(m.IsBanned) ,
}
return r
}

func(m *PB_GroupMember)ToFlat() *PB_GroupMember_Flat {
r := &PB_GroupMember_Flat{
    OrderId:  int(m.OrderId) ,
    GroupId:  int(m.GroupId) ,
    UserId:  int(m.UserId) ,
    ByUserId:  int(m.ByUserId) ,
    GroupRole:  int(m.GroupRole) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_GroupOrderdUser)ToFlat() *PB_GroupOrderdUser_Flat {
r := &PB_GroupOrderdUser_Flat{
    OrderId:  int(m.OrderId) ,
    GroupId:  int(m.GroupId) ,
    UserId:  int(m.UserId) ,
}
return r
}

func(m *PB_GroupPinedMsg)ToFlat() *PB_GroupPinedMsg_Flat {
r := &PB_GroupPinedMsg_Flat{
    MsgId:  int(m.MsgId) ,
    MsgPb:  []byte(m.MsgPb) ,
}
return r
}

func(m *PB_FileMsg)ToFlat() *PB_FileMsg_Flat {
r := &PB_FileMsg_Flat{
    Id:  int(m.Id) ,
    AccessHash:  int(m.AccessHash) ,
    FileType:  int(m.FileType) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Extension:  m.Extension ,
    UserId:  int(m.UserId) ,
    DataThumb:  []byte(m.DataThumb) ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *PB_FilePost)ToFlat() *PB_FilePost_Flat {
r := &PB_FilePost_Flat{
    Id:  int(m.Id) ,
    AccessHash:  int(m.AccessHash) ,
    FileType:  int(m.FileType) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Extension:  m.Extension ,
    UserId:  int(m.UserId) ,
    DataThumb:  []byte(m.DataThumb) ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *PB_ActionFanout)ToFlat() *PB_ActionFanout_Flat {
r := &PB_ActionFanout_Flat{
    OrderId:  int(m.OrderId) ,
    ForUserId:  int(m.ForUserId) ,
    ActionId:  int(m.ActionId) ,
    ActorUserId:  int(m.ActorUserId) ,
}
return r
}

func(m *PB_HomeFanout)ToFlat() *PB_HomeFanout_Flat {
r := &PB_HomeFanout_Flat{
    OrderId:  int(m.OrderId) ,
    ForUserId:  int(m.ForUserId) ,
    PostId:  int(m.PostId) ,
    PostUserId:  int(m.PostUserId) ,
    ResharedId:  int(m.ResharedId) ,
}
return r
}

func(m *PB_SuggestedTopPosts)ToFlat() *PB_SuggestedTopPosts_Flat {
r := &PB_SuggestedTopPosts_Flat{
    Id:  int(m.Id) ,
    PostId:  int(m.PostId) ,
}
return r
}

func(m *PB_SuggestedUser)ToFlat() *PB_SuggestedUser_Flat {
r := &PB_SuggestedUser_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    TargetId:  int(m.TargetId) ,
    Weight:  float32(m.Weight) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_PushChat)ToFlat() *PB_PushChat_Flat {
r := &PB_PushChat_Flat{
    PushId:  int(m.PushId) ,
    ToUserId:  int(m.ToUserId) ,
    PushTypeId:  int(m.PushTypeId) ,
    RoomKey:  m.RoomKey ,
    ChatKey:  m.ChatKey ,
    Seq:  int(m.Seq) ,
    UnseenCount:  int(m.UnseenCount) ,
    FromHighMessageId:  int(m.FromHighMessageId) ,
    ToLowMessageId:  int(m.ToLowMessageId) ,
    MessageId:  int(m.MessageId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessagePb:  []byte(m.MessagePb) ,
    MessageJson:  m.MessageJson ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_HTTPRPCLog)ToFlat() *PB_HTTPRPCLog_Flat {
r := &PB_HTTPRPCLog_Flat{
    Id:  int(m.Id) ,
    Time:  m.Time ,
    MethodFull:  m.MethodFull ,
    MethodParent:  m.MethodParent ,
    UserId:  int(m.UserId) ,
    SessionId:  m.SessionId ,
    StatusCode:  int(m.StatusCode) ,
    InputSize:  int(m.InputSize) ,
    OutputSize:  int(m.OutputSize) ,
    ReqestJson:  m.ReqestJson ,
    ResponseJson:  m.ResponseJson ,
    ReqestParamJson:  m.ReqestParamJson ,
    ResponseMsgJson:  m.ResponseMsgJson ,
}
return r
}

func(m *PB_MetricLog)ToFlat() *PB_MetricLog_Flat {
r := &PB_MetricLog_Flat{
    Id:  int(m.Id) ,
    InstanceId:  int(m.InstanceId) ,
    StartFrom:  m.StartFrom ,
    EndTo:  m.EndTo ,
    StartTime:  int(m.StartTime) ,
    Duration:  m.Duration ,
    MetericsJson:  m.MetericsJson ,
}
return r
}

func(m *PB_XfileServiceInfoLog)ToFlat() *PB_XfileServiceInfoLog_Flat {
r := &PB_XfileServiceInfoLog_Flat{
    Id:  int(m.Id) ,
    InstanceId:  int(m.InstanceId) ,
    Url:  m.Url ,
    CreatedTime:  m.CreatedTime ,
}
return r
}

func(m *PB_XfileServiceMetricLog)ToFlat() *PB_XfileServiceMetricLog_Flat {
r := &PB_XfileServiceMetricLog_Flat{
    Id:  int(m.Id) ,
    InstanceId:  int(m.InstanceId) ,
    MetricJson:  m.MetricJson ,
}
return r
}

func(m *PB_XfileServiceRequestLog)ToFlat() *PB_XfileServiceRequestLog_Flat {
r := &PB_XfileServiceRequestLog_Flat{
    Id:  int(m.Id) ,
    LocalSeq:  int(m.LocalSeq) ,
    InstanceId:  int(m.InstanceId) ,
    Url:  m.Url ,
    HttpCode:  int(m.HttpCode) ,
    CreatedTime:  m.CreatedTime ,
}
return r
}

func(m *PB_InvalidateCache)ToFlat() *PB_InvalidateCache_Flat {
r := &PB_InvalidateCache_Flat{
    OrderId:  int(m.OrderId) ,
    CacheKey:  m.CacheKey ,
}
return r
}

func(m *PB_MediaView)ToFlat() *PB_MediaView_Flat {
r := &PB_MediaView_Flat{
}
return r
}

func(m *PB_ActionView)ToFlat() *PB_ActionView_Flat {
r := &PB_ActionView_Flat{
    ActionId:  int(m.ActionId) ,
    ActorUserId:  int(m.ActorUserId) ,
    ActionTypeEnum:  int(m.ActionTypeEnum) ,
    PeerUserId:  int(m.PeerUserId) ,
    PostId:  int(m.PostId) ,
    CommentId:  int(m.CommentId) ,
    Murmur64Hash:  int(m.Murmur64Hash) ,
    CreatedTime:  int(m.CreatedTime) ,





}
return r
}

func(m *PB_NotifyView)ToFlat() *PB_NotifyView_Flat {
r := &PB_NotifyView_Flat{
    NotifyId:  int(m.NotifyId) ,
    ForUserId:  int(m.ForUserId) ,
    ActorUserId:  int(m.ActorUserId) ,
    NotiyTypeEnum:  int(m.NotiyTypeEnum) ,
    PostId:  int(m.PostId) ,
    CommentId:  int(m.CommentId) ,
    PeerUserId:  int(m.PeerUserId) ,
    Murmur64Hash:  int(m.Murmur64Hash) ,
    SeenStatus:  int(m.SeenStatus) ,
    CreatedTime:  int(m.CreatedTime) ,



}
return r
}

func(m *PB_CommentView)ToFlat() *PB_CommentView_Flat {
r := &PB_CommentView_Flat{
    CommentId:  int(m.CommentId) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    Text:  m.Text ,
    LikesCount:  int(m.LikesCount) ,
    CreatedTime:  int(m.CreatedTime) ,

}
return r
}

func(m *PB_PostView)ToFlat() *PB_PostView_Flat {
r := &PB_PostView_Flat{
    PostId:  int(m.PostId) ,
    UserId:  int(m.UserId) ,

    Text:  m.Text ,
    RichText:  m.RichText ,
    MediaCount:  int(m.MediaCount) ,
    SharedTo:  int(m.SharedTo) ,
    DisableComment:  int(m.DisableComment) ,
    HasTag:  int(m.HasTag) ,
    CommentsCount:  int(m.CommentsCount) ,
    LikesCount:  int(m.LikesCount) ,
    ViewsCount:  int(m.ViewsCount) ,
    EditedTime:  int(m.EditedTime) ,
    CreatedTime:  int(m.CreatedTime) ,
    ReSharedPostId:  int(m.ReSharedPostId) ,
    DidILiked:  m.DidILiked ,
    DidIReShared:  m.DidIReShared ,




}
return r
}

func(m *PB_ChatView)ToFlat() *PB_ChatView_Flat {
r := &PB_ChatView_Flat{
    ChatId:  int(m.ChatId) ,
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomType:  int(m.RoomType) ,
    UserId:  int(m.UserId) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    HashTagId:  int(m.HashTagId) ,
    StartedByMe:  int(m.StartedByMe) ,
    Title:  m.Title ,
    PinTime:  int(m.PinTime) ,
    FromMsgId:  int(m.FromMsgId) ,
    Seq:  int(m.Seq) ,
    LastMsgId:  int(m.LastMsgId) ,
    LastMsgStatus:  int(m.LastMsgStatus) ,
    SeenSeq:  int(m.SeenSeq) ,
    SeenMsgId:  int(m.SeenMsgId) ,
    Left:  int(m.Left) ,
    Creator:  int(m.Creator) ,
    Kicked:  int(m.Kicked) ,
    Admin:  int(m.Admin) ,
    Deactivated:  int(m.Deactivated) ,
    VersionTime:  int(m.VersionTime) ,
    SortTime:  int(m.SortTime) ,
    CreatedTime:  int(m.CreatedTime) ,
    DraftText:  m.DraftText ,
    DratReplyToMsgId:  int(m.DratReplyToMsgId) ,
    IsMute:  int(m.IsMute) ,




}
return r
}

func(m *PB_GroupView)ToFlat() *PB_GroupView_Flat {
r := &PB_GroupView_Flat{
    GroupId:  int(m.GroupId) ,
    GroupKey:  m.GroupKey ,
    GroupName:  m.GroupName ,
    UserName:  m.UserName ,
    IsSuperGroup:  int(m.IsSuperGroup) ,
    HashTagId:  int(m.HashTagId) ,
    CreatorUserId:  int(m.CreatorUserId) ,
    GroupPrivacy:  int(m.GroupPrivacy) ,
    HistoryViewAble:  int(m.HistoryViewAble) ,
    Seq:  int(m.Seq) ,
    LastMsgId:  int(m.LastMsgId) ,
    PinedMsgId:  int(m.PinedMsgId) ,
    AvatarRefId:  int(m.AvatarRefId) ,
    AvatarCount:  int(m.AvatarCount) ,
    About:  m.About ,
    InviteLink:  m.InviteLink ,
    MembersCount:  int(m.MembersCount) ,
    SortTime:  int(m.SortTime) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_GroupMemberView)ToFlat() *PB_GroupMemberView_Flat {
r := &PB_GroupMemberView_Flat{
    OrderId:  int(m.OrderId) ,
    GroupId:  int(m.GroupId) ,
    UserId:  int(m.UserId) ,
    ByUserId:  int(m.ByUserId) ,
    GroupRole:  int(m.GroupRole) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_MessageView)ToFlat() *PB_MessageView_Flat {
r := &PB_MessageView_Flat{
    RoomKey:  m.RoomKey ,
    MessageId:  int(m.MessageId) ,
    UserId:  int(m.UserId) ,
    FileRefId:  int(m.FileRefId) ,
    MessageType:  int(m.MessageType) ,
    Text:  m.Text ,
    Hiden:  int(m.Hiden) ,
    Seq:  int(m.Seq) ,
    ForwardedMsgId:  int(m.ForwardedMsgId) ,
    PostId:  int(m.PostId) ,
    StickerId:  int(m.StickerId) ,
    CreatedTime:  int(m.CreatedTime) ,
    DeliveredTime:  int(m.DeliveredTime) ,
    SeenTime:  int(m.SeenTime) ,
    DeliviryStatus:  int(m.DeliviryStatus) ,
    ReplyToMessageId:  int(m.ReplyToMessageId) ,
    ViewsCount:  int(m.ViewsCount) ,
    EditTime:  int(m.EditTime) ,
    Ttl:  int(m.Ttl) ,

}
return r
}

func(m *PB_FileRedView)ToFlat() *PB_FileRedView_Flat {
r := &PB_FileRedView_Flat{
    FileRefId:  int(m.FileRefId) ,
    UserId:  int(m.UserId) ,
    Name:  m.Name ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    UrlSource:  m.UrlSource ,
}
return r
}

func(m *PB_UserView)ToFlat() *PB_UserView_Flat {
r := &PB_UserView_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    AvatarRefId:  int(m.AvatarRefId) ,
    ProfilePrivacy:  int(m.ProfilePrivacy) ,
    Phone:  int(m.Phone) ,
    About:  m.About ,
    FollowersCount:  int(m.FollowersCount) ,
    FollowingCount:  int(m.FollowingCount) ,
    PostsCount:  int(m.PostsCount) ,
    MediaCount:  int(m.MediaCount) ,

    LastActiveTime:  int(m.LastActiveTime) ,
    LastActiveTimeShow:  m.LastActiveTimeShow ,

}
return r
}

func(m *PB_SettingNotificationView)ToFlat() *PB_SettingNotificationView_Flat {
r := &PB_SettingNotificationView_Flat{
}
return r
}

func(m *PB_AppConfig)ToFlat() *PB_AppConfig_Flat {
r := &PB_AppConfig_Flat{
    DeprecatedClient:  m.DeprecatedClient ,
    HasNewUpdate:  m.HasNewUpdate ,
}
return r
}

func(m *PB_UserProfileView)ToFlat() *PB_UserProfileView_Flat {
r := &PB_UserProfileView_Flat{
}
return r
}

func(m *PB_UserViewRowify)ToFlat() *PB_UserViewRowify_Flat {
r := &PB_UserViewRowify_Flat{
    Id:  int(m.Id) ,
    CreatedTime:  int(m.CreatedTime) ,

}
return r
}

func(m *PB_PostViewRowify)ToFlat() *PB_PostViewRowify_Flat {
r := &PB_PostViewRowify_Flat{
    Id:  int(m.Id) ,

}
return r
}

func(m *PB_TagView)ToFlat() *PB_TagView_Flat {
r := &PB_TagView_Flat{
    TagId:  int(m.TagId) ,
    Name:  m.Name ,
    Count:  int(m.Count) ,
    TagStatusEnum:  int(m.TagStatusEnum) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_TopTagWithSamplePosts)ToFlat() *PB_TopTagWithSamplePosts_Flat {
r := &PB_TopTagWithSamplePosts_Flat{


}
return r
}

func(m *PB_SelfUserView)ToFlat() *PB_SelfUserView_Flat {
r := &PB_SelfUserView_Flat{

    ProfilePrivacy:  int(m.ProfilePrivacy) ,
    OnlinePrivacy:  int(m.OnlinePrivacy) ,
    CallPrivacy:  int(m.CallPrivacy) ,
    AddToGroupPrivacy:  int(m.AddToGroupPrivacy) ,
    SeenMessagePrivacy:  int(m.SeenMessagePrivacy) ,

}
return r
}

func(m *PB_Error)ToFlat() *PB_Error_Flat {
r := &PB_Error_Flat{

    ShowError:  m.ShowError ,
    ErrorMessage:  m.ErrorMessage ,
}
return r
}



///// from_flat ///

func(m *PB_RoomsChanges_Flat)ToPB() *PB_RoomsChanges {
r := &PB_RoomsChanges{
    VersionTime:  uint64(m.VersionTime) ,


}
return r
}

func(m *PB_PushChanges_Flat)ToPB() *PB_PushChanges {
r := &PB_PushChanges{


    InvalidateCacheForRoomKeys:  m.InvalidateCacheForRoomKeys ,
    InvalidateAllChatCache:  int32(m.InvalidateAllChatCache) ,
    InvalidateAllSocialCache:  int32(m.InvalidateAllSocialCache) ,
}
return r
}

func(m *PB_CommandToServer_Flat)ToPB() *PB_CommandToServer {
r := &PB_CommandToServer{
    ClientCallId:  int64(m.ClientCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  m.Data ,
}
return r
}

func(m *PB_CommandToClient_Flat)ToPB() *PB_CommandToClient {
r := &PB_CommandToClient{
    ServerCallId:  int64(m.ServerCallId) ,
    Command:  m.Command ,
    RespondReached:  m.RespondReached ,
    Data:  m.Data ,
}
return r
}

func(m *PB_CommandReachedToServer_Flat)ToPB() *PB_CommandReachedToServer {
r := &PB_CommandReachedToServer{
    ClientCallId:  int64(m.ClientCallId) ,
}
return r
}

func(m *PB_CommandReachedToClient_Flat)ToPB() *PB_CommandReachedToClient {
r := &PB_CommandReachedToClient{
    ServerCallId:  int64(m.ServerCallId) ,
}
return r
}

func(m *PB_ResponseToClient_Flat)ToPB() *PB_ResponseToClient {
r := &PB_ResponseToClient{
    ClientCallId:  int64(m.ClientCallId) ,
    PBClass:  m.PBClass ,
    RpcFullName:  m.RpcFullName ,
    Data:  m.Data ,
}
return r
}

func(m *RPC_Auth_Types_Flat)ToPB() *RPC_Auth_Types {
r := &RPC_Auth_Types{
}
return r
}

func(m *PB_RPC_Chat_Types_Flat)ToPB() *PB_RPC_Chat_Types {
r := &PB_RPC_Chat_Types{
}
return r
}

func(m *RPC_General_Types_Flat)ToPB() *RPC_General_Types {
r := &RPC_General_Types{
}
return r
}

func(m *RPC_Page_Types_Flat)ToPB() *RPC_Page_Types {
r := &RPC_Page_Types{
}
return r
}

func(m *RPC_Social_Types_Flat)ToPB() *RPC_Social_Types {
r := &RPC_Social_Types{
}
return r
}

func(m *RPC_User_Types_Flat)ToPB() *RPC_User_Types {
r := &RPC_User_Types{
}
return r
}

func(m *PB_Action_Flat)ToPB() *PB_Action {
r := &PB_Action{
    ActionId:  int64(m.ActionId) ,
    ActorUserId:  int32(m.ActorUserId) ,
    ActionType:  int32(m.ActionType) ,
    PeerUserId:  int32(m.PeerUserId) ,
    PostId:  int64(m.PostId) ,
    CommentId:  int64(m.CommentId) ,
    Murmur64Hash:  int64(m.Murmur64Hash) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_Blocked_Flat)ToPB() *PB_Blocked {
r := &PB_Blocked{
    Id:  int64(m.Id) ,
    UserId:  int32(m.UserId) ,
    BlockedUserId:  int32(m.BlockedUserId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_Comment_Flat)ToPB() *PB_Comment {
r := &PB_Comment{
    CommentId:  int64(m.CommentId) ,
    UserId:  int32(m.UserId) ,
    PostId:  int64(m.PostId) ,
    Text:  m.Text ,
    LikesCount:  int32(m.LikesCount) ,
    IsEdited:  int32(m.IsEdited) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_CommentDeleted_Flat)ToPB() *PB_CommentDeleted {
r := &PB_CommentDeleted{
    CommentId:  int64(m.CommentId) ,
    UserId:  int32(m.UserId) ,
}
return r
}

func(m *PB_Event_Flat)ToPB() *PB_Event {
r := &PB_Event{
    EventId:  int64(m.EventId) ,
    EventType:  int32(m.EventType) ,
    ByUserId:  int64(m.ByUserId) ,
    PeerUserId:  int64(m.PeerUserId) ,
    PostId:  int64(m.PostId) ,
    CommentId:  int64(m.CommentId) ,
    HashTagId:  int64(m.HashTagId) ,
    GroupId:  int64(m.GroupId) ,
    ActionId:  int64(m.ActionId) ,
    ChatId:  int64(m.ChatId) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int64(m.MessageId) ,
    ReSharedId:  int64(m.ReSharedId) ,
    Murmur64Hash:  int64(m.Murmur64Hash) ,
}
return r
}

func(m *PB_Followed_Flat)ToPB() *PB_Followed {
r := &PB_Followed{
    Id:  int64(m.Id) ,
    UserId:  int32(m.UserId) ,
    FollowedUserId:  int32(m.FollowedUserId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_Likes_Flat)ToPB() *PB_Likes {
r := &PB_Likes{
    Id:  int64(m.Id) ,
    PostId:  int64(m.PostId) ,
    UserId:  int32(m.UserId) ,
    PostType:  int32(m.PostType) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_Notify_Flat)ToPB() *PB_Notify {
r := &PB_Notify{
    NotifyId:  int64(m.NotifyId) ,
    ForUserId:  int32(m.ForUserId) ,
    ActorUserId:  int32(m.ActorUserId) ,
    NotifyType:  int32(m.NotifyType) ,
    PostId:  int64(m.PostId) ,
    CommentId:  int64(m.CommentId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    Murmur64Hash:  int64(m.Murmur64Hash) ,
    SeenStatus:  int32(m.SeenStatus) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_NotifyRemoved_Flat)ToPB() *PB_NotifyRemoved {
r := &PB_NotifyRemoved{
    Murmur64Hash:  int64(m.Murmur64Hash) ,
    ForUserId:  int32(m.ForUserId) ,
    Id:  int64(m.Id) ,
}
return r
}

func(m *PB_PhoneContacts_Flat)ToPB() *PB_PhoneContacts {
r := &PB_PhoneContacts{
    Id:  int64(m.Id) ,
    UserId:  int32(m.UserId) ,
    ClientId:  int64(m.ClientId) ,
    Phone:  m.Phone ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
}
return r
}

func(m *PB_Post_Flat)ToPB() *PB_Post {
r := &PB_Post{
    PostId:  int64(m.PostId) ,
    UserId:  int32(m.UserId) ,
    PostType:  int32(m.PostType) ,
    MediaId:  int64(m.MediaId) ,
    FileRefId:  int64(m.FileRefId) ,
    PostKey:  m.PostKey ,
    Text:  m.Text ,
    RichText:  m.RichText ,
    MediaCount:  int32(m.MediaCount) ,
    SharedTo:  int32(m.SharedTo) ,
    DisableComment:  int32(m.DisableComment) ,
    Via:  int32(m.Via) ,
    Seq:  int32(m.Seq) ,
    CommentsCount:  int32(m.CommentsCount) ,
    LikesCount:  int32(m.LikesCount) ,
    ViewsCount:  int32(m.ViewsCount) ,
    EditedTime:  int32(m.EditedTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_PostCount_Flat)ToPB() *PB_PostCount {
r := &PB_PostCount{
    PostId:  int64(m.PostId) ,
    CommentsCount:  int32(m.CommentsCount) ,
    LikesCount:  int32(m.LikesCount) ,
    ViewsCount:  int64(m.ViewsCount) ,
    ReSharedCount:  int32(m.ReSharedCount) ,
    ChatSharedCount:  int32(m.ChatSharedCount) ,
}
return r
}

func(m *PB_PostDeleted_Flat)ToPB() *PB_PostDeleted {
r := &PB_PostDeleted{
    PostId:  int64(m.PostId) ,
    UserId:  int32(m.UserId) ,
}
return r
}

func(m *PB_PostKeys_Flat)ToPB() *PB_PostKeys {
r := &PB_PostKeys{
    Id:  int32(m.Id) ,
    PostKeyStr:  m.PostKeyStr ,
    Used:  int32(m.Used) ,
    RandShard:  int32(m.RandShard) ,
}
return r
}

func(m *PB_PostLink_Flat)ToPB() *PB_PostLink {
r := &PB_PostLink{
    LinkId:  int64(m.LinkId) ,
    LinkUrl:  m.LinkUrl ,
}
return r
}

func(m *PB_PostMedia_Flat)ToPB() *PB_PostMedia {
r := &PB_PostMedia{
    MediaId:  int64(m.MediaId) ,
    UserId:  int32(m.UserId) ,
    PostId:  int64(m.PostId) ,
    AlbumId:  int32(m.AlbumId) ,
    MediaTypeEnum:  int32(m.MediaTypeEnum) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Size:  int32(m.Size) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    Md5Hash:  m.Md5Hash ,
    Color:  m.Color ,
    CreatedTime:  int32(m.CreatedTime) ,
    ViewCount:  int32(m.ViewCount) ,
    Extra:  m.Extra ,
}
return r
}

func(m *PB_PostReshared_Flat)ToPB() *PB_PostReshared {
r := &PB_PostReshared{
    ResharedId:  int64(m.ResharedId) ,
    PostId:  int64(m.PostId) ,
    ByUserId:  int32(m.ByUserId) ,
    PostUserId:  int32(m.PostUserId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_ProfileAll_Flat)ToPB() *PB_ProfileAll {
r := &PB_ProfileAll{
    Id:  int64(m.Id) ,
    UserId:  int32(m.UserId) ,
    PostId:  int64(m.PostId) ,
    IsReShared:  int32(m.IsReShared) ,
}
return r
}

func(m *PB_ProfileMedia_Flat)ToPB() *PB_ProfileMedia {
r := &PB_ProfileMedia{
    Id:  int64(m.Id) ,
    UserId:  int32(m.UserId) ,
    PostId:  int64(m.PostId) ,
    PostType:  int32(m.PostType) ,
}
return r
}

func(m *PB_ProfileMentioned_Flat)ToPB() *PB_ProfileMentioned {
r := &PB_ProfileMentioned{
    Id:  int64(m.Id) ,
    ForUserId:  int32(m.ForUserId) ,
    PostId:  int64(m.PostId) ,
    PostUserId:  int32(m.PostUserId) ,
    PostType:  int32(m.PostType) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_Session_Flat)ToPB() *PB_Session {
r := &PB_Session{
    Id:  int64(m.Id) ,
    SessionUuid:  m.SessionUuid ,
    UserId:  int32(m.UserId) ,
    LastIpAddress:  m.LastIpAddress ,
    AppVersion:  int32(m.AppVersion) ,
    ActiveTime:  int32(m.ActiveTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_SettingNotifications_Flat)ToPB() *PB_SettingNotifications {
r := &PB_SettingNotifications{
    UserId:  int32(m.UserId) ,
    SocialLedOn:  int32(m.SocialLedOn) ,
    SocialLedColor:  m.SocialLedColor ,
    ReqestToFollowYou:  int32(m.ReqestToFollowYou) ,
    FollowedYou:  int32(m.FollowedYou) ,
    AccptedYourFollowRequest:  int32(m.AccptedYourFollowRequest) ,
    YourPostLiked:  int32(m.YourPostLiked) ,
    YourPostCommented:  int32(m.YourPostCommented) ,
    MenthenedYouInPost:  int32(m.MenthenedYouInPost) ,
    MenthenedYouInComment:  int32(m.MenthenedYouInComment) ,
    YourContactsJoined:  int32(m.YourContactsJoined) ,
    DirectMessage:  int32(m.DirectMessage) ,
    DirectAlert:  int32(m.DirectAlert) ,
    DirectPerview:  int32(m.DirectPerview) ,
    DirectLedOn:  int32(m.DirectLedOn) ,
    DirectLedColor:  int32(m.DirectLedColor) ,
    DirectVibrate:  int32(m.DirectVibrate) ,
    DirectPopup:  int32(m.DirectPopup) ,
    DirectSound:  int32(m.DirectSound) ,
    DirectPriority:  int32(m.DirectPriority) ,
}
return r
}

func(m *PB_Sms_Flat)ToPB() *PB_Sms {
r := &PB_Sms{
    Id:  int32(m.Id) ,
    Hash:  m.Hash ,
    AppUuid:  m.AppUuid ,
    ClientPhone:  m.ClientPhone ,
    GenratedCode:  int32(m.GenratedCode) ,
    SmsSenderNumber:  m.SmsSenderNumber ,
    SmsSendStatues:  m.SmsSendStatues ,
    SmsHttpBody:  m.SmsHttpBody ,
    Err:  m.Err ,
    Carrier:  m.Carrier ,
    Country:  m.Country ,
    IsValidPhone:  int32(m.IsValidPhone) ,
    IsConfirmed:  int32(m.IsConfirmed) ,
    IsLogin:  int32(m.IsLogin) ,
    IsRegister:  int32(m.IsRegister) ,
    RetriedCount:  int32(m.RetriedCount) ,
    TTL:  int32(m.TTL) ,
}
return r
}

func(m *PB_Tag_Flat)ToPB() *PB_Tag {
r := &PB_Tag{
    TagId:  int64(m.TagId) ,
    Name:  m.Name ,
    Count:  int32(m.Count) ,
    TagStatusEnum:  int32(m.TagStatusEnum) ,
    IsBlocked:  int32(m.IsBlocked) ,
    GroupId:  int32(m.GroupId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_TagPost_Flat)ToPB() *PB_TagPost {
r := &PB_TagPost{
    Id:  int64(m.Id) ,
    TagId:  int32(m.TagId) ,
    PostId:  int32(m.PostId) ,
    PostType:  int32(m.PostType) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_TriggerLog_Flat)ToPB() *PB_TriggerLog {
r := &PB_TriggerLog{
    Id:  int64(m.Id) ,
    ModelName:  m.ModelName ,
    ChangeType:  m.ChangeType ,
    TargetId:  int64(m.TargetId) ,
    TargetStr:  m.TargetStr ,
    CreatedSe:  int32(m.CreatedSe) ,
}
return r
}

func(m *PB_User_Flat)ToPB() *PB_User {
r := &PB_User{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
    UserNameLower:  m.UserNameLower ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    IsVerified:  int32(m.IsVerified) ,
    AvatarId:  int64(m.AvatarId) ,
    AccessHash:  int32(m.AccessHash) ,
    ProfilePrivacy:  int32(m.ProfilePrivacy) ,
    OnlinePrivacy:  int32(m.OnlinePrivacy) ,
    CallPrivacy:  int32(m.CallPrivacy) ,
    AddToGroupPrivacy:  int32(m.AddToGroupPrivacy) ,
    SeenMessagePrivacy:  int32(m.SeenMessagePrivacy) ,
    Phone:  m.Phone ,
    Email:  m.Email ,
    About:  m.About ,
    PasswordHash:  m.PasswordHash ,
    PasswordSalt:  m.PasswordSalt ,
    PostSeq:  int32(m.PostSeq) ,
    FollowersCount:  int32(m.FollowersCount) ,
    FollowingCount:  int32(m.FollowingCount) ,
    PostsCount:  int32(m.PostsCount) ,
    MediaCount:  int32(m.MediaCount) ,
    PhotoCount:  int32(m.PhotoCount) ,
    VideoCount:  int32(m.VideoCount) ,
    GifCount:  int32(m.GifCount) ,
    AudioCount:  int32(m.AudioCount) ,
    VoiceCount:  int32(m.VoiceCount) ,
    FileCount:  int32(m.FileCount) ,
    LinkCount:  int32(m.LinkCount) ,
    BoardCount:  int32(m.BoardCount) ,
    PinedCount:  int32(m.PinedCount) ,
    LikesCount:  int32(m.LikesCount) ,
    ResharedCount:  int32(m.ResharedCount) ,
    LastPostTime:  int32(m.LastPostTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
    VersionTime:  int32(m.VersionTime) ,
    IsDeleted:  int32(m.IsDeleted) ,
    IsBanned:  int32(m.IsBanned) ,
}
return r
}

func(m *PB_UserRelation_Flat)ToPB() *PB_UserRelation {
r := &PB_UserRelation{
    RelNanoId:  int64(m.RelNanoId) ,
    UserId:  int32(m.UserId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    Follwing:  int32(m.Follwing) ,
    Followed:  int32(m.Followed) ,
    InContacts:  int32(m.InContacts) ,
    MutualContact:  int32(m.MutualContact) ,
    IsFavorite:  int32(m.IsFavorite) ,
    Notify:  int32(m.Notify) ,
}
return r
}

func(m *PB_Chat_Flat)ToPB() *PB_Chat {
r := &PB_Chat{
    ChatId:  int64(m.ChatId) ,
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomType:  int32(m.RoomType) ,
    UserId:  int32(m.UserId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    HashTagId:  int64(m.HashTagId) ,
    Title:  m.Title ,
    PinTimeMs:  int64(m.PinTimeMs) ,
    FromMsgId:  int64(m.FromMsgId) ,
    UnseenCount:  int32(m.UnseenCount) ,
    Seq:  int32(m.Seq) ,
    LastMsgId:  int64(m.LastMsgId) ,
    LastMyMsgStatus:  int32(m.LastMyMsgStatus) ,
    MyLastSeenSeq:  int32(m.MyLastSeenSeq) ,
    MyLastSeenMsgId:  int64(m.MyLastSeenMsgId) ,
    PeerLastSeenMsgId:  int64(m.PeerLastSeenMsgId) ,
    MyLastDeliveredSeq:  int32(m.MyLastDeliveredSeq) ,
    MyLastDeliveredMsgId:  int64(m.MyLastDeliveredMsgId) ,
    PeerLastDeliveredMsgId:  int64(m.PeerLastDeliveredMsgId) ,
    IsActive:  int32(m.IsActive) ,
    IsLeft:  int32(m.IsLeft) ,
    IsCreator:  int32(m.IsCreator) ,
    IsKicked:  int32(m.IsKicked) ,
    IsAdmin:  int32(m.IsAdmin) ,
    IsDeactivated:  int32(m.IsDeactivated) ,
    IsMuted:  int32(m.IsMuted) ,
    MuteUntil:  int32(m.MuteUntil) ,
    VersionTimeMs:  int64(m.VersionTimeMs) ,
    VersionSeq:  int32(m.VersionSeq) ,
    OrderTime:  int32(m.OrderTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
    DraftText:  m.DraftText ,
    DratReplyToMsgId:  int64(m.DratReplyToMsgId) ,
}
return r
}

func(m *PB_ChatDeleted_Flat)ToPB() *PB_ChatDeleted {
r := &PB_ChatDeleted{
    ChatId:  int64(m.ChatId) ,
    RoomKey:  m.RoomKey ,
}
return r
}

func(m *PB_ChatLastMessage_Flat)ToPB() *PB_ChatLastMessage {
r := &PB_ChatLastMessage{
    ChatIdGroupId:  m.ChatIdGroupId ,
    LastMsgPb:  m.LastMsgPb ,
}
return r
}

func(m *PB_ChatUserVersion_Flat)ToPB() *PB_ChatUserVersion {
r := &PB_ChatUserVersion{
    UserId:  int32(m.UserId) ,
    ChatId:  int64(m.ChatId) ,
    VersionTimeNano:  int32(m.VersionTimeNano) ,
}
return r
}

func(m *PB_Group_Flat)ToPB() *PB_Group {
r := &PB_Group{
    GroupId:  int64(m.GroupId) ,
    GroupKey:  m.GroupKey ,
    GroupName:  m.GroupName ,
    UserName:  m.UserName ,
    IsSuperGroup:  int32(m.IsSuperGroup) ,
    HashTagId:  int64(m.HashTagId) ,
    CreatorUserId:  int32(m.CreatorUserId) ,
    GroupPrivacy:  int32(m.GroupPrivacy) ,
    HistoryViewAble:  int32(m.HistoryViewAble) ,
    Seq:  int64(m.Seq) ,
    LastMsgId:  int64(m.LastMsgId) ,
    PinedMsgId:  int64(m.PinedMsgId) ,
    AvatarRefId:  int64(m.AvatarRefId) ,
    AvatarCount:  int32(m.AvatarCount) ,
    About:  m.About ,
    InviteLinkHash:  m.InviteLinkHash ,
    MembersCount:  int32(m.MembersCount) ,
    AdminsCount:  int32(m.AdminsCount) ,
    ModeratorCounts:  int32(m.ModeratorCounts) ,
    SortTime:  int32(m.SortTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
    IsMute:  int32(m.IsMute) ,
    IsDeleted:  int32(m.IsDeleted) ,
    IsBanned:  int32(m.IsBanned) ,
}
return r
}

func(m *PB_GroupMember_Flat)ToPB() *PB_GroupMember {
r := &PB_GroupMember{
    OrderId:  int64(m.OrderId) ,
    GroupId:  int64(m.GroupId) ,
    UserId:  int32(m.UserId) ,
    ByUserId:  int32(m.ByUserId) ,
    GroupRole:  int32(m.GroupRole) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_GroupOrderdUser_Flat)ToPB() *PB_GroupOrderdUser {
r := &PB_GroupOrderdUser{
    OrderId:  int64(m.OrderId) ,
    GroupId:  int64(m.GroupId) ,
    UserId:  int32(m.UserId) ,
}
return r
}

func(m *PB_GroupPinedMsg_Flat)ToPB() *PB_GroupPinedMsg {
r := &PB_GroupPinedMsg{
    MsgId:  int64(m.MsgId) ,
    MsgPb:  m.MsgPb ,
}
return r
}

func(m *PB_FileMsg_Flat)ToPB() *PB_FileMsg {
r := &PB_FileMsg{
    Id:  int64(m.Id) ,
    AccessHash:  int32(m.AccessHash) ,
    FileType:  int32(m.FileType) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Extension:  m.Extension ,
    UserId:  int32(m.UserId) ,
    DataThumb:  m.DataThumb ,
    Data:  m.Data ,
}
return r
}

func(m *PB_FilePost_Flat)ToPB() *PB_FilePost {
r := &PB_FilePost{
    Id:  int64(m.Id) ,
    AccessHash:  int32(m.AccessHash) ,
    FileType:  int32(m.FileType) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Extension:  m.Extension ,
    UserId:  int32(m.UserId) ,
    DataThumb:  m.DataThumb ,
    Data:  m.Data ,
}
return r
}

func(m *PB_ActionFanout_Flat)ToPB() *PB_ActionFanout {
r := &PB_ActionFanout{
    OrderId:  int64(m.OrderId) ,
    ForUserId:  int32(m.ForUserId) ,
    ActionId:  int64(m.ActionId) ,
    ActorUserId:  int32(m.ActorUserId) ,
}
return r
}

func(m *PB_HomeFanout_Flat)ToPB() *PB_HomeFanout {
r := &PB_HomeFanout{
    OrderId:  int64(m.OrderId) ,
    ForUserId:  int64(m.ForUserId) ,
    PostId:  int64(m.PostId) ,
    PostUserId:  int64(m.PostUserId) ,
    ResharedId:  int64(m.ResharedId) ,
}
return r
}

func(m *PB_SuggestedTopPosts_Flat)ToPB() *PB_SuggestedTopPosts {
r := &PB_SuggestedTopPosts{
    Id:  int64(m.Id) ,
    PostId:  int64(m.PostId) ,
}
return r
}

func(m *PB_SuggestedUser_Flat)ToPB() *PB_SuggestedUser {
r := &PB_SuggestedUser{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    TargetId:  int32(m.TargetId) ,
    Weight:  m.Weight ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_PushChat_Flat)ToPB() *PB_PushChat {
r := &PB_PushChat{
    PushId:  int64(m.PushId) ,
    ToUserId:  int32(m.ToUserId) ,
    PushTypeId:  int32(m.PushTypeId) ,
    RoomKey:  m.RoomKey ,
    ChatKey:  m.ChatKey ,
    Seq:  int32(m.Seq) ,
    UnseenCount:  int32(m.UnseenCount) ,
    FromHighMessageId:  int64(m.FromHighMessageId) ,
    ToLowMessageId:  int64(m.ToLowMessageId) ,
    MessageId:  int64(m.MessageId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessagePb:  m.MessagePb ,
    MessageJson:  m.MessageJson ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_HTTPRPCLog_Flat)ToPB() *PB_HTTPRPCLog {
r := &PB_HTTPRPCLog{
    Id:  int32(m.Id) ,
    Time:  m.Time ,
    MethodFull:  m.MethodFull ,
    MethodParent:  m.MethodParent ,
    UserId:  int32(m.UserId) ,
    SessionId:  m.SessionId ,
    StatusCode:  int32(m.StatusCode) ,
    InputSize:  int32(m.InputSize) ,
    OutputSize:  int32(m.OutputSize) ,
    ReqestJson:  m.ReqestJson ,
    ResponseJson:  m.ResponseJson ,
    ReqestParamJson:  m.ReqestParamJson ,
    ResponseMsgJson:  m.ResponseMsgJson ,
}
return r
}

func(m *PB_MetricLog_Flat)ToPB() *PB_MetricLog {
r := &PB_MetricLog{
    Id:  int32(m.Id) ,
    InstanceId:  int32(m.InstanceId) ,
    StartFrom:  m.StartFrom ,
    EndTo:  m.EndTo ,
    StartTime:  int32(m.StartTime) ,
    Duration:  m.Duration ,
    MetericsJson:  m.MetericsJson ,
}
return r
}

func(m *PB_XfileServiceInfoLog_Flat)ToPB() *PB_XfileServiceInfoLog {
r := &PB_XfileServiceInfoLog{
    Id:  int64(m.Id) ,
    InstanceId:  int32(m.InstanceId) ,
    Url:  m.Url ,
    CreatedTime:  m.CreatedTime ,
}
return r
}

func(m *PB_XfileServiceMetricLog_Flat)ToPB() *PB_XfileServiceMetricLog {
r := &PB_XfileServiceMetricLog{
    Id:  int64(m.Id) ,
    InstanceId:  int32(m.InstanceId) ,
    MetricJson:  m.MetricJson ,
}
return r
}

func(m *PB_XfileServiceRequestLog_Flat)ToPB() *PB_XfileServiceRequestLog {
r := &PB_XfileServiceRequestLog{
    Id:  int64(m.Id) ,
    LocalSeq:  int32(m.LocalSeq) ,
    InstanceId:  int32(m.InstanceId) ,
    Url:  m.Url ,
    HttpCode:  int32(m.HttpCode) ,
    CreatedTime:  m.CreatedTime ,
}
return r
}

func(m *PB_InvalidateCache_Flat)ToPB() *PB_InvalidateCache {
r := &PB_InvalidateCache{
    OrderId:  int64(m.OrderId) ,
    CacheKey:  m.CacheKey ,
}
return r
}

func(m *PB_MediaView_Flat)ToPB() *PB_MediaView {
r := &PB_MediaView{
}
return r
}

func(m *PB_ActionView_Flat)ToPB() *PB_ActionView {
r := &PB_ActionView{
    ActionId:  int64(m.ActionId) ,
    ActorUserId:  int32(m.ActorUserId) ,
    ActionTypeEnum:  int32(m.ActionTypeEnum) ,
    PeerUserId:  int32(m.PeerUserId) ,
    PostId:  int64(m.PostId) ,
    CommentId:  int64(m.CommentId) ,
    Murmur64Hash:  int64(m.Murmur64Hash) ,
    CreatedTime:  int32(m.CreatedTime) ,





}
return r
}

func(m *PB_NotifyView_Flat)ToPB() *PB_NotifyView {
r := &PB_NotifyView{
    NotifyId:  int64(m.NotifyId) ,
    ForUserId:  int32(m.ForUserId) ,
    ActorUserId:  int32(m.ActorUserId) ,
    NotiyTypeEnum:  int32(m.NotiyTypeEnum) ,
    PostId:  int64(m.PostId) ,
    CommentId:  int64(m.CommentId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    Murmur64Hash:  int64(m.Murmur64Hash) ,
    SeenStatus:  int32(m.SeenStatus) ,
    CreatedTime:  int32(m.CreatedTime) ,



}
return r
}

func(m *PB_CommentView_Flat)ToPB() *PB_CommentView {
r := &PB_CommentView{
    CommentId:  int64(m.CommentId) ,
    UserId:  int32(m.UserId) ,
    PostId:  int64(m.PostId) ,
    Text:  m.Text ,
    LikesCount:  int32(m.LikesCount) ,
    CreatedTime:  int32(m.CreatedTime) ,

}
return r
}

func(m *PB_PostView_Flat)ToPB() *PB_PostView {
r := &PB_PostView{
    PostId:  int64(m.PostId) ,
    UserId:  int32(m.UserId) ,

    Text:  m.Text ,
    RichText:  m.RichText ,
    MediaCount:  int32(m.MediaCount) ,
    SharedTo:  int32(m.SharedTo) ,
    DisableComment:  int32(m.DisableComment) ,
    HasTag:  int32(m.HasTag) ,
    CommentsCount:  int32(m.CommentsCount) ,
    LikesCount:  int32(m.LikesCount) ,
    ViewsCount:  int32(m.ViewsCount) ,
    EditedTime:  int32(m.EditedTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
    ReSharedPostId:  int64(m.ReSharedPostId) ,
    DidILiked:  m.DidILiked ,
    DidIReShared:  m.DidIReShared ,




}
return r
}

func(m *PB_ChatView_Flat)ToPB() *PB_ChatView {
r := &PB_ChatView{
    ChatId:  int64(m.ChatId) ,
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomType:  int32(m.RoomType) ,
    UserId:  int32(m.UserId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    HashTagId:  int64(m.HashTagId) ,
    StartedByMe:  int32(m.StartedByMe) ,
    Title:  m.Title ,
    PinTime:  int64(m.PinTime) ,
    FromMsgId:  int64(m.FromMsgId) ,
    Seq:  int32(m.Seq) ,
    LastMsgId:  int64(m.LastMsgId) ,
    LastMsgStatus:  int32(m.LastMsgStatus) ,
    SeenSeq:  int32(m.SeenSeq) ,
    SeenMsgId:  int64(m.SeenMsgId) ,
    Left:  int32(m.Left) ,
    Creator:  int32(m.Creator) ,
    Kicked:  int32(m.Kicked) ,
    Admin:  int32(m.Admin) ,
    Deactivated:  int32(m.Deactivated) ,
    VersionTime:  int32(m.VersionTime) ,
    SortTime:  int32(m.SortTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
    DraftText:  m.DraftText ,
    DratReplyToMsgId:  int64(m.DratReplyToMsgId) ,
    IsMute:  int32(m.IsMute) ,




}
return r
}

func(m *PB_GroupView_Flat)ToPB() *PB_GroupView {
r := &PB_GroupView{
    GroupId:  int64(m.GroupId) ,
    GroupKey:  m.GroupKey ,
    GroupName:  m.GroupName ,
    UserName:  m.UserName ,
    IsSuperGroup:  int32(m.IsSuperGroup) ,
    HashTagId:  int64(m.HashTagId) ,
    CreatorUserId:  int32(m.CreatorUserId) ,
    GroupPrivacy:  int32(m.GroupPrivacy) ,
    HistoryViewAble:  int32(m.HistoryViewAble) ,
    Seq:  int64(m.Seq) ,
    LastMsgId:  int64(m.LastMsgId) ,
    PinedMsgId:  int64(m.PinedMsgId) ,
    AvatarRefId:  int64(m.AvatarRefId) ,
    AvatarCount:  int32(m.AvatarCount) ,
    About:  m.About ,
    InviteLink:  m.InviteLink ,
    MembersCount:  int32(m.MembersCount) ,
    SortTime:  int32(m.SortTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_GroupMemberView_Flat)ToPB() *PB_GroupMemberView {
r := &PB_GroupMemberView{
    OrderId:  int64(m.OrderId) ,
    GroupId:  int64(m.GroupId) ,
    UserId:  int32(m.UserId) ,
    ByUserId:  int32(m.ByUserId) ,
    GroupRole:  int32(m.GroupRole) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_MessageView_Flat)ToPB() *PB_MessageView {
r := &PB_MessageView{
    RoomKey:  m.RoomKey ,
    MessageId:  int64(m.MessageId) ,
    UserId:  int32(m.UserId) ,
    FileRefId:  int64(m.FileRefId) ,
    MessageType:  int32(m.MessageType) ,
    Text:  m.Text ,
    Hiden:  int32(m.Hiden) ,
    Seq:  int32(m.Seq) ,
    ForwardedMsgId:  int64(m.ForwardedMsgId) ,
    PostId:  int64(m.PostId) ,
    StickerId:  int64(m.StickerId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    DeliveredTime:  int32(m.DeliveredTime) ,
    SeenTime:  int32(m.SeenTime) ,
    DeliviryStatus:  int32(m.DeliviryStatus) ,
    ReplyToMessageId:  int64(m.ReplyToMessageId) ,
    ViewsCount:  int64(m.ViewsCount) ,
    EditTime:  int32(m.EditTime) ,
    Ttl:  int32(m.Ttl) ,

}
return r
}

func(m *PB_FileRedView_Flat)ToPB() *PB_FileRedView {
r := &PB_FileRedView{
    FileRefId:  int64(m.FileRefId) ,
    UserId:  int64(m.UserId) ,
    Name:  m.Name ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    UrlSource:  m.UrlSource ,
}
return r
}

func(m *PB_UserView_Flat)ToPB() *PB_UserView {
r := &PB_UserView{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    AvatarRefId:  int64(m.AvatarRefId) ,
    ProfilePrivacy:  int32(m.ProfilePrivacy) ,
    Phone:  int64(m.Phone) ,
    About:  m.About ,
    FollowersCount:  int32(m.FollowersCount) ,
    FollowingCount:  int32(m.FollowingCount) ,
    PostsCount:  int32(m.PostsCount) ,
    MediaCount:  int32(m.MediaCount) ,

    LastActiveTime:  int32(m.LastActiveTime) ,
    LastActiveTimeShow:  m.LastActiveTimeShow ,

}
return r
}

func(m *PB_SettingNotificationView_Flat)ToPB() *PB_SettingNotificationView {
r := &PB_SettingNotificationView{
}
return r
}

func(m *PB_AppConfig_Flat)ToPB() *PB_AppConfig {
r := &PB_AppConfig{
    DeprecatedClient:  m.DeprecatedClient ,
    HasNewUpdate:  m.HasNewUpdate ,
}
return r
}

func(m *PB_UserProfileView_Flat)ToPB() *PB_UserProfileView {
r := &PB_UserProfileView{
}
return r
}

func(m *PB_UserViewRowify_Flat)ToPB() *PB_UserViewRowify {
r := &PB_UserViewRowify{
    Id:  int64(m.Id) ,
    CreatedTime:  int32(m.CreatedTime) ,

}
return r
}

func(m *PB_PostViewRowify_Flat)ToPB() *PB_PostViewRowify {
r := &PB_PostViewRowify{
    Id:  int64(m.Id) ,

}
return r
}

func(m *PB_TagView_Flat)ToPB() *PB_TagView {
r := &PB_TagView{
    TagId:  int64(m.TagId) ,
    Name:  m.Name ,
    Count:  int32(m.Count) ,
    TagStatusEnum:  int32(m.TagStatusEnum) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_TopTagWithSamplePosts_Flat)ToPB() *PB_TopTagWithSamplePosts {
r := &PB_TopTagWithSamplePosts{


}
return r
}

func(m *PB_SelfUserView_Flat)ToPB() *PB_SelfUserView {
r := &PB_SelfUserView{

    ProfilePrivacy:  int32(m.ProfilePrivacy) ,
    OnlinePrivacy:  int32(m.OnlinePrivacy) ,
    CallPrivacy:  int32(m.CallPrivacy) ,
    AddToGroupPrivacy:  int32(m.AddToGroupPrivacy) ,
    SeenMessagePrivacy:  int32(m.SeenMessagePrivacy) ,

}
return r
}

func(m *PB_Error_Flat)ToPB() *PB_Error {
r := &PB_Error{

    ShowError:  m.ShowError ,
    ErrorMessage:  m.ErrorMessage ,
}
return r
}



///// folding ///

var PB_RoomsChanges__FOlD = &PB_RoomsChanges{
        VersionTime:  0 ,


}


var PB_PushChanges__FOlD = &PB_PushChanges{


        InvalidateCacheForRoomKeys:  "" ,
        InvalidateAllChatCache:  0 ,
        InvalidateAllSocialCache:  0 ,
}


var PB_CommandToServer__FOlD = &PB_CommandToServer{
        ClientCallId:  0 ,
        Command:  "" ,
        RespondReached:  false ,
        Data:  []byte{} ,
}


var PB_CommandToClient__FOlD = &PB_CommandToClient{
        ServerCallId:  0 ,
        Command:  "" ,
        RespondReached:  false ,
        Data:  []byte{} ,
}


var PB_CommandReachedToServer__FOlD = &PB_CommandReachedToServer{
        ClientCallId:  0 ,
}


var PB_CommandReachedToClient__FOlD = &PB_CommandReachedToClient{
        ServerCallId:  0 ,
}


var PB_ResponseToClient__FOlD = &PB_ResponseToClient{
        ClientCallId:  0 ,
        PBClass:  "" ,
        RpcFullName:  "" ,
        Data:  []byte{} ,
}


var RPC_Auth_Types__FOlD = &RPC_Auth_Types{
}


var PB_RPC_Chat_Types__FOlD = &PB_RPC_Chat_Types{
}


var RPC_General_Types__FOlD = &RPC_General_Types{
}


var RPC_Page_Types__FOlD = &RPC_Page_Types{
}


var RPC_Social_Types__FOlD = &RPC_Social_Types{
}


var RPC_User_Types__FOlD = &RPC_User_Types{
}


var PB_Action__FOlD = &PB_Action{
        ActionId:  0 ,
        ActorUserId:  0 ,
        ActionType:  0 ,
        PeerUserId:  0 ,
        PostId:  0 ,
        CommentId:  0 ,
        Murmur64Hash:  0 ,
        CreatedTime:  0 ,
}


var PB_Blocked__FOlD = &PB_Blocked{
        Id:  0 ,
        UserId:  0 ,
        BlockedUserId:  0 ,
        CreatedTime:  0 ,
}


var PB_Comment__FOlD = &PB_Comment{
        CommentId:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        Text:  "" ,
        LikesCount:  0 ,
        IsEdited:  0 ,
        CreatedTime:  0 ,
}


var PB_CommentDeleted__FOlD = &PB_CommentDeleted{
        CommentId:  0 ,
        UserId:  0 ,
}


var PB_Event__FOlD = &PB_Event{
        EventId:  0 ,
        EventType:  0 ,
        ByUserId:  0 ,
        PeerUserId:  0 ,
        PostId:  0 ,
        CommentId:  0 ,
        HashTagId:  0 ,
        GroupId:  0 ,
        ActionId:  0 ,
        ChatId:  0 ,
        ChatKey:  "" ,
        MessageId:  0 ,
        ReSharedId:  0 ,
        Murmur64Hash:  0 ,
}


var PB_Followed__FOlD = &PB_Followed{
        Id:  0 ,
        UserId:  0 ,
        FollowedUserId:  0 ,
        CreatedTime:  0 ,
}


var PB_Likes__FOlD = &PB_Likes{
        Id:  0 ,
        PostId:  0 ,
        UserId:  0 ,
        PostType:  0 ,
        CreatedTime:  0 ,
}


var PB_Notify__FOlD = &PB_Notify{
        NotifyId:  0 ,
        ForUserId:  0 ,
        ActorUserId:  0 ,
        NotifyType:  0 ,
        PostId:  0 ,
        CommentId:  0 ,
        PeerUserId:  0 ,
        Murmur64Hash:  0 ,
        SeenStatus:  0 ,
        CreatedTime:  0 ,
}


var PB_NotifyRemoved__FOlD = &PB_NotifyRemoved{
        Murmur64Hash:  0 ,
        ForUserId:  0 ,
        Id:  0 ,
}


var PB_PhoneContacts__FOlD = &PB_PhoneContacts{
        Id:  0 ,
        UserId:  0 ,
        ClientId:  0 ,
        Phone:  "" ,
        FirstName:  "" ,
        LastName:  "" ,
}


var PB_Post__FOlD = &PB_Post{
        PostId:  0 ,
        UserId:  0 ,
        PostType:  0 ,
        MediaId:  0 ,
        FileRefId:  0 ,
        PostKey:  "" ,
        Text:  "" ,
        RichText:  "" ,
        MediaCount:  0 ,
        SharedTo:  0 ,
        DisableComment:  0 ,
        Via:  0 ,
        Seq:  0 ,
        CommentsCount:  0 ,
        LikesCount:  0 ,
        ViewsCount:  0 ,
        EditedTime:  0 ,
        CreatedTime:  0 ,
}


var PB_PostCount__FOlD = &PB_PostCount{
        PostId:  0 ,
        CommentsCount:  0 ,
        LikesCount:  0 ,
        ViewsCount:  0 ,
        ReSharedCount:  0 ,
        ChatSharedCount:  0 ,
}


var PB_PostDeleted__FOlD = &PB_PostDeleted{
        PostId:  0 ,
        UserId:  0 ,
}


var PB_PostKeys__FOlD = &PB_PostKeys{
        Id:  0 ,
        PostKeyStr:  "" ,
        Used:  0 ,
        RandShard:  0 ,
}


var PB_PostLink__FOlD = &PB_PostLink{
        LinkId:  0 ,
        LinkUrl:  "" ,
}


var PB_PostMedia__FOlD = &PB_PostMedia{
        MediaId:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        AlbumId:  0 ,
        MediaTypeEnum:  0 ,
        Width:  0 ,
        Height:  0 ,
        Size:  0 ,
        Duration:  0 ,
        Extension:  "" ,
        Md5Hash:  "" ,
        Color:  "" ,
        CreatedTime:  0 ,
        ViewCount:  0 ,
        Extra:  "" ,
}


var PB_PostReshared__FOlD = &PB_PostReshared{
        ResharedId:  0 ,
        PostId:  0 ,
        ByUserId:  0 ,
        PostUserId:  0 ,
        CreatedTime:  0 ,
}


var PB_ProfileAll__FOlD = &PB_ProfileAll{
        Id:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        IsReShared:  0 ,
}


var PB_ProfileMedia__FOlD = &PB_ProfileMedia{
        Id:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        PostType:  0 ,
}


var PB_ProfileMentioned__FOlD = &PB_ProfileMentioned{
        Id:  0 ,
        ForUserId:  0 ,
        PostId:  0 ,
        PostUserId:  0 ,
        PostType:  0 ,
        CreatedTime:  0 ,
}


var PB_Session__FOlD = &PB_Session{
        Id:  0 ,
        SessionUuid:  "" ,
        UserId:  0 ,
        LastIpAddress:  "" ,
        AppVersion:  0 ,
        ActiveTime:  0 ,
        CreatedTime:  0 ,
}


var PB_SettingNotifications__FOlD = &PB_SettingNotifications{
        UserId:  0 ,
        SocialLedOn:  0 ,
        SocialLedColor:  "" ,
        ReqestToFollowYou:  0 ,
        FollowedYou:  0 ,
        AccptedYourFollowRequest:  0 ,
        YourPostLiked:  0 ,
        YourPostCommented:  0 ,
        MenthenedYouInPost:  0 ,
        MenthenedYouInComment:  0 ,
        YourContactsJoined:  0 ,
        DirectMessage:  0 ,
        DirectAlert:  0 ,
        DirectPerview:  0 ,
        DirectLedOn:  0 ,
        DirectLedColor:  0 ,
        DirectVibrate:  0 ,
        DirectPopup:  0 ,
        DirectSound:  0 ,
        DirectPriority:  0 ,
}


var PB_Sms__FOlD = &PB_Sms{
        Id:  0 ,
        Hash:  "" ,
        AppUuid:  "" ,
        ClientPhone:  "" ,
        GenratedCode:  0 ,
        SmsSenderNumber:  "" ,
        SmsSendStatues:  "" ,
        SmsHttpBody:  "" ,
        Err:  "" ,
        Carrier:  "" ,
        Country:  []byte{} ,
        IsValidPhone:  0 ,
        IsConfirmed:  0 ,
        IsLogin:  0 ,
        IsRegister:  0 ,
        RetriedCount:  0 ,
        TTL:  0 ,
}


var PB_Tag__FOlD = &PB_Tag{
        TagId:  0 ,
        Name:  "" ,
        Count:  0 ,
        TagStatusEnum:  0 ,
        IsBlocked:  0 ,
        GroupId:  0 ,
        CreatedTime:  0 ,
}


var PB_TagPost__FOlD = &PB_TagPost{
        Id:  0 ,
        TagId:  0 ,
        PostId:  0 ,
        PostType:  0 ,
        CreatedTime:  0 ,
}


var PB_TriggerLog__FOlD = &PB_TriggerLog{
        Id:  0 ,
        ModelName:  "" ,
        ChangeType:  "" ,
        TargetId:  0 ,
        TargetStr:  "" ,
        CreatedSe:  0 ,
}


var PB_User__FOlD = &PB_User{
        UserId:  0 ,
        UserName:  "" ,
        UserNameLower:  "" ,
        FirstName:  "" ,
        LastName:  "" ,
        IsVerified:  0 ,
        AvatarId:  0 ,
        AccessHash:  0 ,
        ProfilePrivacy:  0 ,
        OnlinePrivacy:  0 ,
        CallPrivacy:  0 ,
        AddToGroupPrivacy:  0 ,
        SeenMessagePrivacy:  0 ,
        Phone:  "" ,
        Email:  "" ,
        About:  "" ,
        PasswordHash:  "" ,
        PasswordSalt:  "" ,
        PostSeq:  0 ,
        FollowersCount:  0 ,
        FollowingCount:  0 ,
        PostsCount:  0 ,
        MediaCount:  0 ,
        PhotoCount:  0 ,
        VideoCount:  0 ,
        GifCount:  0 ,
        AudioCount:  0 ,
        VoiceCount:  0 ,
        FileCount:  0 ,
        LinkCount:  0 ,
        BoardCount:  0 ,
        PinedCount:  0 ,
        LikesCount:  0 ,
        ResharedCount:  0 ,
        LastPostTime:  0 ,
        CreatedTime:  0 ,
        VersionTime:  0 ,
        IsDeleted:  0 ,
        IsBanned:  0 ,
}


var PB_UserRelation__FOlD = &PB_UserRelation{
        RelNanoId:  0 ,
        UserId:  0 ,
        PeerUserId:  0 ,
        Follwing:  0 ,
        Followed:  0 ,
        InContacts:  0 ,
        MutualContact:  0 ,
        IsFavorite:  0 ,
        Notify:  0 ,
}


var PB_Chat__FOlD = &PB_Chat{
        ChatId:  0 ,
        ChatKey:  "" ,
        RoomKey:  "" ,
        RoomType:  0 ,
        UserId:  0 ,
        PeerUserId:  0 ,
        GroupId:  0 ,
        HashTagId:  0 ,
        Title:  "" ,
        PinTimeMs:  0 ,
        FromMsgId:  0 ,
        UnseenCount:  0 ,
        Seq:  0 ,
        LastMsgId:  0 ,
        LastMyMsgStatus:  0 ,
        MyLastSeenSeq:  0 ,
        MyLastSeenMsgId:  0 ,
        PeerLastSeenMsgId:  0 ,
        MyLastDeliveredSeq:  0 ,
        MyLastDeliveredMsgId:  0 ,
        PeerLastDeliveredMsgId:  0 ,
        IsActive:  0 ,
        IsLeft:  0 ,
        IsCreator:  0 ,
        IsKicked:  0 ,
        IsAdmin:  0 ,
        IsDeactivated:  0 ,
        IsMuted:  0 ,
        MuteUntil:  0 ,
        VersionTimeMs:  0 ,
        VersionSeq:  0 ,
        OrderTime:  0 ,
        CreatedTime:  0 ,
        DraftText:  "" ,
        DratReplyToMsgId:  0 ,
}


var PB_ChatDeleted__FOlD = &PB_ChatDeleted{
        ChatId:  0 ,
        RoomKey:  "" ,
}


var PB_ChatLastMessage__FOlD = &PB_ChatLastMessage{
        ChatIdGroupId:  "" ,
        LastMsgPb:  []byte{} ,
}


var PB_ChatUserVersion__FOlD = &PB_ChatUserVersion{
        UserId:  0 ,
        ChatId:  0 ,
        VersionTimeNano:  0 ,
}


var PB_Group__FOlD = &PB_Group{
        GroupId:  0 ,
        GroupKey:  "" ,
        GroupName:  "" ,
        UserName:  "" ,
        IsSuperGroup:  0 ,
        HashTagId:  0 ,
        CreatorUserId:  0 ,
        GroupPrivacy:  0 ,
        HistoryViewAble:  0 ,
        Seq:  0 ,
        LastMsgId:  0 ,
        PinedMsgId:  0 ,
        AvatarRefId:  0 ,
        AvatarCount:  0 ,
        About:  "" ,
        InviteLinkHash:  "" ,
        MembersCount:  0 ,
        AdminsCount:  0 ,
        ModeratorCounts:  0 ,
        SortTime:  0 ,
        CreatedTime:  0 ,
        IsMute:  0 ,
        IsDeleted:  0 ,
        IsBanned:  0 ,
}


var PB_GroupMember__FOlD = &PB_GroupMember{
        OrderId:  0 ,
        GroupId:  0 ,
        UserId:  0 ,
        ByUserId:  0 ,
        GroupRole:  0 ,
        CreatedTime:  0 ,
}


var PB_GroupOrderdUser__FOlD = &PB_GroupOrderdUser{
        OrderId:  0 ,
        GroupId:  0 ,
        UserId:  0 ,
}


var PB_GroupPinedMsg__FOlD = &PB_GroupPinedMsg{
        MsgId:  0 ,
        MsgPb:  []byte{} ,
}


var PB_FileMsg__FOlD = &PB_FileMsg{
        Id:  0 ,
        AccessHash:  0 ,
        FileType:  0 ,
        Width:  0 ,
        Height:  0 ,
        Extension:  "" ,
        UserId:  0 ,
        DataThumb:  []byte{} ,
        Data:  []byte{} ,
}


var PB_FilePost__FOlD = &PB_FilePost{
        Id:  0 ,
        AccessHash:  0 ,
        FileType:  0 ,
        Width:  0 ,
        Height:  0 ,
        Extension:  "" ,
        UserId:  0 ,
        DataThumb:  []byte{} ,
        Data:  []byte{} ,
}


var PB_ActionFanout__FOlD = &PB_ActionFanout{
        OrderId:  0 ,
        ForUserId:  0 ,
        ActionId:  0 ,
        ActorUserId:  0 ,
}


var PB_HomeFanout__FOlD = &PB_HomeFanout{
        OrderId:  0 ,
        ForUserId:  0 ,
        PostId:  0 ,
        PostUserId:  0 ,
        ResharedId:  0 ,
}


var PB_SuggestedTopPosts__FOlD = &PB_SuggestedTopPosts{
        Id:  0 ,
        PostId:  0 ,
}


var PB_SuggestedUser__FOlD = &PB_SuggestedUser{
        Id:  0 ,
        UserId:  0 ,
        TargetId:  0 ,
        Weight:  0.0 ,
        CreatedTime:  0 ,
}


var PB_PushChat__FOlD = &PB_PushChat{
        PushId:  0 ,
        ToUserId:  0 ,
        PushTypeId:  0 ,
        RoomKey:  "" ,
        ChatKey:  "" ,
        Seq:  0 ,
        UnseenCount:  0 ,
        FromHighMessageId:  0 ,
        ToLowMessageId:  0 ,
        MessageId:  0 ,
        MessageFileId:  0 ,
        MessagePb:  []byte{} ,
        MessageJson:  "" ,
        CreatedTime:  0 ,
}


var PB_HTTPRPCLog__FOlD = &PB_HTTPRPCLog{
        Id:  0 ,
        Time:  "" ,
        MethodFull:  "" ,
        MethodParent:  "" ,
        UserId:  0 ,
        SessionId:  "" ,
        StatusCode:  0 ,
        InputSize:  0 ,
        OutputSize:  0 ,
        ReqestJson:  "" ,
        ResponseJson:  "" ,
        ReqestParamJson:  "" ,
        ResponseMsgJson:  "" ,
}


var PB_MetricLog__FOlD = &PB_MetricLog{
        Id:  0 ,
        InstanceId:  0 ,
        StartFrom:  "" ,
        EndTo:  "" ,
        StartTime:  0 ,
        Duration:  "" ,
        MetericsJson:  "" ,
}


var PB_XfileServiceInfoLog__FOlD = &PB_XfileServiceInfoLog{
        Id:  0 ,
        InstanceId:  0 ,
        Url:  "" ,
        CreatedTime:  "" ,
}


var PB_XfileServiceMetricLog__FOlD = &PB_XfileServiceMetricLog{
        Id:  0 ,
        InstanceId:  0 ,
        MetricJson:  "" ,
}


var PB_XfileServiceRequestLog__FOlD = &PB_XfileServiceRequestLog{
        Id:  0 ,
        LocalSeq:  0 ,
        InstanceId:  0 ,
        Url:  "" ,
        HttpCode:  0 ,
        CreatedTime:  "" ,
}


var PB_InvalidateCache__FOlD = &PB_InvalidateCache{
        OrderId:  0 ,
        CacheKey:  "" ,
}


var PB_MediaView__FOlD = &PB_MediaView{
}


var PB_ActionView__FOlD = &PB_ActionView{
        ActionId:  0 ,
        ActorUserId:  0 ,
        ActionTypeEnum:  0 ,
        PeerUserId:  0 ,
        PostId:  0 ,
        CommentId:  0 ,
        Murmur64Hash:  0 ,
        CreatedTime:  0 ,





}


var PB_NotifyView__FOlD = &PB_NotifyView{
        NotifyId:  0 ,
        ForUserId:  0 ,
        ActorUserId:  0 ,
        NotiyTypeEnum:  0 ,
        PostId:  0 ,
        CommentId:  0 ,
        PeerUserId:  0 ,
        Murmur64Hash:  0 ,
        SeenStatus:  0 ,
        CreatedTime:  0 ,



}


var PB_CommentView__FOlD = &PB_CommentView{
        CommentId:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        Text:  "" ,
        LikesCount:  0 ,
        CreatedTime:  0 ,

}


var PB_PostView__FOlD = &PB_PostView{
        PostId:  0 ,
        UserId:  0 ,

        Text:  "" ,
        RichText:  "" ,
        MediaCount:  0 ,
        SharedTo:  0 ,
        DisableComment:  0 ,
        HasTag:  0 ,
        CommentsCount:  0 ,
        LikesCount:  0 ,
        ViewsCount:  0 ,
        EditedTime:  0 ,
        CreatedTime:  0 ,
        ReSharedPostId:  0 ,
        DidILiked:  false ,
        DidIReShared:  false ,




}


var PB_ChatView__FOlD = &PB_ChatView{
        ChatId:  0 ,
        ChatKey:  "" ,
        RoomKey:  "" ,
        RoomType:  0 ,
        UserId:  0 ,
        PeerUserId:  0 ,
        GroupId:  0 ,
        HashTagId:  0 ,
        StartedByMe:  0 ,
        Title:  "" ,
        PinTime:  0 ,
        FromMsgId:  0 ,
        Seq:  0 ,
        LastMsgId:  0 ,
        LastMsgStatus:  0 ,
        SeenSeq:  0 ,
        SeenMsgId:  0 ,
        Left:  0 ,
        Creator:  0 ,
        Kicked:  0 ,
        Admin:  0 ,
        Deactivated:  0 ,
        VersionTime:  0 ,
        SortTime:  0 ,
        CreatedTime:  0 ,
        DraftText:  "" ,
        DratReplyToMsgId:  0 ,
        IsMute:  0 ,




}


var PB_GroupView__FOlD = &PB_GroupView{
        GroupId:  0 ,
        GroupKey:  "" ,
        GroupName:  "" ,
        UserName:  "" ,
        IsSuperGroup:  0 ,
        HashTagId:  0 ,
        CreatorUserId:  0 ,
        GroupPrivacy:  0 ,
        HistoryViewAble:  0 ,
        Seq:  0 ,
        LastMsgId:  0 ,
        PinedMsgId:  0 ,
        AvatarRefId:  0 ,
        AvatarCount:  0 ,
        About:  "" ,
        InviteLink:  "" ,
        MembersCount:  0 ,
        SortTime:  0 ,
        CreatedTime:  0 ,
}


var PB_GroupMemberView__FOlD = &PB_GroupMemberView{
        OrderId:  0 ,
        GroupId:  0 ,
        UserId:  0 ,
        ByUserId:  0 ,
        GroupRole:  0 ,
        CreatedTime:  0 ,
}


var PB_MessageView__FOlD = &PB_MessageView{
        RoomKey:  "" ,
        MessageId:  0 ,
        UserId:  0 ,
        FileRefId:  0 ,
        MessageType:  0 ,
        Text:  "" ,
        Hiden:  0 ,
        Seq:  0 ,
        ForwardedMsgId:  0 ,
        PostId:  0 ,
        StickerId:  0 ,
        CreatedTime:  0 ,
        DeliveredTime:  0 ,
        SeenTime:  0 ,
        DeliviryStatus:  0 ,
        ReplyToMessageId:  0 ,
        ViewsCount:  0 ,
        EditTime:  0 ,
        Ttl:  0 ,

}


var PB_FileRedView__FOlD = &PB_FileRedView{
        FileRefId:  0 ,
        UserId:  0 ,
        Name:  "" ,
        Width:  0 ,
        Height:  0 ,
        Duration:  0 ,
        Extension:  "" ,
        UrlSource:  "" ,
}


var PB_UserView__FOlD = &PB_UserView{
        UserId:  0 ,
        UserName:  "" ,
        FirstName:  "" ,
        LastName:  "" ,
        AvatarRefId:  0 ,
        ProfilePrivacy:  0 ,
        Phone:  0 ,
        About:  "" ,
        FollowersCount:  0 ,
        FollowingCount:  0 ,
        PostsCount:  0 ,
        MediaCount:  0 ,

        LastActiveTime:  0 ,
        LastActiveTimeShow:  "" ,

}


var PB_SettingNotificationView__FOlD = &PB_SettingNotificationView{
}


var PB_AppConfig__FOlD = &PB_AppConfig{
        DeprecatedClient:  false ,
        HasNewUpdate:  false ,
}


var PB_UserProfileView__FOlD = &PB_UserProfileView{
}


var PB_UserViewRowify__FOlD = &PB_UserViewRowify{
        Id:  0 ,
        CreatedTime:  0 ,

}


var PB_PostViewRowify__FOlD = &PB_PostViewRowify{
        Id:  0 ,

}


var PB_TagView__FOlD = &PB_TagView{
        TagId:  0 ,
        Name:  "" ,
        Count:  0 ,
        TagStatusEnum:  0 ,
        CreatedTime:  0 ,
}


var PB_TopTagWithSamplePosts__FOlD = &PB_TopTagWithSamplePosts{


}


var PB_SelfUserView__FOlD = &PB_SelfUserView{

        ProfilePrivacy:  0 ,
        OnlinePrivacy:  0 ,
        CallPrivacy:  0 ,
        AddToGroupPrivacy:  0 ,
        SeenMessagePrivacy:  0 ,

}


var PB_Error__FOlD = &PB_Error{

        ShowError:  false ,
        ErrorMessage:  "" ,
}



*/
