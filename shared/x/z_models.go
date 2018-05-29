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
	ActionId     int    `db:"ActionId"`
	Murmur64Hash int    `db:"Murmur64Hash"`
	ChatKey      string `db:"ChatKey"`
	MessageId    int    `db:"MessageId"`
	ReSharedId   int    `db:"ReSharedId"`

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
	ChatKey: "",
	MessageId: 0,
	ReSharedId: 0,
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
	Id        int    `db:"Id"`
	UserId    int    `db:"UserId"`
	ClientId  int    `db:"ClientId"`
	Phone     string `db:"Phone"`
	FirstName string `db:"FirstName"`
	LastName  string `db:"LastName"`

	_exists, _deleted bool
}

/*
:= &x.PhoneContact {
	Id: 0,
	UserId: 0,
	ClientId: 0,
	Phone: "",
	FirstName: "",
	LastName: "",
*/
// post 'Post'.
type Post struct {
	PostId           int    `db:"PostId"`
	UserId           int    `db:"UserId"`
	PostTypeEnum     int    `db:"PostTypeEnum"`
	PostCategoryEnum int    `db:"PostCategoryEnum"`
	MediaId          int    `db:"MediaId"`
	PostKey          string `db:"PostKey"`
	Text             string `db:"Text"`
	RichText         string `db:"RichText"`
	MediaCount       int    `db:"MediaCount"`
	SharedTo         int    `db:"SharedTo"`
	DisableComment   int    `db:"DisableComment"`
	Source           int    `db:"Source"`
	HasTag           int    `db:"HasTag"`
	Seq              int    `db:"Seq"`
	CommentsCount    int    `db:"CommentsCount"`
	LikesCount       int    `db:"LikesCount"`
	ViewsCount       int    `db:"ViewsCount"`
	EditedTime       int    `db:"EditedTime"`
	CreatedTime      int    `db:"CreatedTime"`
	ReSharedPostId   int    `db:"ReSharedPostId"`

	_exists, _deleted bool
}

