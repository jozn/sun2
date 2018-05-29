package ir.ms.pb;

import java.util.HashMap;
import java.util.concurrent.ConcurrentHashMap;
import java.util.Map;

public class RPC_HANDLERS {

public interface RPC_Chat {
    void AddNewMessage( PB_RPC_Chat_Types.AddNewMessage.Response pb, boolean handled);
    void SetRoomActionDoing( PB_RPC_Chat_Types.SetRoomActionDoing.Response pb, boolean handled);
    void GetChatList( PB_RPC_Chat_Types.GetChatList.Response pb, boolean handled);
    void GetChatHistory( PB_RPC_Chat_Types.GetChatHistory.Response pb, boolean handled);
    void PushRoomsChange( PB_RPC_Chat_Types.PushRoomsChange.Response pb, boolean handled);
    void GetRoomsChange( PB_RPC_Chat_Types.GetRoomsChange.Response pb, boolean handled);
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
    public void PushRoomsChange( PB_RPC_Chat_Types.PushRoomsChange.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.PushRoomsChange' ");
    }
  	@Override
    public void GetRoomsChange( PB_RPC_Chat_Types.GetRoomsChange.Response pb, boolean handled){
    	Log.d("RPC", " default empty handler for RPC 'RPC_Chat.GetRoomsChange' ");
    }
  }

	/////////////////////////////////// Maper of Handling methods /////////////////////////////////
	public static interface HandleRowRpcResponse {
		void handle(Object pb,boolean handled);
	}


	
	public static RPC_HANDLERS.RPC_Chat RPC_Chat_Default_Handler = new RPC_HANDLERS.RPC_Chat_Empty();


	public static Map<String,HandleRowRpcResponse > router = new HashMap<>();

	public static Map<String,HandleRowRpcResponse > getRouter(){
		if(router.size() ==0 ){
			initMap();
		}
		return router;
	}

	private synchronized static void initMap(){
	
	  
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
	  
			router.put("RPC_Chat.PushRoomsChange", (pb, handled)->{
				if(pb instanceof PB_RPC_Chat_Types.PushRoomsChange.Response){
					RPC_Chat_Default_Handler.PushRoomsChange((PB_RPC_Chat_Types.PushRoomsChange.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_RPC_Chat_Types.PushRoomsChange.Response in rpc: .PushRoomsChange -- class: " + pb );//.getClass().getName());
				}
			});
	  
			router.put("RPC_Chat.GetRoomsChange", (pb, handled)->{
				if(pb instanceof PB_RPC_Chat_Types.GetRoomsChange.Response){
					RPC_Chat_Default_Handler.GetRoomsChange((PB_RPC_Chat_Types.GetRoomsChange.Response) pb, handled);
				}else{
					Log.d("RPC", " can not convert response object to PB_RPC_Chat_Types.GetRoomsChange.Response in rpc: .GetRoomsChange -- class: " + pb );//.getClass().getName());
				}
			});
	  
	}
	
}
/*

RPC_HANDLERS.RPC_Chat RPC_Chat_Default_Handler = new RPC_HANDLERS.RPC_Chat RPC_Chat_Empty();
	
*/