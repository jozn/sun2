package xconst

const (
	Action_Table                       = "action"
	Action_TableGo                     = "Action"
	Comment_Table                      = "comment"
	Comment_TableGo                    = "Comment"
	CommentDeleted_Table               = "comment_deleted"
	CommentDeleted_TableGo             = "CommentDeleted"
	Event_Table                        = "event"
	Event_TableGo                      = "Event"
	FollowingList_Table                = "following_list"
	FollowingList_TableGo              = "FollowingList"
	FollowingListMember_Table          = "following_list_member"
	FollowingListMember_TableGo        = "FollowingListMember"
	FollowingListMemberRemoved_Table   = "following_list_member_removed"
	FollowingListMemberRemoved_TableGo = "FollowingListMemberRemoved"
	Group_Table                        = "group"
	Group_TableGo                      = "Group"
	GroupMember_Table                  = "group_member"
	GroupMember_TableGo                = "GroupMember"
	GroupMessage_Table                 = "group_message"
	GroupMessage_TableGo               = "GroupMessage"
	Like_Table                         = "likes"
	Like_TableGo                       = "Like"
	Notify_Table                       = "notify"
	Notify_TableGo                     = "Notify"
	NotifyRemoved_Table                = "notify_removed"
	NotifyRemoved_TableGo              = "NotifyRemoved"
	PhoneContact_Table                 = "phone_contacts"
	PhoneContact_TableGo               = "PhoneContact"
	Post_Table                         = "post"
	Post_TableGo                       = "Post"
	PostCount_Table                    = "post_count"
	PostCount_TableGo                  = "PostCount"
	PostDeleted_Table                  = "post_deleted"
	PostDeleted_TableGo                = "PostDeleted"
	PostKey_Table                      = "post_keys"
	PostKey_TableGo                    = "PostKey"
	PostLink_Table                     = "post_link"
	PostLink_TableGo                   = "PostLink"
	PostMedia_Table                    = "post_media"
	PostMedia_TableGo                  = "PostMedia"
	PostMentioned_Table                = "post_mentioned"
	PostMentioned_TableGo              = "PostMentioned"
	PostReshared_Table                 = "post_reshared"
	PostReshared_TableGo               = "PostReshared"
	SearchClicked_Table                = "search_clicked"
	SearchClicked_TableGo              = "SearchClicked"
	Session_Table                      = "session"
	Session_TableGo                    = "Session"
	SettingClient_Table                = "setting_client"
	SettingClient_TableGo              = "SettingClient"
	SettingNotification_Table          = "setting_notifications"
	SettingNotification_TableGo        = "SettingNotification"
	Tag_Table                          = "tag"
	Tag_TableGo                        = "Tag"
	TagPost_Table                      = "tag_post"
	TagPost_TableGo                    = "TagPost"
	TriggerLog_Table                   = "trigger_log"
	TriggerLog_TableGo                 = "TriggerLog"
	User_Table                         = "user"
	User_TableGo                       = "User"
	UserMetaInfo_Table                 = "user_meta_info"
	UserMetaInfo_TableGo               = "UserMetaInfo"
	UserPassword_Table                 = "user_password"
	UserPassword_TableGo               = "UserPassword"
	Chat_Table                         = "chat"
	Chat_TableGo                       = "Chat"
	ChatLastMessage_Table              = "chat_last_message"
	ChatLastMessage_TableGo            = "ChatLastMessage"
	ChatSync_Table                     = "chat_sync"
	ChatSync_TableGo                   = "ChatSync"
	DirectMessage_Table                = "direct_message"
	DirectMessage_TableGo              = "DirectMessage"
	Home_Table                         = "home"
	Home_TableGo                       = "Home"
	MessageFile_Table                  = "message_file"
	MessageFile_TableGo                = "MessageFile"
	FileMsg_Table                      = "file_msg"
	FileMsg_TableGo                    = "FileMsg"
	FilePost_Table                     = "file_post"
	FilePost_TableGo                   = "FilePost"
	ActionFanout_Table                 = "action_fanout"
	ActionFanout_TableGo               = "ActionFanout"
	HomeFanout_Table                   = "home_fanout"
	HomeFanout_TableGo                 = "HomeFanout"
	SuggestedTopPost_Table             = "suggested_top_posts"
	SuggestedTopPost_TableGo           = "SuggestedTopPost"
	SuggestedUser_Table                = "suggested_user"
	SuggestedUser_TableGo              = "SuggestedUser"
)

