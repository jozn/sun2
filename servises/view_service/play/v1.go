package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun2/servises/memcache_service"
    "ms/sun2/servises/view_service"
)

func main() {
	var err error
	base.DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/ms?charset=utf8mb4")
	//DB, err = sqlx.Connect("mysql", "root:123456@localhost:3307/ms?charset=utf8mb4")
	//DB.Exec("SET NAMES 'utf8mb4';")
	helper.NoErr(err)
	base.DB.MapperFunc(func(s string) string { return s })
	res := view_service.Action_GetLastsViews(6, 1, 100, 0)
	helper.PertyPrint(res)

	r := memcache_service.DoesUserFollows(6, 753)
	fmt.Println(r)

}
