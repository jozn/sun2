package ir.ms.pb;

public class PBFlatTypes {

	public class PB_RoomsChanges {
	   public long VersionTime;
	   public PB_ChatView rooms;
	}
	/*
	folding
	PBFlatTypes.PB_RoomsChanges t = new PBFlatTypes.PB_RoomsChanges();
    t.setVersionTime();
    t.setrooms();
	*/

	/*
	PBFlatTypes.PB_RoomsChanges t = new PBFlatTypes.PB_RoomsChanges();
	t.VersionTime = ;
	t.rooms = ;
	*/

	/*
	PB_RoomsChanges t = new PB_RoomsChanges();
	t.VersionTime = m.getVersionTime() ;
	t.rooms = m.getrooms() ;
	*/

	public class PB_PushChanges {
	   public PB_RoomsChanges RoomsChanges;
	   public PB_ChatView ChatView;
	   public String InvalidateCacheForRoomKeys;
	   public int InvalidateAllChatCache;
	   public int InvalidateAllSocialCache;
	}
	/*
	folding
	PBFlatTypes.PB_PushChanges t = new PBFlatTypes.PB_PushChanges();
    t.setRoomsChanges();
    t.setChatView();
    t.setInvalidateCacheForRoomKeys();
    t.setInvalidateAllChatCache();
    t.setInvalidateAllSocialCache();
	*/

	/*
	PBFlatTypes.PB_PushChanges t = new PBFlatTypes.PB_PushChanges();
	t.RoomsChanges = ;
	t.ChatView = ;
	t.InvalidateCacheForRoomKeys = ;
	t.InvalidateAllChatCache = ;
	t.InvalidateAllSocialCache = ;
	*/

	/*
	PB_PushChanges t = new PB_PushChanges();
	t.RoomsChanges = m.getRoomsChanges() ;
	t.ChatView = m.getChatView() ;
	t.InvalidateCacheForRoomKeys = m.getInvalidateCacheForRoomKeys() ;
	t.InvalidateAllChatCache = m.getInvalidateAllChatCache() ;
	t.InvalidateAllSocialCache = m.getInvalidateAllSocialCache() ;
	*/

	public class PB_CommandToServer {
	   public long ClientCallId;
	   public String Command;
	   public boolean RespondReached;
	   public byte[] Data;
	}
	/*
	folding
	PBFlatTypes.PB_CommandToServer t = new PBFlatTypes.PB_CommandToServer();
    t.setClientCallId();
    t.setCommand();
    t.setRespondReached();
    t.setData();
	*/

	/*
	PBFlatTypes.PB_CommandToServer t = new PBFlatTypes.PB_CommandToServer();
	t.ClientCallId = ;
	t.Command = ;
	t.RespondReached = ;
	t.Data = ;
	*/

	/*
	PB_CommandToServer t = new PB_CommandToServer();
	t.ClientCallId = m.getClientCallId() ;
	t.Command = m.getCommand() ;
	t.RespondReached = m.getRespondReached() ;
	t.Data = m.getData() ;
	*/

	public class PB_CommandToClient {
	   public long ServerCallId;
	   public String Command;
	   public boolean RespondReached;
	   public byte[] Data;
	}
	/*
	folding
	PBFlatTypes.PB_CommandToClient t = new PBFlatTypes.PB_CommandToClient();
    t.setServerCallId();
    t.setCommand();
    t.setRespondReached();
    t.setData();
	*/

	/*
	PBFlatTypes.PB_CommandToClient t = new PBFlatTypes.PB_CommandToClient();
	t.ServerCallId = ;
	t.Command = ;
	t.RespondReached = ;
	t.Data = ;
	*/

	/*
	PB_CommandToClient t = new PB_CommandToClient();
	t.ServerCallId = m.getServerCallId() ;
	t.Command = m.getCommand() ;
	t.RespondReached = m.getRespondReached() ;
	t.Data = m.getData() ;
	*/

	public class PB_CommandReachedToServer {
	   public long ClientCallId;
	}
	/*
	folding
	PBFlatTypes.PB_CommandReachedToServer t = new PBFlatTypes.PB_CommandReachedToServer();
    t.setClientCallId();
	*/

	/*
	PBFlatTypes.PB_CommandReachedToServer t = new PBFlatTypes.PB_CommandReachedToServer();
	t.ClientCallId = ;
	*/

	/*
	PB_CommandReachedToServer t = new PB_CommandReachedToServer();
	t.ClientCallId = m.getClientCallId() ;
	*/

	public class PB_CommandReachedToClient {
	   public long ServerCallId;
	}
	/*
	folding
	PBFlatTypes.PB_CommandReachedToClient t = new PBFlatTypes.PB_CommandReachedToClient();
    t.setServerCallId();
	*/

	/*
	PBFlatTypes.PB_CommandReachedToClient t = new PBFlatTypes.PB_CommandReachedToClient();
	t.ServerCallId = ;
	*/

	/*
	PB_CommandReachedToClient t = new PB_CommandReachedToClient();
	t.ServerCallId = m.getServerCallId() ;
	*/

	public class PB_ResponseToClient {
	   public long ClientCallId;
	   public String PBClass;
	   public String RpcFullName;
	   public byte[] Data;
	}
	/*
	folding
	PBFlatTypes.PB_ResponseToClient t = new PBFlatTypes.PB_ResponseToClient();
    t.setClientCallId();
    t.setPBClass();
    t.setRpcFullName();
    t.setData();
	*/

	/*
	PBFlatTypes.PB_ResponseToClient t = new PBFlatTypes.PB_ResponseToClient();
	t.ClientCallId = ;
	t.PBClass = ;
	t.RpcFullName = ;
	t.Data = ;
	*/

	/*
	PB_ResponseToClient t = new PB_ResponseToClient();
	t.ClientCallId = m.getClientCallId() ;
	t.PBClass = m.getPBClass() ;
	t.RpcFullName = m.getRpcFullName() ;
	t.Data = m.getData() ;
	*/

	public class RPC_Auth_Types {
	}
	/*
	folding
	PBFlatTypes.RPC_Auth_Types t = new PBFlatTypes.RPC_Auth_Types();
	*/

	/*
	PBFlatTypes.RPC_Auth_Types t = new PBFlatTypes.RPC_Auth_Types();
	*/

	/*
	RPC_Auth_Types t = new RPC_Auth_Types();
	*/

	public class PB_RPC_Chat_Types {
	}
	/*
	folding
	PBFlatTypes.PB_RPC_Chat_Types t = new PBFlatTypes.PB_RPC_Chat_Types();
	*/

	/*
	PBFlatTypes.PB_RPC_Chat_Types t = new PBFlatTypes.PB_RPC_Chat_Types();
	*/

	/*
	PB_RPC_Chat_Types t = new PB_RPC_Chat_Types();
	*/

	public class RPC_General_Types {
	}
	/*
	folding
	PBFlatTypes.RPC_General_Types t = new PBFlatTypes.RPC_General_Types();
	*/

	/*
	PBFlatTypes.RPC_General_Types t = new PBFlatTypes.RPC_General_Types();
	*/

	/*
	RPC_General_Types t = new RPC_General_Types();
	*/

	public class RPC_Page_Types {
	}
	/*
	folding
	PBFlatTypes.RPC_Page_Types t = new PBFlatTypes.RPC_Page_Types();
	*/

	/*
	PBFlatTypes.RPC_Page_Types t = new PBFlatTypes.RPC_Page_Types();
	*/

	/*
	RPC_Page_Types t = new RPC_Page_Types();
	*/

	public class RPC_Social_Types {
	}
	/*
	folding
	PBFlatTypes.RPC_Social_Types t = new PBFlatTypes.RPC_Social_Types();
	*/

	/*
	PBFlatTypes.RPC_Social_Types t = new PBFlatTypes.RPC_Social_Types();
	*/

	/*
	RPC_Social_Types t = new RPC_Social_Types();
	*/

	public class RPC_User_Types {
	}
	/*
	folding
	PBFlatTypes.RPC_User_Types t = new PBFlatTypes.RPC_User_Types();
	*/

	/*
	PBFlatTypes.RPC_User_Types t = new PBFlatTypes.RPC_User_Types();
	*/

	/*
	RPC_User_Types t = new RPC_User_Types();
	*/

	public class PB_Action {
	   public long ActionId;
	   public int ActorUserId;
	   public int ActionType;
	   public int PeerUserId;
	   public long PostId;
	   public long CommentId;
	   public long Murmur64Hash;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_Action t = new PBFlatTypes.PB_Action();
    t.setActionId();
    t.setActorUserId();
    t.setActionType();
    t.setPeerUserId();
    t.setPostId();
    t.setCommentId();
    t.setMurmur64Hash();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_Action t = new PBFlatTypes.PB_Action();
	t.ActionId = ;
	t.ActorUserId = ;
	t.ActionType = ;
	t.PeerUserId = ;
	t.PostId = ;
	t.CommentId = ;
	t.Murmur64Hash = ;
	t.CreatedTime = ;
	*/

	/*
	PB_Action t = new PB_Action();
	t.ActionId = m.getActionId() ;
	t.ActorUserId = m.getActorUserId() ;
	t.ActionType = m.getActionType() ;
	t.PeerUserId = m.getPeerUserId() ;
	t.PostId = m.getPostId() ;
	t.CommentId = m.getCommentId() ;
	t.Murmur64Hash = m.getMurmur64Hash() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_Blocked {
	   public long Id;
	   public int UserId;
	   public int BlockedUserId;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_Blocked t = new PBFlatTypes.PB_Blocked();
    t.setId();
    t.setUserId();
    t.setBlockedUserId();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_Blocked t = new PBFlatTypes.PB_Blocked();
	t.Id = ;
	t.UserId = ;
	t.BlockedUserId = ;
	t.CreatedTime = ;
	*/

	/*
	PB_Blocked t = new PB_Blocked();
	t.Id = m.getId() ;
	t.UserId = m.getUserId() ;
	t.BlockedUserId = m.getBlockedUserId() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_Comment {
	   public long CommentId;
	   public int UserId;
	   public long PostId;
	   public String Text;
	   public int LikesCount;
	   public int IsEdited;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_Comment t = new PBFlatTypes.PB_Comment();
    t.setCommentId();
    t.setUserId();
    t.setPostId();
    t.setText();
    t.setLikesCount();
    t.setIsEdited();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_Comment t = new PBFlatTypes.PB_Comment();
	t.CommentId = ;
	t.UserId = ;
	t.PostId = ;
	t.Text = ;
	t.LikesCount = ;
	t.IsEdited = ;
	t.CreatedTime = ;
	*/

	/*
	PB_Comment t = new PB_Comment();
	t.CommentId = m.getCommentId() ;
	t.UserId = m.getUserId() ;
	t.PostId = m.getPostId() ;
	t.Text = m.getText() ;
	t.LikesCount = m.getLikesCount() ;
	t.IsEdited = m.getIsEdited() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_CommentDeleted {
	   public long CommentId;
	   public int UserId;
	}
	/*
	folding
	PBFlatTypes.PB_CommentDeleted t = new PBFlatTypes.PB_CommentDeleted();
    t.setCommentId();
    t.setUserId();
	*/

	/*
	PBFlatTypes.PB_CommentDeleted t = new PBFlatTypes.PB_CommentDeleted();
	t.CommentId = ;
	t.UserId = ;
	*/

	/*
	PB_CommentDeleted t = new PB_CommentDeleted();
	t.CommentId = m.getCommentId() ;
	t.UserId = m.getUserId() ;
	*/

	public class PB_Event {
	   public long EventId;
	   public int EventType;
	   public long ByUserId;
	   public long PeerUserId;
	   public long PostId;
	   public long CommentId;
	   public long ActionId;
	   public long Murmur64Hash;
	   public String ChatKey;
	   public long MessageId;
	   public long ReSharedId;
	}
	/*
	folding
	PBFlatTypes.PB_Event t = new PBFlatTypes.PB_Event();
    t.setEventId();
    t.setEventType();
    t.setByUserId();
    t.setPeerUserId();
    t.setPostId();
    t.setCommentId();
    t.setActionId();
    t.setMurmur64Hash();
    t.setChatKey();
    t.setMessageId();
    t.setReSharedId();
	*/

	/*
	PBFlatTypes.PB_Event t = new PBFlatTypes.PB_Event();
	t.EventId = ;
	t.EventType = ;
	t.ByUserId = ;
	t.PeerUserId = ;
	t.PostId = ;
	t.CommentId = ;
	t.ActionId = ;
	t.Murmur64Hash = ;
	t.ChatKey = ;
	t.MessageId = ;
	t.ReSharedId = ;
	*/

