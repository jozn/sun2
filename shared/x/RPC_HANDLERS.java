package ir.ms.pb;

import java.util.HashMap;
import java.util.concurrent.ConcurrentHashMap;
import java.util.Map;

public class RPC_HANDLERS {

public interface RPC_Auth {
    void SendConfirmCode( RPC_Auth_Types.SendConfirmCode.Response pb, boolean handled);
    void ConfirmCode( RPC_Auth_Types.ConfirmCode.Response pb, boolean handled);
    void SingUp( RPC_Auth_Types.SingUp.Response pb, boolean handled);
    void SingIn( RPC_Auth_Types.SingIn.Response pb, boolean handled);
    void LogOut( RPC_Auth_Types.LogOut.Response pb, boolean handled);
}
public interface RPC_Chat {
    void AddNewMessage( PB_RPC_Chat_Types.AddNewMessage.Response pb, boolean handled);
    void SetRoomActionDoing( PB_RPC_Chat_Types.SetRoomActionDoing.Response pb, boolean handled);
    void GetChatList( PB_RPC_Chat_Types.GetChatList.Response pb, boolean handled);
    void GetChatHistory( PB_RPC_Chat_Types.GetChatHistory.Response pb, boolean handled);
    void AddRoomsChange( PB_RPC_Chat_Types.AddRoomsChange.Response pb, boolean handled);
    void GetRoomsChange( PB_RPC_Chat_Types.GetRoomsChange.Response pb, boolean handled);
}
public interface RPC_General {
    void Echo( RPC_General_Types.Echo.Response pb, boolean handled);
    void CheckUserName( RPC_General_Types.CheckUserName.Response pb, boolean handled);
}
public interface RPC_Page {
    void GetCommentsPage( RPC_Page_Types.GetCommentsPage.Response pb, boolean handled);
    void GetHomePage( RPC_Page_Types.GetHomePage.Response pb, boolean handled);
    void GetProfileAbout( RPC_Page_Types.GetProfileAbout.Response pb, boolean handled);
    void GetProfileAllShared( RPC_Page_Types.GetProfileAllShared.Response pb, boolean handled);
    void GetProfileByCategoryPage( RPC_Page_Types.GetProfileByCategoryPage.Response pb, boolean handled);
    void GetLikesPage( RPC_Page_Types.GetLikesPage.Response pb, boolean handled);
    void GetFollowersPage( RPC_Page_Types.GetFollowersPage.Response pb, boolean handled);
    void GetFollowingsPage( RPC_Page_Types.GetFollowingsPage.Response pb, boolean handled);
    void GetNotifiesPage( RPC_Page_Types.GetNotifiesPage.Response pb, boolean handled);
    void GetUserActionsPage( RPC_Page_Types.GetUserActionsPage.Response pb, boolean handled);
    void GetPromotedPostsPage( RPC_Page_Types.GetPromotedPostsPage.Response pb, boolean handled);
    void GetSuggestedUsersPage( RPC_Page_Types.GetSuggestedUsersPage.Response pb, boolean handled);
    void GetSuggestedTagsPage( RPC_Page_Types.GetSuggestedTagsPage.Response pb, boolean handled);
    void GetLastPostsPage( RPC_Page_Types.GetLastPostsPage.Response pb, boolean handled);
    void GetLastTagPage( RPC_Page_Types.GetLastTagPage.Response pb, boolean handled);
    void SearchTagsPage( RPC_Page_Types.SearchTagsPage.Response pb, boolean handled);
    void SearchUsersPage( RPC_Page_Types.SearchUsersPage.Response pb, boolean handled);
}
public interface RPC_Social {
    void AddComment( RPC_Social_Types.AddComment.Response pb, boolean handled);
    void DeleteComment( RPC_Social_Types.DeleteComment.Response pb, boolean handled);
    void EditComment( RPC_Social_Types.EditComment.Response pb, boolean handled);
    void LikeComment( RPC_Social_Types.LikeComment.Response pb, boolean handled);
    void AddPost( RPC_Social_Types.AddPost.Response pb, boolean handled);
    void EditPost( RPC_Social_Types.EditPost.Response pb, boolean handled);
    void DeletePost( RPC_Social_Types.DeletePost.Response pb, boolean handled);
    void ArchivePost( RPC_Social_Types.ArchivePost.Response pb, boolean handled);
    void PromotePost( RPC_Social_Types.PromotePost.Response pb, boolean handled);
    void LikePost( RPC_Social_Types.LikePost.Response pb, boolean handled);
    void UnLikePost( RPC_Social_Types.UnLikePost.Response pb, boolean handled);
    void FollowUser( RPC_Social_Types.FollowUser.Response pb, boolean handled);
    void UnFollowUser( RPC_Social_Types.UnFollowUser.Response pb, boolean handled);
    void BlockUser( RPC_Social_Types.BlockUser.Response pb, boolean handled);
    void UnBlockUser( RPC_Social_Types.UnBlockUser.Response pb, boolean handled);
    void AddSeenPosts( RPC_Social_Types.AddSeenPosts.Response pb, boolean handled);
}
public interface RPC_User {
    void UpdateAbout( RPC_User_Types.UpdateAbout.Response pb, boolean handled);
    void UpdateUserName( RPC_User_Types.UpdateUserName.Response pb, boolean handled);
    void ChangePrivacy( RPC_User_Types.ChangePrivacy.Response pb, boolean handled);
    void ChangeAvatar( RPC_User_Types.ChangeAvatar.Response pb, boolean handled);
}


