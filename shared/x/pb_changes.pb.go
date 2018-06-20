// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_changes.proto

/*
Package x is a generated protocol buffer package.

It is generated from these files:
	pb_changes.proto
	pb_enum.proto
	pb_global.proto
	pb_rpc_auth2.proto
	pb_rpc_chat.proto
	pb_rpc_general.proto
	pb_rpc_page.proto
	pb_rpc_social.proto
	pb_rpc_user.proto
	pb_tables.proto
	pb_views.proto

It has these top-level messages:
	PB_RoomsChanges
	PB_PushChanges
	PB_CommandToServer
	PB_CommandToClient
	PB_CommandReachedToServer
	PB_CommandReachedToClient
	PB_ResponseToClient
	RPC_Auth_Types
	PB_RPC_Chat_Types
	RPC_General_Types
	RPC_Page_Types
	RPC_Social_Types
	RPC_User_Types
	PB_Action
	PB_Blocked
	PB_Comment
	PB_CommentDeleted
	PB_Event
	PB_Followed
	PB_Likes
	PB_Notify
	PB_NotifyRemoved
	PB_PhoneContacts
	PB_Post
	PB_PostCount
	PB_PostDeleted
	PB_PostKeys
	PB_PostLink
	PB_PostMedia
	PB_PostReshared
	PB_ProfileAll
	PB_ProfileMedia
	PB_ProfileMentioned
	PB_Session
	PB_SettingNotifications
	PB_Sms
	PB_Tag
	PB_TagPost
	PB_TriggerLog
	PB_User
	PB_UserRelation
	PB_Chat
	PB_ChatDeleted
	PB_ChatLastMessage
	PB_ChatUserVersion
	PB_Group
	PB_GroupMember
	PB_GroupOrderdUser
	PB_GroupPinedMsg
	PB_FileMsg
	PB_FilePost
	PB_ActionFanout
	PB_HomeFanout
	PB_SuggestedTopPosts
	PB_SuggestedUser
	PB_PushChat
	PB_HTTPRPCLog
	PB_MetricLog
	PB_XfileServiceInfoLog
	PB_XfileServiceMetricLog
	PB_XfileServiceRequestLog
	PB_InvalidateCache
	PB_MediaView
	PB_ActionView
	PB_NotifyView
	PB_CommentView
	PB_PostView
	PB_ChatView
	PB_GroupView
	PB_GroupMemberView
	PB_MessageView
	PB_FileRedView
	PB_UserView
	PB_SettingNotificationView
	PB_AppConfig
	PB_UserProfileView
	PB_UserViewRowify
	PB_PostViewRowify
	PB_TagView
	PB_TopTagWithSamplePosts
	PB_SelfUserView
	PB_Error
*/
package x

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PB_RoomsChanges_Chat_Mute int32

const (
	PB_RoomsChanges_Chat_UNCHANGED_MUTE PB_RoomsChanges_Chat_Mute = 0
	PB_RoomsChanges_Chat_MUTED          PB_RoomsChanges_Chat_Mute = 1
	PB_RoomsChanges_Chat_UNMUTED        PB_RoomsChanges_Chat_Mute = 2
)

var PB_RoomsChanges_Chat_Mute_name = map[int32]string{
	0: "UNCHANGED_MUTE",
	1: "MUTED",
	2: "UNMUTED",
}
var PB_RoomsChanges_Chat_Mute_value = map[string]int32{
	"UNCHANGED_MUTE": 0,
	"MUTED":          1,
	"UNMUTED":        2,
}

func (x PB_RoomsChanges_Chat_Mute) String() string {
	return proto.EnumName(PB_RoomsChanges_Chat_Mute_name, int32(x))
}
func (PB_RoomsChanges_Chat_Mute) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

type PB_RoomsChanges_Chat_Pin int32

const (
	PB_RoomsChanges_Chat_UNCHANGED_PIN PB_RoomsChanges_Chat_Pin = 0
	PB_RoomsChanges_Chat_PINED         PB_RoomsChanges_Chat_Pin = 1
	PB_RoomsChanges_Chat_UNPINED       PB_RoomsChanges_Chat_Pin = 2
)