	/*
	PB_Event t = new PB_Event();
	t.EventId = m.getEventId() ;
	t.EventType = m.getEventType() ;
	t.ByUserId = m.getByUserId() ;
	t.PeerUserId = m.getPeerUserId() ;
	t.PostId = m.getPostId() ;
	t.CommentId = m.getCommentId() ;
	t.ActionId = m.getActionId() ;
	t.Murmur64Hash = m.getMurmur64Hash() ;
	t.ChatKey = m.getChatKey() ;
	t.MessageId = m.getMessageId() ;
	t.ReSharedId = m.getReSharedId() ;
	*/

	public class PB_Followed {
	   public long Id;
	   public int UserId;
	   public int FollowedUserId;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_Followed t = new PBFlatTypes.PB_Followed();
    t.setId();
    t.setUserId();
    t.setFollowedUserId();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_Followed t = new PBFlatTypes.PB_Followed();
	t.Id = ;
	t.UserId = ;
	t.FollowedUserId = ;
	t.CreatedTime = ;
	*/

	/*
	PB_Followed t = new PB_Followed();
	t.Id = m.getId() ;
	t.UserId = m.getUserId() ;
	t.FollowedUserId = m.getFollowedUserId() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_Like {
	   public long Id;
	   public long PostId;
	   public int PostTypeEnum;
	   public int UserId;
	   public int LikeEnum;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_Like t = new PBFlatTypes.PB_Like();
    t.setId();
    t.setPostId();
    t.setPostTypeEnum();
    t.setUserId();
    t.setLikeEnum();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_Like t = new PBFlatTypes.PB_Like();
	t.Id = ;
	t.PostId = ;
	t.PostTypeEnum = ;
	t.UserId = ;
	t.LikeEnum = ;
	t.CreatedTime = ;
	*/

	/*
	PB_Like t = new PB_Like();
	t.Id = m.getId() ;
	t.PostId = m.getPostId() ;
	t.PostTypeEnum = m.getPostTypeEnum() ;
	t.UserId = m.getUserId() ;
	t.LikeEnum = m.getLikeEnum() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_Notify {
	   public long NotifyId;
	   public int ForUserId;
	   public int ActorUserId;
	   public int NotifyType;
	   public long PostId;
	   public long CommentId;
	   public int PeerUserId;
	   public long Murmur64Hash;
	   public int SeenStatus;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_Notify t = new PBFlatTypes.PB_Notify();
    t.setNotifyId();
    t.setForUserId();
    t.setActorUserId();
    t.setNotifyType();
    t.setPostId();
    t.setCommentId();
    t.setPeerUserId();
    t.setMurmur64Hash();
    t.setSeenStatus();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_Notify t = new PBFlatTypes.PB_Notify();
	t.NotifyId = ;
	t.ForUserId = ;
	t.ActorUserId = ;
	t.NotifyType = ;
	t.PostId = ;
	t.CommentId = ;
	t.PeerUserId = ;
	t.Murmur64Hash = ;
	t.SeenStatus = ;
	t.CreatedTime = ;
	*/

	/*
	PB_Notify t = new PB_Notify();
	t.NotifyId = m.getNotifyId() ;
	t.ForUserId = m.getForUserId() ;
	t.ActorUserId = m.getActorUserId() ;
	t.NotifyType = m.getNotifyType() ;
	t.PostId = m.getPostId() ;
	t.CommentId = m.getCommentId() ;
	t.PeerUserId = m.getPeerUserId() ;
	t.Murmur64Hash = m.getMurmur64Hash() ;
	t.SeenStatus = m.getSeenStatus() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_NotifyRemoved {
	   public long Murmur64Hash;
	   public int ForUserId;
	   public long Id;
	}
	/*
	folding
	PBFlatTypes.PB_NotifyRemoved t = new PBFlatTypes.PB_NotifyRemoved();
    t.setMurmur64Hash();
    t.setForUserId();
    t.setId();
	*/

	/*
	PBFlatTypes.PB_NotifyRemoved t = new PBFlatTypes.PB_NotifyRemoved();
	t.Murmur64Hash = ;
	t.ForUserId = ;
	t.Id = ;
	*/

	/*
	PB_NotifyRemoved t = new PB_NotifyRemoved();
	t.Murmur64Hash = m.getMurmur64Hash() ;
	t.ForUserId = m.getForUserId() ;
	t.Id = m.getId() ;
	*/

	public class PB_PhoneContact {
	   public long Id;
	   public int UserId;
	   public long ClientId;
	   public String Phone;
	   public String FirstName;
	   public String LastName;
	}
	/*
	folding
	PBFlatTypes.PB_PhoneContact t = new PBFlatTypes.PB_PhoneContact();
    t.setId();
    t.setUserId();
    t.setClientId();
    t.setPhone();
    t.setFirstName();
    t.setLastName();
	*/

	/*
	PBFlatTypes.PB_PhoneContact t = new PBFlatTypes.PB_PhoneContact();
	t.Id = ;
	t.UserId = ;
	t.ClientId = ;
	t.Phone = ;
	t.FirstName = ;
	t.LastName = ;
	*/

	/*
	PB_PhoneContact t = new PB_PhoneContact();
	t.Id = m.getId() ;
	t.UserId = m.getUserId() ;
	t.ClientId = m.getClientId() ;
	t.Phone = m.getPhone() ;
	t.FirstName = m.getFirstName() ;
	t.LastName = m.getLastName() ;
	*/

	public class PB_Post {
	   public long PostId;
	   public int UserId;
	   public int PostTypeEnum;
	   public int PostCategoryEnum;
	   public long MediaId;
	   public String PostKey;
	   public String Text;
	   public String RichText;
	   public int MediaCount;
	   public int SharedTo;
	   public int DisableComment;
	   public int Source;
	   public int HasTag;
	   public int Seq;
	   public int CommentsCount;
	   public int LikesCount;
	   public int ViewsCount;
	   public int EditedTime;
	   public int CreatedTime;
	   public long ReSharedPostId;
	}
	/*
	folding
	PBFlatTypes.PB_Post t = new PBFlatTypes.PB_Post();
    t.setPostId();
    t.setUserId();
    t.setPostTypeEnum();
    t.setPostCategoryEnum();
    t.setMediaId();
    t.setPostKey();
    t.setText();
    t.setRichText();
    t.setMediaCount();
    t.setSharedTo();
    t.setDisableComment();
    t.setSource();
    t.setHasTag();
    t.setSeq();
    t.setCommentsCount();
    t.setLikesCount();
    t.setViewsCount();
    t.setEditedTime();
    t.setCreatedTime();
    t.setReSharedPostId();
	*/

	/*
	PBFlatTypes.PB_Post t = new PBFlatTypes.PB_Post();
	t.PostId = ;
	t.UserId = ;
	t.PostTypeEnum = ;
	t.PostCategoryEnum = ;
	t.MediaId = ;
	t.PostKey = ;
	t.Text = ;
	t.RichText = ;
	t.MediaCount = ;
	t.SharedTo = ;
	t.DisableComment = ;
	t.Source = ;
	t.HasTag = ;
	t.Seq = ;
	t.CommentsCount = ;
	t.LikesCount = ;
	t.ViewsCount = ;
	t.EditedTime = ;
	t.CreatedTime = ;
	t.ReSharedPostId = ;
	*/

	/*
	PB_Post t = new PB_Post();
	t.PostId = m.getPostId() ;
	t.UserId = m.getUserId() ;
	t.PostTypeEnum = m.getPostTypeEnum() ;
	t.PostCategoryEnum = m.getPostCategoryEnum() ;
	t.MediaId = m.getMediaId() ;
	t.PostKey = m.getPostKey() ;
	t.Text = m.getText() ;
	t.RichText = m.getRichText() ;
	t.MediaCount = m.getMediaCount() ;
	t.SharedTo = m.getSharedTo() ;
	t.DisableComment = m.getDisableComment() ;
	t.Source = m.getSource() ;
	t.HasTag = m.getHasTag() ;
	t.Seq = m.getSeq() ;
	t.CommentsCount = m.getCommentsCount() ;
	t.LikesCount = m.getLikesCount() ;
	t.ViewsCount = m.getViewsCount() ;
	t.EditedTime = m.getEditedTime() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.ReSharedPostId = m.getReSharedPostId() ;
	*/

	public class PB_PostCount {
	   public long PostId;
	   public int CommentsCount;
	   public int LikesCount;
	   public long ViewsCount;
	   public int ReSharedCount;
	   public int ChatSharedCount;
	}
	/*
	folding
	PBFlatTypes.PB_PostCount t = new PBFlatTypes.PB_PostCount();
    t.setPostId();
    t.setCommentsCount();
    t.setLikesCount();
    t.setViewsCount();
    t.setReSharedCount();
    t.setChatSharedCount();
	*/

	/*
	PBFlatTypes.PB_PostCount t = new PBFlatTypes.PB_PostCount();
	t.PostId = ;
	t.CommentsCount = ;
	t.LikesCount = ;
	t.ViewsCount = ;
	t.ReSharedCount = ;
	t.ChatSharedCount = ;
	*/

	/*
	PB_PostCount t = new PB_PostCount();
	t.PostId = m.getPostId() ;
	t.CommentsCount = m.getCommentsCount() ;
	t.LikesCount = m.getLikesCount() ;
	t.ViewsCount = m.getViewsCount() ;
	t.ReSharedCount = m.getReSharedCount() ;
	t.ChatSharedCount = m.getChatSharedCount() ;
	*/

	public class PB_PostDeleted {
	   public long PostId;
	   public int UserId;
	}
	/*
	folding
	PBFlatTypes.PB_PostDeleted t = new PBFlatTypes.PB_PostDeleted();
    t.setPostId();
    t.setUserId();
	*/

	/*
	PBFlatTypes.PB_PostDeleted t = new PBFlatTypes.PB_PostDeleted();
	t.PostId = ;
	t.UserId = ;
	*/

	/*
	PB_PostDeleted t = new PB_PostDeleted();
	t.PostId = m.getPostId() ;
	t.UserId = m.getUserId() ;
	*/

	public class PB_PostKey {
	   public int Id;
	   public String PostKeyStr;
	   public int Used;
	}
	/*
	folding
	PBFlatTypes.PB_PostKey t = new PBFlatTypes.PB_PostKey();
    t.setId();
    t.setPostKeyStr();
    t.setUsed();
	*/

	/*
	PBFlatTypes.PB_PostKey t = new PBFlatTypes.PB_PostKey();
	t.Id = ;
	t.PostKeyStr = ;
	t.Used = ;
	*/

	/*
	PB_PostKey t = new PB_PostKey();
	t.Id = m.getId() ;
	t.PostKeyStr = m.getPostKeyStr() ;
	t.Used = m.getUsed() ;
	*/

	public class PB_PostLink {
	   public long LinkId;
	   public String LinkUrl;
	}
	/*
	folding
	PBFlatTypes.PB_PostLink t = new PBFlatTypes.PB_PostLink();
    t.setLinkId();
    t.setLinkUrl();
	*/

	/*
	PBFlatTypes.PB_PostLink t = new PBFlatTypes.PB_PostLink();
	t.LinkId = ;
	t.LinkUrl = ;
	*/

	/*
	PB_PostLink t = new PB_PostLink();
	t.LinkId = m.getLinkId() ;
	t.LinkUrl = m.getLinkUrl() ;
	*/

	public class PB_PostMedia {
	   public long MediaId;
	   public int UserId;
	   public long PostId;
	   public int AlbumId;
	   public int MediaTypeEnum;
	   public int Width;
	   public int Height;
	   public int Size;
	   public int Duration;
	   public String Extension;
	   public String Md5Hash;
	   public String Color;
	   public int CreatedTime;
	   public int ViewCount;
	   public String Extra;
	}
	/*
	folding
	PBFlatTypes.PB_PostMedia t = new PBFlatTypes.PB_PostMedia();
    t.setMediaId();
    t.setUserId();
    t.setPostId();
    t.setAlbumId();
    t.setMediaTypeEnum();
    t.setWidth();
    t.setHeight();
    t.setSize();
    t.setDuration();
    t.setExtension();
    t.setMd5Hash();
    t.setColor();
    t.setCreatedTime();
    t.setViewCount();
    t.setExtra();
	*/

	/*
	PBFlatTypes.PB_PostMedia t = new PBFlatTypes.PB_PostMedia();
	t.MediaId = ;
	t.UserId = ;
	t.PostId = ;
	t.AlbumId = ;
	t.MediaTypeEnum = ;
	t.Width = ;
	t.Height = ;
	t.Size = ;
	t.Duration = ;
	t.Extension = ;
	t.Md5Hash = ;
	t.Color = ;
	t.CreatedTime = ;
	t.ViewCount = ;
	t.Extra = ;
	*/