/*
:= &x.Post {
	PostId: 0,
	UserId: 0,
	PostTypeEnum: 0,
	PostCategoryEnum: 0,
	MediaId: 0,
	PostKey: "",
	Text: "",
	RichText: "",
	MediaCount: 0,
	SharedTo: 0,
	DisableComment: 0,
	Source: 0,
	HasTag: 0,
	Seq: 0,
	CommentsCount: 0,
	LikesCount: 0,
	ViewsCount: 0,
	EditedTime: 0,
	CreatedTime: 0,
	ReSharedPostId: 0,
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
// post_mentioned 'PostMentioned'.
type PostMentioned struct {
	MentionedId      int `db:"MentionedId"`
	ForUserId        int `db:"ForUserId"`
	PostId           int `db:"PostId"`
	PostUserId       int `db:"PostUserId"`
	PostTypeEnum     int `db:"PostTypeEnum"`
	PostCategoryEnum int `db:"PostCategoryEnum"`
	CreatedTime      int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.PostMentioned {
	MentionedId: 0,
	ForUserId: 0,
	PostId: 0,
	PostUserId: 0,
	PostTypeEnum: 0,
	PostCategoryEnum: 0,
	CreatedTime: 0,
*/
// post_reshared 'PostReshared'.
type PostReshared struct {
	ResharedId       int `db:"ResharedId"`
	ByUserId         int `db:"ByUserId"`
	PostId           int `db:"PostId"`
	PostUserId       int `db:"PostUserId"`
	PostTypeEnum     int `db:"PostTypeEnum"`
	PostCategoryEnum int `db:"PostCategoryEnum"`
	CreatedTime      int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.PostReshared {
	ResharedId: 0,
	ByUserId: 0,
	PostId: 0,
	PostUserId: 0,
	PostTypeEnum: 0,
	PostCategoryEnum: 0,
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
// tag_post 'TagPost'.
type TagPost struct {
	Id               int `db:"Id"`
	TagId            int `db:"TagId"`
	PostId           int `db:"PostId"`
	PostTypeEnum     int `db:"PostTypeEnum"`
	PostCategoryEnum int `db:"PostCategoryEnum"`
	CreatedTime      int `db:"CreatedTime"`

	_exists, _deleted bool
}

/*
:= &x.TagPost {
	Id: 0,
	TagId: 0,
	PostId: 0,
	PostTypeEnum: 0,
	PostCategoryEnum: 0,
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
	UserId         int    `db:"UserId"`
	UserName       string `db:"UserName"`
	UserNameLower  string `db:"UserNameLower"` //nojson
	FirstName      string `db:"FirstName"`
	LastName       string `db:"LastName"`
	IsVerified     int    `db:"IsVerified"`
	AvatarId       int    `db:"AvatarId"`
	ProfilePrivacy int    `db:"ProfilePrivacy"`
	OnlinePrivacy  int    `db:"OnlinePrivacy"`
	Phone          int    `db:"Phone"`
	Email          string `db:"Email"`
	About          string `db:"About"`
	PasswordHash   string `db:"PasswordHash"`
	PasswordSalt   string `db:"PasswordSalt"`
	PostSeq        int    `db:"PostSeq"`
	FollowersCount int    `db:"FollowersCount"`
	FollowingCount int    `db:"FollowingCount"`
	PostsCount     int    `db:"PostsCount"`
	MediaCount     int    `db:"MediaCount"`
	PhotoCount     int    `db:"PhotoCount"`
	VideoCount     int    `db:"VideoCount"`
	GifCount       int    `db:"GifCount"`
	AudioCount     int    `db:"AudioCount"`
	VoiceCount     int    `db:"VoiceCount"`
	FileCount      int    `db:"FileCount"`
	LinkCount      int    `db:"LinkCount"`
	BoardCount     int    `db:"BoardCount"`
	PinedCount     int    `db:"PinedCount"`
	LikesCount     int    `db:"LikesCount"`
	ResharedCount  int    `db:"ResharedCount"`
	LastPostTime   int    `db:"LastPostTime"`
	CreatedTime    int    `db:"CreatedTime"`
	VersionTime    int    `db:"VersionTime"`
	IsDeleted      int    `db:"IsDeleted"`
	IsBanned       int    `db:"IsBanned"`

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
	ProfilePrivacy: 0,
	OnlinePrivacy: 0,
	Phone: 0,
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
	ChatId           int    `db:"ChatId"`
	ChatKey          string `db:"ChatKey"`
	RoomKey          string `db:"RoomKey"`
	RoomType         int    `db:"RoomType"`
	UserId           int    `db:"UserId"`
	PeerUserId       int    `db:"PeerUserId"`
	GroupId          int    `db:"GroupId"`
	HashTagId        int    `db:"HashTagId"`
	StartedByMe      int    `db:"StartedByMe"`
	Title            string `db:"Title"`
	PinTime          int    `db:"PinTime"`
	FromMsgId        int    `db:"FromMsgId"`
	Seq              int    `db:"Seq"`
	UnseenCount      int    `db:"UnseenCount"`
	LastMsgId        int    `db:"LastMsgId"`
	LastMsgStatus    int    `db:"LastMsgStatus"`
	SeenSeq          int    `db:"SeenSeq"`
	SeenMsgId        int    `db:"SeenMsgId"`
	LastMsgIdRecived int    `db:"LastMsgIdRecived"`
	Left             int    `db:"Left"`
	Creator          int    `db:"Creator"`
	Kicked           int    `db:"Kicked"`
	Admin            int    `db:"Admin"`
	Deactivated      int    `db:"Deactivated"`
	VersionTime      int    `db:"VersionTime"`
	OrderTime        int    `db:"OrderTime"`
	CreatedTime      int    `db:"CreatedTime"`
	DraftText        string `db:"DraftText"`
	DratReplyToMsgId int    `db:"DratReplyToMsgId"`

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
	StartedByMe: 0,
	Title: "",
	PinTime: 0,
	FromMsgId: 0,
	Seq: 0,
	UnseenCount: 0,
	LastMsgId: 0,
	LastMsgStatus: 0,
	SeenSeq: 0,
	SeenMsgId: 0,
	LastMsgIdRecived: 0,
	Left: 0,
	Creator: 0,
	Kicked: 0,
	Admin: 0,
	Deactivated: 0,
	VersionTime: 0,
	OrderTime: 0,
	CreatedTime: 0,
	DraftText: "",
	DratReplyToMsgId: 0,
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
// chat_version_order 'ChatVersionOrder'.
type ChatVersionOrder struct {
	VersionTime int `db:"VersionTime"`
	UserId      int `db:"UserId"`
	ChatId      int `db:"ChatId"`
	OrderTime   int `db:"OrderTime"`

	_exists, _deleted bool
}

/*
:= &x.ChatVersionOrder {
	VersionTime: 0,
	UserId: 0,
	ChatId: 0,
	OrderTime: 0,
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
	InviteLink      string `db:"InviteLink"`
	MembersCount    int    `db:"MembersCount"`
	SortTime        int    `db:"SortTime"`
	CreatedTime     int    `db:"CreatedTime"`
	IsMute          int    `db:"IsMute"`

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
	InviteLink: "",
	MembersCount: 0,
	SortTime: 0,
	CreatedTime: 0,
	IsMute: 0,
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
	Comment                bool
	CommentDeleted         bool
	Event                  bool
	Like                   bool
	Notify                 bool
	NotifyRemoved          bool
	PhoneContact           bool
	Post                   bool
	PostCount              bool
	PostDeleted            bool
	PostKey                bool
	PostLink               bool
	PostMedia              bool
	PostMentioned          bool
	PostReshared           bool
	Session                bool
	SettingClient          bool
	SettingNotification    bool
	Tag                    bool
	TagPost                bool
	TriggerLog             bool
	User                   bool
	UserRelation           bool
	Chat                   bool
	ChatLastMessage        bool
	ChatVersionOrder       bool
	Group                  bool
	GroupMember            bool
	GroupOrderdUser        bool
	GroupPinedMsg          bool
	FileMsg                bool
	FilePost               bool
	ActionFanout           bool
	HomeFanout             bool
	SuggestedTopPost       bool
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
	Comment:                true,
	CommentDeleted:         true,
	Event:                  true,
	Like:                   true,
	Notify:                 true,
	NotifyRemoved:          true,
	PhoneContact:           true,
	Post:                   true,
	PostCount:              true,
	PostDeleted:            true,
	PostKey:                true,
	PostLink:               true,
	PostMedia:              true,
	PostMentioned:          true,
	PostReshared:           true,
	Session:                true,
	SettingClient:          true,
	SettingNotification:    true,
	Tag:                    true,
	TagPost:                true,
	TriggerLog:             true,
	User:                   true,
	UserRelation:           true,
	Chat:                   true,
	ChatLastMessage:        true,
	ChatVersionOrder:       true,
	Group:                  true,
	GroupMember:            true,
	GroupOrderdUser:        true,
	GroupPinedMsg:          true,
	FileMsg:                true,
	FilePost:               true,
	ActionFanout:           true,
	HomeFanout:             true,
	SuggestedTopPost:       true,
	SuggestedUser:          true,
	PushChat:               true,
	HTTPRPCLog:             true,
	MetricLog:              true,
	XfileServiceInfoLog:    true,
	XfileServiceMetricLog:  true,
	XfileServiceRequestLog: true,
	InvalidateCache:        true,
}
