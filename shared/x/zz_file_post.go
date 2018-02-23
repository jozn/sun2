package x

// GENERATED BY XO. DO NOT EDIT.
import (
	"errors"
	"strings"
	//"time"
	"ms/sun/helper"
	"strconv"

	"github.com/jmoiron/sqlx"
) // (shortname .TableNameGo "err" "res" "sqlstr" "db" "XOLog") -}}//(schema .Schema .Table.TableName) -}}// .TableNameGo}}// FilePost represents a row from 'sun_file.file_post'.

// Manualy copy this to project
type FilePost__ struct {
	Id        int    `json:"Id"`        // Id -
	FileType  int    `json:"FileType"`  // FileType -
	Extension string `json:"Extension"` // Extension -
	DataThumb []byte `json:"DataThumb"` // DataThumb -
	Data      []byte `json:"Data"`      // Data -
	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FilePost exists in the database.
func (fp *FilePost) Exists() bool {
	return fp._exists
}

// Deleted provides information if the FilePost has been deleted from the database.
func (fp *FilePost) Deleted() bool {
	return fp._deleted
}

// Insert inserts the FilePost to the database.
func (fp *FilePost) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if fp._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO sun_file.file_post (` +
		`Id, FileType, Extension, DataThumb, Data` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, fp.Id, fp.FileType, fp.Extension, fp.DataThumb, fp.Data)
	}
	_, err = db.Exec(sqlstr, fp.Id, fp.FileType, fp.Extension, fp.DataThumb, fp.Data)
	if err != nil {
		return err
	}

	// set existence
	fp._exists = true

	OnFilePost_AfterInsert(fp)

	return nil
}

// Insert inserts the FilePost to the database.
func (fp *FilePost) Replace(db XODB) error {
	var err error

	// sql query

	const sqlstr = `REPLACE INTO sun_file.file_post (` +
		`Id, FileType, Extension, DataThumb, Data` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, fp.Id, fp.FileType, fp.Extension, fp.DataThumb, fp.Data)
	}
	_, err = db.Exec(sqlstr, fp.Id, fp.FileType, fp.Extension, fp.DataThumb, fp.Data)
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return err
	}

	fp._exists = true

	OnFilePost_AfterInsert(fp)

	return nil
}

// Update updates the FilePost in the database.
func (fp *FilePost) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !fp._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if fp._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE sun_file.file_post SET ` +
		`FileType = ?, Extension = ?, DataThumb = ?, Data = ?` +
		` WHERE Id = ?`

	// run query
	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, fp.FileType, fp.Extension, fp.DataThumb, fp.Data, fp.Id)
	}
	_, err = db.Exec(sqlstr, fp.FileType, fp.Extension, fp.DataThumb, fp.Data, fp.Id)

	if LogTableSqlReq.FilePost {
		XOLogErr(err)
	}
	OnFilePost_AfterUpdate(fp)

	return err
}

// Save saves the FilePost to the database.
func (fp *FilePost) Save(db XODB) error {
	if fp.Exists() {
		return fp.Update(db)
	}

	return fp.Replace(db)
}