	/*
	PB_PostMedia t = new PB_PostMedia();
	t.MediaId = m.getMediaId() ;
	t.UserId = m.getUserId() ;
	t.PostId = m.getPostId() ;
	t.AlbumId = m.getAlbumId() ;
	t.MediaTypeEnum = m.getMediaTypeEnum() ;
	t.Width = m.getWidth() ;
	t.Height = m.getHeight() ;
	t.Size = m.getSize() ;
	t.Duration = m.getDuration() ;
	t.Extension = m.getExtension() ;
	t.Md5Hash = m.getMd5Hash() ;
	t.Color = m.getColor() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.ViewCount = m.getViewCount() ;
	t.Extra = m.getExtra() ;
	*/

	public class PB_PostMentioned {
	   public long MentionedId;
	   public int ForUserId;
	   public long PostId;
	   public int PostUserId;
	   public int PostType;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_PostMentioned t = new PBFlatTypes.PB_PostMentioned();
    t.setMentionedId();
    t.setForUserId();
    t.setPostId();
    t.setPostUserId();
    t.setPostType();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_PostMentioned t = new PBFlatTypes.PB_PostMentioned();
	t.MentionedId = ;
	t.ForUserId = ;
	t.PostId = ;
	t.PostUserId = ;
	t.PostType = ;
	t.CreatedTime = ;
	*/

	/*
	PB_PostMentioned t = new PB_PostMentioned();
	t.MentionedId = m.getMentionedId() ;
	t.ForUserId = m.getForUserId() ;
	t.PostId = m.getPostId() ;
	t.PostUserId = m.getPostUserId() ;
	t.PostType = m.getPostType() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_PostReshared {
	   public long ResharedId;
	   public long PostId;
	   public int ByUserId;
	   public int PostUserId;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_PostReshared t = new PBFlatTypes.PB_PostReshared();
    t.setResharedId();
    t.setPostId();
    t.setByUserId();
    t.setPostUserId();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_PostReshared t = new PBFlatTypes.PB_PostReshared();
	t.ResharedId = ;
	t.PostId = ;
	t.ByUserId = ;
	t.PostUserId = ;
	t.CreatedTime = ;
	*/

	/*
	PB_PostReshared t = new PB_PostReshared();
	t.ResharedId = m.getResharedId() ;
	t.PostId = m.getPostId() ;
	t.ByUserId = m.getByUserId() ;
	t.PostUserId = m.getPostUserId() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_Session {
	   public long Id;
	   public String SessionUuid;
	   public int UserId;
	   public String LastIpAddress;
	   public int AppVersion;
	   public int ActiveTime;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_Session t = new PBFlatTypes.PB_Session();
    t.setId();
    t.setSessionUuid();
    t.setUserId();
    t.setLastIpAddress();
    t.setAppVersion();
    t.setActiveTime();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_Session t = new PBFlatTypes.PB_Session();
	t.Id = ;
	t.SessionUuid = ;
	t.UserId = ;
	t.LastIpAddress = ;
	t.AppVersion = ;
	t.ActiveTime = ;
	t.CreatedTime = ;
	*/

	/*
	PB_Session t = new PB_Session();
	t.Id = m.getId() ;
	t.SessionUuid = m.getSessionUuid() ;
	t.UserId = m.getUserId() ;
	t.LastIpAddress = m.getLastIpAddress() ;
	t.AppVersion = m.getAppVersion() ;
	t.ActiveTime = m.getActiveTime() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_SettingNotification {
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
	/*
	folding
	PBFlatTypes.PB_SettingNotification t = new PBFlatTypes.PB_SettingNotification();
    t.setUserId();
    t.setSocialLedOn();
    t.setSocialLedColor();
    t.setReqestToFollowYou();
    t.setFollowedYou();
    t.setAccptedYourFollowRequest();
    t.setYourPostLiked();
    t.setYourPostCommented();
    t.setMenthenedYouInPost();
    t.setMenthenedYouInComment();
    t.setYourContactsJoined();
    t.setDirectMessage();
    t.setDirectAlert();
    t.setDirectPerview();
    t.setDirectLedOn();
    t.setDirectLedColor();
    t.setDirectVibrate();
    t.setDirectPopup();
    t.setDirectSound();
    t.setDirectPriority();
	*/

	/*
	PBFlatTypes.PB_SettingNotification t = new PBFlatTypes.PB_SettingNotification();
	t.UserId = ;
	t.SocialLedOn = ;
	t.SocialLedColor = ;
	t.ReqestToFollowYou = ;
	t.FollowedYou = ;
	t.AccptedYourFollowRequest = ;
	t.YourPostLiked = ;
	t.YourPostCommented = ;
	t.MenthenedYouInPost = ;
	t.MenthenedYouInComment = ;
	t.YourContactsJoined = ;
	t.DirectMessage = ;
	t.DirectAlert = ;
	t.DirectPerview = ;
	t.DirectLedOn = ;
	t.DirectLedColor = ;
	t.DirectVibrate = ;
	t.DirectPopup = ;
	t.DirectSound = ;
	t.DirectPriority = ;
	*/

	/*
	PB_SettingNotification t = new PB_SettingNotification();
	t.UserId = m.getUserId() ;
	t.SocialLedOn = m.getSocialLedOn() ;
	t.SocialLedColor = m.getSocialLedColor() ;
	t.ReqestToFollowYou = m.getReqestToFollowYou() ;
	t.FollowedYou = m.getFollowedYou() ;
	t.AccptedYourFollowRequest = m.getAccptedYourFollowRequest() ;
	t.YourPostLiked = m.getYourPostLiked() ;
	t.YourPostCommented = m.getYourPostCommented() ;
	t.MenthenedYouInPost = m.getMenthenedYouInPost() ;
	t.MenthenedYouInComment = m.getMenthenedYouInComment() ;
	t.YourContactsJoined = m.getYourContactsJoined() ;
	t.DirectMessage = m.getDirectMessage() ;
	t.DirectAlert = m.getDirectAlert() ;
	t.DirectPerview = m.getDirectPerview() ;
	t.DirectLedOn = m.getDirectLedOn() ;
	t.DirectLedColor = m.getDirectLedColor() ;
	t.DirectVibrate = m.getDirectVibrate() ;
	t.DirectPopup = m.getDirectPopup() ;
	t.DirectSound = m.getDirectSound() ;
	t.DirectPriority = m.getDirectPriority() ;
	*/

	public class PB_Tag {
	   public long TagId;
	   public String Name;
	   public int Count;
	   public int TagStatusEnum;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_Tag t = new PBFlatTypes.PB_Tag();
    t.setTagId();
    t.setName();
    t.setCount();
    t.setTagStatusEnum();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_Tag t = new PBFlatTypes.PB_Tag();
	t.TagId = ;
	t.Name = ;
	t.Count = ;
	t.TagStatusEnum = ;
	t.CreatedTime = ;
	*/

	/*
	PB_Tag t = new PB_Tag();
	t.TagId = m.getTagId() ;
	t.Name = m.getName() ;
	t.Count = m.getCount() ;
	t.TagStatusEnum = m.getTagStatusEnum() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_TagPost {
	   public long Id;
	   public int TagId;
	   public int PostId;
	   public int PostTypeEnum;
	   public int PostCategoryEnum;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_TagPost t = new PBFlatTypes.PB_TagPost();
    t.setId();
    t.setTagId();
    t.setPostId();
    t.setPostTypeEnum();
    t.setPostCategoryEnum();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_TagPost t = new PBFlatTypes.PB_TagPost();
	t.Id = ;
	t.TagId = ;
	t.PostId = ;
	t.PostTypeEnum = ;
	t.PostCategoryEnum = ;
	t.CreatedTime = ;
	*/

	/*
	PB_TagPost t = new PB_TagPost();
	t.Id = m.getId() ;
	t.TagId = m.getTagId() ;
	t.PostId = m.getPostId() ;
	t.PostTypeEnum = m.getPostTypeEnum() ;
	t.PostCategoryEnum = m.getPostCategoryEnum() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_TriggerLog {
	   public long Id;
	   public String ModelName;
	   public String ChangeType;
	   public long TargetId;
	   public String TargetStr;
	   public int CreatedSe;
	}
	/*
	folding
	PBFlatTypes.PB_TriggerLog t = new PBFlatTypes.PB_TriggerLog();
    t.setId();
    t.setModelName();
    t.setChangeType();
    t.setTargetId();
    t.setTargetStr();
    t.setCreatedSe();
	*/

	/*
	PBFlatTypes.PB_TriggerLog t = new PBFlatTypes.PB_TriggerLog();
	t.Id = ;
	t.ModelName = ;
	t.ChangeType = ;
	t.TargetId = ;
	t.TargetStr = ;
	t.CreatedSe = ;
	*/

	/*
	PB_TriggerLog t = new PB_TriggerLog();
	t.Id = m.getId() ;
	t.ModelName = m.getModelName() ;
	t.ChangeType = m.getChangeType() ;
	t.TargetId = m.getTargetId() ;
	t.TargetStr = m.getTargetStr() ;
	t.CreatedSe = m.getCreatedSe() ;
	*/

	public class PB_User {
	   public int UserId;
	   public String UserName;
	   public String UserNameLower;
	   public String FirstName;
	   public String LastName;
	   public int IsVerified;
	   public long AvatarId;
	   public int ProfilePrivacy;
	   public int OnlinePrivacy;
	   public int CallPrivacy;
	   public int AddToGroupPrivacy;
	   public int SeenMessagePrivacy;
	   public long Phone;
	   public String Email;
	   public String About;
	   public String PasswordHash;
	   public String PasswordSalt;
	   public int PostSeq;
	   public int FollowersCount;
	   public int FollowingCount;
	   public int PostsCount;
	   public int MediaCount;
	   public int PhotoCount;
	   public int VideoCount;
	   public int GifCount;
	   public int AudioCount;
	   public int VoiceCount;
	   public int FileCount;
	   public int LinkCount;
	   public int BoardCount;
	   public int PinedCount;
	   public int LikesCount;
	   public int ResharedCount;
	   public int LastPostTime;
	   public int CreatedTime;
	   public int VersionTime;
	   public int IsDeleted;
	   public int IsBanned;
	}
	/*
	folding
	PBFlatTypes.PB_User t = new PBFlatTypes.PB_User();
    t.setUserId();
    t.setUserName();
    t.setUserNameLower();
    t.setFirstName();
    t.setLastName();
    t.setIsVerified();
    t.setAvatarId();
    t.setProfilePrivacy();
    t.setOnlinePrivacy();
    t.setCallPrivacy();
    t.setAddToGroupPrivacy();
    t.setSeenMessagePrivacy();
    t.setPhone();
    t.setEmail();
    t.setAbout();
    t.setPasswordHash();
    t.setPasswordSalt();
    t.setPostSeq();
    t.setFollowersCount();
    t.setFollowingCount();
    t.setPostsCount();
    t.setMediaCount();
    t.setPhotoCount();
    t.setVideoCount();
    t.setGifCount();
    t.setAudioCount();
    t.setVoiceCount();
    t.setFileCount();
    t.setLinkCount();
    t.setBoardCount();
    t.setPinedCount();
    t.setLikesCount();
    t.setResharedCount();
    t.setLastPostTime();
    t.setCreatedTime();
    t.setVersionTime();
    t.setIsDeleted();
    t.setIsBanned();
	*/

	/*
	PBFlatTypes.PB_User t = new PBFlatTypes.PB_User();
	t.UserId = ;
	t.UserName = ;
	t.UserNameLower = ;
	t.FirstName = ;
	t.LastName = ;
	t.IsVerified = ;
	t.AvatarId = ;
	t.ProfilePrivacy = ;
	t.OnlinePrivacy = ;
	t.CallPrivacy = ;
	t.AddToGroupPrivacy = ;
	t.SeenMessagePrivacy = ;
	t.Phone = ;
	t.Email = ;
	t.About = ;
	t.PasswordHash = ;
	t.PasswordSalt = ;
	t.PostSeq = ;
	t.FollowersCount = ;
	t.FollowingCount = ;
	t.PostsCount = ;
	t.MediaCount = ;
	t.PhotoCount = ;
	t.VideoCount = ;
	t.GifCount = ;
	t.AudioCount = ;
	t.VoiceCount = ;
	t.FileCount = ;
	t.LinkCount = ;
	t.BoardCount = ;
	t.PinedCount = ;
	t.LikesCount = ;
	t.ResharedCount = ;
	t.LastPostTime = ;
	t.CreatedTime = ;
	t.VersionTime = ;
	t.IsDeleted = ;
	t.IsBanned = ;
	*/

