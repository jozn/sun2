package x

// GENERATED BY XO. DO NOT EDIT.
import (
	"errors"
	"strings"
	//"time"
	"ms/sun/shared/helper"
	"strconv"

	"github.com/jmoiron/sqlx"
) // (shortname .TableNameGo "err" "res" "sqlstr" "db" "XOLog") -}}//(schema .Schema .Table.TableName) -}}// .TableNameGo}}// XfileServiceMetricLog represents a row from 'sun_log.xfile_service_metric_log'.

// Manualy copy this to project
type XfileServiceMetricLog__ struct {
	Id         int    `json:"Id"`         // Id -
	InstanceId int    `json:"InstanceId"` // InstanceId -
	MetricJson string `json:"MetricJson"` // MetricJson -
	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XfileServiceMetricLog exists in the database.
func (xsml *XfileServiceMetricLog) Exists() bool {
	return xsml._exists
}

// Deleted provides information if the XfileServiceMetricLog has been deleted from the database.
func (xsml *XfileServiceMetricLog) Deleted() bool {
	return xsml._deleted
}

// Insert inserts the XfileServiceMetricLog to the database.
func (xsml *XfileServiceMetricLog) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xsml._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO sun_log.xfile_service_metric_log (` +
		`Id, InstanceId, MetricJson` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, xsml.Id, xsml.InstanceId, xsml.MetricJson)
	}
	_, err = db.Exec(sqlstr, xsml.Id, xsml.InstanceId, xsml.MetricJson)
	if err != nil {
		return err
	}

	// set existence
	xsml._exists = true

	OnXfileServiceMetricLog_AfterInsert(xsml)

	return nil
}

// Insert inserts the XfileServiceMetricLog to the database.
func (xsml *XfileServiceMetricLog) Replace(db XODB) error {
	var err error

	// sql query

	const sqlstr = `REPLACE INTO sun_log.xfile_service_metric_log (` +
		`Id, InstanceId, MetricJson` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, xsml.Id, xsml.InstanceId, xsml.MetricJson)
	}
	_, err = db.Exec(sqlstr, xsml.Id, xsml.InstanceId, xsml.MetricJson)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return err
	}

	xsml._exists = true

	OnXfileServiceMetricLog_AfterInsert(xsml)

	return nil
}

// Update updates the XfileServiceMetricLog in the database.
func (xsml *XfileServiceMetricLog) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xsml._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xsml._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE sun_log.xfile_service_metric_log SET ` +
		`InstanceId = ?, MetricJson = ?` +
		` WHERE Id = ?`

	// run query
	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, xsml.InstanceId, xsml.MetricJson, xsml.Id)
	}
	_, err = db.Exec(sqlstr, xsml.InstanceId, xsml.MetricJson, xsml.Id)

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLogErr(err)
	}
	OnXfileServiceMetricLog_AfterUpdate(xsml)

	return err
}

// Save saves the XfileServiceMetricLog to the database.
func (xsml *XfileServiceMetricLog) Save(db XODB) error {
	if xsml.Exists() {
		return xsml.Update(db)
	}

	return xsml.Replace(db)
}

