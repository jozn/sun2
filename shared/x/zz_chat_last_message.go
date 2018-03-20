package x

// GENERATED BY XO. DO NOT EDIT.
import (
	"errors"
	"strings"
	//"time"
	"ms/sun/helper"
	"strconv"

	"github.com/jmoiron/sqlx"
) // (shortname .TableNameGo "err" "res" "sqlstr" "db" "XOLog") -}}//(schema .Schema .Table.TableName) -}}// .TableNameGo}}// ChatLastMessage represents a row from 'sun_chat.chat_last_message'.

// Manualy copy this to project
type ChatLastMessage__ struct {
	ChatKey     string `json:"ChatKey"`     // ChatKey -
	ForUserId   int    `json:"ForUserId"`   // ForUserId -
	LastMsgPb   []byte `json:"LastMsgPb"`   // LastMsgPb -
	LastMsgJson string `json:"LastMsgJson"` // LastMsgJson -
	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ChatLastMessage exists in the database.
func (clm *ChatLastMessage) Exists() bool {
	return clm._exists
}

// Deleted provides information if the ChatLastMessage has been deleted from the database.
func (clm *ChatLastMessage) Deleted() bool {
	return clm._deleted
}

// Insert inserts the ChatLastMessage to the database.
func (clm *ChatLastMessage) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if clm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO sun_chat.chat_last_message (` +
		`ChatKey, ForUserId, LastMsgPb, LastMsgJson` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, clm.ChatKey, clm.ForUserId, clm.LastMsgPb, clm.LastMsgJson)
	}
	_, err = db.Exec(sqlstr, clm.ChatKey, clm.ForUserId, clm.LastMsgPb, clm.LastMsgJson)
	if err != nil {
		return err
	}

	// set existence
	clm._exists = true

	OnChatLastMessage_AfterInsert(clm)

	return nil
}

// Insert inserts the ChatLastMessage to the database.
func (clm *ChatLastMessage) Replace(db XODB) error {
	var err error

	// sql query

	const sqlstr = `REPLACE INTO sun_chat.chat_last_message (` +
		`ChatKey, ForUserId, LastMsgPb, LastMsgJson` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, clm.ChatKey, clm.ForUserId, clm.LastMsgPb, clm.LastMsgJson)
	}
	_, err = db.Exec(sqlstr, clm.ChatKey, clm.ForUserId, clm.LastMsgPb, clm.LastMsgJson)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return err
	}

	clm._exists = true

	OnChatLastMessage_AfterInsert(clm)

	return nil
}

// Update updates the ChatLastMessage in the database.
func (clm *ChatLastMessage) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !clm._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if clm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE sun_chat.chat_last_message SET ` +
		`ForUserId = ?, LastMsgPb = ?, LastMsgJson = ?` +
		` WHERE ChatKey = ?`

	// run query
	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, clm.ForUserId, clm.LastMsgPb, clm.LastMsgJson, clm.ChatKey)
	}
	_, err = db.Exec(sqlstr, clm.ForUserId, clm.LastMsgPb, clm.LastMsgJson, clm.ChatKey)

	if LogTableSqlReq.ChatLastMessage {
		XOLogErr(err)
	}
	OnChatLastMessage_AfterUpdate(clm)

	return err
}

// Save saves the ChatLastMessage to the database.
func (clm *ChatLastMessage) Save(db XODB) error {
	if clm.Exists() {
		return clm.Update(db)
	}

	return clm.Replace(db)
}

