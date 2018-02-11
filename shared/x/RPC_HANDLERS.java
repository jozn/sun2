package ir.ms.pb;

import java.util.HashMap;
import java.util.concurrent.ConcurrentHashMap;
import java.util.Map;

public class RPC_HANDLERS {

public interface RPC_Auth {
    void CheckPhone( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void SendCode( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void SendCodeToSms( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void SendCodeToTelgram( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void SingUp( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void SingIn( PB_UserResponse_CheckUserName2 pb, boolean handled);
    void LogOut( PB_UserResponse_CheckUserName2 pb, boolean handled);
}
public interface RPC_Chat {
    void AddNewMessage( PB_ChatResponse_AddNewMessage pb, boolean handled);
    void SetRoomActionDoing( PB_ChatResponse_SetRoomActionDoing pb, boolean handled);
    void SetMessagesAsReceived( PB_ChatResponse_SetMessagesAsReceived pb, boolean handled);
    void SetMessagesRangeAsSeen( PB_ChatResponse_SetChatMessagesRangeAsSeen pb, boolean handled);
    void DeleteChatHistory( PB_ChatResponse_DeleteChatHistory pb, boolean handled);
    void DeleteMessagesByIds( PB_ChatResponse_DeleteMessagesByIds pb, boolean handled);
    void EditMessage( PB_ChatResponse_EditMessage pb, boolean handled);
    void GetChatList( PB_ChatResponse_GetChatList pb, boolean handled);
    void GetChatHistoryToOlder( PB_ChatResponse_GetChatHistoryToOlder pb, boolean handled);
}
public interface RPC_Other {
    void Echo( PB_OtherResponse_Echo pb, boolean handled);
}
public interface RPC_Page {
    void GetCommentsPage( PB_PageResponse_GetCommentsPage pb, boolean handled);
    void GetHomePage( PB_PageResponse_GetHomePage pb, boolean handled);
    void GetProfilePage( PB_PageResponse_GetProfilePage pb, boolean handled);
    void GetLikesPage( PB_PageResponse_GetLikesPage pb, boolean handled);
    void GetFollowersPage( PB_PageResponse_GetFollowersPage pb, boolean handled);
    void GetFollowingsPage( PB_PageResponse_GetFollowingsPage pb, boolean handled);
    void GetNotifiesPage( PB_PageResponse_GetNotifiesPage pb, boolean handled);
    void GetUserActionsPage( PB_PageResponse_GetUserActionsPage pb, boolean handled);
    void GetSuggestedPostsPage( PB_PageResponse_GetSuggestedPostsPage pb, boolean handled);
    void GetSuggestedUsersPage( PB_PageResponse_GetSuggestedUsersPage pb, boolean handled);
    void GetSuggestedTagsPage( PB_PageResponse_GetSuggestedTagsPage pb, boolean handled);
    void GetLastPostsPage( PB_PageResponse_GetLastPostsPage pb, boolean handled);
    void GetTagPage( PB_PageResponse_GetTagPage pb, boolean handled);
    void SearchTagsPage( PB_PageResponse_SearchTagsPage pb, boolean handled);
    void SearchUsersPage( PB_PageResponse_SearchUsersPage pb, boolean handled);
}
public interface RPC_Search {
    void SearchTags( PB_SearchResponse_AddNewC pb, boolean handled);
}
public interface RPC_Social {
    void AddComment( PB_SocialResponse_AddComment pb, boolean handled);
    void DeleteComment( PB_SocialResponse_DeleteComment pb, boolean handled);
    void AddPost( PB_SocialResponse_AddPost pb, boolean handled);
    void EditPost( PB_SocialResponse_EditPost pb, boolean handled);
    void DeletePost( PB_SocialResponse_DeletePost pb, boolean handled);
    void ArchivePost( PB_SocialResponse_ArchivePost pb, boolean handled);
    void LikePost( PB_SocialResponse_LikePost pb, boolean handled);
    void UnLikePost( PB_SocialResponse_UnLikePost pb, boolean handled);
    void FollowUser( PB_SocialResponse_FollowUser pb, boolean handled);
    void UnFollowUser( PB_SocialResponse_UnFollowUser pb, boolean handled);
}
public interface RPC_User {
    void BlockUser( PB_UserResponse_BlockUser pb, boolean handled);
    void UnBlockUser( PB_UserResponse_UnBlockUser pb, boolean handled);
    void GetBlockedList( PB_UserResponse_BlockedList pb, boolean handled);
    void UpdateAbout( PB_UserResponse_UpdateAbout pb, boolean handled);
    void UpdateUserName( PB_UserResponse_UpdateUserName pb, boolean handled);
    void ChangePrivacy( PB_UserResponseOffline_ChangePrivacy pb, boolean handled);
    void ChangeAvatar( PB_UserResponse_ChangeAvatar pb, boolean handled);
    void CheckUserName( PB_UserResponse_CheckUserName pb, boolean handled);
}


  public static class RPC_Auth_Empty implements RPC_Auth{
  
  	@Override
    public void CheckPhone( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.CheckPhone' ");
    }
  	@Override
    public void SendCode( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SendCode' ");
    }
  	@Override
    public void SendCodeToSms( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SendCodeToSms' ");
    }
  	@Override
    public void SendCodeToTelgram( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SendCodeToTelgram' ");
    }
  	@Override
    public void SingUp( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SingUp' ");
    }
  	@Override
    public void SingIn( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SingIn' ");
    }
  	@Override
    public void LogOut( PB_UserResponse_CheckUserName2 pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.LogOut' ");
    }
  }
  public static class RPC_Chat_Empty implements RPC_Chat{
  
  	@Override
    public void AddNewMessage( PB_ChatResponse_AddNewMessage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.AddNewMessage' ");
    }
  	@Override
    public void SetRoomActionDoing( PB_ChatResponse_SetRoomActionDoing pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.SetRoomActionDoing' ");
    }
  	@Override
    public void SetMessagesAsReceived( PB_ChatResponse_SetMessagesAsReceived pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.SetMessagesAsReceived' ");
    }
  	@Override
    public void SetMessagesRangeAsSeen( PB_ChatResponse_SetChatMessagesRangeAsSeen pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.SetMessagesRangeAsSeen' ");
    }
  	@Override
    public void DeleteChatHistory( PB_ChatResponse_DeleteChatHistory pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.DeleteChatHistory' ");
    }
  	@Override
    public void DeleteMessagesByIds( PB_ChatResponse_DeleteMessagesByIds pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.DeleteMessagesByIds' ");
    }
  	@Override
    public void EditMessage( PB_ChatResponse_EditMessage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.EditMessage' ");
    }
  	@Override
    public void GetChatList( PB_ChatResponse_GetChatList pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.GetChatList' ");
    }
  	@Override
    public void GetChatHistoryToOlder( PB_ChatResponse_GetChatHistoryToOlder pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.GetChatHistoryToOlder' ");
    }
  }
  public static class RPC_Other_Empty implements RPC_Other{
  
  	@Override
    public void Echo( PB_OtherResponse_Echo pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Other.Echo' ");
    }
  }
  public static class RPC_Page_Empty implements RPC_Page{
  
  	@Override
    public void GetCommentsPage( PB_PageResponse_GetCommentsPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetCommentsPage' ");
    }
  	@Override
    public void GetHomePage( PB_PageResponse_GetHomePage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetHomePage' ");
    }
  	@Override
    public void GetProfilePage( PB_PageResponse_GetProfilePage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetProfilePage' ");
    }
  	@Override
    public void GetLikesPage( PB_PageResponse_GetLikesPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetLikesPage' ");
    }
  	@Override
    public void GetFollowersPage( PB_PageResponse_GetFollowersPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetFollowersPage' ");
    }
  	@Override
    public void GetFollowingsPage( PB_PageResponse_GetFollowingsPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetFollowingsPage' ");
    }
  	@Override
    public void GetNotifiesPage( PB_PageResponse_GetNotifiesPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetNotifiesPage' ");
    }
  	@Override
    public void GetUserActionsPage( PB_PageResponse_GetUserActionsPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetUserActionsPage' ");
    }
  	@Override
    public void GetSuggestedPostsPage( PB_PageResponse_GetSuggestedPostsPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetSuggestedPostsPage' ");
    }
  	@Override
    public void GetSuggestedUsersPage( PB_PageResponse_GetSuggestedUsersPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetSuggestedUsersPage' ");
    }
  	@Override
    public void GetSuggestedTagsPage( PB_PageResponse_GetSuggestedTagsPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetSuggestedTagsPage' ");
    }
  	@Override
    public void GetLastPostsPage( PB_PageResponse_GetLastPostsPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetLastPostsPage' ");
    }
  	@Override
    public void GetTagPage( PB_PageResponse_GetTagPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetTagPage' ");
    }
  	@Override
    public void SearchTagsPage( PB_PageResponse_SearchTagsPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.SearchTagsPage' ");
    }
  	@Override
    public void SearchUsersPage( PB_PageResponse_SearchUsersPage pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.SearchUsersPage' ");
    }
  }
  public static class RPC_Search_Empty implements RPC_Search{
  
  	@Override
    public void SearchTags( PB_SearchResponse_AddNewC pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Search.SearchTags' ");
    }
  }
  public static class RPC_Social_Empty implements RPC_Social{
  
  	@Override
    public void AddComment( PB_SocialResponse_AddComment pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.AddComment' ");
    }
  	@Override
    public void DeleteComment( PB_SocialResponse_DeleteComment pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.DeleteComment' ");
    }
  	@Override
    public void AddPost( PB_SocialResponse_AddPost pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.AddPost' ");
    }
  	@Override
    public void EditPost( PB_SocialResponse_EditPost pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.EditPost' ");
    }
  	@Override
    public void DeletePost( PB_SocialResponse_DeletePost pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.DeletePost' ");
    }
  	@Override
    public void ArchivePost( PB_SocialResponse_ArchivePost pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.ArchivePost' ");
    }
  	@Override
    public void LikePost( PB_SocialResponse_LikePost pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.LikePost' ");
    }
  	@Override
    public void UnLikePost( PB_SocialResponse_UnLikePost pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.UnLikePost' ");
    }
  	@Override
    public void FollowUser( PB_SocialResponse_FollowUser pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.FollowUser' ");
    }
  	@Override
    public void UnFollowUser( PB_SocialResponse_UnFollowUser pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.UnFollowUser' ");
    }
  }
  public static class RPC_User_Empty implements RPC_User{
  
  	@Override
    public void BlockUser( PB_UserResponse_BlockUser pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.BlockUser' ");
    }
  	@Override
    public void UnBlockUser( PB_UserResponse_UnBlockUser pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.UnBlockUser' ");
    }
  	@Override
    public void GetBlockedList( PB_UserResponse_BlockedList pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.GetBlockedList' ");
    }
  	@Override
    public void UpdateAbout( PB_UserResponse_UpdateAbout pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.UpdateAbout' ");
    }
  	@Override
    public void UpdateUserName( PB_UserResponse_UpdateUserName pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.UpdateUserName' ");
    }
  	@Override
    public void ChangePrivacy( PB_UserResponseOffline_ChangePrivacy pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.ChangePrivacy' ");
    }
  	@Override
    public void ChangeAvatar( PB_UserResponse_ChangeAvatar pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.ChangeAvatar' ");
    }
  	@Override
    public void CheckUserName( PB_UserResponse_CheckUserName pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.CheckUserName' ");
    }
  }

	/////////////////////////////////// Maper of Handling methods /////////////////////////////////
	public static interface HandleRowRpcResponse {
		void handle(Object pb,boolean handled);
	}


	
	public static RPC_HANDLERS.RPC_Auth RPC_Auth_Default_Handler = new RPC_HANDLERS.RPC_Auth_Empty();
	public static RPC_HANDLERS.RPC_Chat RPC_Chat_Default_Handler = new RPC_HANDLERS.RPC_Chat_Empty();
	public static RPC_HANDLERS.RPC_Other RPC_Other_Default_Handler = new RPC_HANDLERS.RPC_Other_Empty();
	public static RPC_HANDLERS.RPC_Page RPC_Page_Default_Handler = new RPC_HANDLERS.RPC_Page_Empty();
	public static RPC_HANDLERS.RPC_Search RPC_Search_Default_Handler = new RPC_HANDLERS.RPC_Search_Empty();
	public static RPC_HANDLERS.RPC_Social RPC_Social_Default_Handler = new RPC_HANDLERS.RPC_Social_Empty();
	public static RPC_HANDLERS.RPC_User RPC_User_Default_Handler = new RPC_HANDLERS.RPC_User_Empty();


	public static Map<String,HandleRowRpcResponse > router = new HashMap<>();

	public static Map<String,HandleRowRpcResponse > getRouter(){
		if(router.size() ==0 ){
			initMap();
		}
		return router;
	}

	private synchronized static void initMap(){
	
	  
			router.put("RPC_Auth.CheckPhone", (pb, handled)->{
				if(pb instanceof PB_UserResponse_CheckUserName2){
					RPC_Auth_Default_Handler.CheckPhone((PB_UserResponse_CheckUserName2) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_CheckUserName2 in rpc: .CheckPhone -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Auth.SendCode", (pb, handled)->{
				if(pb instanceof PB_UserResponse_CheckUserName2){
					RPC_Auth_Default_Handler.SendCode((PB_UserResponse_CheckUserName2) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_CheckUserName2 in rpc: .SendCode -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Auth.SendCodeToSms", (pb, handled)->{
				if(pb instanceof PB_UserResponse_CheckUserName2){
					RPC_Auth_Default_Handler.SendCodeToSms((PB_UserResponse_CheckUserName2) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_CheckUserName2 in rpc: .SendCodeToSms -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Auth.SendCodeToTelgram", (pb, handled)->{
				if(pb instanceof PB_UserResponse_CheckUserName2){
					RPC_Auth_Default_Handler.SendCodeToTelgram((PB_UserResponse_CheckUserName2) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_CheckUserName2 in rpc: .SendCodeToTelgram -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Auth.SingUp", (pb, handled)->{
				if(pb instanceof PB_UserResponse_CheckUserName2){
					RPC_Auth_Default_Handler.SingUp((PB_UserResponse_CheckUserName2) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_CheckUserName2 in rpc: .SingUp -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Auth.SingIn", (pb, handled)->{
				if(pb instanceof PB_UserResponse_CheckUserName2){
					RPC_Auth_Default_Handler.SingIn((PB_UserResponse_CheckUserName2) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_CheckUserName2 in rpc: .SingIn -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Auth.LogOut", (pb, handled)->{
				if(pb instanceof PB_UserResponse_CheckUserName2){
					RPC_Auth_Default_Handler.LogOut((PB_UserResponse_CheckUserName2) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_CheckUserName2 in rpc: .LogOut -- class: " + pb );//.getClass().getName());
				}
			});
	  
	  
			router.put("RPC_Chat.AddNewMessage", (pb, handled)->{
				if(pb instanceof PB_ChatResponse_AddNewMessage){
					RPC_Chat_Default_Handler.AddNewMessage((PB_ChatResponse_AddNewMessage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_ChatResponse_AddNewMessage in rpc: .AddNewMessage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.SetRoomActionDoing", (pb, handled)->{
				if(pb instanceof PB_ChatResponse_SetRoomActionDoing){
					RPC_Chat_Default_Handler.SetRoomActionDoing((PB_ChatResponse_SetRoomActionDoing) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_ChatResponse_SetRoomActionDoing in rpc: .SetRoomActionDoing -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.SetMessagesAsReceived", (pb, handled)->{
				if(pb instanceof PB_ChatResponse_SetMessagesAsReceived){
					RPC_Chat_Default_Handler.SetMessagesAsReceived((PB_ChatResponse_SetMessagesAsReceived) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_ChatResponse_SetMessagesAsReceived in rpc: .SetMessagesAsReceived -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.SetMessagesRangeAsSeen", (pb, handled)->{
				if(pb instanceof PB_ChatResponse_SetChatMessagesRangeAsSeen){
					RPC_Chat_Default_Handler.SetMessagesRangeAsSeen((PB_ChatResponse_SetChatMessagesRangeAsSeen) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_ChatResponse_SetChatMessagesRangeAsSeen in rpc: .SetMessagesRangeAsSeen -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.DeleteChatHistory", (pb, handled)->{
				if(pb instanceof PB_ChatResponse_DeleteChatHistory){
					RPC_Chat_Default_Handler.DeleteChatHistory((PB_ChatResponse_DeleteChatHistory) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_ChatResponse_DeleteChatHistory in rpc: .DeleteChatHistory -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.DeleteMessagesByIds", (pb, handled)->{
				if(pb instanceof PB_ChatResponse_DeleteMessagesByIds){
					RPC_Chat_Default_Handler.DeleteMessagesByIds((PB_ChatResponse_DeleteMessagesByIds) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_ChatResponse_DeleteMessagesByIds in rpc: .DeleteMessagesByIds -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.EditMessage", (pb, handled)->{
				if(pb instanceof PB_ChatResponse_EditMessage){
					RPC_Chat_Default_Handler.EditMessage((PB_ChatResponse_EditMessage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_ChatResponse_EditMessage in rpc: .EditMessage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.GetChatList", (pb, handled)->{
				if(pb instanceof PB_ChatResponse_GetChatList){
					RPC_Chat_Default_Handler.GetChatList((PB_ChatResponse_GetChatList) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_ChatResponse_GetChatList in rpc: .GetChatList -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.GetChatHistoryToOlder", (pb, handled)->{
				if(pb instanceof PB_ChatResponse_GetChatHistoryToOlder){
					RPC_Chat_Default_Handler.GetChatHistoryToOlder((PB_ChatResponse_GetChatHistoryToOlder) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_ChatResponse_GetChatHistoryToOlder in rpc: .GetChatHistoryToOlder -- class: " + pb );//.getClass().getName());
				}
			});
	  
	  
			router.put("RPC_Other.Echo", (pb, handled)->{
				if(pb instanceof PB_OtherResponse_Echo){
					RPC_Other_Default_Handler.Echo((PB_OtherResponse_Echo) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_OtherResponse_Echo in rpc: .Echo -- class: " + pb );//.getClass().getName());
				}
			});
	  
	  
			router.put("RPC_Page.GetCommentsPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetCommentsPage){
					RPC_Page_Default_Handler.GetCommentsPage((PB_PageResponse_GetCommentsPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetCommentsPage in rpc: .GetCommentsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetHomePage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetHomePage){
					RPC_Page_Default_Handler.GetHomePage((PB_PageResponse_GetHomePage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetHomePage in rpc: .GetHomePage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetProfilePage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetProfilePage){
					RPC_Page_Default_Handler.GetProfilePage((PB_PageResponse_GetProfilePage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetProfilePage in rpc: .GetProfilePage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetLikesPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetLikesPage){
					RPC_Page_Default_Handler.GetLikesPage((PB_PageResponse_GetLikesPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetLikesPage in rpc: .GetLikesPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetFollowersPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetFollowersPage){
					RPC_Page_Default_Handler.GetFollowersPage((PB_PageResponse_GetFollowersPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetFollowersPage in rpc: .GetFollowersPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetFollowingsPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetFollowingsPage){
					RPC_Page_Default_Handler.GetFollowingsPage((PB_PageResponse_GetFollowingsPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetFollowingsPage in rpc: .GetFollowingsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetNotifiesPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetNotifiesPage){
					RPC_Page_Default_Handler.GetNotifiesPage((PB_PageResponse_GetNotifiesPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetNotifiesPage in rpc: .GetNotifiesPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetUserActionsPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetUserActionsPage){
					RPC_Page_Default_Handler.GetUserActionsPage((PB_PageResponse_GetUserActionsPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetUserActionsPage in rpc: .GetUserActionsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetSuggestedPostsPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetSuggestedPostsPage){
					RPC_Page_Default_Handler.GetSuggestedPostsPage((PB_PageResponse_GetSuggestedPostsPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetSuggestedPostsPage in rpc: .GetSuggestedPostsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetSuggestedUsersPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetSuggestedUsersPage){
					RPC_Page_Default_Handler.GetSuggestedUsersPage((PB_PageResponse_GetSuggestedUsersPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetSuggestedUsersPage in rpc: .GetSuggestedUsersPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetSuggestedTagsPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetSuggestedTagsPage){
					RPC_Page_Default_Handler.GetSuggestedTagsPage((PB_PageResponse_GetSuggestedTagsPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetSuggestedTagsPage in rpc: .GetSuggestedTagsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetLastPostsPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetLastPostsPage){
					RPC_Page_Default_Handler.GetLastPostsPage((PB_PageResponse_GetLastPostsPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetLastPostsPage in rpc: .GetLastPostsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetTagPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_GetTagPage){
					RPC_Page_Default_Handler.GetTagPage((PB_PageResponse_GetTagPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_GetTagPage in rpc: .GetTagPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.SearchTagsPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_SearchTagsPage){
					RPC_Page_Default_Handler.SearchTagsPage((PB_PageResponse_SearchTagsPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_SearchTagsPage in rpc: .SearchTagsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.SearchUsersPage", (pb, handled)->{
				if(pb instanceof PB_PageResponse_SearchUsersPage){
					RPC_Page_Default_Handler.SearchUsersPage((PB_PageResponse_SearchUsersPage) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_PageResponse_SearchUsersPage in rpc: .SearchUsersPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
	  
			router.put("RPC_Search.SearchTags", (pb, handled)->{
				if(pb instanceof PB_SearchResponse_AddNewC){
					RPC_Search_Default_Handler.SearchTags((PB_SearchResponse_AddNewC) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_SearchResponse_AddNewC in rpc: .SearchTags -- class: " + pb );//.getClass().getName());
				}
			});
	  
	  
			router.put("RPC_Social.AddComment", (pb, handled)->{
				if(pb instanceof PB_SocialResponse_AddComment){
					RPC_Social_Default_Handler.AddComment((PB_SocialResponse_AddComment) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_SocialResponse_AddComment in rpc: .AddComment -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.DeleteComment", (pb, handled)->{
				if(pb instanceof PB_SocialResponse_DeleteComment){
					RPC_Social_Default_Handler.DeleteComment((PB_SocialResponse_DeleteComment) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_SocialResponse_DeleteComment in rpc: .DeleteComment -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.AddPost", (pb, handled)->{
				if(pb instanceof PB_SocialResponse_AddPost){
					RPC_Social_Default_Handler.AddPost((PB_SocialResponse_AddPost) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_SocialResponse_AddPost in rpc: .AddPost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.EditPost", (pb, handled)->{
				if(pb instanceof PB_SocialResponse_EditPost){
					RPC_Social_Default_Handler.EditPost((PB_SocialResponse_EditPost) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_SocialResponse_EditPost in rpc: .EditPost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.DeletePost", (pb, handled)->{
				if(pb instanceof PB_SocialResponse_DeletePost){
					RPC_Social_Default_Handler.DeletePost((PB_SocialResponse_DeletePost) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_SocialResponse_DeletePost in rpc: .DeletePost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.ArchivePost", (pb, handled)->{
				if(pb instanceof PB_SocialResponse_ArchivePost){
					RPC_Social_Default_Handler.ArchivePost((PB_SocialResponse_ArchivePost) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_SocialResponse_ArchivePost in rpc: .ArchivePost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.LikePost", (pb, handled)->{
				if(pb instanceof PB_SocialResponse_LikePost){
					RPC_Social_Default_Handler.LikePost((PB_SocialResponse_LikePost) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_SocialResponse_LikePost in rpc: .LikePost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.UnLikePost", (pb, handled)->{
				if(pb instanceof PB_SocialResponse_UnLikePost){
					RPC_Social_Default_Handler.UnLikePost((PB_SocialResponse_UnLikePost) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_SocialResponse_UnLikePost in rpc: .UnLikePost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.FollowUser", (pb, handled)->{
				if(pb instanceof PB_SocialResponse_FollowUser){
					RPC_Social_Default_Handler.FollowUser((PB_SocialResponse_FollowUser) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_SocialResponse_FollowUser in rpc: .FollowUser -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.UnFollowUser", (pb, handled)->{
				if(pb instanceof PB_SocialResponse_UnFollowUser){
					RPC_Social_Default_Handler.UnFollowUser((PB_SocialResponse_UnFollowUser) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_SocialResponse_UnFollowUser in rpc: .UnFollowUser -- class: " + pb );//.getClass().getName());
				}
			});
	  
	  
			router.put("RPC_User.BlockUser", (pb, handled)->{
				if(pb instanceof PB_UserResponse_BlockUser){
					RPC_User_Default_Handler.BlockUser((PB_UserResponse_BlockUser) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_BlockUser in rpc: .BlockUser -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_User.UnBlockUser", (pb, handled)->{
				if(pb instanceof PB_UserResponse_UnBlockUser){
					RPC_User_Default_Handler.UnBlockUser((PB_UserResponse_UnBlockUser) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_UnBlockUser in rpc: .UnBlockUser -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_User.GetBlockedList", (pb, handled)->{
				if(pb instanceof PB_UserResponse_BlockedList){
					RPC_User_Default_Handler.GetBlockedList((PB_UserResponse_BlockedList) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_BlockedList in rpc: .GetBlockedList -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_User.UpdateAbout", (pb, handled)->{
				if(pb instanceof PB_UserResponse_UpdateAbout){
					RPC_User_Default_Handler.UpdateAbout((PB_UserResponse_UpdateAbout) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_UpdateAbout in rpc: .UpdateAbout -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_User.UpdateUserName", (pb, handled)->{
				if(pb instanceof PB_UserResponse_UpdateUserName){
					RPC_User_Default_Handler.UpdateUserName((PB_UserResponse_UpdateUserName) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_UpdateUserName in rpc: .UpdateUserName -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_User.ChangePrivacy", (pb, handled)->{
				if(pb instanceof PB_UserResponseOffline_ChangePrivacy){
					RPC_User_Default_Handler.ChangePrivacy((PB_UserResponseOffline_ChangePrivacy) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponseOffline_ChangePrivacy in rpc: .ChangePrivacy -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_User.ChangeAvatar", (pb, handled)->{
				if(pb instanceof PB_UserResponse_ChangeAvatar){
					RPC_User_Default_Handler.ChangeAvatar((PB_UserResponse_ChangeAvatar) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_ChangeAvatar in rpc: .ChangeAvatar -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_User.CheckUserName", (pb, handled)->{
				if(pb instanceof PB_UserResponse_CheckUserName){
					RPC_User_Default_Handler.CheckUserName((PB_UserResponse_CheckUserName) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_UserResponse_CheckUserName in rpc: .CheckUserName -- class: " + pb );//.getClass().getName());
				}
			});
	  
	}
	
}
/*

RPC_HANDLERS.RPC_Auth RPC_Auth_Default_Handler = new RPC_HANDLERS.RPC_Auth RPC_Auth_Empty();
RPC_HANDLERS.RPC_Chat RPC_Chat_Default_Handler = new RPC_HANDLERS.RPC_Chat RPC_Chat_Empty();
RPC_HANDLERS.RPC_Other RPC_Other_Default_Handler = new RPC_HANDLERS.RPC_Other RPC_Other_Empty();
RPC_HANDLERS.RPC_Page RPC_Page_Default_Handler = new RPC_HANDLERS.RPC_Page RPC_Page_Empty();
RPC_HANDLERS.RPC_Search RPC_Search_Default_Handler = new RPC_HANDLERS.RPC_Search RPC_Search_Empty();
RPC_HANDLERS.RPC_Social RPC_Social_Default_Handler = new RPC_HANDLERS.RPC_Social RPC_Social_Empty();
RPC_HANDLERS.RPC_User RPC_User_Default_Handler = new RPC_HANDLERS.RPC_User RPC_User_Empty();
	
*/