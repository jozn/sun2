package com.mardomsara.social.json;

public class J {
	public static class Action {//oridnal: 0
		public int ActionId;
		public int ActorUserId;
		public int ActionTypeEnum;
		public int PeerUserId;
		public int PostId;
		public int CommentId;
		public int Murmur64Hash;
		public int CreatedTime;
	}

	public static class Comment {//oridnal: 1
		public int CommentId;
		public int UserId;
		public int PostId;
		public String Text;
		public int LikesCount;
		public int CreatedTime;
	}

	public static class Event {//oridnal: 2
		public int EventId;
		public int EventType;
		public int ByUserId;
		public int PeerUserId;
		public int PostId;
		public int CommentId;
		public int ActionId;
		public int Murmur64Hash;
	}

	public static class FollowingList {//oridnal: 3
		public int Id;
		public int UserId;
		public int ListType;
		public String Name;
		public int Count;
		public int IsAuto;
		public int IsPimiry;
		public int CreatedTime;
	}

	public static class FollowingListMember {//oridnal: 4
		public int Id;
		public int ListId;
		public int UserId;
		public int FollowedUserId;
		public int CreatedTime;
	}

	public static class FollowingListMemberRemoved {//oridnal: 5
		public int Id;
		public int ListId;
		public int UserId;
		public int UnFollowedUserId;
		public int UpdatedTime;
	}

	public static class Group {//oridnal: 6
		public int GroupId;
		public String GroupName;
		public int MembersCount;
		public int GroupPrivacyEnum;
		public int CreatorUserId;
		public int CreatedTime;
		public int UpdatedMs;
		public int CurrentSeq;
	}

	public static class GroupMember {//oridnal: 7
		public int Id;
		public int GroupId;
		public String GroupKey;
		public int UserId;
		public int ByUserId;
		public int GroupRoleEnumId;
		public int CreatedTime;
	}

	public static class GroupMessage {//oridnal: 8
		public int MessageId;
		public String RoomKey;
		public int UserId;
		public int MessageFileId;
		public int MessageTypeEnum;
		public String Text;
		public int CreatedMs;
		public int DeliveryStatusEnum;
	}

	public static class Like {//oridnal: 9
		public int Id;
		public int PostId;
		public int PostTypeEnum;
		public int UserId;
		public int LikeEnum;
		public int CreatedTime;
	}

	public static class Media {//oridnal: 10
		public int MediaId;
		public int UserId;
		public int PostId;
		public int AlbumId;
		public int MediaTypeEnum;
		public int Width;
		public int Height;
		public int Size;
		public int Duration;
		public String Md5Hash;
		public String Color;
		public int CreatedTime;
	}

	public static class Notify {//oridnal: 11
		public int NotifyId;
		public int ForUserId;
		public int ActorUserId;
		public int NotifyTypeEnum;
		public int PostId;
		public int CommentId;
		public int PeerUserId;
		public int Murmur64Hash;
		public int SeenStatus;
		public int CreatedTime;
	}

	public static class NotifyRemoved {//oridnal: 12
		public int Murmur64Hash;
		public int ForUserId;
		public int Id;
	}

	public static class PhoneContact {//oridnal: 13
		public int Id;
		public int UserId;
		public int Phone;
		public String PhoneDisplayName;
		public String PhoneFamilyName;
		public String PhoneNumber;
		public String PhoneNormalizedNumber;
		public int PhoneContactRowId;
		public int DeviceUuidId;
		public int CreatedTime;
	}

	public static class Post {//oridnal: 14
		public int PostId;
		public int UserId;
		public int PostTypeEnum;
		public int MediaId;
		public String PostKey;
		public String Text;
		public String RichText;
		public int MediaCount;
		public int SharedTo;
		public int DisableComment;
		public int HasTag;
		public int Seq;
		public int CommentsCount;
		public int LikesCount;
		public int ViewsCount;
		public int EditedTime;
		public int CreatedTime;
		public int ReSharedPostId;
	}

	public static class PostKey {//oridnal: 15
		public int Id;
		public String PostKeyStr;
		public int Used;
	}

	public static class SearchClicked {//oridnal: 16
		public int Id;
		public String Query;
		public int ClickType;
		public int TargetId;
		public int UserId;
		public int CreatedTime;
	}

	public static class Session {//oridnal: 17
		public String SessionUuid;
		public int UserId;
		public String ClientUuid;
		public String DeviceUuid;
		public int LastActivityTime;
		public String LastIpAddress;
		public String LastWifiMacAddress;
		public String LastNetworkType;
		public int LastNetworkTypeEnum;
		public int AppVersion;
		public int UpdatedTime;
		public int CreatedTime;
	}

	public static class SettingClient {//oridnal: 18
		public int UserId;
		public int AutoDownloadWifiVoice;
		public int AutoDownloadWifiImage;
		public int AutoDownloadWifiVideo;
		public int AutoDownloadWifiFile;
		public int AutoDownloadWifiMusic;
		public int AutoDownloadWifiGif;
		public int AutoDownloadCellularVoice;
		public int AutoDownloadCellularImage;
		public int AutoDownloadCellularVideo;
		public int AutoDownloadCellularFile;
		public int AutoDownloadCellularMusic;
		public int AutoDownloadCellularGif;
		public int AutoDownloadRoamingVoice;
		public int AutoDownloadRoamingImage;
		public int AutoDownloadRoamingVideo;
		public int AutoDownloadRoamingFile;
		public int AutoDownloadRoamingMusic;
		public int AutoDownloadRoamingGif;
		public int SaveToGallery;
	}

