package x

// action 'Action'.
type Action struct {
	ActionId       int
	ActorUserId    int
	ActionTypeEnum int
	PeerUserId     int
	PostId         int
	CommentId      int
	Murmur64Hash   int
	CreatedTime    int
	Seq            int

	_exists, _deleted bool
}

/*
:= &Action {
	ActionId: 0,
	ActorUserId: 0,
	ActionTypeEnum: 0,
	PeerUserId: 0,
	PostId: 0,
	CommentId: 0,
	Murmur64Hash: 0,
	CreatedTime: 0,
	Seq: 0,
*/
// comment 'Comment'.
type Comment struct {
	CommentId   int
	UserId      int
	PostId      int
	Text        string
	LikesCount  int
	CreatedTime int
	Seq         int

	_exists, _deleted bool
}

/*
:= &Comment {
	CommentId: 0,
	UserId: 0,
	PostId: 0,
	Text: "",
	LikesCount: 0,
	CreatedTime: 0,
	Seq: 0,
*/
// following_list 'FollowingList'.
type FollowingList struct {
	Id          int
	UserId      int
	ListType    int
	Name        string
	Count       int
	IsAuto      int
	IsPimiry    int
	CreatedTime int

	_exists, _deleted bool
}

/*
:= &FollowingList {
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
	Id             int
	ListId         int
	UserId         int
	FollowedUserId int
	CreatedTime    int

	_exists, _deleted bool
}

/*
:= &FollowingListMember {
	Id: 0,
	ListId: 0,
	UserId: 0,
	FollowedUserId: 0,
	CreatedTime: 0,
*/
// following_list_member_removed 'FollowingListMemberRemoved'.
type FollowingListMemberRemoved struct {
	Id               int
	ListId           int
	UserId           int
	UnFollowedUserId int
	UpdatedTime      int

	_exists, _deleted bool
}

/*
:= &FollowingListMemberRemoved {
	Id: 0,
	ListId: 0,
	UserId: 0,
	UnFollowedUserId: 0,
	UpdatedTime: 0,
*/
// group 'Group'.
type Group struct {
	GroupId          int
	GroupName        string
	MembersCount     int
	GroupPrivacyEnum int
	CreatorUserId    int
	CreatedTime      int
	UpdatedMs        int
	CurrentSeq       int

	_exists, _deleted bool
}

/*
:= &Group {
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
	Id              int
	GroupId         int
	GroupKey        string
	UserId          int
	ByUserId        int
	GroupRoleEnumId int
	CreatedTime     int

	_exists, _deleted bool
}

/*
:= &GroupMember {
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
	MessageId          int
	RoomKey            string
	UserId             int
	MessageFileId      int
	MessageTypeEnum    int
	Text               string
	CreatedMs          int
	DeliveryStatusEnum int

	_exists, _deleted bool
}

/*
:= &GroupMessage {
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
	Id           int
	PostId       int
	PostTypeEnum int
	UserId       int
	LikeEnum     int
	CreatedTime  int

	_exists, _deleted bool
}

/*
:= &Like {
	Id: 0,
	PostId: 0,
	PostTypeEnum: 0,
	UserId: 0,
	LikeEnum: 0,
	CreatedTime: 0,
*/
// media 'Media'.
type Media struct {
	MediaId       int
	UserId        int
	PostId        int
	AlbumId       int
	MediaTypeEnum int
	Width         int
	Height        int
	Size          int
	Duration      int
	Md5Hash       string
	Color         string
	CreatedTime   int

	_exists, _deleted bool
}

/*
:= &Media {
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
	NotifyId      int
	ForUserId     int
	ActorUserId   int
	NotiyTypeEnum int
	PostId        int
	CommentId     int
	PeerUserId    int
	Murmur64Hash  int
	SeenStatus    int
	CreatedTime   int
	Seq           int

	_exists, _deleted bool
}

/*
:= &Notify {
	NotifyId: 0,
	ForUserId: 0,
	ActorUserId: 0,
	NotiyTypeEnum: 0,
	PostId: 0,
	CommentId: 0,
	PeerUserId: 0,
	Murmur64Hash: 0,
	SeenStatus: 0,
	CreatedTime: 0,
	Seq: 0,
*/
// notify_removed 'NotifyRemoved'.
type NotifyRemoved struct {
	Murmur64Hash int
	ForUserId    int
	Id           int

	_exists, _deleted bool
}

/*
:= &NotifyRemoved {
	Murmur64Hash: 0,
	ForUserId: 0,
	Id: 0,
*/
// phone_contacts 'PhoneContact'.
type PhoneContact struct {
	Id                    int
	UserId                int
	Phone                 int
	PhoneDisplayName      string
	PhoneFamilyName       string
	PhoneNumber           string
	PhoneNormalizedNumber string
	PhoneContactRowId     int
	DeviceUuidId          int
	CreatedTime           int

	_exists, _deleted bool
}

