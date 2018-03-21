package x

// action 'Action'.
type Action struct {
	ActionId       int `db:"ActionId"`
	ActorUserId    int `db:"ActorUserId"`
	ActionTypeEnum int `db:"ActionTypeEnum"`
	PeerUserId     int `db:"PeerUserId"`
	PostId         int `db:"PostId"`
	CommentId      int `db:"CommentId"`
	Murmur64Hash   int `db:"Murmur64Hash"`
	CreatedTime    int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Action {
	ActionId: 0,
	ActorUserId: 0,
	ActionTypeEnum: 0,
	PeerUserId: 0,
	PostId: 0,
	CommentId: 0,
	Murmur64Hash: 0,
	CreatedTime: 0,
*/
// comment 'Comment'.
type Comment struct {
	CommentId   int    `db:"CommentId"`
	UserId      int    `db:"UserId"`
	PostId      int    `db:"PostId"`
	Text        string `db:"Text"`
	LikesCount  int    `db:"LikesCount"`
	CreatedTime int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Comment {
	CommentId: 0,
	UserId: 0,
	PostId: 0,
	Text: "",
	LikesCount: 0,
	CreatedTime: 0,
*/
// event 'Event'.
type Event struct {
	EventId      int `db:"EventId"`
	EventType    int `db:"EventType"`
	ByUserId     int `db:"ByUserId"`
	PeerUserId   int `db:"PeerUserId"`
	PostId       int `db:"PostId"`
	CommentId    int `db:"CommentId"`
	ActionId     int `db:"ActionId"`
	Murmur64Hash int `db:"Murmur64Hash"`

	_exists, _deleted bool
}

/*
:= &x.Event {
	EventId: 0,
	EventType: 0,
	ByUserId: 0,
	PeerUserId: 0,
	PostId: 0,
	CommentId: 0,
	ActionId: 0,
	Murmur64Hash: 0,
*/
// event2 'Event2'.
type Event2 struct {
	EventId      int `db:"EventId"`
	EventType    int `db:"EventType"`
	ByUserId     int `db:"ByUserId"`
	PeerUserId   int `db:"PeerUserId"`
	PostId       int `db:"PostId"`
	CommentId    int `db:"CommentId"`
	ActionId     int `db:"ActionId"`
	Murmur64Hash int `db:"Murmur64Hash"`

	_exists, _deleted bool
}

/*
:= &x.Event2 {
	EventId: 0,
	EventType: 0,
	ByUserId: 0,
	PeerUserId: 0,
	PostId: 0,
	CommentId: 0,
	ActionId: 0,
	Murmur64Hash: 0,
*/
// following_list 'FollowingList'.
type FollowingList struct {
	Id          int    `db:"Id"`
	UserId      int    `db:"UserId"`
	ListType    int    `db:"ListType"`
	Name        string `db:"Name"`
	Count       int    `db:"Count"`
	IsAuto      int    `db:"IsAuto"`
	IsPimiry    int    `db:"IsPimiry"`
	CreatedTime int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.FollowingList {
	Id: 0,
	UserId: 0,
	ListType: 0,
	Name: "",
	Count: 0,
	IsAuto: 0,
	IsPimiry: 0,
	CreatedTime: 0,
*/
// following_list_member 'FollowingListMember'.
type FollowingListMember struct {
	Id             int `db:"Id"`
	ListId         int `db:"ListId"`
	UserId         int `db:"UserId"`
	FollowedUserId int `db:"FollowedUserId"`
	CreatedTime    int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.FollowingListMember {
	Id: 0,
	ListId: 0,
	UserId: 0,
	FollowedUserId: 0,
	CreatedTime: 0,
*/
// following_list_member_removed 'FollowingListMemberRemoved'.
type FollowingListMemberRemoved struct {
	Id               int `db:"Id"`
	ListId           int `db:"ListId"`
	UserId           int `db:"UserId"`
	UnFollowedUserId int `db:"UnFollowedUserId"`
	UpdatedTime      int `db:"UpdatedTime"`

	_exists, _deleted bool
}

/*
:= &x.FollowingListMemberRemoved {
	Id: 0,
	ListId: 0,
	UserId: 0,
	UnFollowedUserId: 0,
	UpdatedTime: 0,
*/
// group 'Group'.
type Group struct {
	GroupId          int    `db:"GroupId"`
	GroupName        string `db:"GroupName"`
	MembersCount     int    `db:"MembersCount"`
	GroupPrivacyEnum int    `db:"GroupPrivacyEnum"`
	CreatorUserId    int    `db:"CreatorUserId"`
	CreatedTime      int    `db:"CreatedTime"`
	UpdatedMs        int    `db:"UpdatedMs"`
	CurrentSeq       int    `db:"CurrentSeq"`

	_exists, _deleted bool
}

/*
:= &x.Group {
	GroupId: 0,
	GroupName: "",
	MembersCount: 0,
	GroupPrivacyEnum: 0,
	CreatorUserId: 0,
	CreatedTime: 0,
	UpdatedMs: 0,
	CurrentSeq: 0,
*/
// group_member 'GroupMember'.
type GroupMember struct {
	Id              int    `db:"Id"`
	GroupId         int    `db:"GroupId"`
	GroupKey        string `db:"GroupKey"`
	UserId          int    `db:"UserId"`
	ByUserId        int    `db:"ByUserId"`
	GroupRoleEnumId int    `db:"GroupRoleEnumId"`
	CreatedTime     int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.GroupMember {
	Id: 0,
	GroupId: 0,
	GroupKey: "",
	UserId: 0,
	ByUserId: 0,
	GroupRoleEnumId: 0,
	CreatedTime: 0,
*/
// group_message 'GroupMessage'.
type GroupMessage struct {
	MessageId          int    `db:"MessageId"`
	RoomKey            string `db:"RoomKey"`
	UserId             int    `db:"UserId"`
	MessageFileId      int    `db:"MessageFileId"`
	MessageTypeEnum    int    `db:"MessageTypeEnum"`
	Text               string `db:"Text"`
	CreatedMs          int    `db:"CreatedMs"`
	DeliveryStatusEnum int    `db:"DeliveryStatusEnum"`

	_exists, _deleted bool
}

/*
:= &x.GroupMessage {
	MessageId: 0,
	RoomKey: "",
	UserId: 0,
	MessageFileId: 0,
	MessageTypeEnum: 0,
	Text: "",
	CreatedMs: 0,
	DeliveryStatusEnum: 0,
*/
// likes 'Like'.
type Like struct {
	Id           int `db:"Id"`
	PostId       int `db:"PostId"`
	PostTypeEnum int `db:"PostTypeEnum"`
	UserId       int `db:"UserId"`
	LikeEnum     int `db:"LikeEnum"`
	CreatedTime  int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Like {
	Id: 0,
	PostId: 0,
	PostTypeEnum: 0,
	UserId: 0,
	LikeEnum: 0,
	CreatedTime: 0,
*/
// media 'Media'.
type Media struct {
	MediaId       int    `db:"MediaId"`
	UserId        int    `db:"UserId"`
	PostId        int    `db:"PostId"`
	AlbumId       int    `db:"AlbumId"`
	MediaTypeEnum int    `db:"MediaTypeEnum"`
	Width         int    `db:"Width"`
	Height        int    `db:"Height"`
	Size          int    `db:"Size"`
	Duration      int    `db:"Duration"`
	Md5Hash       string `db:"Md5Hash"`
	Color         string `db:"Color"`
	CreatedTime   int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Media {
	MediaId: 0,
	UserId: 0,
	PostId: 0,
	AlbumId: 0,
	MediaTypeEnum: 0,
	Width: 0,
	Height: 0,
	Size: 0,
	Duration: 0,
	Md5Hash: "",
	Color: "",
	CreatedTime: 0,
*/
// notify 'Notify'.
type Notify struct {
	NotifyId       int `db:"NotifyId"`
	ForUserId      int `db:"ForUserId"`
	ActorUserId    int `db:"ActorUserId"`
	NotifyTypeEnum int `db:"NotifyTypeEnum"`
	PostId         int `db:"PostId"`
	CommentId      int `db:"CommentId"`
	PeerUserId     int `db:"PeerUserId"`
	Murmur64Hash   int `db:"Murmur64Hash"`
	SeenStatus     int `db:"SeenStatus"`
	CreatedTime    int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Notify {
	NotifyId: 0,
	ForUserId: 0,
	ActorUserId: 0,
	NotifyTypeEnum: 0,
	PostId: 0,
	CommentId: 0,
	PeerUserId: 0,
	Murmur64Hash: 0,
	SeenStatus: 0,
	CreatedTime: 0,
*/
// notify_removed 'NotifyRemoved'.
type NotifyRemoved struct {
	Murmur64Hash int `db:"Murmur64Hash"`
	ForUserId    int `db:"ForUserId"`
	Id           int `db:"Id"`

	_exists, _deleted bool
}

/*
:= &x.NotifyRemoved {
	Murmur64Hash: 0,
	ForUserId: 0,
	Id: 0,
*/
// phone_contacts 'PhoneContact'.
type PhoneContact struct {
	Id                    int    `db:"Id"`
	UserId                int    `db:"UserId"`
	Phone                 int    `db:"Phone"`
	PhoneDisplayName      string `db:"PhoneDisplayName"`
	PhoneFamilyName       string `db:"PhoneFamilyName"`
	PhoneNumber           string `db:"PhoneNumber"`
	PhoneNormalizedNumber string `db:"PhoneNormalizedNumber"`
	PhoneContactRowId     int    `db:"PhoneContactRowId"`
	DeviceUuidId          int    `db:"DeviceUuidId"`
	CreatedTime           int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.PhoneContact {
	Id: 0,
	UserId: 0,
	Phone: 0,
	PhoneDisplayName: "",
	PhoneFamilyName: "",
	PhoneNumber: "",
	PhoneNormalizedNumber: "",
	PhoneContactRowId: 0,
	DeviceUuidId: 0,
	CreatedTime: 0,
*/
// post 'Post'.
type Post struct {
	PostId         int    `db:"PostId"`
	UserId         int    `db:"UserId"`
	PostTypeEnum   int    `db:"PostTypeEnum"`
	MediaId        int    `db:"MediaId"`
	PostKey        string `db:"PostKey"`
	Text           string `db:"Text"`
	RichText       string `db:"RichText"`
	MediaCount     int    `db:"MediaCount"`
	SharedTo       int    `db:"SharedTo"`
	DisableComment int    `db:"DisableComment"`
	HasTag         int    `db:"HasTag"`
	Seq            int    `db:"Seq"`
	CommentsCount  int    `db:"CommentsCount"`
	LikesCount     int    `db:"LikesCount"`
	ViewsCount     int    `db:"ViewsCount"`
	EditedTime     int    `db:"EditedTime"`
	CreatedTime    int    `db:"CreatedTime"`
	ReSharedPostId int    `db:"ReSharedPostId"`

	_exists, _deleted bool
}

/*
:= &x.Post {
	PostId: 0,
	UserId: 0,
	PostTypeEnum: 0,
	MediaId: 0,
	PostKey: "",
	Text: "",
	RichText: "",
	MediaCount: 0,
	SharedTo: 0,
	DisableComment: 0,
	HasTag: 0,
	Seq: 0,
	CommentsCount: 0,
	LikesCount: 0,
	ViewsCount: 0,
	EditedTime: 0,
	CreatedTime: 0,
	ReSharedPostId: 0,
*/
// post_keys 'PostKey'.
type PostKey struct {
	Id         int    `db:"Id"`
	PostKeyStr string `db:"PostKeyStr"`
	Used       int    `db:"Used"`

	_exists, _deleted bool
}

/*
:= &x.PostKey {
	Id: 0,
	PostKeyStr: "",
	Used: 0,
*/
// search_clicked 'SearchClicked'.
type SearchClicked struct {
	Id          int    `db:"Id"`
	Query       string `db:"Query"`
	ClickType   int    `db:"ClickType"`
	TargetId    int    `db:"TargetId"`
	UserId      int    `db:"UserId"`
	CreatedTime int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.SearchClicked {
	Id: 0,
	Query: "",
	ClickType: 0,
	TargetId: 0,
	UserId: 0,
	CreatedTime: 0,
*/
// session 'Session'.
type Session struct {
	SessionUuid         string `db:"SessionUuid"`
	UserId              int    `db:"UserId"`
	ClientUuid          string `db:"ClientUuid"`
	DeviceUuid          string `db:"DeviceUuid"`
	LastActivityTime    int    `db:"LastActivityTime"`
	LastIpAddress       string `db:"LastIpAddress"`
	LastWifiMacAddress  string `db:"LastWifiMacAddress"`
	LastNetworkType     string `db:"LastNetworkType"`
	LastNetworkTypeEnum int    `db:"LastNetworkTypeEnum"`
	AppVersion          int    `db:"AppVersion"`
	UpdatedTime         int    `db:"UpdatedTime"`
	CreatedTime         int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Session {
	SessionUuid: "",
	UserId: 0,
	ClientUuid: "",
	DeviceUuid: "",
	LastActivityTime: 0,
	LastIpAddress: "",
	LastWifiMacAddress: "",
	LastNetworkType: "",
	LastNetworkTypeEnum: 0,
	AppVersion: 0,
	UpdatedTime: 0,
	CreatedTime: 0,
*/
// setting_client 'SettingClient'.
type SettingClient struct {
	UserId                    int `db:"UserId"`
	AutoDownloadWifiVoice     int `db:"AutoDownloadWifiVoice"`
	AutoDownloadWifiImage     int `db:"AutoDownloadWifiImage"`
	AutoDownloadWifiVideo     int `db:"AutoDownloadWifiVideo"`
	AutoDownloadWifiFile      int `db:"AutoDownloadWifiFile"`
	AutoDownloadWifiMusic     int `db:"AutoDownloadWifiMusic"`
	AutoDownloadWifiGif       int `db:"AutoDownloadWifiGif"`
	AutoDownloadCellularVoice int `db:"AutoDownloadCellularVoice"`
	AutoDownloadCellularImage int `db:"AutoDownloadCellularImage"`
	AutoDownloadCellularVideo int `db:"AutoDownloadCellularVideo"`
	AutoDownloadCellularFile  int `db:"AutoDownloadCellularFile"`
	AutoDownloadCellularMusic int `db:"AutoDownloadCellularMusic"`
	AutoDownloadCellularGif   int `db:"AutoDownloadCellularGif"`
	AutoDownloadRoamingVoice  int `db:"AutoDownloadRoamingVoice"`
	AutoDownloadRoamingImage  int `db:"AutoDownloadRoamingImage"`
	AutoDownloadRoamingVideo  int `db:"AutoDownloadRoamingVideo"`
	AutoDownloadRoamingFile   int `db:"AutoDownloadRoamingFile"`
	AutoDownloadRoamingMusic  int `db:"AutoDownloadRoamingMusic"`
	AutoDownloadRoamingGif    int `db:"AutoDownloadRoamingGif"`
	SaveToGallery             int `db:"SaveToGallery"`

	_exists, _deleted bool
}

/*
:= &x.SettingClient {
	UserId: 0,
	AutoDownloadWifiVoice: 0,
	AutoDownloadWifiImage: 0,
	AutoDownloadWifiVideo: 0,
	AutoDownloadWifiFile: 0,
	AutoDownloadWifiMusic: 0,
	AutoDownloadWifiGif: 0,
	AutoDownloadCellularVoice: 0,
	AutoDownloadCellularImage: 0,
	AutoDownloadCellularVideo: 0,
	AutoDownloadCellularFile: 0,
	AutoDownloadCellularMusic: 0,
	AutoDownloadCellularGif: 0,
	AutoDownloadRoamingVoice: 0,
	AutoDownloadRoamingImage: 0,
	AutoDownloadRoamingVideo: 0,
	AutoDownloadRoamingFile: 0,
	AutoDownloadRoamingMusic: 0,
	AutoDownloadRoamingGif: 0,
	SaveToGallery: 0,
*/
// setting_notifications 'SettingNotification'.
type SettingNotification struct {
	UserId                   int    `db:"UserId"`
	SocialLedOn              int    `db:"SocialLedOn"`
	SocialLedColor           string `db:"SocialLedColor"`
	ReqestToFollowYou        int    `db:"ReqestToFollowYou"`
	FollowedYou              int    `db:"FollowedYou"`
	AccptedYourFollowRequest int    `db:"AccptedYourFollowRequest"`
	YourPostLiked            int    `db:"YourPostLiked"`
	YourPostCommented        int    `db:"YourPostCommented"`
	MenthenedYouInPost       int    `db:"MenthenedYouInPost"`
	MenthenedYouInComment    int    `db:"MenthenedYouInComment"`
	YourContactsJoined       int    `db:"YourContactsJoined"`
	DirectMessage            int    `db:"DirectMessage"`
	DirectAlert              int    `db:"DirectAlert"`
	DirectPerview            int    `db:"DirectPerview"`
	DirectLedOn              int    `db:"DirectLedOn"`
	DirectLedColor           int    `db:"DirectLedColor"`
	DirectVibrate            int    `db:"DirectVibrate"`
	DirectPopup              int    `db:"DirectPopup"`
	DirectSound              int    `db:"DirectSound"`
	DirectPriority           int    `db:"DirectPriority"`

	_exists, _deleted bool
}

/*
:= &x.SettingNotification {
	UserId: 0,
	SocialLedOn: 0,
	SocialLedColor: "",
	ReqestToFollowYou: 0,
	FollowedYou: 0,
	AccptedYourFollowRequest: 0,
	YourPostLiked: 0,
	YourPostCommented: 0,
	MenthenedYouInPost: 0,
	MenthenedYouInComment: 0,
	YourContactsJoined: 0,
	DirectMessage: 0,
	DirectAlert: 0,
	DirectPerview: 0,
	DirectLedOn: 0,
	DirectLedColor: 0,
	DirectVibrate: 0,
	DirectPopup: 0,
	DirectSound: 0,
	DirectPriority: 0,
*/
// suggested_top_posts 'SuggestedTopPost'.
type SuggestedTopPost struct {
	Id     int `db:"Id"`
	PostId int `db:"PostId"`

	_exists, _deleted bool
}

/*
:= &x.SuggestedTopPost {
	Id: 0,
	PostId: 0,
*/
// suggested_user 'SuggestedUser'.
type SuggestedUser struct {
	Id          int     `db:"Id"`
	UserId      int     `db:"UserId"`
	TargetId    int     `db:"TargetId"`
	Weight      float32 `db:"Weight"`
	CreatedTime int     `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.SuggestedUser {
	Id: 0,
	UserId: 0,
	TargetId: 0,
	Weight: float32(0),
	CreatedTime: 0,
*/
// tag 'Tag'.
type Tag struct {
	TagId         int    `db:"TagId"`
	Name          string `db:"Name"`
	Count         int    `db:"Count"`
	TagStatusEnum int    `db:"TagStatusEnum"`
	CreatedTime   int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Tag {
	TagId: 0,
	Name: "",
	Count: 0,
	TagStatusEnum: 0,
	CreatedTime: 0,
*/
// tags_posts 'TagsPost'.
type TagsPost struct {
	Id           int `db:"Id"`
	TagId        int `db:"TagId"`
	PostId       int `db:"PostId"`
	PostTypeEnum int `db:"PostTypeEnum"`
	CreatedTime  int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.TagsPost {
	Id: 0,
	TagId: 0,
	PostId: 0,
	PostTypeEnum: 0,
	CreatedTime: 0,
*/
// trigger_log 'TriggerLog'.
type TriggerLog struct {
	Id         int    `db:"Id"`
	ModelName  string `db:"ModelName"`
	ChangeType string `db:"ChangeType"`
	TargetId   int    `db:"TargetId"`
	TargetStr  string `db:"TargetStr"`
	CreatedSe  int    `db:"CreatedSe"`

	_exists, _deleted bool
}

/*
:= &x.TriggerLog {
	Id: 0,
	ModelName: "",
	ChangeType: "",
	TargetId: 0,
	TargetStr: "",
	CreatedSe: 0,
*/
// user 'User'.
type User struct {
	UserId               int    `db:"UserId"`
	UserName             string `db:"UserName"`
	UserNameLower        string `db:"UserNameLower"` //nojson
	FirstName            string `db:"FirstName"`
	LastName             string `db:"LastName"`
	UserTypeEnum         int    `db:"UserTypeEnum"`
	UserLevelEnum        int    `db:"UserLevelEnum"`
	AvatarId             int    `db:"AvatarId"`
	ProfilePrivacyEnum   int    `db:"ProfilePrivacyEnum"`
	Phone                int    `db:"Phone"`
	About                string `db:"About"`
	Email                string `db:"Email"`
	PasswordHash         string `db:"PasswordHash"`
	PasswordSalt         string `db:"PasswordSalt"`
	PostSeq              int    `db:"PostSeq"`
	FollowersCount       int    `db:"FollowersCount"`
	FollowingCount       int    `db:"FollowingCount"`
	PostsCount           int    `db:"PostsCount"`
	MediaCount           int    `db:"MediaCount"`
	LikesCount           int    `db:"LikesCount"`
	ResharedCount        int    `db:"ResharedCount"`
	LastActionTime       int    `db:"LastActionTime"`
	LastPostTime         int    `db:"LastPostTime"`
	PrimaryFollowingList int    `db:"PrimaryFollowingList"`
	CreatedSe            int    `db:"CreatedSe"`
	UpdatedMs            int    `db:"UpdatedMs"`
	OnlinePrivacyEnum    int    `db:"OnlinePrivacyEnum"`
	LastActivityTime     int    `db:"LastActivityTime"`
	Phone2               string `db:"Phone2"`

	_exists, _deleted bool
}

/*
:= &x.User {
	UserId: 0,
	UserName: "",
	UserNameLower: "",
	FirstName: "",
	LastName: "",
	UserTypeEnum: 0,
	UserLevelEnum: 0,
	AvatarId: 0,
	ProfilePrivacyEnum: 0,
	Phone: 0,
	About: "",
	Email: "",
	PasswordHash: "",
	PasswordSalt: "",
	PostSeq: 0,
	FollowersCount: 0,
	FollowingCount: 0,
	PostsCount: 0,
	MediaCount: 0,
	LikesCount: 0,
	ResharedCount: 0,
	LastActionTime: 0,
	LastPostTime: 0,
	PrimaryFollowingList: 0,
	CreatedSe: 0,
	UpdatedMs: 0,
	OnlinePrivacyEnum: 0,
	LastActivityTime: 0,
	Phone2: "",
*/
// user_meta_info 'UserMetaInfo'.
type UserMetaInfo struct {
	Id                  int `db:"Id"`
	UserId              int `db:"UserId"`
	IsNotificationDirty int `db:"IsNotificationDirty"`
	LastUserRecGen      int `db:"LastUserRecGen"`

	_exists, _deleted bool
}

/*
:= &x.UserMetaInfo {
	Id: 0,
	UserId: 0,
	IsNotificationDirty: 0,
	LastUserRecGen: 0,
*/
// user_password 'UserPassword'.
type UserPassword struct {
	UserId      int    `db:"UserId"`
	Password    string `db:"Password"`
	CreatedTime int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.UserPassword {
	UserId: 0,
	Password: "",
	CreatedTime: 0,
*/
// chat 'Chat'.
type Chat struct {
	ChatKey      string `db:"ChatKey"`
	RoomKey      string `db:"RoomKey"`
	RoomTypeEnum int    `db:"RoomTypeEnum"`
	UserId       int    `db:"UserId"`
	PeerUserId   int    `db:"PeerUserId"`
	GroupId      int    `db:"GroupId"`
	CreatedTime  int    `db:"CreatedTime"`
	Seq          int    `db:"Seq"`
	SeenSeq      int    `db:"SeenSeq"`
	UpdatedMs    int    `db:"UpdatedMs"`

	_exists, _deleted bool
}

/*
:= &x.Chat {
	ChatKey: "",
	RoomKey: "",
	RoomTypeEnum: 0,
	UserId: 0,
	PeerUserId: 0,
	GroupId: 0,
	CreatedTime: 0,
	Seq: 0,
	SeenSeq: 0,
	UpdatedMs: 0,
*/
// chat_last_message 'ChatLastMessage'.
type ChatLastMessage struct {
	ChatKey     string `db:"ChatKey"`
	ForUserId   int    `db:"ForUserId"`
	LastMsgPb   []byte `db:"LastMsgPb"`
	LastMsgJson string `db:"LastMsgJson"`

	_exists, _deleted bool
}

/*
:= &x.ChatLastMessage {
	ChatKey: "",
	ForUserId: 0,
	LastMsgPb: []byte{},
	LastMsgJson: "",
*/
// chat_sync 'ChatSync'.
type ChatSync struct {
	SyncId            int    `db:"SyncId"`
	ToUserId          int    `db:"ToUserId"`
	ChatSyncTypeId    int    `db:"ChatSyncTypeId"`
	RoomKey           string `db:"RoomKey"`
	ChatKey           string `db:"ChatKey"`
	FromHighMessageId int    `db:"FromHighMessageId"`
	ToLowMessageId    int    `db:"ToLowMessageId"`
	MessageId         int    `db:"MessageId"`
	MessagePb         []byte `db:"MessagePb"`
	MessageJson       string `db:"MessageJson"`
	CreatedTime       int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.ChatSync {
	SyncId: 0,
	ToUserId: 0,
	ChatSyncTypeId: 0,
	RoomKey: "",
	ChatKey: "",
	FromHighMessageId: 0,
	ToLowMessageId: 0,
	MessageId: 0,
	MessagePb: []byte{},
	MessageJson: "",
	CreatedTime: 0,
*/
// direct_message 'DirectMessage'.
type DirectMessage struct {
	ChatKey            string `db:"ChatKey"`
	MessageId          int    `db:"MessageId"`
	RoomKey            string `db:"RoomKey"`
	UserId             int    `db:"UserId"`
	MessageFileId      int    `db:"MessageFileId"`
	MessageTypeEnum    int    `db:"MessageTypeEnum"`
	Text               string `db:"Text"`
	CreatedTime        int    `db:"CreatedTime"`
	Seq                int    `db:"Seq"`
	DeliviryStatusEnum int    `db:"DeliviryStatusEnum"`
	ExtraPB            []byte `db:"ExtraPB"`

	_exists, _deleted bool
}

/*
:= &x.DirectMessage {
	ChatKey: "",
	MessageId: 0,
	RoomKey: "",
	UserId: 0,
	MessageFileId: 0,
	MessageTypeEnum: 0,
	Text: "",
	CreatedTime: 0,
	Seq: 0,
	DeliviryStatusEnum: 0,
	ExtraPB: []byte{},
*/
// home 'Home'.
type Home struct {
	Id        int `db:"Id"`
	ForUserId int `db:"ForUserId"`
	PostId    int `db:"PostId"`

	_exists, _deleted bool
}

/*
:= &x.Home {
	Id: 0,
	ForUserId: 0,
	PostId: 0,
*/
// message_file 'MessageFile'.
type MessageFile struct {
	MessageFileId int    `db:"MessageFileId"`
	FileTypeEnum  int    `db:"FileTypeEnum"`
	UserId        int    `db:"UserId"` //orginal user id
	Title         string `db:"Title"`
	Size          int    `db:"Size"`
	Width         int    `db:"Width"`
	Height        int    `db:"Height"`
	Duration      int    `db:"Duration"`
	Extension     string `db:"Extension"`
	Md5Hash       string `db:"Md5Hash"`
	CreatedTime   int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.MessageFile {
	MessageFileId: 0,
	FileTypeEnum: 0,
	UserId: 0,
	Title: "",
	Size: 0,
	Width: 0,
	Height: 0,
	Duration: 0,
	Extension: "",
	Md5Hash: "",
	CreatedTime: 0,
*/
// file_msg 'FileMsg'.
type FileMsg struct {
	Id        int    `db:"Id"`
	FileType  int    `db:"FileType"`
	Extension string `db:"Extension"`
	DataThumb []byte `db:"DataThumb"`
	Data      []byte `db:"Data"`

	_exists, _deleted bool
}

/*
:= &x.FileMsg {
	Id: 0,
	FileType: 0,
	Extension: "",
	DataThumb: []byte{},
	Data: []byte{},
*/
// file_post 'FilePost'.
type FilePost struct {
	Id        int    `db:"Id"`
	FileType  int    `db:"FileType"`
	Extension string `db:"Extension"`
	DataThumb []byte `db:"DataThumb"`
	Data      []byte `db:"Data"`

	_exists, _deleted bool
}

/*
:= &x.FilePost {
	Id: 0,
	FileType: 0,
	Extension: "",
	DataThumb: []byte{},
	Data: []byte{},
*/
// action_fanout 'ActionFanout'.
type ActionFanout struct {
	OrderId     int `db:"OrderId"`
	ForUserId   int `db:"ForUserId"`
	ActionId    int `db:"ActionId"`
	ActorUserId int `db:"ActorUserId"`

	_exists, _deleted bool
}

/*
:= &x.ActionFanout {
	OrderId: 0,
	ForUserId: 0,
	ActionId: 0,
	ActorUserId: 0,
*/
// home_fanout 'HomeFanout'.
type HomeFanout struct {
	OrderId    int `db:"OrderId"`
	ForUserId  int `db:"ForUserId"`
	PostId     int `db:"PostId"`
	PostUserId int `db:"PostUserId"`

	_exists, _deleted bool
}

/*
:= &x.HomeFanout {
	OrderId: 0,
	ForUserId: 0,
	PostId: 0,
	PostUserId: 0,
*/

///////////////// Skip Loging Tables ////////////////
type LogTableSql struct {
	Action                     bool
	Comment                    bool
	Event                      bool
	Event2                     bool
	FollowingList              bool
	FollowingListMember        bool
	FollowingListMemberRemoved bool
	Group                      bool
	GroupMember                bool
	GroupMessage               bool
	Like                       bool
	Media                      bool
	Notify                     bool
	NotifyRemoved              bool
	PhoneContact               bool
	Post                       bool
	PostKey                    bool
	SearchClicked              bool
	Session                    bool
	SettingClient              bool
	SettingNotification        bool
	SuggestedTopPost           bool
	SuggestedUser              bool
	Tag                        bool
	TagsPost                   bool
	TriggerLog                 bool
	User                       bool
	UserMetaInfo               bool
	UserPassword               bool
	Chat                       bool
	ChatLastMessage            bool
	ChatSync                   bool
	DirectMessage              bool
	Home                       bool
	MessageFile                bool
	FileMsg                    bool
	FilePost                   bool
	ActionFanout               bool
	HomeFanout                 bool
}

var LogTableSqlReq = LogTableSql{

	Action:                     true,
	Comment:                    true,
	Event:                      true,
	Event2:                     true,
	FollowingList:              true,
	FollowingListMember:        true,
	FollowingListMemberRemoved: true,
	Group:               true,
	GroupMember:         true,
	GroupMessage:        true,
	Like:                true,
	Media:               true,
	Notify:              true,
	NotifyRemoved:       true,
	PhoneContact:        true,
	Post:                true,
	PostKey:             true,
	SearchClicked:       true,
	Session:             true,
	SettingClient:       true,
	SettingNotification: true,
	SuggestedTopPost:    true,
	SuggestedUser:       true,
	Tag:                 true,
	TagsPost:            true,
	TriggerLog:          true,
	User:                true,
	UserMetaInfo:        true,
	UserPassword:        true,
	Chat:                true,
	ChatLastMessage:     true,
	ChatSync:            true,
	DirectMessage:       true,
	Home:                true,
	MessageFile:         true,
	FileMsg:             true,
	FilePost:            true,
	ActionFanout:        true,
	HomeFanout:          true,
}