	/*
	PB_User t = new PB_User();
	t.UserId = m.getUserId() ;
	t.UserName = m.getUserName() ;
	t.UserNameLower = m.getUserNameLower() ;
	t.FirstName = m.getFirstName() ;
	t.LastName = m.getLastName() ;
	t.IsVerified = m.getIsVerified() ;
	t.AvatarId = m.getAvatarId() ;
	t.ProfilePrivacy = m.getProfilePrivacy() ;
	t.OnlinePrivacy = m.getOnlinePrivacy() ;
	t.CallPrivacy = m.getCallPrivacy() ;
	t.AddToGroupPrivacy = m.getAddToGroupPrivacy() ;
	t.SeenMessagePrivacy = m.getSeenMessagePrivacy() ;
	t.Phone = m.getPhone() ;
	t.Email = m.getEmail() ;
	t.About = m.getAbout() ;
	t.PasswordHash = m.getPasswordHash() ;
	t.PasswordSalt = m.getPasswordSalt() ;
	t.PostSeq = m.getPostSeq() ;
	t.FollowersCount = m.getFollowersCount() ;
	t.FollowingCount = m.getFollowingCount() ;
	t.PostsCount = m.getPostsCount() ;
	t.MediaCount = m.getMediaCount() ;
	t.PhotoCount = m.getPhotoCount() ;
	t.VideoCount = m.getVideoCount() ;
	t.GifCount = m.getGifCount() ;
	t.AudioCount = m.getAudioCount() ;
	t.VoiceCount = m.getVoiceCount() ;
	t.FileCount = m.getFileCount() ;
	t.LinkCount = m.getLinkCount() ;
	t.BoardCount = m.getBoardCount() ;
	t.PinedCount = m.getPinedCount() ;
	t.LikesCount = m.getLikesCount() ;
	t.ResharedCount = m.getResharedCount() ;
	t.LastPostTime = m.getLastPostTime() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.VersionTime = m.getVersionTime() ;
	t.IsDeleted = m.getIsDeleted() ;
	t.IsBanned = m.getIsBanned() ;
	*/

	public class PB_UserRelation {
	   public long RelNanoId;
	   public int UserId;
	   public int PeerUserId;
	   public int Follwing;
	   public int Followed;
	   public int InContacts;
	   public int MutualContact;
	   public int IsFavorite;
	   public int Notify;
	}
	/*
	folding
	PBFlatTypes.PB_UserRelation t = new PBFlatTypes.PB_UserRelation();
    t.setRelNanoId();
    t.setUserId();
    t.setPeerUserId();
    t.setFollwing();
    t.setFollowed();
    t.setInContacts();
    t.setMutualContact();
    t.setIsFavorite();
    t.setNotify();
	*/

	/*
	PBFlatTypes.PB_UserRelation t = new PBFlatTypes.PB_UserRelation();
	t.RelNanoId = ;
	t.UserId = ;
	t.PeerUserId = ;
	t.Follwing = ;
	t.Followed = ;
	t.InContacts = ;
	t.MutualContact = ;
	t.IsFavorite = ;
	t.Notify = ;
	*/

	/*
	PB_UserRelation t = new PB_UserRelation();
	t.RelNanoId = m.getRelNanoId() ;
	t.UserId = m.getUserId() ;
	t.PeerUserId = m.getPeerUserId() ;
	t.Follwing = m.getFollwing() ;
	t.Followed = m.getFollowed() ;
	t.InContacts = m.getInContacts() ;
	t.MutualContact = m.getMutualContact() ;
	t.IsFavorite = m.getIsFavorite() ;
	t.Notify = m.getNotify() ;
	*/

	public class PB_Chat {
	   public long ChatId;
	   public String ChatKey;
	   public String RoomKey;
	   public int RoomType;
	   public int UserId;
	   public int PeerUserId;
	   public long GroupId;
	   public long HashTagId;
	   public String Title;
	   public long PinTimeMs;
	   public long FromMsgId;
	   public int UnseenCount;
	   public int Seq;
	   public long LastMsgId;
	   public int LastMyMsgStatus;
	   public int MyLastSeenSeq;
	   public long MyLastSeenMsgId;
	   public long PeerLastSeenMsgId;
	   public int MyLastDeliveredSeq;
	   public long MyLastDeliveredMsgId;
	   public long PeerLastDeliveredMsgId;
	   public int IsActive;
	   public int IsLeft;
	   public int IsCreator;
	   public int IsKicked;
	   public int IsAdmin;
	   public int IsDeactivated;
	   public int IsMuted;
	   public int MuteUntil;
	   public long VersionTimeMs;
	   public int VersionSeq;
	   public int OrderTime;
	   public int CreatedTime;
	   public String DraftText;
	   public long DratReplyToMsgId;
	}
	/*
	folding
	PBFlatTypes.PB_Chat t = new PBFlatTypes.PB_Chat();
    t.setChatId();
    t.setChatKey();
    t.setRoomKey();
    t.setRoomType();
    t.setUserId();
    t.setPeerUserId();
    t.setGroupId();
    t.setHashTagId();
    t.setTitle();
    t.setPinTimeMs();
    t.setFromMsgId();
    t.setUnseenCount();
    t.setSeq();
    t.setLastMsgId();
    t.setLastMyMsgStatus();
    t.setMyLastSeenSeq();
    t.setMyLastSeenMsgId();
    t.setPeerLastSeenMsgId();
    t.setMyLastDeliveredSeq();
    t.setMyLastDeliveredMsgId();
    t.setPeerLastDeliveredMsgId();
    t.setIsActive();
    t.setIsLeft();
    t.setIsCreator();
    t.setIsKicked();
    t.setIsAdmin();
    t.setIsDeactivated();
    t.setIsMuted();
    t.setMuteUntil();
    t.setVersionTimeMs();
    t.setVersionSeq();
    t.setOrderTime();
    t.setCreatedTime();
    t.setDraftText();
    t.setDratReplyToMsgId();
	*/

	/*
	PBFlatTypes.PB_Chat t = new PBFlatTypes.PB_Chat();
	t.ChatId = ;
	t.ChatKey = ;
	t.RoomKey = ;
	t.RoomType = ;
	t.UserId = ;
	t.PeerUserId = ;
	t.GroupId = ;
	t.HashTagId = ;
	t.Title = ;
	t.PinTimeMs = ;
	t.FromMsgId = ;
	t.UnseenCount = ;
	t.Seq = ;
	t.LastMsgId = ;
	t.LastMyMsgStatus = ;
	t.MyLastSeenSeq = ;
	t.MyLastSeenMsgId = ;
	t.PeerLastSeenMsgId = ;
	t.MyLastDeliveredSeq = ;
	t.MyLastDeliveredMsgId = ;
	t.PeerLastDeliveredMsgId = ;
	t.IsActive = ;
	t.IsLeft = ;
	t.IsCreator = ;
	t.IsKicked = ;
	t.IsAdmin = ;
	t.IsDeactivated = ;
	t.IsMuted = ;
	t.MuteUntil = ;
	t.VersionTimeMs = ;
	t.VersionSeq = ;
	t.OrderTime = ;
	t.CreatedTime = ;
	t.DraftText = ;
	t.DratReplyToMsgId = ;
	*/

	/*
	PB_Chat t = new PB_Chat();
	t.ChatId = m.getChatId() ;
	t.ChatKey = m.getChatKey() ;
	t.RoomKey = m.getRoomKey() ;
	t.RoomType = m.getRoomType() ;
	t.UserId = m.getUserId() ;
	t.PeerUserId = m.getPeerUserId() ;
	t.GroupId = m.getGroupId() ;
	t.HashTagId = m.getHashTagId() ;
	t.Title = m.getTitle() ;
	t.PinTimeMs = m.getPinTimeMs() ;
	t.FromMsgId = m.getFromMsgId() ;
	t.UnseenCount = m.getUnseenCount() ;
	t.Seq = m.getSeq() ;
	t.LastMsgId = m.getLastMsgId() ;
	t.LastMyMsgStatus = m.getLastMyMsgStatus() ;
	t.MyLastSeenSeq = m.getMyLastSeenSeq() ;
	t.MyLastSeenMsgId = m.getMyLastSeenMsgId() ;
	t.PeerLastSeenMsgId = m.getPeerLastSeenMsgId() ;
	t.MyLastDeliveredSeq = m.getMyLastDeliveredSeq() ;
	t.MyLastDeliveredMsgId = m.getMyLastDeliveredMsgId() ;
	t.PeerLastDeliveredMsgId = m.getPeerLastDeliveredMsgId() ;
	t.IsActive = m.getIsActive() ;
	t.IsLeft = m.getIsLeft() ;
	t.IsCreator = m.getIsCreator() ;
	t.IsKicked = m.getIsKicked() ;
	t.IsAdmin = m.getIsAdmin() ;
	t.IsDeactivated = m.getIsDeactivated() ;
	t.IsMuted = m.getIsMuted() ;
	t.MuteUntil = m.getMuteUntil() ;
	t.VersionTimeMs = m.getVersionTimeMs() ;
	t.VersionSeq = m.getVersionSeq() ;
	t.OrderTime = m.getOrderTime() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.DraftText = m.getDraftText() ;
	t.DratReplyToMsgId = m.getDratReplyToMsgId() ;
	*/

	public class PB_ChatDeleted {
	   public long ChatId;
	   public String RoomKey;
	}
	/*
	folding
	PBFlatTypes.PB_ChatDeleted t = new PBFlatTypes.PB_ChatDeleted();
    t.setChatId();
    t.setRoomKey();
	*/

	/*
	PBFlatTypes.PB_ChatDeleted t = new PBFlatTypes.PB_ChatDeleted();
	t.ChatId = ;
	t.RoomKey = ;
	*/

	/*
	PB_ChatDeleted t = new PB_ChatDeleted();
	t.ChatId = m.getChatId() ;
	t.RoomKey = m.getRoomKey() ;
	*/

	public class PB_ChatLastMessage {
	   public String ChatIdGroupId;
	   public byte[] LastMsgPb;
	}
	/*
	folding
	PBFlatTypes.PB_ChatLastMessage t = new PBFlatTypes.PB_ChatLastMessage();
    t.setChatIdGroupId();
    t.setLastMsgPb();
	*/

	/*
	PBFlatTypes.PB_ChatLastMessage t = new PBFlatTypes.PB_ChatLastMessage();
	t.ChatIdGroupId = ;
	t.LastMsgPb = ;
	*/

	/*
	PB_ChatLastMessage t = new PB_ChatLastMessage();
	t.ChatIdGroupId = m.getChatIdGroupId() ;
	t.LastMsgPb = m.getLastMsgPb() ;
	*/

	public class PB_ChatUserVersion {
	   public int UserId;
	   public long ChatId;
	   public int VersionTimeNano;
	}
	/*
	folding
	PBFlatTypes.PB_ChatUserVersion t = new PBFlatTypes.PB_ChatUserVersion();
    t.setUserId();
    t.setChatId();
    t.setVersionTimeNano();
	*/

	/*
	PBFlatTypes.PB_ChatUserVersion t = new PBFlatTypes.PB_ChatUserVersion();
	t.UserId = ;
	t.ChatId = ;
	t.VersionTimeNano = ;
	*/

	/*
	PB_ChatUserVersion t = new PB_ChatUserVersion();
	t.UserId = m.getUserId() ;
	t.ChatId = m.getChatId() ;
	t.VersionTimeNano = m.getVersionTimeNano() ;
	*/

	public class PB_Group {
	   public long GroupId;
	   public String GroupKey;
	   public String GroupName;
	   public String UserName;
	   public int IsSuperGroup;
	   public long HashTagId;
	   public int CreatorUserId;
	   public int GroupPrivacy;
	   public int HistoryViewAble;
	   public long Seq;
	   public long LastMsgId;
	   public long PinedMsgId;
	   public long AvatarRefId;
	   public int AvatarCount;
	   public String About;
	   public String InviteLinkHash;
	   public int MembersCount;
	   public int AdminsCount;
	   public int ModeratorCounts;
	   public int SortTime;
	   public int CreatedTime;
	   public int IsMute;
	   public int IsDeleted;
	   public int IsBanned;
	}
	/*
	folding
	PBFlatTypes.PB_Group t = new PBFlatTypes.PB_Group();
    t.setGroupId();
    t.setGroupKey();
    t.setGroupName();
    t.setUserName();
    t.setIsSuperGroup();
    t.setHashTagId();
    t.setCreatorUserId();
    t.setGroupPrivacy();
    t.setHistoryViewAble();
    t.setSeq();
    t.setLastMsgId();
    t.setPinedMsgId();
    t.setAvatarRefId();
    t.setAvatarCount();
    t.setAbout();
    t.setInviteLinkHash();
    t.setMembersCount();
    t.setAdminsCount();
    t.setModeratorCounts();
    t.setSortTime();
    t.setCreatedTime();
    t.setIsMute();
    t.setIsDeleted();
    t.setIsBanned();
	*/

