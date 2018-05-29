package com.mardomsara.social.json;

public class J {
	public static class Action {//oridnal: 0
		public int ActionId; //ActionId
		public int ActorUserId; //ActorUserId
		public int ActionTypeEnum; //ActionTypeEnum
		public int PeerUserId; //PeerUserId
		public int PostId; //PostId
		public int CommentId; //CommentId
		public int Murmur64Hash; //Murmur64Hash
		public int CreatedTime; //CreatedTime
	}

	public static class Comment {//oridnal: 1
		public int CommentId; //CommentId
		public int UserId; //UserId
		public int PostId; //PostId
		public String Text; //Text
		public int LikesCount; //LikesCount
		public int CreatedTime; //CreatedTime
	}

	public static class CommentDeleted {//oridnal: 2
		public int CommentId; //CommentId
		public int UserId; //UserId
	}

	public static class Event {//oridnal: 3
		public int EventId; //EventId
		public int EventType; //EventType
		public int ByUserId; //ByUserId
		public int PeerUserId; //PeerUserId
		public int PostId; //PostId
		public int CommentId; //CommentId
		public int ActionId; //ActionId
		public int Murmur64Hash; //Murmur64Hash
		public String ChatKey; //ChatKey
		public int MessageId; //MessageId
		public int ReSharedId; //ReSharedId
	}

	public static class Like {//oridnal: 4
		public int Id; //Id
		public int PostId; //PostId
		public int PostTypeEnum; //PostTypeEnum
		public int UserId; //UserId
		public int LikeEnum; //LikeEnum
		public int CreatedTime; //CreatedTime
	}

	public static class Notify {//oridnal: 5
		public int NotifyId; //NotifyId
		public int ForUserId; //ForUserId
		public int ActorUserId; //ActorUserId
		public int NotifyTypeEnum; //NotifyTypeEnum
		public int PostId; //PostId
		public int CommentId; //CommentId
		public int PeerUserId; //PeerUserId
		public int Murmur64Hash; //Murmur64Hash
		public int SeenStatus; //SeenStatus
		public int CreatedTime; //CreatedTime
	}

	public static class NotifyRemoved {//oridnal: 6
		public int Murmur64Hash; //Murmur64Hash
		public int ForUserId; //ForUserId
		public int Id; //Id
	}

	public static class PhoneContact {//oridnal: 7
		public int Id; //Id
		public int UserId; //UserId
		public int ClientId; //ClientId
		public String Phone; //Phone
		public String FirstName; //FirstName
		public String LastName; //LastName
	}

	public static class Post {//oridnal: 8
		public int PostId; //PostId
		public int UserId; //UserId
		public int PostTypeEnum; //PostTypeEnum
		public int PostCategoryEnum; //PostCategoryEnum
		public int MediaId; //MediaId
		public String PostKey; //PostKey
		public String Text; //Text
		public String RichText; //RichText
		public int MediaCount; //MediaCount
		public int SharedTo; //SharedTo
		public int DisableComment; //DisableComment
		public int Source; //Source
		public int HasTag; //HasTag
		public int Seq; //Seq
		public int CommentsCount; //CommentsCount
		public int LikesCount; //LikesCount
		public int ViewsCount; //ViewsCount
		public int EditedTime; //EditedTime
		public int CreatedTime; //CreatedTime
		public int ReSharedPostId; //ReSharedPostId
	}

	public static class PostCount {//oridnal: 9
		public int PostId; //PostId
		public int CommentsCount; //CommentsCount
		public int LikesCount; //LikesCount
		public int ViewsCount; //ViewsCount
		public int ReSharedCount; //ReSharedCount
		public int ChatSharedCount; //ChatSharedCount
	}

	public static class PostDeleted {//oridnal: 10
		public int PostId; //PostId
		public int UserId; //UserId
	}

	public static class PostKey {//oridnal: 11
		public int Id; //Id
		public String PostKeyStr; //PostKeyStr
		public int Used; //Used
	}

	public static class PostLink {//oridnal: 12
		public int LinkId; //LinkId
		public String LinkUrl; //LinkUrl
	}

	public static class PostMedia {//oridnal: 13
		public int MediaId; //MediaId
		public int UserId; //UserId
		public int PostId; //PostId
		public int AlbumId; //AlbumId
		public int MediaTypeEnum; //MediaTypeEnum
		public int Width; //Width
		public int Height; //Height
		public int Size; //Size
		public int Duration; //Duration
		public String Extension; //Extension
		public String Md5Hash; //Md5Hash
		public String Color; //Color
		public int CreatedTime; //CreatedTime
		public int ViewCount; //ViewCount
		public String Extra; //Extra
	}

