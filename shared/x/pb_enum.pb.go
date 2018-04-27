// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_enum.proto

/*
Package x is a generated protocol buffer package.

It is generated from these files:
	pb_enum.proto
	pb_global.proto
	pb_push.proto
	pb_rpc_auth2.proto
	pb_rpc_chat.proto
	pb_rpc_other.proto
	pb_rpc_page.proto
	pb_rpc_search.proto
	pb_rpc_social.proto
	pb_rpc_user.proto
	pb_tables.proto
	pb_updates.proto
	pb_views.proto

It has these top-level messages:
	PB_CommandToServer
	PB_CommandToClient
	PB_CommandReachedToServer
	PB_CommandReachedToClient
	PB_ResponseToClient
	PB_ResponseExtra
	PB_Pager
	PB_UpdateRoomActionDoing2
	PB_UpdateMessageMeta2
	PB_Push
	PB_UserParam_CheckUserName2
	PB_UserResponse_CheckUserName2
	PB_ChatParam_AddNewMessage
	PB_ChatResponse_AddNewMessage
	PB_ChatParam_SetRoomActionDoing
	PB_ChatResponse_SetRoomActionDoing
	PB_ChatParam_SetMessagesAsReceived
	PB_ChatResponse_SetMessagesAsReceived
	PB_ChatParam_SetChatMessagesRangeAsSeen
	PB_ChatResponse_SetChatMessagesRangeAsSeen
	PB_ChatParam_DeleteChatHistory
	PB_ChatResponse_DeleteChatHistory
	PB_ChatParam_DeleteMessagesByIds
	PB_ChatResponse_DeleteMessagesByIds
	PB_ChatParam_EditMessage
	PB_ChatResponse_EditMessage
	PB_ChatParam_GetChatList
	PB_ChatResponse_GetChatList
	PB_ChatParam_GetChatHistoryToOlder
	PB_ChatResponse_GetChatHistoryToOlder
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
	PB_PageParam_GetSuggestedUsersPage
	PB_PageResponse_GetSuggestedUsersPage
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
	PB_Comment
	PB_CommentDeleted
	PB_Event
	PB_FollowingList
	PB_FollowingListMember
	PB_FollowingListMemberRemoved
	PB_Like
	PB_Notify
	PB_NotifyRemoved
	PB_PhoneContact
	PB_Post
	PB_PostCount
	PB_PostDeleted
	PB_PostKey
	PB_PostLink
	PB_PostMedia
	PB_PostMentioned
	PB_PostReshared
	PB_SearchClicked
	PB_Session
	PB_SettingClient
	PB_SettingNotification
	PB_Tag
	PB_TagPost
	PB_TriggerLog
	PB_User
	PB_UserMetaInfo
	PB_UserPassword
	PB_Chat
	PB_ChatLastMessage
	PB_DirectMessage
	PB_Group
	PB_GroupMember
	PB_GroupMessage
	PB_FileMsg
	PB_FilePost
	PB_ActionFanout
	PB_HomeFanout
	PB_SuggestedTopPost
	PB_SuggestedUser
	PB_ChatSync2
	PB_PushChat
	PB_PushChat2
	PB_UpdateRoomActionDoing
	PB_UpdateMessageMeta
	PB_Updates
	PB_PostView
	PB_MediaView
	PB_ActionView
	PB_NotifyView
	PB_CommentView
	PB_UserView
	PB_TopProfileView
	PB_UserViewRowify
	PB_PostViewRowify
	PB_TagView
	PB_TopTagWithSamplePosts
	PB_ChatView
	PB_MessageView
	PB_MessageFileView
	PB_MessageTableExtra
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

// dep
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

type PostCategoryEnum int32

const (
	PostCategoryEnum_PostCat_Text  PostCategoryEnum = 0
	PostCategoryEnum_PostCat_Media PostCategoryEnum = 1
	PostCategoryEnum_PostCat_File  PostCategoryEnum = 2
)

var PostCategoryEnum_name = map[int32]string{
	0: "PostCat_Text",
	1: "PostCat_Media",
	2: "PostCat_File",
}
var PostCategoryEnum_value = map[string]int32{
	"PostCat_Text":  0,
	"PostCat_Media": 1,
	"PostCat_File":  2,
}

func (x PostCategoryEnum) String() string {
	return proto.EnumName(PostCategoryEnum_name, int32(x))
}
func (PostCategoryEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
func (NotifyEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
func (ActionEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type ProfilePrivacyLevelEnum int32

const (
	ProfilePrivacyLevelEnum_NONE ProfilePrivacyLevelEnum = 0
)

var ProfilePrivacyLevelEnum_name = map[int32]string{
	0: "NONE",
}
var ProfilePrivacyLevelEnum_value = map[string]int32{
	"NONE": 0,
}

func (x ProfilePrivacyLevelEnum) String() string {
	return proto.EnumName(ProfilePrivacyLevelEnum_name, int32(x))
}
func (ProfilePrivacyLevelEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type RoomTypeEnum int32

const (
	RoomTypeEnum_UNKNOWN_ROOM_TYPE RoomTypeEnum = 0
	RoomTypeEnum_DIRECT            RoomTypeEnum = 1
	RoomTypeEnum_GROUP             RoomTypeEnum = 2
	//    CHANNEL = 3;
	RoomTypeEnum_BROADCAST RoomTypeEnum = 3
)

var RoomTypeEnum_name = map[int32]string{
	0: "UNKNOWN_ROOM_TYPE",
	1: "DIRECT",
	2: "GROUP",
	3: "BROADCAST",
}
var RoomTypeEnum_value = map[string]int32{
	"UNKNOWN_ROOM_TYPE": 0,
	"DIRECT":            1,
	"GROUP":             2,
	"BROADCAST":         3,
}

func (x RoomTypeEnum) String() string {
	return proto.EnumName(RoomTypeEnum_name, int32(x))
}
func (RoomTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type RoomActionDoingEnum int32

const (
	RoomActionDoingEnum_UNKNOWN_ROOM_ACTION_DOING RoomActionDoingEnum = 0
	RoomActionDoingEnum_CANCEL                    RoomActionDoingEnum = 1
	RoomActionDoingEnum_TYPING                    RoomActionDoingEnum = 2
	RoomActionDoingEnum_SENDING_IMAGE             RoomActionDoingEnum = 3
	RoomActionDoingEnum_CAPTURING_IMAGE           RoomActionDoingEnum = 4
	RoomActionDoingEnum_SENDING_VIDEO             RoomActionDoingEnum = 5
	RoomActionDoingEnum_CAPTURING_VIDEO           RoomActionDoingEnum = 6
	RoomActionDoingEnum_SENDING_AUDIO             RoomActionDoingEnum = 7
	RoomActionDoingEnum_RECORDING_VOICE           RoomActionDoingEnum = 8
	RoomActionDoingEnum_SENDING_VOICE             RoomActionDoingEnum = 9
	RoomActionDoingEnum_SENDING_DOCUMENT          RoomActionDoingEnum = 11
	RoomActionDoingEnum_SENDING_GIF               RoomActionDoingEnum = 12
	RoomActionDoingEnum_SENDING_FILE              RoomActionDoingEnum = 13
	RoomActionDoingEnum_SENDING_LOCATION          RoomActionDoingEnum = 14
	RoomActionDoingEnum_CHOOSING_CONTACT          RoomActionDoingEnum = 15
	RoomActionDoingEnum_PAINTING                  RoomActionDoingEnum = 16
)

var RoomActionDoingEnum_name = map[int32]string{
	0:  "UNKNOWN_ROOM_ACTION_DOING",
	1:  "CANCEL",
	2:  "TYPING",
	3:  "SENDING_IMAGE",
	4:  "CAPTURING_IMAGE",
	5:  "SENDING_VIDEO",
	6:  "CAPTURING_VIDEO",
	7:  "SENDING_AUDIO",
	8:  "RECORDING_VOICE",
	9:  "SENDING_VOICE",
	11: "SENDING_DOCUMENT",
	12: "SENDING_GIF",
	13: "SENDING_FILE",
	14: "SENDING_LOCATION",
	15: "CHOOSING_CONTACT",
	16: "PAINTING",
}
var RoomActionDoingEnum_value = map[string]int32{
	"UNKNOWN_ROOM_ACTION_DOING": 0,
	"CANCEL":                    1,
	"TYPING":                    2,
	"SENDING_IMAGE":             3,
	"CAPTURING_IMAGE":           4,
	"SENDING_VIDEO":             5,
	"CAPTURING_VIDEO":           6,
	"SENDING_AUDIO":             7,
	"RECORDING_VOICE":           8,
	"SENDING_VOICE":             9,
	"SENDING_DOCUMENT":          11,
	"SENDING_GIF":               12,
	"SENDING_FILE":              13,
	"SENDING_LOCATION":          14,
	"CHOOSING_CONTACT":          15,
	"PAINTING":                  16,
}

func (x RoomActionDoingEnum) String() string {
	return proto.EnumName(RoomActionDoingEnum_name, int32(x))
}
func (RoomActionDoingEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type RoomMessageDeliviryStatusEnum int32

const (
	RoomMessageDeliviryStatusEnum_UNKNOWN_MESSAGE_DELIVIRY RoomMessageDeliviryStatusEnum = 0
	RoomMessageDeliviryStatusEnum_NEED_TO_SINK             RoomMessageDeliviryStatusEnum = 1
	RoomMessageDeliviryStatusEnum_FAILED                   RoomMessageDeliviryStatusEnum = 2
	RoomMessageDeliviryStatusEnum_SENDING                  RoomMessageDeliviryStatusEnum = 3
	RoomMessageDeliviryStatusEnum_SENT                     RoomMessageDeliviryStatusEnum = 4
	RoomMessageDeliviryStatusEnum_DELIVERED                RoomMessageDeliviryStatusEnum = 5
	RoomMessageDeliviryStatusEnum_SEEN                     RoomMessageDeliviryStatusEnum = 6
)

var RoomMessageDeliviryStatusEnum_name = map[int32]string{
	0: "UNKNOWN_MESSAGE_DELIVIRY",
	1: "NEED_TO_SINK",
	2: "FAILED",
	3: "SENDING",
	4: "SENT",
	5: "DELIVERED",
	6: "SEEN",
}
var RoomMessageDeliviryStatusEnum_value = map[string]int32{
	"UNKNOWN_MESSAGE_DELIVIRY": 0,
	"NEED_TO_SINK":             1,
	"FAILED":                   2,
	"SENDING":                  3,
	"SENT":                     4,
	"DELIVERED":                5,
	"SEEN":                     6,
}

func (x RoomMessageDeliviryStatusEnum) String() string {
	return proto.EnumName(RoomMessageDeliviryStatusEnum_name, int32(x))
}
func (RoomMessageDeliviryStatusEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{13}
}

func init() {
	proto.RegisterEnum("MediaTypeEnum", MediaTypeEnum_name, MediaTypeEnum_value)
	proto.RegisterEnum("FollowingEnum", FollowingEnum_name, FollowingEnum_value)
	proto.RegisterEnum("UserTypeEnum", UserTypeEnum_name, UserTypeEnum_value)
	proto.RegisterEnum("UserLevelEnum", UserLevelEnum_name, UserLevelEnum_value)
	proto.RegisterEnum("UserOnlinePrivacyEnum", UserOnlinePrivacyEnum_name, UserOnlinePrivacyEnum_value)
	proto.RegisterEnum("UserOnlineStatusEnum", UserOnlineStatusEnum_name, UserOnlineStatusEnum_value)
	proto.RegisterEnum("PostTypeEnum", PostTypeEnum_name, PostTypeEnum_value)
	proto.RegisterEnum("PostCategoryEnum", PostCategoryEnum_name, PostCategoryEnum_value)
	proto.RegisterEnum("NotifyEnum", NotifyEnum_name, NotifyEnum_value)
	proto.RegisterEnum("ActionEnum", ActionEnum_name, ActionEnum_value)
	proto.RegisterEnum("ProfilePrivacyLevelEnum", ProfilePrivacyLevelEnum_name, ProfilePrivacyLevelEnum_value)
	proto.RegisterEnum("RoomTypeEnum", RoomTypeEnum_name, RoomTypeEnum_value)
	proto.RegisterEnum("RoomActionDoingEnum", RoomActionDoingEnum_name, RoomActionDoingEnum_value)
	proto.RegisterEnum("RoomMessageDeliviryStatusEnum", RoomMessageDeliviryStatusEnum_name, RoomMessageDeliviryStatusEnum_value)
}

func init() { proto.RegisterFile("pb_enum.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 874 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x54, 0x5d, 0x6e, 0xe3, 0x36,
	0x10, 0x8e, 0xfc, 0x17, 0x67, 0x62, 0x27, 0x63, 0x27, 0xc1, 0xee, 0x02, 0x5d, 0xf4, 0xa1, 0xe8,
	0x8b, 0x1e, 0x02, 0x14, 0xed, 0x05, 0x18, 0x72, 0x6c, 0x13, 0x96, 0x48, 0x55, 0xa2, 0xe3, 0xb8,
	0x2f, 0x82, 0xb3, 0xd5, 0x06, 0x02, 0x1c, 0x2b, 0xb0, 0x9d, 0x74, 0x03, 0xf4, 0x06, 0xbd, 0x41,
	0x7b, 0xa3, 0x9e, 0xaa, 0x20, 0x65, 0x29, 0x0e, 0xfa, 0x26, 0x7d, 0xfc, 0xe6, 0xe3, 0xcc, 0xf0,
	0x9b, 0x81, 0xfe, 0xd3, 0x7d, 0x9a, 0xad, 0x9f, 0x1f, 0xaf, 0x9f, 0x36, 0xc5, 0xae, 0xf0, 0x7f,
	0x82, 0x7e, 0x98, 0xfd, 0x9e, 0x2f, 0xcd, 0xeb, 0x53, 0x46, 0xeb, 0xe7, 0xc7, 0xe1, 0x39, 0x9c,
	0x86, 0x24, 0x24, 0x4b, 0x65, 0xc8, 0xc6, 0x84, 0x47, 0x6f, 0xc0, 0xad, 0x14, 0xa4, 0xd1, 0xf3,
	0x15, 0xf4, 0x47, 0xc5, 0x6a, 0x55, 0xfc, 0x91, 0xaf, 0x1f, 0x5c, 0xc8, 0x10, 0xce, 0x46, 0x3a,
	0x08, 0xf4, 0x5c, 0xaa, 0x71, 0xaa, 0xb4, 0xb2, 0x51, 0x7d, 0x38, 0xa9, 0x31, 0xf4, 0xec, 0x6f,
	0x4c, 0xbf, 0xce, 0x28, 0x31, 0x24, 0xb0, 0x31, 0x3c, 0x85, 0xe3, 0x9b, 0x40, 0xf3, 0x29, 0x09,
	0x6c, 0xfa, 0x3f, 0x42, 0x6f, 0xb6, 0xcd, 0x36, 0x75, 0x06, 0x5d, 0x68, 0xcd, 0x12, 0x8a, 0xf1,
	0xc8, 0xd2, 0xf8, 0x84, 0x29, 0x45, 0x01, 0x7a, 0xfe, 0x9f, 0xd0, 0xb7, 0xb4, 0x20, 0x7b, 0xc9,
	0x56, 0x8e, 0x87, 0xd0, 0x0b, 0xe8, 0x96, 0x82, 0x54, 0xe9, 0x38, 0x64, 0x41, 0x79, 0x29, 0x8b,
	0xa2, 0x94, 0x89, 0x50, 0xaa, 0xf2, 0xd2, 0x64, 0x96, 0x44, 0x5a, 0x09, 0x77, 0xe9, 0x15, 0x0c,
	0x04, 0x05, 0x64, 0x48, 0xa4, 0x37, 0x8b, 0x54, 0xcf, 0x49, 0x51, 0x8c, 0x4d, 0x2b, 0x53, 0xc1,
	0x32, 0x66, 0x0a, 0x5b, 0xb6, 0x9e, 0x3a, 0xae, 0xc4, 0xda, 0xfe, 0xf7, 0x70, 0x65, 0x6f, 0xd7,
	0xeb, 0x55, 0xbe, 0xce, 0xa2, 0x4d, 0xfe, 0xb2, 0xfc, 0xf2, 0xea, 0xb2, 0xe8, 0x40, 0x83, 0x7e,
	0xc1, 0x23, 0xff, 0x6f, 0x0f, 0x2e, 0xdf, 0x18, 0xc9, 0x6e, 0xb9, 0x7b, 0xde, 0x3a, 0xc2, 0x29,
	0x1c, 0xd3, 0x1d, 0xe3, 0x26, 0x58, 0xe0, 0xd1, 0x10, 0xa0, 0xa3, 0x55, 0x20, 0x15, 0x95, 0xe9,
	0x71, 0xad, 0x14, 0xf1, 0xb2, 0x27, 0x08, 0xbd, 0x11, 0xcd, 0x53, 0xc1, 0x16, 0x49, 0xca, 0xc6,
	0x1a, 0x9b, 0xc3, 0x1e, 0x74, 0x63, 0xe2, 0xa4, 0x6c, 0x68, 0xcb, 0xd2, 0x03, 0x96, 0x98, 0x74,
	0x4e, 0x34, 0xc5, 0xf6, 0xf0, 0x0c, 0xc0, 0xfd, 0x86, 0x5a, 0x99, 0x09, 0x76, 0x86, 0x03, 0xe8,
	0x07, 0x5a, 0x8d, 0x53, 0x23, 0x43, 0x72, 0xf1, 0xc7, 0xb6, 0x91, 0x13, 0x29, 0x08, 0xbb, 0xfe,
	0x3f, 0x1e, 0xf4, 0xa2, 0x62, 0xbb, 0xab, 0x7b, 0x3c, 0x80, 0x7e, 0xa4, 0x13, 0x93, 0xc6, 0x94,
	0x4c, 0x58, 0x4c, 0xa2, 0x6c, 0x9e, 0x83, 0x0c, 0xdd, 0x19, 0xf4, 0xac, 0xbe, 0xfb, 0x8d, 0x26,
	0xda, 0x68, 0x6c, 0xd4, 0xff, 0xa5, 0x0b, 0x9a, 0x35, 0x7d, 0x2c, 0x47, 0x09, 0xb6, 0xea, 0x63,
	0x36, 0x13, 0x52, 0x63, 0xdb, 0xe6, 0x5e, 0x1d, 0x63, 0xa7, 0x26, 0x8f, 0x64, 0x40, 0x78, 0x5c,
	0xff, 0x46, 0x3a, 0x08, 0xb0, 0xeb, 0x4b, 0x40, 0x9b, 0x1c, 0x5f, 0xee, 0xb2, 0x87, 0x62, 0xf3,
	0x5a, 0x3d, 0xee, 0x1e, 0x4b, 0x4d, 0xf6, 0x6d, 0x87, 0x47, 0x2e, 0xe5, 0x3d, 0xe2, 0x1c, 0x8b,
	0xde, 0x21, 0x69, 0x94, 0xaf, 0x32, 0x6c, 0xf8, 0x73, 0x00, 0x55, 0xec, 0xf2, 0xaf, 0xa5, 0xc8,
	0x15, 0x0c, 0x94, 0x36, 0x72, 0xb4, 0x48, 0xdd, 0x75, 0x81, 0x9c, 0xba, 0x4a, 0x3f, 0xc1, 0xd5,
	0x21, 0xcc, 0x75, 0x18, 0x92, 0xb2, 0x8f, 0xe0, 0x0d, 0x3f, 0xc0, 0xc5, 0xfe, 0xa8, 0x74, 0x2f,
	0x89, 0x74, 0xa1, 0x67, 0xd8, 0xf0, 0xef, 0x00, 0xd8, 0x97, 0x5d, 0x5e, 0xac, 0x2b, 0x61, 0xc6,
	0x8d, 0xd4, 0xea, 0x7f, 0xc2, 0x87, 0xf0, 0xa1, 0xf0, 0x47, 0xb8, 0xdc, 0x1f, 0xd5, 0xc2, 0xce,
	0xe4, 0x0d, 0xff, 0x07, 0xf8, 0x10, 0x6d, 0x8a, 0xaf, 0xf9, 0xaa, 0xb2, 0xd5, 0x9b, 0xc3, 0xbb,
	0xd0, 0x2a, 0xc7, 0xc9, 0x9f, 0x42, 0x2f, 0x2e, 0x8a, 0xc7, 0xfa, 0xfd, 0xae, 0x60, 0x30, 0x53,
	0x53, 0xa5, 0xe7, 0x2a, 0x8d, 0xb5, 0x0e, 0x53, 0xb3, 0x88, 0xa8, 0xb4, 0x97, 0x90, 0x31, 0x71,
	0xfb, 0x80, 0x27, 0xd0, 0x1e, 0xc7, 0x7a, 0x16, 0x61, 0xc3, 0xf6, 0xfb, 0x26, 0xd6, 0x4c, 0x70,
	0x96, 0x18, 0x6c, 0xfa, 0xff, 0x36, 0xe0, 0xc2, 0xaa, 0x95, 0x05, 0x89, 0xa2, 0x9a, 0xe3, 0xcf,
	0xf0, 0xe9, 0x9d, 0xe8, 0x3e, 0x61, 0xa1, 0xed, 0x0c, 0x3b, 0x71, 0xce, 0x14, 0xb7, 0xc3, 0x68,
	0xbf, 0xcd, 0x22, 0xb2, 0x78, 0xc3, 0x3e, 0x4c, 0x42, 0x4a, 0xd8, 0xe1, 0x2f, 0x77, 0x46, 0x73,
	0x78, 0x01, 0xe7, 0x9c, 0x45, 0x66, 0x16, 0xbf, 0x81, 0xad, 0x43, 0x5e, 0x69, 0xa2, 0xf6, 0x7b,
	0x5e, 0x09, 0x76, 0x0e, 0x79, 0xa5, 0x9b, 0x8e, 0x2d, 0x2f, 0x26, 0xae, 0xe3, 0x32, 0x58, 0x4b,
	0x4e, 0xd8, 0x7d, 0xa7, 0xe7, 0xa0, 0x93, 0xe1, 0x25, 0x60, 0x05, 0x09, 0xcd, 0x67, 0xb6, 0xfb,
	0x78, 0x6a, 0x37, 0x58, 0x85, 0x5a, 0x3b, 0xf6, 0xac, 0x6f, 0x2a, 0xc0, 0x39, 0xb2, 0x7f, 0x18,
	0x18, 0x68, 0xce, 0x6c, 0xdd, 0x78, 0x66, 0x51, 0x3e, 0xd1, 0x3a, 0xb1, 0x30, 0xd7, 0xca, 0x30,
	0x6e, 0xf0, 0xdc, 0x59, 0x9b, 0x49, 0x65, 0x6c, 0xf5, 0xe8, 0xff, 0xe5, 0xc1, 0x67, 0xdb, 0xcc,
	0x30, 0xdb, 0x6e, 0x97, 0x0f, 0x99, 0xc8, 0x56, 0xf9, 0x4b, 0xbe, 0x79, 0x3d, 0x58, 0x00, 0xdf,
	0xc1, 0xc7, 0xaa, 0xad, 0x21, 0x25, 0x09, 0x1b, 0x53, 0x2a, 0x28, 0x90, 0xb7, 0x32, 0xb6, 0x1b,
	0x01, 0xa1, 0xa7, 0x88, 0x44, 0x6a, 0x74, 0x9a, 0x48, 0x35, 0x2d, 0x7b, 0x3b, 0x62, 0x32, 0xa8,
	0x16, 0xe5, 0x3e, 0x2f, 0x6c, 0x5a, 0x3b, 0x24, 0xb6, 0x22, 0xb7, 0x0b, 0x9c, 0x04, 0xd9, 0xd1,
	0x6d, 0x97, 0x07, 0xa4, 0xb0, 0x73, 0x33, 0x80, 0x6e, 0xbe, 0xb9, 0x7e, 0xdc, 0x5e, 0x3f, 0xdd,
	0x4f, 0x9a, 0x91, 0xf7, 0x9b, 0xf7, 0xed, 0xbe, 0xe3, 0x16, 0xfd, 0xcf, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xb8, 0x78, 0x9d, 0x52, 0xf9, 0x05, 0x00, 0x00,
}