	/*
	PBFlatTypes.PB_Group t = new PBFlatTypes.PB_Group();
	t.GroupId = ;
	t.GroupKey = ;
	t.GroupName = ;
	t.UserName = ;
	t.IsSuperGroup = ;
	t.HashTagId = ;
	t.CreatorUserId = ;
	t.GroupPrivacy = ;
	t.HistoryViewAble = ;
	t.Seq = ;
	t.LastMsgId = ;
	t.PinedMsgId = ;
	t.AvatarRefId = ;
	t.AvatarCount = ;
	t.About = ;
	t.InviteLinkHash = ;
	t.MembersCount = ;
	t.AdminsCount = ;
	t.ModeratorCounts = ;
	t.SortTime = ;
	t.CreatedTime = ;
	t.IsMute = ;
	t.IsDeleted = ;
	t.IsBanned = ;
	*/

	/*
	PB_Group t = new PB_Group();
	t.GroupId = m.getGroupId() ;
	t.GroupKey = m.getGroupKey() ;
	t.GroupName = m.getGroupName() ;
	t.UserName = m.getUserName() ;
	t.IsSuperGroup = m.getIsSuperGroup() ;
	t.HashTagId = m.getHashTagId() ;
	t.CreatorUserId = m.getCreatorUserId() ;
	t.GroupPrivacy = m.getGroupPrivacy() ;
	t.HistoryViewAble = m.getHistoryViewAble() ;
	t.Seq = m.getSeq() ;
	t.LastMsgId = m.getLastMsgId() ;
	t.PinedMsgId = m.getPinedMsgId() ;
	t.AvatarRefId = m.getAvatarRefId() ;
	t.AvatarCount = m.getAvatarCount() ;
	t.About = m.getAbout() ;
	t.InviteLinkHash = m.getInviteLinkHash() ;
	t.MembersCount = m.getMembersCount() ;
	t.AdminsCount = m.getAdminsCount() ;
	t.ModeratorCounts = m.getModeratorCounts() ;
	t.SortTime = m.getSortTime() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.IsMute = m.getIsMute() ;
	t.IsDeleted = m.getIsDeleted() ;
	t.IsBanned = m.getIsBanned() ;
	*/

	public class PB_GroupMember {
	   public long OrderId;
	   public long GroupId;
	   public int UserId;
	   public int ByUserId;
	   public int GroupRole;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_GroupMember t = new PBFlatTypes.PB_GroupMember();
    t.setOrderId();
    t.setGroupId();
    t.setUserId();
    t.setByUserId();
    t.setGroupRole();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_GroupMember t = new PBFlatTypes.PB_GroupMember();
	t.OrderId = ;
	t.GroupId = ;
	t.UserId = ;
	t.ByUserId = ;
	t.GroupRole = ;
	t.CreatedTime = ;
	*/

	/*
	PB_GroupMember t = new PB_GroupMember();
	t.OrderId = m.getOrderId() ;
	t.GroupId = m.getGroupId() ;
	t.UserId = m.getUserId() ;
	t.ByUserId = m.getByUserId() ;
	t.GroupRole = m.getGroupRole() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_GroupOrderdUser {
	   public long OrderId;
	   public long GroupId;
	   public int UserId;
	}
	/*
	folding
	PBFlatTypes.PB_GroupOrderdUser t = new PBFlatTypes.PB_GroupOrderdUser();
    t.setOrderId();
    t.setGroupId();
    t.setUserId();
	*/

	/*
	PBFlatTypes.PB_GroupOrderdUser t = new PBFlatTypes.PB_GroupOrderdUser();
	t.OrderId = ;
	t.GroupId = ;
	t.UserId = ;
	*/

	/*
	PB_GroupOrderdUser t = new PB_GroupOrderdUser();
	t.OrderId = m.getOrderId() ;
	t.GroupId = m.getGroupId() ;
	t.UserId = m.getUserId() ;
	*/

	public class PB_GroupPinedMsg {
	   public long MsgId;
	   public byte[] MsgPb;
	}
	/*
	folding
	PBFlatTypes.PB_GroupPinedMsg t = new PBFlatTypes.PB_GroupPinedMsg();
    t.setMsgId();
    t.setMsgPb();
	*/

	/*
	PBFlatTypes.PB_GroupPinedMsg t = new PBFlatTypes.PB_GroupPinedMsg();
	t.MsgId = ;
	t.MsgPb = ;
	*/

	/*
	PB_GroupPinedMsg t = new PB_GroupPinedMsg();
	t.MsgId = m.getMsgId() ;
	t.MsgPb = m.getMsgPb() ;
	*/

	public class PB_FileMsg {
	   public long Id;
	   public int AccessHash;
	   public int FileType;
	   public int Width;
	   public int Height;
	   public String Extension;
	   public int UserId;
	   public byte[] DataThumb;
	   public byte[] Data;
	}
	/*
	folding
	PBFlatTypes.PB_FileMsg t = new PBFlatTypes.PB_FileMsg();
    t.setId();
    t.setAccessHash();
    t.setFileType();
    t.setWidth();
    t.setHeight();
    t.setExtension();
    t.setUserId();
    t.setDataThumb();
    t.setData();
	*/

	/*
	PBFlatTypes.PB_FileMsg t = new PBFlatTypes.PB_FileMsg();
	t.Id = ;
	t.AccessHash = ;
	t.FileType = ;
	t.Width = ;
	t.Height = ;
	t.Extension = ;
	t.UserId = ;
	t.DataThumb = ;
	t.Data = ;
	*/

	/*
	PB_FileMsg t = new PB_FileMsg();
	t.Id = m.getId() ;
	t.AccessHash = m.getAccessHash() ;
	t.FileType = m.getFileType() ;
	t.Width = m.getWidth() ;
	t.Height = m.getHeight() ;
	t.Extension = m.getExtension() ;
	t.UserId = m.getUserId() ;
	t.DataThumb = m.getDataThumb() ;
	t.Data = m.getData() ;
	*/

	public class PB_FilePost {
	   public long Id;
	   public int AccessHash;
	   public int FileType;
	   public int Width;
	   public int Height;
	   public String Extension;
	   public int UserId;
	   public byte[] DataThumb;
	   public byte[] Data;
	}
	/*
	folding
	PBFlatTypes.PB_FilePost t = new PBFlatTypes.PB_FilePost();
    t.setId();
    t.setAccessHash();
    t.setFileType();
    t.setWidth();
    t.setHeight();
    t.setExtension();
    t.setUserId();
    t.setDataThumb();
    t.setData();
	*/

	/*
	PBFlatTypes.PB_FilePost t = new PBFlatTypes.PB_FilePost();
	t.Id = ;
	t.AccessHash = ;
	t.FileType = ;
	t.Width = ;
	t.Height = ;
	t.Extension = ;
	t.UserId = ;
	t.DataThumb = ;
	t.Data = ;
	*/

	/*
	PB_FilePost t = new PB_FilePost();
	t.Id = m.getId() ;
	t.AccessHash = m.getAccessHash() ;
	t.FileType = m.getFileType() ;
	t.Width = m.getWidth() ;
	t.Height = m.getHeight() ;
	t.Extension = m.getExtension() ;
	t.UserId = m.getUserId() ;
	t.DataThumb = m.getDataThumb() ;
	t.Data = m.getData() ;
	*/

	public class PB_ActionFanout {
	   public long OrderId;
	   public int ForUserId;
	   public long ActionId;
	   public int ActorUserId;
	}
	/*
	folding
	PBFlatTypes.PB_ActionFanout t = new PBFlatTypes.PB_ActionFanout();
    t.setOrderId();
    t.setForUserId();
    t.setActionId();
    t.setActorUserId();
	*/

	/*
	PBFlatTypes.PB_ActionFanout t = new PBFlatTypes.PB_ActionFanout();
	t.OrderId = ;
	t.ForUserId = ;
	t.ActionId = ;
	t.ActorUserId = ;
	*/

	/*
	PB_ActionFanout t = new PB_ActionFanout();
	t.OrderId = m.getOrderId() ;
	t.ForUserId = m.getForUserId() ;
	t.ActionId = m.getActionId() ;
	t.ActorUserId = m.getActorUserId() ;
	*/

	public class PB_HomeFanout {
	   public long OrderId;
	   public long ForUserId;
	   public long PostId;
	   public long PostUserId;
	   public long ResharedId;
	}
	/*
	folding
	PBFlatTypes.PB_HomeFanout t = new PBFlatTypes.PB_HomeFanout();
    t.setOrderId();
    t.setForUserId();
    t.setPostId();
    t.setPostUserId();
    t.setResharedId();
	*/

	/*
	PBFlatTypes.PB_HomeFanout t = new PBFlatTypes.PB_HomeFanout();
	t.OrderId = ;
	t.ForUserId = ;
	t.PostId = ;
	t.PostUserId = ;
	t.ResharedId = ;
	*/

	/*
	PB_HomeFanout t = new PB_HomeFanout();
	t.OrderId = m.getOrderId() ;
	t.ForUserId = m.getForUserId() ;
	t.PostId = m.getPostId() ;
	t.PostUserId = m.getPostUserId() ;
	t.ResharedId = m.getResharedId() ;
	*/

	public class PB_SuggestedTopPost {
	   public long Id;
	   public long PostId;
	}
	/*
	folding
	PBFlatTypes.PB_SuggestedTopPost t = new PBFlatTypes.PB_SuggestedTopPost();
    t.setId();
    t.setPostId();
	*/

	/*
	PBFlatTypes.PB_SuggestedTopPost t = new PBFlatTypes.PB_SuggestedTopPost();
	t.Id = ;
	t.PostId = ;
	*/

	/*
	PB_SuggestedTopPost t = new PB_SuggestedTopPost();
	t.Id = m.getId() ;
	t.PostId = m.getPostId() ;
	*/

	public class PB_SuggestedUser {
	   public int Id;
	   public int UserId;
	   public int TargetId;
	   public float Weight;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_SuggestedUser t = new PBFlatTypes.PB_SuggestedUser();
    t.setId();
    t.setUserId();
    t.setTargetId();
    t.setWeight();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_SuggestedUser t = new PBFlatTypes.PB_SuggestedUser();
	t.Id = ;
	t.UserId = ;
	t.TargetId = ;
	t.Weight = ;
	t.CreatedTime = ;
	*/

	/*
	PB_SuggestedUser t = new PB_SuggestedUser();
	t.Id = m.getId() ;
	t.UserId = m.getUserId() ;
	t.TargetId = m.getTargetId() ;
	t.Weight = m.getWeight() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_PushChat {
	   public long PushId;
	   public int ToUserId;
	   public int PushTypeId;
	   public String RoomKey;
	   public String ChatKey;
	   public int Seq;
	   public int UnseenCount;
	   public long FromHighMessageId;
	   public long ToLowMessageId;
	   public long MessageId;
	   public long MessageFileId;
	   public byte[] MessagePb;
	   public String MessageJson;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_PushChat t = new PBFlatTypes.PB_PushChat();
    t.setPushId();
    t.setToUserId();
    t.setPushTypeId();
    t.setRoomKey();
    t.setChatKey();
    t.setSeq();
    t.setUnseenCount();
    t.setFromHighMessageId();
    t.setToLowMessageId();
    t.setMessageId();
    t.setMessageFileId();
    t.setMessagePb();
    t.setMessageJson();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_PushChat t = new PBFlatTypes.PB_PushChat();
	t.PushId = ;
	t.ToUserId = ;
	t.PushTypeId = ;
	t.RoomKey = ;
	t.ChatKey = ;
	t.Seq = ;
	t.UnseenCount = ;
	t.FromHighMessageId = ;
	t.ToLowMessageId = ;
	t.MessageId = ;
	t.MessageFileId = ;
	t.MessagePb = ;
	t.MessageJson = ;
	t.CreatedTime = ;
	*/

