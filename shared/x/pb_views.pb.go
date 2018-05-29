// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_views.proto

package x

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PB_ChatView struct {
	ChatId           int64         `protobuf:"varint,1,opt,name=ChatId" json:"ChatId,omitempty"`
	ChatKey          string        `protobuf:"bytes,2,opt,name=ChatKey" json:"ChatKey,omitempty"`
	RoomKey          string        `protobuf:"bytes,3,opt,name=RoomKey" json:"RoomKey,omitempty"`
	RoomType         int32         `protobuf:"varint,4,opt,name=RoomType" json:"RoomType,omitempty"`
	UserId           int32         `protobuf:"varint,5,opt,name=UserId" json:"UserId,omitempty"`
	PeerUserId       int32         `protobuf:"varint,6,opt,name=PeerUserId" json:"PeerUserId,omitempty"`
	GroupId          int64         `protobuf:"varint,7,opt,name=GroupId" json:"GroupId,omitempty"`
	HashTagId        int64         `protobuf:"varint,8,opt,name=HashTagId" json:"HashTagId,omitempty"`
	StartedByMe      int32         `protobuf:"varint,9,opt,name=StartedByMe" json:"StartedByMe,omitempty"`
	Title            string        `protobuf:"bytes,10,opt,name=Title" json:"Title,omitempty"`
	PinTime          int64         `protobuf:"varint,11,opt,name=PinTime" json:"PinTime,omitempty"`
	FromMsgId        int64         `protobuf:"varint,12,opt,name=FromMsgId" json:"FromMsgId,omitempty"`
	Seq              int32         `protobuf:"varint,13,opt,name=Seq" json:"Seq,omitempty"`
	LastMsgId        int64         `protobuf:"varint,14,opt,name=LastMsgId" json:"LastMsgId,omitempty"`
	LastMsgStatus    int32         `protobuf:"varint,15,opt,name=LastMsgStatus" json:"LastMsgStatus,omitempty"`
	SeenSeq          int32         `protobuf:"varint,16,opt,name=SeenSeq" json:"SeenSeq,omitempty"`
	SeenMsgId        int64         `protobuf:"varint,17,opt,name=SeenMsgId" json:"SeenMsgId,omitempty"`
	Left             int32         `protobuf:"varint,18,opt,name=Left" json:"Left,omitempty"`
	Creator          int32         `protobuf:"varint,19,opt,name=Creator" json:"Creator,omitempty"`
	Kicked           int32         `protobuf:"varint,20,opt,name=Kicked" json:"Kicked,omitempty"`
	Admin            int32         `protobuf:"varint,21,opt,name=Admin" json:"Admin,omitempty"`
	Deactivated      int32         `protobuf:"varint,22,opt,name=Deactivated" json:"Deactivated,omitempty"`
	VersionTime      int32         `protobuf:"varint,23,opt,name=VersionTime" json:"VersionTime,omitempty"`
	SortTime         int32         `protobuf:"varint,24,opt,name=SortTime" json:"SortTime,omitempty"`
	CreatedTime      int32         `protobuf:"varint,25,opt,name=CreatedTime" json:"CreatedTime,omitempty"`
	DraftText        string        `protobuf:"bytes,26,opt,name=DraftText" json:"DraftText,omitempty"`
	DratReplyToMsgId int64         `protobuf:"varint,27,opt,name=DratReplyToMsgId" json:"DratReplyToMsgId,omitempty"`
	IsMute           int32         `protobuf:"varint,28,opt,name=IsMute" json:"IsMute,omitempty"`
	UserView         *PB_UserView  `protobuf:"bytes,100,opt,name=UserView" json:"UserView,omitempty"`
	GroupView        *PB_GroupView `protobuf:"bytes,101,opt,name=GroupView" json:"GroupView,omitempty"`
	// seeting, notification, group, tag
	FirstUnreadMessage *PB_MessageView `protobuf:"bytes,200,opt,name=FirstUnreadMessage" json:"FirstUnreadMessage,omitempty"`
	LastMessage        *PB_MessageView `protobuf:"bytes,201,opt,name=LastMessage" json:"LastMessage,omitempty"`
}

func (m *PB_ChatView) Reset()                    { *m = PB_ChatView{} }
func (m *PB_ChatView) String() string            { return proto.CompactTextString(m) }
func (*PB_ChatView) ProtoMessage()               {}
func (*PB_ChatView) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *PB_ChatView) GetChatId() int64 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *PB_ChatView) GetChatKey() string {
	if m != nil {
		return m.ChatKey
	}
	return ""
}

func (m *PB_ChatView) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_ChatView) GetRoomType() int32 {
	if m != nil {
		return m.RoomType
	}
	return 0
}

func (m *PB_ChatView) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_ChatView) GetPeerUserId() int32 {
	if m != nil {
		return m.PeerUserId
	}
	return 0
}

func (m *PB_ChatView) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *PB_ChatView) GetHashTagId() int64 {
	if m != nil {
		return m.HashTagId
	}
	return 0
}

func (m *PB_ChatView) GetStartedByMe() int32 {
	if m != nil {
		return m.StartedByMe
	}
	return 0
}

func (m *PB_ChatView) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PB_ChatView) GetPinTime() int64 {
	if m != nil {
		return m.PinTime
	}
	return 0
}