var PB_RoomsChanges_Chat_Pin_name = map[int32]string{
	0: "UNCHANGED_PIN",
	1: "PINED",
	2: "UNPINED",
}
var PB_RoomsChanges_Chat_Pin_value = map[string]int32{
	"UNCHANGED_PIN": 0,
	"PINED":         1,
	"UNPINED":       2,
}

func (x PB_RoomsChanges_Chat_Pin) String() string {
	return proto.EnumName(PB_RoomsChanges_Chat_Pin_name, int32(x))
}
func (PB_RoomsChanges_Chat_Pin) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 1}
}

type PB_RoomsChanges struct {
	VersionTime uint64                  `protobuf:"varint,1,opt,name=VersionTime" json:"VersionTime,omitempty"`
	Chats       []*PB_RoomsChanges_Chat `protobuf:"bytes,2,rep,name=Chats" json:"Chats,omitempty"`
	Rooms       []*PB_ChatView          `protobuf:"bytes,3,rep,name=rooms" json:"rooms,omitempty"`
}

func (m *PB_RoomsChanges) Reset()                    { *m = PB_RoomsChanges{} }
func (m *PB_RoomsChanges) String() string            { return proto.CompactTextString(m) }
func (*PB_RoomsChanges) ProtoMessage()               {}
func (*PB_RoomsChanges) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PB_RoomsChanges) GetVersionTime() uint64 {
	if m != nil {
		return m.VersionTime
	}
	return 0
}

func (m *PB_RoomsChanges) GetChats() []*PB_RoomsChanges_Chat {
	if m != nil {
		return m.Chats
	}
	return nil
}

func (m *PB_RoomsChanges) GetRooms() []*PB_ChatView {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type PB_RoomsChanges_Chat struct {
	ChatId                    int64                                 `protobuf:"varint,1,opt,name=ChatId" json:"ChatId,omitempty"`
	RoomKey                   string                                `protobuf:"bytes,3,opt,name=RoomKey" json:"RoomKey,omitempty"`
	RoomType                  int32                                 `protobuf:"varint,4,opt,name=RoomType" json:"RoomType,omitempty"`
	PeerPush                  int32                                 `protobuf:"varint,5,opt,name=PeerPush" json:"PeerPush,omitempty"`
	ReceivedMessages          []uint64                              `protobuf:"varint,10,rep,packed,name=ReceivedMessages" json:"ReceivedMessages,omitempty"`
	SeenMessages              []uint64                              `protobuf:"varint,11,rep,packed,name=SeenMessages" json:"SeenMessages,omitempty"`
	EditeMessages             []*PB_RoomsChanges_Chat_EditeMessage  `protobuf:"bytes,12,rep,name=EditeMessages" json:"EditeMessages,omitempty"`
	DeleteMessages            []*PB_RoomsChanges_Chat_DeleteMessage `protobuf:"bytes,13,rep,name=DeleteMessages" json:"DeleteMessages,omitempty"`
	ClearHistroyFromMessageId uint64                                `protobuf:"varint,14,opt,name=ClearHistroyFromMessageId" json:"ClearHistroyFromMessageId,omitempty"`
	DeleteChat                int32                                 `protobuf:"varint,15,opt,name=DeleteChat" json:"DeleteChat,omitempty"`
	ChatTitle                 string                                `protobuf:"bytes,16,opt,name=ChatTitle" json:"ChatTitle,omitempty"`
	Muted                     PB_RoomsChanges_Chat_Mute             `protobuf:"varint,20,opt,name=Muted,enum=PB_RoomsChanges_Chat_Mute" json:"Muted,omitempty"`
	MutedUntil                int32                                 `protobuf:"varint,21,opt,name=MutedUntil" json:"MutedUntil,omitempty"`
	Pined                     PB_RoomsChanges_Chat_Pin              `protobuf:"varint,30,opt,name=Pined,enum=PB_RoomsChanges_Chat_Pin" json:"Pined,omitempty"`
	PinTime                   int64                                 `protobuf:"varint,31,opt,name=PinTime" json:"PinTime,omitempty"`
}

func (m *PB_RoomsChanges_Chat) Reset()                    { *m = PB_RoomsChanges_Chat{} }
func (m *PB_RoomsChanges_Chat) String() string            { return proto.CompactTextString(m) }
func (*PB_RoomsChanges_Chat) ProtoMessage()               {}
func (*PB_RoomsChanges_Chat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *PB_RoomsChanges_Chat) GetChatId() int64 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *PB_RoomsChanges_Chat) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_RoomsChanges_Chat) GetRoomType() int32 {
	if m != nil {
		return m.RoomType
	}
	return 0
}