var Action = struct {
	ActionId       string
	ActorUserId    string
	ActionTypeEnum string
	PeerUserId     string
	PostId         string
	CommentId      string
	Murmur64Hash   string
	CreatedTime    string
}{

	ActionId:       "ActionId",
	ActorUserId:    "ActorUserId",
	ActionTypeEnum: "ActionTypeEnum",
	PeerUserId:     "PeerUserId",
	PostId:         "PostId",
	CommentId:      "CommentId",
	Murmur64Hash:   "Murmur64Hash",
	CreatedTime:    "CreatedTime",
}

var Comment = struct {
	CommentId   string
	UserId      string
	PostId      string
	Text        string
	LikesCount  string
	CreatedTime string
}{

	CommentId:   "CommentId",
	UserId:      "UserId",
	PostId:      "PostId",
	Text:        "Text",
	LikesCount:  "LikesCount",
	CreatedTime: "CreatedTime",
}

var CommentDeleted = struct {
	CommentId string
	UserId    string
}{

	CommentId: "CommentId",
	UserId:    "UserId",
}

var Event = struct {
	EventId      string
	EventType    string
	ByUserId     string
	PeerUserId   string
	PostId       string
	CommentId    string
	ActionId     string
	Murmur64Hash string
	ChatKey      string
	MessageId    string
	ReSharedId   string
}{

	EventId:      "EventId",
	EventType:    "EventType",
	ByUserId:     "ByUserId",
	PeerUserId:   "PeerUserId",
	PostId:       "PostId",
	CommentId:    "CommentId",
	ActionId:     "ActionId",
	Murmur64Hash: "Murmur64Hash",
	ChatKey:      "ChatKey",
	MessageId:    "MessageId",
	ReSharedId:   "ReSharedId",
}

var FollowingList = struct {
	Id          string
	UserId      string
	ListType    string
	Name        string
	Count       string
	IsAuto      string
	IsPimiry    string
	CreatedTime string
}{

	Id:          "Id",
	UserId:      "UserId",
	ListType:    "ListType",
	Name:        "Name",
	Count:       "Count",
	IsAuto:      "IsAuto",
	IsPimiry:    "IsPimiry",
	CreatedTime: "CreatedTime",
}

var FollowingListMember = struct {
	Id             string
	ListId         string
	UserId         string
	FollowedUserId string
	CreatedTime    string
}{

	Id:             "Id",
	ListId:         "ListId",
	UserId:         "UserId",
	FollowedUserId: "FollowedUserId",
	CreatedTime:    "CreatedTime",
}

var FollowingListMemberRemoved = struct {
	Id               string
	ListId           string
	UserId           string
	UnFollowedUserId string
	UpdatedTime      string
}{

	Id:               "Id",
	ListId:           "ListId",
	UserId:           "UserId",
	UnFollowedUserId: "UnFollowedUserId",
	UpdatedTime:      "UpdatedTime",
}

var Group = struct {
	GroupId          string
	GroupName        string
	MembersCount     string
	GroupPrivacyEnum string
	CreatorUserId    string
	CreatedTime      string
	UpdatedMs        string
	CurrentSeq       string
}{

	GroupId:          "GroupId",
	GroupName:        "GroupName",
	MembersCount:     "MembersCount",
	GroupPrivacyEnum: "GroupPrivacyEnum",
	CreatorUserId:    "CreatorUserId",
	CreatedTime:      "CreatedTime",
	UpdatedMs:        "UpdatedMs",
	CurrentSeq:       "CurrentSeq",
}

var GroupMember = struct {
	Id              string
	GroupId         string
	GroupKey        string
	UserId          string
	ByUserId        string
	GroupRoleEnumId string
	CreatedTime     string
}{

	Id:              "Id",
	GroupId:         "GroupId",
	GroupKey:        "GroupKey",
	UserId:          "UserId",
	ByUserId:        "ByUserId",
	GroupRoleEnumId: "GroupRoleEnumId",
	CreatedTime:     "CreatedTime",
}