func (m *PB_ChatView) GetFromMsgId() int64 {
	if m != nil {
		return m.FromMsgId
	}
	return 0
}

func (m *PB_ChatView) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PB_ChatView) GetLastMsgId() int64 {
	if m != nil {
		return m.LastMsgId
	}
	return 0
}

func (m *PB_ChatView) GetLastMsgStatus() int32 {
	if m != nil {
		return m.LastMsgStatus
	}
	return 0
}

func (m *PB_ChatView) GetSeenSeq() int32 {
	if m != nil {
		return m.SeenSeq
	}
	return 0
}

func (m *PB_ChatView) GetSeenMsgId() int64 {
	if m != nil {
		return m.SeenMsgId
	}
	return 0
}

func (m *PB_ChatView) GetLeft() int32 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *PB_ChatView) GetCreator() int32 {
	if m != nil {
		return m.Creator
	}
	return 0
}

func (m *PB_ChatView) GetKicked() int32 {
	if m != nil {
		return m.Kicked
	}
	return 0
}

func (m *PB_ChatView) GetAdmin() int32 {
	if m != nil {
		return m.Admin
	}
	return 0
}

func (m *PB_ChatView) GetDeactivated() int32 {
	if m != nil {
		return m.Deactivated
	}
	return 0
}

func (m *PB_ChatView) GetVersionTime() int32 {
	if m != nil {
		return m.VersionTime
	}
	return 0
}

func (m *PB_ChatView) GetSortTime() int32 {
	if m != nil {
		return m.SortTime
	}
	return 0
}

func (m *PB_ChatView) GetCreatedTime() int32 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *PB_ChatView) GetDraftText() string {
	if m != nil {
		return m.DraftText
	}
	return ""
}

func (m *PB_ChatView) GetDratReplyToMsgId() int64 {
	if m != nil {
		return m.DratReplyToMsgId
	}
	return 0
}

func (m *PB_ChatView) GetIsMute() int32 {
	if m != nil {
		return m.IsMute
	}
	return 0
}

func (m *PB_ChatView) GetUserView() *PB_UserView {
	if m != nil {
		return m.UserView
	}
	return nil
}

func (m *PB_ChatView) GetGroupView() *PB_GroupView {
	if m != nil {
		return m.GroupView
	}
	return nil
}

func (m *PB_ChatView) GetFirstUnreadMessage() *PB_MessageView {
	if m != nil {
		return m.FirstUnreadMessage
	}
	return nil
}

func (m *PB_ChatView) GetLastMessage() *PB_MessageView {
	if m != nil {
		return m.LastMessage
	}
	return nil
}

type PB_GroupView struct {
	GroupId         int64  `protobuf:"varint,1,opt,name=GroupId" json:"GroupId,omitempty"`
	GroupKey        string `protobuf:"bytes,2,opt,name=GroupKey" json:"GroupKey,omitempty"`
	GroupName       string `protobuf:"bytes,3,opt,name=GroupName" json:"GroupName,omitempty"`
	UserName        string `protobuf:"bytes,4,opt,name=UserName" json:"UserName,omitempty"`
	IsSuperGroup    int32  `protobuf:"varint,5,opt,name=IsSuperGroup" json:"IsSuperGroup,omitempty"`
	HashTagId       int64  `protobuf:"varint,6,opt,name=HashTagId" json:"HashTagId,omitempty"`
	CreatorUserId   int32  `protobuf:"varint,7,opt,name=CreatorUserId" json:"CreatorUserId,omitempty"`
	GroupPrivacy    int32  `protobuf:"varint,8,opt,name=GroupPrivacy" json:"GroupPrivacy,omitempty"`
	HistoryViewAble int32  `protobuf:"varint,9,opt,name=HistoryViewAble" json:"HistoryViewAble,omitempty"`
	Seq             int64  `protobuf:"varint,10,opt,name=Seq" json:"Seq,omitempty"`
	LastMsgId       int64  `protobuf:"varint,11,opt,name=LastMsgId" json:"LastMsgId,omitempty"`
	PinedMsgId      int64  `protobuf:"varint,12,opt,name=PinedMsgId" json:"PinedMsgId,omitempty"`
	AvatarRefId     int64  `protobuf:"varint,13,opt,name=AvatarRefId" json:"AvatarRefId,omitempty"`
	AvatarCount     int32  `protobuf:"varint,14,opt,name=AvatarCount" json:"AvatarCount,omitempty"`
	About           string `protobuf:"bytes,15,opt,name=About" json:"About,omitempty"`
	InviteLink      string `protobuf:"bytes,16,opt,name=InviteLink" json:"InviteLink,omitempty"`
	MembersCount    int32  `protobuf:"varint,17,opt,name=MembersCount" json:"MembersCount,omitempty"`
	SortTime        int32  `protobuf:"varint,18,opt,name=SortTime" json:"SortTime,omitempty"`
	CreatedTime     int32  `protobuf:"varint,19,opt,name=CreatedTime" json:"CreatedTime,omitempty"`
}

func (m *PB_GroupView) Reset()                    { *m = PB_GroupView{} }
func (m *PB_GroupView) String() string            { return proto.CompactTextString(m) }
func (*PB_GroupView) ProtoMessage()               {}
func (*PB_GroupView) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *PB_GroupView) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *PB_GroupView) GetGroupKey() string {
	if m != nil {
		return m.GroupKey
	}
	return ""
}