  public static class RPC_Auth_Empty implements RPC_Auth{
  
  	@Override
    public void SendConfirmCode( RPC_Auth_Types.SendConfirmCode.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SendConfirmCode' ");
    }
  	@Override
    public void ConfirmCode( RPC_Auth_Types.ConfirmCode.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.ConfirmCode' ");
    }
  	@Override
    public void SingUp( RPC_Auth_Types.SingUp.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SingUp' ");
    }
  	@Override
    public void SingIn( RPC_Auth_Types.SingIn.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.SingIn' ");
    }
  	@Override
    public void LogOut( RPC_Auth_Types.LogOut.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Auth.LogOut' ");
    }
  }
  public static class RPC_Chat_Empty implements RPC_Chat{
  
  	@Override
    public void AddNewMessage( PB_RPC_Chat_Types.AddNewMessage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.AddNewMessage' ");
    }
  	@Override
    public void SetRoomActionDoing( PB_RPC_Chat_Types.SetRoomActionDoing.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.SetRoomActionDoing' ");
    }
  	@Override
    public void GetChatList( PB_RPC_Chat_Types.GetChatList.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.GetChatList' ");
    }
  	@Override
    public void GetChatHistory( PB_RPC_Chat_Types.GetChatHistory.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.GetChatHistory' ");
    }
  	@Override
    public void AddRoomsChange( PB_RPC_Chat_Types.AddRoomsChange.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.AddRoomsChange' ");
    }
  	@Override
    public void GetRoomsChange( PB_RPC_Chat_Types.GetRoomsChange.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.GetRoomsChange' ");
    }
  }
  public static class RPC_General_Empty implements RPC_General{
  
  	@Override
    public void Echo( RPC_General_Types.Echo.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_General.Echo' ");
    }
  	@Override
    public void CheckUserName( RPC_General_Types.CheckUserName.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_General.CheckUserName' ");
    }
  }
  public static class RPC_Page_Empty implements RPC_Page{
  
