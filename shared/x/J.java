package com.mardomsara.social.json;

public class J {
	public static class Action {//oridnal: 0
		public int ActionId; //ActionId
		public int ActorUserId; //ActorUserId
		public int ActionType; //ActionType
		public int PeerUserId; //PeerUserId
		public int PostId; //PostId
		public int CommentId; //CommentId
		public int Murmur64Hash; //Murmur64Hash
		public int CreatedTime; //CreatedTime
	}

	public static class Blocked {//oridnal: 1
		public int Id; //Id
		public int UserId; //UserId
		public int BlockedUserId; //BlockedUserId
		public int CreatedTime; //CreatedTime
	}

	public static class Comment {//oridnal: 2
		public int CommentId; //CommentId
		public int UserId; //UserId
		public int PostId; //PostId
		public String Text; //Text
		public int LikesCount; //LikesCount
		public int IsEdited; //IsEdited
		public int CreatedTime; //CreatedTime
	}

	public static class CommentDeleted {//oridnal: 3
		public int CommentId; //CommentId
		public int UserId; //UserId
	}

	public static class Event {//oridnal: 4
		public int EventId; //EventId
		public int EventType; //EventType
		public int ByUserId; //ByUserId
		public int PeerUserId; //PeerUserId
		public int PostId; //PostId
		public int CommentId; //CommentId
		public int HashTagId; //HashTagId
		public int GroupId; //GroupId
		public int ActionId; //ActionId
		public int ChatId; //ChatId
		public String ChatKey; //ChatKey
		public int MessageId; //MessageId
		public int ReSharedId; //ReSharedId
		public int Murmur64Hash; //Murmur64Hash
	}

	public static class Followed {//oridnal: 5
		public int Id; //Id
		public int UserId; //UserId
		public int FollowedUserId; //FollowedUserId
		public int CreatedTime; //CreatedTime
	}

	public static class Likes {//oridnal: 6
		public int Id; //Id
		public int PostId; //PostId
		public int UserId; //UserId
		public int PostType; //PostType
		public int CreatedTime; //CreatedTime
	}

	public static class Notify {//oridnal: 7
		public int NotifyId; //NotifyId
		public int ForUserId; //ForUserId
		public int ActorUserId; //ActorUserId
		public int NotifyType; //NotifyType
		public int PostId; //PostId
		public int CommentId; //CommentId
		public int PeerUserId; //PeerUserId
		public int Murmur64Hash; //Murmur64Hash
		public int SeenStatus; //SeenStatus
		public int CreatedTime; //CreatedTime
	}

	public static class NotifyRemoved {//oridnal: 8
		public int Murmur64Hash; //Murmur64Hash
		public int ForUserId; //ForUserId
		public int Id; //Id
	}

	public static class PhoneContacts {//oridnal: 9
		public int Id; //Id
		public int UserId; //UserId
		public int ClientId; //ClientId
		public String Phone; //Phone
		public String FirstName; //FirstName
		public String LastName; //LastName
	}

	public static class Post {//oridnal: 10
		public int PostId; //PostId
		public int UserId; //UserId
		public int PostType; //PostType
		public int MediaId; //MediaId
		public int FileRefId; //FileRefId
		public String PostKey; //PostKey
		public String Text; //Text
		public String RichText; //RichText
		public int MediaCount; //MediaCount
		public int SharedTo; //SharedTo
		public int DisableComment; //DisableComment
		public int Via; //Via
		public int Seq; //Seq
		public int CommentsCount; //CommentsCount
		public int LikesCount; //LikesCount
		public int ViewsCount; //ViewsCount
		public int EditedTime; //EditedTime
		public int CreatedTime; //CreatedTime
	}

	public static class PostCount {//oridnal: 11
		public int PostId; //PostId
		public int CommentsCount; //CommentsCount
		public int LikesCount; //LikesCount
		public int ViewsCount; //ViewsCount
		public int ReSharedCount; //ReSharedCount
		public int ChatSharedCount; //ChatSharedCount
	}

	public static class PostDeleted {//oridnal: 12
		public int PostId; //PostId
		public int UserId; //UserId
	}

	public static class PostKeys {//oridnal: 13
		public int Id; //Id
		public String PostKeyStr; //PostKeyStr
		public int Used; //Used
		public int RandShard; //RandShard
	}

	public static class PostLink {//oridnal: 14
		public int LinkId; //LinkId
		public String LinkUrl; //LinkUrl
	}

	public static class PostMedia {//oridnal: 15
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

