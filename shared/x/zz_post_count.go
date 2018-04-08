package x

// GENERATED BY XO. DO NOT EDIT.
import (
	"errors"
	"strings"
	//"time"
	"ms/sun_old/helper"
	"strconv"

	"github.com/jmoiron/sqlx"
) // (shortname .TableNameGo "err" "res" "sqlstr" "db" "XOLog") -}}//(schema .Schema .Table.TableName) -}}// .TableNameGo}}// PostCount represents a row from 'sun.post_count'.

// Manualy copy this to project
type PostCount__ struct {
	PostId     int `json:"PostId"`     // PostId -
	ViewsCount int `json:"ViewsCount"` // ViewsCount -
	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PostCount exists in the database.
func (pc *PostCount) Exists() bool {
	return pc._exists
}

// Deleted provides information if the PostCount has been deleted from the database.
func (pc *PostCount) Deleted() bool {
	return pc._deleted
}

// Insert inserts the PostCount to the database.
func (pc *PostCount) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pc._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO sun.post_count (` +
		`PostId, ViewsCount` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, pc.PostId, pc.ViewsCount)
	}
	_, err = db.Exec(sqlstr, pc.PostId, pc.ViewsCount)
	if err != nil {
		return err
	}

	// set existence
	pc._exists = true

	OnPostCount_AfterInsert(pc)

	return nil
}

// Insert inserts the PostCount to the database.
func (pc *PostCount) Replace(db XODB) error {
	var err error

	// sql query

	const sqlstr = `REPLACE INTO sun.post_count (` +
		`PostId, ViewsCount` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, pc.PostId, pc.ViewsCount)
	}
	_, err = db.Exec(sqlstr, pc.PostId, pc.ViewsCount)
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return err
	}

	pc._exists = true

	OnPostCount_AfterInsert(pc)

	return nil
}

// Update updates the PostCount in the database.
func (pc *PostCount) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pc._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if pc._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE sun.post_count SET ` +
		`ViewsCount = ?` +
		` WHERE PostId = ?`

	// run query
	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, pc.ViewsCount, pc.PostId)
	}
	_, err = db.Exec(sqlstr, pc.ViewsCount, pc.PostId)

	if LogTableSqlReq.PostCount {
		XOLogErr(err)
	}
	OnPostCount_AfterUpdate(pc)

	return err
}

// Save saves the PostCount to the database.
func (pc *PostCount) Save(db XODB) error {
	if pc.Exists() {
		return pc.Update(db)
	}

	return pc.Replace(db)
}

