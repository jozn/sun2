package xconst

const (
	Action_Table                       = "action"
	Action_TableGo                     = "Action"
	Chat_Table                         = "chat"
	Chat_TableGo                       = "Chat"
	Comment_Table                      = "comment"
	Comment_TableGo                    = "Comment"
	DirectMessage_Table                = "direct_message"
	DirectMessage_TableGo              = "DirectMessage"
	DirectOffline_Table                = "direct_offline"
	DirectOffline_TableGo              = "DirectOffline"
	DirectToMessage_Table              = "direct_to_message"
	DirectToMessage_TableGo            = "DirectToMessage"
	Feed_Table                         = "feed"
	Feed_TableGo                       = "Feed"
	FollowingList_Table                = "following_list"
	FollowingList_TableGo              = "FollowingList"
	FollowingListMember_Table          = "following_list_member"
	FollowingListMember_TableGo        = "FollowingListMember"
	FollowingListMemberHistory_Table   = "following_list_member_history"
	FollowingListMemberHistory_TableGo = "FollowingListMemberHistory"
	GeneralLog_Table                   = "general_log"
	GeneralLog_TableGo                 = "GeneralLog"
	Group_Table                        = "group"
	Group_TableGo                      = "Group"
	GroupMember_Table                  = "group_member"
	GroupMember_TableGo                = "GroupMember"
	GroupMessage_Table                 = "group_message"
	GroupMessage_TableGo               = "GroupMessage"
	Like_Table                         = "likes"
	Like_TableGo                       = "Like"
	Media_Table                        = "media"
	Media_TableGo                      = "Media"
	MessageFile_Table                  = "message_file"
	MessageFile_TableGo                = "MessageFile"
	Notify_Table                       = "notify"
	Notify_TableGo                     = "Notify"
	NotifyRemoved_Table                = "notify_removed"
	NotifyRemoved_TableGo              = "NotifyRemoved"
	PhoneContact_Table                 = "phone_contacts"
	PhoneContact_TableGo               = "PhoneContact"
	Post_Table                         = "post"
	Post_TableGo                       = "Post"
	RecommendUser_Table                = "recommend_user"
	RecommendUser_TableGo              = "RecommendUser"
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
	TagsPost_Table                     = "tags_posts"
	TagsPost_TableGo                   = "TagsPost"
	TriggerLog_Table                   = "trigger_log"
	TriggerLog_TableGo                 = "TriggerLog"
	User_Table                         = "user"
	User_TableGo                       = "User"
	UserMetaInfo_Table                 = "user_meta_info"
	UserMetaInfo_TableGo               = "UserMetaInfo"
	UserPassword_Table                 = "user_password"
	UserPassword_TableGo               = "UserPassword"
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
	Seq            string
}{

	ActionId:       "ActionId",
	ActorUserId:    "ActorUserId",
	ActionTypeEnum: "ActionTypeEnum",
	PeerUserId:     "PeerUserId",
	PostId:         "PostId",
	CommentId:      "CommentId",
	Murmur64Hash:   "Murmur64Hash",
	CreatedTime:    "CreatedTime",
	Seq:            "Seq",
}

var Chat = struct {
	ChatKey            string
	RoomKey            string
	RoomTypeEnum       string
	UserId             string
	PeerUserId         string
	GroupId            string
	CreatedTime        string
	StartMessageIdFrom string
	LastMessageId      string
	LastSeenMessageId  string
	UpdatedMs          string
}{

	ChatKey:            "ChatKey",
	RoomKey:            "RoomKey",
	RoomTypeEnum:       "RoomTypeEnum",
	UserId:             "UserId",
	PeerUserId:         "PeerUserId",
	GroupId:            "GroupId",
	CreatedTime:        "CreatedTime",
	StartMessageIdFrom: "StartMessageIdFrom",
	LastMessageId:      "LastMessageId",
	LastSeenMessageId:  "LastSeenMessageId",
	UpdatedMs:          "UpdatedMs",
}

var Comment = struct {
	CommentId   string
	UserId      string
	PostId      string
	Text        string
	LikesCount  string
	CreatedTime string
	Seq         string
}{

	CommentId:   "CommentId",
	UserId:      "UserId",
	PostId:      "PostId",
	Text:        "Text",
	LikesCount:  "LikesCount",
	CreatedTime: "CreatedTime",
	Seq:         "Seq",
}