// Delete deletes the ChatLastMessage from the database.
func (clm *ChatLastMessage) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !clm._exists {
		return nil
	}

	// if deleted, bail
	if clm._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM sun_chat.chat_last_message WHERE ChatKey = ?`

	// run query
	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, clm.ChatKey)
	}
	_, err = db.Exec(sqlstr, clm.ChatKey)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return err
	}

	// set deleted
	clm._deleted = true

	OnChatLastMessage_AfterDelete(clm)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////// Querify gen - ME /////////////////////////////////////////
//.TableNameGo= table name
// _Deleter, _Updater

// orma types
type __ChatLastMessage_Deleter struct {
	wheres   []whereClause
	whereSep string
}

type __ChatLastMessage_Updater struct {
	wheres   []whereClause
	updates  map[string]interface{}
	whereSep string
}

type __ChatLastMessage_Selector struct {
	wheres    []whereClause
	selectCol string
	whereSep  string
	orderBy   string //" order by id desc //for ints
	limit     int
	offset    int
}

func NewChatLastMessage_Deleter() *__ChatLastMessage_Deleter {
	d := __ChatLastMessage_Deleter{whereSep: " AND "}
	return &d
}

func NewChatLastMessage_Updater() *__ChatLastMessage_Updater {
	u := __ChatLastMessage_Updater{whereSep: " AND "}
	u.updates = make(map[string]interface{}, 10)
	return &u
}

func NewChatLastMessage_Selector() *__ChatLastMessage_Selector {
	u := __ChatLastMessage_Selector{whereSep: " AND ", selectCol: "*"}
	return &u
}

/////////////////////////////// Where for all /////////////////////////////
//// for ints all selector updater, deleter

////////ints
func (u *__ChatLastMessage_Deleter) Or() *__ChatLastMessage_Deleter {
	u.whereSep = " OR "
	return u
}

func (u *__ChatLastMessage_Deleter) ForUserId_In(ins []int) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ForUserId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Deleter) ForUserId_Ins(ins ...int) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ForUserId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Deleter) ForUserId_NotIn(ins []int) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ForUserId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__ChatLastMessage_Deleter) ForUserId_Eq(val int) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Deleter) ForUserId_NotEq(val int) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Deleter) ForUserId_LT(val int) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Deleter) ForUserId_LE(val int) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Deleter) ForUserId_GT(val int) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Deleter) ForUserId_GE(val int) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__ChatLastMessage_Updater) Or() *__ChatLastMessage_Updater {
	u.whereSep = " OR "
	return u
}

func (u *__ChatLastMessage_Updater) ForUserId_In(ins []int) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ForUserId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Updater) ForUserId_Ins(ins ...int) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ForUserId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Updater) ForUserId_NotIn(ins []int) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ForUserId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__ChatLastMessage_Updater) ForUserId_Eq(val int) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Updater) ForUserId_NotEq(val int) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Updater) ForUserId_LT(val int) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Updater) ForUserId_LE(val int) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Updater) ForUserId_GT(val int) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Updater) ForUserId_GE(val int) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__ChatLastMessage_Selector) Or() *__ChatLastMessage_Selector {
	u.whereSep = " OR "
	return u
}

func (u *__ChatLastMessage_Selector) ForUserId_In(ins []int) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ForUserId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Selector) ForUserId_Ins(ins ...int) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ForUserId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Selector) ForUserId_NotIn(ins []int) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ForUserId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__ChatLastMessage_Selector) ForUserId_Eq(val int) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Selector) ForUserId_NotEq(val int) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Selector) ForUserId_LT(val int) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Selector) ForUserId_LE(val int) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Selector) ForUserId_GT(val int) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Selector) ForUserId_GE(val int) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

///// for strings //copy of above with type int -> string + rm if eq + $ms_str_cond

////////ints

func (u *__ChatLastMessage_Deleter) ChatKey_In(ins []string) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ChatKey IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Deleter) ChatKey_NotIn(ins []string) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ChatKey NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__ChatLastMessage_Deleter) ChatKey_Like(val string) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ChatKey LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__ChatLastMessage_Deleter) ChatKey_Eq(val string) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ChatKey = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Deleter) ChatKey_NotEq(val string) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ChatKey != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__ChatLastMessage_Deleter) LastMsgJson_In(ins []string) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LastMsgJson IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Deleter) LastMsgJson_NotIn(ins []string) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LastMsgJson NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__ChatLastMessage_Deleter) LastMsgJson_Like(val string) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LastMsgJson LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__ChatLastMessage_Deleter) LastMsgJson_Eq(val string) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LastMsgJson = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Deleter) LastMsgJson_NotEq(val string) *__ChatLastMessage_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LastMsgJson != ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints

func (u *__ChatLastMessage_Updater) ChatKey_In(ins []string) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ChatKey IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Updater) ChatKey_NotIn(ins []string) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ChatKey NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__ChatLastMessage_Updater) ChatKey_Like(val string) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ChatKey LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__ChatLastMessage_Updater) ChatKey_Eq(val string) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ChatKey = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Updater) ChatKey_NotEq(val string) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ChatKey != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__ChatLastMessage_Updater) LastMsgJson_In(ins []string) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LastMsgJson IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Updater) LastMsgJson_NotIn(ins []string) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LastMsgJson NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__ChatLastMessage_Updater) LastMsgJson_Like(val string) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LastMsgJson LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__ChatLastMessage_Updater) LastMsgJson_Eq(val string) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LastMsgJson = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Updater) LastMsgJson_NotEq(val string) *__ChatLastMessage_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LastMsgJson != ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints

func (u *__ChatLastMessage_Selector) ChatKey_In(ins []string) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ChatKey IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Selector) ChatKey_NotIn(ins []string) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ChatKey NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__ChatLastMessage_Selector) ChatKey_Like(val string) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ChatKey LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__ChatLastMessage_Selector) ChatKey_Eq(val string) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ChatKey = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Selector) ChatKey_NotEq(val string) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ChatKey != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__ChatLastMessage_Selector) LastMsgJson_In(ins []string) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LastMsgJson IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__ChatLastMessage_Selector) LastMsgJson_NotIn(ins []string) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LastMsgJson NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__ChatLastMessage_Selector) LastMsgJson_Like(val string) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LastMsgJson LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__ChatLastMessage_Selector) LastMsgJson_Eq(val string) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LastMsgJson = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__ChatLastMessage_Selector) LastMsgJson_NotEq(val string) *__ChatLastMessage_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LastMsgJson != ? "
	d.wheres = append(d.wheres, w)

	return d
}

/// End of wheres for selectors , updators, deletor

/////////////////////////////// Updater /////////////////////////////

//ints

//string
func (u *__ChatLastMessage_Updater) ChatKey(newVal string) *__ChatLastMessage_Updater {
	u.updates[" ChatKey = ? "] = newVal
	return u
}

//ints

func (u *__ChatLastMessage_Updater) ForUserId(newVal int) *__ChatLastMessage_Updater {
	u.updates[" ForUserId = ? "] = newVal
	return u
}

func (u *__ChatLastMessage_Updater) ForUserId_Increment(count int) *__ChatLastMessage_Updater {
	if count > 0 {
		u.updates[" ForUserId = ForUserId+? "] = count
	}

	if count < 0 {
		u.updates[" ForUserId = ForUserId-? "] = -(count) //make it positive
	}

	return u
}

//string

//ints

//string

//ints

//string
func (u *__ChatLastMessage_Updater) LastMsgJson(newVal string) *__ChatLastMessage_Updater {
	u.updates[" LastMsgJson = ? "] = newVal
	return u
}

/////////////////////////////////////////////////////////////////////
/////////////////////// Selector ///////////////////////////////////

//Select_* can just be used with: .GetString() , .GetStringSlice(), .GetInt() ..GetIntSlice()

func (u *__ChatLastMessage_Selector) OrderBy_ChatKey_Desc() *__ChatLastMessage_Selector {
	u.orderBy = " ORDER BY ChatKey DESC "
	return u
}

func (u *__ChatLastMessage_Selector) OrderBy_ChatKey_Asc() *__ChatLastMessage_Selector {
	u.orderBy = " ORDER BY ChatKey ASC "
	return u
}

func (u *__ChatLastMessage_Selector) Select_ChatKey() *__ChatLastMessage_Selector {
	u.selectCol = "ChatKey"
	return u
}

func (u *__ChatLastMessage_Selector) OrderBy_ForUserId_Desc() *__ChatLastMessage_Selector {
	u.orderBy = " ORDER BY ForUserId DESC "
	return u
}

func (u *__ChatLastMessage_Selector) OrderBy_ForUserId_Asc() *__ChatLastMessage_Selector {
	u.orderBy = " ORDER BY ForUserId ASC "
	return u
}

func (u *__ChatLastMessage_Selector) Select_ForUserId() *__ChatLastMessage_Selector {
	u.selectCol = "ForUserId"
	return u
}

func (u *__ChatLastMessage_Selector) OrderBy_LastMsgPb_Desc() *__ChatLastMessage_Selector {
	u.orderBy = " ORDER BY LastMsgPb DESC "
	return u
}

func (u *__ChatLastMessage_Selector) OrderBy_LastMsgPb_Asc() *__ChatLastMessage_Selector {
	u.orderBy = " ORDER BY LastMsgPb ASC "
	return u
}

func (u *__ChatLastMessage_Selector) Select_LastMsgPb() *__ChatLastMessage_Selector {
	u.selectCol = "LastMsgPb"
	return u
}

func (u *__ChatLastMessage_Selector) OrderBy_LastMsgJson_Desc() *__ChatLastMessage_Selector {
	u.orderBy = " ORDER BY LastMsgJson DESC "
	return u
}

func (u *__ChatLastMessage_Selector) OrderBy_LastMsgJson_Asc() *__ChatLastMessage_Selector {
	u.orderBy = " ORDER BY LastMsgJson ASC "
	return u
}

func (u *__ChatLastMessage_Selector) Select_LastMsgJson() *__ChatLastMessage_Selector {
	u.selectCol = "LastMsgJson"
	return u
}

func (u *__ChatLastMessage_Selector) Limit(num int) *__ChatLastMessage_Selector {
	u.limit = num
	return u
}

func (u *__ChatLastMessage_Selector) Offset(num int) *__ChatLastMessage_Selector {
	u.offset = num
	return u
}

func (u *__ChatLastMessage_Selector) Order_Rand() *__ChatLastMessage_Selector {
	u.orderBy = " ORDER BY RAND() "
	return u
}

/////////////////////////  Queryer Selector  //////////////////////////////////
func (u *__ChatLastMessage_Selector) _stoSql() (string, []interface{}) {
	sqlWherrs, whereArgs := whereClusesToSql(u.wheres, u.whereSep)

	sqlstr := "SELECT " + u.selectCol + " FROM sun_chat.chat_last_message"

	if len(strings.Trim(sqlWherrs, " ")) > 0 { //2 for safty
		sqlstr += " WHERE " + sqlWherrs
	}

	if u.orderBy != "" {
		sqlstr += u.orderBy
	}

	if u.limit != 0 {
		sqlstr += " LIMIT " + strconv.Itoa(u.limit)
	}

	if u.offset != 0 {
		sqlstr += " OFFSET " + strconv.Itoa(u.offset)
	}
	return sqlstr, whereArgs
}

func (u *__ChatLastMessage_Selector) GetRow(db *sqlx.DB) (*ChatLastMessage, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, whereArgs)
	}

	row := &ChatLastMessage{}
	//by Sqlx
	err = db.Get(row, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return nil, err
	}

	row._exists = true

	OnChatLastMessage_LoadOne(row)

	return row, nil
}

func (u *__ChatLastMessage_Selector) GetRows(db *sqlx.DB) ([]*ChatLastMessage, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, whereArgs)
	}

	var rows []*ChatLastMessage
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return nil, err
	}

	/*for i:=0;i< len(rows);i++ {
		rows[i]._exists = true
	}*/

	for i := 0; i < len(rows); i++ {
		rows[i]._exists = true
	}

	OnChatLastMessage_LoadMany(rows)

	return rows, nil
}

//dep use GetRows()
func (u *__ChatLastMessage_Selector) GetRows2(db *sqlx.DB) ([]ChatLastMessage, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, whereArgs)
	}
	var rows []*ChatLastMessage
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return nil, err
	}

	/*for i:=0;i< len(rows);i++ {
		rows[i]._exists = true
	}*/

	for i := 0; i < len(rows); i++ {
		rows[i]._exists = true
	}

	OnChatLastMessage_LoadMany(rows)

	rows2 := make([]ChatLastMessage, len(rows))
	for i := 0; i < len(rows); i++ {
		cp := *rows[i]
		rows2[i] = cp
	}

	return rows2, nil
}

func (u *__ChatLastMessage_Selector) GetString(db *sqlx.DB) (string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, whereArgs)
	}

	var res string
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return "", err
	}

	return res, nil
}

func (u *__ChatLastMessage_Selector) GetStringSlice(db *sqlx.DB) ([]string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, whereArgs)
	}
	var rows []string
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__ChatLastMessage_Selector) GetIntSlice(db *sqlx.DB) ([]int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, whereArgs)
	}
	var rows []int
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__ChatLastMessage_Selector) GetInt(db *sqlx.DB) (int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, whereArgs)
	}
	var res int
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return 0, err
	}

	return res, nil
}

/////////////////////////  Queryer Update Delete //////////////////////////////////
func (u *__ChatLastMessage_Updater) Update(db XODB) (int, error) {
	var err error

	var updateArgs []interface{}
	var sqlUpdateArr []string
	for up, newVal := range u.updates {
		sqlUpdateArr = append(sqlUpdateArr, up)
		updateArgs = append(updateArgs, newVal)
	}
	sqlUpdate := strings.Join(sqlUpdateArr, ",")

	sqlWherrs, whereArgs := whereClusesToSql(u.wheres, u.whereSep)

	var allArgs []interface{}
	allArgs = append(allArgs, updateArgs...)
	allArgs = append(allArgs, whereArgs...)

	sqlstr := `UPDATE sun_chat.chat_last_message SET ` + sqlUpdate

	if len(strings.Trim(sqlWherrs, " ")) > 0 { //2 for safty
		sqlstr += " WHERE " + sqlWherrs
	}

	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, allArgs)
	}
	res, err := db.Exec(sqlstr, allArgs...)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return 0, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

func (d *__ChatLastMessage_Deleter) Delete(db XODB) (int, error) {
	var err error
	var wheresArr []string
	for _, w := range d.wheres {
		wheresArr = append(wheresArr, w.condition)
	}
	wheresStr := strings.Join(wheresArr, d.whereSep)

	var args []interface{}
	for _, w := range d.wheres {
		args = append(args, w.args...)
	}

	sqlstr := "DELETE FROM sun_chat.chat_last_message WHERE " + wheresStr

	// run query
	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, args)
	}
	res, err := db.Exec(sqlstr, args...)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return 0, err
	}

	// retrieve id
	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

///////////////////////// Mass insert - replace for  ChatLastMessage ////////////////

func MassInsert_ChatLastMessage(rows []ChatLastMessage, db XODB) error {
	if len(rows) == 0 {
		return errors.New("rows slice should not be empty - inserted nothing")
	}
	var err error
	ln := len(rows)
	//s:= "(?,?,?,?)," //`(?, ?, ?, ?),`
	s := "(?,?,?,?)," //`(?, ?, ?, ?),`
	insVals_ := strings.Repeat(s, ln)
	insVals := insVals_[0 : len(insVals_)-1]
	// sql query
	sqlstr := "INSERT INTO sun_chat.chat_last_message (" +
		"ChatKey, ForUserId, LastMsgPb, LastMsgJson" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.ChatKey)
		vals = append(vals, row.ForUserId)
		vals = append(vals, row.LastMsgPb)
		vals = append(vals, row.LastMsgJson)

	}

	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, " MassInsert len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return err
	}

	return nil
}

func MassReplace_ChatLastMessage(rows []ChatLastMessage, db XODB) error {
	if len(rows) == 0 {
		return errors.New("rows slice should not be empty - inserted nothing")
	}
	var err error
	ln := len(rows)
	//s:= "(?,?,?,?)," //`(?, ?, ?, ?),`
	s := "(?,?,?,?)," //`(?, ?, ?, ?),`
	insVals_ := strings.Repeat(s, ln)
	insVals := insVals_[0 : len(insVals_)-1]
	// sql query
	sqlstr := "REPLACE INTO sun_chat.chat_last_message (" +
		"ChatKey, ForUserId, LastMsgPb, LastMsgJson" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.ChatKey)
		vals = append(vals, row.ForUserId)
		vals = append(vals, row.LastMsgPb)
		vals = append(vals, row.LastMsgJson)

	}

	if LogTableSqlReq.ChatLastMessage {
		XOLog(sqlstr, " MassReplace len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.ChatLastMessage {
			XOLogErr(err)
		}
		return err
	}

	return nil

}

//////////////////// Play ///////////////////////////////

//

//

//

//