func (m *PB_GroupView) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *PB_GroupView) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *PB_GroupView) GetIsSuperGroup() int32 {
	if m != nil {
		return m.IsSuperGroup
	}
	return 0
}

func (m *PB_GroupView) GetHashTagId() int64 {
	if m != nil {
		return m.HashTagId
	}
	return 0
}

func (m *PB_GroupView) GetCreatorUserId() int32 {
	if m != nil {
		return m.CreatorUserId
	}
	return 0
}

func (m *PB_GroupView) GetGroupPrivacy() int32 {
	if m != nil {
		return m.GroupPrivacy
	}
	return 0
}

func (m *PB_GroupView) GetHistoryViewAble() int32 {
	if m != nil {
		return m.HistoryViewAble
	}
	return 0
}

func (m *PB_GroupView) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PB_GroupView) GetLastMsgId() int64 {
	if m != nil {
		return m.LastMsgId
	}
	return 0
}

func (m *PB_GroupView) GetPinedMsgId() int64 {
	if m != nil {
		return m.PinedMsgId
	}
	return 0
}

func (m *PB_GroupView) GetAvatarRefId() int64 {
	if m != nil {
		return m.AvatarRefId
	}
	return 0
}

func (m *PB_GroupView) GetAvatarCount() int32 {
	if m != nil {
		return m.AvatarCount
	}
	return 0
}

func (m *PB_GroupView) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *PB_GroupView) GetInviteLink() string {
	if m != nil {
		return m.InviteLink
	}
	return ""
}

func (m *PB_GroupView) GetMembersCount() int32 {
	if m != nil {
		return m.MembersCount
	}
	return 0
}

func (m *PB_GroupView) GetSortTime() int32 {
	if m != nil {
		return m.SortTime
	}
	return 0
}

func (m *PB_GroupView) GetCreatedTime() int32 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

type PB_GroupMemberView struct {
	OrderId     int64 `protobuf:"varint,1,opt,name=OrderId" json:"OrderId,omitempty"`
	GroupId     int64 `protobuf:"varint,2,opt,name=GroupId" json:"GroupId,omitempty"`
	UserId      int32 `protobuf:"varint,3,opt,name=UserId" json:"UserId,omitempty"`
	ByUserId    int32 `protobuf:"varint,4,opt,name=ByUserId" json:"ByUserId,omitempty"`
	GroupRole   int32 `protobuf:"varint,5,opt,name=GroupRole" json:"GroupRole,omitempty"`
	CreatedTime int32 `protobuf:"varint,6,opt,name=CreatedTime" json:"CreatedTime,omitempty"`
}

func (m *PB_GroupMemberView) Reset()                    { *m = PB_GroupMemberView{} }
func (m *PB_GroupMemberView) String() string            { return proto.CompactTextString(m) }
func (*PB_GroupMemberView) ProtoMessage()               {}
func (*PB_GroupMemberView) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *PB_GroupMemberView) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *PB_GroupMemberView) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *PB_GroupMemberView) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_GroupMemberView) GetByUserId() int32 {
	if m != nil {
		return m.ByUserId
	}
	return 0
}

func (m *PB_GroupMemberView) GetGroupRole() int32 {
	if m != nil {
		return m.GroupRole
	}
	return 0
}

func (m *PB_GroupMemberView) GetCreatedTime() int32 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

type PB_MessageView struct {
	RoomKey          string          `protobuf:"bytes,1,opt,name=RoomKey" json:"RoomKey,omitempty"`
	MessageId        int64           `protobuf:"varint,2,opt,name=MessageId" json:"MessageId,omitempty"`
	UserId           int32           `protobuf:"varint,3,opt,name=UserId" json:"UserId,omitempty"`
	FileRefId        int64           `protobuf:"varint,4,opt,name=FileRefId" json:"FileRefId,omitempty"`
	MessageType      int32           `protobuf:"varint,5,opt,name=MessageType" json:"MessageType,omitempty"`
	Text             string          `protobuf:"bytes,6,opt,name=Text" json:"Text,omitempty"`
	Hiden            int32           `protobuf:"varint,7,opt,name=Hiden" json:"Hiden,omitempty"`
	Seq              int32           `protobuf:"varint,8,opt,name=Seq" json:"Seq,omitempty"`
	ForwardedMsgId   int64           `protobuf:"varint,9,opt,name=ForwardedMsgId" json:"ForwardedMsgId,omitempty"`
	PostId           int64           `protobuf:"varint,10,opt,name=PostId" json:"PostId,omitempty"`
	StickerId        int64           `protobuf:"varint,11,opt,name=StickerId" json:"StickerId,omitempty"`
	CreatedTime      int32           `protobuf:"varint,12,opt,name=CreatedTime" json:"CreatedTime,omitempty"`
	DeliveredTime    int32           `protobuf:"varint,13,opt,name=DeliveredTime" json:"DeliveredTime,omitempty"`
	SeenTime         int32           `protobuf:"varint,14,opt,name=SeenTime" json:"SeenTime,omitempty"`
	DeliviryStatus   int32           `protobuf:"varint,15,opt,name=DeliviryStatus" json:"DeliviryStatus,omitempty"`
	ReplyToMessageId int64           `protobuf:"varint,16,opt,name=ReplyToMessageId" json:"ReplyToMessageId,omitempty"`
	ViewsCount       int64           `protobuf:"varint,17,opt,name=ViewsCount" json:"ViewsCount,omitempty"`
	EditTime         int32           `protobuf:"varint,18,opt,name=EditTime" json:"EditTime,omitempty"`
	Ttl              int32           `protobuf:"varint,19,opt,name=Ttl" json:"Ttl,omitempty"`
	FileRedView      *PB_FileRedView `protobuf:"bytes,50,opt,name=FileRedView" json:"FileRedView,omitempty"`
}