  	@Override
    public void GetCommentsPage( RPC_Page_Types.GetCommentsPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetCommentsPage' ");
    }
  	@Override
    public void GetHomePage( RPC_Page_Types.GetHomePage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetHomePage' ");
    }
  	@Override
    public void GetProfileAbout( RPC_Page_Types.GetProfileAbout.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetProfileAbout' ");
    }
  	@Override
    public void GetProfileAllShared( RPC_Page_Types.GetProfileAllShared.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetProfileAllShared' ");
    }
  	@Override
    public void GetProfileByCategoryPage( RPC_Page_Types.GetProfileByCategoryPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetProfileByCategoryPage' ");
    }
  	@Override
    public void GetLikesPage( RPC_Page_Types.GetLikesPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetLikesPage' ");
    }
  	@Override
    public void GetFollowersPage( RPC_Page_Types.GetFollowersPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetFollowersPage' ");
    }
  	@Override
    public void GetFollowingsPage( RPC_Page_Types.GetFollowingsPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetFollowingsPage' ");
    }
  	@Override
    public void GetNotifiesPage( RPC_Page_Types.GetNotifiesPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetNotifiesPage' ");
    }
  	@Override
    public void GetUserActionsPage( RPC_Page_Types.GetUserActionsPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetUserActionsPage' ");
    }
  	@Override
    public void GetPromotedPostsPage( RPC_Page_Types.GetPromotedPostsPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetPromotedPostsPage' ");
    }
  	@Override
    public void GetSuggestedUsersPage( RPC_Page_Types.GetSuggestedUsersPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetSuggestedUsersPage' ");
    }
  	@Override
    public void GetSuggestedTagsPage( RPC_Page_Types.GetSuggestedTagsPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetSuggestedTagsPage' ");
    }
  	@Override
    public void GetLastPostsPage( RPC_Page_Types.GetLastPostsPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetLastPostsPage' ");
    }
  	@Override
    public void GetLastTagPage( RPC_Page_Types.GetLastTagPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.GetLastTagPage' ");
    }
  	@Override
    public void SearchTagsPage( RPC_Page_Types.SearchTagsPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.SearchTagsPage' ");
    }
  	@Override
    public void SearchUsersPage( RPC_Page_Types.SearchUsersPage.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Page.SearchUsersPage' ");
    }
  }
  public static class RPC_Social_Empty implements RPC_Social{
  
  	@Override
    public void AddComment( RPC_Social_Types.AddComment.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.AddComment' ");
    }
  	@Override
    public void DeleteComment( RPC_Social_Types.DeleteComment.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.DeleteComment' ");
    }
  	@Override
    public void EditComment( RPC_Social_Types.EditComment.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.EditComment' ");
    }
  	@Override
    public void LikeComment( RPC_Social_Types.LikeComment.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.LikeComment' ");
    }
  	@Override
    public void AddPost( RPC_Social_Types.AddPost.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.AddPost' ");
    }
  	@Override
    public void EditPost( RPC_Social_Types.EditPost.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.EditPost' ");
    }
  	@Override
    public void DeletePost( RPC_Social_Types.DeletePost.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.DeletePost' ");
    }
  	@Override
    public void ArchivePost( RPC_Social_Types.ArchivePost.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.ArchivePost' ");
    }
  	@Override
    public void PromotePost( RPC_Social_Types.PromotePost.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.PromotePost' ");
    }
  	@Override
    public void LikePost( RPC_Social_Types.LikePost.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.LikePost' ");
    }
  	@Override
    public void UnLikePost( RPC_Social_Types.UnLikePost.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.UnLikePost' ");
    }
  	@Override
    public void FollowUser( RPC_Social_Types.FollowUser.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.FollowUser' ");
    }
  	@Override
    public void UnFollowUser( RPC_Social_Types.UnFollowUser.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.UnFollowUser' ");
    }
  	@Override
    public void BlockUser( RPC_Social_Types.BlockUser.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.BlockUser' ");
    }
  	@Override
    public void UnBlockUser( RPC_Social_Types.UnBlockUser.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.UnBlockUser' ");
    }
  	@Override
    public void AddSeenPosts( RPC_Social_Types.AddSeenPosts.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Social.AddSeenPosts' ");
    }
  }
  public static class RPC_User_Empty implements RPC_User{
  
