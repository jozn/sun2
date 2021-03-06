// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_enum.proto

package x

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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
func (MediaTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

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
func (FollowingEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

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
func (UserTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

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
func (UserLevelEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type GeneralPrivacyEnum int32

const (
	GeneralPrivacyEnum_UNKNOWN_GENERAL_PRIVACY      GeneralPrivacyEnum = 0
	GeneralPrivacyEnum_ALL_PEOPLE_PRIVACY           GeneralPrivacyEnum = 1
	GeneralPrivacyEnum_NOBODY_PRIVACY               GeneralPrivacyEnum = 2
	GeneralPrivacyEnum_CONTACTS_ONLY_PRIVACY        GeneralPrivacyEnum = 3
	GeneralPrivacyEnum_FOLLOWED_ONLY_PRIVACY        GeneralPrivacyEnum = 4
	GeneralPrivacyEnum_CONTACTS_AND_FOLLOWD_PRIVACY GeneralPrivacyEnum = 5
)

var GeneralPrivacyEnum_name = map[int32]string{
	0: "UNKNOWN_GENERAL_PRIVACY",
	1: "ALL_PEOPLE_PRIVACY",
	2: "NOBODY_PRIVACY",
	3: "CONTACTS_ONLY_PRIVACY",
	4: "FOLLOWED_ONLY_PRIVACY",
	5: "CONTACTS_AND_FOLLOWD_PRIVACY",
}
var GeneralPrivacyEnum_value = map[string]int32{
	"UNKNOWN_GENERAL_PRIVACY":      0,
	"ALL_PEOPLE_PRIVACY":           1,
	"NOBODY_PRIVACY":               2,
	"CONTACTS_ONLY_PRIVACY":        3,
	"FOLLOWED_ONLY_PRIVACY":        4,
	"CONTACTS_AND_FOLLOWD_PRIVACY": 5,
}

func (x GeneralPrivacyEnum) String() string {
	return proto.EnumName(GeneralPrivacyEnum_name, int32(x))
}
func (GeneralPrivacyEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

// DEP
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
func (UserOnlinePrivacyEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

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
func (UserOnlineStatusEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type PostTypeEnum int32

const (
	PostTypeEnum_POST_Type_Unknown PostTypeEnum = 0
	PostTypeEnum_POST_TEXT         PostTypeEnum = 1
	PostTypeEnum_POST_PHOTO        PostTypeEnum = 2
	PostTypeEnum_POST_VIDEO        PostTypeEnum = 3
	PostTypeEnum_POST_GIF          PostTypeEnum = 4
	PostTypeEnum_POST_AUDIO        PostTypeEnum = 5
	PostTypeEnum_POST_FILE         PostTypeEnum = 7
	PostTypeEnum_POST_POLL         PostTypeEnum = 8
	PostTypeEnum_POST_MEDIA        PostTypeEnum = 100
)

var PostTypeEnum_name = map[int32]string{
	0:   "POST_Type_Unknown",
	1:   "POST_TEXT",
	2:   "POST_PHOTO",
	3:   "POST_VIDEO",
	4:   "POST_GIF",
	5:   "POST_AUDIO",
	7:   "POST_FILE",
	8:   "POST_POLL",
	100: "POST_MEDIA",
}
var PostTypeEnum_value = map[string]int32{
	"POST_Type_Unknown": 0,
	"POST_TEXT":         1,
	"POST_PHOTO":        2,
	"POST_VIDEO":        3,
	"POST_GIF":          4,
	"POST_AUDIO":        5,
	"POST_FILE":         7,
	"POST_POLL":         8,
	"POST_MEDIA":        100,
}

func (x PostTypeEnum) String() string {
	return proto.EnumName(PostTypeEnum_name, int32(x))
}
func (PostTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type CategoryEnum int32

const (
	CategoryEnum_Category_RESHARED CategoryEnum = 0
)

var CategoryEnum_name = map[int32]string{
	0: "Category_RESHARED",
}
var CategoryEnum_value = map[string]int32{
	"Category_RESHARED": 0,
}

func (x CategoryEnum) String() string {
	return proto.EnumName(CategoryEnum_name, int32(x))
}
func (CategoryEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

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
func (PostCategoryEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

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
func (NotifyEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

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
func (ActionEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

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
func (ProfilePrivacyLevelEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

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
func (RoomTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

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
func (RoomActionDoingEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

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
	return fileDescriptor1, []int{15}
}

func init() {
	proto.RegisterEnum("MediaTypeEnum", MediaTypeEnum_name, MediaTypeEnum_value)
	proto.RegisterEnum("FollowingEnum", FollowingEnum_name, FollowingEnum_value)
	proto.RegisterEnum("UserTypeEnum", UserTypeEnum_name, UserTypeEnum_value)
	proto.RegisterEnum("UserLevelEnum", UserLevelEnum_name, UserLevelEnum_value)
	proto.RegisterEnum("GeneralPrivacyEnum", GeneralPrivacyEnum_name, GeneralPrivacyEnum_value)
	proto.RegisterEnum("UserOnlinePrivacyEnum", UserOnlinePrivacyEnum_name, UserOnlinePrivacyEnum_value)
	proto.RegisterEnum("UserOnlineStatusEnum", UserOnlineStatusEnum_name, UserOnlineStatusEnum_value)
	proto.RegisterEnum("PostTypeEnum", PostTypeEnum_name, PostTypeEnum_value)
	proto.RegisterEnum("CategoryEnum", CategoryEnum_name, CategoryEnum_value)
	proto.RegisterEnum("PostCategoryEnum", PostCategoryEnum_name, PostCategoryEnum_value)
	proto.RegisterEnum("NotifyEnum", NotifyEnum_name, NotifyEnum_value)
	proto.RegisterEnum("ActionEnum", ActionEnum_name, ActionEnum_value)
	proto.RegisterEnum("ProfilePrivacyLevelEnum", ProfilePrivacyLevelEnum_name, ProfilePrivacyLevelEnum_value)
	proto.RegisterEnum("RoomTypeEnum", RoomTypeEnum_name, RoomTypeEnum_value)
	proto.RegisterEnum("RoomActionDoingEnum", RoomActionDoingEnum_name, RoomActionDoingEnum_value)
	proto.RegisterEnum("RoomMessageDeliviryStatusEnum", RoomMessageDeliviryStatusEnum_name, RoomMessageDeliviryStatusEnum_value)
}

func init() { proto.RegisterFile("pb_enum.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 973 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x55, 0x5d, 0x6e, 0xe3, 0x36,
	0x10, 0xb6, 0xfc, 0x17, 0xef, 0xc4, 0x4e, 0x18, 0x25, 0x69, 0x76, 0xd1, 0x5d, 0xb4, 0x40, 0xd1,
	0x17, 0x3d, 0x04, 0x28, 0xda, 0x0b, 0x30, 0xe4, 0xd8, 0x26, 0x42, 0x91, 0xaa, 0x24, 0xc7, 0x71,
	0x5f, 0x08, 0x67, 0x57, 0x1b, 0x18, 0x75, 0xac, 0xc0, 0x76, 0xb2, 0x1b, 0xa0, 0x37, 0xe8, 0x0d,
	0xfa, 0xd4, 0x4b, 0xf4, 0x12, 0x3d, 0x55, 0x31, 0x52, 0xa4, 0x38, 0xe8, 0x9b, 0xf8, 0xf1, 0x9b,
	0xe1, 0xfc, 0x7c, 0x33, 0x82, 0xc1, 0xfd, 0x8d, 0xcb, 0x56, 0x0f, 0x77, 0xe7, 0xf7, 0xeb, 0x7c,
	0x9b, 0x07, 0x3f, 0xc1, 0x20, 0xcc, 0x3e, 0x2d, 0xe6, 0xe9, 0xd3, 0x7d, 0x86, 0xab, 0x87, 0x3b,
	0xff, 0x10, 0xf6, 0x43, 0x94, 0x8a, 0x3b, 0x15, 0xf2, 0x11, 0xb2, 0xc6, 0x0b, 0x70, 0xa5, 0x24,
	0x5a, 0xe6, 0x05, 0x06, 0x06, 0xc3, 0x7c, 0xb9, 0xcc, 0xbf, 0x2c, 0x56, 0xb7, 0x85, 0x89, 0x0f,
	0x07, 0x43, 0xab, 0xb5, 0x9d, 0x2a, 0x33, 0x72, 0xc6, 0x1a, 0xb2, 0x1a, 0xc0, 0x9b, 0x1a, 0x63,
	0x1e, 0x1d, 0x63, 0xfc, 0x75, 0x82, 0x49, 0x8a, 0x92, 0x35, 0xfd, 0x7d, 0xd8, 0xbb, 0xd0, 0x56,
	0x5c, 0xa2, 0x64, 0xad, 0xe0, 0x47, 0xe8, 0x4f, 0x36, 0xd9, 0xba, 0x8e, 0xa0, 0x07, 0xed, 0x49,
	0x82, 0x31, 0x6b, 0x10, 0x4d, 0x8c, 0xb9, 0x31, 0xa8, 0x99, 0x17, 0xfc, 0x01, 0x03, 0xa2, 0xe9,
	0xec, 0x31, 0x5b, 0x16, 0x3c, 0x06, 0x7d, 0x8d, 0x57, 0xa8, 0x9d, 0xb1, 0x71, 0xc8, 0x75, 0xf9,
	0x28, 0x8f, 0x22, 0xc7, 0x65, 0xa8, 0x4c, 0xf9, 0x68, 0x32, 0x49, 0x22, 0x6b, 0x64, 0xf1, 0xe8,
	0x29, 0x1c, 0x49, 0xd4, 0x98, 0xa2, 0x74, 0x17, 0x33, 0x67, 0xa7, 0x68, 0x30, 0x66, 0x2d, 0x72,
	0x53, 0xc1, 0x2a, 0xe6, 0x86, 0xb5, 0x29, 0x9f, 0xda, 0xae, 0xc4, 0x3a, 0xc1, 0x3f, 0x1e, 0xf8,
	0xa3, 0x6c, 0x95, 0xad, 0xe7, 0xcb, 0x68, 0xbd, 0x78, 0x9c, 0x7f, 0x7c, 0x2a, 0x62, 0xf8, 0x16,
	0xce, 0x26, 0xe6, 0xd2, 0xd8, 0xa9, 0x71, 0x23, 0xf2, 0xc7, 0xb5, 0x8b, 0x62, 0x75, 0xc5, 0xc5,
	0x8c, 0x35, 0xfc, 0x6f, 0xc0, 0xe7, 0x5a, 0xbb, 0x08, 0x6d, 0xa4, 0xb1, 0xc6, 0x3d, 0xf2, 0x6f,
	0xec, 0x85, 0x95, 0xb3, 0x1a, 0x6b, 0xfa, 0xef, 0xe0, 0x54, 0x58, 0x93, 0x72, 0x91, 0x26, 0xce,
	0x1a, 0xfd, 0x72, 0xd5, 0xa2, 0xab, 0xb2, 0x94, 0x28, 0x5f, 0x5f, 0xb5, 0xfd, 0xef, 0xe1, 0x7d,
	0x6d, 0xc5, 0x8d, 0x74, 0x25, 0x4f, 0xd6, 0x8c, 0x4e, 0xf0, 0x1d, 0x9c, 0x52, 0xd5, 0xec, 0x6a,
	0xb9, 0x58, 0x65, 0xbb, 0x91, 0x77, 0xa1, 0x89, 0xbf, 0xb0, 0x46, 0xf0, 0x97, 0x07, 0x27, 0x2f,
	0x8c, 0x64, 0x3b, 0xdf, 0x3e, 0x6c, 0x0a, 0xc2, 0x3e, 0xec, 0xe1, 0x35, 0x17, 0xa9, 0xa6, 0x54,
	0x00, 0xba, 0xd6, 0x68, 0x65, 0xb0, 0x2c, 0xab, 0xb0, 0xc6, 0xa0, 0x28, 0x7b, 0xc9, 0xa0, 0x3f,
	0xc4, 0xa9, 0x93, 0x7c, 0x96, 0x38, 0x3e, 0xb2, 0xac, 0xe5, 0xf7, 0xa1, 0x17, 0xa3, 0x40, 0x43,
	0xa6, 0x6d, 0xa2, 0x6b, 0x9e, 0xa4, 0x6e, 0x8a, 0x78, 0xc9, 0x3a, 0xfe, 0x01, 0x40, 0x71, 0x0c,
	0xad, 0x49, 0xc7, 0xac, 0xeb, 0x1f, 0xc1, 0x40, 0x5b, 0x33, 0x72, 0xa9, 0x0a, 0xb1, 0xb0, 0xdf,
	0x23, 0x01, 0x8c, 0x95, 0x44, 0xd6, 0x0b, 0xfe, 0xf6, 0xa0, 0x1f, 0xe5, 0x9b, 0x6d, 0xad, 0x8d,
	0x53, 0x38, 0x8a, 0x6c, 0x92, 0x3a, 0x02, 0xdc, 0x64, 0xf5, 0xfb, 0x2a, 0xff, 0xb2, 0x2a, 0x1b,
	0x5f, 0xc2, 0x78, 0x9d, 0x32, 0x8f, 0xde, 0x28, 0x8e, 0xd1, 0xd8, 0xa6, 0x96, 0x35, 0xeb, 0x73,
	0xa9, 0xe0, 0x22, 0xc0, 0xe2, 0x3c, 0x52, 0x43, 0xd6, 0xae, 0x6f, 0xf9, 0x44, 0x2a, 0xcb, 0x3a,
	0xb5, 0xb3, 0xa1, 0xd2, 0xc8, 0xf6, 0xea, 0x63, 0x64, 0xb5, 0x66, 0xbd, 0x9a, 0x5d, 0xcc, 0x04,
	0xfb, 0x44, 0xea, 0x15, 0xf3, 0x6d, 0x76, 0x9b, 0xaf, 0x9f, 0xaa, 0x08, 0xab, 0xb3, 0x8b, 0x31,
	0x19, 0xf3, 0x18, 0x25, 0x6b, 0x04, 0x0a, 0x18, 0x25, 0xf2, 0x8a, 0xca, 0xca, 0xe4, 0xc4, 0x7c,
	0xeb, 0xd2, 0xec, 0xeb, 0x96, 0x35, 0xa8, 0x18, 0x15, 0x52, 0x4c, 0x25, 0xf3, 0x76, 0x49, 0xc3,
	0xc5, 0x32, 0x63, 0xcd, 0x60, 0x0a, 0x60, 0xf2, 0xed, 0xe2, 0x73, 0xfd, 0x9e, 0xb1, 0xa9, 0x1a,
	0xce, 0x5c, 0x11, 0x96, 0x56, 0x34, 0x54, 0x0d, 0x12, 0xcd, 0x2e, 0x2c, 0x6c, 0x18, 0xa2, 0xa1,
	0x86, 0x79, 0xfe, 0x19, 0x1c, 0x3f, 0x5f, 0xd5, 0xb2, 0x9a, 0xd9, 0x09, 0x6b, 0x06, 0xd7, 0x00,
	0xfc, 0xe3, 0x76, 0x91, 0xaf, 0x2a, 0xc7, 0x5c, 0xa4, 0xca, 0x9a, 0xff, 0x39, 0xde, 0x85, 0x77,
	0x1d, 0xbf, 0x85, 0x93, 0xe7, 0xab, 0xda, 0x71, 0x31, 0xc8, 0xcd, 0xe0, 0x07, 0x38, 0x8b, 0xd6,
	0xf9, 0xe7, 0xc5, 0xb2, 0x92, 0xe0, 0xcb, 0x14, 0xf7, 0xa0, 0x5d, 0xae, 0x8c, 0xe0, 0x12, 0xfa,
	0x71, 0x9e, 0xdf, 0xed, 0xf6, 0xba, 0x9a, 0xad, 0xd8, 0xda, 0xd0, 0xa5, 0xb3, 0x08, 0x4b, 0x29,
	0x4a, 0x15, 0xa3, 0xa0, 0x46, 0xbf, 0x81, 0xce, 0x28, 0xb6, 0x93, 0x88, 0x35, 0xa9, 0x4d, 0x17,
	0xb1, 0xe5, 0x52, 0xf0, 0x24, 0x65, 0xad, 0xe0, 0xdf, 0x26, 0x1c, 0x93, 0xb7, 0x32, 0x21, 0x99,
	0x57, 0xbb, 0xea, 0x03, 0xbc, 0x7b, 0xe5, 0xf4, 0x39, 0x60, 0x69, 0x69, 0x4f, 0x15, 0xce, 0x05,
	0x37, 0x82, 0x16, 0x0e, 0x7d, 0xa7, 0xb3, 0x88, 0xf0, 0x26, 0x35, 0x26, 0x41, 0x23, 0x69, 0xc1,
	0x95, 0x7b, 0xb1, 0xe5, 0x1f, 0xc3, 0xa1, 0xe0, 0x51, 0x3a, 0x89, 0x5f, 0xc0, 0xf6, 0x2e, 0xaf,
	0x14, 0x5b, 0xe7, 0x35, 0xaf, 0x04, 0xbb, 0xbb, 0xbc, 0x52, 0x76, 0x7b, 0xc4, 0x8b, 0x51, 0xd8,
	0xb8, 0x34, 0xb6, 0x4a, 0x20, 0xeb, 0xbd, 0xf2, 0x57, 0x40, 0x6f, 0xfc, 0x13, 0x60, 0x15, 0x24,
	0xad, 0x98, 0x50, 0xf5, 0xd9, 0x3e, 0x6d, 0xe9, 0x0a, 0x25, 0x55, 0xf7, 0x49, 0x37, 0x15, 0x50,
	0x08, 0x79, 0xb0, 0x6b, 0xa8, 0xad, 0xe0, 0x94, 0x37, 0x3b, 0x20, 0x54, 0x8c, 0xad, 0x4d, 0x08,
	0x7e, 0xde, 0x25, 0xec, 0xb0, 0x98, 0x10, 0xae, 0x4c, 0x4a, 0xd9, 0xb3, 0xe0, 0x4f, 0x0f, 0x3e,
	0x50, 0x31, 0xc3, 0x6c, 0xb3, 0x99, 0xdf, 0x66, 0x32, 0x5b, 0x2e, 0x1e, 0x17, 0xeb, 0xa7, 0x9d,
	0x65, 0xf1, 0x1e, 0xde, 0x56, 0x65, 0x0d, 0x31, 0x49, 0xf8, 0x08, 0x9d, 0x44, 0xad, 0xae, 0x54,
	0x4c, 0xdb, 0x83, 0x41, 0xdf, 0x20, 0x4a, 0x97, 0x5a, 0x97, 0x28, 0x73, 0x59, 0xd6, 0x76, 0xc8,
	0x95, 0xae, 0x7e, 0x06, 0xcf, 0x71, 0xb1, 0x16, 0xc9, 0x21, 0xa1, 0x8c, 0x8a, 0xbd, 0x51, 0xb8,
	0x40, 0x1a, 0xa0, 0x4e, 0x79, 0x81, 0x86, 0x75, 0x2f, 0x8e, 0xa0, 0xb7, 0x58, 0x9f, 0xdf, 0x6d,
	0xce, 0xef, 0x6f, 0xc6, 0xad, 0xc8, 0xfb, 0xcd, 0xfb, 0x7a, 0xd3, 0x2d, 0x7e, 0x66, 0x3f, 0xff,
	0x17, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x95, 0x0e, 0x43, 0xdd, 0x06, 0x00, 0x00,
}