func (m *PB_MessageView) Reset()                    { *m = PB_MessageView{} }
func (m *PB_MessageView) String() string            { return proto.CompactTextString(m) }
func (*PB_MessageView) ProtoMessage()               {}
func (*PB_MessageView) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *PB_MessageView) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *PB_MessageView) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *PB_MessageView) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_MessageView) GetFileRefId() int64 {
	if m != nil {
		return m.FileRefId
	}
	return 0
}

func (m *PB_MessageView) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *PB_MessageView) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *PB_MessageView) GetHiden() int32 {
	if m != nil {
		return m.Hiden
	}
	return 0
}

func (m *PB_MessageView) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PB_MessageView) GetForwardedMsgId() int64 {
	if m != nil {
		return m.ForwardedMsgId
	}
	return 0
}

func (m *PB_MessageView) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *PB_MessageView) GetStickerId() int64 {
	if m != nil {
		return m.StickerId
	}
	return 0
}

func (m *PB_MessageView) GetCreatedTime() int32 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *PB_MessageView) GetDeliveredTime() int32 {
	if m != nil {
		return m.DeliveredTime
	}
	return 0
}

func (m *PB_MessageView) GetSeenTime() int32 {
	if m != nil {
		return m.SeenTime
	}
	return 0
}

func (m *PB_MessageView) GetDeliviryStatus() int32 {
	if m != nil {
		return m.DeliviryStatus
	}
	return 0
}

func (m *PB_MessageView) GetReplyToMessageId() int64 {
	if m != nil {
		return m.ReplyToMessageId
	}
	return 0
}

func (m *PB_MessageView) GetViewsCount() int64 {
	if m != nil {
		return m.ViewsCount
	}
	return 0
}

func (m *PB_MessageView) GetEditTime() int32 {
	if m != nil {
		return m.EditTime
	}
	return 0
}

func (m *PB_MessageView) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *PB_MessageView) GetFileRedView() *PB_FileRedView {
	if m != nil {
		return m.FileRedView
	}
	return nil
}

type PB_FileRedView struct {
	FileRefId int64  `protobuf:"varint,1,opt,name=FileRefId" json:"FileRefId,omitempty"`
	UserId    int64  `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Width     int32  `protobuf:"varint,4,opt,name=Width" json:"Width,omitempty"`
	Height    int32  `protobuf:"varint,5,opt,name=Height" json:"Height,omitempty"`
	Duration  int32  `protobuf:"varint,6,opt,name=Duration" json:"Duration,omitempty"`
	Extension string `protobuf:"bytes,7,opt,name=Extension" json:"Extension,omitempty"`
	UrlSource string `protobuf:"bytes,8,opt,name=UrlSource" json:"UrlSource,omitempty"`
}

func (m *PB_FileRedView) Reset()                    { *m = PB_FileRedView{} }
func (m *PB_FileRedView) String() string            { return proto.CompactTextString(m) }
func (*PB_FileRedView) ProtoMessage()               {}
func (*PB_FileRedView) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *PB_FileRedView) GetFileRefId() int64 {
	if m != nil {
		return m.FileRefId
	}
	return 0
}

func (m *PB_FileRedView) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_FileRedView) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PB_FileRedView) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *PB_FileRedView) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *PB_FileRedView) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *PB_FileRedView) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

func (m *PB_FileRedView) GetUrlSource() string {
	if m != nil {
		return m.UrlSource
	}
	return ""
}

type PB_UserView struct {
	UserId         int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	UserName       string `protobuf:"bytes,2,opt,name=UserName" json:"UserName,omitempty"`
	FirstName      string `protobuf:"bytes,4,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName       string `protobuf:"bytes,5,opt,name=LastName" json:"LastName,omitempty"`
	AvatarRefId    int64  `protobuf:"varint,8,opt,name=AvatarRefId" json:"AvatarRefId,omitempty"`
	ProfilePrivacy int32  `protobuf:"varint,9,opt,name=ProfilePrivacy" json:"ProfilePrivacy,omitempty"`
	Phone          int64  `protobuf:"varint,10,opt,name=Phone" json:"Phone,omitempty"`
	About          string `protobuf:"bytes,11,opt,name=About" json:"About,omitempty"`
	// counters 100 - 200
	FollowersCount int32 `protobuf:"varint,100,opt,name=FollowersCount" json:"FollowersCount,omitempty"`
	FollowingCount int32 `protobuf:"varint,101,opt,name=FollowingCount" json:"FollowingCount,omitempty"`
	PostsCount     int32 `protobuf:"varint,102,opt,name=PostsCount" json:"PostsCount,omitempty"`
	MediaCount     int32 `protobuf:"varint,103,opt,name=MediaCount" json:"MediaCount,omitempty"`
	// last activities
	UserOnlineStatusEnum UserOnlineStatusEnum `protobuf:"varint,200,opt,name=UserOnlineStatusEnum,enum=UserOnlineStatusEnum" json:"UserOnlineStatusEnum,omitempty"`
	LastActiveTime       int32                `protobuf:"varint,201,opt,name=LastActiveTime" json:"LastActiveTime,omitempty"`
	LastActiveTimeShow   string               `protobuf:"bytes,202,opt,name=LastActiveTimeShow" json:"LastActiveTimeShow,omitempty"`
	// with me
	MyFollwing FollowingEnum `protobuf:"varint,300,opt,name=MyFollwing,enum=FollowingEnum" json:"MyFollwing,omitempty"`
}