var GroupMessage = struct {
	MessageId          string
	RoomKey            string
	UserId             string
	MessageFileId      string
	MessageTypeEnum    string
	Text               string
	CreatedMs          string
	DeliveryStatusEnum string
}{

	MessageId:          "MessageId",
	RoomKey:            "RoomKey",
	UserId:             "UserId",
	MessageFileId:      "MessageFileId",
	MessageTypeEnum:    "MessageTypeEnum",
	Text:               "Text",
	CreatedMs:          "CreatedMs",
	DeliveryStatusEnum: "DeliveryStatusEnum",
}

var Like = struct {
	Id           string
	PostId       string
	PostTypeEnum string
	UserId       string
	LikeEnum     string
	CreatedTime  string
}{

	Id:           "Id",
	PostId:       "PostId",
	PostTypeEnum: "PostTypeEnum",
	UserId:       "UserId",
	LikeEnum:     "LikeEnum",
	CreatedTime:  "CreatedTime",
}

var Notify = struct {
	NotifyId       string
	ForUserId      string
	ActorUserId    string
	NotifyTypeEnum string
	PostId         string
	CommentId      string
	PeerUserId     string
	Murmur64Hash   string
	SeenStatus     string
	CreatedTime    string
}{

	NotifyId:       "NotifyId",
	ForUserId:      "ForUserId",
	ActorUserId:    "ActorUserId",
	NotifyTypeEnum: "NotifyTypeEnum",
	PostId:         "PostId",
	CommentId:      "CommentId",
	PeerUserId:     "PeerUserId",
	Murmur64Hash:   "Murmur64Hash",
	SeenStatus:     "SeenStatus",
	CreatedTime:    "CreatedTime",
}

var NotifyRemoved = struct {
	Murmur64Hash string
	ForUserId    string
	Id           string
}{

	Murmur64Hash: "Murmur64Hash",
	ForUserId:    "ForUserId",
	Id:           "Id",
}

var PhoneContact = struct {
	Id                    string
	UserId                string
	Phone                 string
	PhoneDisplayName      string
	PhoneFamilyName       string
	PhoneNumber           string
	PhoneNormalizedNumber string
	PhoneContactRowId     string
	DeviceUuidId          string
	CreatedTime           string
}{

	Id:                    "Id",
	UserId:                "UserId",
	Phone:                 "Phone",
	PhoneDisplayName:      "PhoneDisplayName",
	PhoneFamilyName:       "PhoneFamilyName",
	PhoneNumber:           "PhoneNumber",
	PhoneNormalizedNumber: "PhoneNormalizedNumber",
	PhoneContactRowId:     "PhoneContactRowId",
	DeviceUuidId:          "DeviceUuidId",
	CreatedTime:           "CreatedTime",
}

var Post = struct {
	PostId           string
	UserId           string
	PostTypeEnum     string
	PostCategoryEnum string
	MediaId          string
	PostKey          string
	Text             string
	RichText         string
	MediaCount       string
	SharedTo         string
	DisableComment   string
	Source           string
	HasTag           string
	Seq              string
	CommentsCount    string
	LikesCount       string
	ViewsCount       string
	EditedTime       string
	CreatedTime      string
	ReSharedPostId   string
}{

	PostId:           "PostId",
	UserId:           "UserId",
	PostTypeEnum:     "PostTypeEnum",
	PostCategoryEnum: "PostCategoryEnum",
	MediaId:          "MediaId",
	PostKey:          "PostKey",
	Text:             "Text",
	RichText:         "RichText",
	MediaCount:       "MediaCount",
	SharedTo:         "SharedTo",
	DisableComment:   "DisableComment",
	Source:           "Source",
	HasTag:           "HasTag",
	Seq:              "Seq",
	CommentsCount:    "CommentsCount",
	LikesCount:       "LikesCount",
	ViewsCount:       "ViewsCount",
	EditedTime:       "EditedTime",
	CreatedTime:      "CreatedTime",
	ReSharedPostId:   "ReSharedPostId",
}

var PostCount = struct {
	PostId     string
	ViewsCount string
}{

	PostId:     "PostId",
	ViewsCount: "ViewsCount",
}

