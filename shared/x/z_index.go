package x

import "github.com/jmoiron/sqlx"

//todo this code can be used for multi secondery coulmn -- but this one interfier with secondery_template - in futer merege this two temple to one unifed that uses both primiry keys and othe seconder options and with reloadings

// ActionByActionId Generated from index 'PRIMARY' -- retrieves a row from 'sun.action' as a Action.
func ActionByActionId(db *sqlx.DB, actionId int) (*Action, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.action ` +
		`WHERE ActionId = ?`

	XOLog(sqlstr, actionId)
	a := Action{
		_exists: true,
	}

	err = db.Get(&a, sqlstr, actionId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnAction_LoadOne(&a)

	return &a, nil
}

// CommentByCommentId Generated from index 'PRIMARY' -- retrieves a row from 'sun.comment' as a Comment.
func CommentByCommentId(db *sqlx.DB, commentId int) (*Comment, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.comment ` +
		`WHERE CommentId = ?`

	XOLog(sqlstr, commentId)
	c := Comment{
		_exists: true,
	}

	err = db.Get(&c, sqlstr, commentId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnComment_LoadOne(&c)

	return &c, nil
}

// FollowingListById Generated from index 'PRIMARY' -- retrieves a row from 'sun.following_list' as a FollowingList.
func FollowingListById(db *sqlx.DB, id int) (*FollowingList, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.following_list ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	fl := FollowingList{
		_exists: true,
	}

	err = db.Get(&fl, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFollowingList_LoadOne(&fl)

	return &fl, nil
}

// FollowingListMemberById Generated from index 'PRIMARY' -- retrieves a row from 'sun.following_list_member' as a FollowingListMember.
func FollowingListMemberById(db *sqlx.DB, id int) (*FollowingListMember, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.following_list_member ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	flm := FollowingListMember{
		_exists: true,
	}

	err = db.Get(&flm, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFollowingListMember_LoadOne(&flm)

	return &flm, nil
}

// FollowingListMemberRemovedById Generated from index 'PRIMARY' -- retrieves a row from 'sun.following_list_member_removed' as a FollowingListMemberRemoved.
func FollowingListMemberRemovedById(db *sqlx.DB, id int) (*FollowingListMemberRemoved, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.following_list_member_removed ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	flmr := FollowingListMemberRemoved{
		_exists: true,
	}

	err = db.Get(&flmr, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFollowingListMemberRemoved_LoadOne(&flmr)

	return &flmr, nil
}

// GroupByGroupId Generated from index 'PRIMARY' -- retrieves a row from 'sun.group' as a Group.
func GroupByGroupId(db *sqlx.DB, groupId int) (*Group, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.group ` +
		`WHERE GroupId = ?`

	XOLog(sqlstr, groupId)
	g := Group{
		_exists: true,
	}

	err = db.Get(&g, sqlstr, groupId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroup_LoadOne(&g)

	return &g, nil
}

// GroupMemberById Generated from index 'PRIMARY' -- retrieves a row from 'sun.group_member' as a GroupMember.
func GroupMemberById(db *sqlx.DB, id int) (*GroupMember, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.group_member ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	gm := GroupMember{
		_exists: true,
	}

	err = db.Get(&gm, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroupMember_LoadOne(&gm)

	return &gm, nil
}

// GroupMessageByMessageId Generated from index 'PRIMARY' -- retrieves a row from 'sun.group_message' as a GroupMessage.
func GroupMessageByMessageId(db *sqlx.DB, messageId int) (*GroupMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.group_message ` +
		`WHERE MessageId = ?`

	XOLog(sqlstr, messageId)
	gm := GroupMessage{
		_exists: true,
	}

	err = db.Get(&gm, sqlstr, messageId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroupMessage_LoadOne(&gm)

	return &gm, nil
}

// LikeById Generated from index 'PRIMARY' -- retrieves a row from 'sun.likes' as a Like.
func LikeById(db *sqlx.DB, id int) (*Like, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.likes ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	l := Like{
		_exists: true,
	}

	err = db.Get(&l, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnLike_LoadOne(&l)

	return &l, nil
}

// MediaByMediaId Generated from index 'PRIMARY' -- retrieves a row from 'sun.media' as a Media.
func MediaByMediaId(db *sqlx.DB, mediaId int) (*Media, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.media ` +
		`WHERE MediaId = ?`

	XOLog(sqlstr, mediaId)
	m := Media{
		_exists: true,
	}

	err = db.Get(&m, sqlstr, mediaId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnMedia_LoadOne(&m)

	return &m, nil
}

// NotifyByNotifyId Generated from index 'PRIMARY' -- retrieves a row from 'sun.notify' as a Notify.
func NotifyByNotifyId(db *sqlx.DB, notifyId int) (*Notify, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.notify ` +
		`WHERE NotifyId = ?`

	XOLog(sqlstr, notifyId)
	n := Notify{
		_exists: true,
	}

	err = db.Get(&n, sqlstr, notifyId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnNotify_LoadOne(&n)

	return &n, nil
}

// NotifyRemovedById Generated from index 'PRIMARY' -- retrieves a row from 'sun.notify_removed' as a NotifyRemoved.
func NotifyRemovedById(db *sqlx.DB, id int) (*NotifyRemoved, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.notify_removed ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	nr := NotifyRemoved{
		_exists: true,
	}

	err = db.Get(&nr, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnNotifyRemoved_LoadOne(&nr)

	return &nr, nil
}

// PhoneContactById Generated from index 'PRIMARY' -- retrieves a row from 'sun.phone_contacts' as a PhoneContact.
func PhoneContactById(db *sqlx.DB, id int) (*PhoneContact, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.phone_contacts ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	pc := PhoneContact{
		_exists: true,
	}

	err = db.Get(&pc, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPhoneContact_LoadOne(&pc)

	return &pc, nil
}

// PostByPostId Generated from index 'PRIMARY' -- retrieves a row from 'sun.post' as a Post.
func PostByPostId(db *sqlx.DB, postId int) (*Post, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.post ` +
		`WHERE PostId = ?`

	XOLog(sqlstr, postId)
	p := Post{
		_exists: true,
	}

	err = db.Get(&p, sqlstr, postId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPost_LoadOne(&p)

	return &p, nil
}

// PostKeyById Generated from index 'PRIMARY' -- retrieves a row from 'sun.post_keys' as a PostKey.
func PostKeyById(db *sqlx.DB, id int) (*PostKey, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.post_keys ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	pk := PostKey{
		_exists: true,
	}

	err = db.Get(&pk, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPostKey_LoadOne(&pk)

	return &pk, nil
}

// SearchClickedById Generated from index 'PRIMARY' -- retrieves a row from 'sun.search_clicked' as a SearchClicked.
func SearchClickedById(db *sqlx.DB, id int) (*SearchClicked, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.search_clicked ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	sc := SearchClicked{
		_exists: true,
	}

	err = db.Get(&sc, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSearchClicked_LoadOne(&sc)

	return &sc, nil
}

// SessionById Generated from index 'PRIMARY' -- retrieves a row from 'sun.session' as a Session.
func SessionById(db *sqlx.DB, id int) (*Session, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.session ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	s := Session{
		_exists: true,
	}

	err = db.Get(&s, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSession_LoadOne(&s)

	return &s, nil
}

// SettingClientByUserId Generated from index 'PRIMARY' -- retrieves a row from 'sun.setting_client' as a SettingClient.
func SettingClientByUserId(db *sqlx.DB, userId int) (*SettingClient, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.setting_client ` +
		`WHERE UserId = ?`

	XOLog(sqlstr, userId)
	sc := SettingClient{
		_exists: true,
	}

	err = db.Get(&sc, sqlstr, userId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSettingClient_LoadOne(&sc)

	return &sc, nil
}

// SettingNotificationByUserId Generated from index 'PRIMARY' -- retrieves a row from 'sun.setting_notifications' as a SettingNotification.
func SettingNotificationByUserId(db *sqlx.DB, userId int) (*SettingNotification, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.setting_notifications ` +
		`WHERE UserId = ?`

	XOLog(sqlstr, userId)
	sn := SettingNotification{
		_exists: true,
	}

	err = db.Get(&sn, sqlstr, userId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSettingNotification_LoadOne(&sn)

	return &sn, nil
}

// SuggestedTopPostById Generated from index 'PRIMARY' -- retrieves a row from 'sun.suggested_top_posts' as a SuggestedTopPost.
func SuggestedTopPostById(db *sqlx.DB, id int) (*SuggestedTopPost, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.suggested_top_posts ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	stp := SuggestedTopPost{
		_exists: true,
	}

	err = db.Get(&stp, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSuggestedTopPost_LoadOne(&stp)

	return &stp, nil
}

// SuggestedUserById Generated from index 'PRIMARY' -- retrieves a row from 'sun.suggested_user' as a SuggestedUser.
func SuggestedUserById(db *sqlx.DB, id int) (*SuggestedUser, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.suggested_user ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	su := SuggestedUser{
		_exists: true,
	}

	err = db.Get(&su, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSuggestedUser_LoadOne(&su)

	return &su, nil
}

// TagByTagId Generated from index 'PRIMARY' -- retrieves a row from 'sun.tag' as a Tag.
func TagByTagId(db *sqlx.DB, tagId int) (*Tag, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.tag ` +
		`WHERE TagId = ?`

	XOLog(sqlstr, tagId)
	t := Tag{
		_exists: true,
	}

	err = db.Get(&t, sqlstr, tagId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnTag_LoadOne(&t)

	return &t, nil
}

// TagsPostById Generated from index 'PRIMARY' -- retrieves a row from 'sun.tags_posts' as a TagsPost.
func TagsPostById(db *sqlx.DB, id int) (*TagsPost, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.tags_posts ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	tp := TagsPost{
		_exists: true,
	}

	err = db.Get(&tp, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnTagsPost_LoadOne(&tp)

	return &tp, nil
}

// TriggerLogById Generated from index 'PRIMARY' -- retrieves a row from 'sun.trigger_log' as a TriggerLog.
func TriggerLogById(db *sqlx.DB, id int) (*TriggerLog, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.trigger_log ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	tl := TriggerLog{
		_exists: true,
	}

	err = db.Get(&tl, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnTriggerLog_LoadOne(&tl)

	return &tl, nil
}

// UserByUserId Generated from index 'PRIMARY' -- retrieves a row from 'sun.user' as a User.
func UserByUserId(db *sqlx.DB, userId int) (*User, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.user ` +
		`WHERE UserId = ?`

	XOLog(sqlstr, userId)
	u := User{
		_exists: true,
	}

	err = db.Get(&u, sqlstr, userId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnUser_LoadOne(&u)

	return &u, nil
}

// UserMetaInfoById Generated from index 'PRIMARY' -- retrieves a row from 'sun.user_meta_info' as a UserMetaInfo.
func UserMetaInfoById(db *sqlx.DB, id int) (*UserMetaInfo, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.user_meta_info ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	umi := UserMetaInfo{
		_exists: true,
	}

	err = db.Get(&umi, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnUserMetaInfo_LoadOne(&umi)

	return &umi, nil
}

// UserPasswordByUserId Generated from index 'PRIMARY' -- retrieves a row from 'sun.user_password' as a UserPassword.
func UserPasswordByUserId(db *sqlx.DB, userId int) (*UserPassword, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun.user_password ` +
		`WHERE UserId = ?`

	XOLog(sqlstr, userId)
	up := UserPassword{
		_exists: true,
	}

	err = db.Get(&up, sqlstr, userId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnUserPassword_LoadOne(&up)

	return &up, nil
}

// ChatByChatKey Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.chat' as a Chat.
func ChatByChatKey(db *sqlx.DB, chatKey string) (*Chat, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.chat ` +
		`WHERE ChatKey = ?`

	XOLog(sqlstr, chatKey)
	c := Chat{
		_exists: true,
	}

	err = db.Get(&c, sqlstr, chatKey)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnChat_LoadOne(&c)

	return &c, nil
}

// ChatLastMessageByChatKey Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.chat_last_message' as a ChatLastMessage.
func ChatLastMessageByChatKey(db *sqlx.DB, chatKey string) (*ChatLastMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.chat_last_message ` +
		`WHERE ChatKey = ?`

	XOLog(sqlstr, chatKey)
	clm := ChatLastMessage{
		_exists: true,
	}

	err = db.Get(&clm, sqlstr, chatKey)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnChatLastMessage_LoadOne(&clm)

	return &clm, nil
}

// ChatSyncBySyncId Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.chat_sync' as a ChatSync.
func ChatSyncBySyncId(db *sqlx.DB, syncId int) (*ChatSync, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.chat_sync ` +
		`WHERE SyncId = ?`

	XOLog(sqlstr, syncId)
	cs := ChatSync{
		_exists: true,
	}

	err = db.Get(&cs, sqlstr, syncId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnChatSync_LoadOne(&cs)

	return &cs, nil
}

// DirectMessageByMessageId Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.direct_message' as a DirectMessage.
func DirectMessageByMessageId(db *sqlx.DB, messageId int) (*DirectMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.direct_message ` +
		`WHERE MessageId = ?`

	XOLog(sqlstr, messageId)
	dm := DirectMessage{
		_exists: true,
	}

	err = db.Get(&dm, sqlstr, messageId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnDirectMessage_LoadOne(&dm)

	return &dm, nil
}

// HomeById Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.home' as a Home.
func HomeById(db *sqlx.DB, id int) (*Home, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.home ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	h := Home{
		_exists: true,
	}

	err = db.Get(&h, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnHome_LoadOne(&h)

	return &h, nil
}

// MessageFileByMessageFileId Generated from index 'PRIMARY' -- retrieves a row from 'sun_chat.message_file' as a MessageFile.
func MessageFileByMessageFileId(db *sqlx.DB, messageFileId int) (*MessageFile, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_chat.message_file ` +
		`WHERE MessageFileId = ?`

	XOLog(sqlstr, messageFileId)
	mf := MessageFile{
		_exists: true,
	}

	err = db.Get(&mf, sqlstr, messageFileId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnMessageFile_LoadOne(&mf)

	return &mf, nil
}

// FileMsgById Generated from index 'PRIMARY' -- retrieves a row from 'sun_file.file_msg' as a FileMsg.
func FileMsgById(db *sqlx.DB, id int) (*FileMsg, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_file.file_msg ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	fm := FileMsg{
		_exists: true,
	}

	err = db.Get(&fm, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFileMsg_LoadOne(&fm)

	return &fm, nil
}

// FilePostById Generated from index 'PRIMARY' -- retrieves a row from 'sun_file.file_post' as a FilePost.
func FilePostById(db *sqlx.DB, id int) (*FilePost, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM sun_file.file_post ` +
		`WHERE Id = ?`

	//XOLog(sqlstr, id)
	fp := FilePost{
		_exists: true,
	}

	err = db.Get(&fp, sqlstr, id)
	if err != nil {
		//XOLogErr(err)
		return nil, err
	}

	//OnFilePost_LoadOne(&fp)

	return &fp, nil
}
