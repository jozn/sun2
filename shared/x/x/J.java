package com.mardomsara.social.json;

public class J {
	public static class Activity {
		public int Id;
		public int ActorUserId;
		public int ActionTypeId;
		public int RowId;
		public int RootId;
		public int RefId;
		public int CreatedAt;
	}

	public static class Bucket {
		public int BucketId;
		public String BucketName;
		public int Server1Id;
		public int Server2Id;
		public int BackupServerId;
		public int ContentObjectTypeId;
		public int ContentObjectCount;
		public int CreatedTime;
	}

	public static class Chat {
		public String ChatKey;
		public String RoomKey;
		public int RoomTypeEnumId;
		public int UserId;
		public int PeerUserId;
		public int GroupId;
		public int CreatedSe;
		public int StartMessageIdFrom;
		public int LastSeenMessageId;
		public int UpdatedMs;
		public int LastMessageId;
		public int LastDeletedMessageId;
		public int LastSeqSeen;
		public int LastSeqDelete;
		public int CurrentSeq;
	}

	public static class Comment {
		public int Id;
		public int UserId;
		public int PostId;
		public String Text;
		public int CreatedTime;
	}

	public static class DirectMessage {
		public int MessageId;
		public String MessageKey;
		public String RoomKey;
		public int UserId;
		public int MessageFileId;
		public int MessageTypeEnumId;
		public String Text;
		public int CreatedSe;
		public int PeerReceivedTime;
		public int PeerSeenTime;
		public int DeliviryStatusEnumId;
	}

	public static class DirectToMessage {
		public int Id;
		public String ChatKey;
		public int MessageId;
		public int SourceEnumId;
	}

	public static class DirectUpdate {
		public int DirectUpdateId;
		public int ToUserId;
		public int MessageId;
		public int MessageFileId;
		public int OtherId;
		public String ChatKey;
		public int PeerUserId;
		public int EventType;
		public int RoomLogTypeId;
		public int FromSeq;
		public int ToSeq;
		public UNKNOWN ExtraPB;
		public String ExtraJson;
		public int AtTimeMs;
	}

	public static class FollowingList {
		public int Id;
		public int UserId;
		public int ListType;
		public String Name;
		public int Count;
		public int IsAuto;
		public int IsPimiry;
		public int CreatedTime;
	}

	public static class FollowingListMember {
		public int Id;
		public int ListId;
		public int UserId;
		public int FollowedUserId;
		public int FollowType;
		public int UpdatedTimeMs;
	}

	public static class FollowingListMemberHistory {
		public int Id;
		public int ListId;
		public int UserId;
		public int FollowedUserId;
		public int FollowType;
		public int UpdatedTimeMs;
		public int FollowId;
	}

	public static class GeneralLog {
		public int Id;
		public int ToUserId;
		public int TargetId;
		public int LogTypeId;
		public UNKNOWN ExtraPB;
		public String ExtraJson;
		public int CreatedMs;
	}

	public static class Group {
		public int GroupId;
		public String GroupName;
		public int MembersCount;
		public int GroupPrivacyEnum;
		public int CreatorUserId;
		public int CreatedTime;
		public int UpdatedMs;
		public int CurrentSeq;
	}

	public static class GroupMember {
		public int Id;
		public int GroupId;
		public String GroupKey;
		public int UserId;
		public int ByUserId;
		public int GroupRoleEnumId;
		public int CreatedTime;
	}

	public static class GroupMessage {
		public int MessageId;
		public String RoomKey;
		public int UserId;
		public int MessageFileId;
		public int MessageTypeEnum;
		public String Text;
		public int CreatedMs;
		public int DeliveryStatusEnum;
	}

	public static class GroupToMessage {
		public int Id;
		public int GroupId;
		public int MessageId;
		public int CreatedMs;
		public int StatusEnum;
	}

	public static class Like {
		public int Id;
		public int PostId;
		public int PostTypeId;
		public int UserId;
		public int TypeId;
		public int CreatedTime;
	}

	public static class LogChange {
		public int Id;
		public String T;
	}

	public static class Media {
		public int Id;
		public int UserId;
		public int PostId;
		public int AlbumId;
		public int TypeId;
		public int CreatedTime;
		public String Src;
	}

	public static class MessageFile {
		public int MessageFileId;
		public String MessageFileKey;
		public int OriginalUserId;
		public String Name;
		public int Size;
		public int FileTypeEnumId;
		public int Width;
		public int Height;
		public int Duration;
		public String Extension;
		public String HashMd5;
		public int HashAccess;
		public int CreatedSe;
		public String ServerSrc;
		public String ServerPath;
		public String ServerThumbPath;
		public String BucketId;
		public int ServerId;
		public int CanDel;
	}

	public static class Notification {
		public int Id;
		public int ForUserId;
		public int ActorUserId;
		public int ActionTypeId;
		public int ObjectTypeId;
		public int RowId;
		public int RootId;
		public int RefId;
		public int SeenStatus;
		public int CreatedTime;
	}

	public static class NotificationRemoved {
		public int NotificationId;
		public int ForUserId;
	}

	public static class OldMessage {
		public int Id;
		public int Uid;
		public int UserId;
		public String MessageKey;
		public String RoomKey;
		public int MessageType;
		public int RoomType;
		public int MsgFileId;
		public UNKNOWN DataPB;
		public String Data64;
		public String DataJson;
		public int CreatedTimeMs;
	}