func (m *PB_UserView) Reset()                    { *m = PB_UserView{} }
func (m *PB_UserView) String() string            { return proto.CompactTextString(m) }
func (*PB_UserView) ProtoMessage()               {}
func (*PB_UserView) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *PB_UserView) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PB_UserView) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *PB_UserView) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *PB_UserView) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *PB_UserView) GetAvatarRefId() int64 {
	if m != nil {
		return m.AvatarRefId
	}
	return 0
}

func (m *PB_UserView) GetProfilePrivacy() int32 {
	if m != nil {
		return m.ProfilePrivacy
	}
	return 0
}

func (m *PB_UserView) GetPhone() int64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

func (m *PB_UserView) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *PB_UserView) GetFollowersCount() int32 {
	if m != nil {
		return m.FollowersCount
	}
	return 0
}

func (m *PB_UserView) GetFollowingCount() int32 {
	if m != nil {
		return m.FollowingCount
	}
	return 0
}

func (m *PB_UserView) GetPostsCount() int32 {
	if m != nil {
		return m.PostsCount
	}
	return 0
}

func (m *PB_UserView) GetMediaCount() int32 {
	if m != nil {
		return m.MediaCount
	}
	return 0
}

func (m *PB_UserView) GetUserOnlineStatusEnum() UserOnlineStatusEnum {
	if m != nil {
		return m.UserOnlineStatusEnum
	}
	return UserOnlineStatusEnum_EXACTLY
}

func (m *PB_UserView) GetLastActiveTime() int32 {
	if m != nil {
		return m.LastActiveTime
	}
	return 0
}

func (m *PB_UserView) GetLastActiveTimeShow() string {
	if m != nil {
		return m.LastActiveTimeShow
	}
	return ""
}

func (m *PB_UserView) GetMyFollwing() FollowingEnum {
	if m != nil {
		return m.MyFollwing
	}
	return FollowingEnum_FOLLOWING_NONE
}

func init() {
	proto.RegisterType((*PB_ChatView)(nil), "PB_ChatView")
	proto.RegisterType((*PB_GroupView)(nil), "PB_GroupView")
	proto.RegisterType((*PB_GroupMemberView)(nil), "PB_GroupMemberView")
	proto.RegisterType((*PB_MessageView)(nil), "PB_MessageView")
	proto.RegisterType((*PB_FileRedView)(nil), "PB_FileRedView")
	proto.RegisterType((*PB_UserView)(nil), "PB_UserView")
}

