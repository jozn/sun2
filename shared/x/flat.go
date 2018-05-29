package x

type PB_RoomsChanges_Flat struct {
	VersionTime int
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

type PB_Action_Flat struct {
	ActionId       int
	ActorUserId    int
	ActionTypeEnum int
	PeerUserId     int
	PostId         int
	CommentId      int
	Murmur64Hash   int
	CreatedTime    int
}

//ToPB
func (m *PB_Action) ToFlat() *PB_Action_Flat {
	r := &PB_Action_Flat{
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
func (m *PB_Action_Flat) ToPB() *PB_Action {
	r := &PB_Action{
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
var PB_Action__FOlD = &PB_Action{
	ActionId:       0,
	ActorUserId:    0,
	ActionTypeEnum: 0,
	PeerUserId:     0,
	PostId:         0,
	CommentId:      0,
	Murmur64Hash:   0,
	CreatedTime:    0,
}

type PB_Comment_Flat struct {
	CommentId   int
	UserId      int
	PostId      int
	Text        string
	LikesCount  int
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
	ActionId     int
	Murmur64Hash int
	ChatKey      string
	MessageId    int
	ReSharedId   int
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
		ActionId:     int(m.ActionId),
		Murmur64Hash: int(m.Murmur64Hash),
		ChatKey:      m.ChatKey,
		MessageId:    int(m.MessageId),
		ReSharedId:   int(m.ReSharedId),
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
		ActionId:     int64(m.ActionId),
		Murmur64Hash: int64(m.Murmur64Hash),
		ChatKey:      m.ChatKey,
		MessageId:    int64(m.MessageId),
		ReSharedId:   int64(m.ReSharedId),
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
	ActionId:     0,
	Murmur64Hash: 0,
	ChatKey:      "",
	MessageId:    0,
	ReSharedId:   0,
}

type PB_Like_Flat struct {
	Id           int
	PostId       int
	PostTypeEnum int
	UserId       int
	LikeEnum     int
	CreatedTime  int
}

//ToPB
func (m *PB_Like) ToFlat() *PB_Like_Flat {
	r := &PB_Like_Flat{
		Id:           int(m.Id),
		PostId:       int(m.PostId),
		PostTypeEnum: int(m.PostTypeEnum),
		UserId:       int(m.UserId),
		LikeEnum:     int(m.LikeEnum),
		CreatedTime:  int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Like_Flat) ToPB() *PB_Like {
	r := &PB_Like{
		Id:           int64(m.Id),
		PostId:       int64(m.PostId),
		PostTypeEnum: int32(m.PostTypeEnum),
		UserId:       int32(m.UserId),
		LikeEnum:     int32(m.LikeEnum),
		CreatedTime:  int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Like__FOlD = &PB_Like{
	Id:           0,
	PostId:       0,
	PostTypeEnum: 0,
	UserId:       0,
	LikeEnum:     0,
	CreatedTime:  0,
}

type PB_Notify_Flat struct {
	NotifyId       int
	ForUserId      int
	ActorUserId    int
	NotifyTypeEnum int
	PostId         int
	CommentId      int
	PeerUserId     int
	Murmur64Hash   int
	SeenStatus     int
	CreatedTime    int
}

//ToPB
func (m *PB_Notify) ToFlat() *PB_Notify_Flat {
	r := &PB_Notify_Flat{
		NotifyId:       int(m.NotifyId),
		ForUserId:      int(m.ForUserId),
		ActorUserId:    int(m.ActorUserId),
		NotifyTypeEnum: int(m.NotifyTypeEnum),
		PostId:         int(m.PostId),
		CommentId:      int(m.CommentId),
		PeerUserId:     int(m.PeerUserId),
		Murmur64Hash:   int(m.Murmur64Hash),
		SeenStatus:     int(m.SeenStatus),
		CreatedTime:    int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Notify_Flat) ToPB() *PB_Notify {
	r := &PB_Notify{
		NotifyId:       int64(m.NotifyId),
		ForUserId:      int32(m.ForUserId),
		ActorUserId:    int32(m.ActorUserId),
		NotifyTypeEnum: int32(m.NotifyTypeEnum),
		PostId:         int64(m.PostId),
		CommentId:      int64(m.CommentId),
		PeerUserId:     int32(m.PeerUserId),
		Murmur64Hash:   int64(m.Murmur64Hash),
		SeenStatus:     int32(m.SeenStatus),
		CreatedTime:    int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Notify__FOlD = &PB_Notify{
	NotifyId:       0,
	ForUserId:      0,
	ActorUserId:    0,
	NotifyTypeEnum: 0,
	PostId:         0,
	CommentId:      0,
	PeerUserId:     0,
	Murmur64Hash:   0,
	SeenStatus:     0,
	CreatedTime:    0,
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

type PB_PhoneContact_Flat struct {
	Id        int
	UserId    int
	ClientId  int
	Phone     string
	FirstName string
	LastName  string
}

//ToPB
func (m *PB_PhoneContact) ToFlat() *PB_PhoneContact_Flat {
	r := &PB_PhoneContact_Flat{
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
func (m *PB_PhoneContact_Flat) ToPB() *PB_PhoneContact {
	r := &PB_PhoneContact{
		Id:        int32(m.Id),
		UserId:    int32(m.UserId),
		ClientId:  int64(m.ClientId),
		Phone:     m.Phone,
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}
	return r
}

//folding
var PB_PhoneContact__FOlD = &PB_PhoneContact{
	Id:        0,
	UserId:    0,
	ClientId:  0,
	Phone:     "",
	FirstName: "",
	LastName:  "",
}

type PB_Post_Flat struct {
	PostId           int
	UserId           int
	PostTypeEnum     int
	PostCategoryEnum int
	MediaId          int
	PostKey          string
	Text             string
	RichText         string
	MediaCount       int
	SharedTo         int
	DisableComment   int
	Source           int
	HasTag           int
	Seq              int
	CommentsCount    int
	LikesCount       int
	ViewsCount       int
	EditedTime       int
	CreatedTime      int
	ReSharedPostId   int
}

//ToPB
func (m *PB_Post) ToFlat() *PB_Post_Flat {
	r := &PB_Post_Flat{
		PostId:           int(m.PostId),
		UserId:           int(m.UserId),
		PostTypeEnum:     int(m.PostTypeEnum),
		PostCategoryEnum: int(m.PostCategoryEnum),
		MediaId:          int(m.MediaId),
		PostKey:          m.PostKey,
		Text:             m.Text,
		RichText:         m.RichText,
		MediaCount:       int(m.MediaCount),
		SharedTo:         int(m.SharedTo),
		DisableComment:   int(m.DisableComment),
		Source:           int(m.Source),
		HasTag:           int(m.HasTag),
		Seq:              int(m.Seq),
		CommentsCount:    int(m.CommentsCount),
		LikesCount:       int(m.LikesCount),
		ViewsCount:       int(m.ViewsCount),
		EditedTime:       int(m.EditedTime),
		CreatedTime:      int(m.CreatedTime),
		ReSharedPostId:   int(m.ReSharedPostId),
	}
	return r
}

//ToPB
func (m *PB_Post_Flat) ToPB() *PB_Post {
	r := &PB_Post{
		PostId:           int64(m.PostId),
		UserId:           int32(m.UserId),
		PostTypeEnum:     int32(m.PostTypeEnum),
		PostCategoryEnum: int32(m.PostCategoryEnum),
		MediaId:          int64(m.MediaId),
		PostKey:          m.PostKey,
		Text:             m.Text,
		RichText:         m.RichText,
		MediaCount:       int32(m.MediaCount),
		SharedTo:         int32(m.SharedTo),
		DisableComment:   int32(m.DisableComment),
		Source:           int32(m.Source),
		HasTag:           int32(m.HasTag),
		Seq:              int32(m.Seq),
		CommentsCount:    int32(m.CommentsCount),
		LikesCount:       int32(m.LikesCount),
		ViewsCount:       int32(m.ViewsCount),
		EditedTime:       int32(m.EditedTime),
		CreatedTime:      int32(m.CreatedTime),
		ReSharedPostId:   int64(m.ReSharedPostId),
	}
	return r
}

//folding
var PB_Post__FOlD = &PB_Post{
	PostId:           0,
	UserId:           0,
	PostTypeEnum:     0,
	PostCategoryEnum: 0,
	MediaId:          0,
	PostKey:          "",
	Text:             "",
	RichText:         "",
	MediaCount:       0,
	SharedTo:         0,
	DisableComment:   0,
	Source:           0,
	HasTag:           0,
	Seq:              0,
	CommentsCount:    0,
	LikesCount:       0,
	ViewsCount:       0,
	EditedTime:       0,
	CreatedTime:      0,
	ReSharedPostId:   0,
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

type PB_PostKey_Flat struct {
	Id         int
	PostKeyStr string
	Used       int
}

//ToPB
func (m *PB_PostKey) ToFlat() *PB_PostKey_Flat {
	r := &PB_PostKey_Flat{
		Id:         int(m.Id),
		PostKeyStr: m.PostKeyStr,
		Used:       int(m.Used),
	}
	return r
}

//ToPB
func (m *PB_PostKey_Flat) ToPB() *PB_PostKey {
	r := &PB_PostKey{
		Id:         int32(m.Id),
		PostKeyStr: m.PostKeyStr,
		Used:       int32(m.Used),
	}
	return r
}

//folding
var PB_PostKey__FOlD = &PB_PostKey{
	Id:         0,
	PostKeyStr: "",
	Used:       0,
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

type PB_PostMentioned_Flat struct {
	MentionedId      int
	ForUserId        int
	PostId           int
	PostUserId       int
	PostTypeEnum     int
	PostCategoryEnum int
	CreatedTime      int
}

//ToPB
func (m *PB_PostMentioned) ToFlat() *PB_PostMentioned_Flat {
	r := &PB_PostMentioned_Flat{
		MentionedId:      int(m.MentionedId),
		ForUserId:        int(m.ForUserId),
		PostId:           int(m.PostId),
		PostUserId:       int(m.PostUserId),
		PostTypeEnum:     int(m.PostTypeEnum),
		PostCategoryEnum: int(m.PostCategoryEnum),
		CreatedTime:      int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_PostMentioned_Flat) ToPB() *PB_PostMentioned {
	r := &PB_PostMentioned{
		MentionedId:      int64(m.MentionedId),
		ForUserId:        int32(m.ForUserId),
		PostId:           int64(m.PostId),
		PostUserId:       int32(m.PostUserId),
		PostTypeEnum:     int32(m.PostTypeEnum),
		PostCategoryEnum: int32(m.PostCategoryEnum),
		CreatedTime:      int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_PostMentioned__FOlD = &PB_PostMentioned{
	MentionedId:      0,
	ForUserId:        0,
	PostId:           0,
	PostUserId:       0,
	PostTypeEnum:     0,
	PostCategoryEnum: 0,
	CreatedTime:      0,
}

type PB_PostReshared_Flat struct {
	ResharedId       int
	ByUserId         int
	PostId           int
	PostUserId       int
	PostTypeEnum     int
	PostCategoryEnum int
	CreatedTime      int
}

//ToPB
func (m *PB_PostReshared) ToFlat() *PB_PostReshared_Flat {
	r := &PB_PostReshared_Flat{
		ResharedId:       int(m.ResharedId),
		ByUserId:         int(m.ByUserId),
		PostId:           int(m.PostId),
		PostUserId:       int(m.PostUserId),
		PostTypeEnum:     int(m.PostTypeEnum),
		PostCategoryEnum: int(m.PostCategoryEnum),
		CreatedTime:      int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_PostReshared_Flat) ToPB() *PB_PostReshared {
	r := &PB_PostReshared{
		ResharedId:       int64(m.ResharedId),
		ByUserId:         int32(m.ByUserId),
		PostId:           int64(m.PostId),
		PostUserId:       int32(m.PostUserId),
		PostTypeEnum:     int32(m.PostTypeEnum),
		PostCategoryEnum: int32(m.PostCategoryEnum),
		CreatedTime:      int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_PostReshared__FOlD = &PB_PostReshared{
	ResharedId:       0,
	ByUserId:         0,
	PostId:           0,
	PostUserId:       0,
	PostTypeEnum:     0,
	PostCategoryEnum: 0,
	CreatedTime:      0,
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

type PB_SettingClient_Flat struct {
	UserId                    int
	AutoDownloadWifiVoice     int
	AutoDownloadWifiImage     int
	AutoDownloadWifiVideo     int
	AutoDownloadWifiFile      int
	AutoDownloadWifiMusic     int
	AutoDownloadWifiGif       int
	AutoDownloadCellularVoice int
	AutoDownloadCellularImage int
	AutoDownloadCellularVideo int
	AutoDownloadCellularFile  int
	AutoDownloadCellularMusic int
	AutoDownloadCellularGif   int
	AutoDownloadRoamingVoice  int
	AutoDownloadRoamingImage  int
	AutoDownloadRoamingVideo  int
	AutoDownloadRoamingFile   int
	AutoDownloadRoamingMusic  int
	AutoDownloadRoamingGif    int
	SaveToGallery             int
}

//ToPB
func (m *PB_SettingClient) ToFlat() *PB_SettingClient_Flat {
	r := &PB_SettingClient_Flat{
		UserId:                    int(m.UserId),
		AutoDownloadWifiVoice:     int(m.AutoDownloadWifiVoice),
		AutoDownloadWifiImage:     int(m.AutoDownloadWifiImage),
		AutoDownloadWifiVideo:     int(m.AutoDownloadWifiVideo),
		AutoDownloadWifiFile:      int(m.AutoDownloadWifiFile),
		AutoDownloadWifiMusic:     int(m.AutoDownloadWifiMusic),
		AutoDownloadWifiGif:       int(m.AutoDownloadWifiGif),
		AutoDownloadCellularVoice: int(m.AutoDownloadCellularVoice),
		AutoDownloadCellularImage: int(m.AutoDownloadCellularImage),
		AutoDownloadCellularVideo: int(m.AutoDownloadCellularVideo),
		AutoDownloadCellularFile:  int(m.AutoDownloadCellularFile),
		AutoDownloadCellularMusic: int(m.AutoDownloadCellularMusic),
		AutoDownloadCellularGif:   int(m.AutoDownloadCellularGif),
		AutoDownloadRoamingVoice:  int(m.AutoDownloadRoamingVoice),
		AutoDownloadRoamingImage:  int(m.AutoDownloadRoamingImage),
		AutoDownloadRoamingVideo:  int(m.AutoDownloadRoamingVideo),
		AutoDownloadRoamingFile:   int(m.AutoDownloadRoamingFile),
		AutoDownloadRoamingMusic:  int(m.AutoDownloadRoamingMusic),
		AutoDownloadRoamingGif:    int(m.AutoDownloadRoamingGif),
		SaveToGallery:             int(m.SaveToGallery),
	}
	return r
}

//ToPB
func (m *PB_SettingClient_Flat) ToPB() *PB_SettingClient {
	r := &PB_SettingClient{
		UserId:                    int32(m.UserId),
		AutoDownloadWifiVoice:     int32(m.AutoDownloadWifiVoice),
		AutoDownloadWifiImage:     int32(m.AutoDownloadWifiImage),
		AutoDownloadWifiVideo:     int32(m.AutoDownloadWifiVideo),
		AutoDownloadWifiFile:      int32(m.AutoDownloadWifiFile),
		AutoDownloadWifiMusic:     int32(m.AutoDownloadWifiMusic),
		AutoDownloadWifiGif:       int32(m.AutoDownloadWifiGif),
		AutoDownloadCellularVoice: int32(m.AutoDownloadCellularVoice),
		AutoDownloadCellularImage: int32(m.AutoDownloadCellularImage),
		AutoDownloadCellularVideo: int32(m.AutoDownloadCellularVideo),
		AutoDownloadCellularFile:  int32(m.AutoDownloadCellularFile),
		AutoDownloadCellularMusic: int32(m.AutoDownloadCellularMusic),
		AutoDownloadCellularGif:   int32(m.AutoDownloadCellularGif),
		AutoDownloadRoamingVoice:  int32(m.AutoDownloadRoamingVoice),
		AutoDownloadRoamingImage:  int32(m.AutoDownloadRoamingImage),
		AutoDownloadRoamingVideo:  int32(m.AutoDownloadRoamingVideo),
		AutoDownloadRoamingFile:   int32(m.AutoDownloadRoamingFile),
		AutoDownloadRoamingMusic:  int32(m.AutoDownloadRoamingMusic),
		AutoDownloadRoamingGif:    int32(m.AutoDownloadRoamingGif),
		SaveToGallery:             int32(m.SaveToGallery),
	}
	return r
}

//folding
var PB_SettingClient__FOlD = &PB_SettingClient{
	UserId:                    0,
	AutoDownloadWifiVoice:     0,
	AutoDownloadWifiImage:     0,
	AutoDownloadWifiVideo:     0,
	AutoDownloadWifiFile:      0,
	AutoDownloadWifiMusic:     0,
	AutoDownloadWifiGif:       0,
	AutoDownloadCellularVoice: 0,
	AutoDownloadCellularImage: 0,
	AutoDownloadCellularVideo: 0,
	AutoDownloadCellularFile:  0,
	AutoDownloadCellularMusic: 0,
	AutoDownloadCellularGif:   0,
	AutoDownloadRoamingVoice:  0,
	AutoDownloadRoamingImage:  0,
	AutoDownloadRoamingVideo:  0,
	AutoDownloadRoamingFile:   0,
	AutoDownloadRoamingMusic:  0,
	AutoDownloadRoamingGif:    0,
	SaveToGallery:             0,
}

type PB_SettingNotification_Flat struct {
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
func (m *PB_SettingNotification) ToFlat() *PB_SettingNotification_Flat {
	r := &PB_SettingNotification_Flat{
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
func (m *PB_SettingNotification_Flat) ToPB() *PB_SettingNotification {
	r := &PB_SettingNotification{
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
var PB_SettingNotification__FOlD = &PB_SettingNotification{
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

type PB_Tag_Flat struct {
	TagId         int
	Name          string
	Count         int
	TagStatusEnum int
	CreatedTime   int
}

//ToPB
func (m *PB_Tag) ToFlat() *PB_Tag_Flat {
	r := &PB_Tag_Flat{
		TagId:         int(m.TagId),
		Name:          m.Name,
		Count:         int(m.Count),
		TagStatusEnum: int(m.TagStatusEnum),
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
	CreatedTime:   0,
}

type PB_TagPost_Flat struct {
	Id               int
	TagId            int
	PostId           int
	PostTypeEnum     int
	PostCategoryEnum int
	CreatedTime      int
}

//ToPB
func (m *PB_TagPost) ToFlat() *PB_TagPost_Flat {
	r := &PB_TagPost_Flat{
		Id:               int(m.Id),
		TagId:            int(m.TagId),
		PostId:           int(m.PostId),
		PostTypeEnum:     int(m.PostTypeEnum),
		PostCategoryEnum: int(m.PostCategoryEnum),
		CreatedTime:      int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_TagPost_Flat) ToPB() *PB_TagPost {
	r := &PB_TagPost{
		Id:               int64(m.Id),
		TagId:            int32(m.TagId),
		PostId:           int32(m.PostId),
		PostTypeEnum:     int32(m.PostTypeEnum),
		PostCategoryEnum: int32(m.PostCategoryEnum),
		CreatedTime:      int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_TagPost__FOlD = &PB_TagPost{
	Id:               0,
	TagId:            0,
	PostId:           0,
	PostTypeEnum:     0,
	PostCategoryEnum: 0,
	CreatedTime:      0,
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
	UserId         int
	UserName       string
	UserNameLower  string
	FirstName      string
	LastName       string
	IsVerified     int
	AvatarId       int
	ProfilePrivacy int
	OnlinePrivacy  int
	Phone          int
	Email          string
	About          string
	PasswordHash   string
	PasswordSalt   string
	PostSeq        int
	FollowersCount int
	FollowingCount int
	PostsCount     int
	MediaCount     int
	PhotoCount     int
	VideoCount     int
	GifCount       int
	AudioCount     int
	VoiceCount     int
	FileCount      int
	LinkCount      int
	BoardCount     int
	PinedCount     int
	LikesCount     int
	ResharedCount  int
	LastPostTime   int
	CreatedTime    int
	VersionTime    int
	IsDeleted      int
	IsBanned       int
}

//ToPB
func (m *PB_User) ToFlat() *PB_User_Flat {
	r := &PB_User_Flat{
		UserId:         int(m.UserId),
		UserName:       m.UserName,
		UserNameLower:  m.UserNameLower,
		FirstName:      m.FirstName,
		LastName:       m.LastName,
		IsVerified:     int(m.IsVerified),
		AvatarId:       int(m.AvatarId),
		ProfilePrivacy: int(m.ProfilePrivacy),
		OnlinePrivacy:  int(m.OnlinePrivacy),
		Phone:          int(m.Phone),
		Email:          m.Email,
		About:          m.About,
		PasswordHash:   m.PasswordHash,
		PasswordSalt:   m.PasswordSalt,
		PostSeq:        int(m.PostSeq),
		FollowersCount: int(m.FollowersCount),
		FollowingCount: int(m.FollowingCount),
		PostsCount:     int(m.PostsCount),
		MediaCount:     int(m.MediaCount),
		PhotoCount:     int(m.PhotoCount),
		VideoCount:     int(m.VideoCount),
		GifCount:       int(m.GifCount),
		AudioCount:     int(m.AudioCount),
		VoiceCount:     int(m.VoiceCount),
		FileCount:      int(m.FileCount),
		LinkCount:      int(m.LinkCount),
		BoardCount:     int(m.BoardCount),
		PinedCount:     int(m.PinedCount),
		LikesCount:     int(m.LikesCount),
		ResharedCount:  int(m.ResharedCount),
		LastPostTime:   int(m.LastPostTime),
		CreatedTime:    int(m.CreatedTime),
		VersionTime:    int(m.VersionTime),
		IsDeleted:      int(m.IsDeleted),
		IsBanned:       int(m.IsBanned),
	}
	return r
}

//ToPB
func (m *PB_User_Flat) ToPB() *PB_User {
	r := &PB_User{
		UserId:         int32(m.UserId),
		UserName:       m.UserName,
		UserNameLower:  m.UserNameLower,
		FirstName:      m.FirstName,
		LastName:       m.LastName,
		IsVerified:     int32(m.IsVerified),
		AvatarId:       int64(m.AvatarId),
		ProfilePrivacy: int32(m.ProfilePrivacy),
		OnlinePrivacy:  int32(m.OnlinePrivacy),
		Phone:          int64(m.Phone),
		Email:          m.Email,
		About:          m.About,
		PasswordHash:   m.PasswordHash,
		PasswordSalt:   m.PasswordSalt,
		PostSeq:        int32(m.PostSeq),
		FollowersCount: int32(m.FollowersCount),
		FollowingCount: int32(m.FollowingCount),
		PostsCount:     int32(m.PostsCount),
		MediaCount:     int32(m.MediaCount),
		PhotoCount:     int32(m.PhotoCount),
		VideoCount:     int32(m.VideoCount),
		GifCount:       int32(m.GifCount),
		AudioCount:     int32(m.AudioCount),
		VoiceCount:     int32(m.VoiceCount),
		FileCount:      int32(m.FileCount),
		LinkCount:      int32(m.LinkCount),
		BoardCount:     int32(m.BoardCount),
		PinedCount:     int32(m.PinedCount),
		LikesCount:     int32(m.LikesCount),
		ResharedCount:  int32(m.ResharedCount),
		LastPostTime:   int32(m.LastPostTime),
		CreatedTime:    int32(m.CreatedTime),
		VersionTime:    int32(m.VersionTime),
		IsDeleted:      int32(m.IsDeleted),
		IsBanned:       int32(m.IsBanned),
	}
	return r
}

//folding
var PB_User__FOlD = &PB_User{
	UserId:         0,
	UserName:       "",
	UserNameLower:  "",
	FirstName:      "",
	LastName:       "",
	IsVerified:     0,
	AvatarId:       0,
	ProfilePrivacy: 0,
	OnlinePrivacy:  0,
	Phone:          0,
	Email:          "",
	About:          "",
	PasswordHash:   "",
	PasswordSalt:   "",
	PostSeq:        0,
	FollowersCount: 0,
	FollowingCount: 0,
	PostsCount:     0,
	MediaCount:     0,
	PhotoCount:     0,
	VideoCount:     0,
	GifCount:       0,
	AudioCount:     0,
	VoiceCount:     0,
	FileCount:      0,
	LinkCount:      0,
	BoardCount:     0,
	PinedCount:     0,
	LikesCount:     0,
	ResharedCount:  0,
	LastPostTime:   0,
	CreatedTime:    0,
	VersionTime:    0,
	IsDeleted:      0,
	IsBanned:       0,
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
	ChatId           int
	ChatKey          string
	RoomKey          string
	RoomType         int
	UserId           int
	PeerUserId       int
	GroupId          int
	HashTagId        int
	StartedByMe      int
	Title            string
	PinTime          int
	FromMsgId        int
	Seq              int
	UnseenCount      int
	LastMsgId        int
	LastMsgStatus    int
	SeenSeq          int
	SeenMsgId        int
	LastMsgIdRecived int
	Left             int
	Creator          int
	Kicked           int
	Admin            int
	Deactivated      int
	VersionTime      int
	OrderTime        int
	CreatedTime      int
	DraftText        string
	DratReplyToMsgId int
}

//ToPB
func (m *PB_Chat) ToFlat() *PB_Chat_Flat {
	r := &PB_Chat_Flat{
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
		UnseenCount:      int(m.UnseenCount),
		LastMsgId:        int(m.LastMsgId),
		LastMsgStatus:    int(m.LastMsgStatus),
		SeenSeq:          int(m.SeenSeq),
		SeenMsgId:        int(m.SeenMsgId),
		LastMsgIdRecived: int(m.LastMsgIdRecived),
		Left:             int(m.Left),
		Creator:          int(m.Creator),
		Kicked:           int(m.Kicked),
		Admin:            int(m.Admin),
		Deactivated:      int(m.Deactivated),
		VersionTime:      int(m.VersionTime),
		OrderTime:        int(m.OrderTime),
		CreatedTime:      int(m.CreatedTime),
		DraftText:        m.DraftText,
		DratReplyToMsgId: int(m.DratReplyToMsgId),
	}
	return r
}

//ToPB
func (m *PB_Chat_Flat) ToPB() *PB_Chat {
	r := &PB_Chat{
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
		UnseenCount:      int32(m.UnseenCount),
		LastMsgId:        int64(m.LastMsgId),
		LastMsgStatus:    int32(m.LastMsgStatus),
		SeenSeq:          int32(m.SeenSeq),
		SeenMsgId:        int64(m.SeenMsgId),
		LastMsgIdRecived: int64(m.LastMsgIdRecived),
		Left:             int32(m.Left),
		Creator:          int32(m.Creator),
		Kicked:           int32(m.Kicked),
		Admin:            int32(m.Admin),
		Deactivated:      int32(m.Deactivated),
		VersionTime:      int32(m.VersionTime),
		OrderTime:        int32(m.OrderTime),
		CreatedTime:      int32(m.CreatedTime),
		DraftText:        m.DraftText,
		DratReplyToMsgId: int64(m.DratReplyToMsgId),
	}
	return r
}

//folding
var PB_Chat__FOlD = &PB_Chat{
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
	UnseenCount:      0,
	LastMsgId:        0,
	LastMsgStatus:    0,
	SeenSeq:          0,
	SeenMsgId:        0,
	LastMsgIdRecived: 0,
	Left:             0,
	Creator:          0,
	Kicked:           0,
	Admin:            0,
	Deactivated:      0,
	VersionTime:      0,
	OrderTime:        0,
	CreatedTime:      0,
	DraftText:        "",
	DratReplyToMsgId: 0,
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

type PB_ChatVersionOrder_Flat struct {
	VersionTime int
	UserId      int
	ChatId      int
	OrderTime   int
}

//ToPB
func (m *PB_ChatVersionOrder) ToFlat() *PB_ChatVersionOrder_Flat {
	r := &PB_ChatVersionOrder_Flat{
		VersionTime: int(m.VersionTime),
		UserId:      int(m.UserId),
		ChatId:      int(m.ChatId),
		OrderTime:   int(m.OrderTime),
	}
	return r
}

//ToPB
func (m *PB_ChatVersionOrder_Flat) ToPB() *PB_ChatVersionOrder {
	r := &PB_ChatVersionOrder{
		VersionTime: int64(m.VersionTime),
		UserId:      int32(m.UserId),
		ChatId:      int32(m.ChatId),
		OrderTime:   int32(m.OrderTime),
	}
	return r
}

//folding
var PB_ChatVersionOrder__FOlD = &PB_ChatVersionOrder{
	VersionTime: 0,
	UserId:      0,
	ChatId:      0,
	OrderTime:   0,
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
	InviteLink      string
	MembersCount    int
	SortTime        int
	CreatedTime     int
	IsMute          int
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
		InviteLink:      m.InviteLink,
		MembersCount:    int(m.MembersCount),
		SortTime:        int(m.SortTime),
		CreatedTime:     int(m.CreatedTime),
		IsMute:          int(m.IsMute),
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
		InviteLink:      m.InviteLink,
		MembersCount:    int32(m.MembersCount),
		SortTime:        int32(m.SortTime),
		CreatedTime:     int32(m.CreatedTime),
		IsMute:          int32(m.IsMute),
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
	InviteLink:      "",
	MembersCount:    0,
	SortTime:        0,
	CreatedTime:     0,
	IsMute:          0,
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

type PB_SuggestedTopPost_Flat struct {
	Id     int
	PostId int
}

//ToPB
func (m *PB_SuggestedTopPost) ToFlat() *PB_SuggestedTopPost_Flat {
	r := &PB_SuggestedTopPost_Flat{
		Id:     int(m.Id),
		PostId: int(m.PostId),
	}
	return r
}

//ToPB
func (m *PB_SuggestedTopPost_Flat) ToPB() *PB_SuggestedTopPost {
	r := &PB_SuggestedTopPost{
		Id:     int64(m.Id),
		PostId: int64(m.PostId),
	}
	return r
}

//folding
var PB_SuggestedTopPost__FOlD = &PB_SuggestedTopPost{
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

func(m *PB_RPC_Chat_Types)ToFlat() *PB_RPC_Chat_Types_Flat {
r := &PB_RPC_Chat_Types_Flat{
}
return r
}

func(m *PB_Action)ToFlat() *PB_Action_Flat {
r := &PB_Action_Flat{
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

func(m *PB_Comment)ToFlat() *PB_Comment_Flat {
r := &PB_Comment_Flat{
    CommentId:  int(m.CommentId) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    Text:  m.Text ,
    LikesCount:  int(m.LikesCount) ,
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
    ActionId:  int(m.ActionId) ,
    Murmur64Hash:  int(m.Murmur64Hash) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int(m.MessageId) ,
    ReSharedId:  int(m.ReSharedId) ,
}
return r
}

func(m *PB_Like)ToFlat() *PB_Like_Flat {
r := &PB_Like_Flat{
    Id:  int(m.Id) ,
    PostId:  int(m.PostId) ,
    PostTypeEnum:  int(m.PostTypeEnum) ,
    UserId:  int(m.UserId) ,
    LikeEnum:  int(m.LikeEnum) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_Notify)ToFlat() *PB_Notify_Flat {
r := &PB_Notify_Flat{
    NotifyId:  int(m.NotifyId) ,
    ForUserId:  int(m.ForUserId) ,
    ActorUserId:  int(m.ActorUserId) ,
    NotifyTypeEnum:  int(m.NotifyTypeEnum) ,
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

func(m *PB_PhoneContact)ToFlat() *PB_PhoneContact_Flat {
r := &PB_PhoneContact_Flat{
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
    PostTypeEnum:  int(m.PostTypeEnum) ,
    PostCategoryEnum:  int(m.PostCategoryEnum) ,
    MediaId:  int(m.MediaId) ,
    PostKey:  m.PostKey ,
    Text:  m.Text ,
    RichText:  m.RichText ,
    MediaCount:  int(m.MediaCount) ,
    SharedTo:  int(m.SharedTo) ,
    DisableComment:  int(m.DisableComment) ,
    Source:  int(m.Source) ,
    HasTag:  int(m.HasTag) ,
    Seq:  int(m.Seq) ,
    CommentsCount:  int(m.CommentsCount) ,
    LikesCount:  int(m.LikesCount) ,
    ViewsCount:  int(m.ViewsCount) ,
    EditedTime:  int(m.EditedTime) ,
    CreatedTime:  int(m.CreatedTime) ,
    ReSharedPostId:  int(m.ReSharedPostId) ,
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

func(m *PB_PostKey)ToFlat() *PB_PostKey_Flat {
r := &PB_PostKey_Flat{
    Id:  int(m.Id) ,
    PostKeyStr:  m.PostKeyStr ,
    Used:  int(m.Used) ,
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

func(m *PB_PostMentioned)ToFlat() *PB_PostMentioned_Flat {
r := &PB_PostMentioned_Flat{
    MentionedId:  int(m.MentionedId) ,
    ForUserId:  int(m.ForUserId) ,
    PostId:  int(m.PostId) ,
    PostUserId:  int(m.PostUserId) ,
    PostTypeEnum:  int(m.PostTypeEnum) ,
    PostCategoryEnum:  int(m.PostCategoryEnum) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_PostReshared)ToFlat() *PB_PostReshared_Flat {
r := &PB_PostReshared_Flat{
    ResharedId:  int(m.ResharedId) ,
    ByUserId:  int(m.ByUserId) ,
    PostId:  int(m.PostId) ,
    PostUserId:  int(m.PostUserId) ,
    PostTypeEnum:  int(m.PostTypeEnum) ,
    PostCategoryEnum:  int(m.PostCategoryEnum) ,
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

func(m *PB_SettingClient)ToFlat() *PB_SettingClient_Flat {
r := &PB_SettingClient_Flat{
    UserId:  int(m.UserId) ,
    AutoDownloadWifiVoice:  int(m.AutoDownloadWifiVoice) ,
    AutoDownloadWifiImage:  int(m.AutoDownloadWifiImage) ,
    AutoDownloadWifiVideo:  int(m.AutoDownloadWifiVideo) ,
    AutoDownloadWifiFile:  int(m.AutoDownloadWifiFile) ,
    AutoDownloadWifiMusic:  int(m.AutoDownloadWifiMusic) ,
    AutoDownloadWifiGif:  int(m.AutoDownloadWifiGif) ,
    AutoDownloadCellularVoice:  int(m.AutoDownloadCellularVoice) ,
    AutoDownloadCellularImage:  int(m.AutoDownloadCellularImage) ,
    AutoDownloadCellularVideo:  int(m.AutoDownloadCellularVideo) ,
    AutoDownloadCellularFile:  int(m.AutoDownloadCellularFile) ,
    AutoDownloadCellularMusic:  int(m.AutoDownloadCellularMusic) ,
    AutoDownloadCellularGif:  int(m.AutoDownloadCellularGif) ,
    AutoDownloadRoamingVoice:  int(m.AutoDownloadRoamingVoice) ,
    AutoDownloadRoamingImage:  int(m.AutoDownloadRoamingImage) ,
    AutoDownloadRoamingVideo:  int(m.AutoDownloadRoamingVideo) ,
    AutoDownloadRoamingFile:  int(m.AutoDownloadRoamingFile) ,
    AutoDownloadRoamingMusic:  int(m.AutoDownloadRoamingMusic) ,
    AutoDownloadRoamingGif:  int(m.AutoDownloadRoamingGif) ,
    SaveToGallery:  int(m.SaveToGallery) ,
}
return r
}

func(m *PB_SettingNotification)ToFlat() *PB_SettingNotification_Flat {
r := &PB_SettingNotification_Flat{
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

func(m *PB_Tag)ToFlat() *PB_Tag_Flat {
r := &PB_Tag_Flat{
    TagId:  int(m.TagId) ,
    Name:  m.Name ,
    Count:  int(m.Count) ,
    TagStatusEnum:  int(m.TagStatusEnum) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_TagPost)ToFlat() *PB_TagPost_Flat {
r := &PB_TagPost_Flat{
    Id:  int(m.Id) ,
    TagId:  int(m.TagId) ,
    PostId:  int(m.PostId) ,
    PostTypeEnum:  int(m.PostTypeEnum) ,
    PostCategoryEnum:  int(m.PostCategoryEnum) ,
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
    ProfilePrivacy:  int(m.ProfilePrivacy) ,
    OnlinePrivacy:  int(m.OnlinePrivacy) ,
    Phone:  int(m.Phone) ,
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
    StartedByMe:  int(m.StartedByMe) ,
    Title:  m.Title ,
    PinTime:  int(m.PinTime) ,
    FromMsgId:  int(m.FromMsgId) ,
    Seq:  int(m.Seq) ,
    UnseenCount:  int(m.UnseenCount) ,
    LastMsgId:  int(m.LastMsgId) ,
    LastMsgStatus:  int(m.LastMsgStatus) ,
    SeenSeq:  int(m.SeenSeq) ,
    SeenMsgId:  int(m.SeenMsgId) ,
    LastMsgIdRecived:  int(m.LastMsgIdRecived) ,
    Left:  int(m.Left) ,
    Creator:  int(m.Creator) ,
    Kicked:  int(m.Kicked) ,
    Admin:  int(m.Admin) ,
    Deactivated:  int(m.Deactivated) ,
    VersionTime:  int(m.VersionTime) ,
    OrderTime:  int(m.OrderTime) ,
    CreatedTime:  int(m.CreatedTime) ,
    DraftText:  m.DraftText ,
    DratReplyToMsgId:  int(m.DratReplyToMsgId) ,
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

func(m *PB_ChatVersionOrder)ToFlat() *PB_ChatVersionOrder_Flat {
r := &PB_ChatVersionOrder_Flat{
    VersionTime:  int(m.VersionTime) ,
    UserId:  int(m.UserId) ,
    ChatId:  int(m.ChatId) ,
    OrderTime:  int(m.OrderTime) ,
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
    InviteLink:  m.InviteLink ,
    MembersCount:  int(m.MembersCount) ,
    SortTime:  int(m.SortTime) ,
    CreatedTime:  int(m.CreatedTime) ,
    IsMute:  int(m.IsMute) ,
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

func(m *PB_SuggestedTopPost)ToFlat() *PB_SuggestedTopPost_Flat {
r := &PB_SuggestedTopPost_Flat{
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

func(m *PB_RPC_Chat_Types_Flat)ToPB() *PB_RPC_Chat_Types {
r := &PB_RPC_Chat_Types{
}
return r
}

func(m *PB_Action_Flat)ToPB() *PB_Action {
r := &PB_Action{
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

func(m *PB_Comment_Flat)ToPB() *PB_Comment {
r := &PB_Comment{
    CommentId:  int64(m.CommentId) ,
    UserId:  int32(m.UserId) ,
    PostId:  int64(m.PostId) ,
    Text:  m.Text ,
    LikesCount:  int32(m.LikesCount) ,
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
    ActionId:  int64(m.ActionId) ,
    Murmur64Hash:  int64(m.Murmur64Hash) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int64(m.MessageId) ,
    ReSharedId:  int64(m.ReSharedId) ,
}
return r
}

func(m *PB_Like_Flat)ToPB() *PB_Like {
r := &PB_Like{
    Id:  int64(m.Id) ,
    PostId:  int64(m.PostId) ,
    PostTypeEnum:  int32(m.PostTypeEnum) ,
    UserId:  int32(m.UserId) ,
    LikeEnum:  int32(m.LikeEnum) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_Notify_Flat)ToPB() *PB_Notify {
r := &PB_Notify{
    NotifyId:  int64(m.NotifyId) ,
    ForUserId:  int32(m.ForUserId) ,
    ActorUserId:  int32(m.ActorUserId) ,
    NotifyTypeEnum:  int32(m.NotifyTypeEnum) ,
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

func(m *PB_PhoneContact_Flat)ToPB() *PB_PhoneContact {
r := &PB_PhoneContact{
    Id:  int32(m.Id) ,
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
    PostTypeEnum:  int32(m.PostTypeEnum) ,
    PostCategoryEnum:  int32(m.PostCategoryEnum) ,
    MediaId:  int64(m.MediaId) ,
    PostKey:  m.PostKey ,
    Text:  m.Text ,
    RichText:  m.RichText ,
    MediaCount:  int32(m.MediaCount) ,
    SharedTo:  int32(m.SharedTo) ,
    DisableComment:  int32(m.DisableComment) ,
    Source:  int32(m.Source) ,
    HasTag:  int32(m.HasTag) ,
    Seq:  int32(m.Seq) ,
    CommentsCount:  int32(m.CommentsCount) ,
    LikesCount:  int32(m.LikesCount) ,
    ViewsCount:  int32(m.ViewsCount) ,
    EditedTime:  int32(m.EditedTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
    ReSharedPostId:  int64(m.ReSharedPostId) ,
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

func(m *PB_PostKey_Flat)ToPB() *PB_PostKey {
r := &PB_PostKey{
    Id:  int32(m.Id) ,
    PostKeyStr:  m.PostKeyStr ,
    Used:  int32(m.Used) ,
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

func(m *PB_PostMentioned_Flat)ToPB() *PB_PostMentioned {
r := &PB_PostMentioned{
    MentionedId:  int64(m.MentionedId) ,
    ForUserId:  int32(m.ForUserId) ,
    PostId:  int64(m.PostId) ,
    PostUserId:  int32(m.PostUserId) ,
    PostTypeEnum:  int32(m.PostTypeEnum) ,
    PostCategoryEnum:  int32(m.PostCategoryEnum) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_PostReshared_Flat)ToPB() *PB_PostReshared {
r := &PB_PostReshared{
    ResharedId:  int64(m.ResharedId) ,
    ByUserId:  int32(m.ByUserId) ,
    PostId:  int64(m.PostId) ,
    PostUserId:  int32(m.PostUserId) ,
    PostTypeEnum:  int32(m.PostTypeEnum) ,
    PostCategoryEnum:  int32(m.PostCategoryEnum) ,
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

func(m *PB_SettingClient_Flat)ToPB() *PB_SettingClient {
r := &PB_SettingClient{
    UserId:  int32(m.UserId) ,
    AutoDownloadWifiVoice:  int32(m.AutoDownloadWifiVoice) ,
    AutoDownloadWifiImage:  int32(m.AutoDownloadWifiImage) ,
    AutoDownloadWifiVideo:  int32(m.AutoDownloadWifiVideo) ,
    AutoDownloadWifiFile:  int32(m.AutoDownloadWifiFile) ,
    AutoDownloadWifiMusic:  int32(m.AutoDownloadWifiMusic) ,
    AutoDownloadWifiGif:  int32(m.AutoDownloadWifiGif) ,
    AutoDownloadCellularVoice:  int32(m.AutoDownloadCellularVoice) ,
    AutoDownloadCellularImage:  int32(m.AutoDownloadCellularImage) ,
    AutoDownloadCellularVideo:  int32(m.AutoDownloadCellularVideo) ,
    AutoDownloadCellularFile:  int32(m.AutoDownloadCellularFile) ,
    AutoDownloadCellularMusic:  int32(m.AutoDownloadCellularMusic) ,
    AutoDownloadCellularGif:  int32(m.AutoDownloadCellularGif) ,
    AutoDownloadRoamingVoice:  int32(m.AutoDownloadRoamingVoice) ,
    AutoDownloadRoamingImage:  int32(m.AutoDownloadRoamingImage) ,
    AutoDownloadRoamingVideo:  int32(m.AutoDownloadRoamingVideo) ,
    AutoDownloadRoamingFile:  int32(m.AutoDownloadRoamingFile) ,
    AutoDownloadRoamingMusic:  int32(m.AutoDownloadRoamingMusic) ,
    AutoDownloadRoamingGif:  int32(m.AutoDownloadRoamingGif) ,
    SaveToGallery:  int32(m.SaveToGallery) ,
}
return r
}

func(m *PB_SettingNotification_Flat)ToPB() *PB_SettingNotification {
r := &PB_SettingNotification{
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

func(m *PB_Tag_Flat)ToPB() *PB_Tag {
r := &PB_Tag{
    TagId:  int64(m.TagId) ,
    Name:  m.Name ,
    Count:  int32(m.Count) ,
    TagStatusEnum:  int32(m.TagStatusEnum) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_TagPost_Flat)ToPB() *PB_TagPost {
r := &PB_TagPost{
    Id:  int64(m.Id) ,
    TagId:  int32(m.TagId) ,
    PostId:  int32(m.PostId) ,
    PostTypeEnum:  int32(m.PostTypeEnum) ,
    PostCategoryEnum:  int32(m.PostCategoryEnum) ,
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
    ProfilePrivacy:  int32(m.ProfilePrivacy) ,
    OnlinePrivacy:  int32(m.OnlinePrivacy) ,
    Phone:  int64(m.Phone) ,
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
    StartedByMe:  int32(m.StartedByMe) ,
    Title:  m.Title ,
    PinTime:  int64(m.PinTime) ,
    FromMsgId:  int64(m.FromMsgId) ,
    Seq:  int32(m.Seq) ,
    UnseenCount:  int32(m.UnseenCount) ,
    LastMsgId:  int64(m.LastMsgId) ,
    LastMsgStatus:  int32(m.LastMsgStatus) ,
    SeenSeq:  int32(m.SeenSeq) ,
    SeenMsgId:  int64(m.SeenMsgId) ,
    LastMsgIdRecived:  int64(m.LastMsgIdRecived) ,
    Left:  int32(m.Left) ,
    Creator:  int32(m.Creator) ,
    Kicked:  int32(m.Kicked) ,
    Admin:  int32(m.Admin) ,
    Deactivated:  int32(m.Deactivated) ,
    VersionTime:  int32(m.VersionTime) ,
    OrderTime:  int32(m.OrderTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
    DraftText:  m.DraftText ,
    DratReplyToMsgId:  int64(m.DratReplyToMsgId) ,
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

func(m *PB_ChatVersionOrder_Flat)ToPB() *PB_ChatVersionOrder {
r := &PB_ChatVersionOrder{
    VersionTime:  int64(m.VersionTime) ,
    UserId:  int32(m.UserId) ,
    ChatId:  int32(m.ChatId) ,
    OrderTime:  int32(m.OrderTime) ,
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
    InviteLink:  m.InviteLink ,
    MembersCount:  int32(m.MembersCount) ,
    SortTime:  int32(m.SortTime) ,
    CreatedTime:  int32(m.CreatedTime) ,
    IsMute:  int32(m.IsMute) ,
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

func(m *PB_SuggestedTopPost_Flat)ToPB() *PB_SuggestedTopPost {
r := &PB_SuggestedTopPost{
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


var PB_RPC_Chat_Types__FOlD = &PB_RPC_Chat_Types{
}


var PB_Action__FOlD = &PB_Action{
        ActionId:  0 ,
        ActorUserId:  0 ,
        ActionTypeEnum:  0 ,
        PeerUserId:  0 ,
        PostId:  0 ,
        CommentId:  0 ,
        Murmur64Hash:  0 ,
        CreatedTime:  0 ,
}


var PB_Comment__FOlD = &PB_Comment{
        CommentId:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        Text:  "" ,
        LikesCount:  0 ,
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
        ActionId:  0 ,
        Murmur64Hash:  0 ,
        ChatKey:  "" ,
        MessageId:  0 ,
        ReSharedId:  0 ,
}


var PB_Like__FOlD = &PB_Like{
        Id:  0 ,
        PostId:  0 ,
        PostTypeEnum:  0 ,
        UserId:  0 ,
        LikeEnum:  0 ,
        CreatedTime:  0 ,
}


var PB_Notify__FOlD = &PB_Notify{
        NotifyId:  0 ,
        ForUserId:  0 ,
        ActorUserId:  0 ,
        NotifyTypeEnum:  0 ,
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


var PB_PhoneContact__FOlD = &PB_PhoneContact{
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
        PostTypeEnum:  0 ,
        PostCategoryEnum:  0 ,
        MediaId:  0 ,
        PostKey:  "" ,
        Text:  "" ,
        RichText:  "" ,
        MediaCount:  0 ,
        SharedTo:  0 ,
        DisableComment:  0 ,
        Source:  0 ,
        HasTag:  0 ,
        Seq:  0 ,
        CommentsCount:  0 ,
        LikesCount:  0 ,
        ViewsCount:  0 ,
        EditedTime:  0 ,
        CreatedTime:  0 ,
        ReSharedPostId:  0 ,
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


var PB_PostKey__FOlD = &PB_PostKey{
        Id:  0 ,
        PostKeyStr:  "" ,
        Used:  0 ,
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


var PB_PostMentioned__FOlD = &PB_PostMentioned{
        MentionedId:  0 ,
        ForUserId:  0 ,
        PostId:  0 ,
        PostUserId:  0 ,
        PostTypeEnum:  0 ,
        PostCategoryEnum:  0 ,
        CreatedTime:  0 ,
}


var PB_PostReshared__FOlD = &PB_PostReshared{
        ResharedId:  0 ,
        ByUserId:  0 ,
        PostId:  0 ,
        PostUserId:  0 ,
        PostTypeEnum:  0 ,
        PostCategoryEnum:  0 ,
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


var PB_SettingClient__FOlD = &PB_SettingClient{
        UserId:  0 ,
        AutoDownloadWifiVoice:  0 ,
        AutoDownloadWifiImage:  0 ,
        AutoDownloadWifiVideo:  0 ,
        AutoDownloadWifiFile:  0 ,
        AutoDownloadWifiMusic:  0 ,
        AutoDownloadWifiGif:  0 ,
        AutoDownloadCellularVoice:  0 ,
        AutoDownloadCellularImage:  0 ,
        AutoDownloadCellularVideo:  0 ,
        AutoDownloadCellularFile:  0 ,
        AutoDownloadCellularMusic:  0 ,
        AutoDownloadCellularGif:  0 ,
        AutoDownloadRoamingVoice:  0 ,
        AutoDownloadRoamingImage:  0 ,
        AutoDownloadRoamingVideo:  0 ,
        AutoDownloadRoamingFile:  0 ,
        AutoDownloadRoamingMusic:  0 ,
        AutoDownloadRoamingGif:  0 ,
        SaveToGallery:  0 ,
}


var PB_SettingNotification__FOlD = &PB_SettingNotification{
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


var PB_Tag__FOlD = &PB_Tag{
        TagId:  0 ,
        Name:  "" ,
        Count:  0 ,
        TagStatusEnum:  0 ,
        CreatedTime:  0 ,
}


var PB_TagPost__FOlD = &PB_TagPost{
        Id:  0 ,
        TagId:  0 ,
        PostId:  0 ,
        PostTypeEnum:  0 ,
        PostCategoryEnum:  0 ,
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
        ProfilePrivacy:  0 ,
        OnlinePrivacy:  0 ,
        Phone:  0 ,
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
        StartedByMe:  0 ,
        Title:  "" ,
        PinTime:  0 ,
        FromMsgId:  0 ,
        Seq:  0 ,
        UnseenCount:  0 ,
        LastMsgId:  0 ,
        LastMsgStatus:  0 ,
        SeenSeq:  0 ,
        SeenMsgId:  0 ,
        LastMsgIdRecived:  0 ,
        Left:  0 ,
        Creator:  0 ,
        Kicked:  0 ,
        Admin:  0 ,
        Deactivated:  0 ,
        VersionTime:  0 ,
        OrderTime:  0 ,
        CreatedTime:  0 ,
        DraftText:  "" ,
        DratReplyToMsgId:  0 ,
}


var PB_ChatLastMessage__FOlD = &PB_ChatLastMessage{
        ChatIdGroupId:  "" ,
        LastMsgPb:  []byte{} ,
}


var PB_ChatVersionOrder__FOlD = &PB_ChatVersionOrder{
        VersionTime:  0 ,
        UserId:  0 ,
        ChatId:  0 ,
        OrderTime:  0 ,
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
        InviteLink:  "" ,
        MembersCount:  0 ,
        SortTime:  0 ,
        CreatedTime:  0 ,
        IsMute:  0 ,
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


var PB_SuggestedTopPost__FOlD = &PB_SuggestedTopPost{
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



*/
