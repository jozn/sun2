package ir.ms.pb;

public class PBFlatTypes {

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

	public class PB_ResponseExtra {
	   public long ErrorCode;
	   public String ErrMessage;
	   public String RpcFullName;
	   public byte[] Data;
	}
	/*
	folding
	PBFlatTypes.PB_ResponseExtra t = new PBFlatTypes.PB_ResponseExtra();
    t.setErrorCode();
    t.setErrMessage();
    t.setRpcFullName();
    t.setData();
	*/

	/*
	PBFlatTypes.PB_ResponseExtra t = new PBFlatTypes.PB_ResponseExtra();
	t.ErrorCode = ;
	t.ErrMessage = ;
	t.RpcFullName = ;
	t.Data = ;
	*/

	/*
	PB_ResponseExtra t = new PB_ResponseExtra();
	t.ErrorCode = m.getErrorCode() ;
	t.ErrMessage = m.getErrMessage() ;
	t.RpcFullName = m.getRpcFullName() ;
	t.Data = m.getData() ;
	*/

	public class PB_Pager {
	   public long Page;
	   public long Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_Pager t = new PBFlatTypes.PB_Pager();
    t.setPage();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_Pager t = new PBFlatTypes.PB_Pager();
	t.Page = ;
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_Pager t = new PB_Pager();
	t.Page = m.getPage() ;
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_UserParam_CheckUserName2 {
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_CheckUserName2 t = new PBFlatTypes.PB_UserParam_CheckUserName2();
	*/

	/*
	PBFlatTypes.PB_UserParam_CheckUserName2 t = new PBFlatTypes.PB_UserParam_CheckUserName2();
	*/

	/*
	PB_UserParam_CheckUserName2 t = new PB_UserParam_CheckUserName2();
	*/

	public class PB_UserResponse_CheckUserName2 {
	}
	/*
	folding
	PBFlatTypes.PB_UserResponse_CheckUserName2 t = new PBFlatTypes.PB_UserResponse_CheckUserName2();
	*/

	/*
	PBFlatTypes.PB_UserResponse_CheckUserName2 t = new PBFlatTypes.PB_UserResponse_CheckUserName2();
	*/

	/*
	PB_UserResponse_CheckUserName2 t = new PB_UserResponse_CheckUserName2();
	*/

	public class PB_ChatParam_AddNewMessage {
	   public PB_MessageView MessageView;
	   public byte[] FileBlob;
	}
	/*
	folding
	PBFlatTypes.PB_ChatParam_AddNewMessage t = new PBFlatTypes.PB_ChatParam_AddNewMessage();
    t.setMessageView();
    t.setFileBlob();
	*/

	/*
	PBFlatTypes.PB_ChatParam_AddNewMessage t = new PBFlatTypes.PB_ChatParam_AddNewMessage();
	t.MessageView = ;
	t.FileBlob = ;
	*/

	/*
	PB_ChatParam_AddNewMessage t = new PB_ChatParam_AddNewMessage();
	t.MessageView = m.getMessageView() ;
	t.FileBlob = m.getFileBlob() ;
	*/

	public class PB_ChatResponse_AddNewMessage {
	   public PB_MessageView MessageView;
	}
	/*
	folding
	PBFlatTypes.PB_ChatResponse_AddNewMessage t = new PBFlatTypes.PB_ChatResponse_AddNewMessage();
    t.setMessageView();
	*/

	/*
	PBFlatTypes.PB_ChatResponse_AddNewMessage t = new PBFlatTypes.PB_ChatResponse_AddNewMessage();
	t.MessageView = ;
	*/

	/*
	PB_ChatResponse_AddNewMessage t = new PB_ChatResponse_AddNewMessage();
	t.MessageView = m.getMessageView() ;
	*/

	public class PB_ChatParam_SetRoomActionDoing {
	   public long GroupId;
	   public String DirectRoomKey;
	   public RoomActionDoingEnum ActionType;
	}
	/*
	folding
	PBFlatTypes.PB_ChatParam_SetRoomActionDoing t = new PBFlatTypes.PB_ChatParam_SetRoomActionDoing();
    t.setGroupId();
    t.setDirectRoomKey();
    t.setActionType();
	*/

	/*
	PBFlatTypes.PB_ChatParam_SetRoomActionDoing t = new PBFlatTypes.PB_ChatParam_SetRoomActionDoing();
	t.GroupId = ;
	t.DirectRoomKey = ;
	t.ActionType = ;
	*/

	/*
	PB_ChatParam_SetRoomActionDoing t = new PB_ChatParam_SetRoomActionDoing();
	t.GroupId = m.getGroupId() ;
	t.DirectRoomKey = m.getDirectRoomKey() ;
	t.ActionType = m.getActionType() ;
	*/

	public class PB_ChatResponse_SetRoomActionDoing {
	}
	/*
	folding
	PBFlatTypes.PB_ChatResponse_SetRoomActionDoing t = new PBFlatTypes.PB_ChatResponse_SetRoomActionDoing();
	*/

	/*
	PBFlatTypes.PB_ChatResponse_SetRoomActionDoing t = new PBFlatTypes.PB_ChatResponse_SetRoomActionDoing();
	*/

	/*
	PB_ChatResponse_SetRoomActionDoing t = new PB_ChatResponse_SetRoomActionDoing();
	*/

	public class PB_ChatParam_SetMessagesAsReceived {
	   public String RoomKey;
	   public long MessageIds;
	}
	/*
	folding
	PBFlatTypes.PB_ChatParam_SetMessagesAsReceived t = new PBFlatTypes.PB_ChatParam_SetMessagesAsReceived();
    t.setRoomKey();
    t.setMessageIds();
	*/

	/*
	PBFlatTypes.PB_ChatParam_SetMessagesAsReceived t = new PBFlatTypes.PB_ChatParam_SetMessagesAsReceived();
	t.RoomKey = ;
	t.MessageIds = ;
	*/

	/*
	PB_ChatParam_SetMessagesAsReceived t = new PB_ChatParam_SetMessagesAsReceived();
	t.RoomKey = m.getRoomKey() ;
	t.MessageIds = m.getMessageIds() ;
	*/

	public class PB_ChatResponse_SetMessagesAsReceived {
	}
	/*
	folding
	PBFlatTypes.PB_ChatResponse_SetMessagesAsReceived t = new PBFlatTypes.PB_ChatResponse_SetMessagesAsReceived();
	*/

	/*
	PBFlatTypes.PB_ChatResponse_SetMessagesAsReceived t = new PBFlatTypes.PB_ChatResponse_SetMessagesAsReceived();
	*/

	/*
	PB_ChatResponse_SetMessagesAsReceived t = new PB_ChatResponse_SetMessagesAsReceived();
	*/

	public class PB_ChatParam_SetChatMessagesRangeAsSeen {
	   public String RoomKey;
	   public long FromOlderMessageId;
	   public long TopNewerMessageId;
	   public long SeenTimeMs;
	}
	/*
	folding
	PBFlatTypes.PB_ChatParam_SetChatMessagesRangeAsSeen t = new PBFlatTypes.PB_ChatParam_SetChatMessagesRangeAsSeen();
    t.setRoomKey();
    t.setFromOlderMessageId();
    t.setTopNewerMessageId();
    t.setSeenTimeMs();
	*/

	/*
	PBFlatTypes.PB_ChatParam_SetChatMessagesRangeAsSeen t = new PBFlatTypes.PB_ChatParam_SetChatMessagesRangeAsSeen();
	t.RoomKey = ;
	t.FromOlderMessageId = ;
	t.TopNewerMessageId = ;
	t.SeenTimeMs = ;
	*/

	/*
	PB_ChatParam_SetChatMessagesRangeAsSeen t = new PB_ChatParam_SetChatMessagesRangeAsSeen();
	t.RoomKey = m.getRoomKey() ;
	t.FromOlderMessageId = m.getFromOlderMessageId() ;
	t.TopNewerMessageId = m.getTopNewerMessageId() ;
	t.SeenTimeMs = m.getSeenTimeMs() ;
	*/

	public class PB_ChatResponse_SetChatMessagesRangeAsSeen {
	}
	/*
	folding
	PBFlatTypes.PB_ChatResponse_SetChatMessagesRangeAsSeen t = new PBFlatTypes.PB_ChatResponse_SetChatMessagesRangeAsSeen();
	*/

	/*
	PBFlatTypes.PB_ChatResponse_SetChatMessagesRangeAsSeen t = new PBFlatTypes.PB_ChatResponse_SetChatMessagesRangeAsSeen();
	*/

	/*
	PB_ChatResponse_SetChatMessagesRangeAsSeen t = new PB_ChatResponse_SetChatMessagesRangeAsSeen();
	*/

	public class PB_ChatParam_DeleteChatHistory {
	   public String ChatKey;
	   public long FromMessageId;
	}
	/*
	folding
	PBFlatTypes.PB_ChatParam_DeleteChatHistory t = new PBFlatTypes.PB_ChatParam_DeleteChatHistory();
    t.setChatKey();
    t.setFromMessageId();
	*/

	/*
	PBFlatTypes.PB_ChatParam_DeleteChatHistory t = new PBFlatTypes.PB_ChatParam_DeleteChatHistory();
	t.ChatKey = ;
	t.FromMessageId = ;
	*/

	/*
	PB_ChatParam_DeleteChatHistory t = new PB_ChatParam_DeleteChatHistory();
	t.ChatKey = m.getChatKey() ;
	t.FromMessageId = m.getFromMessageId() ;
	*/

	public class PB_ChatResponse_DeleteChatHistory {
	}
	/*
	folding
	PBFlatTypes.PB_ChatResponse_DeleteChatHistory t = new PBFlatTypes.PB_ChatResponse_DeleteChatHistory();
	*/

	/*
	PBFlatTypes.PB_ChatResponse_DeleteChatHistory t = new PBFlatTypes.PB_ChatResponse_DeleteChatHistory();
	*/

	/*
	PB_ChatResponse_DeleteChatHistory t = new PB_ChatResponse_DeleteChatHistory();
	*/

	public class PB_ChatParam_DeleteMessagesByIds {
	   public String ChatKey;
	   public boolean Both;
	   public long MessagesIds;
	}
	/*
	folding
	PBFlatTypes.PB_ChatParam_DeleteMessagesByIds t = new PBFlatTypes.PB_ChatParam_DeleteMessagesByIds();
    t.setChatKey();
    t.setBoth();
    t.setMessagesIds();
	*/

	/*
	PBFlatTypes.PB_ChatParam_DeleteMessagesByIds t = new PBFlatTypes.PB_ChatParam_DeleteMessagesByIds();
	t.ChatKey = ;
	t.Both = ;
	t.MessagesIds = ;
	*/

	/*
	PB_ChatParam_DeleteMessagesByIds t = new PB_ChatParam_DeleteMessagesByIds();
	t.ChatKey = m.getChatKey() ;
	t.Both = m.getBoth() ;
	t.MessagesIds = m.getMessagesIds() ;
	*/

	public class PB_ChatResponse_DeleteMessagesByIds {
	}
	/*
	folding
	PBFlatTypes.PB_ChatResponse_DeleteMessagesByIds t = new PBFlatTypes.PB_ChatResponse_DeleteMessagesByIds();
	*/

	/*
	PBFlatTypes.PB_ChatResponse_DeleteMessagesByIds t = new PBFlatTypes.PB_ChatResponse_DeleteMessagesByIds();
	*/

	/*
	PB_ChatResponse_DeleteMessagesByIds t = new PB_ChatResponse_DeleteMessagesByIds();
	*/

	public class PB_ChatParam_EditMessage {
	   public String RoomKey;
	   public long MessageId;
	   public String NewText;
	}
	/*
	folding
	PBFlatTypes.PB_ChatParam_EditMessage t = new PBFlatTypes.PB_ChatParam_EditMessage();
    t.setRoomKey();
    t.setMessageId();
    t.setNewText();
	*/

	/*
	PBFlatTypes.PB_ChatParam_EditMessage t = new PBFlatTypes.PB_ChatParam_EditMessage();
	t.RoomKey = ;
	t.MessageId = ;
	t.NewText = ;
	*/

	/*
	PB_ChatParam_EditMessage t = new PB_ChatParam_EditMessage();
	t.RoomKey = m.getRoomKey() ;
	t.MessageId = m.getMessageId() ;
	t.NewText = m.getNewText() ;
	*/

	public class PB_ChatResponse_EditMessage {
	}
	/*
	folding
	PBFlatTypes.PB_ChatResponse_EditMessage t = new PBFlatTypes.PB_ChatResponse_EditMessage();
	*/

	/*
	PBFlatTypes.PB_ChatResponse_EditMessage t = new PBFlatTypes.PB_ChatResponse_EditMessage();
	*/

	/*
	PB_ChatResponse_EditMessage t = new PB_ChatResponse_EditMessage();
	*/

	public class PB_ChatParam_GetChatList {
	}
	/*
	folding
	PBFlatTypes.PB_ChatParam_GetChatList t = new PBFlatTypes.PB_ChatParam_GetChatList();
	*/

	/*
	PBFlatTypes.PB_ChatParam_GetChatList t = new PBFlatTypes.PB_ChatParam_GetChatList();
	*/

	/*
	PB_ChatParam_GetChatList t = new PB_ChatParam_GetChatList();
	*/

	public class PB_ChatResponse_GetChatList {
	   public PB_ChatView Chats;
	}
	/*
	folding
	PBFlatTypes.PB_ChatResponse_GetChatList t = new PBFlatTypes.PB_ChatResponse_GetChatList();
    t.setChats();
	*/

	/*
	PBFlatTypes.PB_ChatResponse_GetChatList t = new PBFlatTypes.PB_ChatResponse_GetChatList();
	t.Chats = ;
	*/

	/*
	PB_ChatResponse_GetChatList t = new PB_ChatResponse_GetChatList();
	t.Chats = m.getChats() ;
	*/

	public class PB_ChatParam_GetChatHistoryToOlder {
	   public String ChatKey;
	   public int Limit;
	   public long FromTopMessageId;
	}
	/*
	folding
	PBFlatTypes.PB_ChatParam_GetChatHistoryToOlder t = new PBFlatTypes.PB_ChatParam_GetChatHistoryToOlder();
    t.setChatKey();
    t.setLimit();
    t.setFromTopMessageId();
	*/

	/*
	PBFlatTypes.PB_ChatParam_GetChatHistoryToOlder t = new PBFlatTypes.PB_ChatParam_GetChatHistoryToOlder();
	t.ChatKey = ;
	t.Limit = ;
	t.FromTopMessageId = ;
	*/

	/*
	PB_ChatParam_GetChatHistoryToOlder t = new PB_ChatParam_GetChatHistoryToOlder();
	t.ChatKey = m.getChatKey() ;
	t.Limit = m.getLimit() ;
	t.FromTopMessageId = m.getFromTopMessageId() ;
	*/

	public class PB_ChatResponse_GetChatHistoryToOlder {
	   public PB_MessageView MessagesViews;
	   public boolean HasMore;
	}
	/*
	folding
	PBFlatTypes.PB_ChatResponse_GetChatHistoryToOlder t = new PBFlatTypes.PB_ChatResponse_GetChatHistoryToOlder();
    t.setMessagesViews();
    t.setHasMore();
	*/

	/*
	PBFlatTypes.PB_ChatResponse_GetChatHistoryToOlder t = new PBFlatTypes.PB_ChatResponse_GetChatHistoryToOlder();
	t.MessagesViews = ;
	t.HasMore = ;
	*/

	/*
	PB_ChatResponse_GetChatHistoryToOlder t = new PB_ChatResponse_GetChatHistoryToOlder();
	t.MessagesViews = m.getMessagesViews() ;
	t.HasMore = m.getHasMore() ;
	*/

	public class PB_OtherParam_Echo {
	   public String Text;
	}
	/*
	folding
	PBFlatTypes.PB_OtherParam_Echo t = new PBFlatTypes.PB_OtherParam_Echo();
    t.setText();
	*/

	/*
	PBFlatTypes.PB_OtherParam_Echo t = new PBFlatTypes.PB_OtherParam_Echo();
	t.Text = ;
	*/

	/*
	PB_OtherParam_Echo t = new PB_OtherParam_Echo();
	t.Text = m.getText() ;
	*/

	public class PB_OtherResponse_Echo {
	   public String Text;
	}
	/*
	folding
	PBFlatTypes.PB_OtherResponse_Echo t = new PBFlatTypes.PB_OtherResponse_Echo();
    t.setText();
	*/

	/*
	PBFlatTypes.PB_OtherResponse_Echo t = new PBFlatTypes.PB_OtherResponse_Echo();
	t.Text = ;
	*/

	/*
	PB_OtherResponse_Echo t = new PB_OtherResponse_Echo();
	t.Text = m.getText() ;
	*/

	public class PB_PageParam_GetCommentsPage {
	   public long PostId;
	   public int Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetCommentsPage t = new PBFlatTypes.PB_PageParam_GetCommentsPage();
    t.setPostId();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetCommentsPage t = new PBFlatTypes.PB_PageParam_GetCommentsPage();
	t.PostId = ;
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_PageParam_GetCommentsPage t = new PB_PageParam_GetCommentsPage();
	t.PostId = m.getPostId() ;
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_PageResponse_GetCommentsPage {
	   public PB_ResponseExtra Extra;
	   public PB_CommentView CommentViewList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetCommentsPage t = new PBFlatTypes.PB_PageResponse_GetCommentsPage();
    t.setExtra();
    t.setCommentViewList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetCommentsPage t = new PBFlatTypes.PB_PageResponse_GetCommentsPage();
	t.Extra = ;
	t.CommentViewList = ;
	*/

	/*
	PB_PageResponse_GetCommentsPage t = new PB_PageResponse_GetCommentsPage();
	t.Extra = m.getExtra() ;
	t.CommentViewList = m.getCommentViewList() ;
	*/

	public class PB_PageParam_GetHomePage {
	   public int Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetHomePage t = new PBFlatTypes.PB_PageParam_GetHomePage();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetHomePage t = new PBFlatTypes.PB_PageParam_GetHomePage();
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_PageParam_GetHomePage t = new PB_PageParam_GetHomePage();
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_PageResponse_GetHomePage {
	   public PB_ResponseExtra Extra;
	   public PB_PostView PostViewList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetHomePage t = new PBFlatTypes.PB_PageResponse_GetHomePage();
    t.setExtra();
    t.setPostViewList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetHomePage t = new PBFlatTypes.PB_PageResponse_GetHomePage();
	t.Extra = ;
	t.PostViewList = ;
	*/

	/*
	PB_PageResponse_GetHomePage t = new PB_PageResponse_GetHomePage();
	t.Extra = m.getExtra() ;
	t.PostViewList = m.getPostViewList() ;
	*/

	public class PB_PageParam_GetProfilePage {
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetProfilePage t = new PBFlatTypes.PB_PageParam_GetProfilePage();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetProfilePage t = new PBFlatTypes.PB_PageParam_GetProfilePage();
	*/

	/*
	PB_PageParam_GetProfilePage t = new PB_PageParam_GetProfilePage();
	*/

	public class PB_PageResponse_GetProfilePage {
	   public PB_ResponseExtra Extra;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetProfilePage t = new PBFlatTypes.PB_PageResponse_GetProfilePage();
    t.setExtra();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetProfilePage t = new PBFlatTypes.PB_PageResponse_GetProfilePage();
	t.Extra = ;
	*/

	/*
	PB_PageResponse_GetProfilePage t = new PB_PageResponse_GetProfilePage();
	t.Extra = m.getExtra() ;
	*/

	public class PB_PageParam_GetLikesPage {
	   public long PostId;
	   public int Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetLikesPage t = new PBFlatTypes.PB_PageParam_GetLikesPage();
    t.setPostId();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetLikesPage t = new PBFlatTypes.PB_PageParam_GetLikesPage();
	t.PostId = ;
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_PageParam_GetLikesPage t = new PB_PageParam_GetLikesPage();
	t.PostId = m.getPostId() ;
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_PageResponse_GetLikesPage {
	   public PB_ResponseExtra Extra;
	   public PB_UserViewRowify UserViewRowifyList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetLikesPage t = new PBFlatTypes.PB_PageResponse_GetLikesPage();
    t.setExtra();
    t.setUserViewRowifyList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetLikesPage t = new PBFlatTypes.PB_PageResponse_GetLikesPage();
	t.Extra = ;
	t.UserViewRowifyList = ;
	*/

	/*
	PB_PageResponse_GetLikesPage t = new PB_PageResponse_GetLikesPage();
	t.Extra = m.getExtra() ;
	t.UserViewRowifyList = m.getUserViewRowifyList() ;
	*/

	public class PB_PageParam_GetFollowersPage {
	   public long UserId;
	   public int Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetFollowersPage t = new PBFlatTypes.PB_PageParam_GetFollowersPage();
    t.setUserId();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetFollowersPage t = new PBFlatTypes.PB_PageParam_GetFollowersPage();
	t.UserId = ;
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_PageParam_GetFollowersPage t = new PB_PageParam_GetFollowersPage();
	t.UserId = m.getUserId() ;
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_PageResponse_GetFollowersPage {
	   public PB_ResponseExtra Extra;
	   public PB_UserViewRowify UserViewRowifyList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetFollowersPage t = new PBFlatTypes.PB_PageResponse_GetFollowersPage();
    t.setExtra();
    t.setUserViewRowifyList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetFollowersPage t = new PBFlatTypes.PB_PageResponse_GetFollowersPage();
	t.Extra = ;
	t.UserViewRowifyList = ;
	*/

	/*
	PB_PageResponse_GetFollowersPage t = new PB_PageResponse_GetFollowersPage();
	t.Extra = m.getExtra() ;
	t.UserViewRowifyList = m.getUserViewRowifyList() ;
	*/

	public class PB_PageParam_GetFollowingsPage {
	   public long UserId;
	   public int Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetFollowingsPage t = new PBFlatTypes.PB_PageParam_GetFollowingsPage();
    t.setUserId();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetFollowingsPage t = new PBFlatTypes.PB_PageParam_GetFollowingsPage();
	t.UserId = ;
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_PageParam_GetFollowingsPage t = new PB_PageParam_GetFollowingsPage();
	t.UserId = m.getUserId() ;
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_PageResponse_GetFollowingsPage {
	   public PB_ResponseExtra Extra;
	   public PB_UserViewRowify UserViewRowifyList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetFollowingsPage t = new PBFlatTypes.PB_PageResponse_GetFollowingsPage();
    t.setExtra();
    t.setUserViewRowifyList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetFollowingsPage t = new PBFlatTypes.PB_PageResponse_GetFollowingsPage();
	t.Extra = ;
	t.UserViewRowifyList = ;
	*/

	/*
	PB_PageResponse_GetFollowingsPage t = new PB_PageResponse_GetFollowingsPage();
	t.Extra = m.getExtra() ;
	t.UserViewRowifyList = m.getUserViewRowifyList() ;
	*/

	public class PB_PageParam_GetNotifiesPage {
	   public int Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetNotifiesPage t = new PBFlatTypes.PB_PageParam_GetNotifiesPage();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetNotifiesPage t = new PBFlatTypes.PB_PageParam_GetNotifiesPage();
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_PageParam_GetNotifiesPage t = new PB_PageParam_GetNotifiesPage();
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_PageResponse_GetNotifiesPage {
	   public PB_ResponseExtra Extra;
	   public PB_NotifyView NotifyViewList;
	   public long RemoveIdsList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetNotifiesPage t = new PBFlatTypes.PB_PageResponse_GetNotifiesPage();
    t.setExtra();
    t.setNotifyViewList();
    t.setRemoveIdsList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetNotifiesPage t = new PBFlatTypes.PB_PageResponse_GetNotifiesPage();
	t.Extra = ;
	t.NotifyViewList = ;
	t.RemoveIdsList = ;
	*/

	/*
	PB_PageResponse_GetNotifiesPage t = new PB_PageResponse_GetNotifiesPage();
	t.Extra = m.getExtra() ;
	t.NotifyViewList = m.getNotifyViewList() ;
	t.RemoveIdsList = m.getRemoveIdsList() ;
	*/

	public class PB_PageParam_GetUserActionsPage {
	   public int Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetUserActionsPage t = new PBFlatTypes.PB_PageParam_GetUserActionsPage();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetUserActionsPage t = new PBFlatTypes.PB_PageParam_GetUserActionsPage();
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_PageParam_GetUserActionsPage t = new PB_PageParam_GetUserActionsPage();
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_PageResponse_GetUserActionsPage {
	   public PB_ResponseExtra Extra;
	   public PB_ActionView ActionViewList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetUserActionsPage t = new PBFlatTypes.PB_PageResponse_GetUserActionsPage();
    t.setExtra();
    t.setActionViewList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetUserActionsPage t = new PBFlatTypes.PB_PageResponse_GetUserActionsPage();
	t.Extra = ;
	t.ActionViewList = ;
	*/

	/*
	PB_PageResponse_GetUserActionsPage t = new PB_PageResponse_GetUserActionsPage();
	t.Extra = m.getExtra() ;
	t.ActionViewList = m.getActionViewList() ;
	*/

	public class PB_PageParam_GetSuggestedPostsPage {
	   public int Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetSuggestedPostsPage t = new PBFlatTypes.PB_PageParam_GetSuggestedPostsPage();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetSuggestedPostsPage t = new PBFlatTypes.PB_PageParam_GetSuggestedPostsPage();
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_PageParam_GetSuggestedPostsPage t = new PB_PageParam_GetSuggestedPostsPage();
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_PageResponse_GetSuggestedPostsPage {
	   public PB_ResponseExtra Extra;
	   public PB_PostViewRowify PostViewRowifyList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetSuggestedPostsPage t = new PBFlatTypes.PB_PageResponse_GetSuggestedPostsPage();
    t.setExtra();
    t.setPostViewRowifyList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetSuggestedPostsPage t = new PBFlatTypes.PB_PageResponse_GetSuggestedPostsPage();
	t.Extra = ;
	t.PostViewRowifyList = ;
	*/

	/*
	PB_PageResponse_GetSuggestedPostsPage t = new PB_PageResponse_GetSuggestedPostsPage();
	t.Extra = m.getExtra() ;
	t.PostViewRowifyList = m.getPostViewRowifyList() ;
	*/

	public class PB_PageParam_GetSuggestedUsersPage {
	   public int Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetSuggestedUsersPage t = new PBFlatTypes.PB_PageParam_GetSuggestedUsersPage();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetSuggestedUsersPage t = new PBFlatTypes.PB_PageParam_GetSuggestedUsersPage();
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_PageParam_GetSuggestedUsersPage t = new PB_PageParam_GetSuggestedUsersPage();
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_PageResponse_GetSuggestedUsersPage {
	   public PB_ResponseExtra Extra;
	   public PB_UserViewRowify UserViewRowifyList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetSuggestedUsersPage t = new PBFlatTypes.PB_PageResponse_GetSuggestedUsersPage();
    t.setExtra();
    t.setUserViewRowifyList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetSuggestedUsersPage t = new PBFlatTypes.PB_PageResponse_GetSuggestedUsersPage();
	t.Extra = ;
	t.UserViewRowifyList = ;
	*/

	/*
	PB_PageResponse_GetSuggestedUsersPage t = new PB_PageResponse_GetSuggestedUsersPage();
	t.Extra = m.getExtra() ;
	t.UserViewRowifyList = m.getUserViewRowifyList() ;
	*/

	public class PB_PageParam_GetSuggestedTagsPage {
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetSuggestedTagsPage t = new PBFlatTypes.PB_PageParam_GetSuggestedTagsPage();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetSuggestedTagsPage t = new PBFlatTypes.PB_PageParam_GetSuggestedTagsPage();
	*/

	/*
	PB_PageParam_GetSuggestedTagsPage t = new PB_PageParam_GetSuggestedTagsPage();
	*/

	public class PB_PageResponse_GetSuggestedTagsPage {
	   public PB_ResponseExtra Extra;
	   public PB_TopTagWithSamplePosts TopTagWithSamplePostsList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetSuggestedTagsPage t = new PBFlatTypes.PB_PageResponse_GetSuggestedTagsPage();
    t.setExtra();
    t.setTopTagWithSamplePostsList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetSuggestedTagsPage t = new PBFlatTypes.PB_PageResponse_GetSuggestedTagsPage();
	t.Extra = ;
	t.TopTagWithSamplePostsList = ;
	*/

	/*
	PB_PageResponse_GetSuggestedTagsPage t = new PB_PageResponse_GetSuggestedTagsPage();
	t.Extra = m.getExtra() ;
	t.TopTagWithSamplePostsList = m.getTopTagWithSamplePostsList() ;
	*/

	public class PB_PageParam_GetLastPostsPage {
	   public int Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetLastPostsPage t = new PBFlatTypes.PB_PageParam_GetLastPostsPage();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetLastPostsPage t = new PBFlatTypes.PB_PageParam_GetLastPostsPage();
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_PageParam_GetLastPostsPage t = new PB_PageParam_GetLastPostsPage();
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_PageResponse_GetLastPostsPage {
	   public PB_ResponseExtra Extra;
	   public PB_PostView PostViewList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetLastPostsPage t = new PBFlatTypes.PB_PageResponse_GetLastPostsPage();
    t.setExtra();
    t.setPostViewList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetLastPostsPage t = new PBFlatTypes.PB_PageResponse_GetLastPostsPage();
	t.Extra = ;
	t.PostViewList = ;
	*/

	/*
	PB_PageResponse_GetLastPostsPage t = new PB_PageResponse_GetLastPostsPage();
	t.Extra = m.getExtra() ;
	t.PostViewList = m.getPostViewList() ;
	*/

	public class PB_PageParam_GetTagPage {
	   public String Tag;
	   public int Limit;
	   public long Last;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_GetTagPage t = new PBFlatTypes.PB_PageParam_GetTagPage();
    t.setTag();
    t.setLimit();
    t.setLast();
	*/

	/*
	PBFlatTypes.PB_PageParam_GetTagPage t = new PBFlatTypes.PB_PageParam_GetTagPage();
	t.Tag = ;
	t.Limit = ;
	t.Last = ;
	*/

	/*
	PB_PageParam_GetTagPage t = new PB_PageParam_GetTagPage();
	t.Tag = m.getTag() ;
	t.Limit = m.getLimit() ;
	t.Last = m.getLast() ;
	*/

	public class PB_PageResponse_GetTagPage {
	   public PB_ResponseExtra Extra;
	   public PB_PostView PostViewList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_GetTagPage t = new PBFlatTypes.PB_PageResponse_GetTagPage();
    t.setExtra();
    t.setPostViewList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_GetTagPage t = new PBFlatTypes.PB_PageResponse_GetTagPage();
	t.Extra = ;
	t.PostViewList = ;
	*/

	/*
	PB_PageResponse_GetTagPage t = new PB_PageResponse_GetTagPage();
	t.Extra = m.getExtra() ;
	t.PostViewList = m.getPostViewList() ;
	*/

	public class PB_PageParam_SearchTagsPage {
	   public String Query;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_SearchTagsPage t = new PBFlatTypes.PB_PageParam_SearchTagsPage();
    t.setQuery();
	*/

	/*
	PBFlatTypes.PB_PageParam_SearchTagsPage t = new PBFlatTypes.PB_PageParam_SearchTagsPage();
	t.Query = ;
	*/

	/*
	PB_PageParam_SearchTagsPage t = new PB_PageParam_SearchTagsPage();
	t.Query = m.getQuery() ;
	*/

	public class PB_PageResponse_SearchTagsPage {
	   public PB_ResponseExtra Extra;
	   public PB_TagView TagViewList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_SearchTagsPage t = new PBFlatTypes.PB_PageResponse_SearchTagsPage();
    t.setExtra();
    t.setTagViewList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_SearchTagsPage t = new PBFlatTypes.PB_PageResponse_SearchTagsPage();
	t.Extra = ;
	t.TagViewList = ;
	*/

	/*
	PB_PageResponse_SearchTagsPage t = new PB_PageResponse_SearchTagsPage();
	t.Extra = m.getExtra() ;
	t.TagViewList = m.getTagViewList() ;
	*/

	public class PB_PageParam_SearchUsersPage {
	   public String Query;
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_SearchUsersPage t = new PBFlatTypes.PB_PageParam_SearchUsersPage();
    t.setQuery();
	*/

	/*
	PBFlatTypes.PB_PageParam_SearchUsersPage t = new PBFlatTypes.PB_PageParam_SearchUsersPage();
	t.Query = ;
	*/

	/*
	PB_PageParam_SearchUsersPage t = new PB_PageParam_SearchUsersPage();
	t.Query = m.getQuery() ;
	*/

	public class PB_PageResponse_SearchUsersPage {
	   public PB_ResponseExtra Extra;
	   public PB_UserView UserViewList;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_SearchUsersPage t = new PBFlatTypes.PB_PageResponse_SearchUsersPage();
    t.setExtra();
    t.setUserViewList();
	*/

	/*
	PBFlatTypes.PB_PageResponse_SearchUsersPage t = new PBFlatTypes.PB_PageResponse_SearchUsersPage();
	t.Extra = ;
	t.UserViewList = ;
	*/

	/*
	PB_PageResponse_SearchUsersPage t = new PB_PageResponse_SearchUsersPage();
	t.Extra = m.getExtra() ;
	t.UserViewList = m.getUserViewList() ;
	*/

	public class PB_PageParam_ {
	}
	/*
	folding
	PBFlatTypes.PB_PageParam_ t = new PBFlatTypes.PB_PageParam_();
	*/

	/*
	PBFlatTypes.PB_PageParam_ t = new PBFlatTypes.PB_PageParam_();
	*/

	/*
	PB_PageParam_ t = new PB_PageParam_();
	*/

	public class PB_PageResponse_ {
	   public PB_ResponseExtra Extra;
	}
	/*
	folding
	PBFlatTypes.PB_PageResponse_ t = new PBFlatTypes.PB_PageResponse_();
    t.setExtra();
	*/

	/*
	PBFlatTypes.PB_PageResponse_ t = new PBFlatTypes.PB_PageResponse_();
	t.Extra = ;
	*/

	/*
	PB_PageResponse_ t = new PB_PageResponse_();
	t.Extra = m.getExtra() ;
	*/

	public class PB_SearchResponse_AddNewC {
	}
	/*
	folding
	PBFlatTypes.PB_SearchResponse_AddNewC t = new PBFlatTypes.PB_SearchResponse_AddNewC();
	*/

	/*
	PBFlatTypes.PB_SearchResponse_AddNewC t = new PBFlatTypes.PB_SearchResponse_AddNewC();
	*/

	/*
	PB_SearchResponse_AddNewC t = new PB_SearchResponse_AddNewC();
	*/

	public class PB_SocialParam_AddComment {
	   public long PostId;
	   public String Text;
	}
	/*
	folding
	PBFlatTypes.PB_SocialParam_AddComment t = new PBFlatTypes.PB_SocialParam_AddComment();
    t.setPostId();
    t.setText();
	*/

	/*
	PBFlatTypes.PB_SocialParam_AddComment t = new PBFlatTypes.PB_SocialParam_AddComment();
	t.PostId = ;
	t.Text = ;
	*/

	/*
	PB_SocialParam_AddComment t = new PB_SocialParam_AddComment();
	t.PostId = m.getPostId() ;
	t.Text = m.getText() ;
	*/

	public class PB_SocialResponse_AddComment {
	   public PB_ResponseExtra Extra;
	   public PB_CommentView Comment;
	}
	/*
	folding
	PBFlatTypes.PB_SocialResponse_AddComment t = new PBFlatTypes.PB_SocialResponse_AddComment();
    t.setExtra();
    t.setComment();
	*/

	/*
	PBFlatTypes.PB_SocialResponse_AddComment t = new PBFlatTypes.PB_SocialResponse_AddComment();
	t.Extra = ;
	t.Comment = ;
	*/

	/*
	PB_SocialResponse_AddComment t = new PB_SocialResponse_AddComment();
	t.Extra = m.getExtra() ;
	t.Comment = m.getComment() ;
	*/

	public class PB_SocialParam_DeleteComment {
	   public long PostId;
	   public long CommentId;
	}
	/*
	folding
	PBFlatTypes.PB_SocialParam_DeleteComment t = new PBFlatTypes.PB_SocialParam_DeleteComment();
    t.setPostId();
    t.setCommentId();
	*/

	/*
	PBFlatTypes.PB_SocialParam_DeleteComment t = new PBFlatTypes.PB_SocialParam_DeleteComment();
	t.PostId = ;
	t.CommentId = ;
	*/

	/*
	PB_SocialParam_DeleteComment t = new PB_SocialParam_DeleteComment();
	t.PostId = m.getPostId() ;
	t.CommentId = m.getCommentId() ;
	*/

	public class PB_SocialResponse_DeleteComment {
	   public PB_ResponseExtra Extra;
	   public boolean Deleted;
	}
	/*
	folding
	PBFlatTypes.PB_SocialResponse_DeleteComment t = new PBFlatTypes.PB_SocialResponse_DeleteComment();
    t.setExtra();
    t.setDeleted();
	*/

	/*
	PBFlatTypes.PB_SocialResponse_DeleteComment t = new PBFlatTypes.PB_SocialResponse_DeleteComment();
	t.Extra = ;
	t.Deleted = ;
	*/

	/*
	PB_SocialResponse_DeleteComment t = new PB_SocialResponse_DeleteComment();
	t.Extra = m.getExtra() ;
	t.Deleted = m.getDeleted() ;
	*/

	public class PB_SocialParam_AddPost {
	   public String Text;
	}
	/*
	folding
	PBFlatTypes.PB_SocialParam_AddPost t = new PBFlatTypes.PB_SocialParam_AddPost();
    t.setText();
	*/

	/*
	PBFlatTypes.PB_SocialParam_AddPost t = new PBFlatTypes.PB_SocialParam_AddPost();
	t.Text = ;
	*/

	/*
	PB_SocialParam_AddPost t = new PB_SocialParam_AddPost();
	t.Text = m.getText() ;
	*/

	public class PB_SocialResponse_AddPost {
	   public PB_ResponseExtra Extra;
	   public PB_PostView PostView;
	}
	/*
	folding
	PBFlatTypes.PB_SocialResponse_AddPost t = new PBFlatTypes.PB_SocialResponse_AddPost();
    t.setExtra();
    t.setPostView();
	*/

	/*
	PBFlatTypes.PB_SocialResponse_AddPost t = new PBFlatTypes.PB_SocialResponse_AddPost();
	t.Extra = ;
	t.PostView = ;
	*/

	/*
	PB_SocialResponse_AddPost t = new PB_SocialResponse_AddPost();
	t.Extra = m.getExtra() ;
	t.PostView = m.getPostView() ;
	*/

	public class PB_SocialParam_EditPost {
	   public long PostId;
	   public String Text;
	}
	/*
	folding
	PBFlatTypes.PB_SocialParam_EditPost t = new PBFlatTypes.PB_SocialParam_EditPost();
    t.setPostId();
    t.setText();
	*/

	/*
	PBFlatTypes.PB_SocialParam_EditPost t = new PBFlatTypes.PB_SocialParam_EditPost();
	t.PostId = ;
	t.Text = ;
	*/

	/*
	PB_SocialParam_EditPost t = new PB_SocialParam_EditPost();
	t.PostId = m.getPostId() ;
	t.Text = m.getText() ;
	*/

	public class PB_SocialResponse_EditPost {
	   public PB_ResponseExtra Extra;
	   public PB_PostView PostView;
	}
	/*
	folding
	PBFlatTypes.PB_SocialResponse_EditPost t = new PBFlatTypes.PB_SocialResponse_EditPost();
    t.setExtra();
    t.setPostView();
	*/

	/*
	PBFlatTypes.PB_SocialResponse_EditPost t = new PBFlatTypes.PB_SocialResponse_EditPost();
	t.Extra = ;
	t.PostView = ;
	*/

	/*
	PB_SocialResponse_EditPost t = new PB_SocialResponse_EditPost();
	t.Extra = m.getExtra() ;
	t.PostView = m.getPostView() ;
	*/

	public class PB_SocialParam_DeletePost {
	   public long PostId;
	}
	/*
	folding
	PBFlatTypes.PB_SocialParam_DeletePost t = new PBFlatTypes.PB_SocialParam_DeletePost();
    t.setPostId();
	*/

	/*
	PBFlatTypes.PB_SocialParam_DeletePost t = new PBFlatTypes.PB_SocialParam_DeletePost();
	t.PostId = ;
	*/

	/*
	PB_SocialParam_DeletePost t = new PB_SocialParam_DeletePost();
	t.PostId = m.getPostId() ;
	*/

	public class PB_SocialResponse_DeletePost {
	   public PB_ResponseExtra Extra;
	   public boolean Done;
	}
	/*
	folding
	PBFlatTypes.PB_SocialResponse_DeletePost t = new PBFlatTypes.PB_SocialResponse_DeletePost();
    t.setExtra();
    t.setDone();
	*/

	/*
	PBFlatTypes.PB_SocialResponse_DeletePost t = new PBFlatTypes.PB_SocialResponse_DeletePost();
	t.Extra = ;
	t.Done = ;
	*/

	/*
	PB_SocialResponse_DeletePost t = new PB_SocialResponse_DeletePost();
	t.Extra = m.getExtra() ;
	t.Done = m.getDone() ;
	*/

	public class PB_SocialParam_ArchivePost {
	   public long PostId;
	}
	/*
	folding
	PBFlatTypes.PB_SocialParam_ArchivePost t = new PBFlatTypes.PB_SocialParam_ArchivePost();
    t.setPostId();
	*/

	/*
	PBFlatTypes.PB_SocialParam_ArchivePost t = new PBFlatTypes.PB_SocialParam_ArchivePost();
	t.PostId = ;
	*/

	/*
	PB_SocialParam_ArchivePost t = new PB_SocialParam_ArchivePost();
	t.PostId = m.getPostId() ;
	*/

	public class PB_SocialResponse_ArchivePost {
	   public PB_ResponseExtra Extra;
	   public boolean Done;
	}
	/*
	folding
	PBFlatTypes.PB_SocialResponse_ArchivePost t = new PBFlatTypes.PB_SocialResponse_ArchivePost();
    t.setExtra();
    t.setDone();
	*/

	/*
	PBFlatTypes.PB_SocialResponse_ArchivePost t = new PBFlatTypes.PB_SocialResponse_ArchivePost();
	t.Extra = ;
	t.Done = ;
	*/

	/*
	PB_SocialResponse_ArchivePost t = new PB_SocialResponse_ArchivePost();
	t.Extra = m.getExtra() ;
	t.Done = m.getDone() ;
	*/

	public class PB_SocialParam_LikePost {
	   public long PostId;
	}
	/*
	folding
	PBFlatTypes.PB_SocialParam_LikePost t = new PBFlatTypes.PB_SocialParam_LikePost();
    t.setPostId();
	*/

	/*
	PBFlatTypes.PB_SocialParam_LikePost t = new PBFlatTypes.PB_SocialParam_LikePost();
	t.PostId = ;
	*/

	/*
	PB_SocialParam_LikePost t = new PB_SocialParam_LikePost();
	t.PostId = m.getPostId() ;
	*/

	public class PB_SocialResponse_LikePost {
	   public PB_ResponseExtra Extra;
	   public boolean Done;
	}
	/*
	folding
	PBFlatTypes.PB_SocialResponse_LikePost t = new PBFlatTypes.PB_SocialResponse_LikePost();
    t.setExtra();
    t.setDone();
	*/

	/*
	PBFlatTypes.PB_SocialResponse_LikePost t = new PBFlatTypes.PB_SocialResponse_LikePost();
	t.Extra = ;
	t.Done = ;
	*/

	/*
	PB_SocialResponse_LikePost t = new PB_SocialResponse_LikePost();
	t.Extra = m.getExtra() ;
	t.Done = m.getDone() ;
	*/

	public class PB_SocialParam_UnLikePost {
	   public long PostId;
	}
	/*
	folding
	PBFlatTypes.PB_SocialParam_UnLikePost t = new PBFlatTypes.PB_SocialParam_UnLikePost();
    t.setPostId();
	*/

	/*
	PBFlatTypes.PB_SocialParam_UnLikePost t = new PBFlatTypes.PB_SocialParam_UnLikePost();
	t.PostId = ;
	*/

	/*
	PB_SocialParam_UnLikePost t = new PB_SocialParam_UnLikePost();
	t.PostId = m.getPostId() ;
	*/

	public class PB_SocialResponse_UnLikePost {
	   public PB_ResponseExtra Extra;
	   public boolean Done;
	}
	/*
	folding
	PBFlatTypes.PB_SocialResponse_UnLikePost t = new PBFlatTypes.PB_SocialResponse_UnLikePost();
    t.setExtra();
    t.setDone();
	*/

	/*
	PBFlatTypes.PB_SocialResponse_UnLikePost t = new PBFlatTypes.PB_SocialResponse_UnLikePost();
	t.Extra = ;
	t.Done = ;
	*/

	/*
	PB_SocialResponse_UnLikePost t = new PB_SocialResponse_UnLikePost();
	t.Extra = m.getExtra() ;
	t.Done = m.getDone() ;
	*/

	public class PB_SocialParam_FollowUser {
	   public long UserId;
	}
	/*
	folding
	PBFlatTypes.PB_SocialParam_FollowUser t = new PBFlatTypes.PB_SocialParam_FollowUser();
    t.setUserId();
	*/

	/*
	PBFlatTypes.PB_SocialParam_FollowUser t = new PBFlatTypes.PB_SocialParam_FollowUser();
	t.UserId = ;
	*/

	/*
	PB_SocialParam_FollowUser t = new PB_SocialParam_FollowUser();
	t.UserId = m.getUserId() ;
	*/

	public class PB_SocialResponse_FollowUser {
	   public PB_ResponseExtra Extra;
	}
	/*
	folding
	PBFlatTypes.PB_SocialResponse_FollowUser t = new PBFlatTypes.PB_SocialResponse_FollowUser();
    t.setExtra();
	*/

	/*
	PBFlatTypes.PB_SocialResponse_FollowUser t = new PBFlatTypes.PB_SocialResponse_FollowUser();
	t.Extra = ;
	*/

	/*
	PB_SocialResponse_FollowUser t = new PB_SocialResponse_FollowUser();
	t.Extra = m.getExtra() ;
	*/

	public class PB_SocialParam_UnFollowUser {
	   public long UserId;
	}
	/*
	folding
	PBFlatTypes.PB_SocialParam_UnFollowUser t = new PBFlatTypes.PB_SocialParam_UnFollowUser();
    t.setUserId();
	*/

	/*
	PBFlatTypes.PB_SocialParam_UnFollowUser t = new PBFlatTypes.PB_SocialParam_UnFollowUser();
	t.UserId = ;
	*/

	/*
	PB_SocialParam_UnFollowUser t = new PB_SocialParam_UnFollowUser();
	t.UserId = m.getUserId() ;
	*/

	public class PB_SocialResponse_UnFollowUser {
	   public PB_ResponseExtra Extra;
	}
	/*
	folding
	PBFlatTypes.PB_SocialResponse_UnFollowUser t = new PBFlatTypes.PB_SocialResponse_UnFollowUser();
    t.setExtra();
	*/

	/*
	PBFlatTypes.PB_SocialResponse_UnFollowUser t = new PBFlatTypes.PB_SocialResponse_UnFollowUser();
	t.Extra = ;
	*/

	/*
	PB_SocialResponse_UnFollowUser t = new PB_SocialResponse_UnFollowUser();
	t.Extra = m.getExtra() ;
	*/

	public class PB_UserParam_BlockUser {
	   public int UserId;
	   public String UserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_BlockUser t = new PBFlatTypes.PB_UserParam_BlockUser();
    t.setUserId();
    t.setUserName();
	*/

	/*
	PBFlatTypes.PB_UserParam_BlockUser t = new PBFlatTypes.PB_UserParam_BlockUser();
	t.UserId = ;
	t.UserName = ;
	*/

	/*
	PB_UserParam_BlockUser t = new PB_UserParam_BlockUser();
	t.UserId = m.getUserId() ;
	t.UserName = m.getUserName() ;
	*/

	public class PB_UserResponse_BlockUser {
	   public int ByUserId;
	   public int TargetUserId;
	   public String TargetUserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserResponse_BlockUser t = new PBFlatTypes.PB_UserResponse_BlockUser();
    t.setByUserId();
    t.setTargetUserId();
    t.setTargetUserName();
	*/

	/*
	PBFlatTypes.PB_UserResponse_BlockUser t = new PBFlatTypes.PB_UserResponse_BlockUser();
	t.ByUserId = ;
	t.TargetUserId = ;
	t.TargetUserName = ;
	*/

	/*
	PB_UserResponse_BlockUser t = new PB_UserResponse_BlockUser();
	t.ByUserId = m.getByUserId() ;
	t.TargetUserId = m.getTargetUserId() ;
	t.TargetUserName = m.getTargetUserName() ;
	*/

	public class PB_UserParam_UnBlockUser {
	   public int Offset;
	   public int Limit;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_UnBlockUser t = new PBFlatTypes.PB_UserParam_UnBlockUser();
    t.setOffset();
    t.setLimit();
	*/

	/*
	PBFlatTypes.PB_UserParam_UnBlockUser t = new PBFlatTypes.PB_UserParam_UnBlockUser();
	t.Offset = ;
	t.Limit = ;
	*/

	/*
	PB_UserParam_UnBlockUser t = new PB_UserParam_UnBlockUser();
	t.Offset = m.getOffset() ;
	t.Limit = m.getLimit() ;
	*/

	public class PB_UserResponse_UnBlockUser {
	   public UserView Users;
	}
	/*
	folding
	PBFlatTypes.PB_UserResponse_UnBlockUser t = new PBFlatTypes.PB_UserResponse_UnBlockUser();
    t.setUsers();
	*/

	/*
	PBFlatTypes.PB_UserResponse_UnBlockUser t = new PBFlatTypes.PB_UserResponse_UnBlockUser();
	t.Users = ;
	*/

	/*
	PB_UserResponse_UnBlockUser t = new PB_UserResponse_UnBlockUser();
	t.Users = m.getUsers() ;
	*/

	public class PB_UserParam_BlockedList {
	   public int UserId;
	   public String UserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_BlockedList t = new PBFlatTypes.PB_UserParam_BlockedList();
    t.setUserId();
    t.setUserName();
	*/

	/*
	PBFlatTypes.PB_UserParam_BlockedList t = new PBFlatTypes.PB_UserParam_BlockedList();
	t.UserId = ;
	t.UserName = ;
	*/

	/*
	PB_UserParam_BlockedList t = new PB_UserParam_BlockedList();
	t.UserId = m.getUserId() ;
	t.UserName = m.getUserName() ;
	*/

	public class PB_UserResponse_BlockedList {
	   public int ByUserId;
	   public int TargetUserId;
	   public String TargetUserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserResponse_BlockedList t = new PBFlatTypes.PB_UserResponse_BlockedList();
    t.setByUserId();
    t.setTargetUserId();
    t.setTargetUserName();
	*/

	/*
	PBFlatTypes.PB_UserResponse_BlockedList t = new PBFlatTypes.PB_UserResponse_BlockedList();
	t.ByUserId = ;
	t.TargetUserId = ;
	t.TargetUserName = ;
	*/

	/*
	PB_UserResponse_BlockedList t = new PB_UserResponse_BlockedList();
	t.ByUserId = m.getByUserId() ;
	t.TargetUserId = m.getTargetUserId() ;
	t.TargetUserName = m.getTargetUserName() ;
	*/

	public class PB_UserParam_UpdateAbout {
	   public String NewAbout;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_UpdateAbout t = new PBFlatTypes.PB_UserParam_UpdateAbout();
    t.setNewAbout();
	*/

	/*
	PBFlatTypes.PB_UserParam_UpdateAbout t = new PBFlatTypes.PB_UserParam_UpdateAbout();
	t.NewAbout = ;
	*/

	/*
	PB_UserParam_UpdateAbout t = new PB_UserParam_UpdateAbout();
	t.NewAbout = m.getNewAbout() ;
	*/

	public class PB_UserResponse_UpdateAbout {
	   public int UserId;
	   public String NewAbout;
	}
	/*
	folding
	PBFlatTypes.PB_UserResponse_UpdateAbout t = new PBFlatTypes.PB_UserResponse_UpdateAbout();
    t.setUserId();
    t.setNewAbout();
	*/

	/*
	PBFlatTypes.PB_UserResponse_UpdateAbout t = new PBFlatTypes.PB_UserResponse_UpdateAbout();
	t.UserId = ;
	t.NewAbout = ;
	*/

	/*
	PB_UserResponse_UpdateAbout t = new PB_UserResponse_UpdateAbout();
	t.UserId = m.getUserId() ;
	t.NewAbout = m.getNewAbout() ;
	*/

	public class PB_UserParam_UpdateUserName {
	   public String NewUserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_UpdateUserName t = new PBFlatTypes.PB_UserParam_UpdateUserName();
    t.setNewUserName();
	*/

	/*
	PBFlatTypes.PB_UserParam_UpdateUserName t = new PBFlatTypes.PB_UserParam_UpdateUserName();
	t.NewUserName = ;
	*/

	/*
	PB_UserParam_UpdateUserName t = new PB_UserParam_UpdateUserName();
	t.NewUserName = m.getNewUserName() ;
	*/

	public class PB_UserResponse_UpdateUserName {
	   public int UserId;
	   public String NewUserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserResponse_UpdateUserName t = new PBFlatTypes.PB_UserResponse_UpdateUserName();
    t.setUserId();
    t.setNewUserName();
	*/

	/*
	PBFlatTypes.PB_UserResponse_UpdateUserName t = new PBFlatTypes.PB_UserResponse_UpdateUserName();
	t.UserId = ;
	t.NewUserName = ;
	*/

	/*
	PB_UserResponse_UpdateUserName t = new PB_UserResponse_UpdateUserName();
	t.UserId = m.getUserId() ;
	t.NewUserName = m.getNewUserName() ;
	*/

	public class PB_UserParam_ChangeAvatar {
	   public boolean None;
	   public byte[] ImageData2;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_ChangeAvatar t = new PBFlatTypes.PB_UserParam_ChangeAvatar();
    t.setNone();
    t.setImageData2();
	*/

	/*
	PBFlatTypes.PB_UserParam_ChangeAvatar t = new PBFlatTypes.PB_UserParam_ChangeAvatar();
	t.None = ;
	t.ImageData2 = ;
	*/

	/*
	PB_UserParam_ChangeAvatar t = new PB_UserParam_ChangeAvatar();
	t.None = m.getNone() ;
	t.ImageData2 = m.getImageData2() ;
	*/

	public class PB_UserResponse_ChangeAvatar {
	}
	/*
	folding
	PBFlatTypes.PB_UserResponse_ChangeAvatar t = new PBFlatTypes.PB_UserResponse_ChangeAvatar();
	*/

	/*
	PBFlatTypes.PB_UserResponse_ChangeAvatar t = new PBFlatTypes.PB_UserResponse_ChangeAvatar();
	*/

	/*
	PB_UserResponse_ChangeAvatar t = new PB_UserResponse_ChangeAvatar();
	*/

	public class PB_UserParam_ChangePrivacy {
	   public ProfilePrivacyLevelEnum Level;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_ChangePrivacy t = new PBFlatTypes.PB_UserParam_ChangePrivacy();
    t.setLevel();
	*/

	/*
	PBFlatTypes.PB_UserParam_ChangePrivacy t = new PBFlatTypes.PB_UserParam_ChangePrivacy();
	t.Level = ;
	*/

	/*
	PB_UserParam_ChangePrivacy t = new PB_UserParam_ChangePrivacy();
	t.Level = m.getLevel() ;
	*/

	public class PB_UserResponseOffline_ChangePrivacy {
	}
	/*
	folding
	PBFlatTypes.PB_UserResponseOffline_ChangePrivacy t = new PBFlatTypes.PB_UserResponseOffline_ChangePrivacy();
	*/

	/*
	PBFlatTypes.PB_UserResponseOffline_ChangePrivacy t = new PBFlatTypes.PB_UserResponseOffline_ChangePrivacy();
	*/

	/*
	PB_UserResponseOffline_ChangePrivacy t = new PB_UserResponseOffline_ChangePrivacy();
	*/

	public class PB_UserParam_CheckUserName {
	   public ProfilePrivacyLevelEnum Level;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_CheckUserName t = new PBFlatTypes.PB_UserParam_CheckUserName();
    t.setLevel();
	*/

	/*
	PBFlatTypes.PB_UserParam_CheckUserName t = new PBFlatTypes.PB_UserParam_CheckUserName();
	t.Level = ;
	*/

	/*
	PB_UserParam_CheckUserName t = new PB_UserParam_CheckUserName();
	t.Level = m.getLevel() ;
	*/

	public class PB_UserResponse_CheckUserName {
	}
	/*
	folding
	PBFlatTypes.PB_UserResponse_CheckUserName t = new PBFlatTypes.PB_UserResponse_CheckUserName();
	*/

	/*
	PBFlatTypes.PB_UserResponse_CheckUserName t = new PBFlatTypes.PB_UserResponse_CheckUserName();
	*/

	/*
	PB_UserResponse_CheckUserName t = new PB_UserResponse_CheckUserName();
	*/

	public class UserView {
	}
	/*
	folding
	PBFlatTypes.UserView t = new PBFlatTypes.UserView();
	*/

	/*
	PBFlatTypes.UserView t = new PBFlatTypes.UserView();
	*/

	/*
	UserView t = new UserView();
	*/

	public class PB_Action {
	   public long ActionId;
	   public int ActorUserId;
	   public int ActionTypeEnum;
	   public int PeerUserId;
	   public long PostId;
	   public long CommentId;
	   public long Murmur64Hash;
	   public int CreatedTime;
	   public int Seq;
	}
	/*
	folding
	PBFlatTypes.PB_Action t = new PBFlatTypes.PB_Action();
    t.setActionId();
    t.setActorUserId();
    t.setActionTypeEnum();
    t.setPeerUserId();
    t.setPostId();
    t.setCommentId();
    t.setMurmur64Hash();
    t.setCreatedTime();
    t.setSeq();
	*/

	/*
	PBFlatTypes.PB_Action t = new PBFlatTypes.PB_Action();
	t.ActionId = ;
	t.ActorUserId = ;
	t.ActionTypeEnum = ;
	t.PeerUserId = ;
	t.PostId = ;
	t.CommentId = ;
	t.Murmur64Hash = ;
	t.CreatedTime = ;
	t.Seq = ;
	*/

	/*
	PB_Action t = new PB_Action();
	t.ActionId = m.getActionId() ;
	t.ActorUserId = m.getActorUserId() ;
	t.ActionTypeEnum = m.getActionTypeEnum() ;
	t.PeerUserId = m.getPeerUserId() ;
	t.PostId = m.getPostId() ;
	t.CommentId = m.getCommentId() ;
	t.Murmur64Hash = m.getMurmur64Hash() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.Seq = m.getSeq() ;
	*/

	public class PB_Comment {
	   public long CommentId;
	   public int UserId;
	   public long PostId;
	   public String Text;
	   public int LikesCount;
	   public int CreatedTime;
	   public int Seq;
	}
	/*
	folding
	PBFlatTypes.PB_Comment t = new PBFlatTypes.PB_Comment();
    t.setCommentId();
    t.setUserId();
    t.setPostId();
    t.setText();
    t.setLikesCount();
    t.setCreatedTime();
    t.setSeq();
	*/

	/*
	PBFlatTypes.PB_Comment t = new PBFlatTypes.PB_Comment();
	t.CommentId = ;
	t.UserId = ;
	t.PostId = ;
	t.Text = ;
	t.LikesCount = ;
	t.CreatedTime = ;
	t.Seq = ;
	*/

	/*
	PB_Comment t = new PB_Comment();
	t.CommentId = m.getCommentId() ;
	t.UserId = m.getUserId() ;
	t.PostId = m.getPostId() ;
	t.Text = m.getText() ;
	t.LikesCount = m.getLikesCount() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.Seq = m.getSeq() ;
	*/

	public class PB_FollowingList {
	   public int Id;
	   public int UserId;
	   public int ListType;
	   public String Name;
	   public int Count;
	   public int IsAuto;
	   public int IsPimiry;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_FollowingList t = new PBFlatTypes.PB_FollowingList();
    t.setId();
    t.setUserId();
    t.setListType();
    t.setName();
    t.setCount();
    t.setIsAuto();
    t.setIsPimiry();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_FollowingList t = new PBFlatTypes.PB_FollowingList();
	t.Id = ;
	t.UserId = ;
	t.ListType = ;
	t.Name = ;
	t.Count = ;
	t.IsAuto = ;
	t.IsPimiry = ;
	t.CreatedTime = ;
	*/

	/*
	PB_FollowingList t = new PB_FollowingList();
	t.Id = m.getId() ;
	t.UserId = m.getUserId() ;
	t.ListType = m.getListType() ;
	t.Name = m.getName() ;
	t.Count = m.getCount() ;
	t.IsAuto = m.getIsAuto() ;
	t.IsPimiry = m.getIsPimiry() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_FollowingListMember {
	   public long Id;
	   public int ListId;
	   public int UserId;
	   public int FollowedUserId;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_FollowingListMember t = new PBFlatTypes.PB_FollowingListMember();
    t.setId();
    t.setListId();
    t.setUserId();
    t.setFollowedUserId();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_FollowingListMember t = new PBFlatTypes.PB_FollowingListMember();
	t.Id = ;
	t.ListId = ;
	t.UserId = ;
	t.FollowedUserId = ;
	t.CreatedTime = ;
	*/

	/*
	PB_FollowingListMember t = new PB_FollowingListMember();
	t.Id = m.getId() ;
	t.ListId = m.getListId() ;
	t.UserId = m.getUserId() ;
	t.FollowedUserId = m.getFollowedUserId() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_FollowingListMemberRemoved {
	   public long Id;
	   public int ListId;
	   public int UserId;
	   public int UnFollowedUserId;
	   public int UpdatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_FollowingListMemberRemoved t = new PBFlatTypes.PB_FollowingListMemberRemoved();
    t.setId();
    t.setListId();
    t.setUserId();
    t.setUnFollowedUserId();
    t.setUpdatedTime();
	*/

	/*
	PBFlatTypes.PB_FollowingListMemberRemoved t = new PBFlatTypes.PB_FollowingListMemberRemoved();
	t.Id = ;
	t.ListId = ;
	t.UserId = ;
	t.UnFollowedUserId = ;
	t.UpdatedTime = ;
	*/

	/*
	PB_FollowingListMemberRemoved t = new PB_FollowingListMemberRemoved();
	t.Id = m.getId() ;
	t.ListId = m.getListId() ;
	t.UserId = m.getUserId() ;
	t.UnFollowedUserId = m.getUnFollowedUserId() ;
	t.UpdatedTime = m.getUpdatedTime() ;
	*/

	public class PB_Group {
	   public long GroupId;
	   public String GroupName;
	   public int MembersCount;
	   public int GroupPrivacyEnum;
	   public int CreatorUserId;
	   public int CreatedTime;
	   public long UpdatedMs;
	   public int CurrentSeq;
	}
	/*
	folding
	PBFlatTypes.PB_Group t = new PBFlatTypes.PB_Group();
    t.setGroupId();
    t.setGroupName();
    t.setMembersCount();
    t.setGroupPrivacyEnum();
    t.setCreatorUserId();
    t.setCreatedTime();
    t.setUpdatedMs();
    t.setCurrentSeq();
	*/

	/*
	PBFlatTypes.PB_Group t = new PBFlatTypes.PB_Group();
	t.GroupId = ;
	t.GroupName = ;
	t.MembersCount = ;
	t.GroupPrivacyEnum = ;
	t.CreatorUserId = ;
	t.CreatedTime = ;
	t.UpdatedMs = ;
	t.CurrentSeq = ;
	*/

	/*
	PB_Group t = new PB_Group();
	t.GroupId = m.getGroupId() ;
	t.GroupName = m.getGroupName() ;
	t.MembersCount = m.getMembersCount() ;
	t.GroupPrivacyEnum = m.getGroupPrivacyEnum() ;
	t.CreatorUserId = m.getCreatorUserId() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.UpdatedMs = m.getUpdatedMs() ;
	t.CurrentSeq = m.getCurrentSeq() ;
	*/

	public class PB_GroupMember {
	   public long Id;
	   public long GroupId;
	   public String GroupKey;
	   public int UserId;
	   public int ByUserId;
	   public int GroupRoleEnumId;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_GroupMember t = new PBFlatTypes.PB_GroupMember();
    t.setId();
    t.setGroupId();
    t.setGroupKey();
    t.setUserId();
    t.setByUserId();
    t.setGroupRoleEnumId();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_GroupMember t = new PBFlatTypes.PB_GroupMember();
	t.Id = ;
	t.GroupId = ;
	t.GroupKey = ;
	t.UserId = ;
	t.ByUserId = ;
	t.GroupRoleEnumId = ;
	t.CreatedTime = ;
	*/

	/*
	PB_GroupMember t = new PB_GroupMember();
	t.Id = m.getId() ;
	t.GroupId = m.getGroupId() ;
	t.GroupKey = m.getGroupKey() ;
	t.UserId = m.getUserId() ;
	t.ByUserId = m.getByUserId() ;
	t.GroupRoleEnumId = m.getGroupRoleEnumId() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_GroupMessage {
	   public long MessageId;
	   public String RoomKey;
	   public int UserId;
	   public long MessageFileId;
	   public int MessageTypeEnum;
	   public String Text;
	   public long CreatedMs;
	   public int DeliveryStatusEnum;
	}
	/*
	folding
	PBFlatTypes.PB_GroupMessage t = new PBFlatTypes.PB_GroupMessage();
    t.setMessageId();
    t.setRoomKey();
    t.setUserId();
    t.setMessageFileId();
    t.setMessageTypeEnum();
    t.setText();
    t.setCreatedMs();
    t.setDeliveryStatusEnum();
	*/

	/*
	PBFlatTypes.PB_GroupMessage t = new PBFlatTypes.PB_GroupMessage();
	t.MessageId = ;
	t.RoomKey = ;
	t.UserId = ;
	t.MessageFileId = ;
	t.MessageTypeEnum = ;
	t.Text = ;
	t.CreatedMs = ;
	t.DeliveryStatusEnum = ;
	*/

	/*
	PB_GroupMessage t = new PB_GroupMessage();
	t.MessageId = m.getMessageId() ;
	t.RoomKey = m.getRoomKey() ;
	t.UserId = m.getUserId() ;
	t.MessageFileId = m.getMessageFileId() ;
	t.MessageTypeEnum = m.getMessageTypeEnum() ;
	t.Text = m.getText() ;
	t.CreatedMs = m.getCreatedMs() ;
	t.DeliveryStatusEnum = m.getDeliveryStatusEnum() ;
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

	public class PB_Media {
	   public long MediaId;
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
	/*
	folding
	PBFlatTypes.PB_Media t = new PBFlatTypes.PB_Media();
    t.setMediaId();
    t.setUserId();
    t.setPostId();
    t.setAlbumId();
    t.setMediaTypeEnum();
    t.setWidth();
    t.setHeight();
    t.setSize();
    t.setDuration();
    t.setMd5Hash();
    t.setColor();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_Media t = new PBFlatTypes.PB_Media();
	t.MediaId = ;
	t.UserId = ;
	t.PostId = ;
	t.AlbumId = ;
	t.MediaTypeEnum = ;
	t.Width = ;
	t.Height = ;
	t.Size = ;
	t.Duration = ;
	t.Md5Hash = ;
	t.Color = ;
	t.CreatedTime = ;
	*/

	/*
	PB_Media t = new PB_Media();
	t.MediaId = m.getMediaId() ;
	t.UserId = m.getUserId() ;
	t.PostId = m.getPostId() ;
	t.AlbumId = m.getAlbumId() ;
	t.MediaTypeEnum = m.getMediaTypeEnum() ;
	t.Width = m.getWidth() ;
	t.Height = m.getHeight() ;
	t.Size = m.getSize() ;
	t.Duration = m.getDuration() ;
	t.Md5Hash = m.getMd5Hash() ;
	t.Color = m.getColor() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_Notify {
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
	   public int Seq;
	}
	/*
	folding
	PBFlatTypes.PB_Notify t = new PBFlatTypes.PB_Notify();
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
    t.setSeq();
	*/

	/*
	PBFlatTypes.PB_Notify t = new PBFlatTypes.PB_Notify();
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
	t.Seq = ;
	*/

	/*
	PB_Notify t = new PB_Notify();
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
	t.Seq = m.getSeq() ;
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
	   public int Id;
	   public int UserId;
	   public long Phone;
	   public String PhoneDisplayName;
	   public String PhoneFamilyName;
	   public String PhoneNumber;
	   public String PhoneNormalizedNumber;
	   public int PhoneContactRowId;
	   public int DeviceUuidId;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_PhoneContact t = new PBFlatTypes.PB_PhoneContact();
    t.setId();
    t.setUserId();
    t.setPhone();
    t.setPhoneDisplayName();
    t.setPhoneFamilyName();
    t.setPhoneNumber();
    t.setPhoneNormalizedNumber();
    t.setPhoneContactRowId();
    t.setDeviceUuidId();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_PhoneContact t = new PBFlatTypes.PB_PhoneContact();
	t.Id = ;
	t.UserId = ;
	t.Phone = ;
	t.PhoneDisplayName = ;
	t.PhoneFamilyName = ;
	t.PhoneNumber = ;
	t.PhoneNormalizedNumber = ;
	t.PhoneContactRowId = ;
	t.DeviceUuidId = ;
	t.CreatedTime = ;
	*/

	/*
	PB_PhoneContact t = new PB_PhoneContact();
	t.Id = m.getId() ;
	t.UserId = m.getUserId() ;
	t.Phone = m.getPhone() ;
	t.PhoneDisplayName = m.getPhoneDisplayName() ;
	t.PhoneFamilyName = m.getPhoneFamilyName() ;
	t.PhoneNumber = m.getPhoneNumber() ;
	t.PhoneNormalizedNumber = m.getPhoneNormalizedNumber() ;
	t.PhoneContactRowId = m.getPhoneContactRowId() ;
	t.DeviceUuidId = m.getDeviceUuidId() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_Post {
	   public long PostId;
	   public int UserId;
	   public int PostTypeEnum;
	   public int MediaId;
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
	}
	/*
	folding
	PBFlatTypes.PB_Post t = new PBFlatTypes.PB_Post();
    t.setPostId();
    t.setUserId();
    t.setPostTypeEnum();
    t.setMediaId();
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
	*/

	/*
	PBFlatTypes.PB_Post t = new PBFlatTypes.PB_Post();
	t.PostId = ;
	t.UserId = ;
	t.PostTypeEnum = ;
	t.MediaId = ;
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
	*/

	/*
	PB_Post t = new PB_Post();
	t.PostId = m.getPostId() ;
	t.UserId = m.getUserId() ;
	t.PostTypeEnum = m.getPostTypeEnum() ;
	t.MediaId = m.getMediaId() ;
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
	*/

	public class PB_PostKey {
	   public int Id;
	   public String Key;
	}
	/*
	folding
	PBFlatTypes.PB_PostKey t = new PBFlatTypes.PB_PostKey();
    t.setId();
    t.setKey();
	*/

	/*
	PBFlatTypes.PB_PostKey t = new PBFlatTypes.PB_PostKey();
	t.Id = ;
	t.Key = ;
	*/

	/*
	PB_PostKey t = new PB_PostKey();
	t.Id = m.getId() ;
	t.Key = m.getKey() ;
	*/

	public class PB_SearchClicked {
	   public long Id;
	   public String Query;
	   public int ClickType;
	   public int TargetId;
	   public int UserId;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_SearchClicked t = new PBFlatTypes.PB_SearchClicked();
    t.setId();
    t.setQuery();
    t.setClickType();
    t.setTargetId();
    t.setUserId();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_SearchClicked t = new PBFlatTypes.PB_SearchClicked();
	t.Id = ;
	t.Query = ;
	t.ClickType = ;
	t.TargetId = ;
	t.UserId = ;
	t.CreatedTime = ;
	*/

	/*
	PB_SearchClicked t = new PB_SearchClicked();
	t.Id = m.getId() ;
	t.Query = m.getQuery() ;
	t.ClickType = m.getClickType() ;
	t.TargetId = m.getTargetId() ;
	t.UserId = m.getUserId() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_Session {
	   public long Id;
	   public int UserId;
	   public String SessionUuid;
	   public String ClientUuid;
	   public String DeviceUuid;
	   public int LastActivityTime;
	   public String LastIpAddress;
	   public String LastWifiMacAddress;
	   public String LastNetworkType;
	   public int LastNetworkTypeEnumId;
	   public int AppVersion;
	   public int UpdatedTime;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_Session t = new PBFlatTypes.PB_Session();
    t.setId();
    t.setUserId();
    t.setSessionUuid();
    t.setClientUuid();
    t.setDeviceUuid();
    t.setLastActivityTime();
    t.setLastIpAddress();
    t.setLastWifiMacAddress();
    t.setLastNetworkType();
    t.setLastNetworkTypeEnumId();
    t.setAppVersion();
    t.setUpdatedTime();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_Session t = new PBFlatTypes.PB_Session();
	t.Id = ;
	t.UserId = ;
	t.SessionUuid = ;
	t.ClientUuid = ;
	t.DeviceUuid = ;
	t.LastActivityTime = ;
	t.LastIpAddress = ;
	t.LastWifiMacAddress = ;
	t.LastNetworkType = ;
	t.LastNetworkTypeEnumId = ;
	t.AppVersion = ;
	t.UpdatedTime = ;
	t.CreatedTime = ;
	*/

	/*
	PB_Session t = new PB_Session();
	t.Id = m.getId() ;
	t.UserId = m.getUserId() ;
	t.SessionUuid = m.getSessionUuid() ;
	t.ClientUuid = m.getClientUuid() ;
	t.DeviceUuid = m.getDeviceUuid() ;
	t.LastActivityTime = m.getLastActivityTime() ;
	t.LastIpAddress = m.getLastIpAddress() ;
	t.LastWifiMacAddress = m.getLastWifiMacAddress() ;
	t.LastNetworkType = m.getLastNetworkType() ;
	t.LastNetworkTypeEnumId = m.getLastNetworkTypeEnumId() ;
	t.AppVersion = m.getAppVersion() ;
	t.UpdatedTime = m.getUpdatedTime() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_SettingClient {
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
	/*
	folding
	PBFlatTypes.PB_SettingClient t = new PBFlatTypes.PB_SettingClient();
    t.setUserId();
    t.setAutoDownloadWifiVoice();
    t.setAutoDownloadWifiImage();
    t.setAutoDownloadWifiVideo();
    t.setAutoDownloadWifiFile();
    t.setAutoDownloadWifiMusic();
    t.setAutoDownloadWifiGif();
    t.setAutoDownloadCellularVoice();
    t.setAutoDownloadCellularImage();
    t.setAutoDownloadCellularVideo();
    t.setAutoDownloadCellularFile();
    t.setAutoDownloadCellularMusic();
    t.setAutoDownloadCellularGif();
    t.setAutoDownloadRoamingVoice();
    t.setAutoDownloadRoamingImage();
    t.setAutoDownloadRoamingVideo();
    t.setAutoDownloadRoamingFile();
    t.setAutoDownloadRoamingMusic();
    t.setAutoDownloadRoamingGif();
    t.setSaveToGallery();
	*/

	/*
	PBFlatTypes.PB_SettingClient t = new PBFlatTypes.PB_SettingClient();
	t.UserId = ;
	t.AutoDownloadWifiVoice = ;
	t.AutoDownloadWifiImage = ;
	t.AutoDownloadWifiVideo = ;
	t.AutoDownloadWifiFile = ;
	t.AutoDownloadWifiMusic = ;
	t.AutoDownloadWifiGif = ;
	t.AutoDownloadCellularVoice = ;
	t.AutoDownloadCellularImage = ;
	t.AutoDownloadCellularVideo = ;
	t.AutoDownloadCellularFile = ;
	t.AutoDownloadCellularMusic = ;
	t.AutoDownloadCellularGif = ;
	t.AutoDownloadRoamingVoice = ;
	t.AutoDownloadRoamingImage = ;
	t.AutoDownloadRoamingVideo = ;
	t.AutoDownloadRoamingFile = ;
	t.AutoDownloadRoamingMusic = ;
	t.AutoDownloadRoamingGif = ;
	t.SaveToGallery = ;
	*/

	/*
	PB_SettingClient t = new PB_SettingClient();
	t.UserId = m.getUserId() ;
	t.AutoDownloadWifiVoice = m.getAutoDownloadWifiVoice() ;
	t.AutoDownloadWifiImage = m.getAutoDownloadWifiImage() ;
	t.AutoDownloadWifiVideo = m.getAutoDownloadWifiVideo() ;
	t.AutoDownloadWifiFile = m.getAutoDownloadWifiFile() ;
	t.AutoDownloadWifiMusic = m.getAutoDownloadWifiMusic() ;
	t.AutoDownloadWifiGif = m.getAutoDownloadWifiGif() ;
	t.AutoDownloadCellularVoice = m.getAutoDownloadCellularVoice() ;
	t.AutoDownloadCellularImage = m.getAutoDownloadCellularImage() ;
	t.AutoDownloadCellularVideo = m.getAutoDownloadCellularVideo() ;
	t.AutoDownloadCellularFile = m.getAutoDownloadCellularFile() ;
	t.AutoDownloadCellularMusic = m.getAutoDownloadCellularMusic() ;
	t.AutoDownloadCellularGif = m.getAutoDownloadCellularGif() ;
	t.AutoDownloadRoamingVoice = m.getAutoDownloadRoamingVoice() ;
	t.AutoDownloadRoamingImage = m.getAutoDownloadRoamingImage() ;
	t.AutoDownloadRoamingVideo = m.getAutoDownloadRoamingVideo() ;
	t.AutoDownloadRoamingFile = m.getAutoDownloadRoamingFile() ;
	t.AutoDownloadRoamingMusic = m.getAutoDownloadRoamingMusic() ;
	t.AutoDownloadRoamingGif = m.getAutoDownloadRoamingGif() ;
	t.SaveToGallery = m.getSaveToGallery() ;
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

	public class PB_TagsPost {
	   public long Id;
	   public int TagId;
	   public int PostId;
	   public int PostTypeEnum;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_TagsPost t = new PBFlatTypes.PB_TagsPost();
    t.setId();
    t.setTagId();
    t.setPostId();
    t.setPostTypeEnum();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_TagsPost t = new PBFlatTypes.PB_TagsPost();
	t.Id = ;
	t.TagId = ;
	t.PostId = ;
	t.PostTypeEnum = ;
	t.CreatedTime = ;
	*/

	/*
	PB_TagsPost t = new PB_TagsPost();
	t.Id = m.getId() ;
	t.TagId = m.getTagId() ;
	t.PostId = m.getPostId() ;
	t.PostTypeEnum = m.getPostTypeEnum() ;
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
	   public int UserTypeEnum;
	   public int UserLevelEnum;
	   public long AvatarId;
	   public int ProfilePrivacyEnum;
	   public long Phone;
	   public String About;
	   public String Email;
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
	   public int CreatedSe;
	   public long UpdatedMs;
	   public int OnlinePrivacyEnum;
	   public int LastActivityTime;
	   public String Phone2;
	}
	/*
	folding
	PBFlatTypes.PB_User t = new PBFlatTypes.PB_User();
    t.setUserId();
    t.setUserName();
    t.setUserNameLower();
    t.setFirstName();
    t.setLastName();
    t.setUserTypeEnum();
    t.setUserLevelEnum();
    t.setAvatarId();
    t.setProfilePrivacyEnum();
    t.setPhone();
    t.setAbout();
    t.setEmail();
    t.setPasswordHash();
    t.setPasswordSalt();
    t.setFollowersCount();
    t.setFollowingCount();
    t.setPostsCount();
    t.setMediaCount();
    t.setLikesCount();
    t.setResharedCount();
    t.setLastActionTime();
    t.setLastPostTime();
    t.setPrimaryFollowingList();
    t.setCreatedSe();
    t.setUpdatedMs();
    t.setOnlinePrivacyEnum();
    t.setLastActivityTime();
    t.setPhone2();
	*/

	/*
	PBFlatTypes.PB_User t = new PBFlatTypes.PB_User();
	t.UserId = ;
	t.UserName = ;
	t.UserNameLower = ;
	t.FirstName = ;
	t.LastName = ;
	t.UserTypeEnum = ;
	t.UserLevelEnum = ;
	t.AvatarId = ;
	t.ProfilePrivacyEnum = ;
	t.Phone = ;
	t.About = ;
	t.Email = ;
	t.PasswordHash = ;
	t.PasswordSalt = ;
	t.FollowersCount = ;
	t.FollowingCount = ;
	t.PostsCount = ;
	t.MediaCount = ;
	t.LikesCount = ;
	t.ResharedCount = ;
	t.LastActionTime = ;
	t.LastPostTime = ;
	t.PrimaryFollowingList = ;
	t.CreatedSe = ;
	t.UpdatedMs = ;
	t.OnlinePrivacyEnum = ;
	t.LastActivityTime = ;
	t.Phone2 = ;
	*/

	/*
	PB_User t = new PB_User();
	t.UserId = m.getUserId() ;
	t.UserName = m.getUserName() ;
	t.UserNameLower = m.getUserNameLower() ;
	t.FirstName = m.getFirstName() ;
	t.LastName = m.getLastName() ;
	t.UserTypeEnum = m.getUserTypeEnum() ;
	t.UserLevelEnum = m.getUserLevelEnum() ;
	t.AvatarId = m.getAvatarId() ;
	t.ProfilePrivacyEnum = m.getProfilePrivacyEnum() ;
	t.Phone = m.getPhone() ;
	t.About = m.getAbout() ;
	t.Email = m.getEmail() ;
	t.PasswordHash = m.getPasswordHash() ;
	t.PasswordSalt = m.getPasswordSalt() ;
	t.FollowersCount = m.getFollowersCount() ;
	t.FollowingCount = m.getFollowingCount() ;
	t.PostsCount = m.getPostsCount() ;
	t.MediaCount = m.getMediaCount() ;
	t.LikesCount = m.getLikesCount() ;
	t.ResharedCount = m.getResharedCount() ;
	t.LastActionTime = m.getLastActionTime() ;
	t.LastPostTime = m.getLastPostTime() ;
	t.PrimaryFollowingList = m.getPrimaryFollowingList() ;
	t.CreatedSe = m.getCreatedSe() ;
	t.UpdatedMs = m.getUpdatedMs() ;
	t.OnlinePrivacyEnum = m.getOnlinePrivacyEnum() ;
	t.LastActivityTime = m.getLastActivityTime() ;
	t.Phone2 = m.getPhone2() ;
	*/

	public class PB_UserMetaInfo {
	   public int Id;
	   public int UserId;
	   public int IsNotificationDirty;
	   public int LastUserRecGen;
	}
	/*
	folding
	PBFlatTypes.PB_UserMetaInfo t = new PBFlatTypes.PB_UserMetaInfo();
    t.setId();
    t.setUserId();
    t.setIsNotificationDirty();
    t.setLastUserRecGen();
	*/

	/*
	PBFlatTypes.PB_UserMetaInfo t = new PBFlatTypes.PB_UserMetaInfo();
	t.Id = ;
	t.UserId = ;
	t.IsNotificationDirty = ;
	t.LastUserRecGen = ;
	*/

	/*
	PB_UserMetaInfo t = new PB_UserMetaInfo();
	t.Id = m.getId() ;
	t.UserId = m.getUserId() ;
	t.IsNotificationDirty = m.getIsNotificationDirty() ;
	t.LastUserRecGen = m.getLastUserRecGen() ;
	*/

	public class PB_UserPassword {
	   public int UserId;
	   public String Password;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_UserPassword t = new PBFlatTypes.PB_UserPassword();
    t.setUserId();
    t.setPassword();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_UserPassword t = new PBFlatTypes.PB_UserPassword();
	t.UserId = ;
	t.Password = ;
	t.CreatedTime = ;
	*/

	/*
	PB_UserPassword t = new PB_UserPassword();
	t.UserId = m.getUserId() ;
	t.Password = m.getPassword() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_Chat {
	   public String ChatKey;
	   public String RoomKey;
	   public int RoomTypeEnum;
	   public int UserId;
	   public int PeerUserId;
	   public long GroupId;
	   public int CreatedTime;
	   public int Seq;
	   public int SeenSeq;
	   public long UpdatedMs;
	}
	/*
	folding
	PBFlatTypes.PB_Chat t = new PBFlatTypes.PB_Chat();
    t.setChatKey();
    t.setRoomKey();
    t.setRoomTypeEnum();
    t.setUserId();
    t.setPeerUserId();
    t.setGroupId();
    t.setCreatedTime();
    t.setSeq();
    t.setSeenSeq();
    t.setUpdatedMs();
	*/

	/*
	PBFlatTypes.PB_Chat t = new PBFlatTypes.PB_Chat();
	t.ChatKey = ;
	t.RoomKey = ;
	t.RoomTypeEnum = ;
	t.UserId = ;
	t.PeerUserId = ;
	t.GroupId = ;
	t.CreatedTime = ;
	t.Seq = ;
	t.SeenSeq = ;
	t.UpdatedMs = ;
	*/

	/*
	PB_Chat t = new PB_Chat();
	t.ChatKey = m.getChatKey() ;
	t.RoomKey = m.getRoomKey() ;
	t.RoomTypeEnum = m.getRoomTypeEnum() ;
	t.UserId = m.getUserId() ;
	t.PeerUserId = m.getPeerUserId() ;
	t.GroupId = m.getGroupId() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.Seq = m.getSeq() ;
	t.SeenSeq = m.getSeenSeq() ;
	t.UpdatedMs = m.getUpdatedMs() ;
	*/

	public class PB_ChatLastMessage {
	   public String ChatKey;
	   public byte[] LastMsgPb;
	   public String LastMsgJson;
	}
	/*
	folding
	PBFlatTypes.PB_ChatLastMessage t = new PBFlatTypes.PB_ChatLastMessage();
    t.setChatKey();
    t.setLastMsgPb();
    t.setLastMsgJson();
	*/

	/*
	PBFlatTypes.PB_ChatLastMessage t = new PBFlatTypes.PB_ChatLastMessage();
	t.ChatKey = ;
	t.LastMsgPb = ;
	t.LastMsgJson = ;
	*/

	/*
	PB_ChatLastMessage t = new PB_ChatLastMessage();
	t.ChatKey = m.getChatKey() ;
	t.LastMsgPb = m.getLastMsgPb() ;
	t.LastMsgJson = m.getLastMsgJson() ;
	*/

	public class PB_ChatSync {
	   public long SyncId;
	   public int ToUserId;
	   public int ChatSyncTypeId;
	   public String ChatKey;
	   public long MessageId;
	   public byte[] MessagePb;
	   public String MessageJson;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_ChatSync t = new PBFlatTypes.PB_ChatSync();
    t.setSyncId();
    t.setToUserId();
    t.setChatSyncTypeId();
    t.setChatKey();
    t.setMessageId();
    t.setMessagePb();
    t.setMessageJson();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_ChatSync t = new PBFlatTypes.PB_ChatSync();
	t.SyncId = ;
	t.ToUserId = ;
	t.ChatSyncTypeId = ;
	t.ChatKey = ;
	t.MessageId = ;
	t.MessagePb = ;
	t.MessageJson = ;
	t.CreatedTime = ;
	*/

	/*
	PB_ChatSync t = new PB_ChatSync();
	t.SyncId = m.getSyncId() ;
	t.ToUserId = m.getToUserId() ;
	t.ChatSyncTypeId = m.getChatSyncTypeId() ;
	t.ChatKey = m.getChatKey() ;
	t.MessageId = m.getMessageId() ;
	t.MessagePb = m.getMessagePb() ;
	t.MessageJson = m.getMessageJson() ;
	t.CreatedTime = m.getCreatedTime() ;
	*/

	public class PB_DirectMessage {
	   public String ChatKey;
	   public long MessageId;
	   public String RoomKey;
	   public int UserId;
	   public long MessageFileId;
	   public int MessageTypeEnum;
	   public String Text;
	   public int CreatedTime;
	   public int Seq;
	   public int DeliviryStatusEnum;
	   public byte[] ExtraPB;
	}
	/*
	folding
	PBFlatTypes.PB_DirectMessage t = new PBFlatTypes.PB_DirectMessage();
    t.setChatKey();
    t.setMessageId();
    t.setRoomKey();
    t.setUserId();
    t.setMessageFileId();
    t.setMessageTypeEnum();
    t.setText();
    t.setCreatedTime();
    t.setSeq();
    t.setDeliviryStatusEnum();
    t.setExtraPB();
	*/

	/*
	PBFlatTypes.PB_DirectMessage t = new PBFlatTypes.PB_DirectMessage();
	t.ChatKey = ;
	t.MessageId = ;
	t.RoomKey = ;
	t.UserId = ;
	t.MessageFileId = ;
	t.MessageTypeEnum = ;
	t.Text = ;
	t.CreatedTime = ;
	t.Seq = ;
	t.DeliviryStatusEnum = ;
	t.ExtraPB = ;
	*/

	/*
	PB_DirectMessage t = new PB_DirectMessage();
	t.ChatKey = m.getChatKey() ;
	t.MessageId = m.getMessageId() ;
	t.RoomKey = m.getRoomKey() ;
	t.UserId = m.getUserId() ;
	t.MessageFileId = m.getMessageFileId() ;
	t.MessageTypeEnum = m.getMessageTypeEnum() ;
	t.Text = m.getText() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.Seq = m.getSeq() ;
	t.DeliviryStatusEnum = m.getDeliviryStatusEnum() ;
	t.ExtraPB = m.getExtraPB() ;
	*/

	public class PB_MessageFile {
	   public long MessageFileId;
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
	/*
	folding
	PBFlatTypes.PB_MessageFile t = new PBFlatTypes.PB_MessageFile();
    t.setMessageFileId();
    t.setFileTypeEnum();
    t.setUserId();
    t.setTitle();
    t.setSize();
    t.setWidth();
    t.setHeight();
    t.setDuration();
    t.setExtension();
    t.setMd5Hash();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_MessageFile t = new PBFlatTypes.PB_MessageFile();
	t.MessageFileId = ;
	t.FileTypeEnum = ;
	t.UserId = ;
	t.Title = ;
	t.Size = ;
	t.Width = ;
	t.Height = ;
	t.Duration = ;
	t.Extension = ;
	t.Md5Hash = ;
	t.CreatedTime = ;
	*/

	/*
	PB_MessageFile t = new PB_MessageFile();
	t.MessageFileId = m.getMessageFileId() ;
	t.FileTypeEnum = m.getFileTypeEnum() ;
	t.UserId = m.getUserId() ;
	t.Title = m.getTitle() ;
	t.Size = m.getSize() ;
	t.Width = m.getWidth() ;
	t.Height = m.getHeight() ;
	t.Duration = m.getDuration() ;
	t.Extension = m.getExtension() ;
	t.Md5Hash = m.getMd5Hash() ;
	t.CreatedTime = m.getCreatedTime() ;
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

	public class PB_MediaView {
	   public long MediaId;
	   public int UserId;
	   public int PostId;
	   public int AlbumId;
	   public int MediaTypeEnum;
	   public int Width;
	   public int Height;
	   public int Size;
	   public int Duration;
	   public String Color;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_MediaView t = new PBFlatTypes.PB_MediaView();
    t.setMediaId();
    t.setUserId();
    t.setPostId();
    t.setAlbumId();
    t.setMediaTypeEnum();
    t.setWidth();
    t.setHeight();
    t.setSize();
    t.setDuration();
    t.setColor();
    t.setCreatedTime();
	*/

	/*
	PBFlatTypes.PB_MediaView t = new PBFlatTypes.PB_MediaView();
	t.MediaId = ;
	t.UserId = ;
	t.PostId = ;
	t.AlbumId = ;
	t.MediaTypeEnum = ;
	t.Width = ;
	t.Height = ;
	t.Size = ;
	t.Duration = ;
	t.Color = ;
	t.CreatedTime = ;
	*/

	/*
	PB_MediaView t = new PB_MediaView();
	t.MediaId = m.getMediaId() ;
	t.UserId = m.getUserId() ;
	t.PostId = m.getPostId() ;
	t.AlbumId = m.getAlbumId() ;
	t.MediaTypeEnum = m.getMediaTypeEnum() ;
	t.Width = m.getWidth() ;
	t.Height = m.getHeight() ;
	t.Size = m.getSize() ;
	t.Duration = m.getDuration() ;
	t.Color = m.getColor() ;
	t.CreatedTime = m.getCreatedTime() ;
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

	public class PB_UserView {
	   public int UserId;
	   public String UserName;
	   public String FirstName;
	   public String LastName;
	   public UserTypeEnum UserTypeEnum;
	   public UserLevelEnum UserLevelEnum;
	   public long AvatarId;
	   public int ProfilePrivacyEnum;
	   public long Phone;
	   public String About;
	   public int FollowersCount;
	   public int FollowingCount;
	   public int PostsCount;
	   public int MediaCount;
	   public int LikesCount;
	   public int ResharedCount;
	   public UserOnlineStatusEnum UserOnlineStatusEnum;
	   public int LastActiveTime;
	   public FollowingEnum MyFollwing;
	}
	/*
	folding
	PBFlatTypes.PB_UserView t = new PBFlatTypes.PB_UserView();
    t.setUserId();
    t.setUserName();
    t.setFirstName();
    t.setLastName();
    t.setUserTypeEnum();
    t.setUserLevelEnum();
    t.setAvatarId();
    t.setProfilePrivacyEnum();
    t.setPhone();
    t.setAbout();
    t.setFollowersCount();
    t.setFollowingCount();
    t.setPostsCount();
    t.setMediaCount();
    t.setLikesCount();
    t.setResharedCount();
    t.setUserOnlineStatusEnum();
    t.setLastActiveTime();
    t.setMyFollwing();
	*/

	/*
	PBFlatTypes.PB_UserView t = new PBFlatTypes.PB_UserView();
	t.UserId = ;
	t.UserName = ;
	t.FirstName = ;
	t.LastName = ;
	t.UserTypeEnum = ;
	t.UserLevelEnum = ;
	t.AvatarId = ;
	t.ProfilePrivacyEnum = ;
	t.Phone = ;
	t.About = ;
	t.FollowersCount = ;
	t.FollowingCount = ;
	t.PostsCount = ;
	t.MediaCount = ;
	t.LikesCount = ;
	t.ResharedCount = ;
	t.UserOnlineStatusEnum = ;
	t.LastActiveTime = ;
	t.MyFollwing = ;
	*/

	/*
	PB_UserView t = new PB_UserView();
	t.UserId = m.getUserId() ;
	t.UserName = m.getUserName() ;
	t.FirstName = m.getFirstName() ;
	t.LastName = m.getLastName() ;
	t.UserTypeEnum = m.getUserTypeEnum() ;
	t.UserLevelEnum = m.getUserLevelEnum() ;
	t.AvatarId = m.getAvatarId() ;
	t.ProfilePrivacyEnum = m.getProfilePrivacyEnum() ;
	t.Phone = m.getPhone() ;
	t.About = m.getAbout() ;
	t.FollowersCount = m.getFollowersCount() ;
	t.FollowingCount = m.getFollowingCount() ;
	t.PostsCount = m.getPostsCount() ;
	t.MediaCount = m.getMediaCount() ;
	t.LikesCount = m.getLikesCount() ;
	t.ResharedCount = m.getResharedCount() ;
	t.UserOnlineStatusEnum = m.getUserOnlineStatusEnum() ;
	t.LastActiveTime = m.getLastActiveTime() ;
	t.MyFollwing = m.getMyFollwing() ;
	*/

	public class PB_TopProfileView {
	   public PB_UserView User;
	}
	/*
	folding
	PBFlatTypes.PB_TopProfileView t = new PBFlatTypes.PB_TopProfileView();
    t.setUser();
	*/

	/*
	PBFlatTypes.PB_TopProfileView t = new PBFlatTypes.PB_TopProfileView();
	t.User = ;
	*/

	/*
	PB_TopProfileView t = new PB_TopProfileView();
	t.User = m.getUser() ;
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

	public class PB_ChatView {
	   public String ChatKey;
	   public String RoomKey;
	   public int RoomTypeEnum;
	   public int UserId;
	   public int PeerUserId;
	   public long GroupId;
	   public int CreatedTime;
	   public int Seq;
	   public int SeenSeq;
	   public long UpdatedMs;
	   public PB_UserView UserView;
	   public PB_MessageView FirstUnreadMessage;
	   public PB_MessageView LastMessage;
	}
	/*
	folding
	PBFlatTypes.PB_ChatView t = new PBFlatTypes.PB_ChatView();
    t.setChatKey();
    t.setRoomKey();
    t.setRoomTypeEnum();
    t.setUserId();
    t.setPeerUserId();
    t.setGroupId();
    t.setCreatedTime();
    t.setSeq();
    t.setSeenSeq();
    t.setUpdatedMs();
    t.setUserView();
    t.setFirstUnreadMessage();
    t.setLastMessage();
	*/

	/*
	PBFlatTypes.PB_ChatView t = new PBFlatTypes.PB_ChatView();
	t.ChatKey = ;
	t.RoomKey = ;
	t.RoomTypeEnum = ;
	t.UserId = ;
	t.PeerUserId = ;
	t.GroupId = ;
	t.CreatedTime = ;
	t.Seq = ;
	t.SeenSeq = ;
	t.UpdatedMs = ;
	t.UserView = ;
	t.FirstUnreadMessage = ;
	t.LastMessage = ;
	*/

	/*
	PB_ChatView t = new PB_ChatView();
	t.ChatKey = m.getChatKey() ;
	t.RoomKey = m.getRoomKey() ;
	t.RoomTypeEnum = m.getRoomTypeEnum() ;
	t.UserId = m.getUserId() ;
	t.PeerUserId = m.getPeerUserId() ;
	t.GroupId = m.getGroupId() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.Seq = m.getSeq() ;
	t.SeenSeq = m.getSeenSeq() ;
	t.UpdatedMs = m.getUpdatedMs() ;
	t.UserView = m.getUserView() ;
	t.FirstUnreadMessage = m.getFirstUnreadMessage() ;
	t.LastMessage = m.getLastMessage() ;
	*/

	public class PB_MessageView {
	   public String ChatKey;
	   public long MessageId;
	   public String RoomKey;
	   public int UserId;
	   public long MessageFileId;
	   public int MessageTypeEnum;
	   public String Text;
	   public int CreatedTime;
	   public int Seq;
	   public int DeliviryStatusEnum;
	   public PB_UserView UserView;
	   public PB_MessageFileView MessageFileView;
	}
	/*
	folding
	PBFlatTypes.PB_MessageView t = new PBFlatTypes.PB_MessageView();
    t.setChatKey();
    t.setMessageId();
    t.setRoomKey();
    t.setUserId();
    t.setMessageFileId();
    t.setMessageTypeEnum();
    t.setText();
    t.setCreatedTime();
    t.setSeq();
    t.setDeliviryStatusEnum();
    t.setUserView();
    t.setMessageFileView();
	*/

	/*
	PBFlatTypes.PB_MessageView t = new PBFlatTypes.PB_MessageView();
	t.ChatKey = ;
	t.MessageId = ;
	t.RoomKey = ;
	t.UserId = ;
	t.MessageFileId = ;
	t.MessageTypeEnum = ;
	t.Text = ;
	t.CreatedTime = ;
	t.Seq = ;
	t.DeliviryStatusEnum = ;
	t.UserView = ;
	t.MessageFileView = ;
	*/

	/*
	PB_MessageView t = new PB_MessageView();
	t.ChatKey = m.getChatKey() ;
	t.MessageId = m.getMessageId() ;
	t.RoomKey = m.getRoomKey() ;
	t.UserId = m.getUserId() ;
	t.MessageFileId = m.getMessageFileId() ;
	t.MessageTypeEnum = m.getMessageTypeEnum() ;
	t.Text = m.getText() ;
	t.CreatedTime = m.getCreatedTime() ;
	t.Seq = m.getSeq() ;
	t.DeliviryStatusEnum = m.getDeliviryStatusEnum() ;
	t.UserView = m.getUserView() ;
	t.MessageFileView = m.getMessageFileView() ;
	*/

	public class PB_MessageFileView {
	   public long MessageFileId;
	   public int FileTypeEnum;
	   public int Size;
	   public int Width;
	   public int Height;
	   public int Duration;
	   public String Extension;
	}
	/*
	folding
	PBFlatTypes.PB_MessageFileView t = new PBFlatTypes.PB_MessageFileView();
    t.setMessageFileId();
    t.setFileTypeEnum();
    t.setSize();
    t.setWidth();
    t.setHeight();
    t.setDuration();
    t.setExtension();
	*/

	/*
	PBFlatTypes.PB_MessageFileView t = new PBFlatTypes.PB_MessageFileView();
	t.MessageFileId = ;
	t.FileTypeEnum = ;
	t.Size = ;
	t.Width = ;
	t.Height = ;
	t.Duration = ;
	t.Extension = ;
	*/

	/*
	PB_MessageFileView t = new PB_MessageFileView();
	t.MessageFileId = m.getMessageFileId() ;
	t.FileTypeEnum = m.getFileTypeEnum() ;
	t.Size = m.getSize() ;
	t.Width = m.getWidth() ;
	t.Height = m.getHeight() ;
	t.Duration = m.getDuration() ;
	t.Extension = m.getExtension() ;
	*/

	public class PB_MessageTableExtra {
	   public PB_MessageFileView MessageFileView;
	}
	/*
	folding
	PBFlatTypes.PB_MessageTableExtra t = new PBFlatTypes.PB_MessageTableExtra();
    t.setMessageFileView();
	*/

	/*
	PBFlatTypes.PB_MessageTableExtra t = new PBFlatTypes.PB_MessageTableExtra();
	t.MessageFileView = ;
	*/

	/*
	PB_MessageTableExtra t = new PB_MessageTableExtra();
	t.MessageFileView = m.getMessageFileView() ;
	*/

	
}

/*

RPC_HANDLERS.RPC_Auth RPC_Auth_Handeler = null;
RPC_HANDLERS.RPC_Chat RPC_Chat_Handeler = null;
RPC_HANDLERS.RPC_Other RPC_Other_Handeler = null;
RPC_HANDLERS.RPC_Page RPC_Page_Handeler = null;
RPC_HANDLERS.RPC_Search RPC_Search_Handeler = null;
RPC_HANDLERS.RPC_Social RPC_Social_Handeler = null;
RPC_HANDLERS.RPC_User RPC_User_Handeler = null;
	
*/