	public static class PostMentioned {//oridnal: 14
		public int MentionedId; //MentionedId
		public int ForUserId; //ForUserId
		public int PostId; //PostId
		public int PostUserId; //PostUserId
		public int PostTypeEnum; //PostTypeEnum
		public int PostCategoryEnum; //PostCategoryEnum
		public int CreatedTime; //CreatedTime
	}

	public static class PostReshared {//oridnal: 15
		public int ResharedId; //ResharedId
		public int ByUserId; //ByUserId
		public int PostId; //PostId
		public int PostUserId; //PostUserId
		public int PostTypeEnum; //PostTypeEnum
		public int PostCategoryEnum; //PostCategoryEnum
		public int CreatedTime; //CreatedTime
	}

	public static class Session {//oridnal: 16
		public int Id; //Id
		public String SessionUuid; //SessionUuid
		public int UserId; //UserId
		public String LastIpAddress; //LastIpAddress
		public int AppVersion; //AppVersion
		public int ActiveTime; //ActiveTime
		public int CreatedTime; //CreatedTime
	}

	public static class SettingClient {//oridnal: 17
		public int UserId; //UserId
		public int AutoDownloadWifiVoice; //AutoDownloadWifiVoice
		public int AutoDownloadWifiImage; //AutoDownloadWifiImage
		public int AutoDownloadWifiVideo; //AutoDownloadWifiVideo
		public int AutoDownloadWifiFile; //AutoDownloadWifiFile
		public int AutoDownloadWifiMusic; //AutoDownloadWifiMusic
		public int AutoDownloadWifiGif; //AutoDownloadWifiGif
		public int AutoDownloadCellularVoice; //AutoDownloadCellularVoice
		public int AutoDownloadCellularImage; //AutoDownloadCellularImage
		public int AutoDownloadCellularVideo; //AutoDownloadCellularVideo
		public int AutoDownloadCellularFile; //AutoDownloadCellularFile
		public int AutoDownloadCellularMusic; //AutoDownloadCellularMusic
		public int AutoDownloadCellularGif; //AutoDownloadCellularGif
		public int AutoDownloadRoamingVoice; //AutoDownloadRoamingVoice
		public int AutoDownloadRoamingImage; //AutoDownloadRoamingImage
		public int AutoDownloadRoamingVideo; //AutoDownloadRoamingVideo
		public int AutoDownloadRoamingFile; //AutoDownloadRoamingFile
		public int AutoDownloadRoamingMusic; //AutoDownloadRoamingMusic
		public int AutoDownloadRoamingGif; //AutoDownloadRoamingGif
		public int SaveToGallery; //SaveToGallery
	}

	public static class SettingNotification {//oridnal: 18
		public int UserId; //UserId
		public int SocialLedOn; //SocialLedOn
		public String SocialLedColor; //SocialLedColor
		public int ReqestToFollowYou; //ReqestToFollowYou
		public int FollowedYou; //FollowedYou
		public int AccptedYourFollowRequest; //AccptedYourFollowRequest
		public int YourPostLiked; //YourPostLiked
		public int YourPostCommented; //YourPostCommented
		public int MenthenedYouInPost; //MenthenedYouInPost
		public int MenthenedYouInComment; //MenthenedYouInComment
		public int YourContactsJoined; //YourContactsJoined
		public int DirectMessage; //DirectMessage
		public int DirectAlert; //DirectAlert
		public int DirectPerview; //DirectPerview
		public int DirectLedOn; //DirectLedOn
		public int DirectLedColor; //DirectLedColor
		public int DirectVibrate; //DirectVibrate
		public int DirectPopup; //DirectPopup
		public int DirectSound; //DirectSound
		public int DirectPriority; //DirectPriority
	}

	public static class Tag {//oridnal: 19
		public int TagId; //TagId
		public String Name; //Name
		public int Count; //Count
		public int TagStatusEnum; //TagStatusEnum
		public int CreatedTime; //CreatedTime
	}

	public static class TagPost {//oridnal: 20
		public int Id; //Id
		public int TagId; //TagId
		public int PostId; //PostId
		public int PostTypeEnum; //PostTypeEnum
		public int PostCategoryEnum; //PostCategoryEnum
		public int CreatedTime; //CreatedTime
	}

