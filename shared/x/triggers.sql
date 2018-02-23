
################################ Action ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS action_OnCreateLogger $$
CREATE TRIGGER action_OnCreateLogger AFTER INSERT ON action
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Action","INSERT",NEW.ActionId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS action_OnUpdateLogger $$
CREATE TRIGGER action_OnUpdateLogger AFTER UPDATE ON action
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Action","UPDATE",NEW.ActionId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS action_OnDeleteLogger $$
CREATE TRIGGER action_OnDeleteLogger AFTER DELETE ON action
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Action","DELETE",OLD.ActionId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Comment ######################################

delimiter $$
DROP TRIGGER IF EXISTS comment_OnCreateLogger $$
CREATE TRIGGER comment_OnCreateLogger AFTER INSERT ON comment
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Comment","INSERT",NEW.CommentId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS comment_OnUpdateLogger $$
CREATE TRIGGER comment_OnUpdateLogger AFTER UPDATE ON comment
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Comment","UPDATE",NEW.CommentId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS comment_OnDeleteLogger $$
CREATE TRIGGER comment_OnDeleteLogger AFTER DELETE ON comment
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Comment","DELETE",OLD.CommentId, UNIX_TIMESTAMP(NOW()));
  END;
$$


delimiter ;
################################ Event ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS event_OnCreateLogger $$
CREATE TRIGGER event_OnCreateLogger AFTER INSERT ON event
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Event","INSERT",NEW.EventId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS event_OnUpdateLogger $$
CREATE TRIGGER event_OnUpdateLogger AFTER UPDATE ON event
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Event","UPDATE",NEW.EventId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS event_OnDeleteLogger $$
CREATE TRIGGER event_OnDeleteLogger AFTER DELETE ON event
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Event","DELETE",OLD.EventId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ FollowingList ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS following_list_OnCreateLogger $$
CREATE TRIGGER following_list_OnCreateLogger AFTER INSERT ON following_list
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingList","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS following_list_OnUpdateLogger $$
CREATE TRIGGER following_list_OnUpdateLogger AFTER UPDATE ON following_list
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingList","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS following_list_OnDeleteLogger $$
CREATE TRIGGER following_list_OnDeleteLogger AFTER DELETE ON following_list
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingList","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ FollowingListMember ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS following_list_member_OnCreateLogger $$
CREATE TRIGGER following_list_member_OnCreateLogger AFTER INSERT ON following_list_member
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMember","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS following_list_member_OnUpdateLogger $$
CREATE TRIGGER following_list_member_OnUpdateLogger AFTER UPDATE ON following_list_member
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMember","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS following_list_member_OnDeleteLogger $$
CREATE TRIGGER following_list_member_OnDeleteLogger AFTER DELETE ON following_list_member
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMember","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ FollowingListMemberRemoved ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS following_list_member_removed_OnCreateLogger $$
CREATE TRIGGER following_list_member_removed_OnCreateLogger AFTER INSERT ON following_list_member_removed
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMemberRemoved","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS following_list_member_removed_OnUpdateLogger $$
CREATE TRIGGER following_list_member_removed_OnUpdateLogger AFTER UPDATE ON following_list_member_removed
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMemberRemoved","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS following_list_member_removed_OnDeleteLogger $$
CREATE TRIGGER following_list_member_removed_OnDeleteLogger AFTER DELETE ON following_list_member_removed
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMemberRemoved","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Group ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS group_OnCreateLogger $$
CREATE TRIGGER group_OnCreateLogger AFTER INSERT ON group
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Group","INSERT",NEW.GroupId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS group_OnUpdateLogger $$
CREATE TRIGGER group_OnUpdateLogger AFTER UPDATE ON group
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Group","UPDATE",NEW.GroupId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS group_OnDeleteLogger $$
CREATE TRIGGER group_OnDeleteLogger AFTER DELETE ON group
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Group","DELETE",OLD.GroupId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ GroupMember ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS group_member_OnCreateLogger $$
CREATE TRIGGER group_member_OnCreateLogger AFTER INSERT ON group_member
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMember","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS group_member_OnUpdateLogger $$
CREATE TRIGGER group_member_OnUpdateLogger AFTER UPDATE ON group_member
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMember","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS group_member_OnDeleteLogger $$
CREATE TRIGGER group_member_OnDeleteLogger AFTER DELETE ON group_member
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMember","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ GroupMessage ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS group_message_OnCreateLogger $$
CREATE TRIGGER group_message_OnCreateLogger AFTER INSERT ON group_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMessage","INSERT",NEW.MessageId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS group_message_OnUpdateLogger $$
CREATE TRIGGER group_message_OnUpdateLogger AFTER UPDATE ON group_message
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMessage","UPDATE",NEW.MessageId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS group_message_OnDeleteLogger $$
CREATE TRIGGER group_message_OnDeleteLogger AFTER DELETE ON group_message
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMessage","DELETE",OLD.MessageId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Like ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS likes_OnCreateLogger $$
CREATE TRIGGER likes_OnCreateLogger AFTER INSERT ON likes
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Like","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS likes_OnUpdateLogger $$
CREATE TRIGGER likes_OnUpdateLogger AFTER UPDATE ON likes
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Like","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS likes_OnDeleteLogger $$
CREATE TRIGGER likes_OnDeleteLogger AFTER DELETE ON likes
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Like","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Media ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS media_OnCreateLogger $$
CREATE TRIGGER media_OnCreateLogger AFTER INSERT ON media
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Media","INSERT",NEW.MediaId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS media_OnUpdateLogger $$
CREATE TRIGGER media_OnUpdateLogger AFTER UPDATE ON media
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Media","UPDATE",NEW.MediaId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS media_OnDeleteLogger $$
CREATE TRIGGER media_OnDeleteLogger AFTER DELETE ON media
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Media","DELETE",OLD.MediaId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Notify ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS notify_OnCreateLogger $$
CREATE TRIGGER notify_OnCreateLogger AFTER INSERT ON notify
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Notify","INSERT",NEW.NotifyId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS notify_OnUpdateLogger $$
CREATE TRIGGER notify_OnUpdateLogger AFTER UPDATE ON notify
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Notify","UPDATE",NEW.NotifyId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS notify_OnDeleteLogger $$
CREATE TRIGGER notify_OnDeleteLogger AFTER DELETE ON notify
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Notify","DELETE",OLD.NotifyId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ NotifyRemoved ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS notify_removed_OnCreateLogger $$
CREATE TRIGGER notify_removed_OnCreateLogger AFTER INSERT ON notify_removed
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("NotifyRemoved","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS notify_removed_OnUpdateLogger $$
CREATE TRIGGER notify_removed_OnUpdateLogger AFTER UPDATE ON notify_removed
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("NotifyRemoved","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS notify_removed_OnDeleteLogger $$
CREATE TRIGGER notify_removed_OnDeleteLogger AFTER DELETE ON notify_removed
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("NotifyRemoved","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ PhoneContact ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS phone_contacts_OnCreateLogger $$
CREATE TRIGGER phone_contacts_OnCreateLogger AFTER INSERT ON phone_contacts
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PhoneContact","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS phone_contacts_OnUpdateLogger $$
CREATE TRIGGER phone_contacts_OnUpdateLogger AFTER UPDATE ON phone_contacts
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PhoneContact","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS phone_contacts_OnDeleteLogger $$
CREATE TRIGGER phone_contacts_OnDeleteLogger AFTER DELETE ON phone_contacts
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PhoneContact","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Post ######################################

delimiter $$
DROP TRIGGER IF EXISTS post_OnCreateLogger $$
CREATE TRIGGER post_OnCreateLogger AFTER INSERT ON post
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Post","INSERT",NEW.PostId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS post_OnUpdateLogger $$
CREATE TRIGGER post_OnUpdateLogger AFTER UPDATE ON post
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Post","UPDATE",NEW.PostId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS post_OnDeleteLogger $$
CREATE TRIGGER post_OnDeleteLogger AFTER DELETE ON post
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Post","DELETE",OLD.PostId, UNIX_TIMESTAMP(NOW()));
  END;
$$


delimiter ;
################################ PostKey ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS post_keys_OnCreateLogger $$
CREATE TRIGGER post_keys_OnCreateLogger AFTER INSERT ON post_keys
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostKey","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS post_keys_OnUpdateLogger $$
CREATE TRIGGER post_keys_OnUpdateLogger AFTER UPDATE ON post_keys
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostKey","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS post_keys_OnDeleteLogger $$
CREATE TRIGGER post_keys_OnDeleteLogger AFTER DELETE ON post_keys
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostKey","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ SearchClicked ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS search_clicked_OnCreateLogger $$
CREATE TRIGGER search_clicked_OnCreateLogger AFTER INSERT ON search_clicked
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SearchClicked","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS search_clicked_OnUpdateLogger $$
CREATE TRIGGER search_clicked_OnUpdateLogger AFTER UPDATE ON search_clicked
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SearchClicked","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS search_clicked_OnDeleteLogger $$
CREATE TRIGGER search_clicked_OnDeleteLogger AFTER DELETE ON search_clicked
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SearchClicked","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Session ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS session_OnCreateLogger $$
CREATE TRIGGER session_OnCreateLogger AFTER INSERT ON session
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Session","INSERT",NEW.SessionUuid, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS session_OnUpdateLogger $$
CREATE TRIGGER session_OnUpdateLogger AFTER UPDATE ON session
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Session","UPDATE",NEW.SessionUuid, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS session_OnDeleteLogger $$
CREATE TRIGGER session_OnDeleteLogger AFTER DELETE ON session
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Session","DELETE",OLD.SessionUuid, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ SettingClient ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS setting_client_OnCreateLogger $$
CREATE TRIGGER setting_client_OnCreateLogger AFTER INSERT ON setting_client
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingClient","INSERT",NEW.UserId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS setting_client_OnUpdateLogger $$
CREATE TRIGGER setting_client_OnUpdateLogger AFTER UPDATE ON setting_client
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingClient","UPDATE",NEW.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS setting_client_OnDeleteLogger $$
CREATE TRIGGER setting_client_OnDeleteLogger AFTER DELETE ON setting_client
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingClient","DELETE",OLD.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ SettingNotification ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS setting_notifications_OnCreateLogger $$
CREATE TRIGGER setting_notifications_OnCreateLogger AFTER INSERT ON setting_notifications
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingNotification","INSERT",NEW.UserId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS setting_notifications_OnUpdateLogger $$
CREATE TRIGGER setting_notifications_OnUpdateLogger AFTER UPDATE ON setting_notifications
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingNotification","UPDATE",NEW.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS setting_notifications_OnDeleteLogger $$
CREATE TRIGGER setting_notifications_OnDeleteLogger AFTER DELETE ON setting_notifications
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingNotification","DELETE",OLD.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ SuggestedTopPost ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS suggested_top_posts_OnCreateLogger $$
CREATE TRIGGER suggested_top_posts_OnCreateLogger AFTER INSERT ON suggested_top_posts
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SuggestedTopPost","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS suggested_top_posts_OnUpdateLogger $$
CREATE TRIGGER suggested_top_posts_OnUpdateLogger AFTER UPDATE ON suggested_top_posts
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SuggestedTopPost","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS suggested_top_posts_OnDeleteLogger $$
CREATE TRIGGER suggested_top_posts_OnDeleteLogger AFTER DELETE ON suggested_top_posts
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SuggestedTopPost","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ SuggestedUser ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS suggested_user_OnCreateLogger $$
CREATE TRIGGER suggested_user_OnCreateLogger AFTER INSERT ON suggested_user
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SuggestedUser","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS suggested_user_OnUpdateLogger $$
CREATE TRIGGER suggested_user_OnUpdateLogger AFTER UPDATE ON suggested_user
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SuggestedUser","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS suggested_user_OnDeleteLogger $$
CREATE TRIGGER suggested_user_OnDeleteLogger AFTER DELETE ON suggested_user
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SuggestedUser","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Tag ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS tag_OnCreateLogger $$
CREATE TRIGGER tag_OnCreateLogger AFTER INSERT ON tag
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Tag","INSERT",NEW.TagId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS tag_OnUpdateLogger $$
CREATE TRIGGER tag_OnUpdateLogger AFTER UPDATE ON tag
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Tag","UPDATE",NEW.TagId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS tag_OnDeleteLogger $$
CREATE TRIGGER tag_OnDeleteLogger AFTER DELETE ON tag
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Tag","DELETE",OLD.TagId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ TagsPost ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS tags_posts_OnCreateLogger $$
CREATE TRIGGER tags_posts_OnCreateLogger AFTER INSERT ON tags_posts
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TagsPost","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS tags_posts_OnUpdateLogger $$
CREATE TRIGGER tags_posts_OnUpdateLogger AFTER UPDATE ON tags_posts
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TagsPost","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS tags_posts_OnDeleteLogger $$
CREATE TRIGGER tags_posts_OnDeleteLogger AFTER DELETE ON tags_posts
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TagsPost","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ TriggerLog ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS trigger_log_OnCreateLogger $$
CREATE TRIGGER trigger_log_OnCreateLogger AFTER INSERT ON trigger_log
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TriggerLog","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS trigger_log_OnUpdateLogger $$
CREATE TRIGGER trigger_log_OnUpdateLogger AFTER UPDATE ON trigger_log
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TriggerLog","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS trigger_log_OnDeleteLogger $$
CREATE TRIGGER trigger_log_OnDeleteLogger AFTER DELETE ON trigger_log
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TriggerLog","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ User ######################################

delimiter $$
DROP TRIGGER IF EXISTS user_OnCreateLogger $$
CREATE TRIGGER user_OnCreateLogger AFTER INSERT ON user
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("User","INSERT",NEW.UserId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS user_OnUpdateLogger $$
CREATE TRIGGER user_OnUpdateLogger AFTER UPDATE ON user
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("User","UPDATE",NEW.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS user_OnDeleteLogger $$
CREATE TRIGGER user_OnDeleteLogger AFTER DELETE ON user
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("User","DELETE",OLD.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$


delimiter ;
################################ UserMetaInfo ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS user_meta_info_OnCreateLogger $$
CREATE TRIGGER user_meta_info_OnCreateLogger AFTER INSERT ON user_meta_info
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserMetaInfo","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS user_meta_info_OnUpdateLogger $$
CREATE TRIGGER user_meta_info_OnUpdateLogger AFTER UPDATE ON user_meta_info
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserMetaInfo","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS user_meta_info_OnDeleteLogger $$
CREATE TRIGGER user_meta_info_OnDeleteLogger AFTER DELETE ON user_meta_info
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserMetaInfo","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ UserPassword ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS user_password_OnCreateLogger $$
CREATE TRIGGER user_password_OnCreateLogger AFTER INSERT ON user_password
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserPassword","INSERT",NEW.UserId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS user_password_OnUpdateLogger $$
CREATE TRIGGER user_password_OnUpdateLogger AFTER UPDATE ON user_password
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserPassword","UPDATE",NEW.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS user_password_OnDeleteLogger $$
CREATE TRIGGER user_password_OnDeleteLogger AFTER DELETE ON user_password
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserPassword","DELETE",OLD.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Chat ######################################

delimiter $$
DROP TRIGGER IF EXISTS chat_OnCreateLogger $$
CREATE TRIGGER chat_OnCreateLogger AFTER INSERT ON chat
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Chat","INSERT",NEW.ChatKey, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS chat_OnUpdateLogger $$
CREATE TRIGGER chat_OnUpdateLogger AFTER UPDATE ON chat
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Chat","UPDATE",NEW.ChatKey, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS chat_OnDeleteLogger $$
CREATE TRIGGER chat_OnDeleteLogger AFTER DELETE ON chat
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Chat","DELETE",OLD.ChatKey, UNIX_TIMESTAMP(NOW()));
  END;
$$


delimiter ;
################################ ChatLastMessage ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS chat_last_message_OnCreateLogger $$
CREATE TRIGGER chat_last_message_OnCreateLogger AFTER INSERT ON chat_last_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("ChatLastMessage","INSERT",NEW.ChatKey, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS chat_last_message_OnUpdateLogger $$
CREATE TRIGGER chat_last_message_OnUpdateLogger AFTER UPDATE ON chat_last_message
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("ChatLastMessage","UPDATE",NEW.ChatKey, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS chat_last_message_OnDeleteLogger $$
CREATE TRIGGER chat_last_message_OnDeleteLogger AFTER DELETE ON chat_last_message
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("ChatLastMessage","DELETE",OLD.ChatKey, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ ChatSync ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS chat_sync_OnCreateLogger $$
CREATE TRIGGER chat_sync_OnCreateLogger AFTER INSERT ON chat_sync
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ChatSync","INSERT",NEW.SyncId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS chat_sync_OnUpdateLogger $$
CREATE TRIGGER chat_sync_OnUpdateLogger AFTER UPDATE ON chat_sync
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ChatSync","UPDATE",NEW.SyncId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS chat_sync_OnDeleteLogger $$
CREATE TRIGGER chat_sync_OnDeleteLogger AFTER DELETE ON chat_sync
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ChatSync","DELETE",OLD.SyncId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ DirectMessage ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS direct_message_OnCreateLogger $$
CREATE TRIGGER direct_message_OnCreateLogger AFTER INSERT ON direct_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectMessage","INSERT",NEW.MessageId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS direct_message_OnUpdateLogger $$
CREATE TRIGGER direct_message_OnUpdateLogger AFTER UPDATE ON direct_message
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectMessage","UPDATE",NEW.MessageId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS direct_message_OnDeleteLogger $$
CREATE TRIGGER direct_message_OnDeleteLogger AFTER DELETE ON direct_message
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectMessage","DELETE",OLD.MessageId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Home ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS home_OnCreateLogger $$
CREATE TRIGGER home_OnCreateLogger AFTER INSERT ON home
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Home","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS home_OnUpdateLogger $$
CREATE TRIGGER home_OnUpdateLogger AFTER UPDATE ON home
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Home","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS home_OnDeleteLogger $$
CREATE TRIGGER home_OnDeleteLogger AFTER DELETE ON home
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Home","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ MessageFile ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS message_file_OnCreateLogger $$
CREATE TRIGGER message_file_OnCreateLogger AFTER INSERT ON message_file
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("MessageFile","INSERT",NEW.MessageFileId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS message_file_OnUpdateLogger $$
CREATE TRIGGER message_file_OnUpdateLogger AFTER UPDATE ON message_file
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("MessageFile","UPDATE",NEW.MessageFileId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS message_file_OnDeleteLogger $$
CREATE TRIGGER message_file_OnDeleteLogger AFTER DELETE ON message_file
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("MessageFile","DELETE",OLD.MessageFileId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ FileMsg ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS file_msg_OnCreateLogger $$
CREATE TRIGGER file_msg_OnCreateLogger AFTER INSERT ON file_msg
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FileMsg","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS file_msg_OnUpdateLogger $$
CREATE TRIGGER file_msg_OnUpdateLogger AFTER UPDATE ON file_msg
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FileMsg","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS file_msg_OnDeleteLogger $$
CREATE TRIGGER file_msg_OnDeleteLogger AFTER DELETE ON file_msg
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FileMsg","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ FilePost ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS file_post_OnCreateLogger $$
CREATE TRIGGER file_post_OnCreateLogger AFTER INSERT ON file_post
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FilePost","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS file_post_OnUpdateLogger $$
CREATE TRIGGER file_post_OnUpdateLogger AFTER UPDATE ON file_post
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FilePost","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS file_post_OnDeleteLogger $$
CREATE TRIGGER file_post_OnDeleteLogger AFTER DELETE ON file_post
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FilePost","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/

###############################################################################################
################################## Delete of all triggers #####################################
/*

### Action ##
DROP TRIGGER IF EXISTS action_OnCreateLogger ;
DROP TRIGGER IF EXISTS action_OnUpdateLogger ;
DROP TRIGGER IF EXISTS action_OnDeleteLogger ;
### Comment ##
DROP TRIGGER IF EXISTS comment_OnCreateLogger ;
DROP TRIGGER IF EXISTS comment_OnUpdateLogger ;
DROP TRIGGER IF EXISTS comment_OnDeleteLogger ;
### Event ##
DROP TRIGGER IF EXISTS event_OnCreateLogger ;
DROP TRIGGER IF EXISTS event_OnUpdateLogger ;
DROP TRIGGER IF EXISTS event_OnDeleteLogger ;
### FollowingList ##
DROP TRIGGER IF EXISTS following_list_OnCreateLogger ;
DROP TRIGGER IF EXISTS following_list_OnUpdateLogger ;
DROP TRIGGER IF EXISTS following_list_OnDeleteLogger ;
### FollowingListMember ##
DROP TRIGGER IF EXISTS following_list_member_OnCreateLogger ;
DROP TRIGGER IF EXISTS following_list_member_OnUpdateLogger ;
DROP TRIGGER IF EXISTS following_list_member_OnDeleteLogger ;
### FollowingListMemberRemoved ##
DROP TRIGGER IF EXISTS following_list_member_removed_OnCreateLogger ;
DROP TRIGGER IF EXISTS following_list_member_removed_OnUpdateLogger ;
DROP TRIGGER IF EXISTS following_list_member_removed_OnDeleteLogger ;
### Group ##
DROP TRIGGER IF EXISTS group_OnCreateLogger ;
DROP TRIGGER IF EXISTS group_OnUpdateLogger ;
DROP TRIGGER IF EXISTS group_OnDeleteLogger ;
### GroupMember ##
DROP TRIGGER IF EXISTS group_member_OnCreateLogger ;
DROP TRIGGER IF EXISTS group_member_OnUpdateLogger ;
DROP TRIGGER IF EXISTS group_member_OnDeleteLogger ;
### GroupMessage ##
DROP TRIGGER IF EXISTS group_message_OnCreateLogger ;
DROP TRIGGER IF EXISTS group_message_OnUpdateLogger ;
DROP TRIGGER IF EXISTS group_message_OnDeleteLogger ;
### Like ##
DROP TRIGGER IF EXISTS likes_OnCreateLogger ;
DROP TRIGGER IF EXISTS likes_OnUpdateLogger ;
DROP TRIGGER IF EXISTS likes_OnDeleteLogger ;
### Media ##
DROP TRIGGER IF EXISTS media_OnCreateLogger ;
DROP TRIGGER IF EXISTS media_OnUpdateLogger ;
DROP TRIGGER IF EXISTS media_OnDeleteLogger ;
### Notify ##
DROP TRIGGER IF EXISTS notify_OnCreateLogger ;
DROP TRIGGER IF EXISTS notify_OnUpdateLogger ;
DROP TRIGGER IF EXISTS notify_OnDeleteLogger ;
### NotifyRemoved ##
DROP TRIGGER IF EXISTS notify_removed_OnCreateLogger ;
DROP TRIGGER IF EXISTS notify_removed_OnUpdateLogger ;
DROP TRIGGER IF EXISTS notify_removed_OnDeleteLogger ;
### PhoneContact ##
DROP TRIGGER IF EXISTS phone_contacts_OnCreateLogger ;
DROP TRIGGER IF EXISTS phone_contacts_OnUpdateLogger ;
DROP TRIGGER IF EXISTS phone_contacts_OnDeleteLogger ;
### Post ##
DROP TRIGGER IF EXISTS post_OnCreateLogger ;
DROP TRIGGER IF EXISTS post_OnUpdateLogger ;
DROP TRIGGER IF EXISTS post_OnDeleteLogger ;
### PostKey ##
DROP TRIGGER IF EXISTS post_keys_OnCreateLogger ;
DROP TRIGGER IF EXISTS post_keys_OnUpdateLogger ;
DROP TRIGGER IF EXISTS post_keys_OnDeleteLogger ;
### SearchClicked ##
DROP TRIGGER IF EXISTS search_clicked_OnCreateLogger ;
DROP TRIGGER IF EXISTS search_clicked_OnUpdateLogger ;
DROP TRIGGER IF EXISTS search_clicked_OnDeleteLogger ;
### Session ##
DROP TRIGGER IF EXISTS session_OnCreateLogger ;
DROP TRIGGER IF EXISTS session_OnUpdateLogger ;
DROP TRIGGER IF EXISTS session_OnDeleteLogger ;
### SettingClient ##
DROP TRIGGER IF EXISTS setting_client_OnCreateLogger ;
DROP TRIGGER IF EXISTS setting_client_OnUpdateLogger ;
DROP TRIGGER IF EXISTS setting_client_OnDeleteLogger ;
### SettingNotification ##
DROP TRIGGER IF EXISTS setting_notifications_OnCreateLogger ;
DROP TRIGGER IF EXISTS setting_notifications_OnUpdateLogger ;
DROP TRIGGER IF EXISTS setting_notifications_OnDeleteLogger ;
### SuggestedTopPost ##
DROP TRIGGER IF EXISTS suggested_top_posts_OnCreateLogger ;
DROP TRIGGER IF EXISTS suggested_top_posts_OnUpdateLogger ;
DROP TRIGGER IF EXISTS suggested_top_posts_OnDeleteLogger ;
### SuggestedUser ##
DROP TRIGGER IF EXISTS suggested_user_OnCreateLogger ;
DROP TRIGGER IF EXISTS suggested_user_OnUpdateLogger ;
DROP TRIGGER IF EXISTS suggested_user_OnDeleteLogger ;
### Tag ##
DROP TRIGGER IF EXISTS tag_OnCreateLogger ;
DROP TRIGGER IF EXISTS tag_OnUpdateLogger ;
DROP TRIGGER IF EXISTS tag_OnDeleteLogger ;
### TagsPost ##
DROP TRIGGER IF EXISTS tags_posts_OnCreateLogger ;
DROP TRIGGER IF EXISTS tags_posts_OnUpdateLogger ;
DROP TRIGGER IF EXISTS tags_posts_OnDeleteLogger ;
### TriggerLog ##
DROP TRIGGER IF EXISTS trigger_log_OnCreateLogger ;
DROP TRIGGER IF EXISTS trigger_log_OnUpdateLogger ;
DROP TRIGGER IF EXISTS trigger_log_OnDeleteLogger ;
### User ##
DROP TRIGGER IF EXISTS user_OnCreateLogger ;
DROP TRIGGER IF EXISTS user_OnUpdateLogger ;
DROP TRIGGER IF EXISTS user_OnDeleteLogger ;
### UserMetaInfo ##
DROP TRIGGER IF EXISTS user_meta_info_OnCreateLogger ;
DROP TRIGGER IF EXISTS user_meta_info_OnUpdateLogger ;
DROP TRIGGER IF EXISTS user_meta_info_OnDeleteLogger ;
### UserPassword ##
DROP TRIGGER IF EXISTS user_password_OnCreateLogger ;
DROP TRIGGER IF EXISTS user_password_OnUpdateLogger ;
DROP TRIGGER IF EXISTS user_password_OnDeleteLogger ;
### Chat ##
DROP TRIGGER IF EXISTS chat_OnCreateLogger ;
DROP TRIGGER IF EXISTS chat_OnUpdateLogger ;
DROP TRIGGER IF EXISTS chat_OnDeleteLogger ;
### ChatLastMessage ##
DROP TRIGGER IF EXISTS chat_last_message_OnCreateLogger ;
DROP TRIGGER IF EXISTS chat_last_message_OnUpdateLogger ;
DROP TRIGGER IF EXISTS chat_last_message_OnDeleteLogger ;
### ChatSync ##
DROP TRIGGER IF EXISTS chat_sync_OnCreateLogger ;
DROP TRIGGER IF EXISTS chat_sync_OnUpdateLogger ;
DROP TRIGGER IF EXISTS chat_sync_OnDeleteLogger ;
### DirectMessage ##
DROP TRIGGER IF EXISTS direct_message_OnCreateLogger ;
DROP TRIGGER IF EXISTS direct_message_OnUpdateLogger ;
DROP TRIGGER IF EXISTS direct_message_OnDeleteLogger ;
### Home ##
DROP TRIGGER IF EXISTS home_OnCreateLogger ;
DROP TRIGGER IF EXISTS home_OnUpdateLogger ;
DROP TRIGGER IF EXISTS home_OnDeleteLogger ;
### MessageFile ##
DROP TRIGGER IF EXISTS message_file_OnCreateLogger ;
DROP TRIGGER IF EXISTS message_file_OnUpdateLogger ;
DROP TRIGGER IF EXISTS message_file_OnDeleteLogger ;
### FileMsg ##
DROP TRIGGER IF EXISTS file_msg_OnCreateLogger ;
DROP TRIGGER IF EXISTS file_msg_OnUpdateLogger ;
DROP TRIGGER IF EXISTS file_msg_OnDeleteLogger ;
### FilePost ##
DROP TRIGGER IF EXISTS file_post_OnCreateLogger ;
DROP TRIGGER IF EXISTS file_post_OnUpdateLogger ;
DROP TRIGGER IF EXISTS file_post_OnDeleteLogger ;
*/