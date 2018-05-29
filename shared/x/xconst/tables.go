package xconst

const (
	Action_Table                   = "action"
	Action_TableGo                 = "Action"
	Comment_Table                  = "comment"
	Comment_TableGo                = "Comment"
	CommentDeleted_Table           = "comment_deleted"
	CommentDeleted_TableGo         = "CommentDeleted"
	Event_Table                    = "event"
	Event_TableGo                  = "Event"
	Like_Table                     = "likes"
	Like_TableGo                   = "Like"
	Notify_Table                   = "notify"
	Notify_TableGo                 = "Notify"
	NotifyRemoved_Table            = "notify_removed"
	NotifyRemoved_TableGo          = "NotifyRemoved"
	PhoneContact_Table             = "phone_contacts"
	PhoneContact_TableGo           = "PhoneContact"
	Post_Table                     = "post"
	Post_TableGo                   = "Post"
	PostCount_Table                = "post_count"
	PostCount_TableGo              = "PostCount"
	PostDeleted_Table              = "post_deleted"
	PostDeleted_TableGo            = "PostDeleted"
	PostKey_Table                  = "post_keys"
	PostKey_TableGo                = "PostKey"
	PostLink_Table                 = "post_link"
	PostLink_TableGo               = "PostLink"
	PostMedia_Table                = "post_media"
	PostMedia_TableGo              = "PostMedia"
	PostMentioned_Table            = "post_mentioned"
	PostMentioned_TableGo          = "PostMentioned"
	PostReshared_Table             = "post_reshared"
	PostReshared_TableGo           = "PostReshared"
	Session_Table                  = "session"
	Session_TableGo                = "Session"
	SettingClient_Table            = "setting_client"
	SettingClient_TableGo          = "SettingClient"
	SettingNotification_Table      = "setting_notifications"
	SettingNotification_TableGo    = "SettingNotification"
	Tag_Table                      = "tag"
	Tag_TableGo                    = "Tag"
	TagPost_Table                  = "tag_post"
	TagPost_TableGo                = "TagPost"
	TriggerLog_Table               = "trigger_log"
	TriggerLog_TableGo             = "TriggerLog"
	User_Table                     = "user"
	User_TableGo                   = "User"
	UserRelation_Table             = "user_relation"
	UserRelation_TableGo           = "UserRelation"
	Chat_Table                     = "chat"
	Chat_TableGo                   = "Chat"
	ChatLastMessage_Table          = "chat_last_message"
	ChatLastMessage_TableGo        = "ChatLastMessage"
	ChatVersionOrder_Table         = "chat_version_order"
	ChatVersionOrder_TableGo       = "ChatVersionOrder"
	Group_Table                    = "group"
	Group_TableGo                  = "Group"
	GroupMember_Table              = "group_member"
	GroupMember_TableGo            = "GroupMember"
	GroupOrderdUser_Table          = "group_orderd_user"
	GroupOrderdUser_TableGo        = "GroupOrderdUser"
	GroupPinedMsg_Table            = "group_pined_msg"
	GroupPinedMsg_TableGo          = "GroupPinedMsg"
	FileMsg_Table                  = "file_msg"
	FileMsg_TableGo                = "FileMsg"
	FilePost_Table                 = "file_post"
	FilePost_TableGo               = "FilePost"
	ActionFanout_Table             = "action_fanout"
	ActionFanout_TableGo           = "ActionFanout"
	HomeFanout_Table               = "home_fanout"
	HomeFanout_TableGo             = "HomeFanout"
	SuggestedTopPost_Table         = "suggested_top_posts"
	SuggestedTopPost_TableGo       = "SuggestedTopPost"
	SuggestedUser_Table            = "suggested_user"
	SuggestedUser_TableGo          = "SuggestedUser"
	PushChat_Table                 = "push_chat"
	PushChat_TableGo               = "PushChat"
	HTTPRPCLog_Table               = "http_rpc_log"
	HTTPRPCLog_TableGo             = "HTTPRPCLog"
	MetricLog_Table                = "metric_log"
	MetricLog_TableGo              = "MetricLog"
	XfileServiceInfoLog_Table      = "xfile_service_info_log"
	XfileServiceInfoLog_TableGo    = "XfileServiceInfoLog"
	XfileServiceMetricLog_Table    = "xfile_service_metric_log"
	XfileServiceMetricLog_TableGo  = "XfileServiceMetricLog"
	XfileServiceRequestLog_Table   = "xfile_service_request_log"
	XfileServiceRequestLog_TableGo = "XfileServiceRequestLog"
	InvalidateCache_Table          = "invalidate_cache"
	InvalidateCache_TableGo        = "InvalidateCache"
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
	Id        string
	UserId    string
	ClientId  string
	Phone     string
	FirstName string
	LastName  string
}{

	Id:        "Id",
	UserId:    "UserId",
	ClientId:  "ClientId",
	Phone:     "Phone",
	FirstName: "FirstName",
	LastName:  "LastName",
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
	PostId          string
	CommentsCount   string
	LikesCount      string
	ViewsCount      string
	ReSharedCount   string
	ChatSharedCount string
}{

	PostId:          "PostId",
	CommentsCount:   "CommentsCount",
	LikesCount:      "LikesCount",
	ViewsCount:      "ViewsCount",
	ReSharedCount:   "ReSharedCount",
	ChatSharedCount: "ChatSharedCount",
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

var Session = struct {
	Id            string
	SessionUuid   string
	UserId        string
	LastIpAddress string
	AppVersion    string
	ActiveTime    string
	CreatedTime   string
}{

	Id:            "Id",
	SessionUuid:   "SessionUuid",
	UserId:        "UserId",
	LastIpAddress: "LastIpAddress",
	AppVersion:    "AppVersion",
	ActiveTime:    "ActiveTime",
	CreatedTime:   "CreatedTime",
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
	UserId         string
	UserName       string
	UserNameLower  string
	FirstName      string
	LastName       string
	IsVerified     string
	AvatarId       string
	ProfilePrivacy string
	OnlinePrivacy  string
	Phone          string
	Email          string
	About          string
	PasswordHash   string
	PasswordSalt   string
	PostSeq        string
	FollowersCount string
	FollowingCount string
	PostsCount     string
	MediaCount     string
	PhotoCount     string
	VideoCount     string
	GifCount       string
	AudioCount     string
	VoiceCount     string
	FileCount      string
	LinkCount      string
	BoardCount     string
	PinedCount     string
	LikesCount     string
	ResharedCount  string
	LastPostTime   string
	CreatedTime    string
	VersionTime    string
	IsDeleted      string
	IsBanned       string
}{

	UserId:         "UserId",
	UserName:       "UserName",
	UserNameLower:  "UserNameLower",
	FirstName:      "FirstName",
	LastName:       "LastName",
	IsVerified:     "IsVerified",
	AvatarId:       "AvatarId",
	ProfilePrivacy: "ProfilePrivacy",
	OnlinePrivacy:  "OnlinePrivacy",
	Phone:          "Phone",
	Email:          "Email",
	About:          "About",
	PasswordHash:   "PasswordHash",
	PasswordSalt:   "PasswordSalt",
	PostSeq:        "PostSeq",
	FollowersCount: "FollowersCount",
	FollowingCount: "FollowingCount",
	PostsCount:     "PostsCount",
	MediaCount:     "MediaCount",
	PhotoCount:     "PhotoCount",
	VideoCount:     "VideoCount",
	GifCount:       "GifCount",
	AudioCount:     "AudioCount",
	VoiceCount:     "VoiceCount",
	FileCount:      "FileCount",
	LinkCount:      "LinkCount",
	BoardCount:     "BoardCount",
	PinedCount:     "PinedCount",
	LikesCount:     "LikesCount",
	ResharedCount:  "ResharedCount",
	LastPostTime:   "LastPostTime",
	CreatedTime:    "CreatedTime",
	VersionTime:    "VersionTime",
	IsDeleted:      "IsDeleted",
	IsBanned:       "IsBanned",
}

var UserRelation = struct {
	RelNanoId     string
	UserId        string
	PeerUserId    string
	Follwing      string
	Followed      string
	InContacts    string
	MutualContact string
	IsFavorite    string
	Notify        string
}{

	RelNanoId:     "RelNanoId",
	UserId:        "UserId",
	PeerUserId:    "PeerUserId",
	Follwing:      "Follwing",
	Followed:      "Followed",
	InContacts:    "InContacts",
	MutualContact: "MutualContact",
	IsFavorite:    "IsFavorite",
	Notify:        "Notify",
}

var Chat = struct {
	ChatId           string
	ChatKey          string
	RoomKey          string
	RoomType         string
	UserId           string
	PeerUserId       string
	GroupId          string
	HashTagId        string
	StartedByMe      string
	Title            string
	PinTime          string
	FromMsgId        string
	Seq              string
	UnseenCount      string
	LastMsgId        string
	LastMsgStatus    string
	SeenSeq          string
	SeenMsgId        string
	LastMsgIdRecived string
	Left             string
	Creator          string
	Kicked           string
	Admin            string
	Deactivated      string
	VersionTime      string
	OrderTime        string
	CreatedTime      string
	DraftText        string
	DratReplyToMsgId string
}{

	ChatId:           "ChatId",
	ChatKey:          "ChatKey",
	RoomKey:          "RoomKey",
	RoomType:         "RoomType",
	UserId:           "UserId",
	PeerUserId:       "PeerUserId",
	GroupId:          "GroupId",
	HashTagId:        "HashTagId",
	StartedByMe:      "StartedByMe",
	Title:            "Title",
	PinTime:          "PinTime",
	FromMsgId:        "FromMsgId",
	Seq:              "Seq",
	UnseenCount:      "UnseenCount",
	LastMsgId:        "LastMsgId",
	LastMsgStatus:    "LastMsgStatus",
	SeenSeq:          "SeenSeq",
	SeenMsgId:        "SeenMsgId",
	LastMsgIdRecived: "LastMsgIdRecived",
	Left:             "Left",
	Creator:          "Creator",
	Kicked:           "Kicked",
	Admin:            "Admin",
	Deactivated:      "Deactivated",
	VersionTime:      "VersionTime",
	OrderTime:        "OrderTime",
	CreatedTime:      "CreatedTime",
	DraftText:        "DraftText",
	DratReplyToMsgId: "DratReplyToMsgId",
}

var ChatLastMessage = struct {
	ChatIdGroupId string
	LastMsgPb     string
}{

	ChatIdGroupId: "ChatIdGroupId",
	LastMsgPb:     "LastMsgPb",
}

var ChatVersionOrder = struct {
	VersionTime string
	UserId      string
	ChatId      string
	OrderTime   string
}{

	VersionTime: "VersionTime",
	UserId:      "UserId",
	ChatId:      "ChatId",
	OrderTime:   "OrderTime",
}

var Group = struct {
	GroupId         string
	GroupKey        string
	GroupName       string
	UserName        string
	IsSuperGroup    string
	HashTagId       string
	CreatorUserId   string
	GroupPrivacy    string
	HistoryViewAble string
	Seq             string
	LastMsgId       string
	PinedMsgId      string
	AvatarRefId     string
	AvatarCount     string
	About           string
	InviteLink      string
	MembersCount    string
	SortTime        string
	CreatedTime     string
	IsMute          string
}{

	GroupId:         "GroupId",
	GroupKey:        "GroupKey",
	GroupName:       "GroupName",
	UserName:        "UserName",
	IsSuperGroup:    "IsSuperGroup",
	HashTagId:       "HashTagId",
	CreatorUserId:   "CreatorUserId",
	GroupPrivacy:    "GroupPrivacy",
	HistoryViewAble: "HistoryViewAble",
	Seq:             "Seq",
	LastMsgId:       "LastMsgId",
	PinedMsgId:      "PinedMsgId",
	AvatarRefId:     "AvatarRefId",
	AvatarCount:     "AvatarCount",
	About:           "About",
	InviteLink:      "InviteLink",
	MembersCount:    "MembersCount",
	SortTime:        "SortTime",
	CreatedTime:     "CreatedTime",
	IsMute:          "IsMute",
}

var GroupMember = struct {
	OrderId     string
	GroupId     string
	UserId      string
	ByUserId    string
	GroupRole   string
	CreatedTime string
}{

	OrderId:     "OrderId",
	GroupId:     "GroupId",
	UserId:      "UserId",
	ByUserId:    "ByUserId",
	GroupRole:   "GroupRole",
	CreatedTime: "CreatedTime",
}

var GroupOrderdUser = struct {
	OrderId string
	GroupId string
	UserId  string
}{

	OrderId: "OrderId",
	GroupId: "GroupId",
	UserId:  "UserId",
}

var GroupPinedMsg = struct {
	MsgId string
	MsgPb string
}{

	MsgId: "MsgId",
	MsgPb: "MsgPb",
}

var FileMsg = struct {
	Id         string
	AccessHash string
	FileType   string
	Width      string
	Height     string
	Extension  string
	UserId     string
	DataThumb  string
	Data       string
}{

	Id:         "Id",
	AccessHash: "AccessHash",
	FileType:   "FileType",
	Width:      "Width",
	Height:     "Height",
	Extension:  "Extension",
	UserId:     "UserId",
	DataThumb:  "DataThumb",
	Data:       "Data",
}

var FilePost = struct {
	Id         string
	AccessHash string
	FileType   string
	Width      string
	Height     string
	Extension  string
	UserId     string
	DataThumb  string
	Data       string
}{

	Id:         "Id",
	AccessHash: "AccessHash",
	FileType:   "FileType",
	Width:      "Width",
	Height:     "Height",
	Extension:  "Extension",
	UserId:     "UserId",
	DataThumb:  "DataThumb",
	Data:       "Data",
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
	ResharedId string
}{

	OrderId:    "OrderId",
	ForUserId:  "ForUserId",
	PostId:     "PostId",
	PostUserId: "PostUserId",
	ResharedId: "ResharedId",
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

var PushChat = struct {
	PushId            string
	ToUserId          string
	PushTypeId        string
	RoomKey           string
	ChatKey           string
	Seq               string
	UnseenCount       string
	FromHighMessageId string
	ToLowMessageId    string
	MessageId         string
	MessageFileId     string
	MessagePb         string
	MessageJson       string
	CreatedTime       string
}{

	PushId:            "PushId",
	ToUserId:          "ToUserId",
	PushTypeId:        "PushTypeId",
	RoomKey:           "RoomKey",
	ChatKey:           "ChatKey",
	Seq:               "Seq",
	UnseenCount:       "UnseenCount",
	FromHighMessageId: "FromHighMessageId",
	ToLowMessageId:    "ToLowMessageId",
	MessageId:         "MessageId",
	MessageFileId:     "MessageFileId",
	MessagePb:         "MessagePb",
	MessageJson:       "MessageJson",
	CreatedTime:       "CreatedTime",
}

var HTTPRPCLog = struct {
	Id              string
	Time            string
	MethodFull      string
	MethodParent    string
	UserId          string
	SessionId       string
	StatusCode      string
	InputSize       string
	OutputSize      string
	ReqestJson      string
	ResponseJson    string
	ReqestParamJson string
	ResponseMsgJson string
}{

	Id:              "Id",
	Time:            "Time",
	MethodFull:      "MethodFull",
	MethodParent:    "MethodParent",
	UserId:          "UserId",
	SessionId:       "SessionId",
	StatusCode:      "StatusCode",
	InputSize:       "InputSize",
	OutputSize:      "OutputSize",
	ReqestJson:      "ReqestJson",
	ResponseJson:    "ResponseJson",
	ReqestParamJson: "ReqestParamJson",
	ResponseMsgJson: "ResponseMsgJson",
}

var MetricLog = struct {
	Id           string
	InstanceId   string
	StartFrom    string
	EndTo        string
	StartTime    string
	Duration     string
	MetericsJson string
}{

	Id:           "Id",
	InstanceId:   "InstanceId",
	StartFrom:    "StartFrom",
	EndTo:        "EndTo",
	StartTime:    "StartTime",
	Duration:     "Duration",
	MetericsJson: "MetericsJson",
}

var XfileServiceInfoLog = struct {
	Id          string
	InstanceId  string
	Url         string
	CreatedTime string
}{

	Id:          "Id",
	InstanceId:  "InstanceId",
	Url:         "Url",
	CreatedTime: "CreatedTime",
}

var XfileServiceMetricLog = struct {
	Id         string
	InstanceId string
	MetricJson string
}{

	Id:         "Id",
	InstanceId: "InstanceId",
	MetricJson: "MetricJson",
}

var XfileServiceRequestLog = struct {
	Id          string
	LocalSeq    string
	InstanceId  string
	Url         string
	HttpCode    string
	CreatedTime string
}{

	Id:          "Id",
	LocalSeq:    "LocalSeq",
	InstanceId:  "InstanceId",
	Url:         "Url",
	HttpCode:    "HttpCode",
	CreatedTime: "CreatedTime",
}

var InvalidateCache = struct {
	OrderId  string
	CacheKey string
}{

	OrderId:  "OrderId",
	CacheKey: "CacheKey",
}