	public static class TriggerLog {//oridnal: 21
		public int Id; //Id
		public String ModelName; //ModelName
		public String ChangeType; //ChangeType
		public int TargetId; //TargetId
		public String TargetStr; //TargetStr
		public int CreatedSe; //CreatedSe
	}

	public static class User {//oridnal: 22
		public int UserId; //UserId
		public String UserName; //UserName
		public String UserNameLower; //UserNameLower
		public String FirstName; //FirstName
		public String LastName; //LastName
		public int IsVerified; //IsVerified
		public int AvatarId; //AvatarId
		public int ProfilePrivacy; //ProfilePrivacy
		public int OnlinePrivacy; //OnlinePrivacy
		public int Phone; //Phone
		public String Email; //Email
		public String About; //About
		public String PasswordHash; //PasswordHash
		public String PasswordSalt; //PasswordSalt
		public int PostSeq; //PostSeq
		public int FollowersCount; //FollowersCount
		public int FollowingCount; //FollowingCount
		public int PostsCount; //PostsCount
		public int MediaCount; //MediaCount
		public int PhotoCount; //PhotoCount
		public int VideoCount; //VideoCount
		public int GifCount; //GifCount
		public int AudioCount; //AudioCount
		public int VoiceCount; //VoiceCount
		public int FileCount; //FileCount
		public int LinkCount; //LinkCount
		public int BoardCount; //BoardCount
		public int PinedCount; //PinedCount
		public int LikesCount; //LikesCount
		public int ResharedCount; //ResharedCount
		public int LastPostTime; //LastPostTime
		public int CreatedTime; //CreatedTime
		public int VersionTime; //VersionTime
		public int IsDeleted; //IsDeleted
		public int IsBanned; //IsBanned
	}

	public static class UserRelation {//oridnal: 23
		public int RelNanoId; //RelNanoId
		public int UserId; //UserId
		public int PeerUserId; //PeerUserId
		public int Follwing; //Follwing
		public int Followed; //Followed
		public int InContacts; //InContacts
		public int MutualContact; //MutualContact
		public int IsFavorite; //IsFavorite
		public int Notify; //Notify
	}

	public static class Chat {//oridnal: 24
		public int ChatId; //ChatId
		public String ChatKey; //ChatKey
		public String RoomKey; //RoomKey
		public int RoomType; //RoomType
		public int UserId; //UserId
		public int PeerUserId; //PeerUserId
		public int GroupId; //GroupId
		public int HashTagId; //HashTagId
		public int StartedByMe; //StartedByMe
		public String Title; //Title
		public int PinTime; //PinTime
		public int FromMsgId; //FromMsgId
		public int Seq; //Seq
		public int UnseenCount; //UnseenCount
		public int LastMsgId; //LastMsgId
		public int LastMsgStatus; //LastMsgStatus
		public int SeenSeq; //SeenSeq
		public int SeenMsgId; //SeenMsgId
		public int LastMsgIdRecived; //LastMsgIdRecived
		public int Left; //Left
		public int Creator; //Creator
		public int Kicked; //Kicked
		public int Admin; //Admin
		public int Deactivated; //Deactivated
		public int VersionTime; //VersionTime
		public int OrderTime; //OrderTime
		public int CreatedTime; //CreatedTime
		public String DraftText; //DraftText
		public int DratReplyToMsgId; //DratReplyToMsgId
	}

	public static class ChatLastMessage {//oridnal: 25
		public String ChatIdGroupId; //ChatIdGroupId
		public UNKNOWN LastMsgPb; //LastMsgPb
	}

	public static class ChatVersionOrder {//oridnal: 26
		public int VersionTime; //VersionTime
		public int UserId; //UserId
		public int ChatId; //ChatId
		public int OrderTime; //OrderTime
	}

	public static class Group {//oridnal: 27
		public int GroupId; //GroupId
		public String GroupKey; //GroupKey
		public String GroupName; //GroupName
		public String UserName; //UserName
		public int IsSuperGroup; //IsSuperGroup
		public int HashTagId; //HashTagId
		public int CreatorUserId; //CreatorUserId
		public int GroupPrivacy; //GroupPrivacy
		public int HistoryViewAble; //HistoryViewAble
		public int Seq; //Seq
		public int LastMsgId; //LastMsgId
		public int PinedMsgId; //PinedMsgId
		public int AvatarRefId; //AvatarRefId
		public int AvatarCount; //AvatarCount
		public String About; //About
		public String InviteLink; //InviteLink
		public int MembersCount; //MembersCount
		public int SortTime; //SortTime
		public int CreatedTime; //CreatedTime
		public int IsMute; //IsMute
	}