var PostDeleted = struct {
	PostId string
	UserId string
}{

	PostId: "PostId",
	UserId: "UserId",
}

var PostKey = struct {
	Id         string
	PostKeyStr string
	Used       string
}{

	Id:         "Id",
	PostKeyStr: "PostKeyStr",
	Used:       "Used",
}

var PostLink = struct {
	LinkId  string
	LinkUrl string
}{

	LinkId:  "LinkId",
	LinkUrl: "LinkUrl",
}

var PostMedia = struct {
	MediaId       string
	UserId        string
	PostId        string
	AlbumId       string
	MediaTypeEnum string
	Width         string
	Height        string
	Size          string
	Duration      string
	Extension     string
	Md5Hash       string
	Color         string
	CreatedTime   string
	ViewCount     string
	Extra         string
}{

	MediaId:       "MediaId",
	UserId:        "UserId",
	PostId:        "PostId",
	AlbumId:       "AlbumId",
	MediaTypeEnum: "MediaTypeEnum",
	Width:         "Width",
	Height:        "Height",
	Size:          "Size",
	Duration:      "Duration",
	Extension:     "Extension",
	Md5Hash:       "Md5Hash",
	Color:         "Color",
	CreatedTime:   "CreatedTime",
	ViewCount:     "ViewCount",
	Extra:         "Extra",
}

var PostMentioned = struct {
	MentionedId      string
	ForUserId        string
	PostId           string
	PostUserId       string
	PostTypeEnum     string
	PostCategoryEnum string
	CreatedTime      string
}{

	MentionedId:      "MentionedId",
	ForUserId:        "ForUserId",
	PostId:           "PostId",
	PostUserId:       "PostUserId",
	PostTypeEnum:     "PostTypeEnum",
	PostCategoryEnum: "PostCategoryEnum",
	CreatedTime:      "CreatedTime",
}

var PostReshared = struct {
	ResharedId       string
	ByUserId         string
	PostId           string
	PostUserId       string
	PostTypeEnum     string
	PostCategoryEnum string
	CreatedTime      string
}{

	ResharedId:       "ResharedId",
	ByUserId:         "ByUserId",
	PostId:           "PostId",
	PostUserId:       "PostUserId",
	PostTypeEnum:     "PostTypeEnum",
	PostCategoryEnum: "PostCategoryEnum",
	CreatedTime:      "CreatedTime",
}

var SearchClicked = struct {
	Id          string
	Query       string
	ClickType   string
	TargetId    string
	UserId      string
	CreatedTime string
}{

	Id:          "Id",
	Query:       "Query",
	ClickType:   "ClickType",
	TargetId:    "TargetId",
	UserId:      "UserId",
	CreatedTime: "CreatedTime",
}

var Session = struct {
	SessionUuid         string
	UserId              string
	ClientUuid          string
	DeviceUuid          string
	LastActivityTime    string
	LastIpAddress       string
	LastWifiMacAddress  string
	LastNetworkType     string
	LastNetworkTypeEnum string
	AppVersion          string
	UpdatedTime         string
	CreatedTime         string
}{

	SessionUuid:         "SessionUuid",
	UserId:              "UserId",
	ClientUuid:          "ClientUuid",
	DeviceUuid:          "DeviceUuid",
	LastActivityTime:    "LastActivityTime",
	LastIpAddress:       "LastIpAddress",
	LastWifiMacAddress:  "LastWifiMacAddress",
	LastNetworkType:     "LastNetworkType",
	LastNetworkTypeEnum: "LastNetworkTypeEnum",
	AppVersion:          "AppVersion",
	UpdatedTime:         "UpdatedTime",
	CreatedTime:         "CreatedTime",
}