	public static class SettingNotification {//oridnal: 19
		public int UserId;
		public int SocialLedOn;
		public String SocialLedColor;
		public int ReqestToFollowYou;
		public int FollowedYou;
		public int AccptedYourFollowRequest;
		public int YourPostLiked;
		public int YourPostCommented;
		public int MenthenedYouInPost;
		public int MenthenedYouInComment;
		public int YourContactsJoined;
		public int DirectMessage;
		public int DirectAlert;
		public int DirectPerview;
		public int DirectLedOn;
		public int DirectLedColor;
		public int DirectVibrate;
		public int DirectPopup;
		public int DirectSound;
		public int DirectPriority;
	}

	public static class SuggestedTopPost {//oridnal: 20
		public int Id;
		public int PostId;
	}

	public static class SuggestedUser {//oridnal: 21
		public int Id;
		public int UserId;
		public int TargetId;
		public float Weight;
		public int CreatedTime;
	}

	public static class Tag {//oridnal: 22
		public int TagId;
		public String Name;
		public int Count;
		public int TagStatusEnum;
		public int CreatedTime;
	}

	public static class TagsPost {//oridnal: 23
		public int Id;
		public int TagId;
		public int PostId;
		public int PostTypeEnum;
		public int CreatedTime;
	}

	public static class TriggerLog {//oridnal: 24
		public int Id;
		public String ModelName;
		public String ChangeType;
		public int TargetId;
		public String TargetStr;
		public int CreatedSe;
	}

	public static class User {//oridnal: 25
		public int UserId;
		public String UserName;
		public String UserNameLower;
		public String FirstName;
		public String LastName;
		public int UserTypeEnum;
		public int UserLevelEnum;
		public int AvatarId;
		public int ProfilePrivacyEnum;
		public int Phone;
		public String About;
		public String Email;
		public String PasswordHash;
		public String PasswordSalt;
		public int PostSeq;
		public int FollowersCount;
		public int FollowingCount;
		public int PostsCount;
		public int MediaCount;
		public int LikesCount;
		public int ResharedCount;
		public int LastActionTime;
		public int LastPostTime;
		public int PrimaryFollowingList;
		public int CreatedSe;
		public int UpdatedMs;
		public int OnlinePrivacyEnum;
		public int LastActivityTime;
		public String Phone2;
	}

	public static class UserMetaInfo {//oridnal: 26
		public int Id;
		public int UserId;
		public int IsNotificationDirty;
		public int LastUserRecGen;
	}

	public static class UserPassword {//oridnal: 27
		public int UserId;
		public String Password;
		public int CreatedTime;
	}

	public static class Chat {//oridnal: 28
		public String ChatKey;
		public String RoomKey;
		public int RoomTypeEnum;
		public int UserId;
		public int PeerUserId;
		public int GroupId;
		public int CreatedTime;
		public int Seq;
		public int SeenSeq;
		public int UpdatedMs;
	}

	public static class ChatLastMessage {//oridnal: 29
		public String ChatKey;
		public int ForUserId;
		public UNKNOWN LastMsgPb;
		public String LastMsgJson;
	}

	public static class ChatSync {//oridnal: 30
		public int SyncId;
		public int ToUserId;
		public int ChatSyncTypeId;
		public String RoomKey;
		public String ChatKey;
		public int FromHighMessageId;
		public int ToLowMessageId;
		public int MessageId;
		public UNKNOWN MessagePb;
		public String MessageJson;
		public int CreatedTime;
	}

	public static class DirectMessage {//oridnal: 31
		public String ChatKey;
		public int MessageId;
		public String RoomKey;
		public int UserId;
		public int MessageFileId;
		public int MessageTypeEnum;
		public String Text;
		public int CreatedTime;
		public int Seq;
		public int DeliviryStatusEnum;
		public UNKNOWN ExtraPB;
	}

	public static class Home {//oridnal: 32
		public int Id;
		public int ForUserId;
		public int PostId;
	}

	public static class MessageFile {//oridnal: 33
		public int MessageFileId;
		public int FileTypeEnum;
		public int UserId;
		public String Title;
		public int Size;
		public int Width;
		public int Height;
		public int Duration;
		public String Extension;
		public String Md5Hash;
		public int CreatedTime;
	}

	public static class FileMsg {//oridnal: 34
		public int Id;
		public int FileType;
		public String Extension;
		public UNKNOWN DataThumb;
		public UNKNOWN Data;
	}

	public static class FilePost {//oridnal: 35
		public int Id;
		public int FileType;
		public String Extension;
		public UNKNOWN DataThumb;
		public UNKNOWN Data;
	}

	public static class ActionFanout {//oridnal: 36
		public int OrderId;
		public int ForUserId;
		public int ActionId;
		public int ActorUserId;
		public int Murmur64Hash;
	}

	public static class HomeFanout {//oridnal: 37
		public int OrderId;
		public int ForUserId;
		public int PostId;
		public int PostUserId;
	}

}