	public static class PostPromoted {//oridnal: 16
		public int PromoteId; //PromoteId
		public int PostId; //PostId
		public int ByUserId; //ByUserId
		public int PostUserId; //PostUserId
		public String BazzarUuid; //BazzarUuid
		public String Package; //Package
		public int EndTime; //EndTime
		public int CreatedTime; //CreatedTime
	}

	public static class PostReshared {//oridnal: 17
		public int ResharedId; //ResharedId
		public int PostId; //PostId
		public int ByUserId; //ByUserId
		public int PostUserId; //PostUserId
		public int CreatedTime; //CreatedTime
	}

	public static class ProfileAll {//oridnal: 18
		public int Id; //Id
		public int UserId; //UserId
		public int PostId; //PostId
		public int IsReShared; //IsReShared
	}

	public static class ProfileMedia {//oridnal: 19
		public int Id; //Id
		public int UserId; //UserId
		public int PostId; //PostId
		public int PostType; //PostType
	}

	public static class ProfileMentioned {//oridnal: 20
		public int Id; //Id
		public int ForUserId; //ForUserId
		public int PostId; //PostId
		public int PostUserId; //PostUserId
		public int PostType; //PostType
		public int CreatedTime; //CreatedTime
	}

	public static class Session {//oridnal: 21
		public int Id; //Id
		public String SessionUuid; //SessionUuid
		public int UserId; //UserId
		public String LastIpAddress; //LastIpAddress
		public int AppVersion; //AppVersion
		public int ActiveTime; //ActiveTime
		public int CreatedTime; //CreatedTime
	}

	public static class SettingNotifications {//oridnal: 22
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

	public static class Sms {//oridnal: 23
		public int Id; //Id
		public String Hash; //Hash
		public String AppUuid; //AppUuid
		public String ClientPhone; //ClientPhone
		public int GenratedCode; //GenratedCode
		public String SmsSenderNumber; //SmsSenderNumber
		public String SmsSendStatues; //SmsSendStatues
		public String SmsHttpBody; //SmsHttpBody
		public String Err; //Err
		public String Carrier; //Carrier
		public UNKNOWN Country; //Country
		public int IsValidPhone; //IsValidPhone
		public int IsConfirmed; //IsConfirmed
		public int IsLogin; //IsLogin
		public int IsRegister; //IsRegister
		public int RetriedCount; //RetriedCount
		public int TTL; //TTL
	}

	public static class Tag {//oridnal: 24
		public int TagId; //TagId
		public String Name; //Name
		public int Count; //Count
		public int TagStatusEnum; //TagStatusEnum
		public int IsBlocked; //IsBlocked
		public int GroupId; //GroupId
		public int CreatedTime; //CreatedTime
	}

	public static class TagPost {//oridnal: 25
		public int Id; //Id
		public int TagId; //TagId
		public int PostId; //PostId
		public int PostType; //PostType
		public int CreatedTime; //CreatedTime
	}

	public static class TriggerLog {//oridnal: 26
		public int Id; //Id
		public String ModelName; //ModelName
		public String ChangeType; //ChangeType
		public int TargetId; //TargetId
		public String TargetStr; //TargetStr
		public int CreatedSe; //CreatedSe
	}

	public static class User {//oridnal: 27
		public int UserId; //UserId
		public String UserName; //UserName
		public String UserNameLower; //UserNameLower
		public String FirstName; //FirstName
		public String LastName; //LastName
		public int IsVerified; //IsVerified
		public int AvatarId; //AvatarId
		public int AccessHash; //AccessHash
		public int ProfilePrivacy; //ProfilePrivacy
		public int OnlinePrivacy; //OnlinePrivacy
		public int CallPrivacy; //CallPrivacy
		public int AddToGroupPrivacy; //AddToGroupPrivacy
		public int SeenMessagePrivacy; //SeenMessagePrivacy
		public String Phone; //Phone
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

	public static class UserRelation {//oridnal: 28
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