	/*
	PB_PushChat t = new PB_PushChat();
	t.PushId = m.getPushId() ;
	t.ToUserId = m.getToUserId() ;
	t.PushTypeId = m.getPushTypeId() ;
	t.RoomKey = m.getRoomKey() ;
	t.ChatKey = m.getChatKey() ;
	t.Seq = m.getSeq() ;
	t.UnseenCount = m.getUnseenCount() ;
	t.FromHighMessageId = m.getFromHighMessageId() ;
	t.ToLowMessageId = m.getToLowMessageId() ;
	t.MessageId = m.getMessageId() ;
	t.MessageFileId = m.getMessageFileId() ;
	t.MessagePb = m.getMessagePb() ;
	t.MessageJson = m.getMessageJson() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_HTTPRPCLog {
	   public int Id;
	   public String Time;
	   public String MethodFull;
	   public String MethodParent;
	   public int UserId;
	   public String SessionId;
	   public int StatusCode;
	   public int InputSize;
	   public int OutputSize;
	   public String ReqestJson;
	   public String ResponseJson;
	   public String ReqestParamJson;
	   public String ResponseMsgJson;
	}
	/*
	folding
	PBFlatTypes.PB_HTTPRPCLog t = new PBFlatTypes.PB_HTTPRPCLog();
    t.setId();
    t.setTime();
    t.setMethodFull();
    t.setMethodParent();
    t.setUserId();
    t.setSessionId();
    t.setStatusCode();
    t.setInputSize();
    t.setOutputSize();
    t.setReqestJson();
    t.setResponseJson();
    t.setReqestParamJson();
    t.setResponseMsgJson();
	*/

	/*
	PBFlatTypes.PB_HTTPRPCLog t = new PBFlatTypes.PB_HTTPRPCLog();
	t.Id = ;
	t.Time = ;
	t.MethodFull = ;
	t.MethodParent = ;
	t.UserId = ;
	t.SessionId = ;
	t.StatusCode = ;
	t.InputSize = ;
	t.OutputSize = ;
	t.ReqestJson = ;
	t.ResponseJson = ;
	t.ReqestParamJson = ;
	t.ResponseMsgJson = ;
	*/

	/*
	PB_HTTPRPCLog t = new PB_HTTPRPCLog();
	t.Id = m.getId() ;
	t.Time = m.getTime() ;
	t.MethodFull = m.getMethodFull() ;
	t.MethodParent = m.getMethodParent() ;
	t.UserId = m.getUserId() ;
	t.SessionId = m.getSessionId() ;
	t.StatusCode = m.getStatusCode() ;
	t.InputSize = m.getInputSize() ;
	t.OutputSize = m.getOutputSize() ;
	t.ReqestJson = m.getReqestJson() ;
	t.ResponseJson = m.getResponseJson() ;
	t.ReqestParamJson = m.getReqestParamJson() ;
	t.ResponseMsgJson = m.getResponseMsgJson() ;
	*/

	public class PB_MetricLog {
	   public int Id;
	   public int InstanceId;
	   public String StartFrom;
	   public String EndTo;
	   public int StartTime;
	   public String Duration;
	   public String MetericsJson;
	}
	/*
	folding
	PBFlatTypes.PB_MetricLog t = new PBFlatTypes.PB_MetricLog();
    t.setId();
    t.setInstanceId();
    t.setStartFrom();
    t.setEndTo();
    t.setStartTime();
    t.setDuration();
    t.setMetericsJson();
	*/

	/*
	PBFlatTypes.PB_MetricLog t = new PBFlatTypes.PB_MetricLog();
	t.Id = ;
	t.InstanceId = ;
	t.StartFrom = ;
	t.EndTo = ;
	t.StartTime = ;
	t.Duration = ;
	t.MetericsJson = ;
	*/

	/*
	PB_MetricLog t = new PB_MetricLog();
	t.Id = m.getId() ;
	t.InstanceId = m.getInstanceId() ;
	t.StartFrom = m.getStartFrom() ;
	t.EndTo = m.getEndTo() ;
	t.StartTime = m.getStartTime() ;
	t.Duration = m.getDuration() ;
	t.MetericsJson = m.getMetericsJson() ;
	*/

	public class PB_XfileServiceInfoLog {
	   public long Id;
	   public int InstanceId;
	   public String Url;
	   public String CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_XfileServiceInfoLog t = new PBFlatTypes.PB_XfileServiceInfoLog();
    t.setId();
    t.setInstanceId();
    t.setUrl();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_XfileServiceInfoLog t = new PBFlatTypes.PB_XfileServiceInfoLog();
	t.Id = ;
	t.InstanceId = ;
	t.Url = ;
	t.CreatedTime = ;
	*/

	/*
	PB_XfileServiceInfoLog t = new PB_XfileServiceInfoLog();
	t.Id = m.getId() ;
	t.InstanceId = m.getInstanceId() ;
	t.Url = m.getUrl() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_XfileServiceMetricLog {
	   public long Id;
	   public int InstanceId;
	   public String MetricJson;
	}
	/*
	folding
	PBFlatTypes.PB_XfileServiceMetricLog t = new PBFlatTypes.PB_XfileServiceMetricLog();
    t.setId();
    t.setInstanceId();
    t.setMetricJson();
	*/

	/*
	PBFlatTypes.PB_XfileServiceMetricLog t = new PBFlatTypes.PB_XfileServiceMetricLog();
	t.Id = ;
	t.InstanceId = ;
	t.MetricJson = ;
	*/

	/*
	PB_XfileServiceMetricLog t = new PB_XfileServiceMetricLog();
	t.Id = m.getId() ;
	t.InstanceId = m.getInstanceId() ;
	t.MetricJson = m.getMetricJson() ;
	*/

	public class PB_XfileServiceRequestLog {
	   public long Id;
	   public int LocalSeq;
	   public int InstanceId;
	   public String Url;
	   public int HttpCode;
	   public String CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_XfileServiceRequestLog t = new PBFlatTypes.PB_XfileServiceRequestLog();
    t.setId();
    t.setLocalSeq();
    t.setInstanceId();
    t.setUrl();
    t.setHttpCode();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_XfileServiceRequestLog t = new PBFlatTypes.PB_XfileServiceRequestLog();
	t.Id = ;
	t.LocalSeq = ;
	t.InstanceId = ;
	t.Url = ;
	t.HttpCode = ;
	t.CreatedTime = ;
	*/

	/*
	PB_XfileServiceRequestLog t = new PB_XfileServiceRequestLog();
	t.Id = m.getId() ;
	t.LocalSeq = m.getLocalSeq() ;
	t.InstanceId = m.getInstanceId() ;
	t.Url = m.getUrl() ;
	t.HttpCode = m.getHttpCode() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_InvalidateCache {
	   public long OrderId;
	   public String CacheKey;
	}
	/*
	folding
	PBFlatTypes.PB_InvalidateCache t = new PBFlatTypes.PB_InvalidateCache();
    t.setOrderId();
    t.setCacheKey();
	*/

	/*
	PBFlatTypes.PB_InvalidateCache t = new PBFlatTypes.PB_InvalidateCache();
	t.OrderId = ;
	t.CacheKey = ;
	*/

	/*
	PB_InvalidateCache t = new PB_InvalidateCache();
	t.OrderId = m.getOrderId() ;
	t.CacheKey = m.getCacheKey() ;
	*/

	public class PB_MediaView {
	}
	/*
	folding
	PBFlatTypes.PB_MediaView t = new PBFlatTypes.PB_MediaView();
	*/

	/*
	PBFlatTypes.PB_MediaView t = new PBFlatTypes.PB_MediaView();
	*/

	/*
	PB_MediaView t = new PB_MediaView();
	*/

	public class PB_ActionView {
	   public long ActionId;
	   public int ActorUserId;
	   public int ActionTypeEnum;
	   public int PeerUserId;
	   public long PostId;
	   public long CommentId;
	   public long Murmur64Hash;
	   public int CreatedTime;
	   public PB_UserView ActorUserView;
	   public PB_PostView PostView;
	   public PB_CommentView CommentView;
	   public PB_UserView FollowedUserView;
	   public PB_UserView ContentOwenerUserView;
	}
	/*
	folding
	PBFlatTypes.PB_ActionView t = new PBFlatTypes.PB_ActionView();
    t.setActionId();
    t.setActorUserId();
    t.setActionTypeEnum();
    t.setPeerUserId();
    t.setPostId();
    t.setCommentId();
    t.setMurmur64Hash();
    t.setCreatedTime();
    t.setActorUserView();
    t.setPostView();
    t.setCommentView();
    t.setFollowedUserView();
    t.setContentOwenerUserView();
	*/

	/*
	PBFlatTypes.PB_ActionView t = new PBFlatTypes.PB_ActionView();
	t.ActionId = ;
	t.ActorUserId = ;
	t.ActionTypeEnum = ;
	t.PeerUserId = ;
	t.PostId = ;
	t.CommentId = ;
	t.Murmur64Hash = ;
	t.CreatedTime = ;
	t.ActorUserView = ;
	t.PostView = ;
	t.CommentView = ;
	t.FollowedUserView = ;
	t.ContentOwenerUserView = ;
	*/

	/*
	PB_ActionView t = new PB_ActionView();
	t.ActionId = m.getActionId() ;
	t.ActorUserId = m.getActorUserId() ;
	t.ActionTypeEnum = m.getActionTypeEnum() ;
	t.PeerUserId = m.getPeerUserId() ;
	t.PostId = m.getPostId() ;
	t.CommentId = m.getCommentId() ;
	t.Murmur64Hash = m.getMurmur64Hash() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.ActorUserView = m.getActorUserView() ;
	t.PostView = m.getPostView() ;
	t.CommentView = m.getCommentView() ;
	t.FollowedUserView = m.getFollowedUserView() ;
	t.ContentOwenerUserView = m.getContentOwenerUserView() ;
	*/

	public class PB_NotifyView {
	   public long NotifyId;
	   public int ForUserId;
	   public int ActorUserId;
	   public int NotiyTypeEnum;
	   public long PostId;
	   public long CommentId;
	   public int PeerUserId;
	   public long Murmur64Hash;
	   public int SeenStatus;
	   public int CreatedTime;
	   public PB_UserView ActorUserView;
	   public PB_PostView PostView;
	   public PB_CommentView CommentView;
	}
	/*
	folding
	PBFlatTypes.PB_NotifyView t = new PBFlatTypes.PB_NotifyView();
    t.setNotifyId();
    t.setForUserId();
    t.setActorUserId();
    t.setNotiyTypeEnum();
    t.setPostId();
    t.setCommentId();
    t.setPeerUserId();
    t.setMurmur64Hash();
    t.setSeenStatus();
    t.setCreatedTime();
    t.setActorUserView();
    t.setPostView();
    t.setCommentView();
	*/

	/*
	PBFlatTypes.PB_NotifyView t = new PBFlatTypes.PB_NotifyView();
	t.NotifyId = ;
	t.ForUserId = ;
	t.ActorUserId = ;
	t.NotiyTypeEnum = ;
	t.PostId = ;
	t.CommentId = ;
	t.PeerUserId = ;
	t.Murmur64Hash = ;
	t.SeenStatus = ;
	t.CreatedTime = ;
	t.ActorUserView = ;
	t.PostView = ;
	t.CommentView = ;
	*/

	/*
	PB_NotifyView t = new PB_NotifyView();
	t.NotifyId = m.getNotifyId() ;
	t.ForUserId = m.getForUserId() ;
	t.ActorUserId = m.getActorUserId() ;
	t.NotiyTypeEnum = m.getNotiyTypeEnum() ;
	t.PostId = m.getPostId() ;
	t.CommentId = m.getCommentId() ;
	t.PeerUserId = m.getPeerUserId() ;
	t.Murmur64Hash = m.getMurmur64Hash() ;
	t.SeenStatus = m.getSeenStatus() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.ActorUserView = m.getActorUserView() ;
	t.PostView = m.getPostView() ;
	t.CommentView = m.getCommentView() ;
	*/

	public class PB_CommentView {
	   public long CommentId;
	   public int UserId;
	   public long PostId;
	   public String Text;
	   public int LikesCount;
	   public int CreatedTime;
	   public PB_UserView SenderUserView;
	}
	/*
	folding
	PBFlatTypes.PB_CommentView t = new PBFlatTypes.PB_CommentView();
    t.setCommentId();
    t.setUserId();
    t.setPostId();
    t.setText();
    t.setLikesCount();
    t.setCreatedTime();
    t.setSenderUserView();
	*/

	/*
	PBFlatTypes.PB_CommentView t = new PBFlatTypes.PB_CommentView();
	t.CommentId = ;
	t.UserId = ;
	t.PostId = ;
	t.Text = ;
	t.LikesCount = ;
	t.CreatedTime = ;
	t.SenderUserView = ;
	*/

	/*
	PB_CommentView t = new PB_CommentView();
	t.CommentId = m.getCommentId() ;
	t.UserId = m.getUserId() ;
	t.PostId = m.getPostId() ;
	t.Text = m.getText() ;
	t.LikesCount = m.getLikesCount() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.SenderUserView = m.getSenderUserView() ;
	*/