var SettingClient = struct {
	UserId                    string
	AutoDownloadWifiVoice     string
	AutoDownloadWifiImage     string
	AutoDownloadWifiVideo     string
	AutoDownloadWifiFile      string
	AutoDownloadWifiMusic     string
	AutoDownloadWifiGif       string
	AutoDownloadCellularVoice string
	AutoDownloadCellularImage string
	AutoDownloadCellularVideo string
	AutoDownloadCellularFile  string
	AutoDownloadCellularMusic string
	AutoDownloadCellularGif   string
	AutoDownloadRoamingVoice  string
	AutoDownloadRoamingImage  string
	AutoDownloadRoamingVideo  string
	AutoDownloadRoamingFile   string
	AutoDownloadRoamingMusic  string
	AutoDownloadRoamingGif    string
	SaveToGallery             string
}{

	UserId:                    "UserId",
	AutoDownloadWifiVoice:     "AutoDownloadWifiVoice",
	AutoDownloadWifiImage:     "AutoDownloadWifiImage",
	AutoDownloadWifiVideo:     "AutoDownloadWifiVideo",
	AutoDownloadWifiFile:      "AutoDownloadWifiFile",
	AutoDownloadWifiMusic:     "AutoDownloadWifiMusic",
	AutoDownloadWifiGif:       "AutoDownloadWifiGif",
	AutoDownloadCellularVoice: "AutoDownloadCellularVoice",
	AutoDownloadCellularImage: "AutoDownloadCellularImage",
	AutoDownloadCellularVideo: "AutoDownloadCellularVideo",
	AutoDownloadCellularFile:  "AutoDownloadCellularFile",
	AutoDownloadCellularMusic: "AutoDownloadCellularMusic",
	AutoDownloadCellularGif:   "AutoDownloadCellularGif",
	AutoDownloadRoamingVoice:  "AutoDownloadRoamingVoice",
	AutoDownloadRoamingImage:  "AutoDownloadRoamingImage",
	AutoDownloadRoamingVideo:  "AutoDownloadRoamingVideo",
	AutoDownloadRoamingFile:   "AutoDownloadRoamingFile",
	AutoDownloadRoamingMusic:  "AutoDownloadRoamingMusic",
	AutoDownloadRoamingGif:    "AutoDownloadRoamingGif",
	SaveToGallery:             "SaveToGallery",
}

var SettingNotification = struct {
	UserId                   string
	SocialLedOn              string
	SocialLedColor           string
	ReqestToFollowYou        string
	FollowedYou              string
	AccptedYourFollowRequest string
	YourPostLiked            string
	YourPostCommented        string
	MenthenedYouInPost       string
	MenthenedYouInComment    string
	YourContactsJoined       string
	DirectMessage            string
	DirectAlert              string
	DirectPerview            string
	DirectLedOn              string
	DirectLedColor           string
	DirectVibrate            string
	DirectPopup              string
	DirectSound              string
	DirectPriority           string
}{

	UserId:                   "UserId",
	SocialLedOn:              "SocialLedOn",
	SocialLedColor:           "SocialLedColor",
	ReqestToFollowYou:        "ReqestToFollowYou",
	FollowedYou:              "FollowedYou",
	AccptedYourFollowRequest: "AccptedYourFollowRequest",
	YourPostLiked:            "YourPostLiked",
	YourPostCommented:        "YourPostCommented",
	MenthenedYouInPost:       "MenthenedYouInPost",
	MenthenedYouInComment:    "MenthenedYouInComment",
	YourContactsJoined:       "YourContactsJoined",
	DirectMessage:            "DirectMessage",
	DirectAlert:              "DirectAlert",
	DirectPerview:            "DirectPerview",
	DirectLedOn:              "DirectLedOn",
	DirectLedColor:           "DirectLedColor",
	DirectVibrate:            "DirectVibrate",
	DirectPopup:              "DirectPopup",
	DirectSound:              "DirectSound",
	DirectPriority:           "DirectPriority",
}

var Tag = struct {
	TagId         string
	Name          string
	Count         string
	TagStatusEnum string
	CreatedTime   string
}{

	TagId:         "TagId",
	Name:          "Name",
	Count:         "Count",
	TagStatusEnum: "TagStatusEnum",
	CreatedTime:   "CreatedTime",
}

var TagPost = struct {
	Id               string
	TagId            string
	PostId           string
	PostTypeEnum     string
	PostCategoryEnum string
	CreatedTime      string
}{

	Id:               "Id",
	TagId:            "TagId",
	PostId:           "PostId",
	PostTypeEnum:     "PostTypeEnum",
	PostCategoryEnum: "PostCategoryEnum",
	CreatedTime:      "CreatedTime",
}

