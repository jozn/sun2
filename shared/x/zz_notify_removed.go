package x

// GENERATED BY XO. DO NOT EDIT.
import (
	"errors"
	"strings"
	//"time"
	"ms/sun/helper"
	"strconv"

	"github.com/jmoiron/sqlx"
) // (shortname .TableNameGo "err" "res" "sqlstr" "db" "XOLog") -}}//(schema .Schema .Table.TableName) -}}// .TableNameGo}}// NotifyRemoved represents a row from 'sun.notify_removed'.

// Manualy copy this to project
type NotifyRemoved__ struct {
	Murmur64Hash int `json:"Murmur64Hash"` // Murmur64Hash -
	ForUserId    int `json:"ForUserId"`    // ForUserId -
	Id           int `json:"Id"`           // Id -
	// xo fields
	_exists, _deleted bool
}

// Exists determines if the NotifyRemoved exists in the database.
func (nr *NotifyRemoved) Exists() bool {
	return nr._exists
}

// Deleted provides information if the NotifyRemoved has been deleted from the database.
func (nr *NotifyRemoved) Deleted() bool {
	return nr._deleted
}

// Insert inserts the NotifyRemoved to the database.
func (nr *NotifyRemoved) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if nr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO sun.notify_removed (` +
		`Murmur64Hash, ForUserId` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, nr.Murmur64Hash, nr.ForUserId)
	}
	res, err := db.Exec(sqlstr, nr.Murmur64Hash, nr.ForUserId)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return err
	}

	// set primary key and existence
	nr.Id = int(id)
	nr._exists = true

	OnNotifyRemoved_AfterInsert(nr)

	return nil
}

// Insert inserts the NotifyRemoved to the database.
func (nr *NotifyRemoved) Replace(db XODB) error {
	var err error

	// sql query

	const sqlstr = `REPLACE INTO sun.notify_removed (` +
		`Murmur64Hash, ForUserId` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, nr.Murmur64Hash, nr.ForUserId)
	}
	res, err := db.Exec(sqlstr, nr.Murmur64Hash, nr.ForUserId)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return err
	}

	// set primary key and existence
	nr.Id = int(id)
	nr._exists = true

	OnNotifyRemoved_AfterInsert(nr)

	return nil
}

// Update updates the NotifyRemoved in the database.
func (nr *NotifyRemoved) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !nr._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if nr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE sun.notify_removed SET ` +
		`Murmur64Hash = ?, ForUserId = ?` +
		` WHERE Id = ?`

	// run query
	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, nr.Murmur64Hash, nr.ForUserId, nr.Id)
	}
	_, err = db.Exec(sqlstr, nr.Murmur64Hash, nr.ForUserId, nr.Id)

	if LogTableSqlReq.NotifyRemoved {
		XOLogErr(err)
	}
	OnNotifyRemoved_AfterUpdate(nr)

	return err
}

// Save saves the NotifyRemoved to the database.
func (nr *NotifyRemoved) Save(db XODB) error {
	if nr.Exists() {
		return nr.Update(db)
	}

	return nr.Replace(db)
}