  	@Override
    public void UpdateAbout( RPC_User_Types.UpdateAbout.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.UpdateAbout' ");
    }
  	@Override
    public void UpdateUserName( RPC_User_Types.UpdateUserName.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.UpdateUserName' ");
    }
  	@Override
    public void ChangePrivacy( RPC_User_Types.ChangePrivacy.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.ChangePrivacy' ");
    }
  	@Override
    public void ChangeAvatar( RPC_User_Types.ChangeAvatar.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_User.ChangeAvatar' ");
    }
  }

	/////////////////////////////////// Maper of Handling methods /////////////////////////////////
	public static interface HandleRowRpcResponse {
		void handle(Object pb,boolean handled);
	}


	
	public static RPC_HANDLERS.RPC_Auth RPC_Auth_Default_Handler = new RPC_HANDLERS.RPC_Auth_Empty();
	public static RPC_HANDLERS.RPC_Chat RPC_Chat_Default_Handler = new RPC_HANDLERS.RPC_Chat_Empty();
	public static RPC_HANDLERS.RPC_General RPC_General_Default_Handler = new RPC_HANDLERS.RPC_General_Empty();
	public static RPC_HANDLERS.RPC_Page RPC_Page_Default_Handler = new RPC_HANDLERS.RPC_Page_Empty();
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
	
	  
			router.put("RPC_Auth.SendConfirmCode", (pb, handled)->{
				if(pb instanceof RPC_Auth_Types.SendConfirmCode.Response){
					RPC_Auth_Default_Handler.SendConfirmCode((RPC_Auth_Types.SendConfirmCode.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Auth_Types.SendConfirmCode.Response in rpc: .SendConfirmCode -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Auth.ConfirmCode", (pb, handled)->{
				if(pb instanceof RPC_Auth_Types.ConfirmCode.Response){
					RPC_Auth_Default_Handler.ConfirmCode((RPC_Auth_Types.ConfirmCode.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Auth_Types.ConfirmCode.Response in rpc: .ConfirmCode -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Auth.SingUp", (pb, handled)->{
				if(pb instanceof RPC_Auth_Types.SingUp.Response){
					RPC_Auth_Default_Handler.SingUp((RPC_Auth_Types.SingUp.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Auth_Types.SingUp.Response in rpc: .SingUp -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Auth.SingIn", (pb, handled)->{
				if(pb instanceof RPC_Auth_Types.SingIn.Response){
					RPC_Auth_Default_Handler.SingIn((RPC_Auth_Types.SingIn.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Auth_Types.SingIn.Response in rpc: .SingIn -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Auth.LogOut", (pb, handled)->{
				if(pb instanceof RPC_Auth_Types.LogOut.Response){
					RPC_Auth_Default_Handler.LogOut((RPC_Auth_Types.LogOut.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Auth_Types.LogOut.Response in rpc: .LogOut -- class: " + pb );//.getClass().getName());
				}
			});
	  
	  
			router.put("RPC_Chat.AddNewMessage", (pb, handled)->{
				if(pb instanceof PB_RPC_Chat_Types.AddNewMessage.Response){
					RPC_Chat_Default_Handler.AddNewMessage((PB_RPC_Chat_Types.AddNewMessage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_RPC_Chat_Types.AddNewMessage.Response in rpc: .AddNewMessage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.SetRoomActionDoing", (pb, handled)->{
				if(pb instanceof PB_RPC_Chat_Types.SetRoomActionDoing.Response){
					RPC_Chat_Default_Handler.SetRoomActionDoing((PB_RPC_Chat_Types.SetRoomActionDoing.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_RPC_Chat_Types.SetRoomActionDoing.Response in rpc: .SetRoomActionDoing -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.GetChatList", (pb, handled)->{
				if(pb instanceof PB_RPC_Chat_Types.GetChatList.Response){
					RPC_Chat_Default_Handler.GetChatList((PB_RPC_Chat_Types.GetChatList.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_RPC_Chat_Types.GetChatList.Response in rpc: .GetChatList -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.GetChatHistory", (pb, handled)->{
				if(pb instanceof PB_RPC_Chat_Types.GetChatHistory.Response){
					RPC_Chat_Default_Handler.GetChatHistory((PB_RPC_Chat_Types.GetChatHistory.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_RPC_Chat_Types.GetChatHistory.Response in rpc: .GetChatHistory -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.AddRoomsChange", (pb, handled)->{
				if(pb instanceof PB_RPC_Chat_Types.AddRoomsChange.Response){
					RPC_Chat_Default_Handler.AddRoomsChange((PB_RPC_Chat_Types.AddRoomsChange.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_RPC_Chat_Types.AddRoomsChange.Response in rpc: .AddRoomsChange -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.GetRoomsChange", (pb, handled)->{
				if(pb instanceof PB_RPC_Chat_Types.GetRoomsChange.Response){
					RPC_Chat_Default_Handler.GetRoomsChange((PB_RPC_Chat_Types.GetRoomsChange.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_RPC_Chat_Types.GetRoomsChange.Response in rpc: .GetRoomsChange -- class: " + pb );//.getClass().getName());
				}
			});
	  
	  
			router.put("RPC_General.Echo", (pb, handled)->{
				if(pb instanceof RPC_General_Types.Echo.Response){
					RPC_General_Default_Handler.Echo((RPC_General_Types.Echo.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_General_Types.Echo.Response in rpc: .Echo -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_General.CheckUserName", (pb, handled)->{
				if(pb instanceof RPC_General_Types.CheckUserName.Response){
					RPC_General_Default_Handler.CheckUserName((RPC_General_Types.CheckUserName.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_General_Types.CheckUserName.Response in rpc: .CheckUserName -- class: " + pb );//.getClass().getName());
				}
			});
	  
	  
			router.put("RPC_Page.GetCommentsPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetCommentsPage.Response){
					RPC_Page_Default_Handler.GetCommentsPage((RPC_Page_Types.GetCommentsPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetCommentsPage.Response in rpc: .GetCommentsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetHomePage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetHomePage.Response){
					RPC_Page_Default_Handler.GetHomePage((RPC_Page_Types.GetHomePage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetHomePage.Response in rpc: .GetHomePage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetProfileAbout", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetProfileAbout.Response){
					RPC_Page_Default_Handler.GetProfileAbout((RPC_Page_Types.GetProfileAbout.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetProfileAbout.Response in rpc: .GetProfileAbout -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetProfileAllShared", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetProfileAllShared.Response){
					RPC_Page_Default_Handler.GetProfileAllShared((RPC_Page_Types.GetProfileAllShared.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetProfileAllShared.Response in rpc: .GetProfileAllShared -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetProfileByCategoryPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetProfileByCategoryPage.Response){
					RPC_Page_Default_Handler.GetProfileByCategoryPage((RPC_Page_Types.GetProfileByCategoryPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetProfileByCategoryPage.Response in rpc: .GetProfileByCategoryPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetLikesPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetLikesPage.Response){
					RPC_Page_Default_Handler.GetLikesPage((RPC_Page_Types.GetLikesPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetLikesPage.Response in rpc: .GetLikesPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetFollowersPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetFollowersPage.Response){
					RPC_Page_Default_Handler.GetFollowersPage((RPC_Page_Types.GetFollowersPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetFollowersPage.Response in rpc: .GetFollowersPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetFollowingsPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetFollowingsPage.Response){
					RPC_Page_Default_Handler.GetFollowingsPage((RPC_Page_Types.GetFollowingsPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetFollowingsPage.Response in rpc: .GetFollowingsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetNotifiesPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetNotifiesPage.Response){
					RPC_Page_Default_Handler.GetNotifiesPage((RPC_Page_Types.GetNotifiesPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetNotifiesPage.Response in rpc: .GetNotifiesPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetUserActionsPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetUserActionsPage.Response){
					RPC_Page_Default_Handler.GetUserActionsPage((RPC_Page_Types.GetUserActionsPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetUserActionsPage.Response in rpc: .GetUserActionsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetPromotedPostsPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetPromotedPostsPage.Response){
					RPC_Page_Default_Handler.GetPromotedPostsPage((RPC_Page_Types.GetPromotedPostsPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetPromotedPostsPage.Response in rpc: .GetPromotedPostsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetSuggestedUsersPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetSuggestedUsersPage.Response){
					RPC_Page_Default_Handler.GetSuggestedUsersPage((RPC_Page_Types.GetSuggestedUsersPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetSuggestedUsersPage.Response in rpc: .GetSuggestedUsersPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetSuggestedTagsPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetSuggestedTagsPage.Response){
					RPC_Page_Default_Handler.GetSuggestedTagsPage((RPC_Page_Types.GetSuggestedTagsPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetSuggestedTagsPage.Response in rpc: .GetSuggestedTagsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetLastPostsPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetLastPostsPage.Response){
					RPC_Page_Default_Handler.GetLastPostsPage((RPC_Page_Types.GetLastPostsPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetLastPostsPage.Response in rpc: .GetLastPostsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.GetLastTagPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.GetLastTagPage.Response){
					RPC_Page_Default_Handler.GetLastTagPage((RPC_Page_Types.GetLastTagPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.GetLastTagPage.Response in rpc: .GetLastTagPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.SearchTagsPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.SearchTagsPage.Response){
					RPC_Page_Default_Handler.SearchTagsPage((RPC_Page_Types.SearchTagsPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.SearchTagsPage.Response in rpc: .SearchTagsPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Page.SearchUsersPage", (pb, handled)->{
				if(pb instanceof RPC_Page_Types.SearchUsersPage.Response){
					RPC_Page_Default_Handler.SearchUsersPage((RPC_Page_Types.SearchUsersPage.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Page_Types.SearchUsersPage.Response in rpc: .SearchUsersPage -- class: " + pb );//.getClass().getName());
				}
			});
	  
	  
			router.put("RPC_Social.AddComment", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.AddComment.Response){
					RPC_Social_Default_Handler.AddComment((RPC_Social_Types.AddComment.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.AddComment.Response in rpc: .AddComment -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.DeleteComment", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.DeleteComment.Response){
					RPC_Social_Default_Handler.DeleteComment((RPC_Social_Types.DeleteComment.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.DeleteComment.Response in rpc: .DeleteComment -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.EditComment", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.EditComment.Response){
					RPC_Social_Default_Handler.EditComment((RPC_Social_Types.EditComment.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.EditComment.Response in rpc: .EditComment -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.LikeComment", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.LikeComment.Response){
					RPC_Social_Default_Handler.LikeComment((RPC_Social_Types.LikeComment.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.LikeComment.Response in rpc: .LikeComment -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.AddPost", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.AddPost.Response){
					RPC_Social_Default_Handler.AddPost((RPC_Social_Types.AddPost.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.AddPost.Response in rpc: .AddPost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.EditPost", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.EditPost.Response){
					RPC_Social_Default_Handler.EditPost((RPC_Social_Types.EditPost.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.EditPost.Response in rpc: .EditPost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.DeletePost", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.DeletePost.Response){
					RPC_Social_Default_Handler.DeletePost((RPC_Social_Types.DeletePost.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.DeletePost.Response in rpc: .DeletePost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.ArchivePost", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.ArchivePost.Response){
					RPC_Social_Default_Handler.ArchivePost((RPC_Social_Types.ArchivePost.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.ArchivePost.Response in rpc: .ArchivePost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.PromotePost", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.PromotePost.Response){
					RPC_Social_Default_Handler.PromotePost((RPC_Social_Types.PromotePost.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.PromotePost.Response in rpc: .PromotePost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.LikePost", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.LikePost.Response){
					RPC_Social_Default_Handler.LikePost((RPC_Social_Types.LikePost.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.LikePost.Response in rpc: .LikePost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.UnLikePost", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.UnLikePost.Response){
					RPC_Social_Default_Handler.UnLikePost((RPC_Social_Types.UnLikePost.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.UnLikePost.Response in rpc: .UnLikePost -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.FollowUser", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.FollowUser.Response){
					RPC_Social_Default_Handler.FollowUser((RPC_Social_Types.FollowUser.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.FollowUser.Response in rpc: .FollowUser -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.UnFollowUser", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.UnFollowUser.Response){
					RPC_Social_Default_Handler.UnFollowUser((RPC_Social_Types.UnFollowUser.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.UnFollowUser.Response in rpc: .UnFollowUser -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.BlockUser", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.BlockUser.Response){
					RPC_Social_Default_Handler.BlockUser((RPC_Social_Types.BlockUser.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.BlockUser.Response in rpc: .BlockUser -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.UnBlockUser", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.UnBlockUser.Response){
					RPC_Social_Default_Handler.UnBlockUser((RPC_Social_Types.UnBlockUser.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.UnBlockUser.Response in rpc: .UnBlockUser -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Social.AddSeenPosts", (pb, handled)->{
				if(pb instanceof RPC_Social_Types.AddSeenPosts.Response){
					RPC_Social_Default_Handler.AddSeenPosts((RPC_Social_Types.AddSeenPosts.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_Social_Types.AddSeenPosts.Response in rpc: .AddSeenPosts -- class: " + pb );//.getClass().getName());
				}
			});
	  
	  
			router.put("RPC_User.UpdateAbout", (pb, handled)->{
				if(pb instanceof RPC_User_Types.UpdateAbout.Response){
					RPC_User_Default_Handler.UpdateAbout((RPC_User_Types.UpdateAbout.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_User_Types.UpdateAbout.Response in rpc: .UpdateAbout -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_User.UpdateUserName", (pb, handled)->{
				if(pb instanceof RPC_User_Types.UpdateUserName.Response){
					RPC_User_Default_Handler.UpdateUserName((RPC_User_Types.UpdateUserName.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_User_Types.UpdateUserName.Response in rpc: .UpdateUserName -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_User.ChangePrivacy", (pb, handled)->{
				if(pb instanceof RPC_User_Types.ChangePrivacy.Response){
					RPC_User_Default_Handler.ChangePrivacy((RPC_User_Types.ChangePrivacy.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_User_Types.ChangePrivacy.Response in rpc: .ChangePrivacy -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_User.ChangeAvatar", (pb, handled)->{
				if(pb instanceof RPC_User_Types.ChangeAvatar.Response){
					RPC_User_Default_Handler.ChangeAvatar((RPC_User_Types.ChangeAvatar.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to RPC_User_Types.ChangeAvatar.Response in rpc: .ChangeAvatar -- class: " + pb );//.getClass().getName());
				}
			});
	  
	}
	
}
/*

RPC_HANDLERS.RPC_Auth RPC_Auth_Default_Handler = new RPC_HANDLERS.RPC_Auth RPC_Auth_Empty();
RPC_HANDLERS.RPC_Chat RPC_Chat_Default_Handler = new RPC_HANDLERS.RPC_Chat RPC_Chat_Empty();
RPC_HANDLERS.RPC_General RPC_General_Default_Handler = new RPC_HANDLERS.RPC_General RPC_General_Empty();
RPC_HANDLERS.RPC_Page RPC_Page_Default_Handler = new RPC_HANDLERS.RPC_Page RPC_Page_Empty();
RPC_HANDLERS.RPC_Social RPC_Social_Default_Handler = new RPC_HANDLERS.RPC_Social RPC_Social_Empty();
RPC_HANDLERS.RPC_User RPC_User_Default_Handler = new RPC_HANDLERS.RPC_User RPC_User_Empty();
	
*/