	public class PB_PostView {
	   public long PostId;
	   public int UserId;
	   public PostTypeEnum PostTypeEnum;
	   public String Text;
	   public String RichText;
	   public int MediaCount;
	   public int SharedTo;
	   public int DisableComment;
	   public int HasTag;
	   public int CommentsCount;
	   public int LikesCount;
	   public int ViewsCount;
	   public int EditedTime;
	   public int CreatedTime;
	   public long ReSharedPostId;
	   public boolean DidILiked;
	   public boolean DidIReShared;
	   public PB_UserView SenderUserView;
	   public PB_UserView ReSharedUserView;
	   public PB_MediaView MediaView;
	   public PB_MediaView MediaViewList;
	}
	/*
	folding
	PBFlatTypes.PB_PostView t = new PBFlatTypes.PB_PostView();
    t.setPostId();
    t.setUserId();
    t.setPostTypeEnum();
    t.setText();
    t.setRichText();
    t.setMediaCount();
    t.setSharedTo();
    t.setDisableComment();
    t.setHasTag();
    t.setCommentsCount();
    t.setLikesCount();
    t.setViewsCount();
    t.setEditedTime();
    t.setCreatedTime();
    t.setReSharedPostId();
    t.setDidILiked();
    t.setDidIReShared();
    t.setSenderUserView();
    t.setReSharedUserView();
    t.setMediaView();
    t.setMediaViewList();
	*/

	/*
	PBFlatTypes.PB_PostView t = new PBFlatTypes.PB_PostView();
	t.PostId = ;
	t.UserId = ;
	t.PostTypeEnum = ;
	t.Text = ;
	t.RichText = ;
	t.MediaCount = ;
	t.SharedTo = ;
	t.DisableComment = ;
	t.HasTag = ;
	t.CommentsCount = ;
	t.LikesCount = ;
	t.ViewsCount = ;
	t.EditedTime = ;
	t.CreatedTime = ;
	t.ReSharedPostId = ;
	t.DidILiked = ;
	t.DidIReShared = ;
	t.SenderUserView = ;
	t.ReSharedUserView = ;
	t.MediaView = ;
	t.MediaViewList = ;
	*/

	/*
	PB_PostView t = new PB_PostView();
	t.PostId = m.getPostId() ;
	t.UserId = m.getUserId() ;
	t.PostTypeEnum = m.getPostTypeEnum() ;
	t.Text = m.getText() ;
	t.RichText = m.getRichText() ;
	t.MediaCount = m.getMediaCount() ;
	t.SharedTo = m.getSharedTo() ;
	t.DisableComment = m.getDisableComment() ;
	t.HasTag = m.getHasTag() ;
	t.CommentsCount = m.getCommentsCount() ;
	t.LikesCount = m.getLikesCount() ;
	t.ViewsCount = m.getViewsCount() ;
	t.EditedTime = m.getEditedTime() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.ReSharedPostId = m.getReSharedPostId() ;
	t.DidILiked = m.getDidILiked() ;
	t.DidIReShared = m.getDidIReShared() ;
	t.SenderUserView = m.getSenderUserView() ;
	t.ReSharedUserView = m.getReSharedUserView() ;
	t.MediaView = m.getMediaView() ;
	t.MediaViewList = m.getMediaViewList() ;
	*/

	public class PB_ChatView {
	   public long ChatId;
	   public String ChatKey;
	   public String RoomKey;
	   public int RoomType;
	   public int UserId;
	   public int PeerUserId;
	   public long GroupId;
	   public long HashTagId;
	   public int StartedByMe;
	   public String Title;
	   public long PinTime;
	   public long FromMsgId;
	   public int Seq;
	   public long LastMsgId;
	   public int LastMsgStatus;
	   public int SeenSeq;
	   public long SeenMsgId;
	   public int Left;
	   public int Creator;
	   public int Kicked;
	   public int Admin;
	   public int Deactivated;
	   public int VersionTime;
	   public int SortTime;
	   public int CreatedTime;
	   public String DraftText;
	   public long DratReplyToMsgId;
	   public int IsMute;
	   public PB_UserView UserView;
	   public PB_GroupView GroupView;
	   public PB_MessageView FirstUnreadMessage;
	   public PB_MessageView LastMessage;
	}
	/*
	folding
	PBFlatTypes.PB_ChatView t = new PBFlatTypes.PB_ChatView();
    t.setChatId();
    t.setChatKey();
    t.setRoomKey();
    t.setRoomType();
    t.setUserId();
    t.setPeerUserId();
    t.setGroupId();
    t.setHashTagId();
    t.setStartedByMe();
    t.setTitle();
    t.setPinTime();
    t.setFromMsgId();
    t.setSeq();
    t.setLastMsgId();
    t.setLastMsgStatus();
    t.setSeenSeq();
    t.setSeenMsgId();
    t.setLeft();
    t.setCreator();
    t.setKicked();
    t.setAdmin();
    t.setDeactivated();
    t.setVersionTime();
    t.setSortTime();
    t.setCreatedTime();
    t.setDraftText();
    t.setDratReplyToMsgId();
    t.setIsMute();
    t.setUserView();
    t.setGroupView();
    t.setFirstUnreadMessage();
    t.setLastMessage();
	*/

	/*
	PBFlatTypes.PB_ChatView t = new PBFlatTypes.PB_ChatView();
	t.ChatId = ;
	t.ChatKey = ;
	t.RoomKey = ;
	t.RoomType = ;
	t.UserId = ;
	t.PeerUserId = ;
	t.GroupId = ;
	t.HashTagId = ;
	t.StartedByMe = ;
	t.Title = ;
	t.PinTime = ;
	t.FromMsgId = ;
	t.Seq = ;
	t.LastMsgId = ;
	t.LastMsgStatus = ;
	t.SeenSeq = ;
	t.SeenMsgId = ;
	t.Left = ;
	t.Creator = ;
	t.Kicked = ;
	t.Admin = ;
	t.Deactivated = ;
	t.VersionTime = ;
	t.SortTime = ;
	t.CreatedTime = ;
	t.DraftText = ;
	t.DratReplyToMsgId = ;
	t.IsMute = ;
	t.UserView = ;
	t.GroupView = ;
	t.FirstUnreadMessage = ;
	t.LastMessage = ;
	*/

	/*
	PB_ChatView t = new PB_ChatView();
	t.ChatId = m.getChatId() ;
	t.ChatKey = m.getChatKey() ;
	t.RoomKey = m.getRoomKey() ;
	t.RoomType = m.getRoomType() ;
	t.UserId = m.getUserId() ;
	t.PeerUserId = m.getPeerUserId() ;
	t.GroupId = m.getGroupId() ;
	t.HashTagId = m.getHashTagId() ;
	t.StartedByMe = m.getStartedByMe() ;
	t.Title = m.getTitle() ;
	t.PinTime = m.getPinTime() ;
	t.FromMsgId = m.getFromMsgId() ;
	t.Seq = m.getSeq() ;
	t.LastMsgId = m.getLastMsgId() ;
	t.LastMsgStatus = m.getLastMsgStatus() ;
	t.SeenSeq = m.getSeenSeq() ;
	t.SeenMsgId = m.getSeenMsgId() ;
	t.Left = m.getLeft() ;
	t.Creator = m.getCreator() ;
	t.Kicked = m.getKicked() ;
	t.Admin = m.getAdmin() ;
	t.Deactivated = m.getDeactivated() ;
	t.VersionTime = m.getVersionTime() ;
	t.SortTime = m.getSortTime() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.DraftText = m.getDraftText() ;
	t.DratReplyToMsgId = m.getDratReplyToMsgId() ;
	t.IsMute = m.getIsMute() ;
	t.UserView = m.getUserView() ;
	t.GroupView = m.getGroupView() ;
	t.FirstUnreadMessage = m.getFirstUnreadMessage() ;
	t.LastMessage = m.getLastMessage() ;
	*/

	public class PB_GroupView {
	   public long GroupId;
	   public String GroupKey;
	   public String GroupName;
	   public String UserName;
	   public int IsSuperGroup;
	   public long HashTagId;
	   public int CreatorUserId;
	   public int GroupPrivacy;
	   public int HistoryViewAble;
	   public long Seq;
	   public long LastMsgId;
	   public long PinedMsgId;
	   public long AvatarRefId;
	   public int AvatarCount;
	   public String About;
	   public String InviteLink;
	   public int MembersCount;
	   public int SortTime;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_GroupView t = new PBFlatTypes.PB_GroupView();
    t.setGroupId();
    t.setGroupKey();
    t.setGroupName();
    t.setUserName();
    t.setIsSuperGroup();
    t.setHashTagId();
    t.setCreatorUserId();
    t.setGroupPrivacy();
    t.setHistoryViewAble();
    t.setSeq();
    t.setLastMsgId();
    t.setPinedMsgId();
    t.setAvatarRefId();
    t.setAvatarCount();
    t.setAbout();
    t.setInviteLink();
    t.setMembersCount();
    t.setSortTime();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_GroupView t = new PBFlatTypes.PB_GroupView();
	t.GroupId = ;
	t.GroupKey = ;
	t.GroupName = ;
	t.UserName = ;
	t.IsSuperGroup = ;
	t.HashTagId = ;
	t.CreatorUserId = ;
	t.GroupPrivacy = ;
	t.HistoryViewAble = ;
	t.Seq = ;
	t.LastMsgId = ;
	t.PinedMsgId = ;
	t.AvatarRefId = ;
	t.AvatarCount = ;
	t.About = ;
	t.InviteLink = ;
	t.MembersCount = ;
	t.SortTime = ;
	t.CreatedTime = ;
	*/

	/*
	PB_GroupView t = new PB_GroupView();
	t.GroupId = m.getGroupId() ;
	t.GroupKey = m.getGroupKey() ;
	t.GroupName = m.getGroupName() ;
	t.UserName = m.getUserName() ;
	t.IsSuperGroup = m.getIsSuperGroup() ;
	t.HashTagId = m.getHashTagId() ;
	t.CreatorUserId = m.getCreatorUserId() ;
	t.GroupPrivacy = m.getGroupPrivacy() ;
	t.HistoryViewAble = m.getHistoryViewAble() ;
	t.Seq = m.getSeq() ;
	t.LastMsgId = m.getLastMsgId() ;
	t.PinedMsgId = m.getPinedMsgId() ;
	t.AvatarRefId = m.getAvatarRefId() ;
	t.AvatarCount = m.getAvatarCount() ;
	t.About = m.getAbout() ;
	t.InviteLink = m.getInviteLink() ;
	t.MembersCount = m.getMembersCount() ;
	t.SortTime = m.getSortTime() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_GroupMemberView {
	   public long OrderId;
	   public long GroupId;
	   public int UserId;
	   public int ByUserId;
	   public int GroupRole;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_GroupMemberView t = new PBFlatTypes.PB_GroupMemberView();
    t.setOrderId();
    t.setGroupId();
    t.setUserId();
    t.setByUserId();
    t.setGroupRole();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_GroupMemberView t = new PBFlatTypes.PB_GroupMemberView();
	t.OrderId = ;
	t.GroupId = ;
	t.UserId = ;
	t.ByUserId = ;
	t.GroupRole = ;
	t.CreatedTime = ;
	*/

	/*
	PB_GroupMemberView t = new PB_GroupMemberView();
	t.OrderId = m.getOrderId() ;
	t.GroupId = m.getGroupId() ;
	t.UserId = m.getUserId() ;
	t.ByUserId = m.getByUserId() ;
	t.GroupRole = m.getGroupRole() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_MessageView {
	   public String RoomKey;
	   public long MessageId;
	   public int UserId;
	   public long FileRefId;
	   public int MessageType;
	   public String Text;
	   public int Hiden;
	   public int Seq;
	   public long ForwardedMsgId;
	   public long PostId;
	   public long StickerId;
	   public int CreatedTime;
	   public int DeliveredTime;
	   public int SeenTime;
	   public int DeliviryStatus;
	   public long ReplyToMessageId;
	   public long ViewsCount;
	   public int EditTime;
	   public int Ttl;
	   public PB_FileRedView FileRedView;
	}
	/*
	folding
	PBFlatTypes.PB_MessageView t = new PBFlatTypes.PB_MessageView();
    t.setRoomKey();
    t.setMessageId();
    t.setUserId();
    t.setFileRefId();
    t.setMessageType();
    t.setText();
    t.setHiden();
    t.setSeq();
    t.setForwardedMsgId();
    t.setPostId();
    t.setStickerId();
    t.setCreatedTime();
    t.setDeliveredTime();
    t.setSeenTime();
    t.setDeliviryStatus();
    t.setReplyToMessageId();
    t.setViewsCount();
    t.setEditTime();
    t.setTtl();
    t.setFileRedView();
	*/

