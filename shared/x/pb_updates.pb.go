// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_updates.proto

package x

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PB_UpdateRoomActionDoing struct {
	RoomKey    string              `protobuf:"bytes,2,opt,name=RoomKey" json:"RoomKey,omitempty"`
	ActionType RoomActionDoingEnum `protobuf:"varint,3,opt,name=ActionType,enum=RoomActionDoingEnum" json:"ActionType,omitempty"`
}

func (m *PB_UpdateRoomActionDoing) Reset()                    { *m = PB_UpdateRoomActionDoing{} }
func (m *PB_UpdateRoomActionDoing) String() string            { return proto.CompactTextString(m) }
func (*PB_UpdateRoomActionDoing) ProtoMessage()               {}
func (*PB_UpdateRoomActionDoing) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *PB_UpdateRoomActionDoing) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_UpdateRoomActionDoing) GetActionType() RoomActionDoingEnum {
	if m != nil {
		return m.ActionType
	}
	return RoomActionDoingEnum_UNKNOWN_ROOM_ACTION_DOING
}

type PB_UpdateMessageMeta struct {
	RoomKey   string `protobuf:"bytes,1,opt,name=RoomKey" json:"RoomKey,omitempty"`
	MessageId int64  `protobuf:"varint,2,opt,name=MessageId" json:"MessageId,omitempty"`
}

func (m *PB_UpdateMessageMeta) Reset()                    { *m = PB_UpdateMessageMeta{} }
func (m *PB_UpdateMessageMeta) String() string            { return proto.CompactTextString(m) }
func (*PB_UpdateMessageMeta) ProtoMessage()               {}
func (*PB_UpdateMessageMeta) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *PB_UpdateMessageMeta) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_UpdateMessageMeta) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type PB_Updates struct {
	LastChatSyncId int64 `protobuf:"varint,1,opt,name=LastChatSyncId" json:"LastChatSyncId,omitempty"`
	// new mssages
	NewMessages []*PB_MessageView `protobuf:"bytes,5,rep,name=NewMessages" json:"NewMessages,omitempty"`
	Chats       []*PB_ChatView    `protobuf:"bytes,2,rep,name=Chats" json:"Chats,omitempty"`
	// messages meta info
	MessagesReachedServer   []*PB_UpdateMessageMeta `protobuf:"bytes,20,rep,name=MessagesReachedServer" json:"MessagesReachedServer,omitempty"`
	MessagesDeliveredToUser []*PB_UpdateMessageMeta `protobuf:"bytes,21,rep,name=MessagesDeliveredToUser" json:"MessagesDeliveredToUser,omitempty"`
	MessagesSeenByPeer      []*PB_UpdateMessageMeta `protobuf:"bytes,22,rep,name=MessagesSeenByPeer" json:"MessagesSeenByPeer,omitempty"`
	// room updates
	RoomActionDoing []*PB_UpdateRoomActionDoing `protobuf:"bytes,30,rep,name=RoomActionDoing" json:"RoomActionDoing,omitempty"`
}

func (m *PB_Updates) Reset()                    { *m = PB_Updates{} }
func (m *PB_Updates) String() string            { return proto.CompactTextString(m) }
func (*PB_Updates) ProtoMessage()               {}
func (*PB_Updates) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *PB_Updates) GetLastChatSyncId() int64 {
	if m != nil {
		return m.LastChatSyncId
	}
	return 0
}

func (m *PB_Updates) GetNewMessages() []*PB_MessageView {
	if m != nil {
		return m.NewMessages
	}
	return nil
}

func (m *PB_Updates) GetChats() []*PB_ChatView {
	if m != nil {
		return m.Chats
	}
	return nil
}

func (m *PB_Updates) GetMessagesReachedServer() []*PB_UpdateMessageMeta {
	if m != nil {
		return m.MessagesReachedServer
	}
	return nil
}

func (m *PB_Updates) GetMessagesDeliveredToUser() []*PB_UpdateMessageMeta {
	if m != nil {
		return m.MessagesDeliveredToUser
	}
	return nil
}

func (m *PB_Updates) GetMessagesSeenByPeer() []*PB_UpdateMessageMeta {
	if m != nil {
		return m.MessagesSeenByPeer
	}
	return nil
}

func (m *PB_Updates) GetRoomActionDoing() []*PB_UpdateRoomActionDoing {
	if m != nil {
		return m.RoomActionDoing
	}
	return nil
}

func init() {
	proto.RegisterType((*PB_UpdateRoomActionDoing)(nil), "PB_UpdateRoomActionDoing")
	proto.RegisterType((*PB_UpdateMessageMeta)(nil), "PB_UpdateMessageMeta")
	proto.RegisterType((*PB_Updates)(nil), "PB_Updates")
}