func (m *PB_RoomsChanges_Chat) GetPeerPush() int32 {
	if m != nil {
		return m.PeerPush
	}
	return 0
}

func (m *PB_RoomsChanges_Chat) GetReceivedMessages() []uint64 {
	if m != nil {
		return m.ReceivedMessages
	}
	return nil
}

func (m *PB_RoomsChanges_Chat) GetSeenMessages() []uint64 {
	if m != nil {
		return m.SeenMessages
	}
	return nil
}

func (m *PB_RoomsChanges_Chat) GetEditeMessages() []*PB_RoomsChanges_Chat_EditeMessage {
	if m != nil {
		return m.EditeMessages
	}
	return nil
}

func (m *PB_RoomsChanges_Chat) GetDeleteMessages() []*PB_RoomsChanges_Chat_DeleteMessage {
	if m != nil {
		return m.DeleteMessages
	}
	return nil
}

func (m *PB_RoomsChanges_Chat) GetClearHistroyFromMessageId() uint64 {
	if m != nil {
		return m.ClearHistroyFromMessageId
	}
	return 0
}

func (m *PB_RoomsChanges_Chat) GetDeleteChat() int32 {
	if m != nil {
		return m.DeleteChat
	}
	return 0
}

func (m *PB_RoomsChanges_Chat) GetChatTitle() string {
	if m != nil {
		return m.ChatTitle
	}
	return ""
}

func (m *PB_RoomsChanges_Chat) GetMuted() PB_RoomsChanges_Chat_Mute {
	if m != nil {
		return m.Muted
	}
	return PB_RoomsChanges_Chat_UNCHANGED_MUTE
}

func (m *PB_RoomsChanges_Chat) GetMutedUntil() int32 {
	if m != nil {
		return m.MutedUntil
	}
	return 0
}

func (m *PB_RoomsChanges_Chat) GetPined() PB_RoomsChanges_Chat_Pin {
	if m != nil {
		return m.Pined
	}
	return PB_RoomsChanges_Chat_UNCHANGED_PIN
}

func (m *PB_RoomsChanges_Chat) GetPinTime() int64 {
	if m != nil {
		return m.PinTime
	}
	return 0
}

type PB_RoomsChanges_Chat_EditeMessage struct {
	MessageId uint64 `protobuf:"varint,1,opt,name=MessageId" json:"MessageId,omitempty"`
	NewRext   string `protobuf:"bytes,2,opt,name=NewRext" json:"NewRext,omitempty"`
}

func (m *PB_RoomsChanges_Chat_EditeMessage) Reset()         { *m = PB_RoomsChanges_Chat_EditeMessage{} }
func (m *PB_RoomsChanges_Chat_EditeMessage) String() string { return proto.CompactTextString(m) }
func (*PB_RoomsChanges_Chat_EditeMessage) ProtoMessage()    {}
func (*PB_RoomsChanges_Chat_EditeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

func (m *PB_RoomsChanges_Chat_EditeMessage) GetMessageId() uint64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PB_RoomsChanges_Chat_EditeMessage) GetNewRext() string {
	if m != nil {
		return m.NewRext
	}
	return ""
}

type PB_RoomsChanges_Chat_DeleteMessage struct {
	MessageId uint64 `protobuf:"varint,1,opt,name=MessageId" json:"MessageId,omitempty"`
	Both      bool   `protobuf:"varint,2,opt,name=Both" json:"Both,omitempty"`
}

