package x
// GENERATED BY XO. DO NOT EDIT.
import (
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"fmt"
	"regexp"
	"strings"
	//"time"
	 "ms/sun/helper"
    "strconv"
    "github.com/jmoiron/sqlx"
)// (shortname .TableNameGo "err" "res" "sqlstr" "db" "XOLog") -}}//(schema .Schema .Table.TableName) -}}// .TableNameGo}}// Key represents a row from 'sun.keys'.

// Manualy copy this to project
type Key__ struct {
	Key string `json:"Key"` // Key -
	// xo fields
	_exists, _deleted bool
}



// Exists determines if the Key exists in the database.
func (k *Key) Exists() bool {
	return k._exists
}

// Deleted provides information if the Key has been deleted from the database.
func (k *Key) Deleted() bool {
	return k._deleted
}

// Insert inserts the Key to the database.
func (k *Key) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if k._exists {
		return errors.New("insert failed: already exists")
	}


	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO sun.keys (` +
		`Key` +
		`) VALUES (` +
		`?` +
		`)`

	// run query
	XOLog(sqlstr, k.Key)
	_, err = db.Exec(sqlstr, k.Key)
	if err != nil {
		return err
	}

	// set existence
	k._exists = true


	OnKey_AfterInsert(k)

	return nil
}

// Insert inserts the Key to the database.
func (k *Key) Replace(db XODB) error {
	var err error

	// sql query

	const sqlstr = `REPLACE INTO sun.keys (` +
		`Key` +
		`) VALUES (` +
		`?` +
		`)`

	// run query
	XOLog(sqlstr, k.Key)
	_, err = db.Exec(sqlstr, k.Key)
	if err != nil {
		XOLogErr(err)
		return err
	}

	k._exists = true


	OnKey_AfterInsert(k)

	return nil
}

// Update updates the Key in the database.
func (k *Key) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !k._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if k._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE sun.keys SET ` +
		`` +
		` WHERE Key = ?`

	// run query
	XOLog(sqlstr, , k.Key)
	_, err = db.Exec(sqlstr, , k.Key)

	XOLogErr(err)
	OnKey_AfterUpdate(k)

	return err
}

// Save saves the Key to the database.
func (k *Key) Save(db XODB) error {
	if k.Exists() {
		return k.Update(db)
	}

	return k.Replace(db)
}