	public static class GroupMember {//oridnal: 28
		public int OrderId; //OrderId
		public int GroupId; //GroupId
		public int UserId; //UserId
		public int ByUserId; //ByUserId
		public int GroupRole; //GroupRole
		public int CreatedTime; //CreatedTime
	}

	public static class GroupOrderdUser {//oridnal: 29
		public int OrderId; //OrderId
		public int GroupId; //GroupId
		public int UserId; //UserId
	}

	public static class GroupPinedMsg {//oridnal: 30
		public int MsgId; //MsgId
		public UNKNOWN MsgPb; //MsgPb
	}

	public static class FileMsg {//oridnal: 31
		public int Id; //Id
		public int AccessHash; //AccessHash
		public int FileType; //FileType
		public int Width; //Width
		public int Height; //Height
		public String Extension; //Extension
		public int UserId; //UserId
		public UNKNOWN DataThumb; //DataThumb
		public UNKNOWN Data; //Data
	}

	public static class FilePost {//oridnal: 32
		public int Id; //Id
		public int AccessHash; //AccessHash
		public int FileType; //FileType
		public int Width; //Width
		public int Height; //Height
		public String Extension; //Extension
		public int UserId; //UserId
		public UNKNOWN DataThumb; //DataThumb
		public UNKNOWN Data; //Data
	}

	public static class ActionFanout {//oridnal: 33
		public int OrderId; //OrderId
		public int ForUserId; //ForUserId
		public int ActionId; //ActionId
		public int ActorUserId; //ActorUserId
	}

	public static class HomeFanout {//oridnal: 34
		public int OrderId; //OrderId
		public int ForUserId; //ForUserId
		public int PostId; //PostId
		public int PostUserId; //PostUserId
		public int ResharedId; //ResharedId
	}

	public static class SuggestedTopPost {//oridnal: 35
		public int Id; //Id
		public int PostId; //PostId
	}

	public static class SuggestedUser {//oridnal: 36
		public int Id; //Id
		public int UserId; //UserId
		public int TargetId; //TargetId
		public float Weight; //Weight
		public int CreatedTime; //CreatedTime
	}

	public static class PushChat {//oridnal: 37
		public int PushId; //PushId
		public int ToUserId; //ToUserId
		public int PushTypeId; //PushTypeId
		public String RoomKey; //RoomKey
		public String ChatKey; //ChatKey
		public int Seq; //Seq
		public int UnseenCount; //UnseenCount
		public int FromHighMessageId; //FromHighMessageId
		public int ToLowMessageId; //ToLowMessageId
		public int MessageId; //MessageId
		public int MessageFileId; //MessageFileId
		public UNKNOWN MessagePb; //MessagePb
		public String MessageJson; //MessageJson
		public int CreatedTime; //CreatedTime
	}

	public static class HTTPRPCLog {//oridnal: 38
		public int Id; //Id
		public String Time; //Time
		public String MethodFull; //MethodFull
		public String MethodParent; //MethodParent
		public int UserId; //UserId
		public String SessionId; //SessionId
		public int StatusCode; //StatusCode
		public int InputSize; //InputSize
		public int OutputSize; //OutputSize
		public String ReqestJson; //ReqestJson
		public String ResponseJson; //ResponseJson
		public String ReqestParamJson; //ReqestParamJson
		public String ResponseMsgJson; //ResponseMsgJson
	}

	public static class MetricLog {//oridnal: 39
		public int Id; //Id
		public int InstanceId; //InstanceId
		public String StartFrom; //StartFrom
		public String EndTo; //EndTo
		public int StartTime; //StartTime
		public String Duration; //Duration
		public String MetericsJson; //MetericsJson
	}

	public static class XfileServiceInfoLog {//oridnal: 40
		public int Id; //Id
		public int InstanceId; //InstanceId
		public String Url; //Url
		public String CreatedTime; //CreatedTime
	}

	public static class XfileServiceMetricLog {//oridnal: 41
		public int Id; //Id
		public int InstanceId; //InstanceId
		public String MetricJson; //MetricJson
	}

	public static class XfileServiceRequestLog {//oridnal: 42
		public int Id; //Id
		public int LocalSeq; //LocalSeq
		public int InstanceId; //InstanceId
		public String Url; //Url
		public int HttpCode; //HttpCode
		public String CreatedTime; //CreatedTime
	}

	public static class InvalidateCache {//oridnal: 43
		public int OrderId; //OrderId
		public String CacheKey; //CacheKey
	}

}