func (m *PB_RoomsChanges_Chat_DeleteMessage) Reset()         { *m = PB_RoomsChanges_Chat_DeleteMessage{} }
func (m *PB_RoomsChanges_Chat_DeleteMessage) String() string { return proto.CompactTextString(m) }
func (*PB_RoomsChanges_Chat_DeleteMessage) ProtoMessage()    {}
func (*PB_RoomsChanges_Chat_DeleteMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 1}
}

func (m *PB_RoomsChanges_Chat_DeleteMessage) GetMessageId() uint64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PB_RoomsChanges_Chat_DeleteMessage) GetBoth() bool {
	if m != nil {
		return m.Both
	}
	return false
}

type PB_PushChanges struct {
	RoomsChanges               []*PB_RoomsChanges `protobuf:"bytes,1,rep,name=RoomsChanges" json:"RoomsChanges,omitempty"`
	ChatView                   []*PB_ChatView     `protobuf:"bytes,2,rep,name=ChatView" json:"ChatView,omitempty"`
	InvalidateCacheForRoomKeys []string           `protobuf:"bytes,3,rep,name=InvalidateCacheForRoomKeys" json:"InvalidateCacheForRoomKeys,omitempty"`
	InvalidateAllChatCache     int32              `protobuf:"varint,10,opt,name=InvalidateAllChatCache" json:"InvalidateAllChatCache,omitempty"`
	InvalidateAllSocialCache   int32              `protobuf:"varint,11,opt,name=InvalidateAllSocialCache" json:"InvalidateAllSocialCache,omitempty"`
}

func (m *PB_PushChanges) Reset()                    { *m = PB_PushChanges{} }
func (m *PB_PushChanges) String() string            { return proto.CompactTextString(m) }
func (*PB_PushChanges) ProtoMessage()               {}
func (*PB_PushChanges) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PB_PushChanges) GetRoomsChanges() []*PB_RoomsChanges {
	if m != nil {
		return m.RoomsChanges
	}
	return nil
}

func (m *PB_PushChanges) GetChatView() []*PB_ChatView {
	if m != nil {
		return m.ChatView
	}
	return nil
}

func (m *PB_PushChanges) GetInvalidateCacheForRoomKeys() []string {
	if m != nil {
		return m.InvalidateCacheForRoomKeys
	}
	return nil
}

func (m *PB_PushChanges) GetInvalidateAllChatCache() int32 {
	if m != nil {
		return m.InvalidateAllChatCache
	}
	return 0
}

func (m *PB_PushChanges) GetInvalidateAllSocialCache() int32 {
	if m != nil {
		return m.InvalidateAllSocialCache
	}
	return 0
}

func init() {
	proto.RegisterType((*PB_RoomsChanges)(nil), "PB_RoomsChanges")
	proto.RegisterType((*PB_RoomsChanges_Chat)(nil), "PB_RoomsChanges.Chat")
	proto.RegisterType((*PB_RoomsChanges_Chat_EditeMessage)(nil), "PB_RoomsChanges.Chat.EditeMessage")
	proto.RegisterType((*PB_RoomsChanges_Chat_DeleteMessage)(nil), "PB_RoomsChanges.Chat.DeleteMessage")
	proto.RegisterType((*PB_PushChanges)(nil), "PB_PushChanges")
	proto.RegisterEnum("PB_RoomsChanges_Chat_Mute", PB_RoomsChanges_Chat_Mute_name, PB_RoomsChanges_Chat_Mute_value)
	proto.RegisterEnum("PB_RoomsChanges_Chat_Pin", PB_RoomsChanges_Chat_Pin_name, PB_RoomsChanges_Chat_Pin_value)
}

