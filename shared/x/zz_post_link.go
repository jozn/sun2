package x

// GENERATED BY XO. DO NOT EDIT.
import (
	"errors"
	"strings"
	//"time"
	"ms/sun/shared/helper"
	"strconv"

	"github.com/jmoiron/sqlx"
) // (shortname .TableNameGo "err" "res" "sqlstr" "db" "XOLog") -}}//(schema .Schema .Table.TableName) -}}// .TableNameGo}}// PostLink represents a row from 'sun.post_link'.

// Manualy copy this to project
type PostLink__ struct {
	LinkId  int    `json:"LinkId"`  // LinkId -
	LinkUrl string `json:"LinkUrl"` // LinkUrl -
	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PostLink exists in the database.
func (pl *PostLink) Exists() bool {
	return pl._exists
}

// Deleted provides information if the PostLink has been deleted from the database.
func (pl *PostLink) Deleted() bool {
	return pl._deleted
}

// Insert inserts the PostLink to the database.
func (pl *PostLink) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pl._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO sun.post_link (` +
		`LinkId, LinkUrl` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, pl.LinkId, pl.LinkUrl)
	}
	_, err = db.Exec(sqlstr, pl.LinkId, pl.LinkUrl)
	if err != nil {
		return err
	}

	// set existence
	pl._exists = true

	OnPostLink_AfterInsert(pl)

	return nil
}

// Insert inserts the PostLink to the database.
func (pl *PostLink) Replace(db XODB) error {
	var err error

	// sql query

	const sqlstr = `REPLACE INTO sun.post_link (` +
		`LinkId, LinkUrl` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, pl.LinkId, pl.LinkUrl)
	}
	_, err = db.Exec(sqlstr, pl.LinkId, pl.LinkUrl)
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return err
	}

	pl._exists = true

	OnPostLink_AfterInsert(pl)

	return nil
}

// Update updates the PostLink in the database.
func (pl *PostLink) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pl._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if pl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE sun.post_link SET ` +
		`LinkUrl = ?` +
		` WHERE LinkId = ?`

	// run query
	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, pl.LinkUrl, pl.LinkId)
	}
	_, err = db.Exec(sqlstr, pl.LinkUrl, pl.LinkId)

	if LogTableSqlReq.PostLink {
		XOLogErr(err)
	}
	OnPostLink_AfterUpdate(pl)

	return err
}

// Save saves the PostLink to the database.
func (pl *PostLink) Save(db XODB) error {
	if pl.Exists() {
		return pl.Update(db)
	}

	return pl.Replace(db)
}