// Delete deletes the XfileServiceMetricLog from the database.
func (xsml *XfileServiceMetricLog) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xsml._exists {
		return nil
	}

	// if deleted, bail
	if xsml._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM sun_log.xfile_service_metric_log WHERE Id = ?`

	// run query
	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, xsml.Id)
	}
	_, err = db.Exec(sqlstr, xsml.Id)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return err
	}

	// set deleted
	xsml._deleted = true

	OnXfileServiceMetricLog_AfterDelete(xsml)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////// Querify gen - ME /////////////////////////////////////////
//.TableNameGo= table name
// _Deleter, _Updater

// orma types
type __XfileServiceMetricLog_Deleter struct {
	wheres   []whereClause
	whereSep string
}

type __XfileServiceMetricLog_Updater struct {
	wheres   []whereClause
	updates  map[string]interface{}
	whereSep string
}

type __XfileServiceMetricLog_Selector struct {
	wheres    []whereClause
	selectCol string
	whereSep  string
	orderBy   string //" order by id desc //for ints
	limit     int
	offset    int
}

func NewXfileServiceMetricLog_Deleter() *__XfileServiceMetricLog_Deleter {
	d := __XfileServiceMetricLog_Deleter{whereSep: " AND "}
	return &d
}

func NewXfileServiceMetricLog_Updater() *__XfileServiceMetricLog_Updater {
	u := __XfileServiceMetricLog_Updater{whereSep: " AND "}
	u.updates = make(map[string]interface{}, 10)
	return &u
}

func NewXfileServiceMetricLog_Selector() *__XfileServiceMetricLog_Selector {
	u := __XfileServiceMetricLog_Selector{whereSep: " AND ", selectCol: "*"}
	return &u
}

/////////////////////////////// Where for all /////////////////////////////
//// for ints all selector updater, deleter

////////ints
func (u *__XfileServiceMetricLog_Deleter) Or() *__XfileServiceMetricLog_Deleter {
	u.whereSep = " OR "
	return u
}

func (u *__XfileServiceMetricLog_Deleter) Id_In(ins []int) *__XfileServiceMetricLog_Deleter {
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

func (u *__XfileServiceMetricLog_Deleter) Id_Ins(ins ...int) *__XfileServiceMetricLog_Deleter {
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

func (u *__XfileServiceMetricLog_Deleter) Id_NotIn(ins []int) *__XfileServiceMetricLog_Deleter {
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

func (d *__XfileServiceMetricLog_Deleter) Id_Eq(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Deleter) Id_NotEq(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Deleter) Id_LT(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Deleter) Id_LE(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Deleter) Id_GT(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Deleter) Id_GE(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__XfileServiceMetricLog_Deleter) InstanceId_In(ins []int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " InstanceId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__XfileServiceMetricLog_Deleter) InstanceId_Ins(ins ...int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " InstanceId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__XfileServiceMetricLog_Deleter) InstanceId_NotIn(ins []int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " InstanceId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__XfileServiceMetricLog_Deleter) InstanceId_Eq(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Deleter) InstanceId_NotEq(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Deleter) InstanceId_LT(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Deleter) InstanceId_LE(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Deleter) InstanceId_GT(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Deleter) InstanceId_GE(val int) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__XfileServiceMetricLog_Updater) Or() *__XfileServiceMetricLog_Updater {
	u.whereSep = " OR "
	return u
}

func (u *__XfileServiceMetricLog_Updater) Id_In(ins []int) *__XfileServiceMetricLog_Updater {
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

func (u *__XfileServiceMetricLog_Updater) Id_Ins(ins ...int) *__XfileServiceMetricLog_Updater {
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

func (u *__XfileServiceMetricLog_Updater) Id_NotIn(ins []int) *__XfileServiceMetricLog_Updater {
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

func (d *__XfileServiceMetricLog_Updater) Id_Eq(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Updater) Id_NotEq(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Updater) Id_LT(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Updater) Id_LE(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Updater) Id_GT(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Updater) Id_GE(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__XfileServiceMetricLog_Updater) InstanceId_In(ins []int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " InstanceId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__XfileServiceMetricLog_Updater) InstanceId_Ins(ins ...int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " InstanceId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__XfileServiceMetricLog_Updater) InstanceId_NotIn(ins []int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " InstanceId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__XfileServiceMetricLog_Updater) InstanceId_Eq(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Updater) InstanceId_NotEq(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Updater) InstanceId_LT(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Updater) InstanceId_LE(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Updater) InstanceId_GT(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Updater) InstanceId_GE(val int) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints
func (u *__XfileServiceMetricLog_Selector) Or() *__XfileServiceMetricLog_Selector {
	u.whereSep = " OR "
	return u
}

func (u *__XfileServiceMetricLog_Selector) Id_In(ins []int) *__XfileServiceMetricLog_Selector {
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

func (u *__XfileServiceMetricLog_Selector) Id_Ins(ins ...int) *__XfileServiceMetricLog_Selector {
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

func (u *__XfileServiceMetricLog_Selector) Id_NotIn(ins []int) *__XfileServiceMetricLog_Selector {
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

func (d *__XfileServiceMetricLog_Selector) Id_Eq(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Selector) Id_NotEq(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Selector) Id_LT(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Selector) Id_LE(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Selector) Id_GT(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Selector) Id_GE(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " Id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (u *__XfileServiceMetricLog_Selector) InstanceId_In(ins []int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " InstanceId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__XfileServiceMetricLog_Selector) InstanceId_Ins(ins ...int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " InstanceId IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__XfileServiceMetricLog_Selector) InstanceId_NotIn(ins []int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " InstanceId NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__XfileServiceMetricLog_Selector) InstanceId_Eq(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Selector) InstanceId_NotEq(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId != ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Selector) InstanceId_LT(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId < ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Selector) InstanceId_LE(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Selector) InstanceId_GT(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId > ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Selector) InstanceId_GE(val int) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " InstanceId >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

///// for strings //copy of above with type int -> string + rm if eq + $ms_str_cond

////////ints

func (u *__XfileServiceMetricLog_Deleter) MetricJson_In(ins []string) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " MetricJson IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__XfileServiceMetricLog_Deleter) MetricJson_NotIn(ins []string) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " MetricJson NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__XfileServiceMetricLog_Deleter) MetricJson_Like(val string) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " MetricJson LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__XfileServiceMetricLog_Deleter) MetricJson_Eq(val string) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " MetricJson = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Deleter) MetricJson_NotEq(val string) *__XfileServiceMetricLog_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " MetricJson != ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints

func (u *__XfileServiceMetricLog_Updater) MetricJson_In(ins []string) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " MetricJson IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__XfileServiceMetricLog_Updater) MetricJson_NotIn(ins []string) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " MetricJson NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__XfileServiceMetricLog_Updater) MetricJson_Like(val string) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " MetricJson LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__XfileServiceMetricLog_Updater) MetricJson_Eq(val string) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " MetricJson = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Updater) MetricJson_NotEq(val string) *__XfileServiceMetricLog_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " MetricJson != ? "
	d.wheres = append(d.wheres, w)

	return d
}

////////ints

func (u *__XfileServiceMetricLog_Selector) MetricJson_In(ins []string) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " MetricJson IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

func (u *__XfileServiceMetricLog_Selector) MetricJson_NotIn(ins []string) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, i := range ins {
		insWhere = append(insWhere, i)
	}
	w.args = insWhere
	w.condition = " MetricJson NOT IN(" + helper.DbQuestionForSqlIn(len(ins)) + ") "
	u.wheres = append(u.wheres, w)

	return u
}

//must be used like: UserName_like("hamid%")
func (u *__XfileServiceMetricLog_Selector) MetricJson_Like(val string) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " MetricJson LIKE ? "
	u.wheres = append(u.wheres, w)

	return u
}

func (d *__XfileServiceMetricLog_Selector) MetricJson_Eq(val string) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " MetricJson = ? "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *__XfileServiceMetricLog_Selector) MetricJson_NotEq(val string) *__XfileServiceMetricLog_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " MetricJson != ? "
	d.wheres = append(d.wheres, w)

	return d
}

/// End of wheres for selectors , updators, deletor

/////////////////////////////// Updater /////////////////////////////

//ints

func (u *__XfileServiceMetricLog_Updater) Id(newVal int) *__XfileServiceMetricLog_Updater {
	u.updates[" Id = ? "] = newVal
	return u
}

func (u *__XfileServiceMetricLog_Updater) Id_Increment(count int) *__XfileServiceMetricLog_Updater {
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

func (u *__XfileServiceMetricLog_Updater) InstanceId(newVal int) *__XfileServiceMetricLog_Updater {
	u.updates[" InstanceId = ? "] = newVal
	return u
}

func (u *__XfileServiceMetricLog_Updater) InstanceId_Increment(count int) *__XfileServiceMetricLog_Updater {
	if count > 0 {
		u.updates[" InstanceId = InstanceId+? "] = count
	}

	if count < 0 {
		u.updates[" InstanceId = InstanceId-? "] = -(count) //make it positive
	}

	return u
}

//string

//ints

//string
func (u *__XfileServiceMetricLog_Updater) MetricJson(newVal string) *__XfileServiceMetricLog_Updater {
	u.updates[" MetricJson = ? "] = newVal
	return u
}

/////////////////////////////////////////////////////////////////////
/////////////////////// Selector ///////////////////////////////////

//Select_* can just be used with: .GetString() , .GetStringSlice(), .GetInt() ..GetIntSlice()

func (u *__XfileServiceMetricLog_Selector) OrderBy_Id_Desc() *__XfileServiceMetricLog_Selector {
	u.orderBy = " ORDER BY Id DESC "
	return u
}

func (u *__XfileServiceMetricLog_Selector) OrderBy_Id_Asc() *__XfileServiceMetricLog_Selector {
	u.orderBy = " ORDER BY Id ASC "
	return u
}

func (u *__XfileServiceMetricLog_Selector) Select_Id() *__XfileServiceMetricLog_Selector {
	u.selectCol = "Id"
	return u
}

func (u *__XfileServiceMetricLog_Selector) OrderBy_InstanceId_Desc() *__XfileServiceMetricLog_Selector {
	u.orderBy = " ORDER BY InstanceId DESC "
	return u
}

func (u *__XfileServiceMetricLog_Selector) OrderBy_InstanceId_Asc() *__XfileServiceMetricLog_Selector {
	u.orderBy = " ORDER BY InstanceId ASC "
	return u
}

func (u *__XfileServiceMetricLog_Selector) Select_InstanceId() *__XfileServiceMetricLog_Selector {
	u.selectCol = "InstanceId"
	return u
}

func (u *__XfileServiceMetricLog_Selector) OrderBy_MetricJson_Desc() *__XfileServiceMetricLog_Selector {
	u.orderBy = " ORDER BY MetricJson DESC "
	return u
}

func (u *__XfileServiceMetricLog_Selector) OrderBy_MetricJson_Asc() *__XfileServiceMetricLog_Selector {
	u.orderBy = " ORDER BY MetricJson ASC "
	return u
}

func (u *__XfileServiceMetricLog_Selector) Select_MetricJson() *__XfileServiceMetricLog_Selector {
	u.selectCol = "MetricJson"
	return u
}

func (u *__XfileServiceMetricLog_Selector) Limit(num int) *__XfileServiceMetricLog_Selector {
	u.limit = num
	return u
}

func (u *__XfileServiceMetricLog_Selector) Offset(num int) *__XfileServiceMetricLog_Selector {
	u.offset = num
	return u
}

func (u *__XfileServiceMetricLog_Selector) Order_Rand() *__XfileServiceMetricLog_Selector {
	u.orderBy = " ORDER BY RAND() "
	return u
}

/////////////////////////  Queryer Selector  //////////////////////////////////
func (u *__XfileServiceMetricLog_Selector) _stoSql() (string, []interface{}) {
	sqlWherrs, whereArgs := whereClusesToSql(u.wheres, u.whereSep)

	sqlstr := "SELECT " + u.selectCol + " FROM sun_log.xfile_service_metric_log"

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

func (u *__XfileServiceMetricLog_Selector) GetRow(db *sqlx.DB) (*XfileServiceMetricLog, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, whereArgs)
	}

	row := &XfileServiceMetricLog{}
	//by Sqlx
	err = db.Get(row, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return nil, err
	}

	row._exists = true

	OnXfileServiceMetricLog_LoadOne(row)

	return row, nil
}

func (u *__XfileServiceMetricLog_Selector) GetRows(db *sqlx.DB) ([]*XfileServiceMetricLog, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, whereArgs)
	}

	var rows []*XfileServiceMetricLog
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
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

	OnXfileServiceMetricLog_LoadMany(rows)

	return rows, nil
}

//dep use GetRows()
func (u *__XfileServiceMetricLog_Selector) GetRows2(db *sqlx.DB) ([]XfileServiceMetricLog, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, whereArgs)
	}
	var rows []*XfileServiceMetricLog
	//by Sqlx
	err = db.Unsafe().Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
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

	OnXfileServiceMetricLog_LoadMany(rows)

	rows2 := make([]XfileServiceMetricLog, len(rows))
	for i := 0; i < len(rows); i++ {
		cp := *rows[i]
		rows2[i] = cp
	}

	return rows2, nil
}

func (u *__XfileServiceMetricLog_Selector) GetString(db *sqlx.DB) (string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, whereArgs)
	}

	var res string
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return "", err
	}

	return res, nil
}

func (u *__XfileServiceMetricLog_Selector) GetStringSlice(db *sqlx.DB) ([]string, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, whereArgs)
	}
	var rows []string
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__XfileServiceMetricLog_Selector) GetIntSlice(db *sqlx.DB) ([]int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, whereArgs)
	}
	var rows []int
	//by Sqlx
	err = db.Select(&rows, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return nil, err
	}

	return rows, nil
}

func (u *__XfileServiceMetricLog_Selector) GetInt(db *sqlx.DB) (int, error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, whereArgs)
	}
	var res int
	//by Sqlx
	err = db.Get(&res, sqlstr, whereArgs...)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return 0, err
	}

	return res, nil
}

/////////////////////////  Queryer Update Delete //////////////////////////////////
func (u *__XfileServiceMetricLog_Updater) Update(db XODB) (int, error) {
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

	sqlstr := `UPDATE sun_log.xfile_service_metric_log SET ` + sqlUpdate

	if len(strings.Trim(sqlWherrs, " ")) > 0 { //2 for safty
		sqlstr += " WHERE " + sqlWherrs
	}

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, allArgs)
	}
	res, err := db.Exec(sqlstr, allArgs...)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return 0, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

func (d *__XfileServiceMetricLog_Deleter) Delete(db XODB) (int, error) {
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

	sqlstr := "DELETE FROM sun_log.xfile_service_metric_log WHERE " + wheresStr

	// run query
	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, args)
	}
	res, err := db.Exec(sqlstr, args...)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return 0, err
	}

	// retrieve id
	num, err := res.RowsAffected()
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return 0, err
	}

	return int(num), nil
}

///////////////////////// Mass insert - replace for  XfileServiceMetricLog ////////////////

func MassInsert_XfileServiceMetricLog(rows []XfileServiceMetricLog, db XODB) error {
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
	sqlstr := "INSERT INTO sun_log.xfile_service_metric_log (" +
		"Id, InstanceId, MetricJson" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.Id)
		vals = append(vals, row.InstanceId)
		vals = append(vals, row.MetricJson)

	}

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, " MassInsert len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
			XOLogErr(err)
		}
		return err
	}

	return nil
}

func MassReplace_XfileServiceMetricLog(rows []XfileServiceMetricLog, db XODB) error {
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
	sqlstr := "REPLACE INTO sun_log.xfile_service_metric_log (" +
		"Id, InstanceId, MetricJson" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{}, 0, ln*5) //5 fields

	for _, row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.Id)
		vals = append(vals, row.InstanceId)
		vals = append(vals, row.MetricJson)

	}

	if LogTableSqlReq.XfileServiceMetricLog {
		XOLog(sqlstr, " MassReplace len = ", ln, vals)
	}
	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		if LogTableSqlReq.XfileServiceMetricLog {
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