func init() { proto.RegisterFile("pb_changes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x6e, 0xda, 0x4c,
	0x10, 0xfd, 0x8c, 0x71, 0x02, 0xc3, 0x4f, 0xc8, 0xea, 0x4b, 0xb4, 0xb1, 0xaa, 0xd4, 0x72, 0x6f,
	0xac, 0x56, 0xa2, 0x51, 0x5a, 0xf5, 0xa2, 0xaa, 0x2a, 0x05, 0x92, 0x14, 0x14, 0x05, 0x59, 0x1b,
	0xc8, 0x45, 0x6f, 0x90, 0xc1, 0xa3, 0xb0, 0x92, 0xb1, 0x91, 0xed, 0xfc, 0xbd, 0x4f, 0x5f, 0xac,
	0x77, 0x7d, 0x8c, 0x6a, 0x77, 0x0d, 0xc6, 0x29, 0xb4, 0x57, 0xde, 0x33, 0xe7, 0xcc, 0x19, 0xef,
	0xec, 0xec, 0x42, 0x6b, 0x31, 0x19, 0x4f, 0x67, 0x5e, 0x78, 0x87, 0x49, 0x7b, 0x11, 0x47, 0x69,
	0x64, 0x36, 0x16, 0x93, 0x31, 0x86, 0xf7, 0xf3, 0x0c, 0x36, 0x17, 0x93, 0xf1, 0x03, 0xc7, 0xc7,
	0x8c, 0xb6, 0x7f, 0xed, 0xc2, 0x9e, 0xdb, 0x19, 0xb3, 0x28, 0x9a, 0x27, 0x5d, 0x95, 0x48, 0x2c,
	0xa8, 0xdd, 0x62, 0x9c, 0xf0, 0x28, 0x1c, 0xf2, 0x39, 0x52, 0xcd, 0xd2, 0x9c, 0x32, 0x5b, 0x0f,
	0x91, 0x77, 0x60, 0x74, 0x67, 0x5e, 0x9a, 0xd0, 0x92, 0xa5, 0x3b, 0xb5, 0xd3, 0x83, 0xf6, 0x0b,
	0x8b, 0xb6, 0x60, 0x99, 0xd2, 0x10, 0x1b, 0x8c, 0x58, 0x70, 0x54, 0x97, 0xe2, 0xba, 0x10, 0x0b,
	0xe6, 0x96, 0xe3, 0x23, 0x53, 0x94, 0xf9, 0x73, 0x07, 0xca, 0x22, 0x46, 0x0e, 0x61, 0x47, 0x7c,
	0xfb, 0xbe, 0x2c, 0xab, 0xb3, 0x0c, 0x11, 0x0a, 0xbb, 0xa2, 0xc0, 0x15, 0x3e, 0x53, 0xdd, 0xd2,
	0x9c, 0x2a, 0x5b, 0x42, 0x62, 0x42, 0x45, 0x2c, 0x87, 0xcf, 0x0b, 0xa4, 0x65, 0x4b, 0x73, 0x0c,
	0xb6, 0xc2, 0x82, 0x73, 0x11, 0x63, 0xf7, 0x3e, 0x99, 0x51, 0x43, 0x71, 0x4b, 0x4c, 0xde, 0x42,
	0x8b, 0xe1, 0x14, 0xf9, 0x03, 0xfa, 0xd7, 0x98, 0x24, 0xde, 0x1d, 0x26, 0x14, 0x2c, 0xdd, 0x29,
	0xb3, 0x3f, 0xe2, 0xc4, 0x86, 0xfa, 0x0d, 0x62, 0xb8, 0xd2, 0xd5, 0xa4, 0xae, 0x10, 0x23, 0x3d,
	0x68, 0x5c, 0xf8, 0x3c, 0xc5, 0x95, 0xa8, 0x2e, 0xb7, 0x6b, 0x6f, 0xec, 0x4d, 0x7b, 0x5d, 0xca,
	0x8a, 0x89, 0xe4, 0x0a, 0x9a, 0xe7, 0x18, 0xe0, 0x9a, 0x55, 0x43, 0x5a, 0xbd, 0xd9, 0x6c, 0x55,
	0xd0, 0xb2, 0x17, 0xa9, 0xe4, 0x0b, 0x1c, 0x75, 0x03, 0xf4, 0xe2, 0x1e, 0x4f, 0xd2, 0x38, 0x7a,
	0xbe, 0x8c, 0xa3, 0x79, 0xc6, 0xf5, 0x7d, 0xda, 0x94, 0x47, 0xbb, 0x5d, 0x40, 0x8e, 0x01, 0x94,
	0x9f, 0xa8, 0x44, 0xf7, 0x64, 0x0b, 0xd7, 0x22, 0xe4, 0x15, 0x54, 0xc5, 0x77, 0xc8, 0xd3, 0x00,
	0x69, 0x4b, 0x1e, 0x4c, 0x1e, 0x20, 0x27, 0x60, 0x5c, 0xdf, 0xa7, 0xe8, 0xd3, 0xff, 0x2d, 0xcd,
	0x69, 0x9e, 0x9a, 0x9b, 0xff, 0x5f, 0x48, 0x98, 0x12, 0x8a, 0x7a, 0x72, 0x31, 0x0a, 0x53, 0x1e,
	0xd0, 0x03, 0x55, 0x2f, 0x8f, 0x90, 0xf7, 0x60, 0xb8, 0x3c, 0x44, 0x9f, 0x1e, 0x4b, 0xc7, 0xa3,
	0xcd, 0x8e, 0x2e, 0x0f, 0x99, 0xd2, 0x89, 0xb9, 0x71, 0xb9, 0x9a, 0xe3, 0xd7, 0x72, 0xa0, 0x96,
	0xd0, 0xbc, 0x84, 0xfa, 0x7a, 0xdb, 0xc5, 0x56, 0xf2, 0xc6, 0xa8, 0x99, 0xcf, 0x03, 0xc2, 0x67,
	0x80, 0x8f, 0x0c, 0x9f, 0x52, 0x5a, 0x52, 0xf3, 0x97, 0x41, 0xf3, 0x0c, 0x1a, 0x85, 0x96, 0xff,
	0xc3, 0x88, 0x40, 0xb9, 0x13, 0xa5, 0x33, 0xe9, 0x52, 0x61, 0x72, 0x6d, 0x9f, 0x42, 0x59, 0xec,
	0x91, 0x10, 0x68, 0x8e, 0x06, 0xdd, 0xde, 0xd9, 0xe0, 0xdb, 0xc5, 0xf9, 0xf8, 0x7a, 0x34, 0xbc,
	0x68, 0xfd, 0x47, 0xaa, 0x60, 0x88, 0xd5, 0x79, 0x4b, 0x23, 0x35, 0xd8, 0x1d, 0x0d, 0x14, 0x28,
	0xd9, 0x27, 0xa0, 0xbb, 0x3c, 0x24, 0xfb, 0xd0, 0xc8, 0x53, 0xdc, 0xfe, 0x40, 0x65, 0xb8, 0xfd,
	0x41, 0x9e, 0xa1, 0x40, 0xc9, 0xfe, 0x51, 0x82, 0xa6, 0xdb, 0x19, 0x8b, 0xe1, 0x5f, 0xde, 0xf4,
	0x8f, 0x50, 0x5f, 0xef, 0x1e, 0xd5, 0xe4, 0x9c, 0xb5, 0x5e, 0x76, 0x95, 0x15, 0x54, 0xc4, 0x81,
	0xca, 0xf2, 0xfe, 0x66, 0x0f, 0x40, 0xf1, 0x4e, 0xaf, 0x58, 0xf2, 0x15, 0xcc, 0x7e, 0xf8, 0xe0,
	0x05, 0xdc, 0xf7, 0x52, 0xec, 0x7a, 0xd3, 0x19, 0x5e, 0x46, 0x71, 0x76, 0x71, 0xd5, 0x7b, 0x50,
	0x65, 0x7f, 0x51, 0x90, 0x4f, 0x70, 0x98, 0xb3, 0x67, 0x41, 0x20, 0x8c, 0xa5, 0x88, 0x82, 0x1c,
	0x8d, 0x2d, 0x2c, 0xf9, 0x0c, 0xb4, 0xc0, 0xdc, 0x44, 0x53, 0xee, 0x05, 0x2a, 0xb3, 0x26, 0x33,
	0xb7, 0xf2, 0x9d, 0x7d, 0xa8, 0xf0, 0xb8, 0x3d, 0x4f, 0xda, 0x8b, 0x49, 0x4f, 0x77, 0xb5, 0xef,
	0xda, 0xd3, 0x64, 0x47, 0xbe, 0x95, 0x1f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x80, 0xa1, 0x5e,
	0x29, 0x5e, 0x05, 0x00, 0x00,
}
