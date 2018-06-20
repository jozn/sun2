package x

// action 'Action'.
type Action struct {
	ActionId     int `db:"ActionId"`
	ActorUserId  int `db:"ActorUserId"`
	ActionType   int `db:"ActionType"`
	PeerUserId   int `db:"PeerUserId"`
	PostId       int `db:"PostId"`
	CommentId    int `db:"CommentId"`
	Murmur64Hash int `db:"Murmur64Hash"`
	CreatedTime  int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Action {
	ActionId: 0,
	ActorUserId: 0,
	ActionType: 0,
	PeerUserId: 0,
	PostId: 0,
	CommentId: 0,
	Murmur64Hash: 0,
	CreatedTime: 0,
*/
// blocked 'Blocked'.
type Blocked struct {
	Id            int `db:"Id"`
	UserId        int `db:"UserId"`
	BlockedUserId int `db:"BlockedUserId"`
	CreatedTime   int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Blocked {
	Id: 0,
	UserId: 0,
	BlockedUserId: 0,
	CreatedTime: 0,
*/
// comment 'Comment'.
type Comment struct {
	CommentId   int    `db:"CommentId"`
	UserId      int    `db:"UserId"`
	PostId      int    `db:"PostId"`
	Text        string `db:"Text"`
	LikesCount  int    `db:"LikesCount"`
	IsEdited    int    `db:"IsEdited"`
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
	IsEdited: 0,
	CreatedTime: 0,
*/
// comment_deleted 'CommentDeleted'.
type CommentDeleted struct {
	CommentId int `db:"CommentId"`
	UserId    int `db:"UserId"`

	_exists, _deleted bool
}

/*
:= &x.CommentDeleted {
	CommentId: 0,
	UserId: 0,
*/
// event 'Event'.
type Event struct {
	EventId      int    `db:"EventId"`
	EventType    int    `db:"EventType"`
	ByUserId     int    `db:"ByUserId"`
	PeerUserId   int    `db:"PeerUserId"`
	PostId       int    `db:"PostId"`
	CommentId    int    `db:"CommentId"`
	HashTagId    int    `db:"HashTagId"`
	GroupId      int    `db:"GroupId"`
	ActionId     int    `db:"ActionId"`
	ChatId       int    `db:"ChatId"`
	ChatKey      string `db:"ChatKey"`
	MessageId    int    `db:"MessageId"`
	ReSharedId   int    `db:"ReSharedId"`
	Murmur64Hash int    `db:"Murmur64Hash"`

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
	HashTagId: 0,
	GroupId: 0,
	ActionId: 0,
	ChatId: 0,
	ChatKey: "",
	MessageId: 0,
	ReSharedId: 0,
	Murmur64Hash: 0,
*/
// followed 'Followed'.
type Followed struct {
	Id             int `db:"Id"`
	UserId         int `db:"UserId"`
	FollowedUserId int `db:"FollowedUserId"`
	CreatedTime    int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Followed {
	Id: 0,
	UserId: 0,
	FollowedUserId: 0,
	CreatedTime: 0,
*/
// likes 'Likes'.
type Likes struct {
	Id          int `db:"Id"`
	PostId      int `db:"PostId"`
	UserId      int `db:"UserId"`
	PostType    int `db:"PostType"`
	CreatedTime int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Likes {
	Id: 0,
	PostId: 0,
	UserId: 0,
	PostType: 0,
	CreatedTime: 0,
*/
// notify 'Notify'.
type Notify struct {
	NotifyId     int `db:"NotifyId"`
	ForUserId    int `db:"ForUserId"`
	ActorUserId  int `db:"ActorUserId"`
	NotifyType   int `db:"NotifyType"`
	PostId       int `db:"PostId"`
	CommentId    int `db:"CommentId"`
	PeerUserId   int `db:"PeerUserId"`
	Murmur64Hash int `db:"Murmur64Hash"`
	SeenStatus   int `db:"SeenStatus"`
	CreatedTime  int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Notify {
	NotifyId: 0,
	ForUserId: 0,
	ActorUserId: 0,
	NotifyType: 0,
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
// phone_contacts 'PhoneContacts'.
type PhoneContacts struct {
	Id        int    `db:"Id"`
	UserId    int    `db:"UserId"`
	ClientId  int    `db:"ClientId"`
	Phone     string `db:"Phone"`
	FirstName string `db:"FirstName"`
	LastName  string `db:"LastName"`

	_exists, _deleted bool
}

/*
:= &x.PhoneContacts {
	Id: 0,
	UserId: 0,
	ClientId: 0,
	Phone: "",
	FirstName: "",
	LastName: "",
*/
// post 'Post'.
type Post struct {
	PostId         int    `db:"PostId"`
	UserId         int    `db:"UserId"`
	PostType       int    `db:"PostType"`
	MediaId        int    `db:"MediaId"`
	FileRefId      int    `db:"FileRefId"`
	PostKey        string `db:"PostKey"`
	Text           string `db:"Text"`
	RichText       string `db:"RichText"`
	MediaCount     int    `db:"MediaCount"`
	SharedTo       int    `db:"SharedTo"`
	DisableComment int    `db:"DisableComment"`
	Via            int    `db:"Via"`
	Seq            int    `db:"Seq"`
	CommentsCount  int    `db:"CommentsCount"`
	LikesCount     int    `db:"LikesCount"`
	ViewsCount     int    `db:"ViewsCount"`
	EditedTime     int    `db:"EditedTime"`
	CreatedTime    int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Post {
	PostId: 0,
	UserId: 0,
	PostType: 0,
	MediaId: 0,
	FileRefId: 0,
	PostKey: "",
	Text: "",
	RichText: "",
	MediaCount: 0,
	SharedTo: 0,
	DisableComment: 0,
	Via: 0,
	Seq: 0,
	CommentsCount: 0,
	LikesCount: 0,
	ViewsCount: 0,
	EditedTime: 0,
	CreatedTime: 0,
*/
// post_count 'PostCount'.
type PostCount struct {
	PostId          int `db:"PostId"`
	CommentsCount   int `db:"CommentsCount"`
	LikesCount      int `db:"LikesCount"`
	ViewsCount      int `db:"ViewsCount"`
	ReSharedCount   int `db:"ReSharedCount"`
	ChatSharedCount int `db:"ChatSharedCount"`

	_exists, _deleted bool
}

/*
:= &x.PostCount {
	PostId: 0,
	CommentsCount: 0,
	LikesCount: 0,
	ViewsCount: 0,
	ReSharedCount: 0,
	ChatSharedCount: 0,
*/
// post_deleted 'PostDeleted'.
type PostDeleted struct {
	PostId int `db:"PostId"`
	UserId int `db:"UserId"`

	_exists, _deleted bool
}

/*
:= &x.PostDeleted {
	PostId: 0,
	UserId: 0,
*/
// post_keys 'PostKeys'.
type PostKeys struct {
	Id         int    `db:"Id"`
	PostKeyStr string `db:"PostKeyStr"`
	Used       int    `db:"Used"`
	RandShard  int    `db:"RandShard"` //1-10000

	_exists, _deleted bool
}

/*
:= &x.PostKeys {
	Id: 0,
	PostKeyStr: "",
	Used: 0,
	RandShard: 0,
*/
// post_link 'PostLink'.
type PostLink struct {
	LinkId  int    `db:"LinkId"`
	LinkUrl string `db:"LinkUrl"`

	_exists, _deleted bool
}

/*
:= &x.PostLink {
	LinkId: 0,
	LinkUrl: "",
*/
// post_media 'PostMedia'.
type PostMedia struct {
	MediaId       int    `db:"MediaId"`
	UserId        int    `db:"UserId"`
	PostId        int    `db:"PostId"`
	AlbumId       int    `db:"AlbumId"`
	MediaTypeEnum int    `db:"MediaTypeEnum"`
	Width         int    `db:"Width"`
	Height        int    `db:"Height"`
	Size          int    `db:"Size"`
	Duration      int    `db:"Duration"`
	Extension     string `db:"Extension"`
	Md5Hash       string `db:"Md5Hash"`
	Color         string `db:"Color"`
	CreatedTime   int    `db:"CreatedTime"`
	ViewCount     int    `db:"ViewCount"`
	Extra         string `db:"Extra"`

	_exists, _deleted bool
}

/*
:= &x.PostMedia {
	MediaId: 0,
	UserId: 0,
	PostId: 0,
	AlbumId: 0,
	MediaTypeEnum: 0,
	Width: 0,
	Height: 0,
	Size: 0,
	Duration: 0,
	Extension: "",
	Md5Hash: "",
	Color: "",
	CreatedTime: 0,
	ViewCount: 0,
	Extra: "",
*/
// post_promoted 'PostPromoted'.
type PostPromoted struct {
	PromoteId   int    `db:"PromoteId"`
	PostId      int    `db:"PostId"`
	ByUserId    int    `db:"ByUserId"`
	PostUserId  int    `db:"PostUserId"`
	BazzarUuid  string `db:"BazzarUuid"`
	Package     string `db:"Package"`
	EndTime     int    `db:"EndTime"`
	CreatedTime int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.PostPromoted {
	PromoteId: 0,
	PostId: 0,
	ByUserId: 0,
	PostUserId: 0,
	BazzarUuid: "",
	Package: "",
	EndTime: 0,
	CreatedTime: 0,
*/
// post_reshared 'PostReshared'.
type PostReshared struct {
	ResharedId  int `db:"ResharedId"`
	PostId      int `db:"PostId"`
	ByUserId    int `db:"ByUserId"`
	PostUserId  int `db:"PostUserId"`
	CreatedTime int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.PostReshared {
	ResharedId: 0,
	PostId: 0,
	ByUserId: 0,
	PostUserId: 0,
	CreatedTime: 0,
*/
// profile_all 'ProfileAll'.
type ProfileAll struct {
	Id         int `db:"Id"`
	UserId     int `db:"UserId"`
	PostId     int `db:"PostId"`
	IsReShared int `db:"IsReShared"`

	_exists, _deleted bool
}

/*
:= &x.ProfileAll {
	Id: 0,
	UserId: 0,
	PostId: 0,
	IsReShared: 0,
*/
// profile_media 'ProfileMedia'.
type ProfileMedia struct {
	Id       int `db:"Id"`
	UserId   int `db:"UserId"`
	PostId   int `db:"PostId"`
	PostType int `db:"PostType"`

	_exists, _deleted bool
}

/*
:= &x.ProfileMedia {
	Id: 0,
	UserId: 0,
	PostId: 0,
	PostType: 0,
*/
// profile_mentioned 'ProfileMentioned'.
type ProfileMentioned struct {
	Id          int `db:"Id"`
	ForUserId   int `db:"ForUserId"`
	PostId      int `db:"PostId"`
	PostUserId  int `db:"PostUserId"`
	PostType    int `db:"PostType"`
	CreatedTime int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.ProfileMentioned {
	Id: 0,
	ForUserId: 0,
	PostId: 0,
	PostUserId: 0,
	PostType: 0,
	CreatedTime: 0,
*/
// session 'Session'.
type Session struct {
	Id            int    `db:"Id"`
	SessionUuid   string `db:"SessionUuid"`
	UserId        int    `db:"UserId"`
	LastIpAddress string `db:"LastIpAddress"`
	AppVersion    int    `db:"AppVersion"`
	ActiveTime    int    `db:"ActiveTime"`
	CreatedTime   int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Session {
	Id: 0,
	SessionUuid: "",
	UserId: 0,
	LastIpAddress: "",
	AppVersion: 0,
	ActiveTime: 0,
	CreatedTime: 0,
*/
// setting_notifications 'SettingNotifications'.
type SettingNotifications struct {
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
:= &x.SettingNotifications {
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
// sms 'Sms'.
type Sms struct {
	Id              int    `db:"Id"`
	Hash            string `db:"Hash"`
	AppUuid         string `db:"AppUuid"`
	ClientPhone     string `db:"ClientPhone"`
	GenratedCode    int    `db:"GenratedCode"`
	SmsSenderNumber string `db:"SmsSenderNumber"`
	SmsSendStatues  string `db:"SmsSendStatues"`
	SmsHttpBody     string `db:"SmsHttpBody"`
	Err             string `db:"Err"`
	Carrier         string `db:"Carrier"`
	Country         []byte `db:"Country"`
	IsValidPhone    int    `db:"IsValidPhone"`
	IsConfirmed     int    `db:"IsConfirmed"`
	IsLogin         int    `db:"IsLogin"`
	IsRegister      int    `db:"IsRegister"`
	RetriedCount    int    `db:"RetriedCount"`
	TTL             int    `db:"TTL"`

	_exists, _deleted bool
}

/*
:= &x.Sms {
	Id: 0,
	Hash: "",
	AppUuid: "",
	ClientPhone: "",
	GenratedCode: 0,
	SmsSenderNumber: "",
	SmsSendStatues: "",
	SmsHttpBody: "",
	Err: "",
	Carrier: "",
	Country: []byte{},
	IsValidPhone: 0,
	IsConfirmed: 0,
	IsLogin: 0,
	IsRegister: 0,
	RetriedCount: 0,
	TTL: 0,
*/
// tag 'Tag'.
type Tag struct {
	TagId         int    `db:"TagId"`
	Name          string `db:"Name"`
	Count         int    `db:"Count"`
	TagStatusEnum int    `db:"TagStatusEnum"` //del
	IsBlocked     int    `db:"IsBlocked"`
	GroupId       int    `db:"GroupId"`
	CreatedTime   int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.Tag {
	TagId: 0,
	Name: "",
	Count: 0,
	TagStatusEnum: 0,
	IsBlocked: 0,
	GroupId: 0,
	CreatedTime: 0,
*/
// tag_post 'TagPost'.
type TagPost struct {
	Id          int `db:"Id"`
	TagId       int `db:"TagId"`
	PostId      int `db:"PostId"`
	PostType    int `db:"PostType"`
	CreatedTime int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.TagPost {
	Id: 0,
	TagId: 0,
	PostId: 0,
	PostType: 0,
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
	UserId             int    `db:"UserId"`
	UserName           string `db:"UserName"`
	UserNameLower      string `db:"UserNameLower"` //nojson
	FirstName          string `db:"FirstName"`
	LastName           string `db:"LastName"`
	IsVerified         int    `db:"IsVerified"`
	AvatarId           int    `db:"AvatarId"`
	AccessHash         int    `db:"AccessHash"`
	ProfilePrivacy     int    `db:"ProfilePrivacy"`
	OnlinePrivacy      int    `db:"OnlinePrivacy"`
	CallPrivacy        int    `db:"CallPrivacy"`
	AddToGroupPrivacy  int    `db:"AddToGroupPrivacy"`
	SeenMessagePrivacy int    `db:"SeenMessagePrivacy"`
	Phone              string `db:"Phone"`
	Email              string `db:"Email"`
	About              string `db:"About"`
	PasswordHash       string `db:"PasswordHash"`
	PasswordSalt       string `db:"PasswordSalt"`
	PostSeq            int    `db:"PostSeq"`
	FollowersCount     int    `db:"FollowersCount"`
	FollowingCount     int    `db:"FollowingCount"`
	PostsCount         int    `db:"PostsCount"`
	MediaCount         int    `db:"MediaCount"`
	PhotoCount         int    `db:"PhotoCount"`
	VideoCount         int    `db:"VideoCount"`
	GifCount           int    `db:"GifCount"`
	AudioCount         int    `db:"AudioCount"`
	VoiceCount         int    `db:"VoiceCount"`
	FileCount          int    `db:"FileCount"`
	LinkCount          int    `db:"LinkCount"`
	BoardCount         int    `db:"BoardCount"`
	PinedCount         int    `db:"PinedCount"`
	LikesCount         int    `db:"LikesCount"`
	ResharedCount      int    `db:"ResharedCount"`
	LastPostTime       int    `db:"LastPostTime"`
	CreatedTime        int    `db:"CreatedTime"`
	VersionTime        int    `db:"VersionTime"`
	IsDeleted          int    `db:"IsDeleted"`
	IsBanned           int    `db:"IsBanned"`

	_exists, _deleted bool
}

/*
:= &x.User {
	UserId: 0,
	UserName: "",
	UserNameLower: "",
	FirstName: "",
	LastName: "",
	IsVerified: 0,
	AvatarId: 0,
	AccessHash: 0,
	ProfilePrivacy: 0,
	OnlinePrivacy: 0,
	CallPrivacy: 0,
	AddToGroupPrivacy: 0,
	SeenMessagePrivacy: 0,
	Phone: "",
	Email: "",
	About: "",
	PasswordHash: "",
	PasswordSalt: "",
	PostSeq: 0,
	FollowersCount: 0,
	FollowingCount: 0,
	PostsCount: 0,
	MediaCount: 0,
	PhotoCount: 0,
	VideoCount: 0,
	GifCount: 0,
	AudioCount: 0,
	VoiceCount: 0,
	FileCount: 0,
	LinkCount: 0,
	BoardCount: 0,
	PinedCount: 0,
	LikesCount: 0,
	ResharedCount: 0,
	LastPostTime: 0,
	CreatedTime: 0,
	VersionTime: 0,
	IsDeleted: 0,
	IsBanned: 0,
*/
// user_relation 'UserRelation'.
type UserRelation struct {
	RelNanoId     int `db:"RelNanoId"`
	UserId        int `db:"UserId"`
	PeerUserId    int `db:"PeerUserId"`
	Follwing      int `db:"Follwing"`
	Followed      int `db:"Followed"`
	InContacts    int `db:"InContacts"`
	MutualContact int `db:"MutualContact"`
	IsFavorite    int `db:"IsFavorite"`
	Notify        int `db:"Notify"`

	_exists, _deleted bool
}

/*
:= &x.UserRelation {
	RelNanoId: 0,
	UserId: 0,
	PeerUserId: 0,
	Follwing: 0,
	Followed: 0,
	InContacts: 0,
	MutualContact: 0,
	IsFavorite: 0,
	Notify: 0,
*/
// chat 'Chat'.
type Chat struct {
	ChatId                 int    `db:"ChatId"`
	ChatKey                string `db:"ChatKey"`
	RoomKey                string `db:"RoomKey"`
	RoomType               int    `db:"RoomType"`
	UserId                 int    `db:"UserId"`
	PeerUserId             int    `db:"PeerUserId"`
	GroupId                int    `db:"GroupId"`
	HashTagId              int    `db:"HashTagId"`
	Title                  string `db:"Title"`
	PinTimeMs              int    `db:"PinTimeMs"`
	FromMsgId              int    `db:"FromMsgId"`
	UnseenCount            int    `db:"UnseenCount"`
	Seq                    int    `db:"Seq"`
	LastMsgId              int    `db:"LastMsgId"`
	LastMyMsgStatus        int    `db:"LastMyMsgStatus"`
	MyLastSeenSeq          int    `db:"MyLastSeenSeq"`
	MyLastSeenMsgId        int    `db:"MyLastSeenMsgId"`
	PeerLastSeenMsgId      int    `db:"PeerLastSeenMsgId"`
	MyLastDeliveredSeq     int    `db:"MyLastDeliveredSeq"`
	MyLastDeliveredMsgId   int    `db:"MyLastDeliveredMsgId"`
	PeerLastDeliveredMsgId int    `db:"PeerLastDeliveredMsgId"`
	IsActive               int    `db:"IsActive"`
	IsLeft                 int    `db:"IsLeft"`
	IsCreator              int    `db:"IsCreator"`
	IsKicked               int    `db:"IsKicked"`
	IsAdmin                int    `db:"IsAdmin"`
	IsDeactivated          int    `db:"IsDeactivated"`
	IsMuted                int    `db:"IsMuted"`
	MuteUntil              int    `db:"MuteUntil"`
	VersionTimeMs          int    `db:"VersionTimeMs"`
	VersionSeq             int    `db:"VersionSeq"`
	OrderTime              int    `db:"OrderTime"`
	CreatedTime            int    `db:"CreatedTime"`
	DraftText              string `db:"DraftText"`
	DratReplyToMsgId       int    `db:"DratReplyToMsgId"`

	_exists, _deleted bool
}

/*
:= &x.Chat {
	ChatId: 0,
	ChatKey: "",
	RoomKey: "",
	RoomType: 0,
	UserId: 0,
	PeerUserId: 0,
	GroupId: 0,
	HashTagId: 0,
	Title: "",
	PinTimeMs: 0,
	FromMsgId: 0,
	UnseenCount: 0,
	Seq: 0,
	LastMsgId: 0,
	LastMyMsgStatus: 0,
	MyLastSeenSeq: 0,
	MyLastSeenMsgId: 0,
	PeerLastSeenMsgId: 0,
	MyLastDeliveredSeq: 0,
	MyLastDeliveredMsgId: 0,
	PeerLastDeliveredMsgId: 0,
	IsActive: 0,
	IsLeft: 0,
	IsCreator: 0,
	IsKicked: 0,
	IsAdmin: 0,
	IsDeactivated: 0,
	IsMuted: 0,
	MuteUntil: 0,
	VersionTimeMs: 0,
	VersionSeq: 0,
	OrderTime: 0,
	CreatedTime: 0,
	DraftText: "",
	DratReplyToMsgId: 0,
*/
// chat_deleted 'ChatDeleted'.
type ChatDeleted struct {
	ChatId  int    `db:"ChatId"`
	RoomKey string `db:"RoomKey"`

	_exists, _deleted bool
}

/*
:= &x.ChatDeleted {
	ChatId: 0,
	RoomKey: "",
*/
// chat_last_message 'ChatLastMessage'.
type ChatLastMessage struct {
	ChatIdGroupId string `db:"ChatIdGroupId"`
	LastMsgPb     []byte `db:"LastMsgPb"`

	_exists, _deleted bool
}

/*
:= &x.ChatLastMessage {
	ChatIdGroupId: "",
	LastMsgPb: []byte{},
*/
// chat_user_version 'ChatUserVersion'.
type ChatUserVersion struct {
	UserId          int `db:"UserId"`
	ChatId          int `db:"ChatId"`
	VersionTimeNano int `db:"VersionTimeNano"`

	_exists, _deleted bool
}

/*
:= &x.ChatUserVersion {
	UserId: 0,
	ChatId: 0,
	VersionTimeNano: 0,
*/
// group 'Group'.
type Group struct {
	GroupId         int    `db:"GroupId"`
	GroupKey        string `db:"GroupKey"`
	GroupName       string `db:"GroupName"`
	UserName        string `db:"UserName"`
	IsSuperGroup    int    `db:"IsSuperGroup"`
	HashTagId       int    `db:"HashTagId"`
	CreatorUserId   int    `db:"CreatorUserId"`
	GroupPrivacy    int    `db:"GroupPrivacy"`
	HistoryViewAble int    `db:"HistoryViewAble"`
	Seq             int    `db:"Seq"`
	LastMsgId       int    `db:"LastMsgId"`
	PinedMsgId      int    `db:"PinedMsgId"`
	AvatarRefId     int    `db:"AvatarRefId"`
	AvatarCount     int    `db:"AvatarCount"`
	About           string `db:"About"`
	InviteLinkHash  string `db:"InviteLinkHash"`
	MembersCount    int    `db:"MembersCount"`
	AdminsCount     int    `db:"AdminsCount"`
	ModeratorCounts int    `db:"ModeratorCounts"`
	SortTime        int    `db:"SortTime"`
	CreatedTime     int    `db:"CreatedTime"`
	IsMute          int    `db:"IsMute"`
	IsDeleted       int    `db:"IsDeleted"`
	IsBanned        int    `db:"IsBanned"`

	_exists, _deleted bool
}

/*
:= &x.Group {
	GroupId: 0,
	GroupKey: "",
	GroupName: "",
	UserName: "",
	IsSuperGroup: 0,
	HashTagId: 0,
	CreatorUserId: 0,
	GroupPrivacy: 0,
	HistoryViewAble: 0,
	Seq: 0,
	LastMsgId: 0,
	PinedMsgId: 0,
	AvatarRefId: 0,
	AvatarCount: 0,
	About: "",
	InviteLinkHash: "",
	MembersCount: 0,
	AdminsCount: 0,
	ModeratorCounts: 0,
	SortTime: 0,
	CreatedTime: 0,
	IsMute: 0,
	IsDeleted: 0,
	IsBanned: 0,
*/
// group_member 'GroupMember'.
type GroupMember struct {
	OrderId     int `db:"OrderId"`
	GroupId     int `db:"GroupId"`
	UserId      int `db:"UserId"`
	ByUserId    int `db:"ByUserId"`
	GroupRole   int `db:"GroupRole"`
	CreatedTime int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.GroupMember {
	OrderId: 0,
	GroupId: 0,
	UserId: 0,
	ByUserId: 0,
	GroupRole: 0,
	CreatedTime: 0,
*/
// group_orderd_user 'GroupOrderdUser'.
type GroupOrderdUser struct {
	OrderId int `db:"OrderId"`
	GroupId int `db:"GroupId"`
	UserId  int `db:"UserId"`

	_exists, _deleted bool
}

/*
:= &x.GroupOrderdUser {
	OrderId: 0,
	GroupId: 0,
	UserId: 0,
*/
// group_pined_msg 'GroupPinedMsg'.
type GroupPinedMsg struct {
	MsgId int    `db:"MsgId"`
	MsgPb []byte `db:"MsgPb"`

	_exists, _deleted bool
}

/*
:= &x.GroupPinedMsg {
	MsgId: 0,
	MsgPb: []byte{},
*/
// file_msg 'FileMsg'.
type FileMsg struct {
	Id         int    `db:"Id"`
	AccessHash int    `db:"AccessHash"`
	FileType   int    `db:"FileType"`
	Width      int    `db:"Width"`
	Height     int    `db:"Height"`
	Extension  string `db:"Extension"`
	UserId     int    `db:"UserId"`
	DataThumb  []byte `db:"DataThumb"`
	Data       []byte `db:"Data"`

	_exists, _deleted bool
}

/*
:= &x.FileMsg {
	Id: 0,
	AccessHash: 0,
	FileType: 0,
	Width: 0,
	Height: 0,
	Extension: "",
	UserId: 0,
	DataThumb: []byte{},
	Data: []byte{},
*/
// file_post 'FilePost'.
type FilePost struct {
	Id         int    `db:"Id"`
	AccessHash int    `db:"AccessHash"`
	FileType   int    `db:"FileType"`
	Width      int    `db:"Width"`
	Height     int    `db:"Height"`
	Extension  string `db:"Extension"`
	UserId     int    `db:"UserId"`
	DataThumb  []byte `db:"DataThumb"`
	Data       []byte `db:"Data"`

	_exists, _deleted bool
}

/*
:= &x.FilePost {
	Id: 0,
	AccessHash: 0,
	FileType: 0,
	Width: 0,
	Height: 0,
	Extension: "",
	UserId: 0,
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
	ResharedId int `db:"ResharedId"`

	_exists, _deleted bool
}

/*
:= &x.HomeFanout {
	OrderId: 0,
	ForUserId: 0,
	PostId: 0,
	PostUserId: 0,
	ResharedId: 0,
*/
// suggested_top_posts 'SuggestedTopPosts'.
type SuggestedTopPosts struct {
	Id     int `db:"Id"`
	PostId int `db:"PostId"`

	_exists, _deleted bool
}

/*
:= &x.SuggestedTopPosts {
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
// push_chat 'PushChat'.
type PushChat struct {
	PushId            int    `db:"PushId"`
	ToUserId          int    `db:"ToUserId"`
	PushTypeId        int    `db:"PushTypeId"`
	RoomKey           string `db:"RoomKey"`
	ChatKey           string `db:"ChatKey"`
	Seq               int    `db:"Seq"`
	UnseenCount       int    `db:"UnseenCount"`
	FromHighMessageId int    `db:"FromHighMessageId"`
	ToLowMessageId    int    `db:"ToLowMessageId"`
	MessageId         int    `db:"MessageId"`
	MessageFileId     int    `db:"MessageFileId"`
	MessagePb         []byte `db:"MessagePb"`
	MessageJson       string `db:"MessageJson"`
	CreatedTime       int    `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.PushChat {
	PushId: 0,
	ToUserId: 0,
	PushTypeId: 0,
	RoomKey: "",
	ChatKey: "",
	Seq: 0,
	UnseenCount: 0,
	FromHighMessageId: 0,
	ToLowMessageId: 0,
	MessageId: 0,
	MessageFileId: 0,
	MessagePb: []byte{},
	MessageJson: "",
	CreatedTime: 0,
*/
// http_rpc_log 'HTTPRPCLog'.
type HTTPRPCLog struct {
	Id              int    `db:"Id"`
	Time            string `db:"Time"`
	MethodFull      string `db:"MethodFull"`
	MethodParent    string `db:"MethodParent"`
	UserId          int    `db:"UserId"`
	SessionId       string `db:"SessionId"`
	StatusCode      int    `db:"StatusCode"`
	InputSize       int    `db:"InputSize"`
	OutputSize      int    `db:"OutputSize"`
	ReqestJson      string `db:"ReqestJson"`
	ResponseJson    string `db:"ResponseJson"`
	ReqestParamJson string `db:"ReqestParamJson"`
	ResponseMsgJson string `db:"ResponseMsgJson"`

	_exists, _deleted bool
}

/*
:= &x.HTTPRPCLog {
	Id: 0,
	Time: "",
	MethodFull: "",
	MethodParent: "",
	UserId: 0,
	SessionId: "",
	StatusCode: 0,
	InputSize: 0,
	OutputSize: 0,
	ReqestJson: "",
	ResponseJson: "",
	ReqestParamJson: "",
	ResponseMsgJson: "",
*/
// metric_log 'MetricLog'.
type MetricLog struct {
	Id           int    `db:"Id"`
	InstanceId   int    `db:"InstanceId"`
	StartFrom    string `db:"StartFrom"`
	EndTo        string `db:"EndTo"`
	StartTime    int    `db:"StartTime"`
	Duration     string `db:"Duration"`
	MetericsJson string `db:"MetericsJson"`

	_exists, _deleted bool
}

/*
:= &x.MetricLog {
	Id: 0,
	InstanceId: 0,
	StartFrom: "",
	EndTo: "",
	StartTime: 0,
	Duration: "",
	MetericsJson: "",
*/
// xfile_service_info_log 'XfileServiceInfoLog'.
type XfileServiceInfoLog struct {
	Id          int    `db:"Id"`
	InstanceId  int    `db:"InstanceId"`
	Url         string `db:"Url"`
	CreatedTime string `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.XfileServiceInfoLog {
	Id: 0,
	InstanceId: 0,
	Url: "",
	CreatedTime: "",
*/
// xfile_service_metric_log 'XfileServiceMetricLog'.
type XfileServiceMetricLog struct {
	Id         int    `db:"Id"`
	InstanceId int    `db:"InstanceId"`
	MetricJson string `db:"MetricJson"`

	_exists, _deleted bool
}

/*
:= &x.XfileServiceMetricLog {
	Id: 0,
	InstanceId: 0,
	MetricJson: "",
*/
// xfile_service_request_log 'XfileServiceRequestLog'.
type XfileServiceRequestLog struct {
	Id          int    `db:"Id"`
	LocalSeq    int    `db:"LocalSeq"`
	InstanceId  int    `db:"InstanceId"`
	Url         string `db:"Url"`
	HttpCode    int    `db:"HttpCode"`
	CreatedTime string `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.XfileServiceRequestLog {
	Id: 0,
	LocalSeq: 0,
	InstanceId: 0,
	Url: "",
	HttpCode: 0,
	CreatedTime: "",
*/
// invalidate_cache 'InvalidateCache'.
type InvalidateCache struct {
	OrderId  int    `db:"OrderId"`
	CacheKey string `db:"CacheKey"`

	_exists, _deleted bool
}

/*
:= &x.InvalidateCache {
	OrderId: 0,
	CacheKey: "",
*/

///////////////// Skip Loging Tables ////////////////
type LogTableSql struct {
	Action                 bool
	Blocked                bool
	Comment                bool
	CommentDeleted         bool
	Event                  bool
	Followed               bool
	Likes                  bool
	Notify                 bool
	NotifyRemoved          bool
	PhoneContacts          bool
	Post                   bool
	PostCount              bool
	PostDeleted            bool
	PostKeys               bool
	PostLink               bool
	PostMedia              bool
	PostPromoted           bool
	PostReshared           bool
	ProfileAll             bool
	ProfileMedia           bool
	ProfileMentioned       bool
	Session                bool
	SettingNotifications   bool
	Sms                    bool
	Tag                    bool
	TagPost                bool
	TriggerLog             bool
	User                   bool
	UserRelation           bool
	Chat                   bool
	ChatDeleted            bool
	ChatLastMessage        bool
	ChatUserVersion        bool
	Group                  bool
	GroupMember            bool
	GroupOrderdUser        bool
	GroupPinedMsg          bool
	FileMsg                bool
	FilePost               bool
	ActionFanout           bool
	HomeFanout             bool
	SuggestedTopPosts      bool
	SuggestedUser          bool
	PushChat               bool
	HTTPRPCLog             bool
	MetricLog              bool
	XfileServiceInfoLog    bool
	XfileServiceMetricLog  bool
	XfileServiceRequestLog bool
	InvalidateCache        bool
}

var LogTableSqlReq = LogTableSql{

	Action:                 true,
	Blocked:                true,
	Comment:                true,
	CommentDeleted:         true,
	Event:                  true,
	Followed:               true,
	Likes:                  true,
	Notify:                 true,
	NotifyRemoved:          true,
	PhoneContacts:          true,
	Post:                   true,
	PostCount:              true,
	PostDeleted:            true,
	PostKeys:               true,
	PostLink:               true,
	PostMedia:              true,
	PostPromoted:           true,
	PostReshared:           true,
	ProfileAll:             true,
	ProfileMedia:           true,
	ProfileMentioned:       true,
	Session:                true,
	SettingNotifications:   true,
	Sms:                    true,
	Tag:                    true,
	TagPost:                true,
	TriggerLog:             true,
	User:                   true,
	UserRelation:           true,
	Chat:                   true,
	ChatDeleted:            true,
	ChatLastMessage:        true,
	ChatUserVersion:        true,
	Group:                  true,
	GroupMember:            true,
	GroupOrderdUser:        true,
	GroupPinedMsg:          true,
	FileMsg:                true,
	FilePost:               true,
	ActionFanout:           true,
	HomeFanout:             true,
	SuggestedTopPosts:      true,
	SuggestedUser:          true,
	PushChat:               true,
	HTTPRPCLog:             true,
	MetricLog:              true,
	XfileServiceInfoLog:    true,
	XfileServiceMetricLog:  true,
	XfileServiceRequestLog: true,
	InvalidateCache:        true,
}