	/*
	PBFlatTypes.PB_MessageView t = new PBFlatTypes.PB_MessageView();
	t.RoomKey = ;
	t.MessageId = ;
	t.UserId = ;
	t.FileRefId = ;
	t.MessageType = ;
	t.Text = ;
	t.Hiden = ;
	t.Seq = ;
	t.ForwardedMsgId = ;
	t.PostId = ;
	t.StickerId = ;
	t.CreatedTime = ;
	t.DeliveredTime = ;
	t.SeenTime = ;
	t.DeliviryStatus = ;
	t.ReplyToMessageId = ;
	t.ViewsCount = ;
	t.EditTime = ;
	t.Ttl = ;
	t.FileRedView = ;
	*/

	/*
	PB_MessageView t = new PB_MessageView();
	t.RoomKey = m.getRoomKey() ;
	t.MessageId = m.getMessageId() ;
	t.UserId = m.getUserId() ;
	t.FileRefId = m.getFileRefId() ;
	t.MessageType = m.getMessageType() ;
	t.Text = m.getText() ;
	t.Hiden = m.getHiden() ;
	t.Seq = m.getSeq() ;
	t.ForwardedMsgId = m.getForwardedMsgId() ;
	t.PostId = m.getPostId() ;
	t.StickerId = m.getStickerId() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.DeliveredTime = m.getDeliveredTime() ;
	t.SeenTime = m.getSeenTime() ;
	t.DeliviryStatus = m.getDeliviryStatus() ;
	t.ReplyToMessageId = m.getReplyToMessageId() ;
	t.ViewsCount = m.getViewsCount() ;
	t.EditTime = m.getEditTime() ;
	t.Ttl = m.getTtl() ;
	t.FileRedView = m.getFileRedView() ;
	*/

	public class PB_FileRedView {
	   public long FileRefId;
	   public long UserId;
	   public String Name;
	   public int Width;
	   public int Height;
	   public int Duration;
	   public String Extension;
	   public String UrlSource;
	}
	/*
	folding
	PBFlatTypes.PB_FileRedView t = new PBFlatTypes.PB_FileRedView();
    t.setFileRefId();
    t.setUserId();
    t.setName();
    t.setWidth();
    t.setHeight();
    t.setDuration();
    t.setExtension();
    t.setUrlSource();
	*/

	/*
	PBFlatTypes.PB_FileRedView t = new PBFlatTypes.PB_FileRedView();
	t.FileRefId = ;
	t.UserId = ;
	t.Name = ;
	t.Width = ;
	t.Height = ;
	t.Duration = ;
	t.Extension = ;
	t.UrlSource = ;
	*/

	/*
	PB_FileRedView t = new PB_FileRedView();
	t.FileRefId = m.getFileRefId() ;
	t.UserId = m.getUserId() ;
	t.Name = m.getName() ;
	t.Width = m.getWidth() ;
	t.Height = m.getHeight() ;
	t.Duration = m.getDuration() ;
	t.Extension = m.getExtension() ;
	t.UrlSource = m.getUrlSource() ;
	*/

	public class PB_UserView {
	   public int UserId;
	   public String UserName;
	   public String FirstName;
	   public String LastName;
	   public long AvatarRefId;
	   public int ProfilePrivacy;
	   public long Phone;
	   public String About;
	   public int FollowersCount;
	   public int FollowingCount;
	   public int PostsCount;
	   public int MediaCount;
	   public UserOnlineStatusEnum UserOnlineStatusEnum;
	   public int LastActiveTime;
	   public String LastActiveTimeShow;
	   public FollowingEnum MyFollwing;
	}
	/*
	folding
	PBFlatTypes.PB_UserView t = new PBFlatTypes.PB_UserView();
    t.setUserId();
    t.setUserName();
    t.setFirstName();
    t.setLastName();
    t.setAvatarRefId();
    t.setProfilePrivacy();
    t.setPhone();
    t.setAbout();
    t.setFollowersCount();
    t.setFollowingCount();
    t.setPostsCount();
    t.setMediaCount();
    t.setUserOnlineStatusEnum();
    t.setLastActiveTime();
    t.setLastActiveTimeShow();
    t.setMyFollwing();
	*/

	/*
	PBFlatTypes.PB_UserView t = new PBFlatTypes.PB_UserView();
	t.UserId = ;
	t.UserName = ;
	t.FirstName = ;
	t.LastName = ;
	t.AvatarRefId = ;
	t.ProfilePrivacy = ;
	t.Phone = ;
	t.About = ;
	t.FollowersCount = ;
	t.FollowingCount = ;
	t.PostsCount = ;
	t.MediaCount = ;
	t.UserOnlineStatusEnum = ;
	t.LastActiveTime = ;
	t.LastActiveTimeShow = ;
	t.MyFollwing = ;
	*/

	/*
	PB_UserView t = new PB_UserView();
	t.UserId = m.getUserId() ;
	t.UserName = m.getUserName() ;
	t.FirstName = m.getFirstName() ;
	t.LastName = m.getLastName() ;
	t.AvatarRefId = m.getAvatarRefId() ;
	t.ProfilePrivacy = m.getProfilePrivacy() ;
	t.Phone = m.getPhone() ;
	t.About = m.getAbout() ;
	t.FollowersCount = m.getFollowersCount() ;
	t.FollowingCount = m.getFollowingCount() ;
	t.PostsCount = m.getPostsCount() ;
	t.MediaCount = m.getMediaCount() ;
	t.UserOnlineStatusEnum = m.getUserOnlineStatusEnum() ;
	t.LastActiveTime = m.getLastActiveTime() ;
	t.LastActiveTimeShow = m.getLastActiveTimeShow() ;
	t.MyFollwing = m.getMyFollwing() ;
	*/

	public class PB_SettingNotificationView {
	}
	/*
	folding
	PBFlatTypes.PB_SettingNotificationView t = new PBFlatTypes.PB_SettingNotificationView();
	*/

	/*
	PBFlatTypes.PB_SettingNotificationView t = new PBFlatTypes.PB_SettingNotificationView();
	*/

	/*
	PB_SettingNotificationView t = new PB_SettingNotificationView();
	*/

	public class PB_AppConfig {
	   public boolean DeprecatedClient;
	   public boolean HasNewUpdate;
	}
	/*
	folding
	PBFlatTypes.PB_AppConfig t = new PBFlatTypes.PB_AppConfig();
    t.setDeprecatedClient();
    t.setHasNewUpdate();
	*/

	/*
	PBFlatTypes.PB_AppConfig t = new PBFlatTypes.PB_AppConfig();
	t.DeprecatedClient = ;
	t.HasNewUpdate = ;
	*/

	/*
	PB_AppConfig t = new PB_AppConfig();
	t.DeprecatedClient = m.getDeprecatedClient() ;
	t.HasNewUpdate = m.getHasNewUpdate() ;
	*/

	public class PB_UserProfileView {
	}
	/*
	folding
	PBFlatTypes.PB_UserProfileView t = new PBFlatTypes.PB_UserProfileView();
	*/

	/*
	PBFlatTypes.PB_UserProfileView t = new PBFlatTypes.PB_UserProfileView();
	*/

	/*
	PB_UserProfileView t = new PB_UserProfileView();
	*/

	public class PB_UserViewRowify {
	   public long Id;
	   public int CreatedTime;
	   public PB_UserView UserView;
	}
	/*
	folding
	PBFlatTypes.PB_UserViewRowify t = new PBFlatTypes.PB_UserViewRowify();
    t.setId();
    t.setCreatedTime();
    t.setUserView();
	*/

	/*
	PBFlatTypes.PB_UserViewRowify t = new PBFlatTypes.PB_UserViewRowify();
	t.Id = ;
	t.CreatedTime = ;
	t.UserView = ;
	*/

	/*
	PB_UserViewRowify t = new PB_UserViewRowify();
	t.Id = m.getId() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.UserView = m.getUserView() ;
	*/

	public class PB_PostViewRowify {
	   public long Id;
	   public PB_PostView PostView;
	}
	/*
	folding
	PBFlatTypes.PB_PostViewRowify t = new PBFlatTypes.PB_PostViewRowify();
    t.setId();
    t.setPostView();
	*/

	/*
	PBFlatTypes.PB_PostViewRowify t = new PBFlatTypes.PB_PostViewRowify();
	t.Id = ;
	t.PostView = ;
	*/

	/*
	PB_PostViewRowify t = new PB_PostViewRowify();
	t.Id = m.getId() ;
	t.PostView = m.getPostView() ;
	*/

	public class PB_TagView {
	   public long TagId;
	   public String Name;
	   public int Count;
	   public int TagStatusEnum;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_TagView t = new PBFlatTypes.PB_TagView();
    t.setTagId();
    t.setName();
    t.setCount();
    t.setTagStatusEnum();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_TagView t = new PBFlatTypes.PB_TagView();
	t.TagId = ;
	t.Name = ;
	t.Count = ;
	t.TagStatusEnum = ;
	t.CreatedTime = ;
	*/

	/*
	PB_TagView t = new PB_TagView();
	t.TagId = m.getTagId() ;
	t.Name = m.getName() ;
	t.Count = m.getCount() ;
	t.TagStatusEnum = m.getTagStatusEnum() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_TopTagWithSamplePosts {
	   public PB_TagView TagView;
	   public PB_PostView PostViewList;
	}
	/*
	folding
	PBFlatTypes.PB_TopTagWithSamplePosts t = new PBFlatTypes.PB_TopTagWithSamplePosts();
    t.setTagView();
    t.setPostViewList();
	*/

	/*
	PBFlatTypes.PB_TopTagWithSamplePosts t = new PBFlatTypes.PB_TopTagWithSamplePosts();
	t.TagView = ;
	t.PostViewList = ;
	*/

	/*
	PB_TopTagWithSamplePosts t = new PB_TopTagWithSamplePosts();
	t.TagView = m.getTagView() ;
	t.PostViewList = m.getPostViewList() ;
	*/

	public class PB_SelfUserView {
	   public PB_UserView UserView;
	   public int ProfilePrivacy;
	   public int OnlinePrivacy;
	   public int CallPrivacy;
	   public int AddToGroupPrivacy;
	   public int SeenMessagePrivacy;
	   public PB_SettingNotificationView SettingNotification;
	}
	/*
	folding
	PBFlatTypes.PB_SelfUserView t = new PBFlatTypes.PB_SelfUserView();
    t.setUserView();
    t.setProfilePrivacy();
    t.setOnlinePrivacy();
    t.setCallPrivacy();
    t.setAddToGroupPrivacy();
    t.setSeenMessagePrivacy();
    t.setSettingNotification();
	*/

	/*
	PBFlatTypes.PB_SelfUserView t = new PBFlatTypes.PB_SelfUserView();
	t.UserView = ;
	t.ProfilePrivacy = ;
	t.OnlinePrivacy = ;
	t.CallPrivacy = ;
	t.AddToGroupPrivacy = ;
	t.SeenMessagePrivacy = ;
	t.SettingNotification = ;
	*/

	/*
	PB_SelfUserView t = new PB_SelfUserView();
	t.UserView = m.getUserView() ;
	t.ProfilePrivacy = m.getProfilePrivacy() ;
	t.OnlinePrivacy = m.getOnlinePrivacy() ;
	t.CallPrivacy = m.getCallPrivacy() ;
	t.AddToGroupPrivacy = m.getAddToGroupPrivacy() ;
	t.SeenMessagePrivacy = m.getSeenMessagePrivacy() ;
	t.SettingNotification = m.getSettingNotification() ;
	*/

	public class PB_Error {
	   public ServerErrors Error;
	   public boolean ShowError;
	   public String ErrorMessage;
	}
	/*
	folding
	PBFlatTypes.PB_Error t = new PBFlatTypes.PB_Error();
    t.setError();
    t.setShowError();
    t.setErrorMessage();
	*/

	/*
	PBFlatTypes.PB_Error t = new PBFlatTypes.PB_Error();
	t.Error = ;
	t.ShowError = ;
	t.ErrorMessage = ;
	*/

	/*
	PB_Error t = new PB_Error();
	t.Error = m.getError() ;
	t.ShowError = m.getShowError() ;
	t.ErrorMessage = m.getErrorMessage() ;
	*/

	
}

/*

RPC_HANDLERS.RPC_Auth RPC_Auth_Handeler = null;
RPC_HANDLERS.RPC_Chat RPC_Chat_Handeler = null;
RPC_HANDLERS.RPC_General RPC_General_Handeler = null;
RPC_HANDLERS.RPC_Page RPC_Page_Handeler = null;
RPC_HANDLERS.RPC_Social RPC_Social_Handeler = null;
RPC_HANDLERS.RPC_User RPC_User_Handeler = null;
	
*/