/*
:= &PhoneContact {
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
	PostId         int
	UserId         int
	PostTypeEnum   int
	MediaId        int
	Text           string
	RichText       string
	MediaCount     int
	SharedTo       int
	DisableComment int
	HasTag         int
	CommentsCount  int
	LikesCount     int
	ViewsCount     int
	EditedTime     int
	CreatedTime    int
	ReSharedPostId int

	_exists, _deleted bool
}

/*
:= &Post {
	PostId: 0,
	UserId: 0,
	PostTypeEnum: 0,
	MediaId: 0,
	Text: "",
	RichText: "",
	MediaCount: 0,
	SharedTo: 0,
	DisableComment: 0,
	HasTag: 0,
	CommentsCount: 0,
	LikesCount: 0,
	ViewsCount: 0,
	EditedTime: 0,
	CreatedTime: 0,
	ReSharedPostId: 0,
*/
// post_keys 'PostKey'.
type PostKey struct {
	Id  int
	Key string

	_exists, _deleted bool
}

/*
:= &PostKey {
	Id: 0,
	Key: "",
*/
// search_clicked 'SearchClicked'.
type SearchClicked struct {
	Id          int
	Query       string
	ClickType   int
	TargetId    int
	UserId      int
	CreatedTime int

	_exists, _deleted bool
}

/*
:= &SearchClicked {
	Id: 0,
	Query: "",
	ClickType: 0,
	TargetId: 0,
	UserId: 0,
	CreatedTime: 0,
*/
// session 'Session'.
type Session struct {
	Id                    int
	UserId                int
	SessionUuid           string
	ClientUuid            string
	DeviceUuid            string
	LastActivityTime      int
	LastIpAddress         string
	LastWifiMacAddress    string
	LastNetworkType       string
	LastNetworkTypeEnumId int
	AppVersion            int
	UpdatedTime           int
	CreatedTime           int

	_exists, _deleted bool
}

/*
:= &Session {
	Id: 0,
	UserId: 0,
	SessionUuid: "",
	ClientUuid: "",
	DeviceUuid: "",
	LastActivityTime: 0,
	LastIpAddress: "",
	LastWifiMacAddress: "",
	LastNetworkType: "",
	LastNetworkTypeEnumId: 0,
	AppVersion: 0,
	UpdatedTime: 0,
	CreatedTime: 0,
*/
// setting_client 'SettingClient'.
type SettingClient struct {
	UserId                    int
	AutoDownloadWifiVoice     int
	AutoDownloadWifiImage     int
	AutoDownloadWifiVideo     int
	AutoDownloadWifiFile      int
	AutoDownloadWifiMusic     int
	AutoDownloadWifiGif       int
	AutoDownloadCellularVoice int
	AutoDownloadCellularImage int
	AutoDownloadCellularVideo int
	AutoDownloadCellularFile  int
	AutoDownloadCellularMusic int
	AutoDownloadCellularGif   int
	AutoDownloadRoamingVoice  int
	AutoDownloadRoamingImage  int
	AutoDownloadRoamingVideo  int
	AutoDownloadRoamingFile   int
	AutoDownloadRoamingMusic  int
	AutoDownloadRoamingGif    int
	SaveToGallery             int

	_exists, _deleted bool
}

/*
:= &SettingClient {
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
	UserId                   int
	SocialLedOn              int
	SocialLedColor           string
	ReqestToFollowYou        int
	FollowedYou              int
	AccptedYourFollowRequest int
	YourPostLiked            int
	YourPostCommented        int
	MenthenedYouInPost       int
	MenthenedYouInComment    int
	YourContactsJoined       int
	DirectMessage            int
	DirectAlert              int
	DirectPerview            int
	DirectLedOn              int
	DirectLedColor           int
	DirectVibrate            int
	DirectPopup              int
	DirectSound              int
	DirectPriority           int

	_exists, _deleted bool
}

/*
:= &SettingNotification {
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
	Id     int
	PostId int

	_exists, _deleted bool
}

/*
:= &SuggestedTopPost {
	Id: 0,
	PostId: 0,
*/
// suggested_user 'SuggestedUser'.
type SuggestedUser struct {
	Id          int
	UserId      int
	TargetId    int
	Weight      float32
	CreatedTime int

	_exists, _deleted bool
}

/*
:= &SuggestedUser {
	Id: 0,
	UserId: 0,
	TargetId: 0,
	Weight: float32(0),
	CreatedTime: 0,
*/
// tag 'Tag'.
type Tag struct {
	TagId         int
	Name          string
	Count         int
	TagStatusEnum int
	CreatedTime   int

	_exists, _deleted bool
}