var DirectMessage = struct {
	MessageId            string
	MessageKey           string
	RoomKey              string
	UserId               string
	MessageFileId        string
	MessageTypeEnumId    string
	Text                 string
	CreatedSe            string
	PeerReceivedTime     string
	PeerSeenTime         string
	DeliviryStatusEnumId string
}{

	MessageId:            "MessageId",
	MessageKey:           "MessageKey",
	RoomKey:              "RoomKey",
	UserId:               "UserId",
	MessageFileId:        "MessageFileId",
	MessageTypeEnumId:    "MessageTypeEnumId",
	Text:                 "Text",
	CreatedSe:            "CreatedSe",
	PeerReceivedTime:     "PeerReceivedTime",
	PeerSeenTime:         "PeerSeenTime",
	DeliviryStatusEnumId: "DeliviryStatusEnumId",
}

var DirectOffline = struct {
	DirectOfflineId string
	ToUserId        string
	ChatKey         string
	MessageId       string
	MessageFileId   string
	PBClass         string
	DataPB          string
	DataJson        string
	DataTemp        string
	AtTimeMs        string
}{

	DirectOfflineId: "DirectOfflineId",
	ToUserId:        "ToUserId",
	ChatKey:         "ChatKey",
	MessageId:       "MessageId",
	MessageFileId:   "MessageFileId",
	PBClass:         "PBClass",
	DataPB:          "DataPB",
	DataJson:        "DataJson",
	DataTemp:        "DataTemp",
	AtTimeMs:        "AtTimeMs",
}

var DirectToMessage = struct {
	Id           string
	ChatKey      string
	MessageId    string
	SourceEnumId string
}{

	Id:           "Id",
	ChatKey:      "ChatKey",
	MessageId:    "MessageId",
	SourceEnumId: "SourceEnumId",
}

var Feed = struct {
	UserId string
	PostId string
	FeedId string
}{

	UserId: "UserId",
	PostId: "PostId",
	FeedId: "FeedId",
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
	FollowType     string
	UpdatedTimeMs  string
}{

	Id:             "Id",
	ListId:         "ListId",
	UserId:         "UserId",
	FollowedUserId: "FollowedUserId",
	FollowType:     "FollowType",
	UpdatedTimeMs:  "UpdatedTimeMs",
}

var FollowingListMemberHistory = struct {
	Id             string
	ListId         string
	UserId         string
	FollowedUserId string
	FollowType     string
	UpdatedTimeMs  string
	FollowId       string
}{

	Id:             "Id",
	ListId:         "ListId",
	UserId:         "UserId",
	FollowedUserId: "FollowedUserId",
	FollowType:     "FollowType",
	UpdatedTimeMs:  "UpdatedTimeMs",
	FollowId:       "FollowId",
}