// Delete deletes the PostLink from the database.
func (pl *PostLink) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pl._exists {
		return nil
	}

	// if deleted, bail
	if pl._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM sun.post_link WHERE LinkId = ?`

	// run query
	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, pl.LinkId)
	}
	_, err = db.Exec(sqlstr, pl.LinkId)
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return err
	}

	// set deleted
	pl._deleted = true

	OnPostLink_AfterDelete(pl)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////// Querify gen - ME /////////////////////////////////////////
//.TableNameGo= table name
// _Deleter, _Updater

// orma types
type __PostLink_Deleter struct {
	wheres   []whereClause
	whereSep string
}

type __PostLink_Updater struct {
	wheres   []whereClause
	updates  map[string]interface{}
	whereSep string
}

type __PostLink_Selector struct {
	wheres    []whereClause
	selectCol string
	whereSep  string
	orderBy   string //" order by id desc //for ints
	limit     int
	offset    int
}

func NewPostLink_Deleter() *__PostLink_Deleter {
	d := __PostLink_Deleter{whereSep: " AND "}
	return &d
}

func NewPostLink_Updater() *__PostLink_Updater {
	u := __PostLink_Updater{whereSep: " AND "}
	u.updates = make(map[string]interface{}, 10)
	return &u
}

func NewPostLink_Selector() *__PostLink_Selector {
	u := __PostLink_Selector{whereSep: " AND ", selectCol: "*"}
	return &u
}

/////////////////////////////// Where for all /////////////////////////////
//// for ints all selector updater, deleter

////////ints
func (u *__PostLink_Deleter) Or() *__PostLink_Deleter {
	u.whereSep = " OR "
	return u
}

func (u *__PostLink_Deleter) LinkId_In(ins []int) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostLink_Deleter) LinkId_Ins(ins ...int) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostLink_Deleter) LinkId_NotIn(ins []int) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__PostLink_Deleter) LinkId_Eq(val int) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Deleter) LinkId_NotEq(val int) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Deleter) LinkId_LT(val int) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Deleter) LinkId_LE(val int) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Deleter) LinkId_GT(val int) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Deleter) LinkId_GE(val int) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__PostLink_Updater) Or() *__PostLink_Updater {
	u.whereSep = " OR "
	return u
}

func (u *__PostLink_Updater) LinkId_In(ins []int) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostLink_Updater) LinkId_Ins(ins ...int) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostLink_Updater) LinkId_NotIn(ins []int) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__PostLink_Updater) LinkId_Eq(val int) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Updater) LinkId_NotEq(val int) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Updater) LinkId_LT(val int) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Updater) LinkId_LE(val int) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Updater) LinkId_GT(val int) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Updater) LinkId_GE(val int) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__PostLink_Selector) Or() *__PostLink_Selector {
	u.whereSep = " OR "
	return u
}

func (u *__PostLink_Selector) LinkId_In(ins []int) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostLink_Selector) LinkId_Ins(ins ...int) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostLink_Selector) LinkId_NotIn(ins []int) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__PostLink_Selector) LinkId_Eq(val int) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Selector) LinkId_NotEq(val int) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Selector) LinkId_LT(val int) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Selector) LinkId_LE(val int) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Selector) LinkId_GT(val int) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Selector) LinkId_GE(val int) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

///// for strings //copy of above with type int -> string + rm if eq + $ms_str_cond

////////ints

func (u *__PostLink_Deleter) LinkUrl_In(ins []string) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkUrl IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostLink_Deleter) LinkUrl_NotIn(ins []string) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkUrl NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__PostLink_Deleter) LinkUrl_Like(val string) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkUrl LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__PostLink_Deleter) LinkUrl_Eq(val string) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkUrl = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Deleter) LinkUrl_NotEq(val string) *__PostLink_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkUrl != ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints

func (u *__PostLink_Updater) LinkUrl_In(ins []string) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkUrl IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostLink_Updater) LinkUrl_NotIn(ins []string) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkUrl NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__PostLink_Updater) LinkUrl_Like(val string) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkUrl LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__PostLink_Updater) LinkUrl_Eq(val string) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkUrl = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Updater) LinkUrl_NotEq(val string) *__PostLink_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkUrl != ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints

func (u *__PostLink_Selector) LinkUrl_In(ins []string) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkUrl IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__PostLink_Selector) LinkUrl_NotIn(ins []string) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " LinkUrl NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__PostLink_Selector) LinkUrl_Like(val string) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkUrl LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__PostLink_Selector) LinkUrl_Eq(val string) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkUrl = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__PostLink_Selector) LinkUrl_NotEq(val string) *__PostLink_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " LinkUrl != ? "
	d.wheres = append(d.wheres, w)

	return d
}

/// End of wheres for selectors , updators, deletor

/////////////////////////////// Updater /////////////////////////////

//ints

func (u *__PostLink_Updater) LinkId(newVal int) *__PostLink_Updater {
	u.updates[" LinkId = ? "] = newVal
	return u
}

func (u *__PostLink_Updater) LinkId_Increment(count int) *__PostLink_Updater {
	if count > 0 {
		u.updates[" LinkId = LinkId+? "] = count
	}

	if count < 0 {
		u.updates[" LinkId = LinkId-? "] = -(count) //make it positive
	}

	return u
}

//string

//ints

//string
func (u *__PostLink_Updater) LinkUrl(newVal string) *__PostLink_Updater {
	u.updates[" LinkUrl = ? "] = newVal
	return u
}

/////////////////////////////////////////////////////////////////////
/////////////////////// Selector ///////////////////////////////////

//Select_* can just be used with: .GetString() , .GetStringSlice(), .GetInt() ..GetIntSlice()

func (u *__PostLink_Selector) OrderBy_LinkId_Desc() *__PostLink_Selector {
	u.orderBy = " ORDER BY LinkId DESC "
	return u
}

func (u *__PostLink_Selector) OrderBy_LinkId_Asc() *__PostLink_Selector {
	u.orderBy = " ORDER BY LinkId ASC "
	return u
}

func (u *__PostLink_Selector) Select_LinkId() *__PostLink_Selector {
	u.selectCol = "LinkId"
	return u
}

func (u *__PostLink_Selector) OrderBy_LinkUrl_Desc() *__PostLink_Selector {
	u.orderBy = " ORDER BY LinkUrl DESC "
	return u
}

func (u *__PostLink_Selector) OrderBy_LinkUrl_Asc() *__PostLink_Selector {
	u.orderBy = " ORDER BY LinkUrl ASC "
	return u
}

func (u *__PostLink_Selector) Select_LinkUrl() *__PostLink_Selector {
	u.selectCol = "LinkUrl"
	return u
}

func (u *__PostLink_Selector) Limit(num int) *__PostLink_Selector {
	u.limit = num
	return u
}

func (u *__PostLink_Selector) Offset(num int) *__PostLink_Selector {
	u.offset = num
	return u
}

func (u *__PostLink_Selector) Order_Rand() *__PostLink_Selector {
	u.orderBy = " ORDER BY RAND() "
	return u
}

/////////////////////////  Queryer Selector  //////////////////////////////////
func (u *__PostLink_Selector) _stoSql() (string, []interface{}) {
	sqlWherrs, whereArgs := whereClusesToSql(u.wheres, u.whereSep)

	sqlstr := "SELECT " + u.selectCol + " FROM sun.post_link"

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

func (u *__PostLink_Selector) GetRow(db *sqlx.DB) (*PostLink, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, whereArgs)
	}

	row := &PostLink{}
	//by Sqlx
	err = db.Get(row, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return nil, err
	}

	row._exists = true

	OnPostLink_LoadOne(row)

	return row, nil
}

func (u *__PostLink_Selector) GetRows(db *sqlx.DB) ([]*PostLink, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, whereArgs)
	}

	var rows []*PostLink
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostLink {
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

	OnPostLink_LoadMany(rows)

	return rows, nil
}

//dep use GetRows()
func (u *__PostLink_Selector) GetRows2(db *sqlx.DB) ([]PostLink, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, whereArgs)
	}
	var rows []*PostLink
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostLink {
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

	OnPostLink_LoadMany(rows)

	rows2 := make([]PostLink, len(rows))
	for i := 0; i < len(rows); i++ {
		cp := *rows[i]
		rows2[i] = cp
	}

	return rows2, nil
}

func (u *__PostLink_Selector) GetString(db *sqlx.DB) (string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, whereArgs)
	}

	var res string
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return "", err
	}

	return res, nil
}

func (u *__PostLink_Selector) GetStringSlice(db *sqlx.DB) ([]string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, whereArgs)
	}
	var rows []string
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__PostLink_Selector) GetIntSlice(db *sqlx.DB) ([]int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, whereArgs)
	}
	var rows []int
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__PostLink_Selector) GetInt(db *sqlx.DB) (int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, whereArgs)
	}
	var res int
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return 0, err
	}

	return res, nil
}

/////////////////////////  Queryer Update Delete //////////////////////////////////
func (u *__PostLink_Updater) Update(db XODB) (int, error) {
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

	sqlstr := `UPDATE sun.post_link SET ` + sqlUpdate

	if len(strings.Trim(sqlWherrs, " ")) > 0 { //2 for safty
		sqlstr += " WHERE " + sqlWherrs
	}

	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, allArgs)
	}
	res, err := db.Exec(sqlstr, allArgs...)
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return 0, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

func (d *__PostLink_Deleter) Delete(db XODB) (int, error) {
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

	sqlstr := "DELETE FROM sun.post_link WHERE " + wheresStr

	// run query
	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, args)
	}
	res, err := db.Exec(sqlstr, args...)
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return 0, err
	}

	// retrieve id
	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

///////////////////////// Mass insert - replace for  PostLink ////////////////

func MassInsert_PostLink(rows []PostLink, db XODB) error {
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
	sqlstr := "INSERT INTO sun.post_link (" +
		"LinkId, LinkUrl" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.LinkId)
		vals = append(vals, row.LinkUrl)

	}

	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, " MassInsert len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return err
	}

	return nil
}

func MassReplace_PostLink(rows []PostLink, db XODB) error {
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
	sqlstr := "REPLACE INTO sun.post_link (" +
		"LinkId, LinkUrl" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.LinkId)
		vals = append(vals, row.LinkUrl)

	}

	if LogTableSqlReq.PostLink {
		XOLog(sqlstr, " MassReplace len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.PostLink {
			XOLogErr(err)
		}
		return err
	}

	return nil

}

//////////////////// Play ///////////////////////////////

//

//