// Delete deletes the FilePost from the database.
func (fp *FilePost) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !fp._exists {
		return nil
	}

	// if deleted, bail
	if fp._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM sun_file.file_post WHERE Id = ?`

	// run query
	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, fp.Id)
	}
	_, err = db.Exec(sqlstr, fp.Id)
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return err
	}

	// set deleted
	fp._deleted = true

	OnFilePost_AfterDelete(fp)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////// Querify gen - ME /////////////////////////////////////////
//.TableNameGo= table name
// _Deleter, _Updater

// orma types
type __FilePost_Deleter struct {
	wheres   []whereClause
	whereSep string
}

type __FilePost_Updater struct {
	wheres   []whereClause
	updates  map[string]interface{}
	whereSep string
}

type __FilePost_Selector struct {
	wheres    []whereClause
	selectCol string
	whereSep  string
	orderBy   string //" order by id desc //for ints
	limit     int
	offset    int
}

func NewFilePost_Deleter() *__FilePost_Deleter {
	d := __FilePost_Deleter{whereSep: " AND "}
	return &d
}

func NewFilePost_Updater() *__FilePost_Updater {
	u := __FilePost_Updater{whereSep: " AND "}
	u.updates = make(map[string]interface{}, 10)
	return &u
}

func NewFilePost_Selector() *__FilePost_Selector {
	u := __FilePost_Selector{whereSep: " AND ", selectCol: "*"}
	return &u
}

/////////////////////////////// Where for all /////////////////////////////
//// for ints all selector updater, deleter

////////ints
func (u *__FilePost_Deleter) Or() *__FilePost_Deleter {
	u.whereSep = " OR "
	return u
}

func (u *__FilePost_Deleter) Id_In(ins []int) *__FilePost_Deleter {
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

func (u *__FilePost_Deleter) Id_Ins(ins ...int) *__FilePost_Deleter {
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

func (u *__FilePost_Deleter) Id_NotIn(ins []int) *__FilePost_Deleter {
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

func (d *__FilePost_Deleter) Id_Eq(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Deleter) Id_NotEq(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Deleter) Id_LT(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Deleter) Id_LE(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Deleter) Id_GT(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Deleter) Id_GE(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__FilePost_Deleter) FileType_In(ins []int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " FileType IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__FilePost_Deleter) FileType_Ins(ins ...int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " FileType IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__FilePost_Deleter) FileType_NotIn(ins []int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " FileType NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__FilePost_Deleter) FileType_Eq(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Deleter) FileType_NotEq(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Deleter) FileType_LT(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Deleter) FileType_LE(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Deleter) FileType_GT(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Deleter) FileType_GE(val int) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__FilePost_Updater) Or() *__FilePost_Updater {
	u.whereSep = " OR "
	return u
}

func (u *__FilePost_Updater) Id_In(ins []int) *__FilePost_Updater {
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

func (u *__FilePost_Updater) Id_Ins(ins ...int) *__FilePost_Updater {
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

func (u *__FilePost_Updater) Id_NotIn(ins []int) *__FilePost_Updater {
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

func (d *__FilePost_Updater) Id_Eq(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Updater) Id_NotEq(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Updater) Id_LT(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Updater) Id_LE(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Updater) Id_GT(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Updater) Id_GE(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__FilePost_Updater) FileType_In(ins []int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " FileType IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__FilePost_Updater) FileType_Ins(ins ...int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " FileType IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__FilePost_Updater) FileType_NotIn(ins []int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " FileType NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__FilePost_Updater) FileType_Eq(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Updater) FileType_NotEq(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Updater) FileType_LT(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Updater) FileType_LE(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Updater) FileType_GT(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Updater) FileType_GE(val int) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__FilePost_Selector) Or() *__FilePost_Selector {
	u.whereSep = " OR "
	return u
}

func (u *__FilePost_Selector) Id_In(ins []int) *__FilePost_Selector {
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

func (u *__FilePost_Selector) Id_Ins(ins ...int) *__FilePost_Selector {
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

func (u *__FilePost_Selector) Id_NotIn(ins []int) *__FilePost_Selector {
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

func (d *__FilePost_Selector) Id_Eq(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Selector) Id_NotEq(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Selector) Id_LT(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Selector) Id_LE(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Selector) Id_GT(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Selector) Id_GE(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__FilePost_Selector) FileType_In(ins []int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " FileType IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__FilePost_Selector) FileType_Ins(ins ...int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " FileType IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__FilePost_Selector) FileType_NotIn(ins []int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " FileType NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__FilePost_Selector) FileType_Eq(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Selector) FileType_NotEq(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Selector) FileType_LT(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Selector) FileType_LE(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Selector) FileType_GT(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Selector) FileType_GE(val int) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " FileType >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

///// for strings //copy of above with type int -> string + rm if eq + $ms_str_cond

////////ints

func (u *__FilePost_Deleter) Extension_In(ins []string) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Extension IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__FilePost_Deleter) Extension_NotIn(ins []string) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Extension NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__FilePost_Deleter) Extension_Like(val string) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Extension LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__FilePost_Deleter) Extension_Eq(val string) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Deleter) Extension_NotEq(val string) *__FilePost_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Extension != ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints

func (u *__FilePost_Updater) Extension_In(ins []string) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Extension IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__FilePost_Updater) Extension_NotIn(ins []string) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Extension NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__FilePost_Updater) Extension_Like(val string) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Extension LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__FilePost_Updater) Extension_Eq(val string) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Updater) Extension_NotEq(val string) *__FilePost_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Extension != ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints

func (u *__FilePost_Selector) Extension_In(ins []string) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Extension IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__FilePost_Selector) Extension_NotIn(ins []string) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " Extension NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__FilePost_Selector) Extension_Like(val string) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Extension LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__FilePost_Selector) Extension_Eq(val string) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Extension = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__FilePost_Selector) Extension_NotEq(val string) *__FilePost_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Extension != ? "
	d.wheres = append(d.wheres, w)

	return d
}

/// End of wheres for selectors , updators, deletor

/////////////////////////////// Updater /////////////////////////////

//ints

func (u *__FilePost_Updater) Id(newVal int) *__FilePost_Updater {
	u.updates[" Id = ? "] = newVal
	return u
}

func (u *__FilePost_Updater) Id_Increment(count int) *__FilePost_Updater {
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

func (u *__FilePost_Updater) FileType(newVal int) *__FilePost_Updater {
	u.updates[" FileType = ? "] = newVal
	return u
}

func (u *__FilePost_Updater) FileType_Increment(count int) *__FilePost_Updater {
	if count > 0 {
		u.updates[" FileType = FileType+? "] = count
	}

	if count < 0 {
		u.updates[" FileType = FileType-? "] = -(count) //make it positive
	}

	return u
}

//string

//ints

//string
func (u *__FilePost_Updater) Extension(newVal string) *__FilePost_Updater {
	u.updates[" Extension = ? "] = newVal
	return u
}

//ints

//string

//ints

//string

/////////////////////////////////////////////////////////////////////
/////////////////////// Selector ///////////////////////////////////

//Select_* can just be used with: .GetString() , .GetStringSlice(), .GetInt() ..GetIntSlice()

func (u *__FilePost_Selector) OrderBy_Id_Desc() *__FilePost_Selector {
	u.orderBy = " ORDER BY Id DESC "
	return u
}

func (u *__FilePost_Selector) OrderBy_Id_Asc() *__FilePost_Selector {
	u.orderBy = " ORDER BY Id ASC "
	return u
}

func (u *__FilePost_Selector) Select_Id() *__FilePost_Selector {
	u.selectCol = "Id"
	return u
}

func (u *__FilePost_Selector) OrderBy_FileType_Desc() *__FilePost_Selector {
	u.orderBy = " ORDER BY FileType DESC "
	return u
}

func (u *__FilePost_Selector) OrderBy_FileType_Asc() *__FilePost_Selector {
	u.orderBy = " ORDER BY FileType ASC "
	return u
}

func (u *__FilePost_Selector) Select_FileType() *__FilePost_Selector {
	u.selectCol = "FileType"
	return u
}

func (u *__FilePost_Selector) OrderBy_Extension_Desc() *__FilePost_Selector {
	u.orderBy = " ORDER BY Extension DESC "
	return u
}

func (u *__FilePost_Selector) OrderBy_Extension_Asc() *__FilePost_Selector {
	u.orderBy = " ORDER BY Extension ASC "
	return u
}

func (u *__FilePost_Selector) Select_Extension() *__FilePost_Selector {
	u.selectCol = "Extension"
	return u
}

func (u *__FilePost_Selector) OrderBy_DataThumb_Desc() *__FilePost_Selector {
	u.orderBy = " ORDER BY DataThumb DESC "
	return u
}

func (u *__FilePost_Selector) OrderBy_DataThumb_Asc() *__FilePost_Selector {
	u.orderBy = " ORDER BY DataThumb ASC "
	return u
}

func (u *__FilePost_Selector) Select_DataThumb() *__FilePost_Selector {
	u.selectCol = "DataThumb"
	return u
}

func (u *__FilePost_Selector) OrderBy_Data_Desc() *__FilePost_Selector {
	u.orderBy = " ORDER BY Data DESC "
	return u
}

func (u *__FilePost_Selector) OrderBy_Data_Asc() *__FilePost_Selector {
	u.orderBy = " ORDER BY Data ASC "
	return u
}

func (u *__FilePost_Selector) Select_Data() *__FilePost_Selector {
	u.selectCol = "Data"
	return u
}

func (u *__FilePost_Selector) Limit(num int) *__FilePost_Selector {
	u.limit = num
	return u
}

func (u *__FilePost_Selector) Offset(num int) *__FilePost_Selector {
	u.offset = num
	return u
}

/////////////////////////  Queryer Selector  //////////////////////////////////
func (u *__FilePost_Selector) _stoSql() (string, []interface{}) {
	sqlWherrs, whereArgs := whereClusesToSql(u.wheres, u.whereSep)

	sqlstr := "SELECT " + u.selectCol + " FROM sun_file.file_post"

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

func (u *__FilePost_Selector) GetRow(db *sqlx.DB) (*FilePost, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, whereArgs)
	}

	row := &FilePost{}
	//by Sqlx
	err = db.Get(row, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return nil, err
	}

	row._exists = true

	OnFilePost_LoadOne(row)

	return row, nil
}

func (u *__FilePost_Selector) GetRows(db *sqlx.DB) ([]*FilePost, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, whereArgs)
	}

	var rows []*FilePost
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.FilePost {
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

	OnFilePost_LoadMany(rows)

	return rows, nil
}

//dep use GetRows()
func (u *__FilePost_Selector) GetRows2(db *sqlx.DB) ([]FilePost, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, whereArgs)
	}
	var rows []*FilePost
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.FilePost {
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

	OnFilePost_LoadMany(rows)

	rows2 := make([]FilePost, len(rows))
	for i := 0; i < len(rows); i++ {
		cp := *rows[i]
		rows2[i] = cp
	}

	return rows2, nil
}

func (u *__FilePost_Selector) GetString(db *sqlx.DB) (string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, whereArgs)
	}

	var res string
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return "", err
	}

	return res, nil
}

func (u *__FilePost_Selector) GetStringSlice(db *sqlx.DB) ([]string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, whereArgs)
	}
	var rows []string
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__FilePost_Selector) GetIntSlice(db *sqlx.DB) ([]int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, whereArgs)
	}
	var rows []int
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__FilePost_Selector) GetInt(db *sqlx.DB) (int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, whereArgs)
	}
	var res int
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return 0, err
	}

	return res, nil
}

/////////////////////////  Queryer Update Delete //////////////////////////////////
func (u *__FilePost_Updater) Update(db XODB) (int, error) {
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

	sqlstr := `UPDATE sun_file.file_post SET ` + sqlUpdate

	if len(strings.Trim(sqlWherrs, " ")) > 0 { //2 for safty
		sqlstr += " WHERE " + sqlWherrs
	}

	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, allArgs)
	}
	res, err := db.Exec(sqlstr, allArgs...)
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return 0, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

func (d *__FilePost_Deleter) Delete(db XODB) (int, error) {
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

	sqlstr := "DELETE FROM sun_file.file_post WHERE " + wheresStr

	// run query
	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, args)
	}
	res, err := db.Exec(sqlstr, args...)
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return 0, err
	}

	// retrieve id
	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

///////////////////////// Mass insert - replace for  FilePost ////////////////

func MassInsert_FilePost(rows []FilePost, db XODB) error {
	if len(rows) == 0 {
		return errors.New("rows slice should not be empty - inserted nothing")
	}
	var err error
	ln := len(rows)
	//s:= "(?,?,?,?,?)," //`(?, ?, ?, ?),`
	s := "(?,?,?,?,?)," //`(?, ?, ?, ?),`
	insVals_ := strings.Repeat(s, ln)
	insVals := insVals_[0 : len(insVals_)-1]
	// sql query
	sqlstr := "INSERT INTO sun_file.file_post (" +
		"Id, FileType, Extension, DataThumb, Data" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.Id)
		vals = append(vals, row.FileType)
		vals = append(vals, row.Extension)
		vals = append(vals, row.DataThumb)
		vals = append(vals, row.Data)

	}

	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, " MassInsert len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.FilePost {
			XOLogErr(err)
		}
		return err
	}

	return nil
}

func MassReplace_FilePost(rows []FilePost, db XODB) error {
	var err error
	ln := len(rows)
	s := "(?,?,?,?,?)," //`(?, ?, ?, ?),`
	insVals_ := strings.Repeat(s, ln)
	insVals := insVals_[0 : len(insVals_)-1]
	// sql query
	sqlstr := "REPLACE INTO sun_file.file_post (" +
		"Id, FileType, Extension, DataThumb, Data" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.Id)
		vals = append(vals, row.FileType)
		vals = append(vals, row.Extension)
		vals = append(vals, row.DataThumb)
		vals = append(vals, row.Data)

	}

	if LogTableSqlReq.FilePost {
		XOLog(sqlstr, " MassReplace len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.FilePost {
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

//