// Delete deletes the PostCount from the database.
func (pc *PostCount) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pc._exists {
		return nil
	}

	// if deleted, bail
	if pc._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM sun.post_count WHERE PostId = ?`

	// run query
	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, pc.PostId)
	}
	_, err = db.Exec(sqlstr, pc.PostId)
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return err
	}

	// set deleted
	pc._deleted = true

	OnPostCount_AfterDelete(pc)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////// Querify gen - ME /////////////////////////////////////////
//.TableNameGo= table name
// _Deleter, _Updater

// orma types
type __PostCount_Deleter struct {
	wheres   []whereClause
	whereSep string
}

type __PostCount_Updater struct {
	wheres   []whereClause
	updates  map[string]interface{}
	whereSep string
}

type __PostCount_Selector struct {
	wheres    []whereClause
	selectCol string
	whereSep  string
	orderBy   string //" order by id desc //for ints
	limit     int
	offset    int
}

func NewPostCount_Deleter() *__PostCount_Deleter {
	d := __PostCount_Deleter{whereSep: " AND "}
	return &d
}

func NewPostCount_Updater() *__PostCount_Updater {
	u := __PostCount_Updater{whereSep: " AND "}
	u.updates = make(map[string]interface{}, 10)
	return &u
}

func NewPostCount_Selector() *__PostCount_Selector {
	u := __PostCount_Selector{whereSep: " AND ", selectCol: "*"}
	return &u
}

/////////////////////////////// Where for all /////////////////////////////
//// for ints all selector updater, deleter

////////ints
func (u *__PostCount_Deleter) Or() *__PostCount_Deleter {
	u.whereSep = " OR "
	return u
}

func (u *__PostCount_Deleter) PostId_In(ins []int) *__PostCount_Deleter {
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

func (u *__PostCount_Deleter) PostId_Ins(ins ...int) *__PostCount_Deleter {
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

func (u *__PostCount_Deleter) PostId_NotIn(ins []int) *__PostCount_Deleter {
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

func (d *__PostCount_Deleter) PostId_Eq(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Deleter) PostId_NotEq(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Deleter) PostId_LT(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Deleter) PostId_LE(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Deleter) PostId_GT(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Deleter) PostId_GE(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__PostCount_Deleter) ViewsCount_In(ins []int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ViewsCount IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostCount_Deleter) ViewsCount_Ins(ins ...int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ViewsCount IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostCount_Deleter) ViewsCount_NotIn(ins []int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ViewsCount NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__PostCount_Deleter) ViewsCount_Eq(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Deleter) ViewsCount_NotEq(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Deleter) ViewsCount_LT(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Deleter) ViewsCount_LE(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Deleter) ViewsCount_GT(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Deleter) ViewsCount_GE(val int) *__PostCount_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__PostCount_Updater) Or() *__PostCount_Updater {
	u.whereSep = " OR "
	return u
}

func (u *__PostCount_Updater) PostId_In(ins []int) *__PostCount_Updater {
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

func (u *__PostCount_Updater) PostId_Ins(ins ...int) *__PostCount_Updater {
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

func (u *__PostCount_Updater) PostId_NotIn(ins []int) *__PostCount_Updater {
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

func (d *__PostCount_Updater) PostId_Eq(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Updater) PostId_NotEq(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Updater) PostId_LT(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Updater) PostId_LE(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Updater) PostId_GT(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Updater) PostId_GE(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__PostCount_Updater) ViewsCount_In(ins []int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ViewsCount IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostCount_Updater) ViewsCount_Ins(ins ...int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ViewsCount IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostCount_Updater) ViewsCount_NotIn(ins []int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ViewsCount NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__PostCount_Updater) ViewsCount_Eq(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Updater) ViewsCount_NotEq(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Updater) ViewsCount_LT(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Updater) ViewsCount_LE(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Updater) ViewsCount_GT(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Updater) ViewsCount_GE(val int) *__PostCount_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__PostCount_Selector) Or() *__PostCount_Selector {
	u.whereSep = " OR "
	return u
}

func (u *__PostCount_Selector) PostId_In(ins []int) *__PostCount_Selector {
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

func (u *__PostCount_Selector) PostId_Ins(ins ...int) *__PostCount_Selector {
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

func (u *__PostCount_Selector) PostId_NotIn(ins []int) *__PostCount_Selector {
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

func (d *__PostCount_Selector) PostId_Eq(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Selector) PostId_NotEq(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Selector) PostId_LT(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Selector) PostId_LE(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Selector) PostId_GT(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Selector) PostId_GE(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " PostId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__PostCount_Selector) ViewsCount_In(ins []int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ViewsCount IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostCount_Selector) ViewsCount_Ins(ins ...int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ViewsCount IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostCount_Selector) ViewsCount_NotIn(ins []int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " ViewsCount NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__PostCount_Selector) ViewsCount_Eq(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Selector) ViewsCount_NotEq(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Selector) ViewsCount_LT(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Selector) ViewsCount_LE(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Selector) ViewsCount_GT(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostCount_Selector) ViewsCount_GE(val int) *__PostCount_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " ViewsCount >= ? "
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

func (u *__PostCount_Updater) PostId(newVal int) *__PostCount_Updater {
	u.updates[" PostId = ? "] = newVal
	return u
}

func (u *__PostCount_Updater) PostId_Increment(count int) *__PostCount_Updater {
	if count > 0 {
		u.updates[" PostId = PostId+? "] = count
	}

	if count < 0 {
		u.updates[" PostId = PostId-? "] = -(count) //make it positive
	}

	return u
}

//string

//ints

func (u *__PostCount_Updater) ViewsCount(newVal int) *__PostCount_Updater {
	u.updates[" ViewsCount = ? "] = newVal
	return u
}

func (u *__PostCount_Updater) ViewsCount_Increment(count int) *__PostCount_Updater {
	if count > 0 {
		u.updates[" ViewsCount = ViewsCount+? "] = count
	}

	if count < 0 {
		u.updates[" ViewsCount = ViewsCount-? "] = -(count) //make it positive
	}

	return u
}

//string

/////////////////////////////////////////////////////////////////////
/////////////////////// Selector ///////////////////////////////////

//Select_* can just be used with: .GetString() , .GetStringSlice(), .GetInt() ..GetIntSlice()

func (u *__PostCount_Selector) OrderBy_PostId_Desc() *__PostCount_Selector {
	u.orderBy = " ORDER BY PostId DESC "
	return u
}

func (u *__PostCount_Selector) OrderBy_PostId_Asc() *__PostCount_Selector {
	u.orderBy = " ORDER BY PostId ASC "
	return u
}

func (u *__PostCount_Selector) Select_PostId() *__PostCount_Selector {
	u.selectCol = "PostId"
	return u
}

func (u *__PostCount_Selector) OrderBy_ViewsCount_Desc() *__PostCount_Selector {
	u.orderBy = " ORDER BY ViewsCount DESC "
	return u
}

func (u *__PostCount_Selector) OrderBy_ViewsCount_Asc() *__PostCount_Selector {
	u.orderBy = " ORDER BY ViewsCount ASC "
	return u
}

func (u *__PostCount_Selector) Select_ViewsCount() *__PostCount_Selector {
	u.selectCol = "ViewsCount"
	return u
}

func (u *__PostCount_Selector) Limit(num int) *__PostCount_Selector {
	u.limit = num
	return u
}

func (u *__PostCount_Selector) Offset(num int) *__PostCount_Selector {
	u.offset = num
	return u
}

func (u *__PostCount_Selector) Order_Rand() *__PostCount_Selector {
	u.orderBy = " ORDER BY RAND() "
	return u
}

/////////////////////////  Queryer Selector  //////////////////////////////////
func (u *__PostCount_Selector) _stoSql() (string, []interface{}) {
	sqlWherrs, whereArgs := whereClusesToSql(u.wheres, u.whereSep)

	sqlstr := "SELECT " + u.selectCol + " FROM sun.post_count"

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

func (u *__PostCount_Selector) GetRow(db *sqlx.DB) (*PostCount, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, whereArgs)
	}

	row := &PostCount{}
	//by Sqlx
	err = db.Get(row, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return nil, err
	}

	row._exists = true

	OnPostCount_LoadOne(row)

	return row, nil
}

func (u *__PostCount_Selector) GetRows(db *sqlx.DB) ([]*PostCount, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, whereArgs)
	}

	var rows []*PostCount
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostCount {
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

	OnPostCount_LoadMany(rows)

	return rows, nil
}

//dep use GetRows()
func (u *__PostCount_Selector) GetRows2(db *sqlx.DB) ([]PostCount, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, whereArgs)
	}
	var rows []*PostCount
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostCount {
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

	OnPostCount_LoadMany(rows)

	rows2 := make([]PostCount, len(rows))
	for i := 0; i < len(rows); i++ {
		cp := *rows[i]
		rows2[i] = cp
	}

	return rows2, nil
}

func (u *__PostCount_Selector) GetString(db *sqlx.DB) (string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, whereArgs)
	}

	var res string
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return "", err
	}

	return res, nil
}

func (u *__PostCount_Selector) GetStringSlice(db *sqlx.DB) ([]string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, whereArgs)
	}
	var rows []string
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__PostCount_Selector) GetIntSlice(db *sqlx.DB) ([]int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, whereArgs)
	}
	var rows []int
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__PostCount_Selector) GetInt(db *sqlx.DB) (int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, whereArgs)
	}
	var res int
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return 0, err
	}

	return res, nil
}

/////////////////////////  Queryer Update Delete //////////////////////////////////
func (u *__PostCount_Updater) Update(db XODB) (int, error) {
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

	sqlstr := `UPDATE sun.post_count SET ` + sqlUpdate

	if len(strings.Trim(sqlWherrs, " ")) > 0 { //2 for safty
		sqlstr += " WHERE " + sqlWherrs
	}

	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, allArgs)
	}
	res, err := db.Exec(sqlstr, allArgs...)
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return 0, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

func (d *__PostCount_Deleter) Delete(db XODB) (int, error) {
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

	sqlstr := "DELETE FROM sun.post_count WHERE " + wheresStr

	// run query
	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, args)
	}
	res, err := db.Exec(sqlstr, args...)
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return 0, err
	}

	// retrieve id
	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

///////////////////////// Mass insert - replace for  PostCount ////////////////

func MassInsert_PostCount(rows []PostCount, db XODB) error {
	if len(rows) == 0 {
		return errors.New("rows slice should not be empty - inserted nothing")
	}
	var err error
	ln := len(rows)
	//s:= "(?,?)," //`(?, ?, ?, ?),`
	s := "(?,?)," //`(?, ?, ?, ?),`
	insVals_ := strings.Repeat(s, ln)
	insVals := insVals_[0 : len(insVals_)-1]
	// sql query
	sqlstr := "INSERT INTO sun.post_count (" +
		"PostId, ViewsCount" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.PostId)
		vals = append(vals, row.ViewsCount)

	}

	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, " MassInsert len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return err
	}

	return nil
}

func MassReplace_PostCount(rows []PostCount, db XODB) error {
	if len(rows) == 0 {
		return errors.New("rows slice should not be empty - inserted nothing")
	}
	var err error
	ln := len(rows)
	//s:= "(?,?)," //`(?, ?, ?, ?),`
	s := "(?,?)," //`(?, ?, ?, ?),`
	insVals_ := strings.Repeat(s, ln)
	insVals := insVals_[0 : len(insVals_)-1]
	// sql query
	sqlstr := "REPLACE INTO sun.post_count (" +
		"PostId, ViewsCount" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.PostId)
		vals = append(vals, row.ViewsCount)

	}

	if LogTableSqlReq.PostCount {
		XOLog(sqlstr, " MassReplace len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.PostCount {
			XOLogErr(err)
		}
		return err
	}

	return nil

}

//////////////////// Play ///////////////////////////////

//

//