var TriggerLog = struct {
	Id         string
	ModelName  string
	ChangeType string
	TargetId   string
	TargetStr  string
	CreatedSe  string
}{

	Id:         "Id",
	ModelName:  "ModelName",
	ChangeType: "ChangeType",
	TargetId:   "TargetId",
	TargetStr:  "TargetStr",
	CreatedSe:  "CreatedSe",
}

var User = struct {
	UserId               string
	UserName             string
	UserNameLower        string
	FirstName            string
	LastName             string
	UserTypeEnum         string
	UserLevelEnum        string
	AvatarId             string
	ProfilePrivacyEnum   string
	Phone                string
	About                string
	Email                string
	PasswordHash         string
	PasswordSalt         string
	PostSeq              string
	FollowersCount       string
	FollowingCount       string
	PostsCount           string
	MediaCount           string
	PhotoCount           string
	VideoCount           string
	GifCount             string
	AudioCount           string
	VoiceCount           string
	FileCount            string
	LinkCount            string
	BoardCount           string
	PinedCount           string
	LikesCount           string
	ResharedCount        string
	LastActionTime       string
	LastPostTime         string
	PrimaryFollowingList string
	CreatedSe            string
	UpdatedMs            string
	OnlinePrivacyEnum    string
	LastActivityTime     string
	Phone2               string
}{

	UserId:               "UserId",
	UserName:             "UserName",
	UserNameLower:        "UserNameLower",
	FirstName:            "FirstName",
	LastName:             "LastName",
	UserTypeEnum:         "UserTypeEnum",
	UserLevelEnum:        "UserLevelEnum",
	AvatarId:             "AvatarId",
	ProfilePrivacyEnum:   "ProfilePrivacyEnum",
	Phone:                "Phone",
	About:                "About",
	Email:                "Email",
	PasswordHash:         "PasswordHash",
	PasswordSalt:         "PasswordSalt",
	PostSeq:              "PostSeq",
	FollowersCount:       "FollowersCount",
	FollowingCount:       "FollowingCount",
	PostsCount:           "PostsCount",
	MediaCount:           "MediaCount",
	PhotoCount:           "PhotoCount",
	VideoCount:           "VideoCount",
	GifCount:             "GifCount",
	AudioCount:           "AudioCount",
	VoiceCount:           "VoiceCount",
	FileCount:            "FileCount",
	LinkCount:            "LinkCount",
	BoardCount:           "BoardCount",
	PinedCount:           "PinedCount",
	LikesCount:           "LikesCount",
	ResharedCount:        "ResharedCount",
	LastActionTime:       "LastActionTime",
	LastPostTime:         "LastPostTime",
	PrimaryFollowingList: "PrimaryFollowingList",
	CreatedSe:            "CreatedSe",
	UpdatedMs:            "UpdatedMs",
	OnlinePrivacyEnum:    "OnlinePrivacyEnum",
	LastActivityTime:     "LastActivityTime",
	Phone2:               "Phone2",
}

var UserMetaInfo = struct {
	Id                  string
	UserId              string
	IsNotificationDirty string
	LastUserRecGen      string
}{

	Id:                  "Id",
	UserId:              "UserId",
	IsNotificationDirty: "IsNotificationDirty",
	LastUserRecGen:      "LastUserRecGen",
}

var UserPassword = struct {
	UserId      string
	Password    string
	CreatedTime string
}{

	UserId:      "UserId",
	Password:    "Password",
	CreatedTime: "CreatedTime",
}

var Chat = struct {
	ChatKey      string
	RoomKey      string
	RoomTypeEnum string
	UserId       string
	PeerUserId   string
	GroupId      string
	CreatedTime  string
	Seq          string
	SeenSeq      string
	UpdatedMs    string
}{

	ChatKey:      "ChatKey",
	RoomKey:      "RoomKey",
	RoomTypeEnum: "RoomTypeEnum",
	UserId:       "UserId",
	PeerUserId:   "PeerUserId",
	GroupId:      "GroupId",
	CreatedTime:  "CreatedTime",
	Seq:          "Seq",
	SeenSeq:      "SeenSeq",
	UpdatedMs:    "UpdatedMs",
}

