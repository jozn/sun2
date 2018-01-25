package x

import "ms/sun/helper"

type GeoLocation_Flat struct {
	Lat float64
	Lon float64
}

//ToPB
func (m *GeoLocation) ToFlat() *GeoLocation_Flat {
	r := &GeoLocation_Flat{
		Lat: float64(m.Lat),
		Lon: float64(m.Lon),
	}
	return r
}

//ToPB
func (m *GeoLocation_Flat) ToPB() *GeoLocation {
	r := &GeoLocation{
		Lat: m.Lat,
		Lon: m.Lon,
	}
	return r
}

//folding
var GeoLocation__FOlD = &GeoLocation{
	Lat: 0.0,
	Lon: 0.0,
}

type RoomMessageLog_Flat struct {
	typ          RoomMessageLogEnum
	TargetUserId int
	ByUserId     int
}

//ToPB
func (m *RoomMessageLog) ToFlat() *RoomMessageLog_Flat {
	r := &RoomMessageLog_Flat{

		TargetUserId: int(m.TargetUserId),
		ByUserId:     int(m.ByUserId),
	}
	return r
}

//ToPB
func (m *RoomMessageLog_Flat) ToPB() *RoomMessageLog {
	r := &RoomMessageLog{

		TargetUserId: uint64(m.TargetUserId),
		ByUserId:     uint64(m.ByUserId),
	}
	return r
}

//folding
var RoomMessageLog__FOlD = &RoomMessageLog{

	TargetUserId: 0,
	ByUserId:     0,
}

type RoomMessageForwardFrom_Flat struct {
	RoomId       int
	MessageId    int
	RoomTypeEnum int
}

//ToPB
func (m *RoomMessageForwardFrom) ToFlat() *RoomMessageForwardFrom_Flat {
	r := &RoomMessageForwardFrom_Flat{
		RoomId:       int(m.RoomId),
		MessageId:    int(m.MessageId),
		RoomTypeEnum: int(m.RoomTypeEnum),
	}
	return r
}

//ToPB
func (m *RoomMessageForwardFrom_Flat) ToPB() *RoomMessageForwardFrom {
	r := &RoomMessageForwardFrom{
		RoomId:       uint64(m.RoomId),
		MessageId:    uint64(m.MessageId),
		RoomTypeEnum: uint32(m.RoomTypeEnum),
	}
	return r
}

//folding
var RoomMessageForwardFrom__FOlD = &RoomMessageForwardFrom{
	RoomId:       0,
	MessageId:    0,
	RoomTypeEnum: 0,
}

type RoomDraft_Flat struct {
	Message string
	ReplyTo int
}

//ToPB
func (m *RoomDraft) ToFlat() *RoomDraft_Flat {
	r := &RoomDraft_Flat{
		Message: m.Message,
		ReplyTo: int(m.ReplyTo),
	}
	return r
}

//ToPB
func (m *RoomDraft_Flat) ToPB() *RoomDraft {
	r := &RoomDraft{
		Message: m.Message,
		ReplyTo: uint64(m.ReplyTo),
	}
	return r
}

//folding
var RoomDraft__FOlD = &RoomDraft{
	Message: "",
	ReplyTo: 0,
}

type ChatRoom_Flat struct {
}

//ToPB
func (m *ChatRoom) ToFlat() *ChatRoom_Flat {
	r := &ChatRoom_Flat{}
	return r
}

//ToPB
func (m *ChatRoom_Flat) ToPB() *ChatRoom {
	r := &ChatRoom{}
	return r
}

//folding
var ChatRoom__FOlD = &ChatRoom{}

type Pagination_Flat struct {
	Offset int
	Limit  int
}

//ToPB
func (m *Pagination) ToFlat() *Pagination_Flat {
	r := &Pagination_Flat{
		Offset: int(m.Offset),
		Limit:  int(m.Limit),
	}
	return r
}

//ToPB
func (m *Pagination_Flat) ToPB() *Pagination {
	r := &Pagination{
		Offset: uint32(m.Offset),
		Limit:  uint32(m.Limit),
	}
	return r
}

