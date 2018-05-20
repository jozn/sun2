package dbs

import (
    "github.com/jmoiron/sqlx"
    "fmt"
    "strings"
    _ "github.com/lib/pq"
)
var DB_PG *sqlx.DB
func init()  {
    var err error
    DB_PG, err = sqlx.Connect("postgres", "postgresql://root@localhost:26257?sslmode=disable")
    fmt.Println(DB_PG, err)
    //on PG we must lowercase coulmns names unlike the Myql which is upper case
    DB_PG.MapperFunc(func(s string) string { return strings.ToLower(s) })
    DB_PG = DB_PG.Unsafe()
}