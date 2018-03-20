package x

// GENERATED BY XO. DO NOT EDIT.
import (
	"errors"
	"strings"
	//"time"
	"ms/sun/helper"
	"strconv"

	"github.com/jmoiron/sqlx"
) // (shortname .TableNameGo "err" "res" "sqlstr" "db" "XOLog") -}}//(schema .Schema .Table.TableName) -}}// .TableNameGo}}// Home represents a row from 'sun_chat.home'.

// Manualy copy this to project
type Home__ struct {
	Id        int `json:"Id"`        // Id -
	ForUserId int `json:"ForUserId"` // ForUserId -
	PostId    int `json:"PostId"`    // PostId -
	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Home exists in the database.
func (h *Home) Exists() bool {
	return h._exists
}

// Deleted provides information if the Home has been deleted from the database.
func (h *Home) Deleted() bool {
	return h._deleted
}

// Insert inserts the Home to the database.
func (h *Home) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if h._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO sun_chat.home (` +
		`Id, ForUserId, PostId` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	if LogTableSqlReq.Home {
		XOLog(sqlstr, h.Id, h.ForUserId, h.PostId)
	}
	_, err = db.Exec(sqlstr, h.Id, h.ForUserId, h.PostId)
	if err != nil {
		return err
	}

	// set existence
	h._exists = true

	OnHome_AfterInsert(h)

	return nil
}

// Insert inserts the Home to the database.
func (h *Home) Replace(db XODB) error {
	var err error

	// sql query

	const sqlstr = `REPLACE INTO sun_chat.home (` +
		`Id, ForUserId, PostId` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	if LogTableSqlReq.Home {
		XOLog(sqlstr, h.Id, h.ForUserId, h.PostId)
	}
	_, err = db.Exec(sqlstr, h.Id, h.ForUserId, h.PostId)
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return err
	}

	h._exists = true

	OnHome_AfterInsert(h)

	return nil
}

// Update updates the Home in the database.
func (h *Home) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !h._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if h._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE sun_chat.home SET ` +
		`ForUserId = ?, PostId = ?` +
		` WHERE Id = ?`

	// run query
	if LogTableSqlReq.Home {
		XOLog(sqlstr, h.ForUserId, h.PostId, h.Id)
	}
	_, err = db.Exec(sqlstr, h.ForUserId, h.PostId, h.Id)

	if LogTableSqlReq.Home {
		XOLogErr(err)
	}
	OnHome_AfterUpdate(h)

	return err
}

// Save saves the Home to the database.
func (h *Home) Save(db XODB) error {
	if h.Exists() {
		return h.Update(db)
	}

	return h.Replace(db)
}

// Delete deletes the Home from the database.
func (h *Home) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !h._exists {
		return nil
	}

	// if deleted, bail
	if h._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM sun_chat.home WHERE Id = ?`

	// run query
	if LogTableSqlReq.Home {
		XOLog(sqlstr, h.Id)
	}
	_, err = db.Exec(sqlstr, h.Id)
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return err
	}

	// set deleted
	h._deleted = true

	OnHome_AfterDelete(h)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////// Querify gen - ME /////////////////////////////////////////
//.TableNameGo= table name
// _Deleter, _Updater

// orma types
type __Home_Deleter struct {
	wheres   []whereClause
	whereSep string
}

type __Home_Updater struct {
	wheres   []whereClause
	updates  map[string]interface{}
	whereSep string
}

type __Home_Selector struct {
	wheres    []whereClause
	selectCol string
	whereSep  string
	orderBy   string //" order by id desc //for ints
	limit     int
	offset    int
}

func NewHome_Deleter() *__Home_Deleter {
	d := __Home_Deleter{whereSep: " AND "}
	return &d
}

func NewHome_Updater() *__Home_Updater {
	u := __Home_Updater{whereSep: " AND "}
	u.updates = make(map[string]interface{}, 10)
	return &u
}

func NewHome_Selector() *__Home_Selector {
	u := __Home_Selector{whereSep: " AND ", selectCol: "*"}
	return &u
}

/////////////////////////////// Where for all /////////////////////////////
//// for ints all selector updater, deleter

////////ints
func (u *__Home_Deleter) Or() *__Home_Deleter {
	u.whereSep = " OR "
	return u
}

func (u *__Home_Deleter) Id_In(ins []int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Id IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Deleter) Id_Ins(ins ...int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Id IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Deleter) Id_NotIn(ins []int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Id NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__Home_Deleter) Id_Eq(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) Id_NotEq(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) Id_LT(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) Id_LE(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) Id_GT(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) Id_GE(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__Home_Deleter) ForUserId_In(ins []int) *__Home_Deleter {
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

func (u *__Home_Deleter) ForUserId_Ins(ins ...int) *__Home_Deleter {
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

func (u *__Home_Deleter) ForUserId_NotIn(ins []int) *__Home_Deleter {
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

func (d *__Home_Deleter) ForUserId_Eq(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) ForUserId_NotEq(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) ForUserId_LT(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) ForUserId_LE(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) ForUserId_GT(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) ForUserId_GE(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__Home_Deleter) PostId_In(ins []int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " PostId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Deleter) PostId_Ins(ins ...int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " PostId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Deleter) PostId_NotIn(ins []int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " PostId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__Home_Deleter) PostId_Eq(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) PostId_NotEq(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) PostId_LT(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) PostId_LE(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) PostId_GT(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Deleter) PostId_GE(val int) *__Home_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__Home_Updater) Or() *__Home_Updater {
	u.whereSep = " OR "
	return u
}

func (u *__Home_Updater) Id_In(ins []int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Id IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Updater) Id_Ins(ins ...int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Id IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Updater) Id_NotIn(ins []int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Id NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__Home_Updater) Id_Eq(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) Id_NotEq(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) Id_LT(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) Id_LE(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) Id_GT(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) Id_GE(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__Home_Updater) ForUserId_In(ins []int) *__Home_Updater {
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

func (u *__Home_Updater) ForUserId_Ins(ins ...int) *__Home_Updater {
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

func (u *__Home_Updater) ForUserId_NotIn(ins []int) *__Home_Updater {
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

func (d *__Home_Updater) ForUserId_Eq(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) ForUserId_NotEq(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) ForUserId_LT(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) ForUserId_LE(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) ForUserId_GT(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) ForUserId_GE(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__Home_Updater) PostId_In(ins []int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " PostId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Updater) PostId_Ins(ins ...int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " PostId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Updater) PostId_NotIn(ins []int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " PostId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__Home_Updater) PostId_Eq(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) PostId_NotEq(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) PostId_LT(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) PostId_LE(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) PostId_GT(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Updater) PostId_GE(val int) *__Home_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__Home_Selector) Or() *__Home_Selector {
	u.whereSep = " OR "
	return u
}

func (u *__Home_Selector) Id_In(ins []int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Id IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Selector) Id_Ins(ins ...int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Id IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Selector) Id_NotIn(ins []int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Id NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__Home_Selector) Id_Eq(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) Id_NotEq(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) Id_LT(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) Id_LE(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) Id_GT(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) Id_GE(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__Home_Selector) ForUserId_In(ins []int) *__Home_Selector {
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

func (u *__Home_Selector) ForUserId_Ins(ins ...int) *__Home_Selector {
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

func (u *__Home_Selector) ForUserId_NotIn(ins []int) *__Home_Selector {
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

func (d *__Home_Selector) ForUserId_Eq(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) ForUserId_NotEq(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) ForUserId_LT(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) ForUserId_LE(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) ForUserId_GT(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) ForUserId_GE(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__Home_Selector) PostId_In(ins []int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " PostId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Selector) PostId_Ins(ins ...int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " PostId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__Home_Selector) PostId_NotIn(ins []int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " PostId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__Home_Selector) PostId_Eq(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) PostId_NotEq(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) PostId_LT(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) PostId_LE(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) PostId_GT(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__Home_Selector) PostId_GE(val int) *__Home_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

///// for strings //copy of above with type int -> string + rm if eq + $ms_str_cond

////////ints

////////ints

////////ints

/// End of wheres for selectors , updators, deletor

/////////////////////////////// Updater /////////////////////////////

//ints

func (u *__Home_Updater) Id(newVal int) *__Home_Updater {
	u.updates[" Id = ? "] = newVal
	return u
}

func (u *__Home_Updater) Id_Increment(count int) *__Home_Updater {
	if count > 0 {
		u.updates[" Id = Id+? "] = count
	}

	if count < 0 {
		u.updates[" Id = Id-? "] = -(count) //make it positive
	}

	return u
}

//string

//ints

func (u *__Home_Updater) ForUserId(newVal int) *__Home_Updater {
	u.updates[" ForUserId = ? "] = newVal
	return u
}

func (u *__Home_Updater) ForUserId_Increment(count int) *__Home_Updater {
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

func (u *__Home_Updater) PostId(newVal int) *__Home_Updater {
	u.updates[" PostId = ? "] = newVal
	return u
}

func (u *__Home_Updater) PostId_Increment(count int) *__Home_Updater {
	if count > 0 {
		u.updates[" PostId = PostId+? "] = count
	}

	if count < 0 {
		u.updates[" PostId = PostId-? "] = -(count) //make it positive
	}

	return u
}

//string

/////////////////////////////////////////////////////////////////////
/////////////////////// Selector ///////////////////////////////////

//Select_* can just be used with: .GetString() , .GetStringSlice(), .GetInt() ..GetIntSlice()

func (u *__Home_Selector) OrderBy_Id_Desc() *__Home_Selector {
	u.orderBy = " ORDER BY Id DESC "
	return u
}

func (u *__Home_Selector) OrderBy_Id_Asc() *__Home_Selector {
	u.orderBy = " ORDER BY Id ASC "
	return u
}

func (u *__Home_Selector) Select_Id() *__Home_Selector {
	u.selectCol = "Id"
	return u
}

func (u *__Home_Selector) OrderBy_ForUserId_Desc() *__Home_Selector {
	u.orderBy = " ORDER BY ForUserId DESC "
	return u
}

func (u *__Home_Selector) OrderBy_ForUserId_Asc() *__Home_Selector {
	u.orderBy = " ORDER BY ForUserId ASC "
	return u
}

func (u *__Home_Selector) Select_ForUserId() *__Home_Selector {
	u.selectCol = "ForUserId"
	return u
}

func (u *__Home_Selector) OrderBy_PostId_Desc() *__Home_Selector {
	u.orderBy = " ORDER BY PostId DESC "
	return u
}

func (u *__Home_Selector) OrderBy_PostId_Asc() *__Home_Selector {
	u.orderBy = " ORDER BY PostId ASC "
	return u
}

func (u *__Home_Selector) Select_PostId() *__Home_Selector {
	u.selectCol = "PostId"
	return u
}

func (u *__Home_Selector) Limit(num int) *__Home_Selector {
	u.limit = num
	return u
}

func (u *__Home_Selector) Offset(num int) *__Home_Selector {
	u.offset = num
	return u
}

func (u *__Home_Selector) Order_Rand() *__Home_Selector {
	u.orderBy = " ORDER BY RAND() "
	return u
}

/////////////////////////  Queryer Selector  //////////////////////////////////
func (u *__Home_Selector) _stoSql() (string, []interface{}) {
	sqlWherrs, whereArgs := whereClusesToSql(u.wheres, u.whereSep)

	sqlstr := "SELECT " + u.selectCol + " FROM sun_chat.home"

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

func (u *__Home_Selector) GetRow(db *sqlx.DB) (*Home, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.Home {
		XOLog(sqlstr, whereArgs)
	}

	row := &Home{}
	//by Sqlx
	err = db.Get(row, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return nil, err
	}

	row._exists = true

	OnHome_LoadOne(row)

	return row, nil
}

func (u *__Home_Selector) GetRows(db *sqlx.DB) ([]*Home, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.Home {
		XOLog(sqlstr, whereArgs)
	}

	var rows []*Home
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.Home {
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

	OnHome_LoadMany(rows)

	return rows, nil
}

//dep use GetRows()
func (u *__Home_Selector) GetRows2(db *sqlx.DB) ([]Home, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.Home {
		XOLog(sqlstr, whereArgs)
	}
	var rows []*Home
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.Home {
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

	OnHome_LoadMany(rows)

	rows2 := make([]Home, len(rows))
	for i := 0; i < len(rows); i++ {
		cp := *rows[i]
		rows2[i] = cp
	}

	return rows2, nil
}

func (u *__Home_Selector) GetString(db *sqlx.DB) (string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.Home {
		XOLog(sqlstr, whereArgs)
	}

	var res string
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return "", err
	}

	return res, nil
}

func (u *__Home_Selector) GetStringSlice(db *sqlx.DB) ([]string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.Home {
		XOLog(sqlstr, whereArgs)
	}
	var rows []string
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__Home_Selector) GetIntSlice(db *sqlx.DB) ([]int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.Home {
		XOLog(sqlstr, whereArgs)
	}
	var rows []int
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__Home_Selector) GetInt(db *sqlx.DB) (int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.Home {
		XOLog(sqlstr, whereArgs)
	}
	var res int
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return 0, err
	}

	return res, nil
}

/////////////////////////  Queryer Update Delete //////////////////////////////////
func (u *__Home_Updater) Update(db XODB) (int, error) {
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

	sqlstr := `UPDATE sun_chat.home SET ` + sqlUpdate

	if len(strings.Trim(sqlWherrs, " ")) > 0 { //2 for safty
		sqlstr += " WHERE " + sqlWherrs
	}

	if LogTableSqlReq.Home {
		XOLog(sqlstr, allArgs)
	}
	res, err := db.Exec(sqlstr, allArgs...)
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return 0, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

func (d *__Home_Deleter) Delete(db XODB) (int, error) {
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

	sqlstr := "DELETE FROM sun_chat.home WHERE " + wheresStr

	// run query
	if LogTableSqlReq.Home {
		XOLog(sqlstr, args)
	}
	res, err := db.Exec(sqlstr, args...)
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return 0, err
	}

	// retrieve id
	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

///////////////////////// Mass insert - replace for  Home ////////////////

func MassInsert_Home(rows []Home, db XODB) error {
	if len(rows) == 0 {
		return errors.New("rows slice should not be empty - inserted nothing")
	}
	var err error
	ln := len(rows)
	//s:= "(?,?,?)," //`(?, ?, ?, ?),`
	s := "(?,?,?)," //`(?, ?, ?, ?),`
	insVals_ := strings.Repeat(s, ln)
	insVals := insVals_[0 : len(insVals_)-1]
	// sql query
	sqlstr := "INSERT INTO sun_chat.home (" +
		"Id, ForUserId, PostId" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.Id)
		vals = append(vals, row.ForUserId)
		vals = append(vals, row.PostId)

	}

	if LogTableSqlReq.Home {
		XOLog(sqlstr, " MassInsert len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.Home {
			XOLogErr(err)
		}
		return err
	}

	return nil
}

func MassReplace_Home(rows []Home, db XODB) error {
	if len(rows) == 0 {
		return errors.New("rows slice should not be empty - inserted nothing")
	}
	var err error
	ln := len(rows)
	//s:= "(?,?,?)," //`(?, ?, ?, ?),`
	s := "(?,?,?)," //`(?, ?, ?, ?),`
	insVals_ := strings.Repeat(s, ln)
	insVals := insVals_[0 : len(insVals_)-1]
	// sql query
	sqlstr := "REPLACE INTO sun_chat.home (" +
		"Id, ForUserId, PostId" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.Id)
		vals = append(vals, row.ForUserId)
		vals = append(vals, row.PostId)

	}

	if LogTableSqlReq.Home {
		XOLog(sqlstr, " MassReplace len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.Home {
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
