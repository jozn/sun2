// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_enum.proto

/*
Package x is a generated protocol buffer package.

It is generated from these files:
	pb_enum.proto
	pb_enums.proto
	pb_global.proto
	pb_rpc_auth2.proto
	pb_rpc_chat.proto
	pb_rpc_other.proto
	pb_rpc_page.proto
	pb_rpc_search.proto
	pb_rpc_social.proto
	pb_rpc_user.proto
	pb_tables.proto
	pb_view_social.proto
	pb_views.proto

It has these top-level messages:
	GeoLocation
	RoomMessageLog
	RoomMessageForwardFrom
	RoomDraft
	ChatRoom
	Pagination
	PB_CommandToServer
	PB_CommandToClient
	PB_CommandReachedToServer
	PB_CommandReachedToClient
	PB_ResponseToClient
	PB_ResponseExtra
	PB_Pager
	PB_UserParam_CheckUserName2
	PB_UserResponse_CheckUserName2
	PB_ChatParam_AddNewMessage
	PB_ChatResponse_AddNewMessage
	PB_ChatParam_SetRoomActionDoing
	PB_ChatResponse_SetRoomActionDoing
	PB_ChatParam_SetChatMessagesRangeAsSeen
	PB_ChatResponse_SetChatMessagesRangeAsSeen
	PB_ChatParam_DeleteChatHistory
	PB_ChatResponse_DeleteChatHistory
	PB_ChatParam_DeleteMessagesByIds
	PB_ChatResponse_DeleteMessagesByIds
	PB_ChatParam_SetMessagesAsReceived
	PB_ChatResponse_SetMessagesAsReceived
	PB_ChatParam_EditMessage
	PB_ChatResponse_EditMessage
	PB_ChatParam_GetChatList
	PB_ChatResponse_GetChatList
	PB_ChatParam_GetChatHistoryToOlder
	PB_ChatResponse_GetChatHistoryToOlder
	PB_ChatParam_GetFreshAllDirectMessagesList
	PB_ChatResponse_GetFreshAllDirectMessagesList
	PB_OtherParam_Echo
	PB_OtherResponse_Echo
	PB_PageParam_GetCommentsPage
	PB_PageResponse_GetCommentsPage
	PB_PageParam_GetHomePage
	PB_PageResponse_GetHomePage
	PB_PageParam_GetProfilePage
	PB_PageResponse_GetProfilePage
	PB_PageParam_GetLikesPage
	PB_PageResponse_GetLikesPage
	PB_PageParam_GetFollowersPage
	PB_PageResponse_GetFollowersPage
	PB_PageParam_GetFollowingsPage
	PB_PageResponse_GetFollowingsPage
	PB_PageParam_GetNotifiesPage
	PB_PageResponse_GetNotifiesPage
	PB_PageParam_GetUserActionsPage
	PB_PageResponse_GetUserActionsPage
	PB_PageParam_GetSuggestedPostsPage
	PB_PageResponse_GetSuggestedPostsPage
	PB_PageParam_GetSuggestedUsrsPage
	PB_PageResponse_GetSuggestedUsrsPage
	PB_PageParam_GetSuggestedTagsPage
	PB_PageResponse_GetSuggestedTagsPage
	PB_PageParam_GetLastPostsPage
	PB_PageResponse_GetLastPostsPage
	PB_PageParam_GetTagPage
	PB_PageResponse_GetTagPage
	PB_PageParam_SearchTagsPage
	PB_PageResponse_SearchTagsPage
	PB_PageParam_SearchUsersPage
	PB_PageResponse_SearchUsersPage
	PB_PageParam_
	PB_PageResponse_
	PB_SearchResponse_AddNewC
	PB_SocialParam_AddComment
	PB_SocialResponse_AddComment
	PB_SocialParam_DeleteComment
	PB_SocialResponse_DeleteComment
	PB_SocialParam_AddPost
	PB_SocialResponse_AddPost
	PB_SocialParam_EditPost
	PB_SocialResponse_EditPost
	PB_SocialParam_DeletePost
	PB_SocialResponse_DeletePost
	PB_SocialParam_ArchivePost
	PB_SocialResponse_ArchivePost
	PB_SocialParam_LikePost
	PB_SocialResponse_LikePost
	PB_SocialParam_UnLikePost
	PB_SocialResponse_UnLikePost
	PB_SocialParam_FollowUser
	PB_SocialResponse_FollowUser
	PB_SocialParam_UnFollowUser
	PB_SocialResponse_UnFollowUser
	PB_UserParam_BlockUser
	PB_UserResponse_BlockUser
	PB_UserParam_UnBlockUser
	PB_UserResponse_UnBlockUser
	PB_UserParam_BlockedList
	PB_UserResponse_BlockedList
	PB_UserParam_UpdateAbout
	PB_UserResponse_UpdateAbout
	PB_UserParam_UpdateUserName
	PB_UserResponse_UpdateUserName
	PB_UserParam_ChangeAvatar
	PB_UserResponse_ChangeAvatar
	PB_UserParam_ChangePrivacy
	PB_UserResponseOffline_ChangePrivacy
	PB_UserParam_CheckUserName
	PB_UserResponse_CheckUserName
	UserView
	PB_Action
	PB_Chat
	PB_Comment
	PB_DirectMessage
	PB_DirectOffline
	PB_DirectToMessage
	PB_Feed
	PB_FollowingList
	PB_FollowingListMember
	PB_FollowingListMemberHistory
	PB_GeneralLog
	PB_Group
	PB_GroupMember
	PB_GroupMessage
	PB_Like
	PB_Media
	PB_MessageFile
	PB_Notify
	PB_NotifyRemoved
	PB_PhoneContact
	PB_Post
	PB_RecommendUser
	PB_SearchClicked
	PB_Session
	PB_SettingClient
	PB_SettingNotification
	PB_Tag
	PB_TagsPost
	PB_TriggerLog
	PB_User
	PB_UserMetaInfo
	PB_UserPassword
	PB_PostView
	PB_MediaView
	PB_ActionView
	PB_NotifyView
	PB_CommentView
	PB_UserView
	PB_TopProfileView
	PB_ChatView
	PB_MessageView
	PB_MessageFileView
	PB_UserView3
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

type MediaTypeEnum int32

const (
	MediaTypeEnum_MEDIA_IMAGE MediaTypeEnum = 0
	MediaTypeEnum_MEDIA_VIDEO MediaTypeEnum = 1
)

var MediaTypeEnum_name = map[int32]string{
	0: "MEDIA_IMAGE",
	1: "MEDIA_VIDEO",
}
var MediaTypeEnum_value = map[string]int32{
	"MEDIA_IMAGE": 0,
	"MEDIA_VIDEO": 1,
}

func (x MediaTypeEnum) String() string {
	return proto.EnumName(MediaTypeEnum_name, int32(x))
}
func (MediaTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FollowingEnum int32

const (
	FollowingEnum_FOLLOWING_NONE FollowingEnum = 0
	FollowingEnum_FOLLOWING      FollowingEnum = 1
	FollowingEnum_REQUESTED      FollowingEnum = 2
	FollowingEnum_BLOCKED        FollowingEnum = 3
)

var FollowingEnum_name = map[int32]string{
	0: "FOLLOWING_NONE",
	1: "FOLLOWING",
	2: "REQUESTED",
	3: "BLOCKED",
}
var FollowingEnum_value = map[string]int32{
	"FOLLOWING_NONE": 0,
	"FOLLOWING":      1,
	"REQUESTED":      2,
	"BLOCKED":        3,
}

func (x FollowingEnum) String() string {
	return proto.EnumName(FollowingEnum_name, int32(x))
}
func (FollowingEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type UserTypeEnum int32

const (
	UserTypeEnum_USER    UserTypeEnum = 0
	UserTypeEnum_CHANNEL UserTypeEnum = 1
)

var UserTypeEnum_name = map[int32]string{
	0: "USER",
	1: "CHANNEL",
}
var UserTypeEnum_value = map[string]int32{
	"USER":    0,
	"CHANNEL": 1,
}

func (x UserTypeEnum) String() string {
	return proto.EnumName(UserTypeEnum_name, int32(x))
}
func (UserTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type UserLevelEnum int32

const (
	UserLevelEnum_LEVEL_NORMAL      UserLevelEnum = 0
	UserLevelEnum_APP_ADMIN         UserLevelEnum = 1
	UserLevelEnum_SUSPONDED         UserLevelEnum = 2
	UserLevelEnum_DELETED_BY_OWENER UserLevelEnum = 3
	UserLevelEnum_DELETED_IRAN      UserLevelEnum = 4
	UserLevelEnum_SUSPONDED_IRAN    UserLevelEnum = 5
)

var UserLevelEnum_name = map[int32]string{
	0: "LEVEL_NORMAL",
	1: "APP_ADMIN",
	2: "SUSPONDED",
	3: "DELETED_BY_OWENER",
	4: "DELETED_IRAN",
	5: "SUSPONDED_IRAN",
}
var UserLevelEnum_value = map[string]int32{
	"LEVEL_NORMAL":      0,
	"APP_ADMIN":         1,
	"SUSPONDED":         2,
	"DELETED_BY_OWENER": 3,
	"DELETED_IRAN":      4,
	"SUSPONDED_IRAN":    5,
}

func (x UserLevelEnum) String() string {
	return proto.EnumName(UserLevelEnum_name, int32(x))
}
func (UserLevelEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type UserOnlinePrivacyEnum int32

const (
	UserOnlinePrivacyEnum_E4 UserOnlinePrivacyEnum = 0
)

var UserOnlinePrivacyEnum_name = map[int32]string{
	0: "E4",
}
var UserOnlinePrivacyEnum_value = map[string]int32{
	"E4": 0,
}

func (x UserOnlinePrivacyEnum) String() string {
	return proto.EnumName(UserOnlinePrivacyEnum_name, int32(x))
}
func (UserOnlinePrivacyEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type UserOnlineStatusEnum int32

const (
	UserOnlineStatusEnum_EXACTLY       UserOnlineStatusEnum = 0
	UserOnlineStatusEnum_ONLINE        UserOnlineStatusEnum = 1
	UserOnlineStatusEnum_CONNECTED     UserOnlineStatusEnum = 2
	UserOnlineStatusEnum_FEW_DAYS_AGO  UserOnlineStatusEnum = 3
	UserOnlineStatusEnum_RECENTLY      UserOnlineStatusEnum = 4
	UserOnlineStatusEnum_LAST_WEEK     UserOnlineStatusEnum = 5
	UserOnlineStatusEnum_LAST_MONTH    UserOnlineStatusEnum = 6
	UserOnlineStatusEnum_LONG_TIME_AGO UserOnlineStatusEnum = 7
	UserOnlineStatusEnum_HIDE          UserOnlineStatusEnum = 8
)

var UserOnlineStatusEnum_name = map[int32]string{
	0: "EXACTLY",
	1: "ONLINE",
	2: "CONNECTED",
	3: "FEW_DAYS_AGO",
	4: "RECENTLY",
	5: "LAST_WEEK",
	6: "LAST_MONTH",
	7: "LONG_TIME_AGO",
	8: "HIDE",
}
var UserOnlineStatusEnum_value = map[string]int32{
	"EXACTLY":       0,
	"ONLINE":        1,
	"CONNECTED":     2,
	"FEW_DAYS_AGO":  3,
	"RECENTLY":      4,
	"LAST_WEEK":     5,
	"LAST_MONTH":    6,
	"LONG_TIME_AGO": 7,
	"HIDE":          8,
}

func (x UserOnlineStatusEnum) String() string {
	return proto.EnumName(UserOnlineStatusEnum_name, int32(x))
}
func (UserOnlineStatusEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type PostTypeEnum int32

const (
	PostTypeEnum_POST_RESHARED PostTypeEnum = 0
	PostTypeEnum_POST_TEXT     PostTypeEnum = 1
	PostTypeEnum_POST_PHOTO    PostTypeEnum = 2
	PostTypeEnum_POST_VIDEO    PostTypeEnum = 3
	PostTypeEnum_POST_GIFS     PostTypeEnum = 4
	PostTypeEnum_POST_AUDIO    PostTypeEnum = 5
	PostTypeEnum_POST_GIF      PostTypeEnum = 6
	PostTypeEnum_POST_FILE     PostTypeEnum = 7
	PostTypeEnum_POST_POLL     PostTypeEnum = 8
)

var PostTypeEnum_name = map[int32]string{
	0: "POST_RESHARED",
	1: "POST_TEXT",
	2: "POST_PHOTO",
	3: "POST_VIDEO",
	4: "POST_GIFS",
	5: "POST_AUDIO",
	6: "POST_GIF",
	7: "POST_FILE",
	8: "POST_POLL",
}
var PostTypeEnum_value = map[string]int32{
	"POST_RESHARED": 0,
	"POST_TEXT":     1,
	"POST_PHOTO":    2,
	"POST_VIDEO":    3,
	"POST_GIFS":     4,
	"POST_AUDIO":    5,
	"POST_GIF":      6,
	"POST_FILE":     7,
	"POST_POLL":     8,
}

func (x PostTypeEnum) String() string {
	return proto.EnumName(PostTypeEnum_name, int32(x))
}
func (PostTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type NotifyEnum int32

const (
	NotifyEnum_NOTIFY_POST_LIKED     NotifyEnum = 0
	NotifyEnum_NOTIFY_POST_COMMENTED NotifyEnum = 1
	NotifyEnum_NOTIFY_FOLLOWED_YOU   NotifyEnum = 2
)

var NotifyEnum_name = map[int32]string{
	0: "NOTIFY_POST_LIKED",
	1: "NOTIFY_POST_COMMENTED",
	2: "NOTIFY_FOLLOWED_YOU",
}
var NotifyEnum_value = map[string]int32{
	"NOTIFY_POST_LIKED":     0,
	"NOTIFY_POST_COMMENTED": 1,
	"NOTIFY_FOLLOWED_YOU":   2,
}

func (x NotifyEnum) String() string {
	return proto.EnumName(NotifyEnum_name, int32(x))
}
func (NotifyEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ActionEnum int32

const (
	ActionEnum_ACTION_POST_LIKED     ActionEnum = 0
	ActionEnum_ACTION_POST_COMMENTED ActionEnum = 1
	ActionEnum_ACTION_FOLLOWED_USER  ActionEnum = 2
)

var ActionEnum_name = map[int32]string{
	0: "ACTION_POST_LIKED",
	1: "ACTION_POST_COMMENTED",
	2: "ACTION_FOLLOWED_USER",
}
var ActionEnum_value = map[string]int32{
	"ACTION_POST_LIKED":     0,
	"ACTION_POST_COMMENTED": 1,
	"ACTION_FOLLOWED_USER":  2,
}

func (x ActionEnum) String() string {
	return proto.EnumName(ActionEnum_name, int32(x))
}
func (ActionEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterEnum("MediaTypeEnum", MediaTypeEnum_name, MediaTypeEnum_value)
	proto.RegisterEnum("FollowingEnum", FollowingEnum_name, FollowingEnum_value)
	proto.RegisterEnum("UserTypeEnum", UserTypeEnum_name, UserTypeEnum_value)
	proto.RegisterEnum("UserLevelEnum", UserLevelEnum_name, UserLevelEnum_value)
	proto.RegisterEnum("UserOnlinePrivacyEnum", UserOnlinePrivacyEnum_name, UserOnlinePrivacyEnum_value)
	proto.RegisterEnum("UserOnlineStatusEnum", UserOnlineStatusEnum_name, UserOnlineStatusEnum_value)
	proto.RegisterEnum("PostTypeEnum", PostTypeEnum_name, PostTypeEnum_value)
	proto.RegisterEnum("NotifyEnum", NotifyEnum_name, NotifyEnum_value)
	proto.RegisterEnum("ActionEnum", ActionEnum_name, ActionEnum_value)
}

func init() { proto.RegisterFile("pb_enum.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x5d, 0x6f, 0xda, 0x30,
	0x14, 0x25, 0x40, 0x81, 0xdd, 0x42, 0x6b, 0xb2, 0xa2, 0x7d, 0xbc, 0xec, 0x69, 0x2f, 0x79, 0xa8,
	0x34, 0x6d, 0x7f, 0xc0, 0x8d, 0x2f, 0x60, 0xd5, 0xb1, 0xb3, 0xc4, 0x94, 0xb2, 0x17, 0x0b, 0xba,
	0x6c, 0x8a, 0x44, 0x09, 0x82, 0xb4, 0x5d, 0xa5, 0xfd, 0x92, 0xed, 0xcf, 0x4e, 0x76, 0x68, 0x5a,
	0x69, 0x8f, 0xe7, 0xdc, 0x73, 0xcf, 0x3d, 0x39, 0x72, 0x60, 0xb0, 0x5d, 0x99, 0x6c, 0x73, 0x77,
	0x7b, 0xbe, 0xdd, 0x15, 0x65, 0xf1, 0xfe, 0x74, 0xbb, 0x32, 0xe5, 0x72, 0xb5, 0xce, 0xf6, 0x07,
	0xe2, 0x64, 0xbb, 0x32, 0xf7, 0x79, 0xf6, 0x70, 0xc0, 0xc1, 0x27, 0x18, 0x44, 0xd9, 0xf7, 0x7c,
	0xa9, 0x1f, 0xb7, 0x19, 0x6e, 0xee, 0x6e, 0xfd, 0x53, 0x38, 0x8e, 0x90, 0x71, 0x6a, 0x78, 0x44,
	0x27, 0x48, 0x1a, 0xcf, 0xc4, 0x15, 0x67, 0xa8, 0x88, 0x17, 0x48, 0x18, 0x8c, 0x8b, 0xf5, 0xba,
	0x78, 0xc8, 0x37, 0x3f, 0xdd, 0x8a, 0x0f, 0x27, 0x63, 0x25, 0x84, 0x9a, 0x73, 0x39, 0x31, 0x52,
	0x49, 0xbb, 0x35, 0x80, 0x57, 0x35, 0x47, 0x3c, 0x0b, 0x13, 0xfc, 0x3a, 0xc3, 0x54, 0x23, 0x23,
	0x4d, 0xff, 0x18, 0xba, 0x17, 0x42, 0x85, 0x97, 0xc8, 0x48, 0x2b, 0xf8, 0x08, 0xfd, 0xd9, 0x3e,
	0xdb, 0xd5, 0x09, 0x7a, 0xd0, 0x9e, 0xa5, 0x98, 0x90, 0x86, 0x95, 0x85, 0x53, 0x2a, 0x25, 0x0a,
	0xe2, 0x05, 0xbf, 0x61, 0x60, 0x65, 0x22, 0xbb, 0xcf, 0xd6, 0x4e, 0x47, 0xa0, 0x2f, 0xf0, 0x0a,
	0x85, 0x91, 0x2a, 0x89, 0xa8, 0xa8, 0x8e, 0xd2, 0x38, 0x36, 0x94, 0x45, 0x5c, 0x56, 0x47, 0xd3,
	0x59, 0x1a, 0x2b, 0xc9, 0xdc, 0xd1, 0x11, 0x0c, 0x19, 0x0a, 0xd4, 0xc8, 0xcc, 0xc5, 0xc2, 0xa8,
	0x39, 0x4a, 0x4c, 0x48, 0xcb, 0xda, 0x3c, 0xd1, 0x3c, 0xa1, 0x92, 0xb4, 0xed, 0xf7, 0xd4, 0x7b,
	0x15, 0x77, 0x14, 0x7c, 0x80, 0x91, 0xbd, 0xae, 0x36, 0xeb, 0x7c, 0x93, 0xc5, 0xbb, 0xfc, 0x7e,
	0x79, 0xf3, 0xe8, 0x52, 0x74, 0xa0, 0x89, 0x5f, 0x48, 0x23, 0xf8, 0xe3, 0xc1, 0xd9, 0xb3, 0x22,
	0x2d, 0x97, 0xe5, 0xdd, 0xde, 0x09, 0x8e, 0xa1, 0x8b, 0xd7, 0x34, 0xd4, 0x62, 0x41, 0x1a, 0x3e,
	0x40, 0x47, 0x49, 0xc1, 0x25, 0x56, 0xf1, 0x42, 0x25, 0x25, 0x86, 0x55, 0x27, 0x04, 0xfa, 0x63,
	0x9c, 0x1b, 0x46, 0x17, 0xa9, 0xa1, 0x13, 0x45, 0x5a, 0x7e, 0x1f, 0x7a, 0x09, 0x86, 0x28, 0xed,
	0x6a, 0xdb, 0xca, 0x05, 0x4d, 0xb5, 0x99, 0x23, 0x5e, 0x92, 0x23, 0xff, 0x04, 0xc0, 0xc1, 0x48,
	0x49, 0x3d, 0x25, 0x1d, 0x7f, 0x08, 0x03, 0xa1, 0xe4, 0xc4, 0x68, 0x1e, 0xa1, 0xdb, 0xef, 0xda,
	0x22, 0xa7, 0x9c, 0x21, 0xe9, 0x05, 0x7f, 0x3d, 0xe8, 0xc7, 0xc5, 0xbe, 0xac, 0x3b, 0x1e, 0xc2,
	0x20, 0x56, 0xa9, 0x36, 0x09, 0xa6, 0x53, 0x9a, 0x20, 0xab, 0xca, 0x73, 0x94, 0xc6, 0x6b, 0x4d,
	0x3c, 0xeb, 0xef, 0x60, 0x3c, 0x55, 0x5a, 0x91, 0x66, 0x8d, 0xab, 0x57, 0xd0, 0xaa, 0xe5, 0x13,
	0x3e, 0x4e, 0x49, 0xbb, 0x1e, 0xd3, 0x19, 0xe3, 0x8a, 0x1c, 0xd9, 0xec, 0x4f, 0x63, 0xd2, 0xa9,
	0xc5, 0x63, 0x2e, 0x90, 0x74, 0x6b, 0x18, 0x2b, 0x21, 0x48, 0x2f, 0x98, 0x03, 0xc8, 0xa2, 0xcc,
	0x7f, 0x54, 0x85, 0x8e, 0x60, 0x28, 0x95, 0xe6, 0xe3, 0x85, 0x71, 0x1a, 0xc1, 0x2f, 0x5d, 0xbc,
	0x77, 0x30, 0x7a, 0x49, 0x87, 0x2a, 0x8a, 0x50, 0xda, 0xe6, 0x3c, 0xff, 0x0d, 0xbc, 0x3e, 0x8c,
	0xaa, 0x27, 0x87, 0xcc, 0x2c, 0xd4, 0x8c, 0x34, 0x83, 0x6b, 0x00, 0x7a, 0x53, 0xe6, 0xc5, 0xe6,
	0xc9, 0x98, 0x86, 0x9a, 0x2b, 0xf9, 0x9f, 0xf1, 0x4b, 0xfa, 0xa5, 0xf1, 0x5b, 0x38, 0x3b, 0x8c,
	0x6a, 0x63, 0xf7, 0x32, 0x9b, 0x17, 0x43, 0xe8, 0xe5, 0xbb, 0xf3, 0xdb, 0xfd, 0xf9, 0x76, 0x35,
	0x6d, 0xc5, 0xde, 0x37, 0xef, 0xd7, 0xaa, 0xe3, 0x7e, 0xa8, 0xcf, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x80, 0x98, 0x4c, 0xcd, 0x82, 0x03, 0x00, 0x00,
}