	public static class Chat {//oridnal: 29
		public int ChatId; //ChatId
		public String ChatKey; //ChatKey
		public String RoomKey; //RoomKey
		public int RoomType; //RoomType
		public int UserId; //UserId
		public int PeerUserId; //PeerUserId
		public int GroupId; //GroupId
		public int HashTagId; //HashTagId
		public String Title; //Title
		public int PinTimeMs; //PinTimeMs
		public int FromMsgId; //FromMsgId
		public int UnseenCount; //UnseenCount
		public int Seq; //Seq
		public int LastMsgId; //LastMsgId
		public int LastMyMsgStatus; //LastMyMsgStatus
		public int MyLastSeenSeq; //MyLastSeenSeq
		public int MyLastSeenMsgId; //MyLastSeenMsgId
		public int PeerLastSeenMsgId; //PeerLastSeenMsgId
		public int MyLastDeliveredSeq; //MyLastDeliveredSeq
		public int MyLastDeliveredMsgId; //MyLastDeliveredMsgId
		public int PeerLastDeliveredMsgId; //PeerLastDeliveredMsgId
		public int IsActive; //IsActive
		public int IsLeft; //IsLeft
		public int IsCreator; //IsCreator
		public int IsKicked; //IsKicked
		public int IsAdmin; //IsAdmin
		public int IsDeactivated; //IsDeactivated
		public int IsMuted; //IsMuted
		public int MuteUntil; //MuteUntil
		public int VersionTimeMs; //VersionTimeMs
		public int VersionSeq; //VersionSeq
		public int OrderTime; //OrderTime
		public int CreatedTime; //CreatedTime
		public String DraftText; //DraftText
		public int DratReplyToMsgId; //DratReplyToMsgId
	}

	public static class ChatDeleted {//oridnal: 30
		public int ChatId; //ChatId
		public String RoomKey; //RoomKey
	}

	public static class ChatLastMessage {//oridnal: 31
		public String ChatIdGroupId; //ChatIdGroupId
		public UNKNOWN LastMsgPb; //LastMsgPb
	}

	public static class ChatUserVersion {//oridnal: 32
		public int UserId; //UserId
		public int ChatId; //ChatId
		public int VersionTimeNano; //VersionTimeNano
	}

	public static class Group {//oridnal: 33
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
		public String InviteLinkHash; //InviteLinkHash
		public int MembersCount; //MembersCount
		public int AdminsCount; //AdminsCount
		public int ModeratorCounts; //ModeratorCounts
		public int SortTime; //SortTime
		public int CreatedTime; //CreatedTime
		public int IsMute; //IsMute
		public int IsDeleted; //IsDeleted
		public int IsBanned; //IsBanned
	}

	public static class GroupMember {//oridnal: 34
		public int OrderId; //OrderId
		public int GroupId; //GroupId
		public int UserId; //UserId
		public int ByUserId; //ByUserId
		public int GroupRole; //GroupRole
		public int CreatedTime; //CreatedTime
	}

	public static class GroupOrderdUser {//oridnal: 35
		public int OrderId; //OrderId
		public int GroupId; //GroupId
		public int UserId; //UserId
	}

	public static class GroupPinedMsg {//oridnal: 36
		public int MsgId; //MsgId
		public UNKNOWN MsgPb; //MsgPb
	}

	public static class FileMsg {//oridnal: 37
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

	public static class FilePost {//oridnal: 38
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

	public static class ActionFanout {//oridnal: 39
		public int OrderId; //OrderId
		public int ForUserId; //ForUserId
		public int ActionId; //ActionId
		public int ActorUserId; //ActorUserId
	}

	public static class HomeFanout {//oridnal: 40
		public int OrderId; //OrderId
		public int ForUserId; //ForUserId
		public int PostId; //PostId
		public int PostUserId; //PostUserId
		public int ResharedId; //ResharedId
	}

	public static class SuggestedTopPosts {//oridnal: 41
		public int Id; //Id
		public int PostId; //PostId
	}

	public static class SuggestedUser {//oridnal: 42
		public int Id; //Id
		public int UserId; //UserId
		public int TargetId; //TargetId
		public float Weight; //Weight
		public int CreatedTime; //CreatedTime
	}

	public static class PushChat {//oridnal: 43
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

	public static class HTTPRPCLog {//oridnal: 44
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

	public static class MetricLog {//oridnal: 45
		public int Id; //Id
		public int InstanceId; //InstanceId
		public String StartFrom; //StartFrom
		public String EndTo; //EndTo
		public int StartTime; //StartTime
		public String Duration; //Duration
		public String MetericsJson; //MetericsJson
	}

	public static class XfileServiceInfoLog {//oridnal: 46
		public int Id; //Id
		public int InstanceId; //InstanceId
		public String Url; //Url
		public String CreatedTime; //CreatedTime
	}

	public static class XfileServiceMetricLog {//oridnal: 47
		public int Id; //Id
		public int InstanceId; //InstanceId
		public String MetricJson; //MetricJson
	}

	public static class XfileServiceRequestLog {//oridnal: 48
		public int Id; //Id
		public int LocalSeq; //LocalSeq
		public int InstanceId; //InstanceId
		public String Url; //Url
		public int HttpCode; //HttpCode
		public String CreatedTime; //CreatedTime
	}

	public static class InvalidateCache {//oridnal: 49
		public int OrderId; //OrderId
		public String CacheKey; //CacheKey
	}

}