// Delete deletes the Key from the database.
func (k *Key) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !k._exists {
		return nil
	}

	// if deleted, bail
	if k._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM sun.keys WHERE Key = ?`

	// run query
	XOLog(sqlstr, k.Key)
	_, err = db.Exec(sqlstr, k.Key)
	if err != nil {
		XOLogErr(err)
		return err
	}

	// set deleted
	k._deleted = true

	OnKey_AfterDelete(k)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////// Querify gen - ME /////////////////////////////////////////
//.TableNameGo= table name
 // _Deleter, _Updater

// orma types
type __Key_Deleter struct {
	wheres   []whereClause
    whereSep string
}

type __Key_Updater struct {
	wheres   []whereClause
	updates   map[string]interface{}
    whereSep string
}

type __Key_Selector struct {
    wheres   []whereClause
    selectCol string
    whereSep  string
    orderBy string//" order by id desc //for ints
    limit int
    offset int
}

func NewKey_Deleter()  *__Key_Deleter {
	    d := __Key_Deleter {whereSep: " AND "}
	    return &d
}

func NewKey_Updater()  *__Key_Updater {
	    u := __Key_Updater {whereSep: " AND "}
	    u.updates =  make(map[string]interface{},10)
	    return &u
}

func NewKey_Selector()  *__Key_Selector {
	    u := __Key_Selector {whereSep: " AND ",selectCol: "*"}
	    return &u
}
/////////////////////////////// Where for all /////////////////////////////
//// for ints all selector updater, deleter

		
			////////ints
func (u *__Key_Deleter) Or () *__Key_Deleter {
    u.whereSep = " OR "
    return u
}


		
			////////ints
func (u *__Key_Updater) Or () *__Key_Updater {
    u.whereSep = " OR "
    return u
}


		
			////////ints
func (u *__Key_Selector) Or () *__Key_Selector {
    u.whereSep = " OR "
    return u
}




///// for strings //copy of above with type int -> string + rm if eq + $ms_str_cond

		
			////////ints

func (u *__Key_Deleter) Key_In (ins []string) *__Key_Deleter {
    w := whereClause{}
    var insWhere []interface{}
    for _, i:= range ins {
        insWhere = append(insWhere,i)
    }
    w.args = insWhere
    w.condition = " Key IN("+helper.DbQuestionForSqlIn(len(ins))+") "
    u.wheres = append(u.wheres, w)

    return u
}

func (u *__Key_Deleter) Key_NotIn (ins []string) *__Key_Deleter {
    w := whereClause{}
    var insWhere []interface{}
    for _, i:= range ins {
        insWhere = append(insWhere,i)
    }
    w.args = insWhere
    w.condition = " Key NOT IN("+helper.DbQuestionForSqlIn(len(ins))+") "
    u.wheres = append(u.wheres, w)

    return u
}

//must be used like: UserName_like("hamid%")
func (u *__Key_Deleter) Key_Like (val string) *__Key_Deleter {
    w := whereClause{}
    var insWhere []interface{}
    insWhere = append(insWhere,val)
    w.args = insWhere
    w.condition = " Key LIKE ? "
    u.wheres = append(u.wheres, w)

    return u
}

func (d *__Key_Deleter) Key_Eq (val string) *__Key_Deleter {
    w := whereClause{}
    var insWhere []interface{}
    insWhere = append(insWhere,val)
    w.args = insWhere
    w.condition = " Key = ? "
    d.wheres = append(d.wheres, w)

    return d
}

func (d *__Key_Deleter) Key_NotEq (val string) *__Key_Deleter {
    w := whereClause{}
    var insWhere []interface{}
    insWhere = append(insWhere,val)
    w.args = insWhere
    w.condition = " Key != ? "
    d.wheres = append(d.wheres, w)

    return d
}


		
			////////ints

func (u *__Key_Updater) Key_In (ins []string) *__Key_Updater {
    w := whereClause{}
    var insWhere []interface{}
    for _, i:= range ins {
        insWhere = append(insWhere,i)
    }
    w.args = insWhere
    w.condition = " Key IN("+helper.DbQuestionForSqlIn(len(ins))+") "
    u.wheres = append(u.wheres, w)

    return u
}

func (u *__Key_Updater) Key_NotIn (ins []string) *__Key_Updater {
    w := whereClause{}
    var insWhere []interface{}
    for _, i:= range ins {
        insWhere = append(insWhere,i)
    }
    w.args = insWhere
    w.condition = " Key NOT IN("+helper.DbQuestionForSqlIn(len(ins))+") "
    u.wheres = append(u.wheres, w)

    return u
}

//must be used like: UserName_like("hamid%")
func (u *__Key_Updater) Key_Like (val string) *__Key_Updater {
    w := whereClause{}
    var insWhere []interface{}
    insWhere = append(insWhere,val)
    w.args = insWhere
    w.condition = " Key LIKE ? "
    u.wheres = append(u.wheres, w)

    return u
}

func (d *__Key_Updater) Key_Eq (val string) *__Key_Updater {
    w := whereClause{}
    var insWhere []interface{}
    insWhere = append(insWhere,val)
    w.args = insWhere
    w.condition = " Key = ? "
    d.wheres = append(d.wheres, w)

    return d
}

func (d *__Key_Updater) Key_NotEq (val string) *__Key_Updater {
    w := whereClause{}
    var insWhere []interface{}
    insWhere = append(insWhere,val)
    w.args = insWhere
    w.condition = " Key != ? "
    d.wheres = append(d.wheres, w)

    return d
}


		
			////////ints

func (u *__Key_Selector) Key_In (ins []string) *__Key_Selector {
    w := whereClause{}
    var insWhere []interface{}
    for _, i:= range ins {
        insWhere = append(insWhere,i)
    }
    w.args = insWhere
    w.condition = " Key IN("+helper.DbQuestionForSqlIn(len(ins))+") "
    u.wheres = append(u.wheres, w)

    return u
}

func (u *__Key_Selector) Key_NotIn (ins []string) *__Key_Selector {
    w := whereClause{}
    var insWhere []interface{}
    for _, i:= range ins {
        insWhere = append(insWhere,i)
    }
    w.args = insWhere
    w.condition = " Key NOT IN("+helper.DbQuestionForSqlIn(len(ins))+") "
    u.wheres = append(u.wheres, w)

    return u
}

//must be used like: UserName_like("hamid%")
func (u *__Key_Selector) Key_Like (val string) *__Key_Selector {
    w := whereClause{}
    var insWhere []interface{}
    insWhere = append(insWhere,val)
    w.args = insWhere
    w.condition = " Key LIKE ? "
    u.wheres = append(u.wheres, w)

    return u
}

func (d *__Key_Selector) Key_Eq (val string) *__Key_Selector {
    w := whereClause{}
    var insWhere []interface{}
    insWhere = append(insWhere,val)
    w.args = insWhere
    w.condition = " Key = ? "
    d.wheres = append(d.wheres, w)

    return d
}

func (d *__Key_Selector) Key_NotEq (val string) *__Key_Selector {
    w := whereClause{}
    var insWhere []interface{}
    insWhere = append(insWhere,val)
    w.args = insWhere
    w.condition = " Key != ? "
    d.wheres = append(d.wheres, w)

    return d
}


/// End of wheres for selectors , updators, deletor





/////////////////////////////// Updater /////////////////////////////



	//ints

	//string
func (u *__Key_Updater)Key (newVal string) *__Key_Updater {
    u.updates[" Key = ? "] = newVal
    return u
}


/////////////////////////////////////////////////////////////////////
/////////////////////// Selector ///////////////////////////////////


//Select_* can just be used with: .GetString() , .GetStringSlice(), .GetInt() ..GetIntSlice()

func (u *__Key_Selector) OrderBy_Key_Desc () *__Key_Selector {
    u.orderBy = " ORDER BY Key DESC "
    return u
}

func (u *__Key_Selector) OrderBy_Key_Asc () *__Key_Selector {
    u.orderBy = " ORDER BY Key ASC "
    return u
}

func (u *__Key_Selector) Select_Key () *__Key_Selector {
    u.selectCol = "Key"
    return u
}

func (u *__Key_Selector) Limit(num int) *__Key_Selector {
    u.limit = num
    return u
}

func (u *__Key_Selector) Offset(num int) *__Key_Selector {
    u.offset = num
    return u
}


/////////////////////////  Queryer Selector  //////////////////////////////////
func (u *__Key_Selector)_stoSql ()  (string,[]interface{}) {
	sqlWherrs, whereArgs := whereClusesToSql(u.wheres,u.whereSep)

	sqlstr := "SELECT " +u.selectCol +" FROM sun.keys"

	if len( strings.Trim(sqlWherrs," ") ) > 0 {//2 for safty
		sqlstr += " WHERE "+ sqlWherrs
	}

	if u.orderBy != ""{
        sqlstr += u.orderBy
    }

    if u.limit != 0 {
        sqlstr += " LIMIT " + strconv.Itoa(u.limit)
    }

    if u.offset != 0 {
        sqlstr += " OFFSET " + strconv.Itoa(u.offset)
    }
    return sqlstr,whereArgs
}

func (u *__Key_Selector) GetRow (db *sqlx.DB) (*Key,error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	XOLog(sqlstr,whereArgs )

	row := &Key{}
	//by Sqlx
	err = db.Get(row ,sqlstr, whereArgs...)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	row._exists = true

	OnKey_LoadOne(row)

	return row, nil
}

func (u *__Key_Selector) GetRows (db *sqlx.DB) ([]*Key,error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	XOLog(sqlstr,whereArgs )

	var rows []*Key
	//by Sqlx
	err = db.Unsafe().Select(&rows ,sqlstr, whereArgs...)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	/*for i:=0;i< len(rows);i++ {
		rows[i]._exists = true
	}*/

	for i:=0;i< len(rows);i++ {
		rows[i]._exists = true
	}

	OnKey_LoadMany(rows)

	return rows, nil
}

//dep use GetRows()
func (u *__Key_Selector) GetRows2 (db *sqlx.DB) ([]Key,error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	XOLog(sqlstr,whereArgs )

	var rows []*Key
	//by Sqlx
	err = db.Unsafe().Select(&rows ,sqlstr, whereArgs...)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	/*for i:=0;i< len(rows);i++ {
		rows[i]._exists = true
	}*/

	for i:=0;i< len(rows);i++ {
		rows[i]._exists = true
	}

	OnKey_LoadMany(rows)

	rows2 := make([]Key, len(rows))
	for i:=0;i< len(rows);i++ {
		cp := *rows[i]
		rows2[i]= cp
	}

	return rows2, nil
}



func (u *__Key_Selector) GetString (db *sqlx.DB) (string,error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	XOLog(sqlstr,whereArgs )

	var res string
	//by Sqlx
	err = db.Get(&res ,sqlstr, whereArgs...)
	if err != nil {
		XOLogErr(err)
		return "", err
	}

	return res, nil
}

func (u *__Key_Selector) GetStringSlice (db *sqlx.DB) ([]string,error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	XOLog(sqlstr,whereArgs )

	var rows []string
	//by Sqlx
	err = db.Select(&rows ,sqlstr, whereArgs...)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	return rows, nil
}

func (u *__Key_Selector) GetIntSlice (db *sqlx.DB) ([]int,error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	XOLog(sqlstr,whereArgs )

	var rows []int
	//by Sqlx
	err = db.Select(&rows ,sqlstr, whereArgs...)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	return rows, nil
}

func (u *__Key_Selector) GetInt (db *sqlx.DB) (int,error) {
	var err error

	sqlstr, whereArgs := u._stoSql()

	XOLog(sqlstr,whereArgs )

	var res int
	//by Sqlx
	err = db.Get(&res ,sqlstr, whereArgs...)
	if err != nil {
		XOLogErr(err)
		return 0, err
	}

	return res, nil
}

/////////////////////////  Queryer Update Delete //////////////////////////////////
func (u *__Key_Updater)Update (db XODB) (int,error) {
    var err error

    var updateArgs []interface{}
    var sqlUpdateArr  []string
    for up, newVal := range u.updates {
        sqlUpdateArr = append(sqlUpdateArr, up)
        updateArgs = append(updateArgs, newVal)
    }
    sqlUpdate:= strings.Join(sqlUpdateArr, ",")

    sqlWherrs, whereArgs := whereClusesToSql(u.wheres,u.whereSep)

    var allArgs []interface{}
    allArgs = append(allArgs, updateArgs...)
    allArgs = append(allArgs, whereArgs...)

    sqlstr := `UPDATE sun.keys SET ` + sqlUpdate

    if len( strings.Trim(sqlWherrs," ") ) > 0 {//2 for safty
		sqlstr += " WHERE " +sqlWherrs
	}

    XOLog(sqlstr,allArgs)
    res, err := db.Exec(sqlstr, allArgs...)
    if err != nil {
    	XOLogErr(err)
        return 0,err
    }

    num, err := res.RowsAffected()
    if err != nil {
    	XOLogErr(err)
        return 0,err
    }

    return int(num),nil
}

func (d *__Key_Deleter)Delete (db XODB) (int,error) {
    var err error
    var wheresArr []string
    for _,w := range d.wheres{
        wheresArr = append(wheresArr,w.condition)
    }
    wheresStr := strings.Join(wheresArr, d.whereSep)

    var args []interface{}
    for _,w := range d.wheres{
        args = append(args,w.args...)
    }

    sqlstr := "DELETE FROM sun.keys WHERE " + wheresStr

    // run query
    XOLog(sqlstr, args)
    res, err := db.Exec(sqlstr, args...)
    if err != nil {
    	XOLogErr(err)
        return 0,err
    }

    // retrieve id
    num, err := res.RowsAffected()
    if err != nil {
    	XOLogErr(err)
        return 0,err
    }

    return int(num),nil
}

///////////////////////// Mass insert - replace for  Key ////////////////

func MassInsert_Key(rows []Key ,db XODB) error {
	if len(rows) == 0 {
		return errors.New("rows slice should not be empty - inserted nothing")
	}
	var err error
	ln := len(rows)
	//s:= "(?)," //`(?, ?, ?, ?),`
	s:= "(?)," //`(?, ?, ?, ?),`
	insVals_:= strings.Repeat(s, ln)
	insVals := insVals_[0:len(insVals_)-1]
	// sql query
	sqlstr := "INSERT INTO sun.keys (" +
		"Key" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{},0, ln * 5)//5 fields

	for _,row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.Key) 

	}

	XOLog(sqlstr, " MassInsert len = ", ln, vals)

	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		XOLogErr(err)
		return err
	}

	return nil
}

func MassReplace_Key(rows []Key ,db XODB) error {
	var err error
	ln := len(rows)
	s:= "(?)," //`(?, ?, ?, ?),`
	insVals_:= strings.Repeat(s, ln)
	insVals := insVals_[0:len(insVals_)-1]
	// sql query
	sqlstr := "REPLACE INTO sun.keys (" +
		"Key" +
		") VALUES " + insVals

	// run query
	vals := make([]interface{},0, ln * 5)//5 fields

	for _,row := range rows {
		// vals = append(vals,row.UserId)
		vals = append(vals, row.Key) 

	}

	XOLog(sqlstr, " MassReplace len = ", ln , vals)

	_, err = db.Exec(sqlstr, vals...)
	if err != nil {
		XOLogErr(err)
		return err
	}

	return nil
}



//////////////////// Play ///////////////////////////////

			//