func init() { proto.RegisterFile("pb_updates.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x53, 0x1a, 0x54, 0x06, 0x05, 0xdd, 0x80, 0xae, 0xc6, 0x18, 0xd2, 0x83, 0xe1, 0xd4,
	0x44, 0xf4, 0x0b, 0x58, 0x20, 0x91, 0x20, 0xd8, 0x2c, 0xe0, 0xc1, 0x4b, 0xd3, 0xd2, 0x09, 0xd4,
	0xd8, 0x3f, 0xe9, 0x16, 0xb0, 0x1f, 0xcc, 0xef, 0x67, 0xba, 0xb4, 0x88, 0x0d, 0x70, 0x9c, 0xf7,
	0x7e, 0xef, 0x4d, 0x9b, 0x1d, 0x38, 0x0f, 0x2c, 0x63, 0x11, 0xd8, 0x66, 0x84, 0x5c, 0x0d, 0x42,
	0x3f, 0xf2, 0x6f, 0xce, 0x02, 0xcb, 0x40, 0x6f, 0xe1, 0xa6, 0x63, 0x25, 0xb0, 0x8c, 0xa5, 0x83,
	0xab, 0xd4, 0x56, 0x3e, 0x81, 0xea, 0x9a, 0x31, 0x11, 0x11, 0xe6, 0xfb, 0xee, 0xf3, 0x34, 0x72,
	0x7c, 0xaf, 0xe3, 0x3b, 0xde, 0x8c, 0x50, 0x38, 0x4e, 0xa4, 0x3e, 0xc6, 0xb4, 0xd0, 0x90, 0x9a,
	0x25, 0x96, 0x8d, 0xe4, 0x09, 0x60, 0x0d, 0x8e, 0xe3, 0x00, 0xa9, 0xdc, 0x90, 0x9a, 0x95, 0x56,
	0x4d, 0xcd, 0xe5, 0xbb, 0xde, 0xc2, 0x65, 0x5b, 0x9c, 0x32, 0x84, 0xda, 0x66, 0xd7, 0x00, 0x39,
	0x37, 0x67, 0x38, 0xc0, 0xc8, 0xdc, 0xde, 0x23, 0xfd, 0xdf, 0x73, 0x0b, 0xa5, 0x14, 0xec, 0xd9,
	0xe2, 0x1b, 0x64, 0xf6, 0x27, 0x28, 0x3f, 0x32, 0xc0, 0xa6, 0x90, 0x93, 0x7b, 0xa8, 0xbc, 0x9a,
	0x3c, 0x6a, 0xcf, 0xcd, 0x68, 0x14, 0x7b, 0xd3, 0x9e, 0x2d, 0xda, 0x64, 0x96, 0x53, 0xc9, 0x03,
	0x94, 0x87, 0xb8, 0x4a, 0x6b, 0x38, 0x2d, 0x36, 0xe4, 0x66, 0xb9, 0x55, 0x55, 0x75, 0xcd, 0x48,
	0xb5, 0x77, 0x07, 0x57, 0x6c, 0x9b, 0x21, 0x0a, 0x14, 0x93, 0x02, 0x4e, 0x0b, 0x02, 0x3e, 0x4d,
	0xe0, 0x44, 0x10, 0xe4, 0xda, 0x22, 0x7d, 0xa8, 0x67, 0x3c, 0x43, 0x73, 0x3a, 0x47, 0x7b, 0x84,
	0xe1, 0x12, 0x43, 0x5a, 0x13, 0x99, 0xba, 0xba, 0xeb, 0xdf, 0xd9, 0xee, 0x0c, 0x79, 0x83, 0xab,
	0xcc, 0xe8, 0xe0, 0x97, 0xb3, 0xc4, 0x10, 0xed, 0xb1, 0x3f, 0xe1, 0x18, 0xd2, 0xfa, 0xa1, 0xba,
	0x7d, 0x29, 0xd2, 0x05, 0x92, 0x59, 0x23, 0x44, 0x4f, 0x8b, 0x75, 0xc4, 0x90, 0x5e, 0x1e, 0xea,
	0xda, 0x11, 0x20, 0x6d, 0xa8, 0xe6, 0x5e, 0x99, 0xde, 0x89, 0x8e, 0x6b, 0x75, 0xdf, 0x19, 0xb1,
	0x7c, 0x42, 0xbb, 0x80, 0x13, 0x27, 0x54, 0x5d, 0xae, 0x06, 0xd6, 0x8b, 0xac, 0x4b, 0x1f, 0xd2,
	0xb7, 0x75, 0x24, 0xae, 0xf1, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x77, 0x71, 0xb2, 0xbc, 0xc0,
	0x02, 0x00, 0x00,
}
