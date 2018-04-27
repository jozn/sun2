package base

import (
	"github.com/jmoiron/sqlx"
	"ms/sun_old/helper"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sqlx.DB
var __DEV__ bool = true

type AppErr error

func init()  {
    DefultConnectToMysql()
}

func DefultConnectToMysql() {
    var err error
	DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/sun?charset=utf8mb4")
	helper.NoErr(err)
	DB.MapperFunc(func(s string) string { return s })


}
