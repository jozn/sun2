
################################ Action ######################################

delimiter $$
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


delimiter ;
################################ Blocked ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS blocked_OnCreateLogger $$
CREATE TRIGGER blocked_OnCreateLogger AFTER INSERT ON blocked
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Blocked","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS blocked_OnUpdateLogger $$
CREATE TRIGGER blocked_OnUpdateLogger AFTER UPDATE ON blocked
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Blocked","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS blocked_OnDeleteLogger $$
CREATE TRIGGER blocked_OnDeleteLogger AFTER DELETE ON blocked
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Blocked","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
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
################################ CommentDeleted ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS comment_deleted_OnCreateLogger $$
CREATE TRIGGER comment_deleted_OnCreateLogger AFTER INSERT ON comment_deleted
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("CommentDeleted","INSERT",NEW.CommentId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS comment_deleted_OnUpdateLogger $$
CREATE TRIGGER comment_deleted_OnUpdateLogger AFTER UPDATE ON comment_deleted
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("CommentDeleted","UPDATE",NEW.CommentId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS comment_deleted_OnDeleteLogger $$
CREATE TRIGGER comment_deleted_OnDeleteLogger AFTER DELETE ON comment_deleted
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("CommentDeleted","DELETE",OLD.CommentId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
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
################################ Followed ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS followed_OnCreateLogger $$
CREATE TRIGGER followed_OnCreateLogger AFTER INSERT ON followed
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Followed","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS followed_OnUpdateLogger $$
CREATE TRIGGER followed_OnUpdateLogger AFTER UPDATE ON followed
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Followed","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS followed_OnDeleteLogger $$
CREATE TRIGGER followed_OnDeleteLogger AFTER DELETE ON followed
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Followed","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Likes ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS likes_OnCreateLogger $$
CREATE TRIGGER likes_OnCreateLogger AFTER INSERT ON likes
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Likes","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS likes_OnUpdateLogger $$
CREATE TRIGGER likes_OnUpdateLogger AFTER UPDATE ON likes
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Likes","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS likes_OnDeleteLogger $$
CREATE TRIGGER likes_OnDeleteLogger AFTER DELETE ON likes
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Likes","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
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
################################ PhoneContacts ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS phone_contacts_OnCreateLogger $$
CREATE TRIGGER phone_contacts_OnCreateLogger AFTER INSERT ON phone_contacts
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PhoneContacts","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS phone_contacts_OnUpdateLogger $$
CREATE TRIGGER phone_contacts_OnUpdateLogger AFTER UPDATE ON phone_contacts
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PhoneContacts","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS phone_contacts_OnDeleteLogger $$
CREATE TRIGGER phone_contacts_OnDeleteLogger AFTER DELETE ON phone_contacts
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PhoneContacts","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
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
################################ PostCount ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS post_count_OnCreateLogger $$
CREATE TRIGGER post_count_OnCreateLogger AFTER INSERT ON post_count
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostCount","INSERT",NEW.PostId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS post_count_OnUpdateLogger $$
CREATE TRIGGER post_count_OnUpdateLogger AFTER UPDATE ON post_count
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostCount","UPDATE",NEW.PostId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS post_count_OnDeleteLogger $$
CREATE TRIGGER post_count_OnDeleteLogger AFTER DELETE ON post_count
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostCount","DELETE",OLD.PostId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ PostDeleted ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS post_deleted_OnCreateLogger $$
CREATE TRIGGER post_deleted_OnCreateLogger AFTER INSERT ON post_deleted
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostDeleted","INSERT",NEW.PostId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS post_deleted_OnUpdateLogger $$
CREATE TRIGGER post_deleted_OnUpdateLogger AFTER UPDATE ON post_deleted
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostDeleted","UPDATE",NEW.PostId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS post_deleted_OnDeleteLogger $$
CREATE TRIGGER post_deleted_OnDeleteLogger AFTER DELETE ON post_deleted
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostDeleted","DELETE",OLD.PostId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ PostKeys ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS post_keys_OnCreateLogger $$
CREATE TRIGGER post_keys_OnCreateLogger AFTER INSERT ON post_keys
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostKeys","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS post_keys_OnUpdateLogger $$
CREATE TRIGGER post_keys_OnUpdateLogger AFTER UPDATE ON post_keys
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostKeys","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS post_keys_OnDeleteLogger $$
CREATE TRIGGER post_keys_OnDeleteLogger AFTER DELETE ON post_keys
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostKeys","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ PostLink ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS post_link_OnCreateLogger $$
CREATE TRIGGER post_link_OnCreateLogger AFTER INSERT ON post_link
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostLink","INSERT",NEW.LinkId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS post_link_OnUpdateLogger $$
CREATE TRIGGER post_link_OnUpdateLogger AFTER UPDATE ON post_link
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostLink","UPDATE",NEW.LinkId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS post_link_OnDeleteLogger $$
CREATE TRIGGER post_link_OnDeleteLogger AFTER DELETE ON post_link
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostLink","DELETE",OLD.LinkId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ PostMedia ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS post_media_OnCreateLogger $$
CREATE TRIGGER post_media_OnCreateLogger AFTER INSERT ON post_media
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostMedia","INSERT",NEW.MediaId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS post_media_OnUpdateLogger $$
CREATE TRIGGER post_media_OnUpdateLogger AFTER UPDATE ON post_media
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostMedia","UPDATE",NEW.MediaId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS post_media_OnDeleteLogger $$
CREATE TRIGGER post_media_OnDeleteLogger AFTER DELETE ON post_media
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostMedia","DELETE",OLD.MediaId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ PostPromoted ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS post_promoted_OnCreateLogger $$
CREATE TRIGGER post_promoted_OnCreateLogger AFTER INSERT ON post_promoted
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostPromoted","INSERT",NEW.PromoteId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS post_promoted_OnUpdateLogger $$
CREATE TRIGGER post_promoted_OnUpdateLogger AFTER UPDATE ON post_promoted
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostPromoted","UPDATE",NEW.PromoteId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS post_promoted_OnDeleteLogger $$
CREATE TRIGGER post_promoted_OnDeleteLogger AFTER DELETE ON post_promoted
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostPromoted","DELETE",OLD.PromoteId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ PostReshared ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS post_reshared_OnCreateLogger $$
CREATE TRIGGER post_reshared_OnCreateLogger AFTER INSERT ON post_reshared
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostReshared","INSERT",NEW.ResharedId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS post_reshared_OnUpdateLogger $$
CREATE TRIGGER post_reshared_OnUpdateLogger AFTER UPDATE ON post_reshared
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostReshared","UPDATE",NEW.ResharedId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS post_reshared_OnDeleteLogger $$
CREATE TRIGGER post_reshared_OnDeleteLogger AFTER DELETE ON post_reshared
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PostReshared","DELETE",OLD.ResharedId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ ProfileAll ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS profile_all_OnCreateLogger $$
CREATE TRIGGER profile_all_OnCreateLogger AFTER INSERT ON profile_all
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ProfileAll","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS profile_all_OnUpdateLogger $$
CREATE TRIGGER profile_all_OnUpdateLogger AFTER UPDATE ON profile_all
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ProfileAll","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS profile_all_OnDeleteLogger $$
CREATE TRIGGER profile_all_OnDeleteLogger AFTER DELETE ON profile_all
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ProfileAll","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ ProfileMedia ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS profile_media_OnCreateLogger $$
CREATE TRIGGER profile_media_OnCreateLogger AFTER INSERT ON profile_media
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ProfileMedia","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS profile_media_OnUpdateLogger $$
CREATE TRIGGER profile_media_OnUpdateLogger AFTER UPDATE ON profile_media
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ProfileMedia","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS profile_media_OnDeleteLogger $$
CREATE TRIGGER profile_media_OnDeleteLogger AFTER DELETE ON profile_media
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ProfileMedia","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ ProfileMentioned ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS profile_mentioned_OnCreateLogger $$
CREATE TRIGGER profile_mentioned_OnCreateLogger AFTER INSERT ON profile_mentioned
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ProfileMentioned","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS profile_mentioned_OnUpdateLogger $$
CREATE TRIGGER profile_mentioned_OnUpdateLogger AFTER UPDATE ON profile_mentioned
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ProfileMentioned","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS profile_mentioned_OnDeleteLogger $$
CREATE TRIGGER profile_mentioned_OnDeleteLogger AFTER DELETE ON profile_mentioned
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ProfileMentioned","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Session ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS session_OnCreateLogger $$
CREATE TRIGGER session_OnCreateLogger AFTER INSERT ON session
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Session","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS session_OnUpdateLogger $$
CREATE TRIGGER session_OnUpdateLogger AFTER UPDATE ON session
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Session","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS session_OnDeleteLogger $$
CREATE TRIGGER session_OnDeleteLogger AFTER DELETE ON session
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Session","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ SettingNotifications ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS setting_notifications_OnCreateLogger $$
CREATE TRIGGER setting_notifications_OnCreateLogger AFTER INSERT ON setting_notifications
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingNotifications","INSERT",NEW.UserId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS setting_notifications_OnUpdateLogger $$
CREATE TRIGGER setting_notifications_OnUpdateLogger AFTER UPDATE ON setting_notifications
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingNotifications","UPDATE",NEW.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS setting_notifications_OnDeleteLogger $$
CREATE TRIGGER setting_notifications_OnDeleteLogger AFTER DELETE ON setting_notifications
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingNotifications","DELETE",OLD.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Sms ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS sms_OnCreateLogger $$
CREATE TRIGGER sms_OnCreateLogger AFTER INSERT ON sms
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Sms","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS sms_OnUpdateLogger $$
CREATE TRIGGER sms_OnUpdateLogger AFTER UPDATE ON sms
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Sms","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS sms_OnDeleteLogger $$
CREATE TRIGGER sms_OnDeleteLogger AFTER DELETE ON sms
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Sms","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
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
################################ TagPost ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS tag_post_OnCreateLogger $$
CREATE TRIGGER tag_post_OnCreateLogger AFTER INSERT ON tag_post
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TagPost","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS tag_post_OnUpdateLogger $$
CREATE TRIGGER tag_post_OnUpdateLogger AFTER UPDATE ON tag_post
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TagPost","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS tag_post_OnDeleteLogger $$
CREATE TRIGGER tag_post_OnDeleteLogger AFTER DELETE ON tag_post
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TagPost","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
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
################################ UserRelation ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS user_relation_OnCreateLogger $$
CREATE TRIGGER user_relation_OnCreateLogger AFTER INSERT ON user_relation
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserRelation","INSERT",NEW.RelNanoId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS user_relation_OnUpdateLogger $$
CREATE TRIGGER user_relation_OnUpdateLogger AFTER UPDATE ON user_relation
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserRelation","UPDATE",NEW.RelNanoId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS user_relation_OnDeleteLogger $$
CREATE TRIGGER user_relation_OnDeleteLogger AFTER DELETE ON user_relation
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserRelation","DELETE",OLD.RelNanoId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Chat ######################################

delimiter $$
DROP TRIGGER IF EXISTS chat_OnCreateLogger $$
CREATE TRIGGER chat_OnCreateLogger AFTER INSERT ON chat
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Chat","INSERT",NEW.ChatId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS chat_OnUpdateLogger $$
CREATE TRIGGER chat_OnUpdateLogger AFTER UPDATE ON chat
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Chat","UPDATE",NEW.ChatId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS chat_OnDeleteLogger $$
CREATE TRIGGER chat_OnDeleteLogger AFTER DELETE ON chat
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Chat","DELETE",OLD.ChatId, UNIX_TIMESTAMP(NOW()));
  END;
$$


delimiter ;
################################ ChatDeleted ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS chat_deleted_OnCreateLogger $$
CREATE TRIGGER chat_deleted_OnCreateLogger AFTER INSERT ON chat_deleted
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ChatDeleted","INSERT",NEW.ChatId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS chat_deleted_OnUpdateLogger $$
CREATE TRIGGER chat_deleted_OnUpdateLogger AFTER UPDATE ON chat_deleted
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ChatDeleted","UPDATE",NEW.ChatId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS chat_deleted_OnDeleteLogger $$
CREATE TRIGGER chat_deleted_OnDeleteLogger AFTER DELETE ON chat_deleted
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ChatDeleted","DELETE",OLD.ChatId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ ChatLastMessage ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS chat_last_message_OnCreateLogger $$
CREATE TRIGGER chat_last_message_OnCreateLogger AFTER INSERT ON chat_last_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("ChatLastMessage","INSERT",NEW.ChatIdGroupId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS chat_last_message_OnUpdateLogger $$
CREATE TRIGGER chat_last_message_OnUpdateLogger AFTER UPDATE ON chat_last_message
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("ChatLastMessage","UPDATE",NEW.ChatIdGroupId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS chat_last_message_OnDeleteLogger $$
CREATE TRIGGER chat_last_message_OnDeleteLogger AFTER DELETE ON chat_last_message
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("ChatLastMessage","DELETE",OLD.ChatIdGroupId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ ChatUserVersion ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS chat_user_version_OnCreateLogger $$
CREATE TRIGGER chat_user_version_OnCreateLogger AFTER INSERT ON chat_user_version
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ChatUserVersion","INSERT",NEW.VersionTimeNano, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS chat_user_version_OnUpdateLogger $$
CREATE TRIGGER chat_user_version_OnUpdateLogger AFTER UPDATE ON chat_user_version
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ChatUserVersion","UPDATE",NEW.VersionTimeNano, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS chat_user_version_OnDeleteLogger $$
CREATE TRIGGER chat_user_version_OnDeleteLogger AFTER DELETE ON chat_user_version
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ChatUserVersion","DELETE",OLD.VersionTimeNano, UNIX_TIMESTAMP(NOW()));
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
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMember","INSERT",NEW.OrderId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS group_member_OnUpdateLogger $$
CREATE TRIGGER group_member_OnUpdateLogger AFTER UPDATE ON group_member
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMember","UPDATE",NEW.OrderId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS group_member_OnDeleteLogger $$
CREATE TRIGGER group_member_OnDeleteLogger AFTER DELETE ON group_member
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMember","DELETE",OLD.OrderId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ GroupOrderdUser ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS group_orderd_user_OnCreateLogger $$
CREATE TRIGGER group_orderd_user_OnCreateLogger AFTER INSERT ON group_orderd_user
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupOrderdUser","INSERT",NEW.OrderId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS group_orderd_user_OnUpdateLogger $$
CREATE TRIGGER group_orderd_user_OnUpdateLogger AFTER UPDATE ON group_orderd_user
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupOrderdUser","UPDATE",NEW.OrderId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS group_orderd_user_OnDeleteLogger $$
CREATE TRIGGER group_orderd_user_OnDeleteLogger AFTER DELETE ON group_orderd_user
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupOrderdUser","DELETE",OLD.OrderId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ GroupPinedMsg ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS group_pined_msg_OnCreateLogger $$
CREATE TRIGGER group_pined_msg_OnCreateLogger AFTER INSERT ON group_pined_msg
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupPinedMsg","INSERT",NEW.MsgId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS group_pined_msg_OnUpdateLogger $$
CREATE TRIGGER group_pined_msg_OnUpdateLogger AFTER UPDATE ON group_pined_msg
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupPinedMsg","UPDATE",NEW.MsgId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS group_pined_msg_OnDeleteLogger $$
CREATE TRIGGER group_pined_msg_OnDeleteLogger AFTER DELETE ON group_pined_msg
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupPinedMsg","DELETE",OLD.MsgId, UNIX_TIMESTAMP(NOW()));
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
################################ ActionFanout ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS action_fanout_OnCreateLogger $$
CREATE TRIGGER action_fanout_OnCreateLogger AFTER INSERT ON action_fanout
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ActionFanout","INSERT",NEW.OrderId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS action_fanout_OnUpdateLogger $$
CREATE TRIGGER action_fanout_OnUpdateLogger AFTER UPDATE ON action_fanout
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ActionFanout","UPDATE",NEW.OrderId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS action_fanout_OnDeleteLogger $$
CREATE TRIGGER action_fanout_OnDeleteLogger AFTER DELETE ON action_fanout
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("ActionFanout","DELETE",OLD.OrderId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ HomeFanout ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS home_fanout_OnCreateLogger $$
CREATE TRIGGER home_fanout_OnCreateLogger AFTER INSERT ON home_fanout
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("HomeFanout","INSERT",NEW.OrderId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS home_fanout_OnUpdateLogger $$
CREATE TRIGGER home_fanout_OnUpdateLogger AFTER UPDATE ON home_fanout
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("HomeFanout","UPDATE",NEW.OrderId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS home_fanout_OnDeleteLogger $$
CREATE TRIGGER home_fanout_OnDeleteLogger AFTER DELETE ON home_fanout
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("HomeFanout","DELETE",OLD.OrderId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ SuggestedTopPosts ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS suggested_top_posts_OnCreateLogger $$
CREATE TRIGGER suggested_top_posts_OnCreateLogger AFTER INSERT ON suggested_top_posts
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SuggestedTopPosts","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS suggested_top_posts_OnUpdateLogger $$
CREATE TRIGGER suggested_top_posts_OnUpdateLogger AFTER UPDATE ON suggested_top_posts
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SuggestedTopPosts","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS suggested_top_posts_OnDeleteLogger $$
CREATE TRIGGER suggested_top_posts_OnDeleteLogger AFTER DELETE ON suggested_top_posts
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SuggestedTopPosts","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
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
################################ PushChat ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS push_chat_OnCreateLogger $$
CREATE TRIGGER push_chat_OnCreateLogger AFTER INSERT ON push_chat
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PushChat","INSERT",NEW.PushId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS push_chat_OnUpdateLogger $$
CREATE TRIGGER push_chat_OnUpdateLogger AFTER UPDATE ON push_chat
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PushChat","UPDATE",NEW.PushId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS push_chat_OnDeleteLogger $$
CREATE TRIGGER push_chat_OnDeleteLogger AFTER DELETE ON push_chat
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PushChat","DELETE",OLD.PushId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ HTTPRPCLog ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS http_rpc_log_OnCreateLogger $$
CREATE TRIGGER http_rpc_log_OnCreateLogger AFTER INSERT ON http_rpc_log
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("HTTPRPCLog","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS http_rpc_log_OnUpdateLogger $$
CREATE TRIGGER http_rpc_log_OnUpdateLogger AFTER UPDATE ON http_rpc_log
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("HTTPRPCLog","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS http_rpc_log_OnDeleteLogger $$
CREATE TRIGGER http_rpc_log_OnDeleteLogger AFTER DELETE ON http_rpc_log
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("HTTPRPCLog","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ MetricLog ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS metric_log_OnCreateLogger $$
CREATE TRIGGER metric_log_OnCreateLogger AFTER INSERT ON metric_log
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("MetricLog","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS metric_log_OnUpdateLogger $$
CREATE TRIGGER metric_log_OnUpdateLogger AFTER UPDATE ON metric_log
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("MetricLog","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS metric_log_OnDeleteLogger $$
CREATE TRIGGER metric_log_OnDeleteLogger AFTER DELETE ON metric_log
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("MetricLog","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ XfileServiceInfoLog ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS xfile_service_info_log_OnCreateLogger $$
CREATE TRIGGER xfile_service_info_log_OnCreateLogger AFTER INSERT ON xfile_service_info_log
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("XfileServiceInfoLog","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS xfile_service_info_log_OnUpdateLogger $$
CREATE TRIGGER xfile_service_info_log_OnUpdateLogger AFTER UPDATE ON xfile_service_info_log
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("XfileServiceInfoLog","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS xfile_service_info_log_OnDeleteLogger $$
CREATE TRIGGER xfile_service_info_log_OnDeleteLogger AFTER DELETE ON xfile_service_info_log
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("XfileServiceInfoLog","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ XfileServiceMetricLog ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS xfile_service_metric_log_OnCreateLogger $$
CREATE TRIGGER xfile_service_metric_log_OnCreateLogger AFTER INSERT ON xfile_service_metric_log
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("XfileServiceMetricLog","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS xfile_service_metric_log_OnUpdateLogger $$
CREATE TRIGGER xfile_service_metric_log_OnUpdateLogger AFTER UPDATE ON xfile_service_metric_log
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("XfileServiceMetricLog","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS xfile_service_metric_log_OnDeleteLogger $$
CREATE TRIGGER xfile_service_metric_log_OnDeleteLogger AFTER DELETE ON xfile_service_metric_log
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("XfileServiceMetricLog","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ XfileServiceRequestLog ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS xfile_service_request_log_OnCreateLogger $$
CREATE TRIGGER xfile_service_request_log_OnCreateLogger AFTER INSERT ON xfile_service_request_log
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("XfileServiceRequestLog","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS xfile_service_request_log_OnUpdateLogger $$
CREATE TRIGGER xfile_service_request_log_OnUpdateLogger AFTER UPDATE ON xfile_service_request_log
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("XfileServiceRequestLog","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS xfile_service_request_log_OnDeleteLogger $$
CREATE TRIGGER xfile_service_request_log_OnDeleteLogger AFTER DELETE ON xfile_service_request_log
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("XfileServiceRequestLog","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ InvalidateCache ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS invalidate_cache_OnCreateLogger $$
CREATE TRIGGER invalidate_cache_OnCreateLogger AFTER INSERT ON invalidate_cache
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("InvalidateCache","INSERT",NEW.OrderId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS invalidate_cache_OnUpdateLogger $$
CREATE TRIGGER invalidate_cache_OnUpdateLogger AFTER UPDATE ON invalidate_cache
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("InvalidateCache","UPDATE",NEW.OrderId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS invalidate_cache_OnDeleteLogger $$
CREATE TRIGGER invalidate_cache_OnDeleteLogger AFTER DELETE ON invalidate_cache
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("InvalidateCache","DELETE",OLD.OrderId, UNIX_TIMESTAMP(NOW()));
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
### Blocked ##
DROP TRIGGER IF EXISTS blocked_OnCreateLogger ;
DROP TRIGGER IF EXISTS blocked_OnUpdateLogger ;
DROP TRIGGER IF EXISTS blocked_OnDeleteLogger ;
### Comment ##
DROP TRIGGER IF EXISTS comment_OnCreateLogger ;
DROP TRIGGER IF EXISTS comment_OnUpdateLogger ;
DROP TRIGGER IF EXISTS comment_OnDeleteLogger ;
### CommentDeleted ##
DROP TRIGGER IF EXISTS comment_deleted_OnCreateLogger ;
DROP TRIGGER IF EXISTS comment_deleted_OnUpdateLogger ;
DROP TRIGGER IF EXISTS comment_deleted_OnDeleteLogger ;
### Event ##
DROP TRIGGER IF EXISTS event_OnCreateLogger ;
DROP TRIGGER IF EXISTS event_OnUpdateLogger ;
DROP TRIGGER IF EXISTS event_OnDeleteLogger ;
### Followed ##
DROP TRIGGER IF EXISTS followed_OnCreateLogger ;
DROP TRIGGER IF EXISTS followed_OnUpdateLogger ;
DROP TRIGGER IF EXISTS followed_OnDeleteLogger ;
### Likes ##
DROP TRIGGER IF EXISTS likes_OnCreateLogger ;
DROP TRIGGER IF EXISTS likes_OnUpdateLogger ;
DROP TRIGGER IF EXISTS likes_OnDeleteLogger ;
### Notify ##
DROP TRIGGER IF EXISTS notify_OnCreateLogger ;
DROP TRIGGER IF EXISTS notify_OnUpdateLogger ;
DROP TRIGGER IF EXISTS notify_OnDeleteLogger ;
### NotifyRemoved ##
DROP TRIGGER IF EXISTS notify_removed_OnCreateLogger ;
DROP TRIGGER IF EXISTS notify_removed_OnUpdateLogger ;
DROP TRIGGER IF EXISTS notify_removed_OnDeleteLogger ;
### PhoneContacts ##
DROP TRIGGER IF EXISTS phone_contacts_OnCreateLogger ;
DROP TRIGGER IF EXISTS phone_contacts_OnUpdateLogger ;
DROP TRIGGER IF EXISTS phone_contacts_OnDeleteLogger ;
### Post ##
DROP TRIGGER IF EXISTS post_OnCreateLogger ;
DROP TRIGGER IF EXISTS post_OnUpdateLogger ;
DROP TRIGGER IF EXISTS post_OnDeleteLogger ;
### PostCount ##
DROP TRIGGER IF EXISTS post_count_OnCreateLogger ;
DROP TRIGGER IF EXISTS post_count_OnUpdateLogger ;
DROP TRIGGER IF EXISTS post_count_OnDeleteLogger ;
### PostDeleted ##
DROP TRIGGER IF EXISTS post_deleted_OnCreateLogger ;
DROP TRIGGER IF EXISTS post_deleted_OnUpdateLogger ;
DROP TRIGGER IF EXISTS post_deleted_OnDeleteLogger ;
### PostKeys ##
DROP TRIGGER IF EXISTS post_keys_OnCreateLogger ;
DROP TRIGGER IF EXISTS post_keys_OnUpdateLogger ;
DROP TRIGGER IF EXISTS post_keys_OnDeleteLogger ;
### PostLink ##
DROP TRIGGER IF EXISTS post_link_OnCreateLogger ;
DROP TRIGGER IF EXISTS post_link_OnUpdateLogger ;
DROP TRIGGER IF EXISTS post_link_OnDeleteLogger ;
### PostMedia ##
DROP TRIGGER IF EXISTS post_media_OnCreateLogger ;
DROP TRIGGER IF EXISTS post_media_OnUpdateLogger ;
DROP TRIGGER IF EXISTS post_media_OnDeleteLogger ;
### PostPromoted ##
DROP TRIGGER IF EXISTS post_promoted_OnCreateLogger ;
DROP TRIGGER IF EXISTS post_promoted_OnUpdateLogger ;
DROP TRIGGER IF EXISTS post_promoted_OnDeleteLogger ;
### PostReshared ##
DROP TRIGGER IF EXISTS post_reshared_OnCreateLogger ;
DROP TRIGGER IF EXISTS post_reshared_OnUpdateLogger ;
DROP TRIGGER IF EXISTS post_reshared_OnDeleteLogger ;
### ProfileAll ##
DROP TRIGGER IF EXISTS profile_all_OnCreateLogger ;
DROP TRIGGER IF EXISTS profile_all_OnUpdateLogger ;
DROP TRIGGER IF EXISTS profile_all_OnDeleteLogger ;
### ProfileMedia ##
DROP TRIGGER IF EXISTS profile_media_OnCreateLogger ;
DROP TRIGGER IF EXISTS profile_media_OnUpdateLogger ;
DROP TRIGGER IF EXISTS profile_media_OnDeleteLogger ;
### ProfileMentioned ##
DROP TRIGGER IF EXISTS profile_mentioned_OnCreateLogger ;
DROP TRIGGER IF EXISTS profile_mentioned_OnUpdateLogger ;
DROP TRIGGER IF EXISTS profile_mentioned_OnDeleteLogger ;
### Session ##
DROP TRIGGER IF EXISTS session_OnCreateLogger ;
DROP TRIGGER IF EXISTS session_OnUpdateLogger ;
DROP TRIGGER IF EXISTS session_OnDeleteLogger ;
### SettingNotifications ##
DROP TRIGGER IF EXISTS setting_notifications_OnCreateLogger ;
DROP TRIGGER IF EXISTS setting_notifications_OnUpdateLogger ;
DROP TRIGGER IF EXISTS setting_notifications_OnDeleteLogger ;
### Sms ##
DROP TRIGGER IF EXISTS sms_OnCreateLogger ;
DROP TRIGGER IF EXISTS sms_OnUpdateLogger ;
DROP TRIGGER IF EXISTS sms_OnDeleteLogger ;
### Tag ##
DROP TRIGGER IF EXISTS tag_OnCreateLogger ;
DROP TRIGGER IF EXISTS tag_OnUpdateLogger ;
DROP TRIGGER IF EXISTS tag_OnDeleteLogger ;
### TagPost ##
DROP TRIGGER IF EXISTS tag_post_OnCreateLogger ;
DROP TRIGGER IF EXISTS tag_post_OnUpdateLogger ;
DROP TRIGGER IF EXISTS tag_post_OnDeleteLogger ;
### TriggerLog ##
DROP TRIGGER IF EXISTS trigger_log_OnCreateLogger ;
DROP TRIGGER IF EXISTS trigger_log_OnUpdateLogger ;
DROP TRIGGER IF EXISTS trigger_log_OnDeleteLogger ;
### User ##
DROP TRIGGER IF EXISTS user_OnCreateLogger ;
DROP TRIGGER IF EXISTS user_OnUpdateLogger ;
DROP TRIGGER IF EXISTS user_OnDeleteLogger ;
### UserRelation ##
DROP TRIGGER IF EXISTS user_relation_OnCreateLogger ;
DROP TRIGGER IF EXISTS user_relation_OnUpdateLogger ;
DROP TRIGGER IF EXISTS user_relation_OnDeleteLogger ;
### Chat ##
DROP TRIGGER IF EXISTS chat_OnCreateLogger ;
DROP TRIGGER IF EXISTS chat_OnUpdateLogger ;
DROP TRIGGER IF EXISTS chat_OnDeleteLogger ;
### ChatDeleted ##
DROP TRIGGER IF EXISTS chat_deleted_OnCreateLogger ;
DROP TRIGGER IF EXISTS chat_deleted_OnUpdateLogger ;
DROP TRIGGER IF EXISTS chat_deleted_OnDeleteLogger ;
### ChatLastMessage ##
DROP TRIGGER IF EXISTS chat_last_message_OnCreateLogger ;
DROP TRIGGER IF EXISTS chat_last_message_OnUpdateLogger ;
DROP TRIGGER IF EXISTS chat_last_message_OnDeleteLogger ;
### ChatUserVersion ##
DROP TRIGGER IF EXISTS chat_user_version_OnCreateLogger ;
DROP TRIGGER IF EXISTS chat_user_version_OnUpdateLogger ;
DROP TRIGGER IF EXISTS chat_user_version_OnDeleteLogger ;
### Group ##
DROP TRIGGER IF EXISTS group_OnCreateLogger ;
DROP TRIGGER IF EXISTS group_OnUpdateLogger ;
DROP TRIGGER IF EXISTS group_OnDeleteLogger ;
### GroupMember ##
DROP TRIGGER IF EXISTS group_member_OnCreateLogger ;
DROP TRIGGER IF EXISTS group_member_OnUpdateLogger ;
DROP TRIGGER IF EXISTS group_member_OnDeleteLogger ;
### GroupOrderdUser ##
DROP TRIGGER IF EXISTS group_orderd_user_OnCreateLogger ;
DROP TRIGGER IF EXISTS group_orderd_user_OnUpdateLogger ;
DROP TRIGGER IF EXISTS group_orderd_user_OnDeleteLogger ;
### GroupPinedMsg ##
DROP TRIGGER IF EXISTS group_pined_msg_OnCreateLogger ;
DROP TRIGGER IF EXISTS group_pined_msg_OnUpdateLogger ;
DROP TRIGGER IF EXISTS group_pined_msg_OnDeleteLogger ;
### FileMsg ##
DROP TRIGGER IF EXISTS file_msg_OnCreateLogger ;
DROP TRIGGER IF EXISTS file_msg_OnUpdateLogger ;
DROP TRIGGER IF EXISTS file_msg_OnDeleteLogger ;
### FilePost ##
DROP TRIGGER IF EXISTS file_post_OnCreateLogger ;
DROP TRIGGER IF EXISTS file_post_OnUpdateLogger ;
DROP TRIGGER IF EXISTS file_post_OnDeleteLogger ;
### ActionFanout ##
DROP TRIGGER IF EXISTS action_fanout_OnCreateLogger ;
DROP TRIGGER IF EXISTS action_fanout_OnUpdateLogger ;
DROP TRIGGER IF EXISTS action_fanout_OnDeleteLogger ;
### HomeFanout ##
DROP TRIGGER IF EXISTS home_fanout_OnCreateLogger ;
DROP TRIGGER IF EXISTS home_fanout_OnUpdateLogger ;
DROP TRIGGER IF EXISTS home_fanout_OnDeleteLogger ;
### SuggestedTopPosts ##
DROP TRIGGER IF EXISTS suggested_top_posts_OnCreateLogger ;
DROP TRIGGER IF EXISTS suggested_top_posts_OnUpdateLogger ;
DROP TRIGGER IF EXISTS suggested_top_posts_OnDeleteLogger ;
### SuggestedUser ##
DROP TRIGGER IF EXISTS suggested_user_OnCreateLogger ;
DROP TRIGGER IF EXISTS suggested_user_OnUpdateLogger ;
DROP TRIGGER IF EXISTS suggested_user_OnDeleteLogger ;
### PushChat ##
DROP TRIGGER IF EXISTS push_chat_OnCreateLogger ;
DROP TRIGGER IF EXISTS push_chat_OnUpdateLogger ;
DROP TRIGGER IF EXISTS push_chat_OnDeleteLogger ;
### HTTPRPCLog ##
DROP TRIGGER IF EXISTS http_rpc_log_OnCreateLogger ;
DROP TRIGGER IF EXISTS http_rpc_log_OnUpdateLogger ;
DROP TRIGGER IF EXISTS http_rpc_log_OnDeleteLogger ;
### MetricLog ##
DROP TRIGGER IF EXISTS metric_log_OnCreateLogger ;
DROP TRIGGER IF EXISTS metric_log_OnUpdateLogger ;
DROP TRIGGER IF EXISTS metric_log_OnDeleteLogger ;
### XfileServiceInfoLog ##
DROP TRIGGER IF EXISTS xfile_service_info_log_OnCreateLogger ;
DROP TRIGGER IF EXISTS xfile_service_info_log_OnUpdateLogger ;
DROP TRIGGER IF EXISTS xfile_service_info_log_OnDeleteLogger ;
### XfileServiceMetricLog ##
DROP TRIGGER IF EXISTS xfile_service_metric_log_OnCreateLogger ;
DROP TRIGGER IF EXISTS xfile_service_metric_log_OnUpdateLogger ;
DROP TRIGGER IF EXISTS xfile_service_metric_log_OnDeleteLogger ;
### XfileServiceRequestLog ##
DROP TRIGGER IF EXISTS xfile_service_request_log_OnCreateLogger ;
DROP TRIGGER IF EXISTS xfile_service_request_log_OnUpdateLogger ;
DROP TRIGGER IF EXISTS xfile_service_request_log_OnDeleteLogger ;
### InvalidateCache ##
DROP TRIGGER IF EXISTS invalidate_cache_OnCreateLogger ;
DROP TRIGGER IF EXISTS invalidate_cache_OnUpdateLogger ;
DROP TRIGGER IF EXISTS invalidate_cache_OnDeleteLogger ;
*/