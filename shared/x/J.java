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

	public static class FollowingList {//oridnal: 4
		public int Id; //Id
		public int UserId; //UserId
		public int ListType; //ListType
		public String Name; //Name
		public int Count; //Count
		public int IsAuto; //IsAuto
		public int IsPimiry; //IsPimiry
		public int CreatedTime; //CreatedTime
	}

	public static class FollowingListMember {//oridnal: 5
		public int Id; //Id
		public int ListId; //ListId
		public int UserId; //UserId
		public int FollowedUserId; //FollowedUserId
		public int CreatedTime; //CreatedTime
	}

	public static class FollowingListMemberRemoved {//oridnal: 6
		public int Id; //Id
		public int ListId; //ListId
		public int UserId; //UserId
		public int UnFollowedUserId; //UnFollowedUserId
		public int UpdatedTime; //UpdatedTime
	}

	public static class Like {//oridnal: 7
		public int Id; //Id
		public int PostId; //PostId
		public int PostTypeEnum; //PostTypeEnum
		public int UserId; //UserId
		public int LikeEnum; //LikeEnum
		public int CreatedTime; //CreatedTime
	}

	public static class Notify {//oridnal: 8
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

	public static class NotifyRemoved {//oridnal: 9
		public int Murmur64Hash; //Murmur64Hash
		public int ForUserId; //ForUserId
		public int Id; //Id
	}

	public static class PhoneContact {//oridnal: 10
		public int Id; //Id
		public int UserId; //UserId
		public int Phone; //Phone
		public String PhoneDisplayName; //PhoneDisplayName
		public String PhoneFamilyName; //PhoneFamilyName
		public String PhoneNumber; //PhoneNumber
		public String PhoneNormalizedNumber; //PhoneNormalizedNumber
		public int PhoneContactRowId; //PhoneContactRowId
		public int DeviceUuidId; //DeviceUuidId
		public int CreatedTime; //CreatedTime
	}

	public static class Post {//oridnal: 11
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

	public static class PostCount {//oridnal: 12
		public int PostId; //PostId
		public int ViewsCount; //ViewsCount
	}

	public static class PostDeleted {//oridnal: 13
		public int PostId; //PostId
		public int UserId; //UserId
	}

	public static class PostKey {//oridnal: 14
		public int Id; //Id
		public String PostKeyStr; //PostKeyStr
		public int Used; //Used
	}

	public static class PostLink {//oridnal: 15
		public int LinkId; //LinkId
		public String LinkUrl; //LinkUrl
	}

	public static class PostMedia {//oridnal: 16
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

	public static class PostMentioned {//oridnal: 17
		public int MentionedId; //MentionedId
		public int ForUserId; //ForUserId
		public int PostId; //PostId
		public int PostUserId; //PostUserId
		public int PostTypeEnum; //PostTypeEnum
		public int PostCategoryEnum; //PostCategoryEnum
		public int CreatedTime; //CreatedTime
	}

	public static class PostReshared {//oridnal: 18
		public int ResharedId; //ResharedId
		public int ByUserId; //ByUserId
		public int PostId; //PostId
		public int PostUserId; //PostUserId
		public int PostTypeEnum; //PostTypeEnum
		public int PostCategoryEnum; //PostCategoryEnum
		public int CreatedTime; //CreatedTime
	}

	public static class SearchClicked {//oridnal: 19
		public int Id; //Id
		public String Query; //Query
		public int ClickType; //ClickType
		public int TargetId; //TargetId
		public int UserId; //UserId
		public int CreatedTime; //CreatedTime
	}

	public static class Session {//oridnal: 20
		public int Id; //Id
		public String SessionUuid; //SessionUuid
		public int UserId; //UserId
		public String LastIpAddress; //LastIpAddress
		public int AppVersion; //AppVersion
		public int ActiveTime; //ActiveTime
		public int CreatedTime; //CreatedTime
	}

	public static class SettingClient {//oridnal: 21
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

	public static class SettingNotification {//oridnal: 22
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

	public static class Tag {//oridnal: 23
		public int TagId; //TagId
		public String Name; //Name
		public int Count; //Count
		public int TagStatusEnum; //TagStatusEnum
		public int CreatedTime; //CreatedTime
	}

	public static class TagPost {//oridnal: 24
		public int Id; //Id
		public int TagId; //TagId
		public int PostId; //PostId
		public int PostTypeEnum; //PostTypeEnum
		public int PostCategoryEnum; //PostCategoryEnum
		public int CreatedTime; //CreatedTime
	}

	public static class TriggerLog {//oridnal: 25
		public int Id; //Id
		public String ModelName; //ModelName
		public String ChangeType; //ChangeType
		public int TargetId; //TargetId
		public String TargetStr; //TargetStr
		public int CreatedSe; //CreatedSe
	}

	public static class User {//oridnal: 26
		public int UserId; //UserId
		public String UserName; //UserName
		public String UserNameLower; //UserNameLower
		public String FirstName; //FirstName
		public String LastName; //LastName
		public int UserTypeEnum; //UserTypeEnum
		public int UserLevelEnum; //UserLevelEnum
		public int AvatarId; //AvatarId
		public int ProfilePrivacyEnum; //ProfilePrivacyEnum
		public int Phone; //Phone
		public String About; //About
		public String Email; //Email
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
		public int LastActionTime; //LastActionTime
		public int LastPostTime; //LastPostTime
		public int PrimaryFollowingList; //PrimaryFollowingList
		public int CreatedSe; //CreatedSe
		public int UpdatedMs; //UpdatedMs
		public int OnlinePrivacyEnum; //OnlinePrivacyEnum
		public int LastActivityTime; //LastActivityTime
		public String Phone2; //Phone2
	}

	public static class UserMetaInfo {//oridnal: 27
		public int Id; //Id
		public int UserId; //UserId
		public int IsNotificationDirty; //IsNotificationDirty
		public int LastUserRecGen; //LastUserRecGen
	}

	public static class UserPassword {//oridnal: 28
		public int UserId; //UserId
		public String Password; //Password
		public int CreatedTime; //CreatedTime
	}

	public static class Chat {//oridnal: 29
		public String ChatKey; //ChatKey
		public String RoomKey; //RoomKey
		public int RoomTypeEnum; //RoomTypeEnum
		public int UserId; //UserId
		public int PeerUserId; //PeerUserId
		public int GroupId; //GroupId
		public int CreatedTime; //CreatedTime
		public int Seq; //Seq
		public int SeenSeq; //SeenSeq
		public int UpdatedMs; //UpdatedMs
	}

	public static class ChatLastMessage {//oridnal: 30
		public String ChatKey; //ChatKey
		public int ForUserId; //ForUserId
		public UNKNOWN LastMsgPb; //LastMsgPb
		public String LastMsgJson; //LastMsgJson
	}

	public static class DirectMessage {//oridnal: 31
		public String ChatKey; //ChatKey
		public int MessageId; //MessageId
		public String RoomKey; //RoomKey
		public int UserId; //UserId
		public int MessageFileId; //MessageFileId
		public int MessageTypeEnum; //MessageTypeEnum
		public String Text; //Text
		public int CreatedTime; //CreatedTime
		public int Seq; //Seq
		public int DeliviryStatusEnum; //DeliviryStatusEnum
		public UNKNOWN ExtraPB; //ExtraPB
	}

	public static class Group {//oridnal: 32
		public int GroupId; //GroupId
		public String GroupName; //GroupName
		public int MembersCount; //MembersCount
		public int GroupPrivacyEnum; //GroupPrivacyEnum
		public int CreatorUserId; //CreatorUserId
		public int CreatedTime; //CreatedTime
		public int UpdatedMs; //UpdatedMs
		public int CurrentSeq; //CurrentSeq
	}

	public static class GroupMember {//oridnal: 33
		public int Id; //Id
		public int GroupId; //GroupId
		public String GroupKey; //GroupKey
		public int UserId; //UserId
		public int ByUserId; //ByUserId
		public int GroupRoleEnumId; //GroupRoleEnumId
		public int CreatedTime; //CreatedTime
	}

	public static class GroupMessage {//oridnal: 34
		public int MessageId; //MessageId
		public String RoomKey; //RoomKey
		public int UserId; //UserId
		public int MessageFileId; //MessageFileId
		public int MessageTypeEnum; //MessageTypeEnum
		public String Text; //Text
		public int CreatedMs; //CreatedMs
		public int DeliveryStatusEnum; //DeliveryStatusEnum
	}

	public static class FileMsg {//oridnal: 35
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

	public static class FilePost {//oridnal: 36
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

	public static class ActionFanout {//oridnal: 37
		public int OrderId; //OrderId
		public int ForUserId; //ForUserId
		public int ActionId; //ActionId
		public int ActorUserId; //ActorUserId
	}

	public static class HomeFanout {//oridnal: 38
		public int OrderId; //OrderId
		public int ForUserId; //ForUserId
		public int PostId; //PostId
		public int PostUserId; //PostUserId
		public int ResharedId; //ResharedId
	}

	public static class SuggestedTopPost {//oridnal: 39
		public int Id; //Id
		public int PostId; //PostId
	}

	public static class SuggestedUser {//oridnal: 40
		public int Id; //Id
		public int UserId; //UserId
		public int TargetId; //TargetId
		public float Weight; //Weight
		public int CreatedTime; //CreatedTime
	}

	public static class ChatSync2 {//oridnal: 41
		public int SyncId; //sync_id
		public int ToUserId; //to_user_id
		public int ChatSyncTypeId; //chat_sync_type_id
		public String RoomKey; //room_key
		public String ChatKey; //chat_key
		public int FromHighMessageId; //from_high_message_id
		public int ToLowMessageId; //to_low_message_id
		public int MessageId; //message_id
		public UNKNOWN MessagePb; //message_pb
		public String MessageJson; //message_json
		public int CreatedTime; //created_time
	}

	public static class LowerTable {//oridnal: 42
		public int Id; //id
		public String Text; //text
		public int TimeStamp; //time_stamp
		public int AnyThingMore; //any_thing_more_
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

	public static class PushChat2 {//oridnal: 44
		public int SyncId; //sync_id
		public int ToUserId; //to_user_id
		public int ChatSyncTypeId; //chat_sync_type_id
		public String RoomKey; //room_key
		public String ChatKey; //chat_key
		public int FromHighMessageId; //from_high_message_id
		public int ToLowMessageId; //to_low_message_id
		public int MessageId; //message_id
		public UNKNOWN MessagePb; //message_pb
		public String MessageJson; //message_json
		public int CreatedTime; //created_time
	}

	public static class HTTPRPCLog {//oridnal: 45
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

	public static class MetricLog {//oridnal: 46
		public int Id; //Id
		public int InstanceId; //InstanceId
		public String StartFrom; //StartFrom
		public String EndTo; //EndTo
		public int StartTime; //StartTime
		public String Duration; //Duration
		public String MetericsJson; //MetericsJson
	}

	public static class XfileServiceInfoLog {//oridnal: 47
		public int Id; //Id
		public int InstanceId; //InstanceId
		public String Url; //Url
		public String CreatedTime; //CreatedTime
	}

	public static class XfileServiceMetricLog {//oridnal: 48
		public int Id; //Id
		public int InstanceId; //InstanceId
		public String MetricJson; //MetricJson
	}

	public static class XfileServiceRequestLog {//oridnal: 49
		public int Id; //Id
		public int LocalSeq; //LocalSeq
		public int InstanceId; //InstanceId
		public String Url; //Url
		public int HttpCode; //HttpCode
		public String CreatedTime; //CreatedTime
	}

	public static class Account {//oridnal: 50
		public int Id; //id
		public double Balance; //balance
	}

	public static class PostCdb {//oridnal: 51
		public int PostId; //post_id
		public int UserId; //user_id
		public int PostTypeEnum; //post_type_enum
		public int PostCategoryEnum; //post_category_enum
		public int MediaId; //media_id
		public String PostKey; //post_key
		public String Text; //text
		public String RichText; //rich_text
		public int MediaCount; //media_count
		public int SharedTo; //shared_to
		public int DisableComment; //disable_comment
		public int Source; //source
		public int HasTag; //has_tag
		public int Seq; //seq
		public int CommentsCount; //comments_count
		public int LikesCount; //likes_count
		public int ViewsCount; //views_count
		public int EditedTime; //edited_time
		public int CreatedTime; //created_time
		public int ReSharedPostId; //re_shared_post_id
	}

}