// Delete deletes the NotifyRemoved from the database.
func (nr *NotifyRemoved) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !nr._exists {
		return nil
	}

	// if deleted, bail
	if nr._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM sun.notify_removed WHERE Id = ?`

	// run query
	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, nr.Id)
	}
	_, err = db.Exec(sqlstr, nr.Id)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return err
	}

	// set deleted
	nr._deleted = true

	OnNotifyRemoved_AfterDelete(nr)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////// Querify gen - ME /////////////////////////////////////////
//.TableNameGo= table name
// _Deleter, _Updater

// orma types
type __NotifyRemoved_Deleter struct {
	wheres   []whereClause
	whereSep string
}

type __NotifyRemoved_Updater struct {
	wheres   []whereClause
	updates  map[string]interface{}
	whereSep string
}

type __NotifyRemoved_Selector struct {
	wheres    []whereClause
	selectCol string
	whereSep  string
	orderBy   string //" order by id desc //for ints
	limit     int
	offset    int
}

func NewNotifyRemoved_Deleter() *__NotifyRemoved_Deleter {
	d := __NotifyRemoved_Deleter{whereSep: " AND "}
	return &d
}

func NewNotifyRemoved_Updater() *__NotifyRemoved_Updater {
	u := __NotifyRemoved_Updater{whereSep: " AND "}
	u.updates = make(map[string]interface{}, 10)
	return &u
}

func NewNotifyRemoved_Selector() *__NotifyRemoved_Selector {
	u := __NotifyRemoved_Selector{whereSep: " AND ", selectCol: "*"}
	return &u
}

/////////////////////////////// Where for all /////////////////////////////
//// for ints all selector updater, deleter

////////ints
func (u *__NotifyRemoved_Deleter) Or() *__NotifyRemoved_Deleter {
	u.whereSep = " OR "
	return u
}

func (u *__NotifyRemoved_Deleter) Murmur64Hash_In(ins []int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Murmur64Hash IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__NotifyRemoved_Deleter) Murmur64Hash_Ins(ins ...int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Murmur64Hash IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__NotifyRemoved_Deleter) Murmur64Hash_NotIn(ins []int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Murmur64Hash NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__NotifyRemoved_Deleter) Murmur64Hash_Eq(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) Murmur64Hash_NotEq(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) Murmur64Hash_LT(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) Murmur64Hash_LE(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) Murmur64Hash_GT(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) Murmur64Hash_GE(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__NotifyRemoved_Deleter) ForUserId_In(ins []int) *__NotifyRemoved_Deleter {
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

func (u *__NotifyRemoved_Deleter) ForUserId_Ins(ins ...int) *__NotifyRemoved_Deleter {
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

func (u *__NotifyRemoved_Deleter) ForUserId_NotIn(ins []int) *__NotifyRemoved_Deleter {
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

func (d *__NotifyRemoved_Deleter) ForUserId_Eq(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) ForUserId_NotEq(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) ForUserId_LT(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) ForUserId_LE(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) ForUserId_GT(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) ForUserId_GE(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__NotifyRemoved_Deleter) Id_In(ins []int) *__NotifyRemoved_Deleter {
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

func (u *__NotifyRemoved_Deleter) Id_Ins(ins ...int) *__NotifyRemoved_Deleter {
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

func (u *__NotifyRemoved_Deleter) Id_NotIn(ins []int) *__NotifyRemoved_Deleter {
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

func (d *__NotifyRemoved_Deleter) Id_Eq(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) Id_NotEq(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) Id_LT(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) Id_LE(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) Id_GT(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Deleter) Id_GE(val int) *__NotifyRemoved_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__NotifyRemoved_Updater) Or() *__NotifyRemoved_Updater {
	u.whereSep = " OR "
	return u
}

func (u *__NotifyRemoved_Updater) Murmur64Hash_In(ins []int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Murmur64Hash IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__NotifyRemoved_Updater) Murmur64Hash_Ins(ins ...int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Murmur64Hash IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__NotifyRemoved_Updater) Murmur64Hash_NotIn(ins []int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Murmur64Hash NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__NotifyRemoved_Updater) Murmur64Hash_Eq(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) Murmur64Hash_NotEq(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) Murmur64Hash_LT(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) Murmur64Hash_LE(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) Murmur64Hash_GT(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) Murmur64Hash_GE(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__NotifyRemoved_Updater) ForUserId_In(ins []int) *__NotifyRemoved_Updater {
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

func (u *__NotifyRemoved_Updater) ForUserId_Ins(ins ...int) *__NotifyRemoved_Updater {
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

func (u *__NotifyRemoved_Updater) ForUserId_NotIn(ins []int) *__NotifyRemoved_Updater {
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

func (d *__NotifyRemoved_Updater) ForUserId_Eq(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) ForUserId_NotEq(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) ForUserId_LT(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) ForUserId_LE(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) ForUserId_GT(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) ForUserId_GE(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__NotifyRemoved_Updater) Id_In(ins []int) *__NotifyRemoved_Updater {
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

func (u *__NotifyRemoved_Updater) Id_Ins(ins ...int) *__NotifyRemoved_Updater {
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

func (u *__NotifyRemoved_Updater) Id_NotIn(ins []int) *__NotifyRemoved_Updater {
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

func (d *__NotifyRemoved_Updater) Id_Eq(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) Id_NotEq(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) Id_LT(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) Id_LE(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) Id_GT(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Updater) Id_GE(val int) *__NotifyRemoved_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__NotifyRemoved_Selector) Or() *__NotifyRemoved_Selector {
	u.whereSep = " OR "
	return u
}

func (u *__NotifyRemoved_Selector) Murmur64Hash_In(ins []int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Murmur64Hash IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__NotifyRemoved_Selector) Murmur64Hash_Ins(ins ...int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Murmur64Hash IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__NotifyRemoved_Selector) Murmur64Hash_NotIn(ins []int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Murmur64Hash NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__NotifyRemoved_Selector) Murmur64Hash_Eq(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) Murmur64Hash_NotEq(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) Murmur64Hash_LT(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) Murmur64Hash_LE(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) Murmur64Hash_GT(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) Murmur64Hash_GE(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Murmur64Hash >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__NotifyRemoved_Selector) ForUserId_In(ins []int) *__NotifyRemoved_Selector {
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

func (u *__NotifyRemoved_Selector) ForUserId_Ins(ins ...int) *__NotifyRemoved_Selector {
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

func (u *__NotifyRemoved_Selector) ForUserId_NotIn(ins []int) *__NotifyRemoved_Selector {
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

func (d *__NotifyRemoved_Selector) ForUserId_Eq(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) ForUserId_NotEq(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) ForUserId_LT(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) ForUserId_LE(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) ForUserId_GT(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) ForUserId_GE(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ForUserId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__NotifyRemoved_Selector) Id_In(ins []int) *__NotifyRemoved_Selector {
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

func (u *__NotifyRemoved_Selector) Id_Ins(ins ...int) *__NotifyRemoved_Selector {
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

func (u *__NotifyRemoved_Selector) Id_NotIn(ins []int) *__NotifyRemoved_Selector {
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

func (d *__NotifyRemoved_Selector) Id_Eq(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) Id_NotEq(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) Id_LT(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) Id_LE(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) Id_GT(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__NotifyRemoved_Selector) Id_GE(val int) *__NotifyRemoved_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
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

func (u *__NotifyRemoved_Updater) Murmur64Hash(newVal int) *__NotifyRemoved_Updater {
	u.updates[" Murmur64Hash = ? "] = newVal
	return u
}

func (u *__NotifyRemoved_Updater) Murmur64Hash_Increment(count int) *__NotifyRemoved_Updater {
	if count > 0 {
		u.updates[" Murmur64Hash = Murmur64Hash+? "] = count
	}

	if count < 0 {
		u.updates[" Murmur64Hash = Murmur64Hash-? "] = -(count) //make it positive
	}

	return u
}

//string

//ints

func (u *__NotifyRemoved_Updater) ForUserId(newVal int) *__NotifyRemoved_Updater {
	u.updates[" ForUserId = ? "] = newVal
	return u
}

func (u *__NotifyRemoved_Updater) ForUserId_Increment(count int) *__NotifyRemoved_Updater {
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

func (u *__NotifyRemoved_Updater) Id(newVal int) *__NotifyRemoved_Updater {
	u.updates[" Id = ? "] = newVal
	return u
}

func (u *__NotifyRemoved_Updater) Id_Increment(count int) *__NotifyRemoved_Updater {
	if count > 0 {
		u.updates[" Id = Id+? "] = count
	}

	if count < 0 {
		u.updates[" Id = Id-? "] = -(count) //make it positive
	}

	return u
}

//string

/////////////////////////////////////////////////////////////////////
/////////////////////// Selector ///////////////////////////////////

//Select_* can just be used with: .GetString() , .GetStringSlice(), .GetInt() ..GetIntSlice()

func (u *__NotifyRemoved_Selector) OrderBy_Murmur64Hash_Desc() *__NotifyRemoved_Selector {
	u.orderBy = " ORDER BY Murmur64Hash DESC "
	return u
}

func (u *__NotifyRemoved_Selector) OrderBy_Murmur64Hash_Asc() *__NotifyRemoved_Selector {
	u.orderBy = " ORDER BY Murmur64Hash ASC "
	return u
}

func (u *__NotifyRemoved_Selector) Select_Murmur64Hash() *__NotifyRemoved_Selector {
	u.selectCol = "Murmur64Hash"
	return u
}

func (u *__NotifyRemoved_Selector) OrderBy_ForUserId_Desc() *__NotifyRemoved_Selector {
	u.orderBy = " ORDER BY ForUserId DESC "
	return u
}

func (u *__NotifyRemoved_Selector) OrderBy_ForUserId_Asc() *__NotifyRemoved_Selector {
	u.orderBy = " ORDER BY ForUserId ASC "
	return u
}

func (u *__NotifyRemoved_Selector) Select_ForUserId() *__NotifyRemoved_Selector {
	u.selectCol = "ForUserId"
	return u
}

func (u *__NotifyRemoved_Selector) OrderBy_Id_Desc() *__NotifyRemoved_Selector {
	u.orderBy = " ORDER BY Id DESC "
	return u
}

func (u *__NotifyRemoved_Selector) OrderBy_Id_Asc() *__NotifyRemoved_Selector {
	u.orderBy = " ORDER BY Id ASC "
	return u
}

func (u *__NotifyRemoved_Selector) Select_Id() *__NotifyRemoved_Selector {
	u.selectCol = "Id"
	return u
}

func (u *__NotifyRemoved_Selector) Limit(num int) *__NotifyRemoved_Selector {
	u.limit = num
	return u
}

func (u *__NotifyRemoved_Selector) Offset(num int) *__NotifyRemoved_Selector {
	u.offset = num
	return u
}

func (u *__NotifyRemoved_Selector) Order_Rand() *__NotifyRemoved_Selector {
	u.orderBy = " ORDER BY RAND() "
	return u
}

/////////////////////////  Queryer Selector  //////////////////////////////////
func (u *__NotifyRemoved_Selector) _stoSql() (string, []interface{}) {
	sqlWherrs, whereArgs := whereClusesToSql(u.wheres, u.whereSep)

	sqlstr := "SELECT " + u.selectCol + " FROM sun.notify_removed"

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

func (u *__NotifyRemoved_Selector) GetRow(db *sqlx.DB) (*NotifyRemoved, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, whereArgs)
	}

	row := &NotifyRemoved{}
	//by Sqlx
	err = db.Get(row, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return nil, err
	}

	row._exists = true

	OnNotifyRemoved_LoadOne(row)

	return row, nil
}

func (u *__NotifyRemoved_Selector) GetRows(db *sqlx.DB) ([]*NotifyRemoved, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, whereArgs)
	}

	var rows []*NotifyRemoved
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
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

	OnNotifyRemoved_LoadMany(rows)

	return rows, nil
}

//dep use GetRows()
func (u *__NotifyRemoved_Selector) GetRows2(db *sqlx.DB) ([]NotifyRemoved, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, whereArgs)
	}
	var rows []*NotifyRemoved
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
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

	OnNotifyRemoved_LoadMany(rows)

	rows2 := make([]NotifyRemoved, len(rows))
	for i := 0; i < len(rows); i++ {
		cp := *rows[i]
		rows2[i] = cp
	}

	return rows2, nil
}

func (u *__NotifyRemoved_Selector) GetString(db *sqlx.DB) (string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, whereArgs)
	}

	var res string
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return "", err
	}

	return res, nil
}

func (u *__NotifyRemoved_Selector) GetStringSlice(db *sqlx.DB) ([]string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, whereArgs)
	}
	var rows []string
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__NotifyRemoved_Selector) GetIntSlice(db *sqlx.DB) ([]int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, whereArgs)
	}
	var rows []int
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__NotifyRemoved_Selector) GetInt(db *sqlx.DB) (int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, whereArgs)
	}
	var res int
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return 0, err
	}

	return res, nil
}

/////////////////////////  Queryer Update Delete //////////////////////////////////
func (u *__NotifyRemoved_Updater) Update(db XODB) (int, error) {
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

	sqlstr := `UPDATE sun.notify_removed SET ` + sqlUpdate

	if len(strings.Trim(sqlWherrs, " ")) > 0 { //2 for safty
		sqlstr += " WHERE " + sqlWherrs
	}

	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, allArgs)
	}
	res, err := db.Exec(sqlstr, allArgs...)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return 0, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

func (d *__NotifyRemoved_Deleter) Delete(db XODB) (int, error) {
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

	sqlstr := "DELETE FROM sun.notify_removed WHERE " + wheresStr

	// run query
	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, args)
	}
	res, err := db.Exec(sqlstr, args...)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return 0, err
	}

	// retrieve id
	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

///////////////////////// Mass insert - replace for  NotifyRemoved ////////////////

func MassInsert_NotifyRemoved(rows []NotifyRemoved, db XODB) error {
	if len(rows) == 0 {
		return errors.New("rows slice should not be empty - inserted nothing")
	}
	var err error
	ln := len(rows)
	//s:= "( ms_question_mark .Columns .PrimaryKey.ColumnName }})," //`(?, ?, ?, ?),`
	s := "(?,?)," //`(?, ?, ?, ?),`
	insVals_ := strings.Repeat(s, ln)
	insVals := insVals_[0 : len(insVals_)-1]
	// sql query
	sqlstr := "INSERT INTO sun.notify_removed (" +
		"Murmur64Hash, ForUserId" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.Murmur64Hash)
		vals = append(vals, row.ForUserId)

	}

	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, " MassInsert len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
			XOLogErr(err)
		}
		return err
	}

	return nil
}

func MassReplace_NotifyRemoved(rows []NotifyRemoved, db XODB) error {
	var err error
	ln := len(rows)
	s := "(?,?)," //`(?, ?, ?, ?),`
	insVals_ := strings.Repeat(s, ln)
	insVals := insVals_[0 : len(insVals_)-1]
	// sql query
	sqlstr := "REPLACE INTO sun.notify_removed (" +
		"Murmur64Hash, ForUserId" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.Murmur64Hash)
		vals = append(vals, row.ForUserId)

	}

	if LogTableSqlReq.NotifyRemoved {
		XOLog(sqlstr, " MassReplace len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.NotifyRemoved {
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