//folding
var Pagination__FOlD = &Pagination{
	Offset: 0,
	Limit:  0,
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

type PB_ResponseExtra_Flat struct {
	ErrorCode   int
	ErrMessage  string
	RpcFullName string
	Data        []byte
}

//ToPB
func (m *PB_ResponseExtra) ToFlat() *PB_ResponseExtra_Flat {
	r := &PB_ResponseExtra_Flat{
		ErrorCode:   int(m.ErrorCode),
		ErrMessage:  m.ErrMessage,
		RpcFullName: m.RpcFullName,
		Data:        []byte(m.Data),
	}
	return r
}

//ToPB
func (m *PB_ResponseExtra_Flat) ToPB() *PB_ResponseExtra {
	r := &PB_ResponseExtra{
		ErrorCode:   int64(m.ErrorCode),
		ErrMessage:  m.ErrMessage,
		RpcFullName: m.RpcFullName,
		Data:        m.Data,
	}
	return r
}

//folding
var PB_ResponseExtra__FOlD = &PB_ResponseExtra{
	ErrorCode:   0,
	ErrMessage:  "",
	RpcFullName: "",
	Data:        []byte{},
}

type PB_Pager_Flat struct {
	Page  int
	Limit int
	Last  int
}

//ToPB
func (m *PB_Pager) ToFlat() *PB_Pager_Flat {
	r := &PB_Pager_Flat{
		Page:  int(m.Page),
		Limit: int(m.Limit),
		Last:  int(m.Last),
	}
	return r
}

//ToPB
func (m *PB_Pager_Flat) ToPB() *PB_Pager {
	r := &PB_Pager{
		Page:  int64(m.Page),
		Limit: int64(m.Limit),
		Last:  int64(m.Last),
	}
	return r
}

//folding
var PB_Pager__FOlD = &PB_Pager{
	Page:  0,
	Limit: 0,
	Last:  0,
}

type PB_UserParam_CheckUserName2_Flat struct {
}

//ToPB
func (m *PB_UserParam_CheckUserName2) ToFlat() *PB_UserParam_CheckUserName2_Flat {
	r := &PB_UserParam_CheckUserName2_Flat{}
	return r
}

//ToPB
func (m *PB_UserParam_CheckUserName2_Flat) ToPB() *PB_UserParam_CheckUserName2 {
	r := &PB_UserParam_CheckUserName2{}
	return r
}

//folding
var PB_UserParam_CheckUserName2__FOlD = &PB_UserParam_CheckUserName2{}

type PB_UserResponse_CheckUserName2_Flat struct {
}

//ToPB
func (m *PB_UserResponse_CheckUserName2) ToFlat() *PB_UserResponse_CheckUserName2_Flat {
	r := &PB_UserResponse_CheckUserName2_Flat{}
	return r
}

//ToPB
func (m *PB_UserResponse_CheckUserName2_Flat) ToPB() *PB_UserResponse_CheckUserName2 {
	r := &PB_UserResponse_CheckUserName2{}
	return r
}

//folding
var PB_UserResponse_CheckUserName2__FOlD = &PB_UserResponse_CheckUserName2{}

type PB_ChatParam_AddNewMessage_Flat struct {
	ReplyToMessageId int
	Blob             []byte
	MessageView      PB_MessageView
}

//ToPB
func (m *PB_ChatParam_AddNewMessage) ToFlat() *PB_ChatParam_AddNewMessage_Flat {
	r := &PB_ChatParam_AddNewMessage_Flat{
		ReplyToMessageId: int(m.ReplyToMessageId),
		Blob:             []byte(m.Blob),
	}
	return r
}

//ToPB
func (m *PB_ChatParam_AddNewMessage_Flat) ToPB() *PB_ChatParam_AddNewMessage {
	r := &PB_ChatParam_AddNewMessage{
		ReplyToMessageId: int64(m.ReplyToMessageId),
		Blob:             m.Blob,
	}
	return r
}

//folding
var PB_ChatParam_AddNewMessage__FOlD = &PB_ChatParam_AddNewMessage{
	ReplyToMessageId: 0,
	Blob:             []byte{},
}

type PB_ChatResponse_AddNewMessage_Flat struct {
	MessageKey       string
	NewMessageId     int
	MessageFileKey   string
	NewMessageFileId int
	AtTime           int
}

//ToPB
func (m *PB_ChatResponse_AddNewMessage) ToFlat() *PB_ChatResponse_AddNewMessage_Flat {
	r := &PB_ChatResponse_AddNewMessage_Flat{
		MessageKey:       m.MessageKey,
		NewMessageId:     int(m.NewMessageId),
		MessageFileKey:   m.MessageFileKey,
		NewMessageFileId: int(m.NewMessageFileId),
		AtTime:           int(m.AtTime),
	}
	return r
}

//ToPB
func (m *PB_ChatResponse_AddNewMessage_Flat) ToPB() *PB_ChatResponse_AddNewMessage {
	r := &PB_ChatResponse_AddNewMessage{
		MessageKey:       m.MessageKey,
		NewMessageId:     int64(m.NewMessageId),
		MessageFileKey:   m.MessageFileKey,
		NewMessageFileId: int64(m.NewMessageFileId),
		AtTime:           int64(m.AtTime),
	}
	return r
}

//folding
var PB_ChatResponse_AddNewMessage__FOlD = &PB_ChatResponse_AddNewMessage{
	MessageKey:       "",
	NewMessageId:     0,
	MessageFileKey:   "",
	NewMessageFileId: 0,
	AtTime:           0,
}

type PB_ChatParam_SetRoomActionDoing_Flat struct {
	GroupId       int
	DirectRoomKey string
	ActionType    RoomActionDoingEnum
}

//ToPB
func (m *PB_ChatParam_SetRoomActionDoing) ToFlat() *PB_ChatParam_SetRoomActionDoing_Flat {
	r := &PB_ChatParam_SetRoomActionDoing_Flat{
		GroupId:       int(m.GroupId),
		DirectRoomKey: m.DirectRoomKey,
	}
	return r
}

//ToPB
func (m *PB_ChatParam_SetRoomActionDoing_Flat) ToPB() *PB_ChatParam_SetRoomActionDoing {
	r := &PB_ChatParam_SetRoomActionDoing{
		GroupId:       int64(m.GroupId),
		DirectRoomKey: m.DirectRoomKey,
	}
	return r
}

//folding
var PB_ChatParam_SetRoomActionDoing__FOlD = &PB_ChatParam_SetRoomActionDoing{
	GroupId:       0,
	DirectRoomKey: "",
}

type PB_ChatResponse_SetRoomActionDoing_Flat struct {
}

//ToPB
func (m *PB_ChatResponse_SetRoomActionDoing) ToFlat() *PB_ChatResponse_SetRoomActionDoing_Flat {
	r := &PB_ChatResponse_SetRoomActionDoing_Flat{}
	return r
}

//ToPB
func (m *PB_ChatResponse_SetRoomActionDoing_Flat) ToPB() *PB_ChatResponse_SetRoomActionDoing {
	r := &PB_ChatResponse_SetRoomActionDoing{}
	return r
}

//folding
var PB_ChatResponse_SetRoomActionDoing__FOlD = &PB_ChatResponse_SetRoomActionDoing{}

type PB_ChatParam_SetChatMessagesRangeAsSeen_Flat struct {
	ChatKey         string
	BottomMessageId int
	TopMessageId    int
	SeenTimeMs      int
}

//ToPB
func (m *PB_ChatParam_SetChatMessagesRangeAsSeen) ToFlat() *PB_ChatParam_SetChatMessagesRangeAsSeen_Flat {
	r := &PB_ChatParam_SetChatMessagesRangeAsSeen_Flat{
		ChatKey:         m.ChatKey,
		BottomMessageId: int(m.BottomMessageId),
		TopMessageId:    int(m.TopMessageId),
		SeenTimeMs:      int(m.SeenTimeMs),
	}
	return r
}

//ToPB
func (m *PB_ChatParam_SetChatMessagesRangeAsSeen_Flat) ToPB() *PB_ChatParam_SetChatMessagesRangeAsSeen {
	r := &PB_ChatParam_SetChatMessagesRangeAsSeen{
		ChatKey:         m.ChatKey,
		BottomMessageId: int64(m.BottomMessageId),
		TopMessageId:    int64(m.TopMessageId),
		SeenTimeMs:      int64(m.SeenTimeMs),
	}
	return r
}

//folding
var PB_ChatParam_SetChatMessagesRangeAsSeen__FOlD = &PB_ChatParam_SetChatMessagesRangeAsSeen{
	ChatKey:         "",
	BottomMessageId: 0,
	TopMessageId:    0,
	SeenTimeMs:      0,
}

type PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat struct {
}

//ToPB
func (m *PB_ChatResponse_SetChatMessagesRangeAsSeen) ToFlat() *PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat {
	r := &PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat{}
	return r
}

//ToPB
func (m *PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat) ToPB() *PB_ChatResponse_SetChatMessagesRangeAsSeen {
	r := &PB_ChatResponse_SetChatMessagesRangeAsSeen{}
	return r
}

//folding
var PB_ChatResponse_SetChatMessagesRangeAsSeen__FOlD = &PB_ChatResponse_SetChatMessagesRangeAsSeen{}

type PB_ChatParam_DeleteChatHistory_Flat struct {
	ChatKey       string
	FromMessageId int
}

//ToPB
func (m *PB_ChatParam_DeleteChatHistory) ToFlat() *PB_ChatParam_DeleteChatHistory_Flat {
	r := &PB_ChatParam_DeleteChatHistory_Flat{
		ChatKey:       m.ChatKey,
		FromMessageId: int(m.FromMessageId),
	}
	return r
}

//ToPB
func (m *PB_ChatParam_DeleteChatHistory_Flat) ToPB() *PB_ChatParam_DeleteChatHistory {
	r := &PB_ChatParam_DeleteChatHistory{
		ChatKey:       m.ChatKey,
		FromMessageId: int64(m.FromMessageId),
	}
	return r
}

//folding
var PB_ChatParam_DeleteChatHistory__FOlD = &PB_ChatParam_DeleteChatHistory{
	ChatKey:       "",
	FromMessageId: 0,
}

type PB_ChatResponse_DeleteChatHistory_Flat struct {
}

//ToPB
func (m *PB_ChatResponse_DeleteChatHistory) ToFlat() *PB_ChatResponse_DeleteChatHistory_Flat {
	r := &PB_ChatResponse_DeleteChatHistory_Flat{}
	return r
}

//ToPB
func (m *PB_ChatResponse_DeleteChatHistory_Flat) ToPB() *PB_ChatResponse_DeleteChatHistory {
	r := &PB_ChatResponse_DeleteChatHistory{}
	return r
}

//folding
var PB_ChatResponse_DeleteChatHistory__FOlD = &PB_ChatResponse_DeleteChatHistory{}

type PB_ChatParam_DeleteMessagesByIds_Flat struct {
	ChatKey     string
	Both        bool
	MessagesIds []int
}

//ToPB
func (m *PB_ChatParam_DeleteMessagesByIds) ToFlat() *PB_ChatParam_DeleteMessagesByIds_Flat {
	r := &PB_ChatParam_DeleteMessagesByIds_Flat{
		ChatKey:     m.ChatKey,
		Both:        m.Both,
		MessagesIds: helper.SliceInt64ToInt(m.MessagesIds),
	}
	return r
}

//ToPB
func (m *PB_ChatParam_DeleteMessagesByIds_Flat) ToPB() *PB_ChatParam_DeleteMessagesByIds {
	r := &PB_ChatParam_DeleteMessagesByIds{
		ChatKey:     m.ChatKey,
		Both:        m.Both,
		MessagesIds: helper.SliceIntToInt64(m.MessagesIds),
	}
	return r
}

//folding
var PB_ChatParam_DeleteMessagesByIds__FOlD = &PB_ChatParam_DeleteMessagesByIds{
	ChatKey: "",
	Both:    false,
}

type PB_ChatResponse_DeleteMessagesByIds_Flat struct {
}

//ToPB
func (m *PB_ChatResponse_DeleteMessagesByIds) ToFlat() *PB_ChatResponse_DeleteMessagesByIds_Flat {
	r := &PB_ChatResponse_DeleteMessagesByIds_Flat{}
	return r
}

//ToPB
func (m *PB_ChatResponse_DeleteMessagesByIds_Flat) ToPB() *PB_ChatResponse_DeleteMessagesByIds {
	r := &PB_ChatResponse_DeleteMessagesByIds{}
	return r
}

//folding
var PB_ChatResponse_DeleteMessagesByIds__FOlD = &PB_ChatResponse_DeleteMessagesByIds{}

type PB_ChatParam_SetMessagesAsReceived_Flat struct {
	ChatRoom   string
	MessageIds []int
}

//ToPB
func (m *PB_ChatParam_SetMessagesAsReceived) ToFlat() *PB_ChatParam_SetMessagesAsReceived_Flat {
	r := &PB_ChatParam_SetMessagesAsReceived_Flat{
		ChatRoom:   m.ChatRoom,
		MessageIds: helper.SliceInt64ToInt(m.MessageIds),
	}
	return r
}

//ToPB
func (m *PB_ChatParam_SetMessagesAsReceived_Flat) ToPB() *PB_ChatParam_SetMessagesAsReceived {
	r := &PB_ChatParam_SetMessagesAsReceived{
		ChatRoom:   m.ChatRoom,
		MessageIds: helper.SliceIntToInt64(m.MessageIds),
	}
	return r
}

//folding
var PB_ChatParam_SetMessagesAsReceived__FOlD = &PB_ChatParam_SetMessagesAsReceived{
	ChatRoom: "",
}

type PB_ChatResponse_SetMessagesAsReceived_Flat struct {
}

//ToPB
func (m *PB_ChatResponse_SetMessagesAsReceived) ToFlat() *PB_ChatResponse_SetMessagesAsReceived_Flat {
	r := &PB_ChatResponse_SetMessagesAsReceived_Flat{}
	return r
}

//ToPB
func (m *PB_ChatResponse_SetMessagesAsReceived_Flat) ToPB() *PB_ChatResponse_SetMessagesAsReceived {
	r := &PB_ChatResponse_SetMessagesAsReceived{}
	return r
}

//folding
var PB_ChatResponse_SetMessagesAsReceived__FOlD = &PB_ChatResponse_SetMessagesAsReceived{}

type PB_ChatParam_EditMessage_Flat struct {
	RoomKey   string
	MessageId int
	NewText   string
}

//ToPB
func (m *PB_ChatParam_EditMessage) ToFlat() *PB_ChatParam_EditMessage_Flat {
	r := &PB_ChatParam_EditMessage_Flat{
		RoomKey:   m.RoomKey,
		MessageId: int(m.MessageId),
		NewText:   m.NewText,
	}
	return r
}

//ToPB
func (m *PB_ChatParam_EditMessage_Flat) ToPB() *PB_ChatParam_EditMessage {
	r := &PB_ChatParam_EditMessage{
		RoomKey:   m.RoomKey,
		MessageId: int64(m.MessageId),
		NewText:   m.NewText,
	}
	return r
}

//folding
var PB_ChatParam_EditMessage__FOlD = &PB_ChatParam_EditMessage{
	RoomKey:   "",
	MessageId: 0,
	NewText:   "",
}

type PB_ChatResponse_EditMessage_Flat struct {
}

//ToPB
func (m *PB_ChatResponse_EditMessage) ToFlat() *PB_ChatResponse_EditMessage_Flat {
	r := &PB_ChatResponse_EditMessage_Flat{}
	return r
}

//ToPB
func (m *PB_ChatResponse_EditMessage_Flat) ToPB() *PB_ChatResponse_EditMessage {
	r := &PB_ChatResponse_EditMessage{}
	return r
}

//folding
var PB_ChatResponse_EditMessage__FOlD = &PB_ChatResponse_EditMessage{}

type PB_ChatParam_GetChatList_Flat struct {
}

//ToPB
func (m *PB_ChatParam_GetChatList) ToFlat() *PB_ChatParam_GetChatList_Flat {
	r := &PB_ChatParam_GetChatList_Flat{}
	return r
}

//ToPB
func (m *PB_ChatParam_GetChatList_Flat) ToPB() *PB_ChatParam_GetChatList {
	r := &PB_ChatParam_GetChatList{}
	return r
}

//folding
var PB_ChatParam_GetChatList__FOlD = &PB_ChatParam_GetChatList{}

type PB_ChatResponse_GetChatList_Flat struct {
	Chats []PB_ChatView
	Users []PB_UserView
}

//ToPB
func (m *PB_ChatResponse_GetChatList) ToFlat() *PB_ChatResponse_GetChatList_Flat {
	r := &PB_ChatResponse_GetChatList_Flat{}
	return r
}

//ToPB
func (m *PB_ChatResponse_GetChatList_Flat) ToPB() *PB_ChatResponse_GetChatList {
	r := &PB_ChatResponse_GetChatList{}
	return r
}

//folding
var PB_ChatResponse_GetChatList__FOlD = &PB_ChatResponse_GetChatList{}

type PB_ChatParam_GetChatHistoryToOlder_Flat struct {
	ChatKey          string
	Limit            int
	FromTopMessageId int
}

//ToPB
func (m *PB_ChatParam_GetChatHistoryToOlder) ToFlat() *PB_ChatParam_GetChatHistoryToOlder_Flat {
	r := &PB_ChatParam_GetChatHistoryToOlder_Flat{
		ChatKey:          m.ChatKey,
		Limit:            int(m.Limit),
		FromTopMessageId: int(m.FromTopMessageId),
	}
	return r
}

//ToPB
func (m *PB_ChatParam_GetChatHistoryToOlder_Flat) ToPB() *PB_ChatParam_GetChatHistoryToOlder {
	r := &PB_ChatParam_GetChatHistoryToOlder{
		ChatKey:          m.ChatKey,
		Limit:            int32(m.Limit),
		FromTopMessageId: int64(m.FromTopMessageId),
	}
	return r
}

//folding
var PB_ChatParam_GetChatHistoryToOlder__FOlD = &PB_ChatParam_GetChatHistoryToOlder{
	ChatKey:          "",
	Limit:            0,
	FromTopMessageId: 0,
}

type PB_ChatResponse_GetChatHistoryToOlder_Flat struct {
	MessagesViews []PB_MessageView
	HasMore       bool
}

//ToPB
func (m *PB_ChatResponse_GetChatHistoryToOlder) ToFlat() *PB_ChatResponse_GetChatHistoryToOlder_Flat {
	r := &PB_ChatResponse_GetChatHistoryToOlder_Flat{

		HasMore: m.HasMore,
	}
	return r
}

//ToPB
func (m *PB_ChatResponse_GetChatHistoryToOlder_Flat) ToPB() *PB_ChatResponse_GetChatHistoryToOlder {
	r := &PB_ChatResponse_GetChatHistoryToOlder{

		HasMore: m.HasMore,
	}
	return r
}

//folding
var PB_ChatResponse_GetChatHistoryToOlder__FOlD = &PB_ChatResponse_GetChatHistoryToOlder{

	HasMore: false,
}

type PB_ChatParam_GetFreshAllDirectMessagesList_Flat struct {
	LowMessageId  int
	HighMessageId int
}

//ToPB
func (m *PB_ChatParam_GetFreshAllDirectMessagesList) ToFlat() *PB_ChatParam_GetFreshAllDirectMessagesList_Flat {
	r := &PB_ChatParam_GetFreshAllDirectMessagesList_Flat{
		LowMessageId:  int(m.LowMessageId),
		HighMessageId: int(m.HighMessageId),
	}
	return r
}

//ToPB
func (m *PB_ChatParam_GetFreshAllDirectMessagesList_Flat) ToPB() *PB_ChatParam_GetFreshAllDirectMessagesList {
	r := &PB_ChatParam_GetFreshAllDirectMessagesList{
		LowMessageId:  int64(m.LowMessageId),
		HighMessageId: int64(m.HighMessageId),
	}
	return r
}

//folding
var PB_ChatParam_GetFreshAllDirectMessagesList__FOlD = &PB_ChatParam_GetFreshAllDirectMessagesList{
	LowMessageId:  0,
	HighMessageId: 0,
}

type PB_ChatResponse_GetFreshAllDirectMessagesList_Flat struct {
	Messages []PB_MessageView
	HasMore  bool
}

//ToPB
func (m *PB_ChatResponse_GetFreshAllDirectMessagesList) ToFlat() *PB_ChatResponse_GetFreshAllDirectMessagesList_Flat {
	r := &PB_ChatResponse_GetFreshAllDirectMessagesList_Flat{

		HasMore: m.HasMore,
	}
	return r
}

//ToPB
func (m *PB_ChatResponse_GetFreshAllDirectMessagesList_Flat) ToPB() *PB_ChatResponse_GetFreshAllDirectMessagesList {
	r := &PB_ChatResponse_GetFreshAllDirectMessagesList{

		HasMore: m.HasMore,
	}
	return r
}

//folding
var PB_ChatResponse_GetFreshAllDirectMessagesList__FOlD = &PB_ChatResponse_GetFreshAllDirectMessagesList{

	HasMore: false,
}

type PB_OtherParam_Echo_Flat struct {
	Text string
}

//ToPB
func (m *PB_OtherParam_Echo) ToFlat() *PB_OtherParam_Echo_Flat {
	r := &PB_OtherParam_Echo_Flat{
		Text: m.Text,
	}
	return r
}

//ToPB
func (m *PB_OtherParam_Echo_Flat) ToPB() *PB_OtherParam_Echo {
	r := &PB_OtherParam_Echo{
		Text: m.Text,
	}
	return r
}

//folding
var PB_OtherParam_Echo__FOlD = &PB_OtherParam_Echo{
	Text: "",
}

type PB_OtherResponse_Echo_Flat struct {
	Text string
}

//ToPB
func (m *PB_OtherResponse_Echo) ToFlat() *PB_OtherResponse_Echo_Flat {
	r := &PB_OtherResponse_Echo_Flat{
		Text: m.Text,
	}
	return r
}

//ToPB
func (m *PB_OtherResponse_Echo_Flat) ToPB() *PB_OtherResponse_Echo {
	r := &PB_OtherResponse_Echo{
		Text: m.Text,
	}
	return r
}

//folding
var PB_OtherResponse_Echo__FOlD = &PB_OtherResponse_Echo{
	Text: "",
}

type PB_PageParam_GetCommentsPage_Flat struct {
	PostId int
	Pager  PB_Pager
}

//ToPB
func (m *PB_PageParam_GetCommentsPage) ToFlat() *PB_PageParam_GetCommentsPage_Flat {
	r := &PB_PageParam_GetCommentsPage_Flat{
		PostId: int(m.PostId),
	}
	return r
}

//ToPB
func (m *PB_PageParam_GetCommentsPage_Flat) ToPB() *PB_PageParam_GetCommentsPage {
	r := &PB_PageParam_GetCommentsPage{
		PostId: int64(m.PostId),
	}
	return r
}

//folding
var PB_PageParam_GetCommentsPage__FOlD = &PB_PageParam_GetCommentsPage{
	PostId: 0,
}

type PB_PageResponse_GetCommentsPage_Flat struct {
	Extra           PB_ResponseExtra
	CommentViewList []PB_CommentView
}

//ToPB
func (m *PB_PageResponse_GetCommentsPage) ToFlat() *PB_PageResponse_GetCommentsPage_Flat {
	r := &PB_PageResponse_GetCommentsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetCommentsPage_Flat) ToPB() *PB_PageResponse_GetCommentsPage {
	r := &PB_PageResponse_GetCommentsPage{}
	return r
}

//folding
var PB_PageResponse_GetCommentsPage__FOlD = &PB_PageResponse_GetCommentsPage{}

type PB_PageParam_GetHomePage_Flat struct {
	Pager PB_Pager
}

//ToPB
func (m *PB_PageParam_GetHomePage) ToFlat() *PB_PageParam_GetHomePage_Flat {
	r := &PB_PageParam_GetHomePage_Flat{}
	return r
}

//ToPB
func (m *PB_PageParam_GetHomePage_Flat) ToPB() *PB_PageParam_GetHomePage {
	r := &PB_PageParam_GetHomePage{}
	return r
}

//folding
var PB_PageParam_GetHomePage__FOlD = &PB_PageParam_GetHomePage{}

type PB_PageResponse_GetHomePage_Flat struct {
	Extra        PB_ResponseExtra
	PostViewList []PB_PostView
}

//ToPB
func (m *PB_PageResponse_GetHomePage) ToFlat() *PB_PageResponse_GetHomePage_Flat {
	r := &PB_PageResponse_GetHomePage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetHomePage_Flat) ToPB() *PB_PageResponse_GetHomePage {
	r := &PB_PageResponse_GetHomePage{}
	return r
}

//folding
var PB_PageResponse_GetHomePage__FOlD = &PB_PageResponse_GetHomePage{}

type PB_PageParam_GetProfilePage_Flat struct {
}

//ToPB
func (m *PB_PageParam_GetProfilePage) ToFlat() *PB_PageParam_GetProfilePage_Flat {
	r := &PB_PageParam_GetProfilePage_Flat{}
	return r
}

//ToPB
func (m *PB_PageParam_GetProfilePage_Flat) ToPB() *PB_PageParam_GetProfilePage {
	r := &PB_PageParam_GetProfilePage{}
	return r
}

//folding
var PB_PageParam_GetProfilePage__FOlD = &PB_PageParam_GetProfilePage{}

type PB_PageResponse_GetProfilePage_Flat struct {
	Extra PB_ResponseExtra
}

//ToPB
func (m *PB_PageResponse_GetProfilePage) ToFlat() *PB_PageResponse_GetProfilePage_Flat {
	r := &PB_PageResponse_GetProfilePage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetProfilePage_Flat) ToPB() *PB_PageResponse_GetProfilePage {
	r := &PB_PageResponse_GetProfilePage{}
	return r
}

//folding
var PB_PageResponse_GetProfilePage__FOlD = &PB_PageResponse_GetProfilePage{}

type PB_PageParam_GetLikesPage_Flat struct {
	PostId int
	Pager  PB_Pager
}

//ToPB
func (m *PB_PageParam_GetLikesPage) ToFlat() *PB_PageParam_GetLikesPage_Flat {
	r := &PB_PageParam_GetLikesPage_Flat{
		PostId: int(m.PostId),
	}
	return r
}

//ToPB
func (m *PB_PageParam_GetLikesPage_Flat) ToPB() *PB_PageParam_GetLikesPage {
	r := &PB_PageParam_GetLikesPage{
		PostId: int64(m.PostId),
	}
	return r
}

//folding
var PB_PageParam_GetLikesPage__FOlD = &PB_PageParam_GetLikesPage{
	PostId: 0,
}

type PB_PageResponse_GetLikesPage_Flat struct {
	Extra        PB_ResponseExtra
	UserViewList []PB_UserView
}

//ToPB
func (m *PB_PageResponse_GetLikesPage) ToFlat() *PB_PageResponse_GetLikesPage_Flat {
	r := &PB_PageResponse_GetLikesPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetLikesPage_Flat) ToPB() *PB_PageResponse_GetLikesPage {
	r := &PB_PageResponse_GetLikesPage{}
	return r
}

//folding
var PB_PageResponse_GetLikesPage__FOlD = &PB_PageResponse_GetLikesPage{}

type PB_PageParam_GetFollowersPage_Flat struct {
	UserId int
	Pager  PB_Pager
}

//ToPB
func (m *PB_PageParam_GetFollowersPage) ToFlat() *PB_PageParam_GetFollowersPage_Flat {
	r := &PB_PageParam_GetFollowersPage_Flat{
		UserId: int(m.UserId),
	}
	return r
}

//ToPB
func (m *PB_PageParam_GetFollowersPage_Flat) ToPB() *PB_PageParam_GetFollowersPage {
	r := &PB_PageParam_GetFollowersPage{
		UserId: int64(m.UserId),
	}
	return r
}

//folding
var PB_PageParam_GetFollowersPage__FOlD = &PB_PageParam_GetFollowersPage{
	UserId: 0,
}

type PB_PageResponse_GetFollowersPage_Flat struct {
	Extra        PB_ResponseExtra
	UserViewList []PB_UserView
}

//ToPB
func (m *PB_PageResponse_GetFollowersPage) ToFlat() *PB_PageResponse_GetFollowersPage_Flat {
	r := &PB_PageResponse_GetFollowersPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetFollowersPage_Flat) ToPB() *PB_PageResponse_GetFollowersPage {
	r := &PB_PageResponse_GetFollowersPage{}
	return r
}

//folding
var PB_PageResponse_GetFollowersPage__FOlD = &PB_PageResponse_GetFollowersPage{}

type PB_PageParam_GetFollowingsPage_Flat struct {
	UserId int
	Pager  PB_Pager
}

//ToPB
func (m *PB_PageParam_GetFollowingsPage) ToFlat() *PB_PageParam_GetFollowingsPage_Flat {
	r := &PB_PageParam_GetFollowingsPage_Flat{
		UserId: int(m.UserId),
	}
	return r
}

//ToPB
func (m *PB_PageParam_GetFollowingsPage_Flat) ToPB() *PB_PageParam_GetFollowingsPage {
	r := &PB_PageParam_GetFollowingsPage{
		UserId: int64(m.UserId),
	}
	return r
}

//folding
var PB_PageParam_GetFollowingsPage__FOlD = &PB_PageParam_GetFollowingsPage{
	UserId: 0,
}

type PB_PageResponse_GetFollowingsPage_Flat struct {
	Extra        PB_ResponseExtra
	UserViewList []PB_UserView
}

//ToPB
func (m *PB_PageResponse_GetFollowingsPage) ToFlat() *PB_PageResponse_GetFollowingsPage_Flat {
	r := &PB_PageResponse_GetFollowingsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetFollowingsPage_Flat) ToPB() *PB_PageResponse_GetFollowingsPage {
	r := &PB_PageResponse_GetFollowingsPage{}
	return r
}

//folding
var PB_PageResponse_GetFollowingsPage__FOlD = &PB_PageResponse_GetFollowingsPage{}

type PB_PageParam_GetNotifiesPage_Flat struct {
	Pager PB_Pager
}

//ToPB
func (m *PB_PageParam_GetNotifiesPage) ToFlat() *PB_PageParam_GetNotifiesPage_Flat {
	r := &PB_PageParam_GetNotifiesPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageParam_GetNotifiesPage_Flat) ToPB() *PB_PageParam_GetNotifiesPage {
	r := &PB_PageParam_GetNotifiesPage{}
	return r
}

//folding
var PB_PageParam_GetNotifiesPage__FOlD = &PB_PageParam_GetNotifiesPage{}

type PB_PageResponse_GetNotifiesPage_Flat struct {
	Extra          PB_ResponseExtra
	NotifyViewList []PB_NotifyView
	RemoveIdsList  []int
}

//ToPB
func (m *PB_PageResponse_GetNotifiesPage) ToFlat() *PB_PageResponse_GetNotifiesPage_Flat {
	r := &PB_PageResponse_GetNotifiesPage_Flat{

		RemoveIdsList: helper.SliceInt64ToInt(m.RemoveIdsList),
	}
	return r
}

//ToPB
func (m *PB_PageResponse_GetNotifiesPage_Flat) ToPB() *PB_PageResponse_GetNotifiesPage {
	r := &PB_PageResponse_GetNotifiesPage{

		RemoveIdsList: helper.SliceIntToInt64(m.RemoveIdsList),
	}
	return r
}

//folding
var PB_PageResponse_GetNotifiesPage__FOlD = &PB_PageResponse_GetNotifiesPage{}

type PB_PageParam_GetUserActionsPage_Flat struct {
	Pager PB_Pager
}

//ToPB
func (m *PB_PageParam_GetUserActionsPage) ToFlat() *PB_PageParam_GetUserActionsPage_Flat {
	r := &PB_PageParam_GetUserActionsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageParam_GetUserActionsPage_Flat) ToPB() *PB_PageParam_GetUserActionsPage {
	r := &PB_PageParam_GetUserActionsPage{}
	return r
}

//folding
var PB_PageParam_GetUserActionsPage__FOlD = &PB_PageParam_GetUserActionsPage{}

type PB_PageResponse_GetUserActionsPage_Flat struct {
	Extra          PB_ResponseExtra
	ActionViewList []PB_ActionView
}

//ToPB
func (m *PB_PageResponse_GetUserActionsPage) ToFlat() *PB_PageResponse_GetUserActionsPage_Flat {
	r := &PB_PageResponse_GetUserActionsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetUserActionsPage_Flat) ToPB() *PB_PageResponse_GetUserActionsPage {
	r := &PB_PageResponse_GetUserActionsPage{}
	return r
}

//folding
var PB_PageResponse_GetUserActionsPage__FOlD = &PB_PageResponse_GetUserActionsPage{}

type PB_PageParam_GetSuggestedPostsPage_Flat struct {
	Pager PB_Pager
}

//ToPB
func (m *PB_PageParam_GetSuggestedPostsPage) ToFlat() *PB_PageParam_GetSuggestedPostsPage_Flat {
	r := &PB_PageParam_GetSuggestedPostsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageParam_GetSuggestedPostsPage_Flat) ToPB() *PB_PageParam_GetSuggestedPostsPage {
	r := &PB_PageParam_GetSuggestedPostsPage{}
	return r
}

//folding
var PB_PageParam_GetSuggestedPostsPage__FOlD = &PB_PageParam_GetSuggestedPostsPage{}

type PB_PageResponse_GetSuggestedPostsPage_Flat struct {
	Extra        PB_ResponseExtra
	PostViewList []PB_PostView
}

//ToPB
func (m *PB_PageResponse_GetSuggestedPostsPage) ToFlat() *PB_PageResponse_GetSuggestedPostsPage_Flat {
	r := &PB_PageResponse_GetSuggestedPostsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetSuggestedPostsPage_Flat) ToPB() *PB_PageResponse_GetSuggestedPostsPage {
	r := &PB_PageResponse_GetSuggestedPostsPage{}
	return r
}

//folding
var PB_PageResponse_GetSuggestedPostsPage__FOlD = &PB_PageResponse_GetSuggestedPostsPage{}

type PB_PageParam_GetSuggestedUsrsPage_Flat struct {
	Pager PB_Pager
}

//ToPB
func (m *PB_PageParam_GetSuggestedUsrsPage) ToFlat() *PB_PageParam_GetSuggestedUsrsPage_Flat {
	r := &PB_PageParam_GetSuggestedUsrsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageParam_GetSuggestedUsrsPage_Flat) ToPB() *PB_PageParam_GetSuggestedUsrsPage {
	r := &PB_PageParam_GetSuggestedUsrsPage{}
	return r
}

//folding
var PB_PageParam_GetSuggestedUsrsPage__FOlD = &PB_PageParam_GetSuggestedUsrsPage{}

type PB_PageResponse_GetSuggestedUsrsPage_Flat struct {
	Extra        PB_ResponseExtra
	UserViewList []PB_UserView
}

//ToPB
func (m *PB_PageResponse_GetSuggestedUsrsPage) ToFlat() *PB_PageResponse_GetSuggestedUsrsPage_Flat {
	r := &PB_PageResponse_GetSuggestedUsrsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetSuggestedUsrsPage_Flat) ToPB() *PB_PageResponse_GetSuggestedUsrsPage {
	r := &PB_PageResponse_GetSuggestedUsrsPage{}
	return r
}

//folding
var PB_PageResponse_GetSuggestedUsrsPage__FOlD = &PB_PageResponse_GetSuggestedUsrsPage{}

type PB_PageParam_GetSuggestedTagsPage_Flat struct {
	Pager PB_Pager
}

//ToPB
func (m *PB_PageParam_GetSuggestedTagsPage) ToFlat() *PB_PageParam_GetSuggestedTagsPage_Flat {
	r := &PB_PageParam_GetSuggestedTagsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageParam_GetSuggestedTagsPage_Flat) ToPB() *PB_PageParam_GetSuggestedTagsPage {
	r := &PB_PageParam_GetSuggestedTagsPage{}
	return r
}

//folding
var PB_PageParam_GetSuggestedTagsPage__FOlD = &PB_PageParam_GetSuggestedTagsPage{}

type PB_PageResponse_GetSuggestedTagsPage_Flat struct {
	Extra PB_ResponseExtra
}

//ToPB
func (m *PB_PageResponse_GetSuggestedTagsPage) ToFlat() *PB_PageResponse_GetSuggestedTagsPage_Flat {
	r := &PB_PageResponse_GetSuggestedTagsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetSuggestedTagsPage_Flat) ToPB() *PB_PageResponse_GetSuggestedTagsPage {
	r := &PB_PageResponse_GetSuggestedTagsPage{}
	return r
}

//folding
var PB_PageResponse_GetSuggestedTagsPage__FOlD = &PB_PageResponse_GetSuggestedTagsPage{}

type PB_PageParam_GetLastPostsPage_Flat struct {
	Pager PB_Pager
}

//ToPB
func (m *PB_PageParam_GetLastPostsPage) ToFlat() *PB_PageParam_GetLastPostsPage_Flat {
	r := &PB_PageParam_GetLastPostsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageParam_GetLastPostsPage_Flat) ToPB() *PB_PageParam_GetLastPostsPage {
	r := &PB_PageParam_GetLastPostsPage{}
	return r
}

//folding
var PB_PageParam_GetLastPostsPage__FOlD = &PB_PageParam_GetLastPostsPage{}

type PB_PageResponse_GetLastPostsPage_Flat struct {
	Extra        PB_ResponseExtra
	UserViewList []PB_UserView
}

//ToPB
func (m *PB_PageResponse_GetLastPostsPage) ToFlat() *PB_PageResponse_GetLastPostsPage_Flat {
	r := &PB_PageResponse_GetLastPostsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetLastPostsPage_Flat) ToPB() *PB_PageResponse_GetLastPostsPage {
	r := &PB_PageResponse_GetLastPostsPage{}
	return r
}

//folding
var PB_PageResponse_GetLastPostsPage__FOlD = &PB_PageResponse_GetLastPostsPage{}

type PB_PageParam_GetTagPage_Flat struct {
	Pager PB_Pager
}

//ToPB
func (m *PB_PageParam_GetTagPage) ToFlat() *PB_PageParam_GetTagPage_Flat {
	r := &PB_PageParam_GetTagPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageParam_GetTagPage_Flat) ToPB() *PB_PageParam_GetTagPage {
	r := &PB_PageParam_GetTagPage{}
	return r
}

//folding
var PB_PageParam_GetTagPage__FOlD = &PB_PageParam_GetTagPage{}

type PB_PageResponse_GetTagPage_Flat struct {
	Extra        PB_ResponseExtra
	PostViewList []PB_PostView
}

//ToPB
func (m *PB_PageResponse_GetTagPage) ToFlat() *PB_PageResponse_GetTagPage_Flat {
	r := &PB_PageResponse_GetTagPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_GetTagPage_Flat) ToPB() *PB_PageResponse_GetTagPage {
	r := &PB_PageResponse_GetTagPage{}
	return r
}

//folding
var PB_PageResponse_GetTagPage__FOlD = &PB_PageResponse_GetTagPage{}

type PB_PageParam_SearchTagsPage_Flat struct {
}

//ToPB
func (m *PB_PageParam_SearchTagsPage) ToFlat() *PB_PageParam_SearchTagsPage_Flat {
	r := &PB_PageParam_SearchTagsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageParam_SearchTagsPage_Flat) ToPB() *PB_PageParam_SearchTagsPage {
	r := &PB_PageParam_SearchTagsPage{}
	return r
}

//folding
var PB_PageParam_SearchTagsPage__FOlD = &PB_PageParam_SearchTagsPage{}

type PB_PageResponse_SearchTagsPage_Flat struct {
	Extra PB_ResponseExtra
}

//ToPB
func (m *PB_PageResponse_SearchTagsPage) ToFlat() *PB_PageResponse_SearchTagsPage_Flat {
	r := &PB_PageResponse_SearchTagsPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_SearchTagsPage_Flat) ToPB() *PB_PageResponse_SearchTagsPage {
	r := &PB_PageResponse_SearchTagsPage{}
	return r
}

//folding
var PB_PageResponse_SearchTagsPage__FOlD = &PB_PageResponse_SearchTagsPage{}

type PB_PageParam_SearchUsersPage_Flat struct {
}

//ToPB
func (m *PB_PageParam_SearchUsersPage) ToFlat() *PB_PageParam_SearchUsersPage_Flat {
	r := &PB_PageParam_SearchUsersPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageParam_SearchUsersPage_Flat) ToPB() *PB_PageParam_SearchUsersPage {
	r := &PB_PageParam_SearchUsersPage{}
	return r
}

//folding
var PB_PageParam_SearchUsersPage__FOlD = &PB_PageParam_SearchUsersPage{}

type PB_PageResponse_SearchUsersPage_Flat struct {
	Extra PB_ResponseExtra
}

//ToPB
func (m *PB_PageResponse_SearchUsersPage) ToFlat() *PB_PageResponse_SearchUsersPage_Flat {
	r := &PB_PageResponse_SearchUsersPage_Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse_SearchUsersPage_Flat) ToPB() *PB_PageResponse_SearchUsersPage {
	r := &PB_PageResponse_SearchUsersPage{}
	return r
}

//folding
var PB_PageResponse_SearchUsersPage__FOlD = &PB_PageResponse_SearchUsersPage{}

type PB_PageParam__Flat struct {
}

//ToPB
func (m *PB_PageParam_) ToFlat() *PB_PageParam__Flat {
	r := &PB_PageParam__Flat{}
	return r
}

//ToPB
func (m *PB_PageParam__Flat) ToPB() *PB_PageParam_ {
	r := &PB_PageParam_{}
	return r
}

//folding
var PB_PageParam___FOlD = &PB_PageParam_{}

type PB_PageResponse__Flat struct {
	Extra PB_ResponseExtra
}

//ToPB
func (m *PB_PageResponse_) ToFlat() *PB_PageResponse__Flat {
	r := &PB_PageResponse__Flat{}
	return r
}

//ToPB
func (m *PB_PageResponse__Flat) ToPB() *PB_PageResponse_ {
	r := &PB_PageResponse_{}
	return r
}

//folding
var PB_PageResponse___FOlD = &PB_PageResponse_{}

type PB_SearchResponse_AddNewC_Flat struct {
}

//ToPB
func (m *PB_SearchResponse_AddNewC) ToFlat() *PB_SearchResponse_AddNewC_Flat {
	r := &PB_SearchResponse_AddNewC_Flat{}
	return r
}

//ToPB
func (m *PB_SearchResponse_AddNewC_Flat) ToPB() *PB_SearchResponse_AddNewC {
	r := &PB_SearchResponse_AddNewC{}
	return r
}

//folding
var PB_SearchResponse_AddNewC__FOlD = &PB_SearchResponse_AddNewC{}

type PB_SocialParam_AddComment_Flat struct {
	PostId int
	Text   string
}

//ToPB
func (m *PB_SocialParam_AddComment) ToFlat() *PB_SocialParam_AddComment_Flat {
	r := &PB_SocialParam_AddComment_Flat{
		PostId: int(m.PostId),
		Text:   m.Text,
	}
	return r
}

//ToPB
func (m *PB_SocialParam_AddComment_Flat) ToPB() *PB_SocialParam_AddComment {
	r := &PB_SocialParam_AddComment{
		PostId: int64(m.PostId),
		Text:   m.Text,
	}
	return r
}

//folding
var PB_SocialParam_AddComment__FOlD = &PB_SocialParam_AddComment{
	PostId: 0,
	Text:   "",
}

type PB_SocialResponse_AddComment_Flat struct {
	Extra   PB_ResponseExtra
	Comment PB_CommentView
}

//ToPB
func (m *PB_SocialResponse_AddComment) ToFlat() *PB_SocialResponse_AddComment_Flat {
	r := &PB_SocialResponse_AddComment_Flat{}
	return r
}

//ToPB
func (m *PB_SocialResponse_AddComment_Flat) ToPB() *PB_SocialResponse_AddComment {
	r := &PB_SocialResponse_AddComment{}
	return r
}

//folding
var PB_SocialResponse_AddComment__FOlD = &PB_SocialResponse_AddComment{}

type PB_SocialParam_DeleteComment_Flat struct {
	PostId    int
	CommentId int
}

//ToPB
func (m *PB_SocialParam_DeleteComment) ToFlat() *PB_SocialParam_DeleteComment_Flat {
	r := &PB_SocialParam_DeleteComment_Flat{
		PostId:    int(m.PostId),
		CommentId: int(m.CommentId),
	}
	return r
}

//ToPB
func (m *PB_SocialParam_DeleteComment_Flat) ToPB() *PB_SocialParam_DeleteComment {
	r := &PB_SocialParam_DeleteComment{
		PostId:    int64(m.PostId),
		CommentId: int64(m.CommentId),
	}
	return r
}

//folding
var PB_SocialParam_DeleteComment__FOlD = &PB_SocialParam_DeleteComment{
	PostId:    0,
	CommentId: 0,
}

type PB_SocialResponse_DeleteComment_Flat struct {
	Extra   PB_ResponseExtra
	Deleted bool
}

//ToPB
func (m *PB_SocialResponse_DeleteComment) ToFlat() *PB_SocialResponse_DeleteComment_Flat {
	r := &PB_SocialResponse_DeleteComment_Flat{

		Deleted: m.Deleted,
	}
	return r
}

//ToPB
func (m *PB_SocialResponse_DeleteComment_Flat) ToPB() *PB_SocialResponse_DeleteComment {
	r := &PB_SocialResponse_DeleteComment{

		Deleted: m.Deleted,
	}
	return r
}

//folding
var PB_SocialResponse_DeleteComment__FOlD = &PB_SocialResponse_DeleteComment{

	Deleted: false,
}

type PB_SocialParam_AddPost_Flat struct {
	Text string
}

//ToPB
func (m *PB_SocialParam_AddPost) ToFlat() *PB_SocialParam_AddPost_Flat {
	r := &PB_SocialParam_AddPost_Flat{
		Text: m.Text,
	}
	return r
}

//ToPB
func (m *PB_SocialParam_AddPost_Flat) ToPB() *PB_SocialParam_AddPost {
	r := &PB_SocialParam_AddPost{
		Text: m.Text,
	}
	return r
}

//folding
var PB_SocialParam_AddPost__FOlD = &PB_SocialParam_AddPost{
	Text: "",
}

type PB_SocialResponse_AddPost_Flat struct {
	Extra    PB_ResponseExtra
	PostView PB_PostView
}

//ToPB
func (m *PB_SocialResponse_AddPost) ToFlat() *PB_SocialResponse_AddPost_Flat {
	r := &PB_SocialResponse_AddPost_Flat{}
	return r
}

//ToPB
func (m *PB_SocialResponse_AddPost_Flat) ToPB() *PB_SocialResponse_AddPost {
	r := &PB_SocialResponse_AddPost{}
	return r
}

//folding
var PB_SocialResponse_AddPost__FOlD = &PB_SocialResponse_AddPost{}

type PB_SocialParam_EditPost_Flat struct {
	PostId int
	Text   string
}

//ToPB
func (m *PB_SocialParam_EditPost) ToFlat() *PB_SocialParam_EditPost_Flat {
	r := &PB_SocialParam_EditPost_Flat{
		PostId: int(m.PostId),
		Text:   m.Text,
	}
	return r
}

//ToPB
func (m *PB_SocialParam_EditPost_Flat) ToPB() *PB_SocialParam_EditPost {
	r := &PB_SocialParam_EditPost{
		PostId: int64(m.PostId),
		Text:   m.Text,
	}
	return r
}

//folding
var PB_SocialParam_EditPost__FOlD = &PB_SocialParam_EditPost{
	PostId: 0,
	Text:   "",
}

type PB_SocialResponse_EditPost_Flat struct {
	Extra    PB_ResponseExtra
	PostView PB_PostView
}

//ToPB
func (m *PB_SocialResponse_EditPost) ToFlat() *PB_SocialResponse_EditPost_Flat {
	r := &PB_SocialResponse_EditPost_Flat{}
	return r
}

//ToPB
func (m *PB_SocialResponse_EditPost_Flat) ToPB() *PB_SocialResponse_EditPost {
	r := &PB_SocialResponse_EditPost{}
	return r
}

//folding
var PB_SocialResponse_EditPost__FOlD = &PB_SocialResponse_EditPost{}

type PB_SocialParam_DeletePost_Flat struct {
	PostId int
}

//ToPB
func (m *PB_SocialParam_DeletePost) ToFlat() *PB_SocialParam_DeletePost_Flat {
	r := &PB_SocialParam_DeletePost_Flat{
		PostId: int(m.PostId),
	}
	return r
}

//ToPB
func (m *PB_SocialParam_DeletePost_Flat) ToPB() *PB_SocialParam_DeletePost {
	r := &PB_SocialParam_DeletePost{
		PostId: int64(m.PostId),
	}
	return r
}

//folding
var PB_SocialParam_DeletePost__FOlD = &PB_SocialParam_DeletePost{
	PostId: 0,
}

type PB_SocialResponse_DeletePost_Flat struct {
	Extra PB_ResponseExtra
	Done  bool
}

//ToPB
func (m *PB_SocialResponse_DeletePost) ToFlat() *PB_SocialResponse_DeletePost_Flat {
	r := &PB_SocialResponse_DeletePost_Flat{

		Done: m.Done,
	}
	return r
}

//ToPB
func (m *PB_SocialResponse_DeletePost_Flat) ToPB() *PB_SocialResponse_DeletePost {
	r := &PB_SocialResponse_DeletePost{

		Done: m.Done,
	}
	return r
}

//folding
var PB_SocialResponse_DeletePost__FOlD = &PB_SocialResponse_DeletePost{

	Done: false,
}

type PB_SocialParam_ArchivePost_Flat struct {
	PostId int
}

//ToPB
func (m *PB_SocialParam_ArchivePost) ToFlat() *PB_SocialParam_ArchivePost_Flat {
	r := &PB_SocialParam_ArchivePost_Flat{
		PostId: int(m.PostId),
	}
	return r
}

//ToPB
func (m *PB_SocialParam_ArchivePost_Flat) ToPB() *PB_SocialParam_ArchivePost {
	r := &PB_SocialParam_ArchivePost{
		PostId: int64(m.PostId),
	}
	return r
}

//folding
var PB_SocialParam_ArchivePost__FOlD = &PB_SocialParam_ArchivePost{
	PostId: 0,
}

type PB_SocialResponse_ArchivePost_Flat struct {
	Extra PB_ResponseExtra
	Done  bool
}

//ToPB
func (m *PB_SocialResponse_ArchivePost) ToFlat() *PB_SocialResponse_ArchivePost_Flat {
	r := &PB_SocialResponse_ArchivePost_Flat{

		Done: m.Done,
	}
	return r
}

//ToPB
func (m *PB_SocialResponse_ArchivePost_Flat) ToPB() *PB_SocialResponse_ArchivePost {
	r := &PB_SocialResponse_ArchivePost{

		Done: m.Done,
	}
	return r
}

//folding
var PB_SocialResponse_ArchivePost__FOlD = &PB_SocialResponse_ArchivePost{

	Done: false,
}

type PB_SocialParam_LikePost_Flat struct {
	PostId int
}

//ToPB
func (m *PB_SocialParam_LikePost) ToFlat() *PB_SocialParam_LikePost_Flat {
	r := &PB_SocialParam_LikePost_Flat{
		PostId: int(m.PostId),
	}
	return r
}

//ToPB
func (m *PB_SocialParam_LikePost_Flat) ToPB() *PB_SocialParam_LikePost {
	r := &PB_SocialParam_LikePost{
		PostId: int64(m.PostId),
	}
	return r
}

//folding
var PB_SocialParam_LikePost__FOlD = &PB_SocialParam_LikePost{
	PostId: 0,
}

type PB_SocialResponse_LikePost_Flat struct {
	Extra PB_ResponseExtra
	Done  bool
}

//ToPB
func (m *PB_SocialResponse_LikePost) ToFlat() *PB_SocialResponse_LikePost_Flat {
	r := &PB_SocialResponse_LikePost_Flat{

		Done: m.Done,
	}
	return r
}

//ToPB
func (m *PB_SocialResponse_LikePost_Flat) ToPB() *PB_SocialResponse_LikePost {
	r := &PB_SocialResponse_LikePost{

		Done: m.Done,
	}
	return r
}

//folding
var PB_SocialResponse_LikePost__FOlD = &PB_SocialResponse_LikePost{

	Done: false,
}

type PB_SocialParam_UnLikePost_Flat struct {
	PostId int
}

//ToPB
func (m *PB_SocialParam_UnLikePost) ToFlat() *PB_SocialParam_UnLikePost_Flat {
	r := &PB_SocialParam_UnLikePost_Flat{
		PostId: int(m.PostId),
	}
	return r
}

//ToPB
func (m *PB_SocialParam_UnLikePost_Flat) ToPB() *PB_SocialParam_UnLikePost {
	r := &PB_SocialParam_UnLikePost{
		PostId: int64(m.PostId),
	}
	return r
}

//folding
var PB_SocialParam_UnLikePost__FOlD = &PB_SocialParam_UnLikePost{
	PostId: 0,
}

type PB_SocialResponse_UnLikePost_Flat struct {
	Extra PB_ResponseExtra
	Done  bool
}

//ToPB
func (m *PB_SocialResponse_UnLikePost) ToFlat() *PB_SocialResponse_UnLikePost_Flat {
	r := &PB_SocialResponse_UnLikePost_Flat{

		Done: m.Done,
	}
	return r
}

//ToPB
func (m *PB_SocialResponse_UnLikePost_Flat) ToPB() *PB_SocialResponse_UnLikePost {
	r := &PB_SocialResponse_UnLikePost{

		Done: m.Done,
	}
	return r
}

//folding
var PB_SocialResponse_UnLikePost__FOlD = &PB_SocialResponse_UnLikePost{

	Done: false,
}

type PB_SocialParam_FollowUser_Flat struct {
	UserId int
}

//ToPB
func (m *PB_SocialParam_FollowUser) ToFlat() *PB_SocialParam_FollowUser_Flat {
	r := &PB_SocialParam_FollowUser_Flat{
		UserId: int(m.UserId),
	}
	return r
}

//ToPB
func (m *PB_SocialParam_FollowUser_Flat) ToPB() *PB_SocialParam_FollowUser {
	r := &PB_SocialParam_FollowUser{
		UserId: int64(m.UserId),
	}
	return r
}

//folding
var PB_SocialParam_FollowUser__FOlD = &PB_SocialParam_FollowUser{
	UserId: 0,
}

type PB_SocialResponse_FollowUser_Flat struct {
	Extra PB_ResponseExtra
}

//ToPB
func (m *PB_SocialResponse_FollowUser) ToFlat() *PB_SocialResponse_FollowUser_Flat {
	r := &PB_SocialResponse_FollowUser_Flat{}
	return r
}

//ToPB
func (m *PB_SocialResponse_FollowUser_Flat) ToPB() *PB_SocialResponse_FollowUser {
	r := &PB_SocialResponse_FollowUser{}
	return r
}

//folding
var PB_SocialResponse_FollowUser__FOlD = &PB_SocialResponse_FollowUser{}

type PB_SocialParam_UnFollowUser_Flat struct {
	UserId int
}

//ToPB
func (m *PB_SocialParam_UnFollowUser) ToFlat() *PB_SocialParam_UnFollowUser_Flat {
	r := &PB_SocialParam_UnFollowUser_Flat{
		UserId: int(m.UserId),
	}
	return r
}

//ToPB
func (m *PB_SocialParam_UnFollowUser_Flat) ToPB() *PB_SocialParam_UnFollowUser {
	r := &PB_SocialParam_UnFollowUser{
		UserId: int64(m.UserId),
	}
	return r
}

//folding
var PB_SocialParam_UnFollowUser__FOlD = &PB_SocialParam_UnFollowUser{
	UserId: 0,
}

type PB_SocialResponse_UnFollowUser_Flat struct {
	Extra PB_ResponseExtra
}

//ToPB
func (m *PB_SocialResponse_UnFollowUser) ToFlat() *PB_SocialResponse_UnFollowUser_Flat {
	r := &PB_SocialResponse_UnFollowUser_Flat{}
	return r
}

//ToPB
func (m *PB_SocialResponse_UnFollowUser_Flat) ToPB() *PB_SocialResponse_UnFollowUser {
	r := &PB_SocialResponse_UnFollowUser{}
	return r
}

//folding
var PB_SocialResponse_UnFollowUser__FOlD = &PB_SocialResponse_UnFollowUser{}

type PB_UserParam_BlockUser_Flat struct {
	UserId   int
	UserName string
}

//ToPB
func (m *PB_UserParam_BlockUser) ToFlat() *PB_UserParam_BlockUser_Flat {
	r := &PB_UserParam_BlockUser_Flat{
		UserId:   int(m.UserId),
		UserName: m.UserName,
	}
	return r
}

//ToPB
func (m *PB_UserParam_BlockUser_Flat) ToPB() *PB_UserParam_BlockUser {
	r := &PB_UserParam_BlockUser{
		UserId:   int32(m.UserId),
		UserName: m.UserName,
	}
	return r
}

//folding
var PB_UserParam_BlockUser__FOlD = &PB_UserParam_BlockUser{
	UserId:   0,
	UserName: "",
}

type PB_UserResponse_BlockUser_Flat struct {
	ByUserId       int
	TargetUserId   int
	TargetUserName string
}

//ToPB
func (m *PB_UserResponse_BlockUser) ToFlat() *PB_UserResponse_BlockUser_Flat {
	r := &PB_UserResponse_BlockUser_Flat{
		ByUserId:       int(m.ByUserId),
		TargetUserId:   int(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

//ToPB
func (m *PB_UserResponse_BlockUser_Flat) ToPB() *PB_UserResponse_BlockUser {
	r := &PB_UserResponse_BlockUser{
		ByUserId:       int32(m.ByUserId),
		TargetUserId:   int32(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

//folding
var PB_UserResponse_BlockUser__FOlD = &PB_UserResponse_BlockUser{
	ByUserId:       0,
	TargetUserId:   0,
	TargetUserName: "",
}

type PB_UserParam_UnBlockUser_Flat struct {
	Offset int
	Limit  int
}

//ToPB
func (m *PB_UserParam_UnBlockUser) ToFlat() *PB_UserParam_UnBlockUser_Flat {
	r := &PB_UserParam_UnBlockUser_Flat{
		Offset: int(m.Offset),
		Limit:  int(m.Limit),
	}
	return r
}

//ToPB
func (m *PB_UserParam_UnBlockUser_Flat) ToPB() *PB_UserParam_UnBlockUser {
	r := &PB_UserParam_UnBlockUser{
		Offset: int32(m.Offset),
		Limit:  int32(m.Limit),
	}
	return r
}

//folding
var PB_UserParam_UnBlockUser__FOlD = &PB_UserParam_UnBlockUser{
	Offset: 0,
	Limit:  0,
}

type PB_UserResponse_UnBlockUser_Flat struct {
	Users []UserView
}

//ToPB
func (m *PB_UserResponse_UnBlockUser) ToFlat() *PB_UserResponse_UnBlockUser_Flat {
	r := &PB_UserResponse_UnBlockUser_Flat{}
	return r
}

//ToPB
func (m *PB_UserResponse_UnBlockUser_Flat) ToPB() *PB_UserResponse_UnBlockUser {
	r := &PB_UserResponse_UnBlockUser{}
	return r
}

//folding
var PB_UserResponse_UnBlockUser__FOlD = &PB_UserResponse_UnBlockUser{}

type PB_UserParam_BlockedList_Flat struct {
	UserId   int
	UserName string
}

//ToPB
func (m *PB_UserParam_BlockedList) ToFlat() *PB_UserParam_BlockedList_Flat {
	r := &PB_UserParam_BlockedList_Flat{
		UserId:   int(m.UserId),
		UserName: m.UserName,
	}
	return r
}

//ToPB
func (m *PB_UserParam_BlockedList_Flat) ToPB() *PB_UserParam_BlockedList {
	r := &PB_UserParam_BlockedList{
		UserId:   int32(m.UserId),
		UserName: m.UserName,
	}
	return r
}

//folding
var PB_UserParam_BlockedList__FOlD = &PB_UserParam_BlockedList{
	UserId:   0,
	UserName: "",
}

type PB_UserResponse_BlockedList_Flat struct {
	ByUserId       int
	TargetUserId   int
	TargetUserName string
}

//ToPB
func (m *PB_UserResponse_BlockedList) ToFlat() *PB_UserResponse_BlockedList_Flat {
	r := &PB_UserResponse_BlockedList_Flat{
		ByUserId:       int(m.ByUserId),
		TargetUserId:   int(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

//ToPB
func (m *PB_UserResponse_BlockedList_Flat) ToPB() *PB_UserResponse_BlockedList {
	r := &PB_UserResponse_BlockedList{
		ByUserId:       int32(m.ByUserId),
		TargetUserId:   int32(m.TargetUserId),
		TargetUserName: m.TargetUserName,
	}
	return r
}

//folding
var PB_UserResponse_BlockedList__FOlD = &PB_UserResponse_BlockedList{
	ByUserId:       0,
	TargetUserId:   0,
	TargetUserName: "",
}

type PB_UserParam_UpdateAbout_Flat struct {
	NewAbout string
}

//ToPB
func (m *PB_UserParam_UpdateAbout) ToFlat() *PB_UserParam_UpdateAbout_Flat {
	r := &PB_UserParam_UpdateAbout_Flat{
		NewAbout: m.NewAbout,
	}
	return r
}

//ToPB
func (m *PB_UserParam_UpdateAbout_Flat) ToPB() *PB_UserParam_UpdateAbout {
	r := &PB_UserParam_UpdateAbout{
		NewAbout: m.NewAbout,
	}
	return r
}

//folding
var PB_UserParam_UpdateAbout__FOlD = &PB_UserParam_UpdateAbout{
	NewAbout: "",
}

type PB_UserResponse_UpdateAbout_Flat struct {
	UserId   int
	NewAbout string
}

//ToPB
func (m *PB_UserResponse_UpdateAbout) ToFlat() *PB_UserResponse_UpdateAbout_Flat {
	r := &PB_UserResponse_UpdateAbout_Flat{
		UserId:   int(m.UserId),
		NewAbout: m.NewAbout,
	}
	return r
}

//ToPB
func (m *PB_UserResponse_UpdateAbout_Flat) ToPB() *PB_UserResponse_UpdateAbout {
	r := &PB_UserResponse_UpdateAbout{
		UserId:   int32(m.UserId),
		NewAbout: m.NewAbout,
	}
	return r
}

//folding
var PB_UserResponse_UpdateAbout__FOlD = &PB_UserResponse_UpdateAbout{
	UserId:   0,
	NewAbout: "",
}

type PB_UserParam_UpdateUserName_Flat struct {
	NewUserName string
}

//ToPB
func (m *PB_UserParam_UpdateUserName) ToFlat() *PB_UserParam_UpdateUserName_Flat {
	r := &PB_UserParam_UpdateUserName_Flat{
		NewUserName: m.NewUserName,
	}
	return r
}

//ToPB
func (m *PB_UserParam_UpdateUserName_Flat) ToPB() *PB_UserParam_UpdateUserName {
	r := &PB_UserParam_UpdateUserName{
		NewUserName: m.NewUserName,
	}
	return r
}

//folding
var PB_UserParam_UpdateUserName__FOlD = &PB_UserParam_UpdateUserName{
	NewUserName: "",
}

type PB_UserResponse_UpdateUserName_Flat struct {
	UserId      int
	NewUserName string
}

//ToPB
func (m *PB_UserResponse_UpdateUserName) ToFlat() *PB_UserResponse_UpdateUserName_Flat {
	r := &PB_UserResponse_UpdateUserName_Flat{
		UserId:      int(m.UserId),
		NewUserName: m.NewUserName,
	}
	return r
}

//ToPB
func (m *PB_UserResponse_UpdateUserName_Flat) ToPB() *PB_UserResponse_UpdateUserName {
	r := &PB_UserResponse_UpdateUserName{
		UserId:      int32(m.UserId),
		NewUserName: m.NewUserName,
	}
	return r
}

//folding
var PB_UserResponse_UpdateUserName__FOlD = &PB_UserResponse_UpdateUserName{
	UserId:      0,
	NewUserName: "",
}

type PB_UserParam_ChangeAvatar_Flat struct {
	None       bool
	ImageData2 []byte
}

//ToPB
func (m *PB_UserParam_ChangeAvatar) ToFlat() *PB_UserParam_ChangeAvatar_Flat {
	r := &PB_UserParam_ChangeAvatar_Flat{
		None:       m.None,
		ImageData2: []byte(m.ImageData2),
	}
	return r
}

//ToPB
func (m *PB_UserParam_ChangeAvatar_Flat) ToPB() *PB_UserParam_ChangeAvatar {
	r := &PB_UserParam_ChangeAvatar{
		None:       m.None,
		ImageData2: m.ImageData2,
	}
	return r
}

//folding
var PB_UserParam_ChangeAvatar__FOlD = &PB_UserParam_ChangeAvatar{
	None:       false,
	ImageData2: []byte{},
}

type PB_UserResponse_ChangeAvatar_Flat struct {
}

//ToPB
func (m *PB_UserResponse_ChangeAvatar) ToFlat() *PB_UserResponse_ChangeAvatar_Flat {
	r := &PB_UserResponse_ChangeAvatar_Flat{}
	return r
}

//ToPB
func (m *PB_UserResponse_ChangeAvatar_Flat) ToPB() *PB_UserResponse_ChangeAvatar {
	r := &PB_UserResponse_ChangeAvatar{}
	return r
}

//folding
var PB_UserResponse_ChangeAvatar__FOlD = &PB_UserResponse_ChangeAvatar{}

type PB_UserParam_ChangePrivacy_Flat struct {
	Level ProfilePrivacyLevelEnum
}

//ToPB
func (m *PB_UserParam_ChangePrivacy) ToFlat() *PB_UserParam_ChangePrivacy_Flat {
	r := &PB_UserParam_ChangePrivacy_Flat{}
	return r
}

//ToPB
func (m *PB_UserParam_ChangePrivacy_Flat) ToPB() *PB_UserParam_ChangePrivacy {
	r := &PB_UserParam_ChangePrivacy{}
	return r
}

//folding
var PB_UserParam_ChangePrivacy__FOlD = &PB_UserParam_ChangePrivacy{}

type PB_UserResponseOffline_ChangePrivacy_Flat struct {
}

//ToPB
func (m *PB_UserResponseOffline_ChangePrivacy) ToFlat() *PB_UserResponseOffline_ChangePrivacy_Flat {
	r := &PB_UserResponseOffline_ChangePrivacy_Flat{}
	return r
}

//ToPB
func (m *PB_UserResponseOffline_ChangePrivacy_Flat) ToPB() *PB_UserResponseOffline_ChangePrivacy {
	r := &PB_UserResponseOffline_ChangePrivacy{}
	return r
}

//folding
var PB_UserResponseOffline_ChangePrivacy__FOlD = &PB_UserResponseOffline_ChangePrivacy{}

type PB_UserParam_CheckUserName_Flat struct {
	Level ProfilePrivacyLevelEnum
}

//ToPB
func (m *PB_UserParam_CheckUserName) ToFlat() *PB_UserParam_CheckUserName_Flat {
	r := &PB_UserParam_CheckUserName_Flat{}
	return r
}

//ToPB
func (m *PB_UserParam_CheckUserName_Flat) ToPB() *PB_UserParam_CheckUserName {
	r := &PB_UserParam_CheckUserName{}
	return r
}

//folding
var PB_UserParam_CheckUserName__FOlD = &PB_UserParam_CheckUserName{}

type PB_UserResponse_CheckUserName_Flat struct {
}

//ToPB
func (m *PB_UserResponse_CheckUserName) ToFlat() *PB_UserResponse_CheckUserName_Flat {
	r := &PB_UserResponse_CheckUserName_Flat{}
	return r
}

//ToPB
func (m *PB_UserResponse_CheckUserName_Flat) ToPB() *PB_UserResponse_CheckUserName {
	r := &PB_UserResponse_CheckUserName{}
	return r
}

//folding
var PB_UserResponse_CheckUserName__FOlD = &PB_UserResponse_CheckUserName{}

type UserView_Flat struct {
}

//ToPB
func (m *UserView) ToFlat() *UserView_Flat {
	r := &UserView_Flat{}
	return r
}

//ToPB
func (m *UserView_Flat) ToPB() *UserView {
	r := &UserView{}
	return r
}

//folding
var UserView__FOlD = &UserView{}

type PB_Action_Flat struct {
	ActionId       int
	ActorUserId    int
	ActionTypeEnum int
	PeerUserId     int
	PostId         int
	CommentId      int
	Murmur64Hash   int
	CreatedTime    int
	Seq            int
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
		Seq:            int(m.Seq),
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
		Seq:            int32(m.Seq),
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
	Seq:            0,
}

type PB_Chat_Flat struct {
	ChatKey            string
	RoomKey            string
	RoomTypeEnum       int
	UserId             int
	PeerUserId         int
	GroupId            int
	CreatedTime        int
	StartMessageIdFrom int
	LastMessageId      int
	LastSeenMessageId  int
	UpdatedMs          int
}

//ToPB
func (m *PB_Chat) ToFlat() *PB_Chat_Flat {
	r := &PB_Chat_Flat{
		ChatKey:            m.ChatKey,
		RoomKey:            m.RoomKey,
		RoomTypeEnum:       int(m.RoomTypeEnum),
		UserId:             int(m.UserId),
		PeerUserId:         int(m.PeerUserId),
		GroupId:            int(m.GroupId),
		CreatedTime:        int(m.CreatedTime),
		StartMessageIdFrom: int(m.StartMessageIdFrom),
		LastMessageId:      int(m.LastMessageId),
		LastSeenMessageId:  int(m.LastSeenMessageId),
		UpdatedMs:          int(m.UpdatedMs),
	}
	return r
}

//ToPB
func (m *PB_Chat_Flat) ToPB() *PB_Chat {
	r := &PB_Chat{
		ChatKey:            m.ChatKey,
		RoomKey:            m.RoomKey,
		RoomTypeEnum:       int32(m.RoomTypeEnum),
		UserId:             int32(m.UserId),
		PeerUserId:         int32(m.PeerUserId),
		GroupId:            int64(m.GroupId),
		CreatedTime:        int32(m.CreatedTime),
		StartMessageIdFrom: int64(m.StartMessageIdFrom),
		LastMessageId:      int64(m.LastMessageId),
		LastSeenMessageId:  int64(m.LastSeenMessageId),
		UpdatedMs:          int64(m.UpdatedMs),
	}
	return r
}

//folding
var PB_Chat__FOlD = &PB_Chat{
	ChatKey:            "",
	RoomKey:            "",
	RoomTypeEnum:       0,
	UserId:             0,
	PeerUserId:         0,
	GroupId:            0,
	CreatedTime:        0,
	StartMessageIdFrom: 0,
	LastMessageId:      0,
	LastSeenMessageId:  0,
	UpdatedMs:          0,
}

type PB_Comment_Flat struct {
	CommentId   int
	UserId      int
	PostId      int
	Text        string
	LikesCount  int
	CreatedTime int
	Seq         int
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
		Seq:         int(m.Seq),
	}
	return r
}

//ToPB
func (m *PB_Comment_Flat) ToPB() *PB_Comment {
	r := &PB_Comment{
		CommentId:   int64(m.CommentId),
		UserId:      int32(m.UserId),
		PostId:      int32(m.PostId),
		Text:        m.Text,
		LikesCount:  int32(m.LikesCount),
		CreatedTime: int32(m.CreatedTime),
		Seq:         int32(m.Seq),
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
	Seq:         0,
}

type PB_DirectMessage_Flat struct {
	MessageId            int
	MessageKey           string
	RoomKey              string
	UserId               int
	MessageFileId        int
	MessageTypeEnumId    int
	Text                 string
	CreatedSe            int
	PeerReceivedTime     int
	PeerSeenTime         int
	DeliviryStatusEnumId int
}

//ToPB
func (m *PB_DirectMessage) ToFlat() *PB_DirectMessage_Flat {
	r := &PB_DirectMessage_Flat{
		MessageId:            int(m.MessageId),
		MessageKey:           m.MessageKey,
		RoomKey:              m.RoomKey,
		UserId:               int(m.UserId),
		MessageFileId:        int(m.MessageFileId),
		MessageTypeEnumId:    int(m.MessageTypeEnumId),
		Text:                 m.Text,
		CreatedSe:            int(m.CreatedSe),
		PeerReceivedTime:     int(m.PeerReceivedTime),
		PeerSeenTime:         int(m.PeerSeenTime),
		DeliviryStatusEnumId: int(m.DeliviryStatusEnumId),
	}
	return r
}

//ToPB
func (m *PB_DirectMessage_Flat) ToPB() *PB_DirectMessage {
	r := &PB_DirectMessage{
		MessageId:            int64(m.MessageId),
		MessageKey:           m.MessageKey,
		RoomKey:              m.RoomKey,
		UserId:               int32(m.UserId),
		MessageFileId:        int64(m.MessageFileId),
		MessageTypeEnumId:    int32(m.MessageTypeEnumId),
		Text:                 m.Text,
		CreatedSe:            int32(m.CreatedSe),
		PeerReceivedTime:     int32(m.PeerReceivedTime),
		PeerSeenTime:         int32(m.PeerSeenTime),
		DeliviryStatusEnumId: int32(m.DeliviryStatusEnumId),
	}
	return r
}

//folding
var PB_DirectMessage__FOlD = &PB_DirectMessage{
	MessageId:            0,
	MessageKey:           "",
	RoomKey:              "",
	UserId:               0,
	MessageFileId:        0,
	MessageTypeEnumId:    0,
	Text:                 "",
	CreatedSe:            0,
	PeerReceivedTime:     0,
	PeerSeenTime:         0,
	DeliviryStatusEnumId: 0,
}

type PB_DirectOffline_Flat struct {
	DirectOfflineId int
	ToUserId        int
	ChatKey         string
	MessageId       int
	MessageFileId   int
	PBClass         string
	DataPB          []byte
	DataJson        string
	DataTemp        string
	AtTimeMs        int
}

//ToPB
func (m *PB_DirectOffline) ToFlat() *PB_DirectOffline_Flat {
	r := &PB_DirectOffline_Flat{
		DirectOfflineId: int(m.DirectOfflineId),
		ToUserId:        int(m.ToUserId),
		ChatKey:         m.ChatKey,
		MessageId:       int(m.MessageId),
		MessageFileId:   int(m.MessageFileId),
		PBClass:         m.PBClass,
		DataPB:          []byte(m.DataPB),
		DataJson:        m.DataJson,
		DataTemp:        m.DataTemp,
		AtTimeMs:        int(m.AtTimeMs),
	}
	return r
}

//ToPB
func (m *PB_DirectOffline_Flat) ToPB() *PB_DirectOffline {
	r := &PB_DirectOffline{
		DirectOfflineId: int64(m.DirectOfflineId),
		ToUserId:        int32(m.ToUserId),
		ChatKey:         m.ChatKey,
		MessageId:       int64(m.MessageId),
		MessageFileId:   int64(m.MessageFileId),
		PBClass:         m.PBClass,
		DataPB:          m.DataPB,
		DataJson:        m.DataJson,
		DataTemp:        m.DataTemp,
		AtTimeMs:        int64(m.AtTimeMs),
	}
	return r
}

//folding
var PB_DirectOffline__FOlD = &PB_DirectOffline{
	DirectOfflineId: 0,
	ToUserId:        0,
	ChatKey:         "",
	MessageId:       0,
	MessageFileId:   0,
	PBClass:         "",
	DataPB:          []byte{},
	DataJson:        "",
	DataTemp:        "",
	AtTimeMs:        0,
}

type PB_DirectToMessage_Flat struct {
	Id           int
	ChatKey      string
	MessageId    int
	SourceEnumId int
}

//ToPB
func (m *PB_DirectToMessage) ToFlat() *PB_DirectToMessage_Flat {
	r := &PB_DirectToMessage_Flat{
		Id:           int(m.Id),
		ChatKey:      m.ChatKey,
		MessageId:    int(m.MessageId),
		SourceEnumId: int(m.SourceEnumId),
	}
	return r
}

//ToPB
func (m *PB_DirectToMessage_Flat) ToPB() *PB_DirectToMessage {
	r := &PB_DirectToMessage{
		Id:           int64(m.Id),
		ChatKey:      m.ChatKey,
		MessageId:    int64(m.MessageId),
		SourceEnumId: int32(m.SourceEnumId),
	}
	return r
}

//folding
var PB_DirectToMessage__FOlD = &PB_DirectToMessage{
	Id:           0,
	ChatKey:      "",
	MessageId:    0,
	SourceEnumId: 0,
}

type PB_Feed_Flat struct {
	UserId int
	PostId int
	FeedId int
}

//ToPB
func (m *PB_Feed) ToFlat() *PB_Feed_Flat {
	r := &PB_Feed_Flat{
		UserId: int(m.UserId),
		PostId: int(m.PostId),
		FeedId: int(m.FeedId),
	}
	return r
}

//ToPB
func (m *PB_Feed_Flat) ToPB() *PB_Feed {
	r := &PB_Feed{
		UserId: int64(m.UserId),
		PostId: int64(m.PostId),
		FeedId: int32(m.FeedId),
	}
	return r
}

//folding
var PB_Feed__FOlD = &PB_Feed{
	UserId: 0,
	PostId: 0,
	FeedId: 0,
}

type PB_FollowingList_Flat struct {
	Id          int
	UserId      int
	ListType    int
	Name        string
	Count       int
	IsAuto      int
	IsPimiry    int
	CreatedTime int
}

//ToPB
func (m *PB_FollowingList) ToFlat() *PB_FollowingList_Flat {
	r := &PB_FollowingList_Flat{
		Id:          int(m.Id),
		UserId:      int(m.UserId),
		ListType:    int(m.ListType),
		Name:        m.Name,
		Count:       int(m.Count),
		IsAuto:      int(m.IsAuto),
		IsPimiry:    int(m.IsPimiry),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_FollowingList_Flat) ToPB() *PB_FollowingList {
	r := &PB_FollowingList{
		Id:          int32(m.Id),
		UserId:      int32(m.UserId),
		ListType:    int32(m.ListType),
		Name:        m.Name,
		Count:       int32(m.Count),
		IsAuto:      int32(m.IsAuto),
		IsPimiry:    int32(m.IsPimiry),
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_FollowingList__FOlD = &PB_FollowingList{
	Id:          0,
	UserId:      0,
	ListType:    0,
	Name:        "",
	Count:       0,
	IsAuto:      0,
	IsPimiry:    0,
	CreatedTime: 0,
}

type PB_FollowingListMember_Flat struct {
	Id             int
	ListId         int
	UserId         int
	FollowedUserId int
	FollowType     int
	UpdatedTimeMs  int
}

//ToPB
func (m *PB_FollowingListMember) ToFlat() *PB_FollowingListMember_Flat {
	r := &PB_FollowingListMember_Flat{
		Id:             int(m.Id),
		ListId:         int(m.ListId),
		UserId:         int(m.UserId),
		FollowedUserId: int(m.FollowedUserId),
		FollowType:     int(m.FollowType),
		UpdatedTimeMs:  int(m.UpdatedTimeMs),
	}
	return r
}

//ToPB
func (m *PB_FollowingListMember_Flat) ToPB() *PB_FollowingListMember {
	r := &PB_FollowingListMember{
		Id:             int64(m.Id),
		ListId:         int32(m.ListId),
		UserId:         int32(m.UserId),
		FollowedUserId: int32(m.FollowedUserId),
		FollowType:     int32(m.FollowType),
		UpdatedTimeMs:  int64(m.UpdatedTimeMs),
	}
	return r
}

//folding
var PB_FollowingListMember__FOlD = &PB_FollowingListMember{
	Id:             0,
	ListId:         0,
	UserId:         0,
	FollowedUserId: 0,
	FollowType:     0,
	UpdatedTimeMs:  0,
}

type PB_FollowingListMemberHistory_Flat struct {
	Id             int
	ListId         int
	UserId         int
	FollowedUserId int
	FollowType     int
	UpdatedTimeMs  int
	FollowId       int
}

//ToPB
func (m *PB_FollowingListMemberHistory) ToFlat() *PB_FollowingListMemberHistory_Flat {
	r := &PB_FollowingListMemberHistory_Flat{
		Id:             int(m.Id),
		ListId:         int(m.ListId),
		UserId:         int(m.UserId),
		FollowedUserId: int(m.FollowedUserId),
		FollowType:     int(m.FollowType),
		UpdatedTimeMs:  int(m.UpdatedTimeMs),
		FollowId:       int(m.FollowId),
	}
	return r
}

//ToPB
func (m *PB_FollowingListMemberHistory_Flat) ToPB() *PB_FollowingListMemberHistory {
	r := &PB_FollowingListMemberHistory{
		Id:             int64(m.Id),
		ListId:         int32(m.ListId),
		UserId:         int32(m.UserId),
		FollowedUserId: int32(m.FollowedUserId),
		FollowType:     int32(m.FollowType),
		UpdatedTimeMs:  int64(m.UpdatedTimeMs),
		FollowId:       int32(m.FollowId),
	}
	return r
}

//folding
var PB_FollowingListMemberHistory__FOlD = &PB_FollowingListMemberHistory{
	Id:             0,
	ListId:         0,
	UserId:         0,
	FollowedUserId: 0,
	FollowType:     0,
	UpdatedTimeMs:  0,
	FollowId:       0,
}

type PB_GeneralLog_Flat struct {
	Id        int
	ToUserId  int
	TargetId  int
	LogTypeId int
	ExtraPB   []byte
	ExtraJson string
	CreatedMs int
}

//ToPB
func (m *PB_GeneralLog) ToFlat() *PB_GeneralLog_Flat {
	r := &PB_GeneralLog_Flat{
		Id:        int(m.Id),
		ToUserId:  int(m.ToUserId),
		TargetId:  int(m.TargetId),
		LogTypeId: int(m.LogTypeId),
		ExtraPB:   []byte(m.ExtraPB),
		ExtraJson: m.ExtraJson,
		CreatedMs: int(m.CreatedMs),
	}
	return r
}

//ToPB
func (m *PB_GeneralLog_Flat) ToPB() *PB_GeneralLog {
	r := &PB_GeneralLog{
		Id:        int64(m.Id),
		ToUserId:  int32(m.ToUserId),
		TargetId:  int32(m.TargetId),
		LogTypeId: int32(m.LogTypeId),
		ExtraPB:   m.ExtraPB,
		ExtraJson: m.ExtraJson,
		CreatedMs: int64(m.CreatedMs),
	}
	return r
}

//folding
var PB_GeneralLog__FOlD = &PB_GeneralLog{
	Id:        0,
	ToUserId:  0,
	TargetId:  0,
	LogTypeId: 0,
	ExtraPB:   []byte{},
	ExtraJson: "",
	CreatedMs: 0,
}

type PB_Group_Flat struct {
	GroupId          int
	GroupName        string
	MembersCount     int
	GroupPrivacyEnum int
	CreatorUserId    int
	CreatedTime      int
	UpdatedMs        int
	CurrentSeq       int
}

//ToPB
func (m *PB_Group) ToFlat() *PB_Group_Flat {
	r := &PB_Group_Flat{
		GroupId:          int(m.GroupId),
		GroupName:        m.GroupName,
		MembersCount:     int(m.MembersCount),
		GroupPrivacyEnum: int(m.GroupPrivacyEnum),
		CreatorUserId:    int(m.CreatorUserId),
		CreatedTime:      int(m.CreatedTime),
		UpdatedMs:        int(m.UpdatedMs),
		CurrentSeq:       int(m.CurrentSeq),
	}
	return r
}

//ToPB
func (m *PB_Group_Flat) ToPB() *PB_Group {
	r := &PB_Group{
		GroupId:          int64(m.GroupId),
		GroupName:        m.GroupName,
		MembersCount:     int32(m.MembersCount),
		GroupPrivacyEnum: int32(m.GroupPrivacyEnum),
		CreatorUserId:    int32(m.CreatorUserId),
		CreatedTime:      int32(m.CreatedTime),
		UpdatedMs:        int64(m.UpdatedMs),
		CurrentSeq:       int32(m.CurrentSeq),
	}
	return r
}

//folding
var PB_Group__FOlD = &PB_Group{
	GroupId:          0,
	GroupName:        "",
	MembersCount:     0,
	GroupPrivacyEnum: 0,
	CreatorUserId:    0,
	CreatedTime:      0,
	UpdatedMs:        0,
	CurrentSeq:       0,
}

type PB_GroupMember_Flat struct {
	Id              int
	GroupId         int
	GroupKey        string
	UserId          int
	ByUserId        int
	GroupRoleEnumId int
	CreatedTime     int
}

//ToPB
func (m *PB_GroupMember) ToFlat() *PB_GroupMember_Flat {
	r := &PB_GroupMember_Flat{
		Id:              int(m.Id),
		GroupId:         int(m.GroupId),
		GroupKey:        m.GroupKey,
		UserId:          int(m.UserId),
		ByUserId:        int(m.ByUserId),
		GroupRoleEnumId: int(m.GroupRoleEnumId),
		CreatedTime:     int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_GroupMember_Flat) ToPB() *PB_GroupMember {
	r := &PB_GroupMember{
		Id:              int64(m.Id),
		GroupId:         int64(m.GroupId),
		GroupKey:        m.GroupKey,
		UserId:          int32(m.UserId),
		ByUserId:        int32(m.ByUserId),
		GroupRoleEnumId: int32(m.GroupRoleEnumId),
		CreatedTime:     int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_GroupMember__FOlD = &PB_GroupMember{
	Id:              0,
	GroupId:         0,
	GroupKey:        "",
	UserId:          0,
	ByUserId:        0,
	GroupRoleEnumId: 0,
	CreatedTime:     0,
}

type PB_GroupMessage_Flat struct {
	MessageId          int
	RoomKey            string
	UserId             int
	MessageFileId      int
	MessageTypeEnum    int
	Text               string
	CreatedMs          int
	DeliveryStatusEnum int
}

//ToPB
func (m *PB_GroupMessage) ToFlat() *PB_GroupMessage_Flat {
	r := &PB_GroupMessage_Flat{
		MessageId:          int(m.MessageId),
		RoomKey:            m.RoomKey,
		UserId:             int(m.UserId),
		MessageFileId:      int(m.MessageFileId),
		MessageTypeEnum:    int(m.MessageTypeEnum),
		Text:               m.Text,
		CreatedMs:          int(m.CreatedMs),
		DeliveryStatusEnum: int(m.DeliveryStatusEnum),
	}
	return r
}

//ToPB
func (m *PB_GroupMessage_Flat) ToPB() *PB_GroupMessage {
	r := &PB_GroupMessage{
		MessageId:          int64(m.MessageId),
		RoomKey:            m.RoomKey,
		UserId:             int32(m.UserId),
		MessageFileId:      int64(m.MessageFileId),
		MessageTypeEnum:    int32(m.MessageTypeEnum),
		Text:               m.Text,
		CreatedMs:          int64(m.CreatedMs),
		DeliveryStatusEnum: int32(m.DeliveryStatusEnum),
	}
	return r
}

//folding
var PB_GroupMessage__FOlD = &PB_GroupMessage{
	MessageId:          0,
	RoomKey:            "",
	UserId:             0,
	MessageFileId:      0,
	MessageTypeEnum:    0,
	Text:               "",
	CreatedMs:          0,
	DeliveryStatusEnum: 0,
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

type PB_Media_Flat struct {
	MediaId       int
	UserId        int
	PostId        int
	AlbumId       int
	MediaTypeEnum int
	Width         int
	Height        int
	Size          int
	Duration      int
	Md5Hash       string
	Color         string
	CreatedTime   int
}

//ToPB
func (m *PB_Media) ToFlat() *PB_Media_Flat {
	r := &PB_Media_Flat{
		MediaId:       int(m.MediaId),
		UserId:        int(m.UserId),
		PostId:        int(m.PostId),
		AlbumId:       int(m.AlbumId),
		MediaTypeEnum: int(m.MediaTypeEnum),
		Width:         int(m.Width),
		Height:        int(m.Height),
		Size:          int(m.Size),
		Duration:      int(m.Duration),
		Md5Hash:       m.Md5Hash,
		Color:         m.Color,
		CreatedTime:   int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Media_Flat) ToPB() *PB_Media {
	r := &PB_Media{
		MediaId:       int64(m.MediaId),
		UserId:        int32(m.UserId),
		PostId:        int32(m.PostId),
		AlbumId:       int32(m.AlbumId),
		MediaTypeEnum: int32(m.MediaTypeEnum),
		Width:         int32(m.Width),
		Height:        int32(m.Height),
		Size:          int32(m.Size),
		Duration:      int32(m.Duration),
		Md5Hash:       m.Md5Hash,
		Color:         m.Color,
		CreatedTime:   int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Media__FOlD = &PB_Media{
	MediaId:       0,
	UserId:        0,
	PostId:        0,
	AlbumId:       0,
	MediaTypeEnum: 0,
	Width:         0,
	Height:        0,
	Size:          0,
	Duration:      0,
	Md5Hash:       "",
	Color:         "",
	CreatedTime:   0,
}

type PB_MessageFile_Flat struct {
	MessageFileId  int
	MessageFileKey string
	UserId         int
	Title          string
	Size           int
	FileTypeEnum   int
	Width          int
	Height         int
	Duration       int
	Extension      string
	Md5Hash        string
	CreatedTime    int
}

//ToPB
func (m *PB_MessageFile) ToFlat() *PB_MessageFile_Flat {
	r := &PB_MessageFile_Flat{
		MessageFileId:  int(m.MessageFileId),
		MessageFileKey: m.MessageFileKey,
		UserId:         int(m.UserId),
		Title:          m.Title,
		Size:           int(m.Size),
		FileTypeEnum:   int(m.FileTypeEnum),
		Width:          int(m.Width),
		Height:         int(m.Height),
		Duration:       int(m.Duration),
		Extension:      m.Extension,
		Md5Hash:        m.Md5Hash,
		CreatedTime:    int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_MessageFile_Flat) ToPB() *PB_MessageFile {
	r := &PB_MessageFile{
		MessageFileId:  int64(m.MessageFileId),
		MessageFileKey: m.MessageFileKey,
		UserId:         int32(m.UserId),
		Title:          m.Title,
		Size:           int32(m.Size),
		FileTypeEnum:   int32(m.FileTypeEnum),
		Width:          int32(m.Width),
		Height:         int32(m.Height),
		Duration:       int32(m.Duration),
		Extension:      m.Extension,
		Md5Hash:        m.Md5Hash,
		CreatedTime:    int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_MessageFile__FOlD = &PB_MessageFile{
	MessageFileId:  0,
	MessageFileKey: "",
	UserId:         0,
	Title:          "",
	Size:           0,
	FileTypeEnum:   0,
	Width:          0,
	Height:         0,
	Duration:       0,
	Extension:      "",
	Md5Hash:        "",
	CreatedTime:    0,
}

type PB_Notify_Flat struct {
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
	Seq           int
}

//ToPB
func (m *PB_Notify) ToFlat() *PB_Notify_Flat {
	r := &PB_Notify_Flat{
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
		Seq:           int(m.Seq),
	}
	return r
}

//ToPB
func (m *PB_Notify_Flat) ToPB() *PB_Notify {
	r := &PB_Notify{
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
		Seq:           int32(m.Seq),
	}
	return r
}

//folding
var PB_Notify__FOlD = &PB_Notify{
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
	Seq:           0,
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
	Id                    int
	PhoneDisplayName      string
	PhoneFamilyName       string
	PhoneNumber           string
	PhoneNormalizedNumber string
	PhoneContactRowId     int
	UserId                int
	DeviceUuidId          int
	CreatedTime           int
	UpdatedTime           int
}

//ToPB
func (m *PB_PhoneContact) ToFlat() *PB_PhoneContact_Flat {
	r := &PB_PhoneContact_Flat{
		Id:                    int(m.Id),
		PhoneDisplayName:      m.PhoneDisplayName,
		PhoneFamilyName:       m.PhoneFamilyName,
		PhoneNumber:           m.PhoneNumber,
		PhoneNormalizedNumber: m.PhoneNormalizedNumber,
		PhoneContactRowId:     int(m.PhoneContactRowId),
		UserId:                int(m.UserId),
		DeviceUuidId:          int(m.DeviceUuidId),
		CreatedTime:           int(m.CreatedTime),
		UpdatedTime:           int(m.UpdatedTime),
	}
	return r
}

//ToPB
func (m *PB_PhoneContact_Flat) ToPB() *PB_PhoneContact {
	r := &PB_PhoneContact{
		Id:                    int32(m.Id),
		PhoneDisplayName:      m.PhoneDisplayName,
		PhoneFamilyName:       m.PhoneFamilyName,
		PhoneNumber:           m.PhoneNumber,
		PhoneNormalizedNumber: m.PhoneNormalizedNumber,
		PhoneContactRowId:     int32(m.PhoneContactRowId),
		UserId:                int32(m.UserId),
		DeviceUuidId:          int32(m.DeviceUuidId),
		CreatedTime:           int32(m.CreatedTime),
		UpdatedTime:           int32(m.UpdatedTime),
	}
	return r
}

//folding
var PB_PhoneContact__FOlD = &PB_PhoneContact{
	Id:                    0,
	PhoneDisplayName:      "",
	PhoneFamilyName:       "",
	PhoneNumber:           "",
	PhoneNormalizedNumber: "",
	PhoneContactRowId:     0,
	UserId:                0,
	DeviceUuidId:          0,
	CreatedTime:           0,
	UpdatedTime:           0,
}

type PB_Post_Flat struct {
	PostId         int
	UserId         int
	PostTypeEnum   int
	MediaId        int
	Text           string
	RichText       string
	MediaCount     int
	SharedTo       int
	DisableComment int
	HasTag         int
	CommentsCount  int
	LikesCount     int
	ViewsCount     int
	EditedTime     int
	CreatedTime    int
	ReSharedPostId int
}

//ToPB
func (m *PB_Post) ToFlat() *PB_Post_Flat {
	r := &PB_Post_Flat{
		PostId:         int(m.PostId),
		UserId:         int(m.UserId),
		PostTypeEnum:   int(m.PostTypeEnum),
		MediaId:        int(m.MediaId),
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
	}
	return r
}

//ToPB
func (m *PB_Post_Flat) ToPB() *PB_Post {
	r := &PB_Post{
		PostId:         int64(m.PostId),
		UserId:         int32(m.UserId),
		PostTypeEnum:   int32(m.PostTypeEnum),
		MediaId:        int32(m.MediaId),
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
	}
	return r
}

//folding
var PB_Post__FOlD = &PB_Post{
	PostId:         0,
	UserId:         0,
	PostTypeEnum:   0,
	MediaId:        0,
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
}

type PB_RecommendUser_Flat struct {
	Id          int
	UserId      int
	TargetId    int
	Weight      float32
	CreatedTime int
}

//ToPB
func (m *PB_RecommendUser) ToFlat() *PB_RecommendUser_Flat {
	r := &PB_RecommendUser_Flat{
		Id:          int(m.Id),
		UserId:      int(m.UserId),
		TargetId:    int(m.TargetId),
		Weight:      float32(m.Weight),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_RecommendUser_Flat) ToPB() *PB_RecommendUser {
	r := &PB_RecommendUser{
		Id:          int32(m.Id),
		UserId:      int32(m.UserId),
		TargetId:    int32(m.TargetId),
		Weight:      m.Weight,
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_RecommendUser__FOlD = &PB_RecommendUser{
	Id:          0,
	UserId:      0,
	TargetId:    0,
	Weight:      0.0,
	CreatedTime: 0,
}

type PB_SearchClicked_Flat struct {
	Id          int
	Query       string
	ClickType   int
	TargetId    int
	UserId      int
	CreatedTime int
}

//ToPB
func (m *PB_SearchClicked) ToFlat() *PB_SearchClicked_Flat {
	r := &PB_SearchClicked_Flat{
		Id:          int(m.Id),
		Query:       m.Query,
		ClickType:   int(m.ClickType),
		TargetId:    int(m.TargetId),
		UserId:      int(m.UserId),
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_SearchClicked_Flat) ToPB() *PB_SearchClicked {
	r := &PB_SearchClicked{
		Id:          int64(m.Id),
		Query:       m.Query,
		ClickType:   int32(m.ClickType),
		TargetId:    int32(m.TargetId),
		UserId:      int32(m.UserId),
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_SearchClicked__FOlD = &PB_SearchClicked{
	Id:          0,
	Query:       "",
	ClickType:   0,
	TargetId:    0,
	UserId:      0,
	CreatedTime: 0,
}

type PB_Session_Flat struct {
	Id                    int
	UserId                int
	SessionUuid           string
	ClientUuid            string
	DeviceUuid            string
	LastActivityTime      int
	LastIpAddress         string
	LastWifiMacAddress    string
	LastNetworkType       string
	LastNetworkTypeEnumId int
	AppVersion            int
	UpdatedTime           int
	CreatedTime           int
}

//ToPB
func (m *PB_Session) ToFlat() *PB_Session_Flat {
	r := &PB_Session_Flat{
		Id:                    int(m.Id),
		UserId:                int(m.UserId),
		SessionUuid:           m.SessionUuid,
		ClientUuid:            m.ClientUuid,
		DeviceUuid:            m.DeviceUuid,
		LastActivityTime:      int(m.LastActivityTime),
		LastIpAddress:         m.LastIpAddress,
		LastWifiMacAddress:    m.LastWifiMacAddress,
		LastNetworkType:       m.LastNetworkType,
		LastNetworkTypeEnumId: int(m.LastNetworkTypeEnumId),
		AppVersion:            int(m.AppVersion),
		UpdatedTime:           int(m.UpdatedTime),
		CreatedTime:           int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_Session_Flat) ToPB() *PB_Session {
	r := &PB_Session{
		Id:                    int64(m.Id),
		UserId:                int32(m.UserId),
		SessionUuid:           m.SessionUuid,
		ClientUuid:            m.ClientUuid,
		DeviceUuid:            m.DeviceUuid,
		LastActivityTime:      int32(m.LastActivityTime),
		LastIpAddress:         m.LastIpAddress,
		LastWifiMacAddress:    m.LastWifiMacAddress,
		LastNetworkType:       m.LastNetworkType,
		LastNetworkTypeEnumId: int32(m.LastNetworkTypeEnumId),
		AppVersion:            int32(m.AppVersion),
		UpdatedTime:           int32(m.UpdatedTime),
		CreatedTime:           int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_Session__FOlD = &PB_Session{
	Id:                    0,
	UserId:                0,
	SessionUuid:           "",
	ClientUuid:            "",
	DeviceUuid:            "",
	LastActivityTime:      0,
	LastIpAddress:         "",
	LastWifiMacAddress:    "",
	LastNetworkType:       "",
	LastNetworkTypeEnumId: 0,
	AppVersion:            0,
	UpdatedTime:           0,
	CreatedTime:           0,
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

type PB_TagsPost_Flat struct {
	Id           int
	TagId        int
	PostId       int
	PostTypeEnum int
	CreatedTime  int
}

//ToPB
func (m *PB_TagsPost) ToFlat() *PB_TagsPost_Flat {
	r := &PB_TagsPost_Flat{
		Id:           int(m.Id),
		TagId:        int(m.TagId),
		PostId:       int(m.PostId),
		PostTypeEnum: int(m.PostTypeEnum),
		CreatedTime:  int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_TagsPost_Flat) ToPB() *PB_TagsPost {
	r := &PB_TagsPost{
		Id:           int64(m.Id),
		TagId:        int32(m.TagId),
		PostId:       int32(m.PostId),
		PostTypeEnum: int32(m.PostTypeEnum),
		CreatedTime:  int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_TagsPost__FOlD = &PB_TagsPost{
	Id:           0,
	TagId:        0,
	PostId:       0,
	PostTypeEnum: 0,
	CreatedTime:  0,
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
	UserId               int
	UserName             string
	UserNameLower        string
	FirstName            string
	LastName             string
	UserTypeEnum         int
	UserLevelEnum        int
	AvatarId             int
	ProfilePrivacyEnum   int
	Phone                int
	About                string
	Email                string
	PasswordHash         string
	PasswordSalt         string
	FollowersCount       int
	FollowingCount       int
	PostsCount           int
	MediaCount           int
	LikesCount           int
	ResharedCount        int
	LastActionTime       int
	LastPostTime         int
	PrimaryFollowingList int
	CreatedSe            int
	UpdatedMs            int
	OnlinePrivacyEnum    int
	LastActivityTime     int
	Phone2               string
}

//ToPB
func (m *PB_User) ToFlat() *PB_User_Flat {
	r := &PB_User_Flat{
		UserId:               int(m.UserId),
		UserName:             m.UserName,
		UserNameLower:        m.UserNameLower,
		FirstName:            m.FirstName,
		LastName:             m.LastName,
		UserTypeEnum:         int(m.UserTypeEnum),
		UserLevelEnum:        int(m.UserLevelEnum),
		AvatarId:             int(m.AvatarId),
		ProfilePrivacyEnum:   int(m.ProfilePrivacyEnum),
		Phone:                int(m.Phone),
		About:                m.About,
		Email:                m.Email,
		PasswordHash:         m.PasswordHash,
		PasswordSalt:         m.PasswordSalt,
		FollowersCount:       int(m.FollowersCount),
		FollowingCount:       int(m.FollowingCount),
		PostsCount:           int(m.PostsCount),
		MediaCount:           int(m.MediaCount),
		LikesCount:           int(m.LikesCount),
		ResharedCount:        int(m.ResharedCount),
		LastActionTime:       int(m.LastActionTime),
		LastPostTime:         int(m.LastPostTime),
		PrimaryFollowingList: int(m.PrimaryFollowingList),
		CreatedSe:            int(m.CreatedSe),
		UpdatedMs:            int(m.UpdatedMs),
		OnlinePrivacyEnum:    int(m.OnlinePrivacyEnum),
		LastActivityTime:     int(m.LastActivityTime),
		Phone2:               m.Phone2,
	}
	return r
}

//ToPB
func (m *PB_User_Flat) ToPB() *PB_User {
	r := &PB_User{
		UserId:               int32(m.UserId),
		UserName:             m.UserName,
		UserNameLower:        m.UserNameLower,
		FirstName:            m.FirstName,
		LastName:             m.LastName,
		UserTypeEnum:         int32(m.UserTypeEnum),
		UserLevelEnum:        int32(m.UserLevelEnum),
		AvatarId:             int64(m.AvatarId),
		ProfilePrivacyEnum:   int32(m.ProfilePrivacyEnum),
		Phone:                int64(m.Phone),
		About:                m.About,
		Email:                m.Email,
		PasswordHash:         m.PasswordHash,
		PasswordSalt:         m.PasswordSalt,
		FollowersCount:       int32(m.FollowersCount),
		FollowingCount:       int32(m.FollowingCount),
		PostsCount:           int32(m.PostsCount),
		MediaCount:           int32(m.MediaCount),
		LikesCount:           int32(m.LikesCount),
		ResharedCount:        int32(m.ResharedCount),
		LastActionTime:       int32(m.LastActionTime),
		LastPostTime:         int32(m.LastPostTime),
		PrimaryFollowingList: int32(m.PrimaryFollowingList),
		CreatedSe:            int32(m.CreatedSe),
		UpdatedMs:            int64(m.UpdatedMs),
		OnlinePrivacyEnum:    int32(m.OnlinePrivacyEnum),
		LastActivityTime:     int32(m.LastActivityTime),
		Phone2:               m.Phone2,
	}
	return r
}

//folding
var PB_User__FOlD = &PB_User{
	UserId:               0,
	UserName:             "",
	UserNameLower:        "",
	FirstName:            "",
	LastName:             "",
	UserTypeEnum:         0,
	UserLevelEnum:        0,
	AvatarId:             0,
	ProfilePrivacyEnum:   0,
	Phone:                0,
	About:                "",
	Email:                "",
	PasswordHash:         "",
	PasswordSalt:         "",
	FollowersCount:       0,
	FollowingCount:       0,
	PostsCount:           0,
	MediaCount:           0,
	LikesCount:           0,
	ResharedCount:        0,
	LastActionTime:       0,
	LastPostTime:         0,
	PrimaryFollowingList: 0,
	CreatedSe:            0,
	UpdatedMs:            0,
	OnlinePrivacyEnum:    0,
	LastActivityTime:     0,
	Phone2:               "",
}

type PB_UserMetaInfo_Flat struct {
	Id                  int
	UserId              int
	IsNotificationDirty int
	LastUserRecGen      int
}

//ToPB
func (m *PB_UserMetaInfo) ToFlat() *PB_UserMetaInfo_Flat {
	r := &PB_UserMetaInfo_Flat{
		Id:                  int(m.Id),
		UserId:              int(m.UserId),
		IsNotificationDirty: int(m.IsNotificationDirty),
		LastUserRecGen:      int(m.LastUserRecGen),
	}
	return r
}

//ToPB
func (m *PB_UserMetaInfo_Flat) ToPB() *PB_UserMetaInfo {
	r := &PB_UserMetaInfo{
		Id:                  int32(m.Id),
		UserId:              int32(m.UserId),
		IsNotificationDirty: int32(m.IsNotificationDirty),
		LastUserRecGen:      int32(m.LastUserRecGen),
	}
	return r
}

//folding
var PB_UserMetaInfo__FOlD = &PB_UserMetaInfo{
	Id:                  0,
	UserId:              0,
	IsNotificationDirty: 0,
	LastUserRecGen:      0,
}

type PB_UserPassword_Flat struct {
	UserId      int
	Password    string
	CreatedTime int
}

//ToPB
func (m *PB_UserPassword) ToFlat() *PB_UserPassword_Flat {
	r := &PB_UserPassword_Flat{
		UserId:      int(m.UserId),
		Password:    m.Password,
		CreatedTime: int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_UserPassword_Flat) ToPB() *PB_UserPassword {
	r := &PB_UserPassword{
		UserId:      int32(m.UserId),
		Password:    m.Password,
		CreatedTime: int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_UserPassword__FOlD = &PB_UserPassword{
	UserId:      0,
	Password:    "",
	CreatedTime: 0,
}

type PB_PostView_Flat struct {
	PostId         int
	UserId         int
	PostTypeEnum   PostTypeEnum
	Text           string
	RichText       string
	MediaCount     int
	SharedTo       int
	DisableComment int
	HasTag         int
	CommentsCount  int
	LikesCount     int
	ViewsCount     int
	EditedTime     int
	CreatedTime    int
	ReSharedPostId int
	DidILiked      bool
	DidIReShared   bool
	SenderUser     PB_UserView
	ReSharedUser   PB_UserView
	Media          PB_MediaView
	MediaList      []PB_MediaView
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

type PB_MediaView_Flat struct {
	MediaId       int
	UserId        int
	PostId        int
	AlbumId       int
	MediaTypeEnum int
	Width         int
	Height        int
	Size          int
	Duration      int
	Color         string
	CreatedTime   int
}

//ToPB
func (m *PB_MediaView) ToFlat() *PB_MediaView_Flat {
	r := &PB_MediaView_Flat{
		MediaId:       int(m.MediaId),
		UserId:        int(m.UserId),
		PostId:        int(m.PostId),
		AlbumId:       int(m.AlbumId),
		MediaTypeEnum: int(m.MediaTypeEnum),
		Width:         int(m.Width),
		Height:        int(m.Height),
		Size:          int(m.Size),
		Duration:      int(m.Duration),
		Color:         m.Color,
		CreatedTime:   int(m.CreatedTime),
	}
	return r
}

//ToPB
func (m *PB_MediaView_Flat) ToPB() *PB_MediaView {
	r := &PB_MediaView{
		MediaId:       int64(m.MediaId),
		UserId:        int32(m.UserId),
		PostId:        int32(m.PostId),
		AlbumId:       int32(m.AlbumId),
		MediaTypeEnum: int32(m.MediaTypeEnum),
		Width:         int32(m.Width),
		Height:        int32(m.Height),
		Size:          int32(m.Size),
		Duration:      int32(m.Duration),
		Color:         m.Color,
		CreatedTime:   int32(m.CreatedTime),
	}
	return r
}

//folding
var PB_MediaView__FOlD = &PB_MediaView{
	MediaId:       0,
	UserId:        0,
	PostId:        0,
	AlbumId:       0,
	MediaTypeEnum: 0,
	Width:         0,
	Height:        0,
	Size:          0,
	Duration:      0,
	Color:         "",
	CreatedTime:   0,
}

type PB_ActionView_Flat struct {
	ActionId          int
	ActorUserId       int
	ActionTypeEnum    int
	PeerUserId        int
	PostId            int
	CommentId         int
	Murmur64Hash      int
	CreatedTime       int
	ActorUser         PB_UserView
	Post              PB_PostView
	Comment           PB_CommentView
	FollowedUser      PB_UserView
	ContentOwenerUser PB_UserView
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
	ActorUser     PB_UserView
	Post          PB_PostView
	Comment       PB_CommentView
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
	CommentId   int
	UserId      int
	PostId      int
	Text        string
	LikesCount  int
	CreatedTime int
	SenderUser  PB_UserView
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
		PostId:      int32(m.PostId),
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

type PB_UserView_Flat struct {
	UserId               int
	UserName             string
	FirstName            string
	LastName             string
	UserTypeEnum         UserTypeEnum
	UserLevelEnum        UserLevelEnum
	AvatarId             int
	ProfilePrivacyEnum   int
	Phone                int
	About                string
	FollowersCount       int
	FollowingCount       int
	PostsCount           int
	MediaCount           int
	LikesCount           int
	ResharedCount        int
	UserOnlineStatusEnum UserOnlineStatusEnum
	LastActiveTime       int
	MyFollwing           FollowingEnum
}

//ToPB
func (m *PB_UserView) ToFlat() *PB_UserView_Flat {
	r := &PB_UserView_Flat{
		UserId:    int(m.UserId),
		UserName:  m.UserName,
		FirstName: m.FirstName,
		LastName:  m.LastName,

		AvatarId:           int(m.AvatarId),
		ProfilePrivacyEnum: int(m.ProfilePrivacyEnum),
		Phone:              int(m.Phone),
		About:              m.About,
		FollowersCount:     int(m.FollowersCount),
		FollowingCount:     int(m.FollowingCount),
		PostsCount:         int(m.PostsCount),
		MediaCount:         int(m.MediaCount),
		LikesCount:         int(m.LikesCount),
		ResharedCount:      int(m.ResharedCount),

		LastActiveTime: int(m.LastActiveTime),
	}
	return r
}

//ToPB
func (m *PB_UserView_Flat) ToPB() *PB_UserView {
	r := &PB_UserView{
		UserId:    int32(m.UserId),
		UserName:  m.UserName,
		FirstName: m.FirstName,
		LastName:  m.LastName,

		AvatarId:           int64(m.AvatarId),
		ProfilePrivacyEnum: int32(m.ProfilePrivacyEnum),
		Phone:              int64(m.Phone),
		About:              m.About,
		FollowersCount:     int32(m.FollowersCount),
		FollowingCount:     int32(m.FollowingCount),
		PostsCount:         int32(m.PostsCount),
		MediaCount:         int32(m.MediaCount),
		LikesCount:         int32(m.LikesCount),
		ResharedCount:      int32(m.ResharedCount),

		LastActiveTime: int32(m.LastActiveTime),
	}
	return r
}

//folding
var PB_UserView__FOlD = &PB_UserView{
	UserId:    0,
	UserName:  "",
	FirstName: "",
	LastName:  "",

	AvatarId:           0,
	ProfilePrivacyEnum: 0,
	Phone:              0,
	About:              "",
	FollowersCount:     0,
	FollowingCount:     0,
	PostsCount:         0,
	MediaCount:         0,
	LikesCount:         0,
	ResharedCount:      0,

	LastActiveTime: 0,
}

type PB_TopProfileView_Flat struct {
	User PB_UserView
}

//ToPB
func (m *PB_TopProfileView) ToFlat() *PB_TopProfileView_Flat {
	r := &PB_TopProfileView_Flat{}
	return r
}

//ToPB
func (m *PB_TopProfileView_Flat) ToPB() *PB_TopProfileView {
	r := &PB_TopProfileView{}
	return r
}

//folding
var PB_TopProfileView__FOlD = &PB_TopProfileView{}

type PB_ChatView_Flat struct {
	ChatKey              string
	RoomKey              string
	RoomTypeEnumId       int
	UserId               int
	PeerUserId           int
	GroupId              int
	CreatedSe            int
	UpdatedMs            int
	LastMessageId        int
	LastDeletedMessageId int
	LastSeenMessageId    int
	LastSeqSeen          int
	LastSeqDelete        int
	CurrentSeq           int
	UserView             PB_UserView3
	SharedMediaCount     int
	UnseenCount          int
	FirstUnreadMessage   PB_MessageView
	LastMessage          PB_MessageView
}

//ToPB
func (m *PB_ChatView) ToFlat() *PB_ChatView_Flat {
	r := &PB_ChatView_Flat{
		ChatKey:              m.ChatKey,
		RoomKey:              m.RoomKey,
		RoomTypeEnumId:       int(m.RoomTypeEnumId),
		UserId:               int(m.UserId),
		PeerUserId:           int(m.PeerUserId),
		GroupId:              int(m.GroupId),
		CreatedSe:            int(m.CreatedSe),
		UpdatedMs:            int(m.UpdatedMs),
		LastMessageId:        int(m.LastMessageId),
		LastDeletedMessageId: int(m.LastDeletedMessageId),
		LastSeenMessageId:    int(m.LastSeenMessageId),
		LastSeqSeen:          int(m.LastSeqSeen),
		LastSeqDelete:        int(m.LastSeqDelete),
		CurrentSeq:           int(m.CurrentSeq),

		SharedMediaCount: int(m.SharedMediaCount),
		UnseenCount:      int(m.UnseenCount),
	}
	return r
}

//ToPB
func (m *PB_ChatView_Flat) ToPB() *PB_ChatView {
	r := &PB_ChatView{
		ChatKey:              m.ChatKey,
		RoomKey:              m.RoomKey,
		RoomTypeEnumId:       int32(m.RoomTypeEnumId),
		UserId:               int32(m.UserId),
		PeerUserId:           int32(m.PeerUserId),
		GroupId:              int64(m.GroupId),
		CreatedSe:            int32(m.CreatedSe),
		UpdatedMs:            int64(m.UpdatedMs),
		LastMessageId:        int64(m.LastMessageId),
		LastDeletedMessageId: int64(m.LastDeletedMessageId),
		LastSeenMessageId:    int64(m.LastSeenMessageId),
		LastSeqSeen:          int32(m.LastSeqSeen),
		LastSeqDelete:        int32(m.LastSeqDelete),
		CurrentSeq:           int32(m.CurrentSeq),

		SharedMediaCount: int32(m.SharedMediaCount),
		UnseenCount:      int32(m.UnseenCount),
	}
	return r
}

//folding
var PB_ChatView__FOlD = &PB_ChatView{
	ChatKey:              "",
	RoomKey:              "",
	RoomTypeEnumId:       0,
	UserId:               0,
	PeerUserId:           0,
	GroupId:              0,
	CreatedSe:            0,
	UpdatedMs:            0,
	LastMessageId:        0,
	LastDeletedMessageId: 0,
	LastSeenMessageId:    0,
	LastSeqSeen:          0,
	LastSeqDelete:        0,
	CurrentSeq:           0,

	SharedMediaCount: 0,
	UnseenCount:      0,
}

type PB_MessageView_Flat struct {
	MessageId            int
	MessageKey           string
	RoomKey              string
	UserId               int
	MessageFileId        int
	MessageTypeEnumId    int
	Text                 string
	CreatedSe            int
	PeerReceivedTime     int
	PeerSeenTime         int
	DeliviryStatusEnumId int
	ChatKey              string
	RoomTypeEnumId       int
	IsByMe               bool
	RemoteId             int
	MessageFileView      PB_MessageFileView
}

//ToPB
func (m *PB_MessageView) ToFlat() *PB_MessageView_Flat {
	r := &PB_MessageView_Flat{
		MessageId:            int(m.MessageId),
		MessageKey:           m.MessageKey,
		RoomKey:              m.RoomKey,
		UserId:               int(m.UserId),
		MessageFileId:        int(m.MessageFileId),
		MessageTypeEnumId:    int(m.MessageTypeEnumId),
		Text:                 m.Text,
		CreatedSe:            int(m.CreatedSe),
		PeerReceivedTime:     int(m.PeerReceivedTime),
		PeerSeenTime:         int(m.PeerSeenTime),
		DeliviryStatusEnumId: int(m.DeliviryStatusEnumId),
		ChatKey:              m.ChatKey,
		RoomTypeEnumId:       int(m.RoomTypeEnumId),
		IsByMe:               m.IsByMe,
		RemoteId:             int(m.RemoteId),
	}
	return r
}

//ToPB
func (m *PB_MessageView_Flat) ToPB() *PB_MessageView {
	r := &PB_MessageView{
		MessageId:            int64(m.MessageId),
		MessageKey:           m.MessageKey,
		RoomKey:              m.RoomKey,
		UserId:               int32(m.UserId),
		MessageFileId:        int64(m.MessageFileId),
		MessageTypeEnumId:    int32(m.MessageTypeEnumId),
		Text:                 m.Text,
		CreatedSe:            int32(m.CreatedSe),
		PeerReceivedTime:     int32(m.PeerReceivedTime),
		PeerSeenTime:         int32(m.PeerSeenTime),
		DeliviryStatusEnumId: int32(m.DeliviryStatusEnumId),
		ChatKey:              m.ChatKey,
		RoomTypeEnumId:       int32(m.RoomTypeEnumId),
		IsByMe:               m.IsByMe,
		RemoteId:             int64(m.RemoteId),
	}
	return r
}

//folding
var PB_MessageView__FOlD = &PB_MessageView{
	MessageId:            0,
	MessageKey:           "",
	RoomKey:              "",
	UserId:               0,
	MessageFileId:        0,
	MessageTypeEnumId:    0,
	Text:                 "",
	CreatedSe:            0,
	PeerReceivedTime:     0,
	PeerSeenTime:         0,
	DeliviryStatusEnumId: 0,
	ChatKey:              "",
	RoomTypeEnumId:       0,
	IsByMe:               false,
	RemoteId:             0,
}

type PB_MessageFileView_Flat struct {
	MessageFileId       int
	MessageFileKey      string
	OriginalUserId      int
	Name                string
	Size                int
	FileTypeEnumId      int
	Width               int
	Height              int
	Duration            int
	Extension           string
	HashMd5             string
	HashAccess          int
	CreatedSe           int
	ServerSrc           string
	ServerPath          string
	ServerThumbPath     string
	BucketId            string
	ServerId            int
	CanDel              int
	ServerThumbLocalSrc string
	RemoteMessageFileId int
	LocalSrc            string
	ThumbLocalSrc       string
	MessageFileStatusId string
}

//ToPB
func (m *PB_MessageFileView) ToFlat() *PB_MessageFileView_Flat {
	r := &PB_MessageFileView_Flat{
		MessageFileId:       int(m.MessageFileId),
		MessageFileKey:      m.MessageFileKey,
		OriginalUserId:      int(m.OriginalUserId),
		Name:                m.Name,
		Size:                int(m.Size),
		FileTypeEnumId:      int(m.FileTypeEnumId),
		Width:               int(m.Width),
		Height:              int(m.Height),
		Duration:            int(m.Duration),
		Extension:           m.Extension,
		HashMd5:             m.HashMd5,
		HashAccess:          int(m.HashAccess),
		CreatedSe:           int(m.CreatedSe),
		ServerSrc:           m.ServerSrc,
		ServerPath:          m.ServerPath,
		ServerThumbPath:     m.ServerThumbPath,
		BucketId:            m.BucketId,
		ServerId:            int(m.ServerId),
		CanDel:              int(m.CanDel),
		ServerThumbLocalSrc: m.ServerThumbLocalSrc,
		RemoteMessageFileId: int(m.RemoteMessageFileId),
		LocalSrc:            m.LocalSrc,
		ThumbLocalSrc:       m.ThumbLocalSrc,
		MessageFileStatusId: m.MessageFileStatusId,
	}
	return r
}

//ToPB
func (m *PB_MessageFileView_Flat) ToPB() *PB_MessageFileView {
	r := &PB_MessageFileView{
		MessageFileId:       int64(m.MessageFileId),
		MessageFileKey:      m.MessageFileKey,
		OriginalUserId:      int32(m.OriginalUserId),
		Name:                m.Name,
		Size:                int32(m.Size),
		FileTypeEnumId:      int32(m.FileTypeEnumId),
		Width:               int32(m.Width),
		Height:              int32(m.Height),
		Duration:            int32(m.Duration),
		Extension:           m.Extension,
		HashMd5:             m.HashMd5,
		HashAccess:          int64(m.HashAccess),
		CreatedSe:           int32(m.CreatedSe),
		ServerSrc:           m.ServerSrc,
		ServerPath:          m.ServerPath,
		ServerThumbPath:     m.ServerThumbPath,
		BucketId:            m.BucketId,
		ServerId:            int32(m.ServerId),
		CanDel:              int32(m.CanDel),
		ServerThumbLocalSrc: m.ServerThumbLocalSrc,
		RemoteMessageFileId: int64(m.RemoteMessageFileId),
		LocalSrc:            m.LocalSrc,
		ThumbLocalSrc:       m.ThumbLocalSrc,
		MessageFileStatusId: m.MessageFileStatusId,
	}
	return r
}

//folding
var PB_MessageFileView__FOlD = &PB_MessageFileView{
	MessageFileId:       0,
	MessageFileKey:      "",
	OriginalUserId:      0,
	Name:                "",
	Size:                0,
	FileTypeEnumId:      0,
	Width:               0,
	Height:              0,
	Duration:            0,
	Extension:           "",
	HashMd5:             "",
	HashAccess:          0,
	CreatedSe:           0,
	ServerSrc:           "",
	ServerPath:          "",
	ServerThumbPath:     "",
	BucketId:            "",
	ServerId:            0,
	CanDel:              0,
	ServerThumbLocalSrc: "",
	RemoteMessageFileId: 0,
	LocalSrc:            "",
	ThumbLocalSrc:       "",
	MessageFileStatusId: "",
}

type PB_UserView3_Flat struct {
	UserId           int
	UserName         string
	FirstName        string
	LastName         string
	About            string
	FullName         string
	AvatarUrl        string
	PrivacyProfile   int
	IsDeleted        int
	FollowersCount   int
	FollowingCount   int
	PostsCount       int
	UpdatedTime      int
	AppVersion       int
	LastActivityTime int
	FollowingType    int
}

//ToPB
func (m *PB_UserView3) ToFlat() *PB_UserView3_Flat {
	r := &PB_UserView3_Flat{
		UserId:           int(m.UserId),
		UserName:         m.UserName,
		FirstName:        m.FirstName,
		LastName:         m.LastName,
		About:            m.About,
		FullName:         m.FullName,
		AvatarUrl:        m.AvatarUrl,
		PrivacyProfile:   int(m.PrivacyProfile),
		IsDeleted:        int(m.IsDeleted),
		FollowersCount:   int(m.FollowersCount),
		FollowingCount:   int(m.FollowingCount),
		PostsCount:       int(m.PostsCount),
		UpdatedTime:      int(m.UpdatedTime),
		AppVersion:       int(m.AppVersion),
		LastActivityTime: int(m.LastActivityTime),
		FollowingType:    int(m.FollowingType),
	}
	return r
}

//ToPB
func (m *PB_UserView3_Flat) ToPB() *PB_UserView3 {
	r := &PB_UserView3{
		UserId:           int32(m.UserId),
		UserName:         m.UserName,
		FirstName:        m.FirstName,
		LastName:         m.LastName,
		About:            m.About,
		FullName:         m.FullName,
		AvatarUrl:        m.AvatarUrl,
		PrivacyProfile:   int32(m.PrivacyProfile),
		IsDeleted:        int32(m.IsDeleted),
		FollowersCount:   int32(m.FollowersCount),
		FollowingCount:   int32(m.FollowingCount),
		PostsCount:       int32(m.PostsCount),
		UpdatedTime:      int32(m.UpdatedTime),
		AppVersion:       int32(m.AppVersion),
		LastActivityTime: int32(m.LastActivityTime),
		FollowingType:    int32(m.FollowingType),
	}
	return r
}

//folding
var PB_UserView3__FOlD = &PB_UserView3{
	UserId:           0,
	UserName:         "",
	FirstName:        "",
	LastName:         "",
	About:            "",
	FullName:         "",
	AvatarUrl:        "",
	PrivacyProfile:   0,
	IsDeleted:        0,
	FollowersCount:   0,
	FollowingCount:   0,
	PostsCount:       0,
	UpdatedTime:      0,
	AppVersion:       0,
	LastActivityTime: 0,
	FollowingType:    0,
}

/*
///// to_flat ///

func(m *GeoLocation)ToFlat() *GeoLocation_Flat {
r := &GeoLocation_Flat{
    Lat:  float64(m.Lat) ,
    Lon:  float64(m.Lon) ,
}
return r
}

func(m *RoomMessageLog)ToFlat() *RoomMessageLog_Flat {
r := &RoomMessageLog_Flat{

    TargetUserId:  int(m.TargetUserId) ,
    ByUserId:  int(m.ByUserId) ,
}
return r
}

func(m *RoomMessageForwardFrom)ToFlat() *RoomMessageForwardFrom_Flat {
r := &RoomMessageForwardFrom_Flat{
    RoomId:  int(m.RoomId) ,
    MessageId:  int(m.MessageId) ,
    RoomTypeEnum:  int(m.RoomTypeEnum) ,
}
return r
}

func(m *RoomDraft)ToFlat() *RoomDraft_Flat {
r := &RoomDraft_Flat{
    Message:  m.Message ,
    ReplyTo:  int(m.ReplyTo) ,
}
return r
}

func(m *ChatRoom)ToFlat() *ChatRoom_Flat {
r := &ChatRoom_Flat{
}
return r
}

func(m *Pagination)ToFlat() *Pagination_Flat {
r := &Pagination_Flat{
    Offset:  int(m.Offset) ,
    Limit:  int(m.Limit) ,
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

func(m *PB_ResponseExtra)ToFlat() *PB_ResponseExtra_Flat {
r := &PB_ResponseExtra_Flat{
    ErrorCode:  int(m.ErrorCode) ,
    ErrMessage:  m.ErrMessage ,
    RpcFullName:  m.RpcFullName ,
    Data:  []byte(m.Data) ,
}
return r
}

func(m *PB_Pager)ToFlat() *PB_Pager_Flat {
r := &PB_Pager_Flat{
    Page:  int(m.Page) ,
    Limit:  int(m.Limit) ,
    Last:  int(m.Last) ,
}
return r
}

func(m *PB_UserParam_CheckUserName2)ToFlat() *PB_UserParam_CheckUserName2_Flat {
r := &PB_UserParam_CheckUserName2_Flat{
}
return r
}

func(m *PB_UserResponse_CheckUserName2)ToFlat() *PB_UserResponse_CheckUserName2_Flat {
r := &PB_UserResponse_CheckUserName2_Flat{
}
return r
}

func(m *PB_ChatParam_AddNewMessage)ToFlat() *PB_ChatParam_AddNewMessage_Flat {
r := &PB_ChatParam_AddNewMessage_Flat{
    ReplyToMessageId:  int(m.ReplyToMessageId) ,
    Blob:  []byte(m.Blob) ,

}
return r
}

func(m *PB_ChatResponse_AddNewMessage)ToFlat() *PB_ChatResponse_AddNewMessage_Flat {
r := &PB_ChatResponse_AddNewMessage_Flat{
    MessageKey:  m.MessageKey ,
    NewMessageId:  int(m.NewMessageId) ,
    MessageFileKey:  m.MessageFileKey ,
    NewMessageFileId:  int(m.NewMessageFileId) ,
    AtTime:  int(m.AtTime) ,
}
return r
}

func(m *PB_ChatParam_SetRoomActionDoing)ToFlat() *PB_ChatParam_SetRoomActionDoing_Flat {
r := &PB_ChatParam_SetRoomActionDoing_Flat{
    GroupId:  int(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,

}
return r
}

func(m *PB_ChatResponse_SetRoomActionDoing)ToFlat() *PB_ChatResponse_SetRoomActionDoing_Flat {
r := &PB_ChatResponse_SetRoomActionDoing_Flat{
}
return r
}

func(m *PB_ChatParam_SetChatMessagesRangeAsSeen)ToFlat() *PB_ChatParam_SetChatMessagesRangeAsSeen_Flat {
r := &PB_ChatParam_SetChatMessagesRangeAsSeen_Flat{
    ChatKey:  m.ChatKey ,
    BottomMessageId:  int(m.BottomMessageId) ,
    TopMessageId:  int(m.TopMessageId) ,
    SeenTimeMs:  int(m.SeenTimeMs) ,
}
return r
}

func(m *PB_ChatResponse_SetChatMessagesRangeAsSeen)ToFlat() *PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat {
r := &PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat{
}
return r
}

func(m *PB_ChatParam_DeleteChatHistory)ToFlat() *PB_ChatParam_DeleteChatHistory_Flat {
r := &PB_ChatParam_DeleteChatHistory_Flat{
    ChatKey:  m.ChatKey ,
    FromMessageId:  int(m.FromMessageId) ,
}
return r
}

func(m *PB_ChatResponse_DeleteChatHistory)ToFlat() *PB_ChatResponse_DeleteChatHistory_Flat {
r := &PB_ChatResponse_DeleteChatHistory_Flat{
}
return r
}

func(m *PB_ChatParam_DeleteMessagesByIds)ToFlat() *PB_ChatParam_DeleteMessagesByIds_Flat {
r := &PB_ChatParam_DeleteMessagesByIds_Flat{
    ChatKey:  m.ChatKey ,
    Both:  m.Both ,
    MessagesIds:  helper.SliceInt64ToInt(m.MessagesIds) ,
}
return r
}

func(m *PB_ChatResponse_DeleteMessagesByIds)ToFlat() *PB_ChatResponse_DeleteMessagesByIds_Flat {
r := &PB_ChatResponse_DeleteMessagesByIds_Flat{
}
return r
}

func(m *PB_ChatParam_SetMessagesAsReceived)ToFlat() *PB_ChatParam_SetMessagesAsReceived_Flat {
r := &PB_ChatParam_SetMessagesAsReceived_Flat{
    ChatRoom:  m.ChatRoom ,
    MessageIds:  helper.SliceInt64ToInt(m.MessageIds) ,
}
return r
}

func(m *PB_ChatResponse_SetMessagesAsReceived)ToFlat() *PB_ChatResponse_SetMessagesAsReceived_Flat {
r := &PB_ChatResponse_SetMessagesAsReceived_Flat{
}
return r
}

func(m *PB_ChatParam_EditMessage)ToFlat() *PB_ChatParam_EditMessage_Flat {
r := &PB_ChatParam_EditMessage_Flat{
    RoomKey:  m.RoomKey ,
    MessageId:  int(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_ChatResponse_EditMessage)ToFlat() *PB_ChatResponse_EditMessage_Flat {
r := &PB_ChatResponse_EditMessage_Flat{
}
return r
}

func(m *PB_ChatParam_GetChatList)ToFlat() *PB_ChatParam_GetChatList_Flat {
r := &PB_ChatParam_GetChatList_Flat{
}
return r
}

func(m *PB_ChatResponse_GetChatList)ToFlat() *PB_ChatResponse_GetChatList_Flat {
r := &PB_ChatResponse_GetChatList_Flat{


}
return r
}

func(m *PB_ChatParam_GetChatHistoryToOlder)ToFlat() *PB_ChatParam_GetChatHistoryToOlder_Flat {
r := &PB_ChatParam_GetChatHistoryToOlder_Flat{
    ChatKey:  m.ChatKey ,
    Limit:  int(m.Limit) ,
    FromTopMessageId:  int(m.FromTopMessageId) ,
}
return r
}

func(m *PB_ChatResponse_GetChatHistoryToOlder)ToFlat() *PB_ChatResponse_GetChatHistoryToOlder_Flat {
r := &PB_ChatResponse_GetChatHistoryToOlder_Flat{

    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_ChatParam_GetFreshAllDirectMessagesList)ToFlat() *PB_ChatParam_GetFreshAllDirectMessagesList_Flat {
r := &PB_ChatParam_GetFreshAllDirectMessagesList_Flat{
    LowMessageId:  int(m.LowMessageId) ,
    HighMessageId:  int(m.HighMessageId) ,
}
return r
}

func(m *PB_ChatResponse_GetFreshAllDirectMessagesList)ToFlat() *PB_ChatResponse_GetFreshAllDirectMessagesList_Flat {
r := &PB_ChatResponse_GetFreshAllDirectMessagesList_Flat{

    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_OtherParam_Echo)ToFlat() *PB_OtherParam_Echo_Flat {
r := &PB_OtherParam_Echo_Flat{
    Text:  m.Text ,
}
return r
}

func(m *PB_OtherResponse_Echo)ToFlat() *PB_OtherResponse_Echo_Flat {
r := &PB_OtherResponse_Echo_Flat{
    Text:  m.Text ,
}
return r
}

func(m *PB_PageParam_GetCommentsPage)ToFlat() *PB_PageParam_GetCommentsPage_Flat {
r := &PB_PageParam_GetCommentsPage_Flat{
    PostId:  int(m.PostId) ,

}
return r
}

func(m *PB_PageResponse_GetCommentsPage)ToFlat() *PB_PageResponse_GetCommentsPage_Flat {
r := &PB_PageResponse_GetCommentsPage_Flat{


}
return r
}

func(m *PB_PageParam_GetHomePage)ToFlat() *PB_PageParam_GetHomePage_Flat {
r := &PB_PageParam_GetHomePage_Flat{

}
return r
}

func(m *PB_PageResponse_GetHomePage)ToFlat() *PB_PageResponse_GetHomePage_Flat {
r := &PB_PageResponse_GetHomePage_Flat{


}
return r
}

func(m *PB_PageParam_GetProfilePage)ToFlat() *PB_PageParam_GetProfilePage_Flat {
r := &PB_PageParam_GetProfilePage_Flat{
}
return r
}

func(m *PB_PageResponse_GetProfilePage)ToFlat() *PB_PageResponse_GetProfilePage_Flat {
r := &PB_PageResponse_GetProfilePage_Flat{

}
return r
}

func(m *PB_PageParam_GetLikesPage)ToFlat() *PB_PageParam_GetLikesPage_Flat {
r := &PB_PageParam_GetLikesPage_Flat{
    PostId:  int(m.PostId) ,

}
return r
}

func(m *PB_PageResponse_GetLikesPage)ToFlat() *PB_PageResponse_GetLikesPage_Flat {
r := &PB_PageResponse_GetLikesPage_Flat{


}
return r
}

func(m *PB_PageParam_GetFollowersPage)ToFlat() *PB_PageParam_GetFollowersPage_Flat {
r := &PB_PageParam_GetFollowersPage_Flat{
    UserId:  int(m.UserId) ,

}
return r
}

func(m *PB_PageResponse_GetFollowersPage)ToFlat() *PB_PageResponse_GetFollowersPage_Flat {
r := &PB_PageResponse_GetFollowersPage_Flat{


}
return r
}

func(m *PB_PageParam_GetFollowingsPage)ToFlat() *PB_PageParam_GetFollowingsPage_Flat {
r := &PB_PageParam_GetFollowingsPage_Flat{
    UserId:  int(m.UserId) ,

}
return r
}

func(m *PB_PageResponse_GetFollowingsPage)ToFlat() *PB_PageResponse_GetFollowingsPage_Flat {
r := &PB_PageResponse_GetFollowingsPage_Flat{


}
return r
}

func(m *PB_PageParam_GetNotifiesPage)ToFlat() *PB_PageParam_GetNotifiesPage_Flat {
r := &PB_PageParam_GetNotifiesPage_Flat{

}
return r
}

func(m *PB_PageResponse_GetNotifiesPage)ToFlat() *PB_PageResponse_GetNotifiesPage_Flat {
r := &PB_PageResponse_GetNotifiesPage_Flat{


    RemoveIdsList:  helper.SliceInt64ToInt(m.RemoveIdsList) ,
}
return r
}

func(m *PB_PageParam_GetUserActionsPage)ToFlat() *PB_PageParam_GetUserActionsPage_Flat {
r := &PB_PageParam_GetUserActionsPage_Flat{

}
return r
}

func(m *PB_PageResponse_GetUserActionsPage)ToFlat() *PB_PageResponse_GetUserActionsPage_Flat {
r := &PB_PageResponse_GetUserActionsPage_Flat{


}
return r
}

func(m *PB_PageParam_GetSuggestedPostsPage)ToFlat() *PB_PageParam_GetSuggestedPostsPage_Flat {
r := &PB_PageParam_GetSuggestedPostsPage_Flat{

}
return r
}

func(m *PB_PageResponse_GetSuggestedPostsPage)ToFlat() *PB_PageResponse_GetSuggestedPostsPage_Flat {
r := &PB_PageResponse_GetSuggestedPostsPage_Flat{


}
return r
}

func(m *PB_PageParam_GetSuggestedUsrsPage)ToFlat() *PB_PageParam_GetSuggestedUsrsPage_Flat {
r := &PB_PageParam_GetSuggestedUsrsPage_Flat{

}
return r
}

func(m *PB_PageResponse_GetSuggestedUsrsPage)ToFlat() *PB_PageResponse_GetSuggestedUsrsPage_Flat {
r := &PB_PageResponse_GetSuggestedUsrsPage_Flat{


}
return r
}

func(m *PB_PageParam_GetSuggestedTagsPage)ToFlat() *PB_PageParam_GetSuggestedTagsPage_Flat {
r := &PB_PageParam_GetSuggestedTagsPage_Flat{

}
return r
}

func(m *PB_PageResponse_GetSuggestedTagsPage)ToFlat() *PB_PageResponse_GetSuggestedTagsPage_Flat {
r := &PB_PageResponse_GetSuggestedTagsPage_Flat{

}
return r
}

func(m *PB_PageParam_GetLastPostsPage)ToFlat() *PB_PageParam_GetLastPostsPage_Flat {
r := &PB_PageParam_GetLastPostsPage_Flat{

}
return r
}

func(m *PB_PageResponse_GetLastPostsPage)ToFlat() *PB_PageResponse_GetLastPostsPage_Flat {
r := &PB_PageResponse_GetLastPostsPage_Flat{


}
return r
}

func(m *PB_PageParam_GetTagPage)ToFlat() *PB_PageParam_GetTagPage_Flat {
r := &PB_PageParam_GetTagPage_Flat{

}
return r
}

func(m *PB_PageResponse_GetTagPage)ToFlat() *PB_PageResponse_GetTagPage_Flat {
r := &PB_PageResponse_GetTagPage_Flat{


}
return r
}

func(m *PB_PageParam_SearchTagsPage)ToFlat() *PB_PageParam_SearchTagsPage_Flat {
r := &PB_PageParam_SearchTagsPage_Flat{
}
return r
}

func(m *PB_PageResponse_SearchTagsPage)ToFlat() *PB_PageResponse_SearchTagsPage_Flat {
r := &PB_PageResponse_SearchTagsPage_Flat{

}
return r
}

func(m *PB_PageParam_SearchUsersPage)ToFlat() *PB_PageParam_SearchUsersPage_Flat {
r := &PB_PageParam_SearchUsersPage_Flat{
}
return r
}

func(m *PB_PageResponse_SearchUsersPage)ToFlat() *PB_PageResponse_SearchUsersPage_Flat {
r := &PB_PageResponse_SearchUsersPage_Flat{

}
return r
}

func(m *PB_PageParam_)ToFlat() *PB_PageParam__Flat {
r := &PB_PageParam__Flat{
}
return r
}

func(m *PB_PageResponse_)ToFlat() *PB_PageResponse__Flat {
r := &PB_PageResponse__Flat{

}
return r
}

func(m *PB_SearchResponse_AddNewC)ToFlat() *PB_SearchResponse_AddNewC_Flat {
r := &PB_SearchResponse_AddNewC_Flat{
}
return r
}

func(m *PB_SocialParam_AddComment)ToFlat() *PB_SocialParam_AddComment_Flat {
r := &PB_SocialParam_AddComment_Flat{
    PostId:  int(m.PostId) ,
    Text:  m.Text ,
}
return r
}

func(m *PB_SocialResponse_AddComment)ToFlat() *PB_SocialResponse_AddComment_Flat {
r := &PB_SocialResponse_AddComment_Flat{


}
return r
}

func(m *PB_SocialParam_DeleteComment)ToFlat() *PB_SocialParam_DeleteComment_Flat {
r := &PB_SocialParam_DeleteComment_Flat{
    PostId:  int(m.PostId) ,
    CommentId:  int(m.CommentId) ,
}
return r
}

func(m *PB_SocialResponse_DeleteComment)ToFlat() *PB_SocialResponse_DeleteComment_Flat {
r := &PB_SocialResponse_DeleteComment_Flat{

    Deleted:  m.Deleted ,
}
return r
}

func(m *PB_SocialParam_AddPost)ToFlat() *PB_SocialParam_AddPost_Flat {
r := &PB_SocialParam_AddPost_Flat{
    Text:  m.Text ,
}
return r
}

func(m *PB_SocialResponse_AddPost)ToFlat() *PB_SocialResponse_AddPost_Flat {
r := &PB_SocialResponse_AddPost_Flat{


}
return r
}

func(m *PB_SocialParam_EditPost)ToFlat() *PB_SocialParam_EditPost_Flat {
r := &PB_SocialParam_EditPost_Flat{
    PostId:  int(m.PostId) ,
    Text:  m.Text ,
}
return r
}

func(m *PB_SocialResponse_EditPost)ToFlat() *PB_SocialResponse_EditPost_Flat {
r := &PB_SocialResponse_EditPost_Flat{


}
return r
}

func(m *PB_SocialParam_DeletePost)ToFlat() *PB_SocialParam_DeletePost_Flat {
r := &PB_SocialParam_DeletePost_Flat{
    PostId:  int(m.PostId) ,
}
return r
}

func(m *PB_SocialResponse_DeletePost)ToFlat() *PB_SocialResponse_DeletePost_Flat {
r := &PB_SocialResponse_DeletePost_Flat{

    Done:  m.Done ,
}
return r
}

func(m *PB_SocialParam_ArchivePost)ToFlat() *PB_SocialParam_ArchivePost_Flat {
r := &PB_SocialParam_ArchivePost_Flat{
    PostId:  int(m.PostId) ,
}
return r
}

func(m *PB_SocialResponse_ArchivePost)ToFlat() *PB_SocialResponse_ArchivePost_Flat {
r := &PB_SocialResponse_ArchivePost_Flat{

    Done:  m.Done ,
}
return r
}

func(m *PB_SocialParam_LikePost)ToFlat() *PB_SocialParam_LikePost_Flat {
r := &PB_SocialParam_LikePost_Flat{
    PostId:  int(m.PostId) ,
}
return r
}

func(m *PB_SocialResponse_LikePost)ToFlat() *PB_SocialResponse_LikePost_Flat {
r := &PB_SocialResponse_LikePost_Flat{

    Done:  m.Done ,
}
return r
}

func(m *PB_SocialParam_UnLikePost)ToFlat() *PB_SocialParam_UnLikePost_Flat {
r := &PB_SocialParam_UnLikePost_Flat{
    PostId:  int(m.PostId) ,
}
return r
}

func(m *PB_SocialResponse_UnLikePost)ToFlat() *PB_SocialResponse_UnLikePost_Flat {
r := &PB_SocialResponse_UnLikePost_Flat{

    Done:  m.Done ,
}
return r
}

func(m *PB_SocialParam_FollowUser)ToFlat() *PB_SocialParam_FollowUser_Flat {
r := &PB_SocialParam_FollowUser_Flat{
    UserId:  int(m.UserId) ,
}
return r
}

func(m *PB_SocialResponse_FollowUser)ToFlat() *PB_SocialResponse_FollowUser_Flat {
r := &PB_SocialResponse_FollowUser_Flat{

}
return r
}

func(m *PB_SocialParam_UnFollowUser)ToFlat() *PB_SocialParam_UnFollowUser_Flat {
r := &PB_SocialParam_UnFollowUser_Flat{
    UserId:  int(m.UserId) ,
}
return r
}

func(m *PB_SocialResponse_UnFollowUser)ToFlat() *PB_SocialResponse_UnFollowUser_Flat {
r := &PB_SocialResponse_UnFollowUser_Flat{

}
return r
}

func(m *PB_UserParam_BlockUser)ToFlat() *PB_UserParam_BlockUser_Flat {
r := &PB_UserParam_BlockUser_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserResponse_BlockUser)ToFlat() *PB_UserResponse_BlockUser_Flat {
r := &PB_UserResponse_BlockUser_Flat{
    ByUserId:  int(m.ByUserId) ,
    TargetUserId:  int(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UnBlockUser)ToFlat() *PB_UserParam_UnBlockUser_Flat {
r := &PB_UserParam_UnBlockUser_Flat{
    Offset:  int(m.Offset) ,
    Limit:  int(m.Limit) ,
}
return r
}

func(m *PB_UserResponse_UnBlockUser)ToFlat() *PB_UserResponse_UnBlockUser_Flat {
r := &PB_UserResponse_UnBlockUser_Flat{

}
return r
}

func(m *PB_UserParam_BlockedList)ToFlat() *PB_UserParam_BlockedList_Flat {
r := &PB_UserParam_BlockedList_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserResponse_BlockedList)ToFlat() *PB_UserResponse_BlockedList_Flat {
r := &PB_UserResponse_BlockedList_Flat{
    ByUserId:  int(m.ByUserId) ,
    TargetUserId:  int(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UpdateAbout)ToFlat() *PB_UserParam_UpdateAbout_Flat {
r := &PB_UserParam_UpdateAbout_Flat{
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserResponse_UpdateAbout)ToFlat() *PB_UserResponse_UpdateAbout_Flat {
r := &PB_UserResponse_UpdateAbout_Flat{
    UserId:  int(m.UserId) ,
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserParam_UpdateUserName)ToFlat() *PB_UserParam_UpdateUserName_Flat {
r := &PB_UserParam_UpdateUserName_Flat{
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserResponse_UpdateUserName)ToFlat() *PB_UserResponse_UpdateUserName_Flat {
r := &PB_UserResponse_UpdateUserName_Flat{
    UserId:  int(m.UserId) ,
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserParam_ChangeAvatar)ToFlat() *PB_UserParam_ChangeAvatar_Flat {
r := &PB_UserParam_ChangeAvatar_Flat{
    None:  m.None ,
    ImageData2:  []byte(m.ImageData2) ,
}
return r
}

func(m *PB_UserResponse_ChangeAvatar)ToFlat() *PB_UserResponse_ChangeAvatar_Flat {
r := &PB_UserResponse_ChangeAvatar_Flat{
}
return r
}

func(m *PB_UserParam_ChangePrivacy)ToFlat() *PB_UserParam_ChangePrivacy_Flat {
r := &PB_UserParam_ChangePrivacy_Flat{

}
return r
}

func(m *PB_UserResponseOffline_ChangePrivacy)ToFlat() *PB_UserResponseOffline_ChangePrivacy_Flat {
r := &PB_UserResponseOffline_ChangePrivacy_Flat{
}
return r
}

func(m *PB_UserParam_CheckUserName)ToFlat() *PB_UserParam_CheckUserName_Flat {
r := &PB_UserParam_CheckUserName_Flat{

}
return r
}

func(m *PB_UserResponse_CheckUserName)ToFlat() *PB_UserResponse_CheckUserName_Flat {
r := &PB_UserResponse_CheckUserName_Flat{
}
return r
}

func(m *UserView)ToFlat() *UserView_Flat {
r := &UserView_Flat{
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
    Seq:  int(m.Seq) ,
}
return r
}

func(m *PB_Chat)ToFlat() *PB_Chat_Flat {
r := &PB_Chat_Flat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnum:  int(m.RoomTypeEnum) ,
    UserId:  int(m.UserId) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    CreatedTime:  int(m.CreatedTime) ,
    StartMessageIdFrom:  int(m.StartMessageIdFrom) ,
    LastMessageId:  int(m.LastMessageId) ,
    LastSeenMessageId:  int(m.LastSeenMessageId) ,
    UpdatedMs:  int(m.UpdatedMs) ,
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
    Seq:  int(m.Seq) ,
}
return r
}

func(m *PB_DirectMessage)ToFlat() *PB_DirectMessage_Flat {
r := &PB_DirectMessage_Flat{
    MessageId:  int(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessageTypeEnumId:  int(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int(m.CreatedSe) ,
    PeerReceivedTime:  int(m.PeerReceivedTime) ,
    PeerSeenTime:  int(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int(m.DeliviryStatusEnumId) ,
}
return r
}

func(m *PB_DirectOffline)ToFlat() *PB_DirectOffline_Flat {
r := &PB_DirectOffline_Flat{
    DirectOfflineId:  int(m.DirectOfflineId) ,
    ToUserId:  int(m.ToUserId) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int(m.MessageId) ,
    MessageFileId:  int(m.MessageFileId) ,
    PBClass:  m.PBClass ,
    DataPB:  []byte(m.DataPB) ,
    DataJson:  m.DataJson ,
    DataTemp:  m.DataTemp ,
    AtTimeMs:  int(m.AtTimeMs) ,
}
return r
}

func(m *PB_DirectToMessage)ToFlat() *PB_DirectToMessage_Flat {
r := &PB_DirectToMessage_Flat{
    Id:  int(m.Id) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int(m.MessageId) ,
    SourceEnumId:  int(m.SourceEnumId) ,
}
return r
}

func(m *PB_Feed)ToFlat() *PB_Feed_Flat {
r := &PB_Feed_Flat{
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    FeedId:  int(m.FeedId) ,
}
return r
}

func(m *PB_FollowingList)ToFlat() *PB_FollowingList_Flat {
r := &PB_FollowingList_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    ListType:  int(m.ListType) ,
    Name:  m.Name ,
    Count:  int(m.Count) ,
    IsAuto:  int(m.IsAuto) ,
    IsPimiry:  int(m.IsPimiry) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_FollowingListMember)ToFlat() *PB_FollowingListMember_Flat {
r := &PB_FollowingListMember_Flat{
    Id:  int(m.Id) ,
    ListId:  int(m.ListId) ,
    UserId:  int(m.UserId) ,
    FollowedUserId:  int(m.FollowedUserId) ,
    FollowType:  int(m.FollowType) ,
    UpdatedTimeMs:  int(m.UpdatedTimeMs) ,
}
return r
}

func(m *PB_FollowingListMemberHistory)ToFlat() *PB_FollowingListMemberHistory_Flat {
r := &PB_FollowingListMemberHistory_Flat{
    Id:  int(m.Id) ,
    ListId:  int(m.ListId) ,
    UserId:  int(m.UserId) ,
    FollowedUserId:  int(m.FollowedUserId) ,
    FollowType:  int(m.FollowType) ,
    UpdatedTimeMs:  int(m.UpdatedTimeMs) ,
    FollowId:  int(m.FollowId) ,
}
return r
}

func(m *PB_GeneralLog)ToFlat() *PB_GeneralLog_Flat {
r := &PB_GeneralLog_Flat{
    Id:  int(m.Id) ,
    ToUserId:  int(m.ToUserId) ,
    TargetId:  int(m.TargetId) ,
    LogTypeId:  int(m.LogTypeId) ,
    ExtraPB:  []byte(m.ExtraPB) ,
    ExtraJson:  m.ExtraJson ,
    CreatedMs:  int(m.CreatedMs) ,
}
return r
}

func(m *PB_Group)ToFlat() *PB_Group_Flat {
r := &PB_Group_Flat{
    GroupId:  int(m.GroupId) ,
    GroupName:  m.GroupName ,
    MembersCount:  int(m.MembersCount) ,
    GroupPrivacyEnum:  int(m.GroupPrivacyEnum) ,
    CreatorUserId:  int(m.CreatorUserId) ,
    CreatedTime:  int(m.CreatedTime) ,
    UpdatedMs:  int(m.UpdatedMs) ,
    CurrentSeq:  int(m.CurrentSeq) ,
}
return r
}

func(m *PB_GroupMember)ToFlat() *PB_GroupMember_Flat {
r := &PB_GroupMember_Flat{
    Id:  int(m.Id) ,
    GroupId:  int(m.GroupId) ,
    GroupKey:  m.GroupKey ,
    UserId:  int(m.UserId) ,
    ByUserId:  int(m.ByUserId) ,
    GroupRoleEnumId:  int(m.GroupRoleEnumId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_GroupMessage)ToFlat() *PB_GroupMessage_Flat {
r := &PB_GroupMessage_Flat{
    MessageId:  int(m.MessageId) ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessageTypeEnum:  int(m.MessageTypeEnum) ,
    Text:  m.Text ,
    CreatedMs:  int(m.CreatedMs) ,
    DeliveryStatusEnum:  int(m.DeliveryStatusEnum) ,
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

func(m *PB_Media)ToFlat() *PB_Media_Flat {
r := &PB_Media_Flat{
    MediaId:  int(m.MediaId) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    AlbumId:  int(m.AlbumId) ,
    MediaTypeEnum:  int(m.MediaTypeEnum) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Size:  int(m.Size) ,
    Duration:  int(m.Duration) ,
    Md5Hash:  m.Md5Hash ,
    Color:  m.Color ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_MessageFile)ToFlat() *PB_MessageFile_Flat {
r := &PB_MessageFile_Flat{
    MessageFileId:  int(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    UserId:  int(m.UserId) ,
    Title:  m.Title ,
    Size:  int(m.Size) ,
    FileTypeEnum:  int(m.FileTypeEnum) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    Md5Hash:  m.Md5Hash ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_Notify)ToFlat() *PB_Notify_Flat {
r := &PB_Notify_Flat{
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
    Seq:  int(m.Seq) ,
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
    PhoneDisplayName:  m.PhoneDisplayName ,
    PhoneFamilyName:  m.PhoneFamilyName ,
    PhoneNumber:  m.PhoneNumber ,
    PhoneNormalizedNumber:  m.PhoneNormalizedNumber ,
    PhoneContactRowId:  int(m.PhoneContactRowId) ,
    UserId:  int(m.UserId) ,
    DeviceUuidId:  int(m.DeviceUuidId) ,
    CreatedTime:  int(m.CreatedTime) ,
    UpdatedTime:  int(m.UpdatedTime) ,
}
return r
}

func(m *PB_Post)ToFlat() *PB_Post_Flat {
r := &PB_Post_Flat{
    PostId:  int(m.PostId) ,
    UserId:  int(m.UserId) ,
    PostTypeEnum:  int(m.PostTypeEnum) ,
    MediaId:  int(m.MediaId) ,
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
}
return r
}

func(m *PB_RecommendUser)ToFlat() *PB_RecommendUser_Flat {
r := &PB_RecommendUser_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    TargetId:  int(m.TargetId) ,
    Weight:  float32(m.Weight) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_SearchClicked)ToFlat() *PB_SearchClicked_Flat {
r := &PB_SearchClicked_Flat{
    Id:  int(m.Id) ,
    Query:  m.Query ,
    ClickType:  int(m.ClickType) ,
    TargetId:  int(m.TargetId) ,
    UserId:  int(m.UserId) ,
    CreatedTime:  int(m.CreatedTime) ,
}
return r
}

func(m *PB_Session)ToFlat() *PB_Session_Flat {
r := &PB_Session_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    SessionUuid:  m.SessionUuid ,
    ClientUuid:  m.ClientUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastActivityTime:  int(m.LastActivityTime) ,
    LastIpAddress:  m.LastIpAddress ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    LastNetworkTypeEnumId:  int(m.LastNetworkTypeEnumId) ,
    AppVersion:  int(m.AppVersion) ,
    UpdatedTime:  int(m.UpdatedTime) ,
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

func(m *PB_TagsPost)ToFlat() *PB_TagsPost_Flat {
r := &PB_TagsPost_Flat{
    Id:  int(m.Id) ,
    TagId:  int(m.TagId) ,
    PostId:  int(m.PostId) ,
    PostTypeEnum:  int(m.PostTypeEnum) ,
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
    UserTypeEnum:  int(m.UserTypeEnum) ,
    UserLevelEnum:  int(m.UserLevelEnum) ,
    AvatarId:  int(m.AvatarId) ,
    ProfilePrivacyEnum:  int(m.ProfilePrivacyEnum) ,
    Phone:  int(m.Phone) ,
    About:  m.About ,
    Email:  m.Email ,
    PasswordHash:  m.PasswordHash ,
    PasswordSalt:  m.PasswordSalt ,
    FollowersCount:  int(m.FollowersCount) ,
    FollowingCount:  int(m.FollowingCount) ,
    PostsCount:  int(m.PostsCount) ,
    MediaCount:  int(m.MediaCount) ,
    LikesCount:  int(m.LikesCount) ,
    ResharedCount:  int(m.ResharedCount) ,
    LastActionTime:  int(m.LastActionTime) ,
    LastPostTime:  int(m.LastPostTime) ,
    PrimaryFollowingList:  int(m.PrimaryFollowingList) ,
    CreatedSe:  int(m.CreatedSe) ,
    UpdatedMs:  int(m.UpdatedMs) ,
    OnlinePrivacyEnum:  int(m.OnlinePrivacyEnum) ,
    LastActivityTime:  int(m.LastActivityTime) ,
    Phone2:  m.Phone2 ,
}
return r
}

func(m *PB_UserMetaInfo)ToFlat() *PB_UserMetaInfo_Flat {
r := &PB_UserMetaInfo_Flat{
    Id:  int(m.Id) ,
    UserId:  int(m.UserId) ,
    IsNotificationDirty:  int(m.IsNotificationDirty) ,
    LastUserRecGen:  int(m.LastUserRecGen) ,
}
return r
}

func(m *PB_UserPassword)ToFlat() *PB_UserPassword_Flat {
r := &PB_UserPassword_Flat{
    UserId:  int(m.UserId) ,
    Password:  m.Password ,
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

func(m *PB_MediaView)ToFlat() *PB_MediaView_Flat {
r := &PB_MediaView_Flat{
    MediaId:  int(m.MediaId) ,
    UserId:  int(m.UserId) ,
    PostId:  int(m.PostId) ,
    AlbumId:  int(m.AlbumId) ,
    MediaTypeEnum:  int(m.MediaTypeEnum) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Size:  int(m.Size) ,
    Duration:  int(m.Duration) ,
    Color:  m.Color ,
    CreatedTime:  int(m.CreatedTime) ,
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

func(m *PB_UserView)ToFlat() *PB_UserView_Flat {
r := &PB_UserView_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,


    AvatarId:  int(m.AvatarId) ,
    ProfilePrivacyEnum:  int(m.ProfilePrivacyEnum) ,
    Phone:  int(m.Phone) ,
    About:  m.About ,
    FollowersCount:  int(m.FollowersCount) ,
    FollowingCount:  int(m.FollowingCount) ,
    PostsCount:  int(m.PostsCount) ,
    MediaCount:  int(m.MediaCount) ,
    LikesCount:  int(m.LikesCount) ,
    ResharedCount:  int(m.ResharedCount) ,

    LastActiveTime:  int(m.LastActiveTime) ,

}
return r
}

func(m *PB_TopProfileView)ToFlat() *PB_TopProfileView_Flat {
r := &PB_TopProfileView_Flat{

}
return r
}

func(m *PB_ChatView)ToFlat() *PB_ChatView_Flat {
r := &PB_ChatView_Flat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnumId:  int(m.RoomTypeEnumId) ,
    UserId:  int(m.UserId) ,
    PeerUserId:  int(m.PeerUserId) ,
    GroupId:  int(m.GroupId) ,
    CreatedSe:  int(m.CreatedSe) ,
    UpdatedMs:  int(m.UpdatedMs) ,
    LastMessageId:  int(m.LastMessageId) ,
    LastDeletedMessageId:  int(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int(m.LastSeenMessageId) ,
    LastSeqSeen:  int(m.LastSeqSeen) ,
    LastSeqDelete:  int(m.LastSeqDelete) ,
    CurrentSeq:  int(m.CurrentSeq) ,

    SharedMediaCount:  int(m.SharedMediaCount) ,
    UnseenCount:  int(m.UnseenCount) ,


}
return r
}

func(m *PB_MessageView)ToFlat() *PB_MessageView_Flat {
r := &PB_MessageView_Flat{
    MessageId:  int(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int(m.UserId) ,
    MessageFileId:  int(m.MessageFileId) ,
    MessageTypeEnumId:  int(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int(m.CreatedSe) ,
    PeerReceivedTime:  int(m.PeerReceivedTime) ,
    PeerSeenTime:  int(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int(m.DeliviryStatusEnumId) ,
    ChatKey:  m.ChatKey ,
    RoomTypeEnumId:  int(m.RoomTypeEnumId) ,
    IsByMe:  m.IsByMe ,
    RemoteId:  int(m.RemoteId) ,

}
return r
}

func(m *PB_MessageFileView)ToFlat() *PB_MessageFileView_Flat {
r := &PB_MessageFileView_Flat{
    MessageFileId:  int(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    OriginalUserId:  int(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int(m.Size) ,
    FileTypeEnumId:  int(m.FileTypeEnumId) ,
    Width:  int(m.Width) ,
    Height:  int(m.Height) ,
    Duration:  int(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int(m.HashAccess) ,
    CreatedSe:  int(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int(m.ServerId) ,
    CanDel:  int(m.CanDel) ,
    ServerThumbLocalSrc:  m.ServerThumbLocalSrc ,
    RemoteMessageFileId:  int(m.RemoteMessageFileId) ,
    LocalSrc:  m.LocalSrc ,
    ThumbLocalSrc:  m.ThumbLocalSrc ,
    MessageFileStatusId:  m.MessageFileStatusId ,
}
return r
}

func(m *PB_UserView3)ToFlat() *PB_UserView3_Flat {
r := &PB_UserView3_Flat{
    UserId:  int(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int(m.PrivacyProfile) ,
    IsDeleted:  int(m.IsDeleted) ,
    FollowersCount:  int(m.FollowersCount) ,
    FollowingCount:  int(m.FollowingCount) ,
    PostsCount:  int(m.PostsCount) ,
    UpdatedTime:  int(m.UpdatedTime) ,
    AppVersion:  int(m.AppVersion) ,
    LastActivityTime:  int(m.LastActivityTime) ,
    FollowingType:  int(m.FollowingType) ,
}
return r
}



///// from_flat ///

func(m *GeoLocation_Flat)ToPB() *GeoLocation {
r := &GeoLocation{
    Lat:  m.Lat ,
    Lon:  m.Lon ,
}
return r
}

func(m *RoomMessageLog_Flat)ToPB() *RoomMessageLog {
r := &RoomMessageLog{

    TargetUserId:  uint64(m.TargetUserId) ,
    ByUserId:  uint64(m.ByUserId) ,
}
return r
}

func(m *RoomMessageForwardFrom_Flat)ToPB() *RoomMessageForwardFrom {
r := &RoomMessageForwardFrom{
    RoomId:  uint64(m.RoomId) ,
    MessageId:  uint64(m.MessageId) ,
    RoomTypeEnum:  uint32(m.RoomTypeEnum) ,
}
return r
}

func(m *RoomDraft_Flat)ToPB() *RoomDraft {
r := &RoomDraft{
    Message:  m.Message ,
    ReplyTo:  uint64(m.ReplyTo) ,
}
return r
}

func(m *ChatRoom_Flat)ToPB() *ChatRoom {
r := &ChatRoom{
}
return r
}

func(m *Pagination_Flat)ToPB() *Pagination {
r := &Pagination{
    Offset:  uint32(m.Offset) ,
    Limit:  uint32(m.Limit) ,
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

func(m *PB_ResponseExtra_Flat)ToPB() *PB_ResponseExtra {
r := &PB_ResponseExtra{
    ErrorCode:  int64(m.ErrorCode) ,
    ErrMessage:  m.ErrMessage ,
    RpcFullName:  m.RpcFullName ,
    Data:  m.Data ,
}
return r
}

func(m *PB_Pager_Flat)ToPB() *PB_Pager {
r := &PB_Pager{
    Page:  int64(m.Page) ,
    Limit:  int64(m.Limit) ,
    Last:  int64(m.Last) ,
}
return r
}

func(m *PB_UserParam_CheckUserName2_Flat)ToPB() *PB_UserParam_CheckUserName2 {
r := &PB_UserParam_CheckUserName2{
}
return r
}

func(m *PB_UserResponse_CheckUserName2_Flat)ToPB() *PB_UserResponse_CheckUserName2 {
r := &PB_UserResponse_CheckUserName2{
}
return r
}

func(m *PB_ChatParam_AddNewMessage_Flat)ToPB() *PB_ChatParam_AddNewMessage {
r := &PB_ChatParam_AddNewMessage{
    ReplyToMessageId:  int64(m.ReplyToMessageId) ,
    Blob:  m.Blob ,

}
return r
}

func(m *PB_ChatResponse_AddNewMessage_Flat)ToPB() *PB_ChatResponse_AddNewMessage {
r := &PB_ChatResponse_AddNewMessage{
    MessageKey:  m.MessageKey ,
    NewMessageId:  int64(m.NewMessageId) ,
    MessageFileKey:  m.MessageFileKey ,
    NewMessageFileId:  int64(m.NewMessageFileId) ,
    AtTime:  int64(m.AtTime) ,
}
return r
}

func(m *PB_ChatParam_SetRoomActionDoing_Flat)ToPB() *PB_ChatParam_SetRoomActionDoing {
r := &PB_ChatParam_SetRoomActionDoing{
    GroupId:  int64(m.GroupId) ,
    DirectRoomKey:  m.DirectRoomKey ,

}
return r
}

func(m *PB_ChatResponse_SetRoomActionDoing_Flat)ToPB() *PB_ChatResponse_SetRoomActionDoing {
r := &PB_ChatResponse_SetRoomActionDoing{
}
return r
}

func(m *PB_ChatParam_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_ChatParam_SetChatMessagesRangeAsSeen {
r := &PB_ChatParam_SetChatMessagesRangeAsSeen{
    ChatKey:  m.ChatKey ,
    BottomMessageId:  int64(m.BottomMessageId) ,
    TopMessageId:  int64(m.TopMessageId) ,
    SeenTimeMs:  int64(m.SeenTimeMs) ,
}
return r
}

func(m *PB_ChatResponse_SetChatMessagesRangeAsSeen_Flat)ToPB() *PB_ChatResponse_SetChatMessagesRangeAsSeen {
r := &PB_ChatResponse_SetChatMessagesRangeAsSeen{
}
return r
}

func(m *PB_ChatParam_DeleteChatHistory_Flat)ToPB() *PB_ChatParam_DeleteChatHistory {
r := &PB_ChatParam_DeleteChatHistory{
    ChatKey:  m.ChatKey ,
    FromMessageId:  int64(m.FromMessageId) ,
}
return r
}

func(m *PB_ChatResponse_DeleteChatHistory_Flat)ToPB() *PB_ChatResponse_DeleteChatHistory {
r := &PB_ChatResponse_DeleteChatHistory{
}
return r
}

func(m *PB_ChatParam_DeleteMessagesByIds_Flat)ToPB() *PB_ChatParam_DeleteMessagesByIds {
r := &PB_ChatParam_DeleteMessagesByIds{
    ChatKey:  m.ChatKey ,
    Both:  m.Both ,
    MessagesIds:  helper.SliceIntToInt64(m.MessagesIds) ,
}
return r
}

func(m *PB_ChatResponse_DeleteMessagesByIds_Flat)ToPB() *PB_ChatResponse_DeleteMessagesByIds {
r := &PB_ChatResponse_DeleteMessagesByIds{
}
return r
}

func(m *PB_ChatParam_SetMessagesAsReceived_Flat)ToPB() *PB_ChatParam_SetMessagesAsReceived {
r := &PB_ChatParam_SetMessagesAsReceived{
    ChatRoom:  m.ChatRoom ,
    MessageIds:  helper.SliceIntToInt64(m.MessageIds) ,
}
return r
}

func(m *PB_ChatResponse_SetMessagesAsReceived_Flat)ToPB() *PB_ChatResponse_SetMessagesAsReceived {
r := &PB_ChatResponse_SetMessagesAsReceived{
}
return r
}

func(m *PB_ChatParam_EditMessage_Flat)ToPB() *PB_ChatParam_EditMessage {
r := &PB_ChatParam_EditMessage{
    RoomKey:  m.RoomKey ,
    MessageId:  int64(m.MessageId) ,
    NewText:  m.NewText ,
}
return r
}

func(m *PB_ChatResponse_EditMessage_Flat)ToPB() *PB_ChatResponse_EditMessage {
r := &PB_ChatResponse_EditMessage{
}
return r
}

func(m *PB_ChatParam_GetChatList_Flat)ToPB() *PB_ChatParam_GetChatList {
r := &PB_ChatParam_GetChatList{
}
return r
}

func(m *PB_ChatResponse_GetChatList_Flat)ToPB() *PB_ChatResponse_GetChatList {
r := &PB_ChatResponse_GetChatList{


}
return r
}

func(m *PB_ChatParam_GetChatHistoryToOlder_Flat)ToPB() *PB_ChatParam_GetChatHistoryToOlder {
r := &PB_ChatParam_GetChatHistoryToOlder{
    ChatKey:  m.ChatKey ,
    Limit:  int32(m.Limit) ,
    FromTopMessageId:  int64(m.FromTopMessageId) ,
}
return r
}

func(m *PB_ChatResponse_GetChatHistoryToOlder_Flat)ToPB() *PB_ChatResponse_GetChatHistoryToOlder {
r := &PB_ChatResponse_GetChatHistoryToOlder{

    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_ChatParam_GetFreshAllDirectMessagesList_Flat)ToPB() *PB_ChatParam_GetFreshAllDirectMessagesList {
r := &PB_ChatParam_GetFreshAllDirectMessagesList{
    LowMessageId:  int64(m.LowMessageId) ,
    HighMessageId:  int64(m.HighMessageId) ,
}
return r
}

func(m *PB_ChatResponse_GetFreshAllDirectMessagesList_Flat)ToPB() *PB_ChatResponse_GetFreshAllDirectMessagesList {
r := &PB_ChatResponse_GetFreshAllDirectMessagesList{

    HasMore:  m.HasMore ,
}
return r
}

func(m *PB_OtherParam_Echo_Flat)ToPB() *PB_OtherParam_Echo {
r := &PB_OtherParam_Echo{
    Text:  m.Text ,
}
return r
}

func(m *PB_OtherResponse_Echo_Flat)ToPB() *PB_OtherResponse_Echo {
r := &PB_OtherResponse_Echo{
    Text:  m.Text ,
}
return r
}

func(m *PB_PageParam_GetCommentsPage_Flat)ToPB() *PB_PageParam_GetCommentsPage {
r := &PB_PageParam_GetCommentsPage{
    PostId:  int64(m.PostId) ,

}
return r
}

func(m *PB_PageResponse_GetCommentsPage_Flat)ToPB() *PB_PageResponse_GetCommentsPage {
r := &PB_PageResponse_GetCommentsPage{


}
return r
}

func(m *PB_PageParam_GetHomePage_Flat)ToPB() *PB_PageParam_GetHomePage {
r := &PB_PageParam_GetHomePage{

}
return r
}

func(m *PB_PageResponse_GetHomePage_Flat)ToPB() *PB_PageResponse_GetHomePage {
r := &PB_PageResponse_GetHomePage{


}
return r
}

func(m *PB_PageParam_GetProfilePage_Flat)ToPB() *PB_PageParam_GetProfilePage {
r := &PB_PageParam_GetProfilePage{
}
return r
}

func(m *PB_PageResponse_GetProfilePage_Flat)ToPB() *PB_PageResponse_GetProfilePage {
r := &PB_PageResponse_GetProfilePage{

}
return r
}

func(m *PB_PageParam_GetLikesPage_Flat)ToPB() *PB_PageParam_GetLikesPage {
r := &PB_PageParam_GetLikesPage{
    PostId:  int64(m.PostId) ,

}
return r
}

func(m *PB_PageResponse_GetLikesPage_Flat)ToPB() *PB_PageResponse_GetLikesPage {
r := &PB_PageResponse_GetLikesPage{


}
return r
}

func(m *PB_PageParam_GetFollowersPage_Flat)ToPB() *PB_PageParam_GetFollowersPage {
r := &PB_PageParam_GetFollowersPage{
    UserId:  int64(m.UserId) ,

}
return r
}

func(m *PB_PageResponse_GetFollowersPage_Flat)ToPB() *PB_PageResponse_GetFollowersPage {
r := &PB_PageResponse_GetFollowersPage{


}
return r
}

func(m *PB_PageParam_GetFollowingsPage_Flat)ToPB() *PB_PageParam_GetFollowingsPage {
r := &PB_PageParam_GetFollowingsPage{
    UserId:  int64(m.UserId) ,

}
return r
}

func(m *PB_PageResponse_GetFollowingsPage_Flat)ToPB() *PB_PageResponse_GetFollowingsPage {
r := &PB_PageResponse_GetFollowingsPage{


}
return r
}

func(m *PB_PageParam_GetNotifiesPage_Flat)ToPB() *PB_PageParam_GetNotifiesPage {
r := &PB_PageParam_GetNotifiesPage{

}
return r
}

func(m *PB_PageResponse_GetNotifiesPage_Flat)ToPB() *PB_PageResponse_GetNotifiesPage {
r := &PB_PageResponse_GetNotifiesPage{


    RemoveIdsList:  helper.SliceIntToInt64(m.RemoveIdsList) ,
}
return r
}

func(m *PB_PageParam_GetUserActionsPage_Flat)ToPB() *PB_PageParam_GetUserActionsPage {
r := &PB_PageParam_GetUserActionsPage{

}
return r
}

func(m *PB_PageResponse_GetUserActionsPage_Flat)ToPB() *PB_PageResponse_GetUserActionsPage {
r := &PB_PageResponse_GetUserActionsPage{


}
return r
}

func(m *PB_PageParam_GetSuggestedPostsPage_Flat)ToPB() *PB_PageParam_GetSuggestedPostsPage {
r := &PB_PageParam_GetSuggestedPostsPage{

}
return r
}

func(m *PB_PageResponse_GetSuggestedPostsPage_Flat)ToPB() *PB_PageResponse_GetSuggestedPostsPage {
r := &PB_PageResponse_GetSuggestedPostsPage{


}
return r
}

func(m *PB_PageParam_GetSuggestedUsrsPage_Flat)ToPB() *PB_PageParam_GetSuggestedUsrsPage {
r := &PB_PageParam_GetSuggestedUsrsPage{

}
return r
}

func(m *PB_PageResponse_GetSuggestedUsrsPage_Flat)ToPB() *PB_PageResponse_GetSuggestedUsrsPage {
r := &PB_PageResponse_GetSuggestedUsrsPage{


}
return r
}

func(m *PB_PageParam_GetSuggestedTagsPage_Flat)ToPB() *PB_PageParam_GetSuggestedTagsPage {
r := &PB_PageParam_GetSuggestedTagsPage{

}
return r
}

func(m *PB_PageResponse_GetSuggestedTagsPage_Flat)ToPB() *PB_PageResponse_GetSuggestedTagsPage {
r := &PB_PageResponse_GetSuggestedTagsPage{

}
return r
}

func(m *PB_PageParam_GetLastPostsPage_Flat)ToPB() *PB_PageParam_GetLastPostsPage {
r := &PB_PageParam_GetLastPostsPage{

}
return r
}

func(m *PB_PageResponse_GetLastPostsPage_Flat)ToPB() *PB_PageResponse_GetLastPostsPage {
r := &PB_PageResponse_GetLastPostsPage{


}
return r
}

func(m *PB_PageParam_GetTagPage_Flat)ToPB() *PB_PageParam_GetTagPage {
r := &PB_PageParam_GetTagPage{

}
return r
}

func(m *PB_PageResponse_GetTagPage_Flat)ToPB() *PB_PageResponse_GetTagPage {
r := &PB_PageResponse_GetTagPage{


}
return r
}

func(m *PB_PageParam_SearchTagsPage_Flat)ToPB() *PB_PageParam_SearchTagsPage {
r := &PB_PageParam_SearchTagsPage{
}
return r
}

func(m *PB_PageResponse_SearchTagsPage_Flat)ToPB() *PB_PageResponse_SearchTagsPage {
r := &PB_PageResponse_SearchTagsPage{

}
return r
}

func(m *PB_PageParam_SearchUsersPage_Flat)ToPB() *PB_PageParam_SearchUsersPage {
r := &PB_PageParam_SearchUsersPage{
}
return r
}

func(m *PB_PageResponse_SearchUsersPage_Flat)ToPB() *PB_PageResponse_SearchUsersPage {
r := &PB_PageResponse_SearchUsersPage{

}
return r
}

func(m *PB_PageParam__Flat)ToPB() *PB_PageParam_ {
r := &PB_PageParam_{
}
return r
}

func(m *PB_PageResponse__Flat)ToPB() *PB_PageResponse_ {
r := &PB_PageResponse_{

}
return r
}

func(m *PB_SearchResponse_AddNewC_Flat)ToPB() *PB_SearchResponse_AddNewC {
r := &PB_SearchResponse_AddNewC{
}
return r
}

func(m *PB_SocialParam_AddComment_Flat)ToPB() *PB_SocialParam_AddComment {
r := &PB_SocialParam_AddComment{
    PostId:  int64(m.PostId) ,
    Text:  m.Text ,
}
return r
}

func(m *PB_SocialResponse_AddComment_Flat)ToPB() *PB_SocialResponse_AddComment {
r := &PB_SocialResponse_AddComment{


}
return r
}

func(m *PB_SocialParam_DeleteComment_Flat)ToPB() *PB_SocialParam_DeleteComment {
r := &PB_SocialParam_DeleteComment{
    PostId:  int64(m.PostId) ,
    CommentId:  int64(m.CommentId) ,
}
return r
}

func(m *PB_SocialResponse_DeleteComment_Flat)ToPB() *PB_SocialResponse_DeleteComment {
r := &PB_SocialResponse_DeleteComment{

    Deleted:  m.Deleted ,
}
return r
}

func(m *PB_SocialParam_AddPost_Flat)ToPB() *PB_SocialParam_AddPost {
r := &PB_SocialParam_AddPost{
    Text:  m.Text ,
}
return r
}

func(m *PB_SocialResponse_AddPost_Flat)ToPB() *PB_SocialResponse_AddPost {
r := &PB_SocialResponse_AddPost{


}
return r
}

func(m *PB_SocialParam_EditPost_Flat)ToPB() *PB_SocialParam_EditPost {
r := &PB_SocialParam_EditPost{
    PostId:  int64(m.PostId) ,
    Text:  m.Text ,
}
return r
}

func(m *PB_SocialResponse_EditPost_Flat)ToPB() *PB_SocialResponse_EditPost {
r := &PB_SocialResponse_EditPost{


}
return r
}

func(m *PB_SocialParam_DeletePost_Flat)ToPB() *PB_SocialParam_DeletePost {
r := &PB_SocialParam_DeletePost{
    PostId:  int64(m.PostId) ,
}
return r
}

func(m *PB_SocialResponse_DeletePost_Flat)ToPB() *PB_SocialResponse_DeletePost {
r := &PB_SocialResponse_DeletePost{

    Done:  m.Done ,
}
return r
}

func(m *PB_SocialParam_ArchivePost_Flat)ToPB() *PB_SocialParam_ArchivePost {
r := &PB_SocialParam_ArchivePost{
    PostId:  int64(m.PostId) ,
}
return r
}

func(m *PB_SocialResponse_ArchivePost_Flat)ToPB() *PB_SocialResponse_ArchivePost {
r := &PB_SocialResponse_ArchivePost{

    Done:  m.Done ,
}
return r
}

func(m *PB_SocialParam_LikePost_Flat)ToPB() *PB_SocialParam_LikePost {
r := &PB_SocialParam_LikePost{
    PostId:  int64(m.PostId) ,
}
return r
}

func(m *PB_SocialResponse_LikePost_Flat)ToPB() *PB_SocialResponse_LikePost {
r := &PB_SocialResponse_LikePost{

    Done:  m.Done ,
}
return r
}

func(m *PB_SocialParam_UnLikePost_Flat)ToPB() *PB_SocialParam_UnLikePost {
r := &PB_SocialParam_UnLikePost{
    PostId:  int64(m.PostId) ,
}
return r
}

func(m *PB_SocialResponse_UnLikePost_Flat)ToPB() *PB_SocialResponse_UnLikePost {
r := &PB_SocialResponse_UnLikePost{

    Done:  m.Done ,
}
return r
}

func(m *PB_SocialParam_FollowUser_Flat)ToPB() *PB_SocialParam_FollowUser {
r := &PB_SocialParam_FollowUser{
    UserId:  int64(m.UserId) ,
}
return r
}

func(m *PB_SocialResponse_FollowUser_Flat)ToPB() *PB_SocialResponse_FollowUser {
r := &PB_SocialResponse_FollowUser{

}
return r
}

func(m *PB_SocialParam_UnFollowUser_Flat)ToPB() *PB_SocialParam_UnFollowUser {
r := &PB_SocialParam_UnFollowUser{
    UserId:  int64(m.UserId) ,
}
return r
}

func(m *PB_SocialResponse_UnFollowUser_Flat)ToPB() *PB_SocialResponse_UnFollowUser {
r := &PB_SocialResponse_UnFollowUser{

}
return r
}

func(m *PB_UserParam_BlockUser_Flat)ToPB() *PB_UserParam_BlockUser {
r := &PB_UserParam_BlockUser{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserResponse_BlockUser_Flat)ToPB() *PB_UserResponse_BlockUser {
r := &PB_UserResponse_BlockUser{
    ByUserId:  int32(m.ByUserId) ,
    TargetUserId:  int32(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UnBlockUser_Flat)ToPB() *PB_UserParam_UnBlockUser {
r := &PB_UserParam_UnBlockUser{
    Offset:  int32(m.Offset) ,
    Limit:  int32(m.Limit) ,
}
return r
}

func(m *PB_UserResponse_UnBlockUser_Flat)ToPB() *PB_UserResponse_UnBlockUser {
r := &PB_UserResponse_UnBlockUser{

}
return r
}

func(m *PB_UserParam_BlockedList_Flat)ToPB() *PB_UserParam_BlockedList {
r := &PB_UserParam_BlockedList{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
}
return r
}

func(m *PB_UserResponse_BlockedList_Flat)ToPB() *PB_UserResponse_BlockedList {
r := &PB_UserResponse_BlockedList{
    ByUserId:  int32(m.ByUserId) ,
    TargetUserId:  int32(m.TargetUserId) ,
    TargetUserName:  m.TargetUserName ,
}
return r
}

func(m *PB_UserParam_UpdateAbout_Flat)ToPB() *PB_UserParam_UpdateAbout {
r := &PB_UserParam_UpdateAbout{
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserResponse_UpdateAbout_Flat)ToPB() *PB_UserResponse_UpdateAbout {
r := &PB_UserResponse_UpdateAbout{
    UserId:  int32(m.UserId) ,
    NewAbout:  m.NewAbout ,
}
return r
}

func(m *PB_UserParam_UpdateUserName_Flat)ToPB() *PB_UserParam_UpdateUserName {
r := &PB_UserParam_UpdateUserName{
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserResponse_UpdateUserName_Flat)ToPB() *PB_UserResponse_UpdateUserName {
r := &PB_UserResponse_UpdateUserName{
    UserId:  int32(m.UserId) ,
    NewUserName:  m.NewUserName ,
}
return r
}

func(m *PB_UserParam_ChangeAvatar_Flat)ToPB() *PB_UserParam_ChangeAvatar {
r := &PB_UserParam_ChangeAvatar{
    None:  m.None ,
    ImageData2:  m.ImageData2 ,
}
return r
}

func(m *PB_UserResponse_ChangeAvatar_Flat)ToPB() *PB_UserResponse_ChangeAvatar {
r := &PB_UserResponse_ChangeAvatar{
}
return r
}

func(m *PB_UserParam_ChangePrivacy_Flat)ToPB() *PB_UserParam_ChangePrivacy {
r := &PB_UserParam_ChangePrivacy{

}
return r
}

func(m *PB_UserResponseOffline_ChangePrivacy_Flat)ToPB() *PB_UserResponseOffline_ChangePrivacy {
r := &PB_UserResponseOffline_ChangePrivacy{
}
return r
}

func(m *PB_UserParam_CheckUserName_Flat)ToPB() *PB_UserParam_CheckUserName {
r := &PB_UserParam_CheckUserName{

}
return r
}

func(m *PB_UserResponse_CheckUserName_Flat)ToPB() *PB_UserResponse_CheckUserName {
r := &PB_UserResponse_CheckUserName{
}
return r
}

func(m *UserView_Flat)ToPB() *UserView {
r := &UserView{
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
    Seq:  int32(m.Seq) ,
}
return r
}

func(m *PB_Chat_Flat)ToPB() *PB_Chat {
r := &PB_Chat{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnum:  int32(m.RoomTypeEnum) ,
    UserId:  int32(m.UserId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    StartMessageIdFrom:  int64(m.StartMessageIdFrom) ,
    LastMessageId:  int64(m.LastMessageId) ,
    LastSeenMessageId:  int64(m.LastSeenMessageId) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
}
return r
}

func(m *PB_Comment_Flat)ToPB() *PB_Comment {
r := &PB_Comment{
    CommentId:  int64(m.CommentId) ,
    UserId:  int32(m.UserId) ,
    PostId:  int32(m.PostId) ,
    Text:  m.Text ,
    LikesCount:  int32(m.LikesCount) ,
    CreatedTime:  int32(m.CreatedTime) ,
    Seq:  int32(m.Seq) ,
}
return r
}

func(m *PB_DirectMessage_Flat)ToPB() *PB_DirectMessage {
r := &PB_DirectMessage{
    MessageId:  int64(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int32(m.UserId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessageTypeEnumId:  int32(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int32(m.CreatedSe) ,
    PeerReceivedTime:  int32(m.PeerReceivedTime) ,
    PeerSeenTime:  int32(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int32(m.DeliviryStatusEnumId) ,
}
return r
}

func(m *PB_DirectOffline_Flat)ToPB() *PB_DirectOffline {
r := &PB_DirectOffline{
    DirectOfflineId:  int64(m.DirectOfflineId) ,
    ToUserId:  int32(m.ToUserId) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int64(m.MessageId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    PBClass:  m.PBClass ,
    DataPB:  m.DataPB ,
    DataJson:  m.DataJson ,
    DataTemp:  m.DataTemp ,
    AtTimeMs:  int64(m.AtTimeMs) ,
}
return r
}

func(m *PB_DirectToMessage_Flat)ToPB() *PB_DirectToMessage {
r := &PB_DirectToMessage{
    Id:  int64(m.Id) ,
    ChatKey:  m.ChatKey ,
    MessageId:  int64(m.MessageId) ,
    SourceEnumId:  int32(m.SourceEnumId) ,
}
return r
}

func(m *PB_Feed_Flat)ToPB() *PB_Feed {
r := &PB_Feed{
    UserId:  int64(m.UserId) ,
    PostId:  int64(m.PostId) ,
    FeedId:  int32(m.FeedId) ,
}
return r
}

func(m *PB_FollowingList_Flat)ToPB() *PB_FollowingList {
r := &PB_FollowingList{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    ListType:  int32(m.ListType) ,
    Name:  m.Name ,
    Count:  int32(m.Count) ,
    IsAuto:  int32(m.IsAuto) ,
    IsPimiry:  int32(m.IsPimiry) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_FollowingListMember_Flat)ToPB() *PB_FollowingListMember {
r := &PB_FollowingListMember{
    Id:  int64(m.Id) ,
    ListId:  int32(m.ListId) ,
    UserId:  int32(m.UserId) ,
    FollowedUserId:  int32(m.FollowedUserId) ,
    FollowType:  int32(m.FollowType) ,
    UpdatedTimeMs:  int64(m.UpdatedTimeMs) ,
}
return r
}

func(m *PB_FollowingListMemberHistory_Flat)ToPB() *PB_FollowingListMemberHistory {
r := &PB_FollowingListMemberHistory{
    Id:  int64(m.Id) ,
    ListId:  int32(m.ListId) ,
    UserId:  int32(m.UserId) ,
    FollowedUserId:  int32(m.FollowedUserId) ,
    FollowType:  int32(m.FollowType) ,
    UpdatedTimeMs:  int64(m.UpdatedTimeMs) ,
    FollowId:  int32(m.FollowId) ,
}
return r
}

func(m *PB_GeneralLog_Flat)ToPB() *PB_GeneralLog {
r := &PB_GeneralLog{
    Id:  int64(m.Id) ,
    ToUserId:  int32(m.ToUserId) ,
    TargetId:  int32(m.TargetId) ,
    LogTypeId:  int32(m.LogTypeId) ,
    ExtraPB:  m.ExtraPB ,
    ExtraJson:  m.ExtraJson ,
    CreatedMs:  int64(m.CreatedMs) ,
}
return r
}

func(m *PB_Group_Flat)ToPB() *PB_Group {
r := &PB_Group{
    GroupId:  int64(m.GroupId) ,
    GroupName:  m.GroupName ,
    MembersCount:  int32(m.MembersCount) ,
    GroupPrivacyEnum:  int32(m.GroupPrivacyEnum) ,
    CreatorUserId:  int32(m.CreatorUserId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
    CurrentSeq:  int32(m.CurrentSeq) ,
}
return r
}

func(m *PB_GroupMember_Flat)ToPB() *PB_GroupMember {
r := &PB_GroupMember{
    Id:  int64(m.Id) ,
    GroupId:  int64(m.GroupId) ,
    GroupKey:  m.GroupKey ,
    UserId:  int32(m.UserId) ,
    ByUserId:  int32(m.ByUserId) ,
    GroupRoleEnumId:  int32(m.GroupRoleEnumId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_GroupMessage_Flat)ToPB() *PB_GroupMessage {
r := &PB_GroupMessage{
    MessageId:  int64(m.MessageId) ,
    RoomKey:  m.RoomKey ,
    UserId:  int32(m.UserId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessageTypeEnum:  int32(m.MessageTypeEnum) ,
    Text:  m.Text ,
    CreatedMs:  int64(m.CreatedMs) ,
    DeliveryStatusEnum:  int32(m.DeliveryStatusEnum) ,
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

func(m *PB_Media_Flat)ToPB() *PB_Media {
r := &PB_Media{
    MediaId:  int64(m.MediaId) ,
    UserId:  int32(m.UserId) ,
    PostId:  int32(m.PostId) ,
    AlbumId:  int32(m.AlbumId) ,
    MediaTypeEnum:  int32(m.MediaTypeEnum) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Size:  int32(m.Size) ,
    Duration:  int32(m.Duration) ,
    Md5Hash:  m.Md5Hash ,
    Color:  m.Color ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_MessageFile_Flat)ToPB() *PB_MessageFile {
r := &PB_MessageFile{
    MessageFileId:  int64(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    UserId:  int32(m.UserId) ,
    Title:  m.Title ,
    Size:  int32(m.Size) ,
    FileTypeEnum:  int32(m.FileTypeEnum) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    Md5Hash:  m.Md5Hash ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_Notify_Flat)ToPB() *PB_Notify {
r := &PB_Notify{
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
    Seq:  int32(m.Seq) ,
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
    PhoneDisplayName:  m.PhoneDisplayName ,
    PhoneFamilyName:  m.PhoneFamilyName ,
    PhoneNumber:  m.PhoneNumber ,
    PhoneNormalizedNumber:  m.PhoneNormalizedNumber ,
    PhoneContactRowId:  int32(m.PhoneContactRowId) ,
    UserId:  int32(m.UserId) ,
    DeviceUuidId:  int32(m.DeviceUuidId) ,
    CreatedTime:  int32(m.CreatedTime) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
}
return r
}

func(m *PB_Post_Flat)ToPB() *PB_Post {
r := &PB_Post{
    PostId:  int64(m.PostId) ,
    UserId:  int32(m.UserId) ,
    PostTypeEnum:  int32(m.PostTypeEnum) ,
    MediaId:  int32(m.MediaId) ,
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
}
return r
}

func(m *PB_RecommendUser_Flat)ToPB() *PB_RecommendUser {
r := &PB_RecommendUser{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    TargetId:  int32(m.TargetId) ,
    Weight:  m.Weight ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_SearchClicked_Flat)ToPB() *PB_SearchClicked {
r := &PB_SearchClicked{
    Id:  int64(m.Id) ,
    Query:  m.Query ,
    ClickType:  int32(m.ClickType) ,
    TargetId:  int32(m.TargetId) ,
    UserId:  int32(m.UserId) ,
    CreatedTime:  int32(m.CreatedTime) ,
}
return r
}

func(m *PB_Session_Flat)ToPB() *PB_Session {
r := &PB_Session{
    Id:  int64(m.Id) ,
    UserId:  int32(m.UserId) ,
    SessionUuid:  m.SessionUuid ,
    ClientUuid:  m.ClientUuid ,
    DeviceUuid:  m.DeviceUuid ,
    LastActivityTime:  int32(m.LastActivityTime) ,
    LastIpAddress:  m.LastIpAddress ,
    LastWifiMacAddress:  m.LastWifiMacAddress ,
    LastNetworkType:  m.LastNetworkType ,
    LastNetworkTypeEnumId:  int32(m.LastNetworkTypeEnumId) ,
    AppVersion:  int32(m.AppVersion) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
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

func(m *PB_TagsPost_Flat)ToPB() *PB_TagsPost {
r := &PB_TagsPost{
    Id:  int64(m.Id) ,
    TagId:  int32(m.TagId) ,
    PostId:  int32(m.PostId) ,
    PostTypeEnum:  int32(m.PostTypeEnum) ,
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
    UserTypeEnum:  int32(m.UserTypeEnum) ,
    UserLevelEnum:  int32(m.UserLevelEnum) ,
    AvatarId:  int64(m.AvatarId) ,
    ProfilePrivacyEnum:  int32(m.ProfilePrivacyEnum) ,
    Phone:  int64(m.Phone) ,
    About:  m.About ,
    Email:  m.Email ,
    PasswordHash:  m.PasswordHash ,
    PasswordSalt:  m.PasswordSalt ,
    FollowersCount:  int32(m.FollowersCount) ,
    FollowingCount:  int32(m.FollowingCount) ,
    PostsCount:  int32(m.PostsCount) ,
    MediaCount:  int32(m.MediaCount) ,
    LikesCount:  int32(m.LikesCount) ,
    ResharedCount:  int32(m.ResharedCount) ,
    LastActionTime:  int32(m.LastActionTime) ,
    LastPostTime:  int32(m.LastPostTime) ,
    PrimaryFollowingList:  int32(m.PrimaryFollowingList) ,
    CreatedSe:  int32(m.CreatedSe) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
    OnlinePrivacyEnum:  int32(m.OnlinePrivacyEnum) ,
    LastActivityTime:  int32(m.LastActivityTime) ,
    Phone2:  m.Phone2 ,
}
return r
}

func(m *PB_UserMetaInfo_Flat)ToPB() *PB_UserMetaInfo {
r := &PB_UserMetaInfo{
    Id:  int32(m.Id) ,
    UserId:  int32(m.UserId) ,
    IsNotificationDirty:  int32(m.IsNotificationDirty) ,
    LastUserRecGen:  int32(m.LastUserRecGen) ,
}
return r
}

func(m *PB_UserPassword_Flat)ToPB() *PB_UserPassword {
r := &PB_UserPassword{
    UserId:  int32(m.UserId) ,
    Password:  m.Password ,
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

func(m *PB_MediaView_Flat)ToPB() *PB_MediaView {
r := &PB_MediaView{
    MediaId:  int64(m.MediaId) ,
    UserId:  int32(m.UserId) ,
    PostId:  int32(m.PostId) ,
    AlbumId:  int32(m.AlbumId) ,
    MediaTypeEnum:  int32(m.MediaTypeEnum) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Size:  int32(m.Size) ,
    Duration:  int32(m.Duration) ,
    Color:  m.Color ,
    CreatedTime:  int32(m.CreatedTime) ,
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
    PostId:  int32(m.PostId) ,
    Text:  m.Text ,
    LikesCount:  int32(m.LikesCount) ,
    CreatedTime:  int32(m.CreatedTime) ,

}
return r
}

func(m *PB_UserView_Flat)ToPB() *PB_UserView {
r := &PB_UserView{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,


    AvatarId:  int64(m.AvatarId) ,
    ProfilePrivacyEnum:  int32(m.ProfilePrivacyEnum) ,
    Phone:  int64(m.Phone) ,
    About:  m.About ,
    FollowersCount:  int32(m.FollowersCount) ,
    FollowingCount:  int32(m.FollowingCount) ,
    PostsCount:  int32(m.PostsCount) ,
    MediaCount:  int32(m.MediaCount) ,
    LikesCount:  int32(m.LikesCount) ,
    ResharedCount:  int32(m.ResharedCount) ,

    LastActiveTime:  int32(m.LastActiveTime) ,

}
return r
}

func(m *PB_TopProfileView_Flat)ToPB() *PB_TopProfileView {
r := &PB_TopProfileView{

}
return r
}

func(m *PB_ChatView_Flat)ToPB() *PB_ChatView {
r := &PB_ChatView{
    ChatKey:  m.ChatKey ,
    RoomKey:  m.RoomKey ,
    RoomTypeEnumId:  int32(m.RoomTypeEnumId) ,
    UserId:  int32(m.UserId) ,
    PeerUserId:  int32(m.PeerUserId) ,
    GroupId:  int64(m.GroupId) ,
    CreatedSe:  int32(m.CreatedSe) ,
    UpdatedMs:  int64(m.UpdatedMs) ,
    LastMessageId:  int64(m.LastMessageId) ,
    LastDeletedMessageId:  int64(m.LastDeletedMessageId) ,
    LastSeenMessageId:  int64(m.LastSeenMessageId) ,
    LastSeqSeen:  int32(m.LastSeqSeen) ,
    LastSeqDelete:  int32(m.LastSeqDelete) ,
    CurrentSeq:  int32(m.CurrentSeq) ,

    SharedMediaCount:  int32(m.SharedMediaCount) ,
    UnseenCount:  int32(m.UnseenCount) ,


}
return r
}

func(m *PB_MessageView_Flat)ToPB() *PB_MessageView {
r := &PB_MessageView{
    MessageId:  int64(m.MessageId) ,
    MessageKey:  m.MessageKey ,
    RoomKey:  m.RoomKey ,
    UserId:  int32(m.UserId) ,
    MessageFileId:  int64(m.MessageFileId) ,
    MessageTypeEnumId:  int32(m.MessageTypeEnumId) ,
    Text:  m.Text ,
    CreatedSe:  int32(m.CreatedSe) ,
    PeerReceivedTime:  int32(m.PeerReceivedTime) ,
    PeerSeenTime:  int32(m.PeerSeenTime) ,
    DeliviryStatusEnumId:  int32(m.DeliviryStatusEnumId) ,
    ChatKey:  m.ChatKey ,
    RoomTypeEnumId:  int32(m.RoomTypeEnumId) ,
    IsByMe:  m.IsByMe ,
    RemoteId:  int64(m.RemoteId) ,

}
return r
}

func(m *PB_MessageFileView_Flat)ToPB() *PB_MessageFileView {
r := &PB_MessageFileView{
    MessageFileId:  int64(m.MessageFileId) ,
    MessageFileKey:  m.MessageFileKey ,
    OriginalUserId:  int32(m.OriginalUserId) ,
    Name:  m.Name ,
    Size:  int32(m.Size) ,
    FileTypeEnumId:  int32(m.FileTypeEnumId) ,
    Width:  int32(m.Width) ,
    Height:  int32(m.Height) ,
    Duration:  int32(m.Duration) ,
    Extension:  m.Extension ,
    HashMd5:  m.HashMd5 ,
    HashAccess:  int64(m.HashAccess) ,
    CreatedSe:  int32(m.CreatedSe) ,
    ServerSrc:  m.ServerSrc ,
    ServerPath:  m.ServerPath ,
    ServerThumbPath:  m.ServerThumbPath ,
    BucketId:  m.BucketId ,
    ServerId:  int32(m.ServerId) ,
    CanDel:  int32(m.CanDel) ,
    ServerThumbLocalSrc:  m.ServerThumbLocalSrc ,
    RemoteMessageFileId:  int64(m.RemoteMessageFileId) ,
    LocalSrc:  m.LocalSrc ,
    ThumbLocalSrc:  m.ThumbLocalSrc ,
    MessageFileStatusId:  m.MessageFileStatusId ,
}
return r
}

func(m *PB_UserView3_Flat)ToPB() *PB_UserView3 {
r := &PB_UserView3{
    UserId:  int32(m.UserId) ,
    UserName:  m.UserName ,
    FirstName:  m.FirstName ,
    LastName:  m.LastName ,
    About:  m.About ,
    FullName:  m.FullName ,
    AvatarUrl:  m.AvatarUrl ,
    PrivacyProfile:  int32(m.PrivacyProfile) ,
    IsDeleted:  int32(m.IsDeleted) ,
    FollowersCount:  int32(m.FollowersCount) ,
    FollowingCount:  int32(m.FollowingCount) ,
    PostsCount:  int32(m.PostsCount) ,
    UpdatedTime:  int32(m.UpdatedTime) ,
    AppVersion:  int32(m.AppVersion) ,
    LastActivityTime:  int32(m.LastActivityTime) ,
    FollowingType:  int32(m.FollowingType) ,
}
return r
}



///// folding ///

var GeoLocation__FOlD = &GeoLocation{
        Lat:  0.0 ,
        Lon:  0.0 ,
}


var RoomMessageLog__FOlD = &RoomMessageLog{

        TargetUserId:  0 ,
        ByUserId:  0 ,
}


var RoomMessageForwardFrom__FOlD = &RoomMessageForwardFrom{
        RoomId:  0 ,
        MessageId:  0 ,
        RoomTypeEnum:  0 ,
}


var RoomDraft__FOlD = &RoomDraft{
        Message:  "" ,
        ReplyTo:  0 ,
}


var ChatRoom__FOlD = &ChatRoom{
}


var Pagination__FOlD = &Pagination{
        Offset:  0 ,
        Limit:  0 ,
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


var PB_ResponseExtra__FOlD = &PB_ResponseExtra{
        ErrorCode:  0 ,
        ErrMessage:  "" ,
        RpcFullName:  "" ,
        Data:  []byte{} ,
}


var PB_Pager__FOlD = &PB_Pager{
        Page:  0 ,
        Limit:  0 ,
        Last:  0 ,
}


var PB_UserParam_CheckUserName2__FOlD = &PB_UserParam_CheckUserName2{
}


var PB_UserResponse_CheckUserName2__FOlD = &PB_UserResponse_CheckUserName2{
}


var PB_ChatParam_AddNewMessage__FOlD = &PB_ChatParam_AddNewMessage{
        ReplyToMessageId:  0 ,
        Blob:  []byte{} ,

}


var PB_ChatResponse_AddNewMessage__FOlD = &PB_ChatResponse_AddNewMessage{
        MessageKey:  "" ,
        NewMessageId:  0 ,
        MessageFileKey:  "" ,
        NewMessageFileId:  0 ,
        AtTime:  0 ,
}


var PB_ChatParam_SetRoomActionDoing__FOlD = &PB_ChatParam_SetRoomActionDoing{
        GroupId:  0 ,
        DirectRoomKey:  "" ,

}


var PB_ChatResponse_SetRoomActionDoing__FOlD = &PB_ChatResponse_SetRoomActionDoing{
}


var PB_ChatParam_SetChatMessagesRangeAsSeen__FOlD = &PB_ChatParam_SetChatMessagesRangeAsSeen{
        ChatKey:  "" ,
        BottomMessageId:  0 ,
        TopMessageId:  0 ,
        SeenTimeMs:  0 ,
}


var PB_ChatResponse_SetChatMessagesRangeAsSeen__FOlD = &PB_ChatResponse_SetChatMessagesRangeAsSeen{
}


var PB_ChatParam_DeleteChatHistory__FOlD = &PB_ChatParam_DeleteChatHistory{
        ChatKey:  "" ,
        FromMessageId:  0 ,
}


var PB_ChatResponse_DeleteChatHistory__FOlD = &PB_ChatResponse_DeleteChatHistory{
}


var PB_ChatParam_DeleteMessagesByIds__FOlD = &PB_ChatParam_DeleteMessagesByIds{
        ChatKey:  "" ,
        Both:  false ,
        MessagesIds:  0 ,
}


var PB_ChatResponse_DeleteMessagesByIds__FOlD = &PB_ChatResponse_DeleteMessagesByIds{
}


var PB_ChatParam_SetMessagesAsReceived__FOlD = &PB_ChatParam_SetMessagesAsReceived{
        ChatRoom:  "" ,
        MessageIds:  0 ,
}


var PB_ChatResponse_SetMessagesAsReceived__FOlD = &PB_ChatResponse_SetMessagesAsReceived{
}


var PB_ChatParam_EditMessage__FOlD = &PB_ChatParam_EditMessage{
        RoomKey:  "" ,
        MessageId:  0 ,
        NewText:  "" ,
}


var PB_ChatResponse_EditMessage__FOlD = &PB_ChatResponse_EditMessage{
}


var PB_ChatParam_GetChatList__FOlD = &PB_ChatParam_GetChatList{
}


var PB_ChatResponse_GetChatList__FOlD = &PB_ChatResponse_GetChatList{


}


var PB_ChatParam_GetChatHistoryToOlder__FOlD = &PB_ChatParam_GetChatHistoryToOlder{
        ChatKey:  "" ,
        Limit:  0 ,
        FromTopMessageId:  0 ,
}


var PB_ChatResponse_GetChatHistoryToOlder__FOlD = &PB_ChatResponse_GetChatHistoryToOlder{

        HasMore:  false ,
}


var PB_ChatParam_GetFreshAllDirectMessagesList__FOlD = &PB_ChatParam_GetFreshAllDirectMessagesList{
        LowMessageId:  0 ,
        HighMessageId:  0 ,
}


var PB_ChatResponse_GetFreshAllDirectMessagesList__FOlD = &PB_ChatResponse_GetFreshAllDirectMessagesList{

        HasMore:  false ,
}


var PB_OtherParam_Echo__FOlD = &PB_OtherParam_Echo{
        Text:  "" ,
}


var PB_OtherResponse_Echo__FOlD = &PB_OtherResponse_Echo{
        Text:  "" ,
}


var PB_PageParam_GetCommentsPage__FOlD = &PB_PageParam_GetCommentsPage{
        PostId:  0 ,

}


var PB_PageResponse_GetCommentsPage__FOlD = &PB_PageResponse_GetCommentsPage{


}


var PB_PageParam_GetHomePage__FOlD = &PB_PageParam_GetHomePage{

}


var PB_PageResponse_GetHomePage__FOlD = &PB_PageResponse_GetHomePage{


}


var PB_PageParam_GetProfilePage__FOlD = &PB_PageParam_GetProfilePage{
}


var PB_PageResponse_GetProfilePage__FOlD = &PB_PageResponse_GetProfilePage{

}


var PB_PageParam_GetLikesPage__FOlD = &PB_PageParam_GetLikesPage{
        PostId:  0 ,

}


var PB_PageResponse_GetLikesPage__FOlD = &PB_PageResponse_GetLikesPage{


}


var PB_PageParam_GetFollowersPage__FOlD = &PB_PageParam_GetFollowersPage{
        UserId:  0 ,

}


var PB_PageResponse_GetFollowersPage__FOlD = &PB_PageResponse_GetFollowersPage{


}


var PB_PageParam_GetFollowingsPage__FOlD = &PB_PageParam_GetFollowingsPage{
        UserId:  0 ,

}


var PB_PageResponse_GetFollowingsPage__FOlD = &PB_PageResponse_GetFollowingsPage{


}


var PB_PageParam_GetNotifiesPage__FOlD = &PB_PageParam_GetNotifiesPage{

}


var PB_PageResponse_GetNotifiesPage__FOlD = &PB_PageResponse_GetNotifiesPage{


        RemoveIdsList:  0 ,
}


var PB_PageParam_GetUserActionsPage__FOlD = &PB_PageParam_GetUserActionsPage{

}


var PB_PageResponse_GetUserActionsPage__FOlD = &PB_PageResponse_GetUserActionsPage{


}


var PB_PageParam_GetSuggestedPostsPage__FOlD = &PB_PageParam_GetSuggestedPostsPage{

}


var PB_PageResponse_GetSuggestedPostsPage__FOlD = &PB_PageResponse_GetSuggestedPostsPage{


}


var PB_PageParam_GetSuggestedUsrsPage__FOlD = &PB_PageParam_GetSuggestedUsrsPage{

}


var PB_PageResponse_GetSuggestedUsrsPage__FOlD = &PB_PageResponse_GetSuggestedUsrsPage{


}


var PB_PageParam_GetSuggestedTagsPage__FOlD = &PB_PageParam_GetSuggestedTagsPage{

}


var PB_PageResponse_GetSuggestedTagsPage__FOlD = &PB_PageResponse_GetSuggestedTagsPage{

}


var PB_PageParam_GetLastPostsPage__FOlD = &PB_PageParam_GetLastPostsPage{

}


var PB_PageResponse_GetLastPostsPage__FOlD = &PB_PageResponse_GetLastPostsPage{


}


var PB_PageParam_GetTagPage__FOlD = &PB_PageParam_GetTagPage{

}


var PB_PageResponse_GetTagPage__FOlD = &PB_PageResponse_GetTagPage{


}


var PB_PageParam_SearchTagsPage__FOlD = &PB_PageParam_SearchTagsPage{
}


var PB_PageResponse_SearchTagsPage__FOlD = &PB_PageResponse_SearchTagsPage{

}


var PB_PageParam_SearchUsersPage__FOlD = &PB_PageParam_SearchUsersPage{
}


var PB_PageResponse_SearchUsersPage__FOlD = &PB_PageResponse_SearchUsersPage{

}


var PB_PageParam___FOlD = &PB_PageParam_{
}


var PB_PageResponse___FOlD = &PB_PageResponse_{

}


var PB_SearchResponse_AddNewC__FOlD = &PB_SearchResponse_AddNewC{
}


var PB_SocialParam_AddComment__FOlD = &PB_SocialParam_AddComment{
        PostId:  0 ,
        Text:  "" ,
}


var PB_SocialResponse_AddComment__FOlD = &PB_SocialResponse_AddComment{


}


var PB_SocialParam_DeleteComment__FOlD = &PB_SocialParam_DeleteComment{
        PostId:  0 ,
        CommentId:  0 ,
}


var PB_SocialResponse_DeleteComment__FOlD = &PB_SocialResponse_DeleteComment{

        Deleted:  false ,
}


var PB_SocialParam_AddPost__FOlD = &PB_SocialParam_AddPost{
        Text:  "" ,
}


var PB_SocialResponse_AddPost__FOlD = &PB_SocialResponse_AddPost{


}


var PB_SocialParam_EditPost__FOlD = &PB_SocialParam_EditPost{
        PostId:  0 ,
        Text:  "" ,
}


var PB_SocialResponse_EditPost__FOlD = &PB_SocialResponse_EditPost{


}


var PB_SocialParam_DeletePost__FOlD = &PB_SocialParam_DeletePost{
        PostId:  0 ,
}


var PB_SocialResponse_DeletePost__FOlD = &PB_SocialResponse_DeletePost{

        Done:  false ,
}


var PB_SocialParam_ArchivePost__FOlD = &PB_SocialParam_ArchivePost{
        PostId:  0 ,
}


var PB_SocialResponse_ArchivePost__FOlD = &PB_SocialResponse_ArchivePost{

        Done:  false ,
}


var PB_SocialParam_LikePost__FOlD = &PB_SocialParam_LikePost{
        PostId:  0 ,
}


var PB_SocialResponse_LikePost__FOlD = &PB_SocialResponse_LikePost{

        Done:  false ,
}


var PB_SocialParam_UnLikePost__FOlD = &PB_SocialParam_UnLikePost{
        PostId:  0 ,
}


var PB_SocialResponse_UnLikePost__FOlD = &PB_SocialResponse_UnLikePost{

        Done:  false ,
}


var PB_SocialParam_FollowUser__FOlD = &PB_SocialParam_FollowUser{
        UserId:  0 ,
}


var PB_SocialResponse_FollowUser__FOlD = &PB_SocialResponse_FollowUser{

}


var PB_SocialParam_UnFollowUser__FOlD = &PB_SocialParam_UnFollowUser{
        UserId:  0 ,
}


var PB_SocialResponse_UnFollowUser__FOlD = &PB_SocialResponse_UnFollowUser{

}


var PB_UserParam_BlockUser__FOlD = &PB_UserParam_BlockUser{
        UserId:  0 ,
        UserName:  "" ,
}


var PB_UserResponse_BlockUser__FOlD = &PB_UserResponse_BlockUser{
        ByUserId:  0 ,
        TargetUserId:  0 ,
        TargetUserName:  "" ,
}


var PB_UserParam_UnBlockUser__FOlD = &PB_UserParam_UnBlockUser{
        Offset:  0 ,
        Limit:  0 ,
}


var PB_UserResponse_UnBlockUser__FOlD = &PB_UserResponse_UnBlockUser{

}


var PB_UserParam_BlockedList__FOlD = &PB_UserParam_BlockedList{
        UserId:  0 ,
        UserName:  "" ,
}


var PB_UserResponse_BlockedList__FOlD = &PB_UserResponse_BlockedList{
        ByUserId:  0 ,
        TargetUserId:  0 ,
        TargetUserName:  "" ,
}


var PB_UserParam_UpdateAbout__FOlD = &PB_UserParam_UpdateAbout{
        NewAbout:  "" ,
}


var PB_UserResponse_UpdateAbout__FOlD = &PB_UserResponse_UpdateAbout{
        UserId:  0 ,
        NewAbout:  "" ,
}


var PB_UserParam_UpdateUserName__FOlD = &PB_UserParam_UpdateUserName{
        NewUserName:  "" ,
}


var PB_UserResponse_UpdateUserName__FOlD = &PB_UserResponse_UpdateUserName{
        UserId:  0 ,
        NewUserName:  "" ,
}


var PB_UserParam_ChangeAvatar__FOlD = &PB_UserParam_ChangeAvatar{
        None:  false ,
        ImageData2:  []byte{} ,
}


var PB_UserResponse_ChangeAvatar__FOlD = &PB_UserResponse_ChangeAvatar{
}


var PB_UserParam_ChangePrivacy__FOlD = &PB_UserParam_ChangePrivacy{

}


var PB_UserResponseOffline_ChangePrivacy__FOlD = &PB_UserResponseOffline_ChangePrivacy{
}


var PB_UserParam_CheckUserName__FOlD = &PB_UserParam_CheckUserName{

}


var PB_UserResponse_CheckUserName__FOlD = &PB_UserResponse_CheckUserName{
}


var UserView__FOlD = &UserView{
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
        Seq:  0 ,
}


var PB_Chat__FOlD = &PB_Chat{
        ChatKey:  "" ,
        RoomKey:  "" ,
        RoomTypeEnum:  0 ,
        UserId:  0 ,
        PeerUserId:  0 ,
        GroupId:  0 ,
        CreatedTime:  0 ,
        StartMessageIdFrom:  0 ,
        LastMessageId:  0 ,
        LastSeenMessageId:  0 ,
        UpdatedMs:  0 ,
}


var PB_Comment__FOlD = &PB_Comment{
        CommentId:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        Text:  "" ,
        LikesCount:  0 ,
        CreatedTime:  0 ,
        Seq:  0 ,
}


var PB_DirectMessage__FOlD = &PB_DirectMessage{
        MessageId:  0 ,
        MessageKey:  "" ,
        RoomKey:  "" ,
        UserId:  0 ,
        MessageFileId:  0 ,
        MessageTypeEnumId:  0 ,
        Text:  "" ,
        CreatedSe:  0 ,
        PeerReceivedTime:  0 ,
        PeerSeenTime:  0 ,
        DeliviryStatusEnumId:  0 ,
}


var PB_DirectOffline__FOlD = &PB_DirectOffline{
        DirectOfflineId:  0 ,
        ToUserId:  0 ,
        ChatKey:  "" ,
        MessageId:  0 ,
        MessageFileId:  0 ,
        PBClass:  "" ,
        DataPB:  []byte{} ,
        DataJson:  "" ,
        DataTemp:  "" ,
        AtTimeMs:  0 ,
}


var PB_DirectToMessage__FOlD = &PB_DirectToMessage{
        Id:  0 ,
        ChatKey:  "" ,
        MessageId:  0 ,
        SourceEnumId:  0 ,
}


var PB_Feed__FOlD = &PB_Feed{
        UserId:  0 ,
        PostId:  0 ,
        FeedId:  0 ,
}


var PB_FollowingList__FOlD = &PB_FollowingList{
        Id:  0 ,
        UserId:  0 ,
        ListType:  0 ,
        Name:  "" ,
        Count:  0 ,
        IsAuto:  0 ,
        IsPimiry:  0 ,
        CreatedTime:  0 ,
}


var PB_FollowingListMember__FOlD = &PB_FollowingListMember{
        Id:  0 ,
        ListId:  0 ,
        UserId:  0 ,
        FollowedUserId:  0 ,
        FollowType:  0 ,
        UpdatedTimeMs:  0 ,
}


var PB_FollowingListMemberHistory__FOlD = &PB_FollowingListMemberHistory{
        Id:  0 ,
        ListId:  0 ,
        UserId:  0 ,
        FollowedUserId:  0 ,
        FollowType:  0 ,
        UpdatedTimeMs:  0 ,
        FollowId:  0 ,
}


var PB_GeneralLog__FOlD = &PB_GeneralLog{
        Id:  0 ,
        ToUserId:  0 ,
        TargetId:  0 ,
        LogTypeId:  0 ,
        ExtraPB:  []byte{} ,
        ExtraJson:  "" ,
        CreatedMs:  0 ,
}


var PB_Group__FOlD = &PB_Group{
        GroupId:  0 ,
        GroupName:  "" ,
        MembersCount:  0 ,
        GroupPrivacyEnum:  0 ,
        CreatorUserId:  0 ,
        CreatedTime:  0 ,
        UpdatedMs:  0 ,
        CurrentSeq:  0 ,
}


var PB_GroupMember__FOlD = &PB_GroupMember{
        Id:  0 ,
        GroupId:  0 ,
        GroupKey:  "" ,
        UserId:  0 ,
        ByUserId:  0 ,
        GroupRoleEnumId:  0 ,
        CreatedTime:  0 ,
}


var PB_GroupMessage__FOlD = &PB_GroupMessage{
        MessageId:  0 ,
        RoomKey:  "" ,
        UserId:  0 ,
        MessageFileId:  0 ,
        MessageTypeEnum:  0 ,
        Text:  "" ,
        CreatedMs:  0 ,
        DeliveryStatusEnum:  0 ,
}


var PB_Like__FOlD = &PB_Like{
        Id:  0 ,
        PostId:  0 ,
        PostTypeEnum:  0 ,
        UserId:  0 ,
        LikeEnum:  0 ,
        CreatedTime:  0 ,
}


var PB_Media__FOlD = &PB_Media{
        MediaId:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        AlbumId:  0 ,
        MediaTypeEnum:  0 ,
        Width:  0 ,
        Height:  0 ,
        Size:  0 ,
        Duration:  0 ,
        Md5Hash:  "" ,
        Color:  "" ,
        CreatedTime:  0 ,
}


var PB_MessageFile__FOlD = &PB_MessageFile{
        MessageFileId:  0 ,
        MessageFileKey:  "" ,
        UserId:  0 ,
        Title:  "" ,
        Size:  0 ,
        FileTypeEnum:  0 ,
        Width:  0 ,
        Height:  0 ,
        Duration:  0 ,
        Extension:  "" ,
        Md5Hash:  "" ,
        CreatedTime:  0 ,
}


var PB_Notify__FOlD = &PB_Notify{
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
        Seq:  0 ,
}


var PB_NotifyRemoved__FOlD = &PB_NotifyRemoved{
        Murmur64Hash:  0 ,
        ForUserId:  0 ,
        Id:  0 ,
}


var PB_PhoneContact__FOlD = &PB_PhoneContact{
        Id:  0 ,
        PhoneDisplayName:  "" ,
        PhoneFamilyName:  "" ,
        PhoneNumber:  "" ,
        PhoneNormalizedNumber:  "" ,
        PhoneContactRowId:  0 ,
        UserId:  0 ,
        DeviceUuidId:  0 ,
        CreatedTime:  0 ,
        UpdatedTime:  0 ,
}


var PB_Post__FOlD = &PB_Post{
        PostId:  0 ,
        UserId:  0 ,
        PostTypeEnum:  0 ,
        MediaId:  0 ,
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
}


var PB_RecommendUser__FOlD = &PB_RecommendUser{
        Id:  0 ,
        UserId:  0 ,
        TargetId:  0 ,
        Weight:  0.0 ,
        CreatedTime:  0 ,
}


var PB_SearchClicked__FOlD = &PB_SearchClicked{
        Id:  0 ,
        Query:  "" ,
        ClickType:  0 ,
        TargetId:  0 ,
        UserId:  0 ,
        CreatedTime:  0 ,
}


var PB_Session__FOlD = &PB_Session{
        Id:  0 ,
        UserId:  0 ,
        SessionUuid:  "" ,
        ClientUuid:  "" ,
        DeviceUuid:  "" ,
        LastActivityTime:  0 ,
        LastIpAddress:  "" ,
        LastWifiMacAddress:  "" ,
        LastNetworkType:  "" ,
        LastNetworkTypeEnumId:  0 ,
        AppVersion:  0 ,
        UpdatedTime:  0 ,
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


var PB_TagsPost__FOlD = &PB_TagsPost{
        Id:  0 ,
        TagId:  0 ,
        PostId:  0 ,
        PostTypeEnum:  0 ,
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
        UserTypeEnum:  0 ,
        UserLevelEnum:  0 ,
        AvatarId:  0 ,
        ProfilePrivacyEnum:  0 ,
        Phone:  0 ,
        About:  "" ,
        Email:  "" ,
        PasswordHash:  "" ,
        PasswordSalt:  "" ,
        FollowersCount:  0 ,
        FollowingCount:  0 ,
        PostsCount:  0 ,
        MediaCount:  0 ,
        LikesCount:  0 ,
        ResharedCount:  0 ,
        LastActionTime:  0 ,
        LastPostTime:  0 ,
        PrimaryFollowingList:  0 ,
        CreatedSe:  0 ,
        UpdatedMs:  0 ,
        OnlinePrivacyEnum:  0 ,
        LastActivityTime:  0 ,
        Phone2:  "" ,
}


var PB_UserMetaInfo__FOlD = &PB_UserMetaInfo{
        Id:  0 ,
        UserId:  0 ,
        IsNotificationDirty:  0 ,
        LastUserRecGen:  0 ,
}


var PB_UserPassword__FOlD = &PB_UserPassword{
        UserId:  0 ,
        Password:  "" ,
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


var PB_MediaView__FOlD = &PB_MediaView{
        MediaId:  0 ,
        UserId:  0 ,
        PostId:  0 ,
        AlbumId:  0 ,
        MediaTypeEnum:  0 ,
        Width:  0 ,
        Height:  0 ,
        Size:  0 ,
        Duration:  0 ,
        Color:  "" ,
        CreatedTime:  0 ,
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


var PB_UserView__FOlD = &PB_UserView{
        UserId:  0 ,
        UserName:  "" ,
        FirstName:  "" ,
        LastName:  "" ,


        AvatarId:  0 ,
        ProfilePrivacyEnum:  0 ,
        Phone:  0 ,
        About:  "" ,
        FollowersCount:  0 ,
        FollowingCount:  0 ,
        PostsCount:  0 ,
        MediaCount:  0 ,
        LikesCount:  0 ,
        ResharedCount:  0 ,

        LastActiveTime:  0 ,

}


var PB_TopProfileView__FOlD = &PB_TopProfileView{

}


var PB_ChatView__FOlD = &PB_ChatView{
        ChatKey:  "" ,
        RoomKey:  "" ,
        RoomTypeEnumId:  0 ,
        UserId:  0 ,
        PeerUserId:  0 ,
        GroupId:  0 ,
        CreatedSe:  0 ,
        UpdatedMs:  0 ,
        LastMessageId:  0 ,
        LastDeletedMessageId:  0 ,
        LastSeenMessageId:  0 ,
        LastSeqSeen:  0 ,
        LastSeqDelete:  0 ,
        CurrentSeq:  0 ,

        SharedMediaCount:  0 ,
        UnseenCount:  0 ,


}


var PB_MessageView__FOlD = &PB_MessageView{
        MessageId:  0 ,
        MessageKey:  "" ,
        RoomKey:  "" ,
        UserId:  0 ,
        MessageFileId:  0 ,
        MessageTypeEnumId:  0 ,
        Text:  "" ,
        CreatedSe:  0 ,
        PeerReceivedTime:  0 ,
        PeerSeenTime:  0 ,
        DeliviryStatusEnumId:  0 ,
        ChatKey:  "" ,
        RoomTypeEnumId:  0 ,
        IsByMe:  false ,
        RemoteId:  0 ,

}


var PB_MessageFileView__FOlD = &PB_MessageFileView{
        MessageFileId:  0 ,
        MessageFileKey:  "" ,
        OriginalUserId:  0 ,
        Name:  "" ,
        Size:  0 ,
        FileTypeEnumId:  0 ,
        Width:  0 ,
        Height:  0 ,
        Duration:  0 ,
        Extension:  "" ,
        HashMd5:  "" ,
        HashAccess:  0 ,
        CreatedSe:  0 ,
        ServerSrc:  "" ,
        ServerPath:  "" ,
        ServerThumbPath:  "" ,
        BucketId:  "" ,
        ServerId:  0 ,
        CanDel:  0 ,
        ServerThumbLocalSrc:  "" ,
        RemoteMessageFileId:  0 ,
        LocalSrc:  "" ,
        ThumbLocalSrc:  "" ,
        MessageFileStatusId:  "" ,
}


var PB_UserView3__FOlD = &PB_UserView3{
        UserId:  0 ,
        UserName:  "" ,
        FirstName:  "" ,
        LastName:  "" ,
        About:  "" ,
        FullName:  "" ,
        AvatarUrl:  "" ,
        PrivacyProfile:  0 ,
        IsDeleted:  0 ,
        FollowersCount:  0 ,
        FollowingCount:  0 ,
        PostsCount:  0 ,
        UpdatedTime:  0 ,
        AppVersion:  0 ,
        LastActivityTime:  0 ,
        FollowingType:  0 ,
}



*/