/*
:= &Tag {
	TagId: 0,
	Name: "",
	Count: 0,
	TagStatusEnum: 0,
	CreatedTime: 0,
*/
// tags_posts 'TagsPost'.
type TagsPost struct {
	Id           int
	TagId        int
	PostId       int
	PostTypeEnum int
	CreatedTime  int

	_exists, _deleted bool
}

/*
:= &TagsPost {
	Id: 0,
	TagId: 0,
	PostId: 0,
	PostTypeEnum: 0,
	CreatedTime: 0,
*/
// trigger_log 'TriggerLog'.
type TriggerLog struct {
	Id         int
	ModelName  string
	ChangeType string
	TargetId   int
	TargetStr  string
	CreatedSe  int

	_exists, _deleted bool
}

/*
:= &TriggerLog {
	Id: 0,
	ModelName: "",
	ChangeType: "",
	TargetId: 0,
	TargetStr: "",
	CreatedSe: 0,
*/
// user 'User'.
type User struct {
	UserId               int
	UserName             string
	UserNameLower        string `json:"-"` //nojson
	FirstName            string
	LastName             string
	UserTypeEnum         int
	UserLevelEnum        int
	AvatarId             int
	ProfilePrivacyEnum   int
	Phone                int
	About                string
	Email                string
	PasswordHash         string
	PasswordSalt         string
	FollowersCount       int
	FollowingCount       int
	PostsCount           int
	MediaCount           int
	LikesCount           int
	ResharedCount        int
	LastActionTime       int
	LastPostTime         int
	PrimaryFollowingList int
	CreatedSe            int
	UpdatedMs            int
	OnlinePrivacyEnum    int
	LastActivityTime     int
	Phone2               string

	_exists, _deleted bool
}

/*
:= &User {
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
	Id                  int
	UserId              int
	IsNotificationDirty int
	LastUserRecGen      int

	_exists, _deleted bool
}

/*
:= &UserMetaInfo {
	Id: 0,
	UserId: 0,
	IsNotificationDirty: 0,
	LastUserRecGen: 0,
*/
// user_password 'UserPassword'.
type UserPassword struct {
	UserId      int
	Password    string
	CreatedTime int

	_exists, _deleted bool
}

/*
:= &UserPassword {
	UserId: 0,
	Password: "",
	CreatedTime: 0,
*/
// chat 'Chat'.
type Chat struct {
	ChatKey      string
	RoomKey      string
	RoomTypeEnum int
	UserId       int
	PeerUserId   int
	GroupId      int
	CreatedTime  int
	Seq          int
	SeenSeq      int
	UpdatedMs    int

	_exists, _deleted bool
}

/*
:= &Chat {
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
	ChatKey     string
	ForUserId   int
	LastMsgPb   []byte
	LastMsgJson string

	_exists, _deleted bool
}

/*
:= &ChatLastMessage {
	ChatKey: "",
	ForUserId: 0,
	LastMsgPb: []byte{},
	LastMsgJson: "",
*/
// chat_sync 'ChatSync'.
type ChatSync struct {
	SyncId            int
	ToUserId          int
	ChatSyncTypeId    int
	RoomKey           string
	ChatKey           string
	FromHighMessageId int
	ToLowMessageId    int
	MessageId         int
	MessagePb         []byte
	MessageJson       string
	CreatedTime       int

	_exists, _deleted bool
}

/*
:= &ChatSync {
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
	ChatKey            string
	MessageId          int
	RoomKey            string
	UserId             int
	MessageFileId      int
	MessageTypeEnum    int
	Text               string
	CreatedTime        int
	Seq                int
	DeliviryStatusEnum int
	ExtraPB            []byte

	_exists, _deleted bool
}

/*
:= &DirectMessage {
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
	Id        int
	ForUserId int
	PostId    int

	_exists, _deleted bool
}

/*
:= &Home {
	Id: 0,
	ForUserId: 0,
	PostId: 0,
*/
// message_file 'MessageFile'.
type MessageFile struct {
	MessageFileId int
	FileTypeEnum  int
	UserId        int //orginal user id
	Title         string
	Size          int
	Width         int
	Height        int
	Duration      int
	Extension     string
	Md5Hash       string
	CreatedTime   int

	_exists, _deleted bool
}

/*
:= &MessageFile {
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
	Id        int
	FileType  int
	Extension string
	DataThumb []byte
	Data      []byte

	_exists, _deleted bool
}

/*
:= &FileMsg {
	Id: 0,
	FileType: 0,
	Extension: "",
	DataThumb: []byte{},
	Data: []byte{},
*/
// file_post 'FilePost'.
type FilePost struct {
	Id        int
	FileType  int
	Extension string
	DataThumb []byte
	Data      []byte

	_exists, _deleted bool
}

/*
:= &FilePost {
	Id: 0,
	FileType: 0,
	Extension: "",
	DataThumb: []byte{},
	Data: []byte{},
*/