var ChatLastMessage = struct {
	ChatKey     string
	ForUserId   string
	LastMsgPb   string
	LastMsgJson string
}{

	ChatKey:     "ChatKey",
	ForUserId:   "ForUserId",
	LastMsgPb:   "LastMsgPb",
	LastMsgJson: "LastMsgJson",
}

var ChatSync = struct {
	SyncId            string
	ToUserId          string
	ChatSyncTypeId    string
	RoomKey           string
	ChatKey           string
	FromHighMessageId string
	ToLowMessageId    string
	MessageId         string
	MessagePb         string
	MessageJson       string
	CreatedTime       string
}{

	SyncId:            "SyncId",
	ToUserId:          "ToUserId",
	ChatSyncTypeId:    "ChatSyncTypeId",
	RoomKey:           "RoomKey",
	ChatKey:           "ChatKey",
	FromHighMessageId: "FromHighMessageId",
	ToLowMessageId:    "ToLowMessageId",
	MessageId:         "MessageId",
	MessagePb:         "MessagePb",
	MessageJson:       "MessageJson",
	CreatedTime:       "CreatedTime",
}

var DirectMessage = struct {
	ChatKey            string
	MessageId          string
	RoomKey            string
	UserId             string
	MessageFileId      string
	MessageTypeEnum    string
	Text               string
	CreatedTime        string
	Seq                string
	DeliviryStatusEnum string
	ExtraPB            string
}{

	ChatKey:            "ChatKey",
	MessageId:          "MessageId",
	RoomKey:            "RoomKey",
	UserId:             "UserId",
	MessageFileId:      "MessageFileId",
	MessageTypeEnum:    "MessageTypeEnum",
	Text:               "Text",
	CreatedTime:        "CreatedTime",
	Seq:                "Seq",
	DeliviryStatusEnum: "DeliviryStatusEnum",
	ExtraPB:            "ExtraPB",
}

var Home = struct {
	Id        string
	ForUserId string
	PostId    string
}{

	Id:        "Id",
	ForUserId: "ForUserId",
	PostId:    "PostId",
}

var MessageFile = struct {
	MessageFileId string
	FileTypeEnum  string
	UserId        string
	Title         string
	Size          string
	Width         string
	Height        string
	Duration      string
	Extension     string
	Md5Hash       string
	CreatedTime   string
}{

	MessageFileId: "MessageFileId",
	FileTypeEnum:  "FileTypeEnum",
	UserId:        "UserId",
	Title:         "Title",
	Size:          "Size",
	Width:         "Width",
	Height:        "Height",
	Duration:      "Duration",
	Extension:     "Extension",
	Md5Hash:       "Md5Hash",
	CreatedTime:   "CreatedTime",
}

var FileMsg = struct {
	Id        string
	FileType  string
	Extension string
	DataThumb string
	Data      string
}{

	Id:        "Id",
	FileType:  "FileType",
	Extension: "Extension",
	DataThumb: "DataThumb",
	Data:      "Data",
}

var FilePost = struct {
	Id        string
	FileType  string
	Extension string
	DataThumb string
	Data      string
}{

	Id:        "Id",
	FileType:  "FileType",
	Extension: "Extension",
	DataThumb: "DataThumb",
	Data:      "Data",
}

var ActionFanout = struct {
	OrderId     string
	ForUserId   string
	ActionId    string
	ActorUserId string
}{

	OrderId:     "OrderId",
	ForUserId:   "ForUserId",
	ActionId:    "ActionId",
	ActorUserId: "ActorUserId",
}

var HomeFanout = struct {
	OrderId    string
	ForUserId  string
	PostId     string
	PostUserId string
}{

	OrderId:    "OrderId",
	ForUserId:  "ForUserId",
	PostId:     "PostId",
	PostUserId: "PostUserId",
}

var SuggestedTopPost = struct {
	Id     string
	PostId string
}{

	Id:     "Id",
	PostId: "PostId",
}

var SuggestedUser = struct {
	Id          string
	UserId      string
	TargetId    string
	Weight      string
	CreatedTime string
}{

	Id:          "Id",
	UserId:      "UserId",
	TargetId:    "TargetId",
	Weight:      "Weight",
	CreatedTime: "CreatedTime",
}