	public static class OldMsgFile {
		public int Id;
		public String Name;
		public int Size;
		public int FileType;
		public String MimeType;
		public int Width;
		public int Height;
		public int Duration;
		public String Extension;
		public UNKNOWN ThumbData;
		public String ThumbData64;
		public String ServerSrc;
		public String ServerPath;
		public int ServerId;
		public int CanDel;
	}

	public static class OldMsgPush {
		public int Id;
		public int Uid;
		public int ToUser;
		public int MsgUid;
		public int CreatedTimeMs;
	}

	public static class OldMsgPushEvent {
		public int Id;
		public int Uid;
		public int ToUserId;
		public int MsgUid;
		public String MsgKey;
		public String RoomKey;
		public int PeerUserId;
		public int EventType;
		public int AtTime;
	}

	public static class PhoneContact {
		public int Id;
		public String PhoneDisplayName;
		public String PhoneFamilyName;
		public String PhoneNumber;
		public String PhoneNormalizedNumber;
		public int PhoneContactRowId;
		public int UserId;
		public int DeviceUuidId;
		public int CreatedTime;
		public int UpdatedTime;
	}

	public static class Photo {
		public int PhotoId;
		public int UserId;
		public int PostId;
		public int AlbumId;
		public int ImageTypeId;
		public String Title;
		public String Src;
		public String PathSrc;
		public int BucketId;
		public int Width;
		public int Height;
		public float Ratio;
		public String HashMd5;
		public String Color;
		public int CreatedTime;
		public int W1080;
		public int W720;
		public int W480;
		public int W320;
		public int W160;
		public int W80;
	}

	public static class Post {
		public int Id;
		public int UserId;
		public int TypeId;
		public String Text;
		public String FormatedText;
		public int MediaCount;
		public int SharedTo;
		public int DisableComment;
		public int HasTag;
		public int LikesCount;
		public int CommentsCount;
		public int EditedTime;
		public int CreatedTime;
	}

	public static class PushEvent {
		public int PushEventId;
		public int ToUserId;
		public int ToDeviceId;
		public int MessageId;
		public int RoomTypeEnum;
		public int RoomId;
		public int PeerUserId;
		public int PushEventTypeEnum;
		public int AtTime;
	}

	public static class PushMessage {
		public int PushMessageId;
		public int ToUserId;
		public int ToDeviceId;
		public int MessageId;
		public int RoomTypeEnum;
		public int CreatedMs;
	}

	public static class RecommendUser {
		public int Id;
		public int UserId;
		public int TargetId;
		public float Weight;
		public int CreatedTime;
	}

	public static class Room {
		public int RoomId;
		public String RoomKey;
		public int RoomTypeEnum;
		public int UserId;
		public int LastSeqSeen;
		public int LastSeqDelete;
		public int PeerUserId;
		public int GroupId;
		public int CreatedTime;
		public int CurrentSeq;
	}

	public static class SearchClicked {
		public int Id;
		public String Query;
		public int ClickType;
		public int TargetId;
		public int UserId;
		public int CreatedAt;
	}

	public static class Session {
		public int Id;
		public int UserId;
		public String SessionUuid;
		public String ClientUuid;
		public String DeviceUuid;
		public int LastActivityTime;
		public String LastIpAddress;
		public String LastWifiMacAddress;
		public String LastNetworkType;
		public int LastNetworkTypeId;
		public int AppVersion;
		public int UpdatedTime;
		public int CreatedTime;
		public int LastSyncDirectUpdateId;
		public int LastSyncGeneralUpdateId;
		public int LastSyncNotifyUpdateId;
	}

	public static class SettingClient {
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

	public static class SettingNotification {
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

	public static class Tag {
		public int Id;
		public String Name;
		public int Count;
		public int IsBlocked;
		public int CreatedTime;
	}

	public static class TagsPost {
		public int Id;
		public int TagId;
		public int PostId;
		public int TypeId;
		public int CreatedTime;
	}

	public static class TestChat {
		public int Id;
		public int Id4;
		public int TimeMs;
		public String Text;
		public String Name;
		public int UserId;
		public int C2;
		public int C3;
		public int C4;
		public int C5;
	}

	public static class User {
		public int Id;
		public String UserName;
		public String UserNameLower;
		public String FirstName;
		public String LastName;
		public String About;
		public String FullName;
		public String AvatarUrl;
		public int PrivacyProfile;
		public String Phone;
		public String Email;
		public int IsDeleted;
		public String PasswordHash;
		public String PasswordSalt;
		public int FollowersCount;
		public int FollowingCount;
		public int PostsCount;
		public int MediaCount;
		public int LikesCount;
		public int ResharedCount;
		public int LastActionTime;
		public int LastPostTime;
		public int PrimaryFollowingList;
		public int CreatedTime;
		public int UpdatedTime;
		public String SessionUuid;
		public String DeviceUuid;
		public String LastWifiMacAddress;
		public String LastNetworkType;
		public int AppVersion;
		public int LastActivityTime;
		public int LastLoginTime;
		public String LastIpAddress;
	}

	public static class UserMetaInfo {
		public int Id;
		public int UserId;
		public int IsNotificationDirty;
		public int LastUserRecGen;
	}

	public static class UserPassword {
		public int UserId;
		public String Password;
		public int CreatedTime;
	}

}