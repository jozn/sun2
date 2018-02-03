package x

import "github.com/jmoiron/sqlx"

//todo this code can be used for multi secondery coulmn -- but this one interfier with secondery_template - in futer merege this two temple to one unifed that uses both primiry keys and othe seconder options and with reloadings

// ActionByActionId Generated from index 'PRIMARY' -- retrieves a row from 'ms.action' as a Action.
func ActionByActionId(db *sqlx.DB, actionId int) (*Action, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.action ` +
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

// ChatByChatKey Generated from index 'PRIMARY' -- retrieves a row from 'ms.chat' as a Chat.
func ChatByChatKey(db *sqlx.DB, chatKey string) (*Chat, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.chat ` +
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

// CommentByCommentId Generated from index 'PRIMARY' -- retrieves a row from 'ms.comment' as a Comment.
func CommentByCommentId(db *sqlx.DB, commentId int) (*Comment, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.comment ` +
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

// DirectMessageByMessageId Generated from index 'PRIMARY' -- retrieves a row from 'ms.direct_message' as a DirectMessage.
func DirectMessageByMessageId(db *sqlx.DB, messageId int) (*DirectMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.direct_message ` +
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

// DirectOfflineByDirectOfflineId Generated from index 'PRIMARY' -- retrieves a row from 'ms.direct_offline' as a DirectOffline.
func DirectOfflineByDirectOfflineId(db *sqlx.DB, directOfflineId int) (*DirectOffline, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.direct_offline ` +
		`WHERE DirectOfflineId = ?`

	XOLog(sqlstr, directOfflineId)
	do := DirectOffline{
		_exists: true,
	}

	err = db.Get(&do, sqlstr, directOfflineId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnDirectOffline_LoadOne(&do)

	return &do, nil
}

// DirectToMessageById Generated from index 'PRIMARY' -- retrieves a row from 'ms.direct_to_message' as a DirectToMessage.
func DirectToMessageById(db *sqlx.DB, id int) (*DirectToMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.direct_to_message ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	dtm := DirectToMessage{
		_exists: true,
	}

	err = db.Get(&dtm, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnDirectToMessage_LoadOne(&dtm)

	return &dtm, nil
}

// FollowingListById Generated from index 'PRIMARY' -- retrieves a row from 'ms.following_list' as a FollowingList.
func FollowingListById(db *sqlx.DB, id int) (*FollowingList, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.following_list ` +
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

// FollowingListMemberById Generated from index 'PRIMARY' -- retrieves a row from 'ms.following_list_member' as a FollowingListMember.
func FollowingListMemberById(db *sqlx.DB, id int) (*FollowingListMember, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.following_list_member ` +
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

// FollowingListMemberHistoryById Generated from index 'PRIMARY' -- retrieves a row from 'ms.following_list_member_history' as a FollowingListMemberHistory.
func FollowingListMemberHistoryById(db *sqlx.DB, id int) (*FollowingListMemberHistory, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.following_list_member_history ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	flmh := FollowingListMemberHistory{
		_exists: true,
	}

	err = db.Get(&flmh, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFollowingListMemberHistory_LoadOne(&flmh)

	return &flmh, nil
}

// GeneralLogById Generated from index 'PRIMARY' -- retrieves a row from 'ms.general_log' as a GeneralLog.
func GeneralLogById(db *sqlx.DB, id int) (*GeneralLog, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.general_log ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	gl := GeneralLog{
		_exists: true,
	}

	err = db.Get(&gl, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGeneralLog_LoadOne(&gl)

	return &gl, nil
}

// GroupByGroupId Generated from index 'PRIMARY' -- retrieves a row from 'ms.group' as a Group.
func GroupByGroupId(db *sqlx.DB, groupId int) (*Group, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.group ` +
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

// GroupMemberById Generated from index 'PRIMARY' -- retrieves a row from 'ms.group_member' as a GroupMember.
func GroupMemberById(db *sqlx.DB, id int) (*GroupMember, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.group_member ` +
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

// GroupMessageByMessageId Generated from index 'PRIMARY' -- retrieves a row from 'ms.group_message' as a GroupMessage.
func GroupMessageByMessageId(db *sqlx.DB, messageId int) (*GroupMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.group_message ` +
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

// LikeById Generated from index 'PRIMARY' -- retrieves a row from 'ms.likes' as a Like.
func LikeById(db *sqlx.DB, id int) (*Like, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.likes ` +
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

// MediaByMediaId Generated from index 'PRIMARY' -- retrieves a row from 'ms.media' as a Media.
func MediaByMediaId(db *sqlx.DB, mediaId int) (*Media, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.media ` +
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

// MessageFileByMessageFileId Generated from index 'PRIMARY' -- retrieves a row from 'ms.message_file' as a MessageFile.
func MessageFileByMessageFileId(db *sqlx.DB, messageFileId int) (*MessageFile, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.message_file ` +
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

// NotifyByNotifyId Generated from index 'PRIMARY' -- retrieves a row from 'ms.notify' as a Notify.
func NotifyByNotifyId(db *sqlx.DB, notifyId int) (*Notify, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.notify ` +
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

// NotifyRemovedById Generated from index 'PRIMARY' -- retrieves a row from 'ms.notify_removed' as a NotifyRemoved.
func NotifyRemovedById(db *sqlx.DB, id int) (*NotifyRemoved, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.notify_removed ` +
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

// PhoneContactById Generated from index 'PRIMARY' -- retrieves a row from 'ms.phone_contacts' as a PhoneContact.
func PhoneContactById(db *sqlx.DB, id int) (*PhoneContact, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.phone_contacts ` +
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

// PhoneContactsCopyById Generated from index 'PRIMARY' -- retrieves a row from 'ms.phone_contacts_copy' as a PhoneContactsCopy.
func PhoneContactsCopyById(db *sqlx.DB, id int) (*PhoneContactsCopy, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.phone_contacts_copy ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	pcc := PhoneContactsCopy{
		_exists: true,
	}

	err = db.Get(&pcc, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPhoneContactsCopy_LoadOne(&pcc)

	return &pcc, nil
}

// PostByPostId Generated from index 'PRIMARY' -- retrieves a row from 'ms.post' as a Post.
func PostByPostId(db *sqlx.DB, postId int) (*Post, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.post ` +
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

// RecommendUserById Generated from index 'PRIMARY' -- retrieves a row from 'ms.recommend_user' as a RecommendUser.
func RecommendUserById(db *sqlx.DB, id int) (*RecommendUser, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.recommend_user ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	ru := RecommendUser{
		_exists: true,
	}

	err = db.Get(&ru, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnRecommendUser_LoadOne(&ru)

	return &ru, nil
}

// SearchClickedById Generated from index 'PRIMARY' -- retrieves a row from 'ms.search_clicked' as a SearchClicked.
func SearchClickedById(db *sqlx.DB, id int) (*SearchClicked, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.search_clicked ` +
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

// SessionById Generated from index 'PRIMARY' -- retrieves a row from 'ms.session' as a Session.
func SessionById(db *sqlx.DB, id int) (*Session, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.session ` +
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

// SettingClientByUserId Generated from index 'PRIMARY' -- retrieves a row from 'ms.setting_client' as a SettingClient.
func SettingClientByUserId(db *sqlx.DB, userId int) (*SettingClient, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.setting_client ` +
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

// SettingNotificationByUserId Generated from index 'PRIMARY' -- retrieves a row from 'ms.setting_notifications' as a SettingNotification.
func SettingNotificationByUserId(db *sqlx.DB, userId int) (*SettingNotification, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.setting_notifications ` +
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

// TagByTagId Generated from index 'PRIMARY' -- retrieves a row from 'ms.tag' as a Tag.
func TagByTagId(db *sqlx.DB, tagId int) (*Tag, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.tag ` +
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

// TagsPostById Generated from index 'PRIMARY' -- retrieves a row from 'ms.tags_posts' as a TagsPost.
func TagsPostById(db *sqlx.DB, id int) (*TagsPost, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.tags_posts ` +
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

// TriggerLogById Generated from index 'PRIMARY' -- retrieves a row from 'ms.trigger_log' as a TriggerLog.
func TriggerLogById(db *sqlx.DB, id int) (*TriggerLog, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.trigger_log ` +
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

// UserByUserId Generated from index 'PRIMARY' -- retrieves a row from 'ms.user' as a User.
func UserByUserId(db *sqlx.DB, userId int) (*User, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.user ` +
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

// UserMetaInfoById Generated from index 'PRIMARY' -- retrieves a row from 'ms.user_meta_info' as a UserMetaInfo.
func UserMetaInfoById(db *sqlx.DB, id int) (*UserMetaInfo, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.user_meta_info ` +
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

// UserPasswordByUserId Generated from index 'PRIMARY' -- retrieves a row from 'ms.user_password' as a UserPassword.
func UserPasswordByUserId(db *sqlx.DB, userId int) (*UserPassword, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.user_password ` +
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