func init() { proto.RegisterFile("pb_views.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 1302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x57, 0x5f, 0x6e, 0xdb, 0x46,
	0x13, 0x07, 0x2d, 0xd9, 0x96, 0x56, 0xb6, 0xec, 0x6c, 0xfe, 0x7c, 0xfb, 0xb9, 0x41, 0x61, 0x18,
	0x41, 0x6a, 0xb4, 0x80, 0x83, 0xba, 0x17, 0xa8, 0x1d, 0x27, 0xb5, 0x9a, 0xa8, 0x11, 0x28, 0x25,
	0x05, 0xfa, 0x62, 0x50, 0xe6, 0xd8, 0x5e, 0x84, 0x22, 0xd5, 0xe5, 0xca, 0x8e, 0xee, 0xd4, 0x23,
	0xf4, 0x00, 0x4d, 0x4f, 0xd0, 0xc7, 0x1c, 0xa1, 0x40, 0x2f, 0x50, 0xcc, 0xcc, 0x92, 0x5c, 0xd2,
	0x4e, 0xdf, 0x76, 0x7e, 0x33, 0xb3, 0x3b, 0x3b, 0xfb, 0x9b, 0x19, 0x52, 0xf4, 0xe7, 0xd3, 0xb3,
	0x6b, 0x0d, 0x37, 0xf9, 0xc1, 0xdc, 0x64, 0x36, 0xdb, 0xd9, 0x9a, 0x4f, 0xcf, 0x6c, 0x34, 0x4d,
	0xa0, 0x00, 0x36, 0xe7, 0xd3, 0x33, 0x48, 0x17, 0x33, 0x16, 0xf7, 0xfe, 0x5e, 0x17, 0xbd, 0xd1,
	0xf1, 0xd9, 0xf3, 0xab, 0xc8, 0xbe, 0xd3, 0x70, 0x23, 0x1f, 0x89, 0x35, 0x5c, 0x0f, 0x62, 0x15,
	0xec, 0x06, 0xfb, 0xad, 0xd0, 0x49, 0x52, 0x89, 0x75, 0x5c, 0xbd, 0x82, 0xa5, 0x5a, 0xd9, 0x0d,
	0xf6, 0xbb, 0x61, 0x21, 0xa2, 0x26, 0xcc, 0xb2, 0x19, 0x6a, 0x5a, 0xac, 0x71, 0xa2, 0xdc, 0x11,
	0x1d, 0x5c, 0x4e, 0x96, 0x73, 0x50, 0xed, 0xdd, 0x60, 0x7f, 0x35, 0x2c, 0x65, 0x3c, 0xe7, 0x6d,
	0x0e, 0x66, 0x10, 0xab, 0x55, 0xd2, 0x38, 0x49, 0x7e, 0x29, 0xc4, 0x08, 0xc0, 0x38, 0xdd, 0x1a,
	0xe9, 0x3c, 0x04, 0x4f, 0xfb, 0xc1, 0x64, 0x8b, 0xf9, 0x20, 0x56, 0xeb, 0x14, 0x60, 0x21, 0xca,
	0xc7, 0xa2, 0x7b, 0x1a, 0xe5, 0x57, 0x93, 0xe8, 0x72, 0x10, 0xab, 0x0e, 0xe9, 0x2a, 0x40, 0xee,
	0x8a, 0xde, 0xd8, 0x46, 0xc6, 0x42, 0x7c, 0xbc, 0x1c, 0x82, 0xea, 0xd2, 0xc6, 0x3e, 0x24, 0x1f,
	0x88, 0xd5, 0x89, 0xb6, 0x09, 0x28, 0x41, 0xb7, 0x60, 0x01, 0xcf, 0x1b, 0xe9, 0x74, 0xa2, 0x67,
	0xa0, 0x7a, 0x7c, 0x9e, 0x13, 0xf1, 0xbc, 0x97, 0x26, 0x9b, 0x0d, 0x73, 0x3c, 0x6f, 0x83, 0xcf,
	0x2b, 0x01, 0xb9, 0x2d, 0x5a, 0x63, 0xf8, 0x55, 0x6d, 0xd2, 0x39, 0xb8, 0x44, 0xfb, 0xd7, 0x51,
	0x6e, 0xd9, 0xbe, 0xcf, 0xf6, 0x25, 0x20, 0x9f, 0x88, 0x4d, 0x27, 0x8c, 0x6d, 0x64, 0x17, 0xb9,
	0xda, 0x22, 0xcf, 0x3a, 0x88, 0xd1, 0x8c, 0x01, 0x52, 0xdc, 0x79, 0x9b, 0xf4, 0x85, 0x88, 0xbb,
	0xe3, 0x92, 0x77, 0xbf, 0xc7, 0xbb, 0x97, 0x80, 0x94, 0xa2, 0xfd, 0x1a, 0x2e, 0xac, 0x92, 0xe4,
	0x44, 0x6b, 0x7a, 0x51, 0x03, 0x91, 0xcd, 0x8c, 0xba, 0xcf, 0x7b, 0x39, 0x11, 0xdf, 0xe6, 0x95,
	0x3e, 0x7f, 0x0f, 0xb1, 0x7a, 0xc0, 0x6f, 0xc3, 0x12, 0x66, 0xe8, 0x28, 0x9e, 0xe9, 0x54, 0x3d,
	0x24, 0x98, 0x05, 0xcc, 0xec, 0x09, 0x44, 0xe7, 0x56, 0x5f, 0x47, 0x16, 0x62, 0xf5, 0x88, 0x33,
	0xeb, 0x41, 0x68, 0xf1, 0x0e, 0x4c, 0xae, 0x33, 0xce, 0xe3, 0xff, 0xd8, 0xc2, 0x83, 0x90, 0x29,
	0xe3, 0xcc, 0x58, 0x52, 0x2b, 0x66, 0x4a, 0x21, 0xa3, 0x37, 0x05, 0x06, 0x31, 0xa9, 0xff, 0xcf,
	0xde, 0x1e, 0x84, 0x77, 0x3f, 0x31, 0xd1, 0x85, 0x9d, 0xc0, 0x07, 0xab, 0x76, 0xe8, 0xf5, 0x2a,
	0x40, 0x7e, 0x2d, 0xb6, 0x4f, 0x4c, 0x64, 0x43, 0x98, 0x27, 0xcb, 0x49, 0xc6, 0x09, 0xfa, 0x82,
	0x12, 0x74, 0x0b, 0xc7, 0x9b, 0x0f, 0xf2, 0xe1, 0xc2, 0x82, 0x7a, 0xcc, 0x37, 0x67, 0x49, 0xee,
	0x8b, 0x0e, 0xf2, 0x0f, 0x2b, 0x44, 0xc5, 0xbb, 0xc1, 0x7e, 0xef, 0x70, 0xe3, 0x60, 0x74, 0x7c,
	0x56, 0x60, 0x61, 0xa9, 0x95, 0xdf, 0x88, 0x2e, 0x11, 0x92, 0x4c, 0x81, 0x4c, 0x37, 0xd1, 0xb4,
	0x04, 0xc3, 0x4a, 0x2f, 0xbf, 0x17, 0xf2, 0xa5, 0x36, 0xb9, 0x7d, 0x9b, 0x1a, 0x88, 0xe2, 0x21,
	0xe4, 0x79, 0x74, 0x09, 0xea, 0x8f, 0x80, 0xdc, 0xb6, 0xd0, 0xcd, 0x61, 0xe4, 0x78, 0x87, 0xad,
	0x3c, 0x14, 0x3d, 0x62, 0x88, 0x73, 0xfd, 0xf8, 0x19, 0x57, 0xdf, 0x68, 0xef, 0x53, 0x5b, 0x6c,
	0xf8, 0x11, 0xf9, 0x35, 0x15, 0xd4, 0x6b, 0x6a, 0x47, 0x74, 0x68, 0x59, 0x95, 0x7d, 0x29, 0x63,
	0xd6, 0x69, 0xfd, 0x53, 0x34, 0x03, 0x57, 0xf9, 0x15, 0x80, 0x9e, 0x98, 0x13, 0x52, 0xb6, 0xd9,
	0xb3, 0x90, 0xe5, 0x9e, 0xd8, 0x18, 0xe4, 0xe3, 0xc5, 0x1c, 0x0c, 0xd9, 0xbb, 0x0e, 0x50, 0xc3,
	0xea, 0xd5, 0xbc, 0xd6, 0xac, 0xe6, 0x27, 0x62, 0xd3, 0x91, 0xd5, 0x35, 0x8a, 0x75, 0xae, 0x96,
	0x1a, 0x88, 0xe7, 0xd0, 0x66, 0x23, 0xa3, 0xaf, 0xa3, 0xf3, 0x25, 0x35, 0x85, 0xd5, 0xb0, 0x86,
	0xc9, 0x7d, 0xb1, 0x75, 0xaa, 0x73, 0x9b, 0x99, 0x25, 0xa6, 0xe2, 0x68, 0x9a, 0x14, 0xbd, 0xa1,
	0x09, 0x17, 0x15, 0x2d, 0x28, 0x96, 0xdb, 0x15, 0xdd, 0x6b, 0x56, 0x34, 0x76, 0x32, 0x9d, 0x42,
	0xec, 0x37, 0x08, 0x0f, 0x41, 0x5e, 0x1f, 0x5d, 0x47, 0x36, 0x32, 0x21, 0x5c, 0x0c, 0x62, 0xea,
	0x14, 0xad, 0xd0, 0x87, 0x2a, 0x8b, 0xe7, 0xd9, 0x22, 0xb5, 0xd4, 0x33, 0x56, 0x43, 0x1f, 0xa2,
	0x8a, 0x9c, 0x66, 0x0b, 0x4b, 0xdd, 0xa2, 0x1b, 0xb2, 0x80, 0x27, 0x0f, 0xd2, 0x6b, 0x6d, 0xe1,
	0xb5, 0x4e, 0xdf, 0x53, 0xa3, 0xe8, 0x86, 0x1e, 0x82, 0x79, 0x19, 0xc2, 0x6c, 0x0a, 0x26, 0xe7,
	0x8d, 0xef, 0x71, 0x5e, 0x7c, 0xac, 0x56, 0x91, 0xf2, 0xbf, 0x2b, 0xf2, 0xfe, 0xad, 0x8a, 0xdc,
	0xfb, 0x3d, 0x10, 0xb2, 0xa0, 0x18, 0x6f, 0x5b, 0x10, 0xed, 0x8d, 0x89, 0xe9, 0xc1, 0x1c, 0xd1,
	0x9c, 0xe8, 0x53, 0x70, 0xa5, 0x4e, 0xc1, 0x6a, 0x50, 0xb4, 0x6a, 0x83, 0x62, 0x47, 0x74, 0x8e,
	0x97, 0x4e, 0xe3, 0x86, 0x4b, 0x21, 0x97, 0xd4, 0x0c, 0xb3, 0x04, 0x1c, 0xbb, 0x2a, 0xa0, 0x19,
	0xfe, 0xda, 0xed, 0xf0, 0xff, 0x69, 0x8b, 0x7e, 0xbd, 0x82, 0xfc, 0x29, 0x17, 0xd4, 0xa7, 0xdc,
	0x63, 0xd1, 0x75, 0x86, 0x65, 0xf0, 0x15, 0xf0, 0xd9, 0xf0, 0x71, 0x7a, 0xe8, 0x04, 0xf8, 0xed,
	0xdb, 0x6e, 0x7a, 0x14, 0x00, 0x86, 0xe8, 0xb6, 0xa0, 0xe1, 0xc9, 0x57, 0xf0, 0x21, 0xec, 0xe8,
	0xd4, 0xee, 0xd6, 0x28, 0x18, 0x5a, 0x23, 0x1b, 0x4e, 0x75, 0x0c, 0xa9, 0xab, 0x06, 0x16, 0x0a,
	0xde, 0x76, 0xaa, 0x49, 0xf4, 0x54, 0xf4, 0x5f, 0x66, 0xe6, 0x26, 0x32, 0x71, 0xc1, 0xce, 0x2e,
	0x05, 0xd0, 0x40, 0x31, 0xf6, 0x51, 0x96, 0xe3, 0xb7, 0x00, 0x93, 0xde, 0x49, 0x34, 0x6b, 0x2c,
	0x8e, 0x04, 0x53, 0xf1, 0xbe, 0x04, 0x9a, 0xe9, 0xdd, 0xb8, 0xdd, 0xaf, 0x9f, 0x88, 0xcd, 0x13,
	0x48, 0xf4, 0x35, 0x18, 0x67, 0xc3, 0x53, 0xb2, 0x0e, 0x12, 0x03, 0x01, 0x78, 0x64, 0xf4, 0x1d,
	0x03, 0x9d, 0x8c, 0x37, 0x20, 0x63, 0x6d, 0x96, 0xb5, 0x71, 0xd9, 0x40, 0xb1, 0xf7, 0x17, 0xfd,
	0xbd, 0x7c, 0xa2, 0x6d, 0xee, 0xfd, 0x4d, 0x1c, 0xab, 0x06, 0x5f, 0xda, 0xab, 0x89, 0x56, 0xe8,
	0x21, 0x18, 0xcf, 0x8b, 0x58, 0xd7, 0x2a, 0xa2, 0x90, 0x31, 0xc7, 0x13, 0x9b, 0xb8, 0x4a, 0xc0,
	0xa5, 0xfc, 0x56, 0xf4, 0xf8, 0x39, 0x63, 0x9a, 0x04, 0x87, 0x55, 0x5f, 0xf6, 0xe0, 0xd0, 0xb7,
	0xd9, 0xfb, 0x14, 0x10, 0xeb, 0x3c, 0xa8, 0xce, 0x92, 0xa0, 0xc9, 0x92, 0x8a, 0x5b, 0x4c, 0xbb,
	0x82, 0x5b, 0x52, 0xb4, 0xbd, 0xa6, 0x4c, 0x6b, 0xe4, 0xc6, 0xcf, 0x3a, 0xb6, 0x57, 0xae, 0x56,
	0x58, 0xc0, 0x1d, 0x4e, 0x41, 0x5f, 0x5e, 0xd9, 0xe2, 0x2b, 0x8c, 0x25, 0xbc, 0xeb, 0xc9, 0xc2,
	0x44, 0x56, 0x67, 0xa9, 0xab, 0x8f, 0x52, 0xc6, 0x98, 0x5e, 0x7c, 0xb0, 0x90, 0xe2, 0xf0, 0x26,
	0xa6, 0x75, 0xc3, 0x0a, 0x40, 0xed, 0x5b, 0x93, 0x8c, 0xb3, 0x85, 0x39, 0x07, 0xe2, 0x5c, 0x37,
	0xac, 0x80, 0xbd, 0xbf, 0xda, 0xf4, 0xb5, 0x59, 0x4e, 0xcb, 0xea, 0x06, 0x41, 0xb3, 0xb8, 0xcb,
	0xe9, 0xb1, 0xd2, 0x98, 0x1e, 0x94, 0x13, 0x93, 0x5b, 0x6f, 0xb4, 0x54, 0x00, 0x7a, 0x62, 0x0b,
	0x26, 0xe5, 0x2a, 0x7b, 0x16, 0x72, 0xb3, 0xe3, 0x76, 0x6e, 0x77, 0xdc, 0xa7, 0xa2, 0x3f, 0x32,
	0xd9, 0x85, 0x4e, 0xa0, 0x98, 0x19, 0x3c, 0x0c, 0x1a, 0x28, 0x66, 0x73, 0x74, 0x95, 0xa5, 0xe0,
	0x0a, 0x83, 0x85, 0xaa, 0x1b, 0xf7, 0xfc, 0x6e, 0x4c, 0xd5, 0x96, 0x24, 0xd9, 0x4d, 0xd9, 0x6f,
	0x63, 0xde, 0xb3, 0x8e, 0x56, 0x76, 0x3a, 0xbd, 0x64, 0x3b, 0xf0, 0xed, 0x0a, 0x94, 0xe6, 0x4a,
	0x96, 0x5b, 0xb7, 0xd7, 0x85, 0xfb, 0x42, 0x2e, 0x11, 0xd4, 0x0f, 0x21, 0xd6, 0x11, 0xeb, 0x2f,
	0x59, 0x5f, 0x21, 0xf2, 0x47, 0xf1, 0x00, 0x73, 0xf9, 0x26, 0x4d, 0x74, 0x0a, 0x5c, 0x27, 0x2f,
	0xd2, 0xc5, 0x8c, 0x3f, 0x3b, 0xfa, 0x87, 0x0f, 0x0f, 0xee, 0xd2, 0x86, 0x77, 0xfa, 0xc8, 0xaf,
	0x44, 0x1f, 0xb3, 0x7b, 0x84, 0x9f, 0x7a, 0x40, 0x95, 0xf1, 0x91, 0x1f, 0xb2, 0x01, 0xcb, 0x67,
	0x42, 0xd6, 0x91, 0xf1, 0x55, 0x76, 0xa3, 0xfe, 0xe4, 0x56, 0x7a, 0x87, 0x4a, 0x3e, 0x13, 0x62,
	0xb8, 0xc4, 0x9b, 0xe3, 0xc5, 0xd5, 0x6f, 0x2b, 0x14, 0x5b, 0xff, 0xa0, 0xcc, 0x05, 0x05, 0xe5,
	0x99, 0x1c, 0xdf, 0x13, 0x1d, 0x6d, 0x0e, 0x66, 0xf9, 0xc1, 0x7c, 0x7a, 0xda, 0x1a, 0x05, 0xbf,
	0x04, 0x1f, 0xa6, 0x6b, 0xf4, 0x8b, 0xf3, 0xdd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0xca,
	0x07, 0x2a, 0x14, 0x0d, 0x00, 0x00,
}