var GeneralLog = struct {
	Id        string
	ToUserId  string
	TargetId  string
	LogTypeId string
	ExtraPB   string
	ExtraJson string
	CreatedMs string
}{

	Id:        "Id",
	ToUserId:  "ToUserId",
	TargetId:  "TargetId",
	LogTypeId: "LogTypeId",
	ExtraPB:   "ExtraPB",
	ExtraJson: "ExtraJson",
	CreatedMs: "CreatedMs",
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

var Media = struct {
	MediaId       string
	UserId        string
	PostId        string
	AlbumId       string
	MediaTypeEnum string
	Width         string
	Height        string
	Size          string
	Duration      string
	Md5Hash       string
	Color         string
	CreatedTime   string
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
	Md5Hash:       "Md5Hash",
	Color:         "Color",
	CreatedTime:   "CreatedTime",
}

var MessageFile = struct {
	MessageFileId  string
	MessageFileKey string
	UserId         string
	Title          string
	Size           string
	FileTypeEnum   string
	Width          string
	Height         string
	Duration       string
	Extension      string
	Md5Hash        string
	CreatedTime    string
}{

	MessageFileId:  "MessageFileId",
	MessageFileKey: "MessageFileKey",
	UserId:         "UserId",
	Title:          "Title",
	Size:           "Size",
	FileTypeEnum:   "FileTypeEnum",
	Width:          "Width",
	Height:         "Height",
	Duration:       "Duration",
	Extension:      "Extension",
	Md5Hash:        "Md5Hash",
	CreatedTime:    "CreatedTime",
}

var Notify = struct {
	NotifyId      string
	ForUserId     string
	ActorUserId   string
	NotiyTypeEnum string
	PostId        string
	CommentId     string
	PeerUserId    string
	Murmur64Hash  string
	SeenStatus    string
	CreatedTime   string
	Seq           string
}{

	NotifyId:      "NotifyId",
	ForUserId:     "ForUserId",
	ActorUserId:   "ActorUserId",
	NotiyTypeEnum: "NotiyTypeEnum",
	PostId:        "PostId",
	CommentId:     "CommentId",
	PeerUserId:    "PeerUserId",
	Murmur64Hash:  "Murmur64Hash",
	SeenStatus:    "SeenStatus",
	CreatedTime:   "CreatedTime",
	Seq:           "Seq",
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
	PhoneDisplayName      string
	PhoneFamilyName       string
	PhoneNumber           string
	PhoneNormalizedNumber string
	PhoneContactRowId     string
	UserId                string
	DeviceUuidId          string
	CreatedTime           string
	UpdatedTime           string
}{

	Id:                    "Id",
	PhoneDisplayName:      "PhoneDisplayName",
	PhoneFamilyName:       "PhoneFamilyName",
	PhoneNumber:           "PhoneNumber",
	PhoneNormalizedNumber: "PhoneNormalizedNumber",
	PhoneContactRowId:     "PhoneContactRowId",
	UserId:                "UserId",
	DeviceUuidId:          "DeviceUuidId",
	CreatedTime:           "CreatedTime",
	UpdatedTime:           "UpdatedTime",
}

var Post = struct {
	PostId         string
	UserId         string
	PostTypeEnum   string
	MediaId        string
	Text           string
	RichText       string
	MediaCount     string
	SharedTo       string
	DisableComment string
	HasTag         string
	CommentsCount  string
	LikesCount     string
	ViewsCount     string
	EditedTime     string
	CreatedTime    string
	ReSharedPostId string
}{

	PostId:         "PostId",
	UserId:         "UserId",
	PostTypeEnum:   "PostTypeEnum",
	MediaId:        "MediaId",
	Text:           "Text",
	RichText:       "RichText",
	MediaCount:     "MediaCount",
	SharedTo:       "SharedTo",
	DisableComment: "DisableComment",
	HasTag:         "HasTag",
	CommentsCount:  "CommentsCount",
	LikesCount:     "LikesCount",
	ViewsCount:     "ViewsCount",
	EditedTime:     "EditedTime",
	CreatedTime:    "CreatedTime",
	ReSharedPostId: "ReSharedPostId",
}

var RecommendUser = struct {
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
	Id                    string
	UserId                string
	SessionUuid           string
	ClientUuid            string
	DeviceUuid            string
	LastActivityTime      string
	LastIpAddress         string
	LastWifiMacAddress    string
	LastNetworkType       string
	LastNetworkTypeEnumId string
	AppVersion            string
	UpdatedTime           string
	CreatedTime           string
}{

	Id:                    "Id",
	UserId:                "UserId",
	SessionUuid:           "SessionUuid",
	ClientUuid:            "ClientUuid",
	DeviceUuid:            "DeviceUuid",
	LastActivityTime:      "LastActivityTime",
	LastIpAddress:         "LastIpAddress",
	LastWifiMacAddress:    "LastWifiMacAddress",
	LastNetworkType:       "LastNetworkType",
	LastNetworkTypeEnumId: "LastNetworkTypeEnumId",
	AppVersion:            "AppVersion",
	UpdatedTime:           "UpdatedTime",
	CreatedTime:           "CreatedTime",
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

var TagsPost = struct {
	Id           string
	TagId        string
	PostId       string
	PostTypeEnum string
	CreatedTime  string
}{

	Id:           "Id",
	TagId:        "TagId",
	PostId:       "PostId",
	PostTypeEnum: "PostTypeEnum",
	CreatedTime:  "CreatedTime",
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
	FollowersCount       string
	FollowingCount       string
	PostsCount           string
	MediaCount           string
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
	FollowersCount:       "FollowersCount",
	FollowingCount:       "FollowingCount",
	PostsCount:           "PostsCount",
	MediaCount:           "MediaCount",
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
