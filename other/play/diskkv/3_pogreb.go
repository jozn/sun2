package main

import (
    "log"

    "github.com/akrylysov/pogreb"
    "ms/sun/shared/helper"
    "fmt"
)

func main() {
    db, err := pogreb.Open("pogreb.test1", nil)
    if err != nil {
        log.Fatal(err)
        return
    }
    write := func() {
        for i := 0; i < 100000; i++ {
            err := db.Put([]byte("key"+helper.IntToStr(i)), []byte(fmt.Sprintf("%d",helper.NanoRowIdSeq())))
            //err := db.Put([]byte("key"+helper.IntToStr(i)), []byte(i)))
            helper.NoErr(err)
            if i% 10000 == 0{
                fmt.Println(i)
            }
        }
    }
    read := func() {
        for i := 0; i < 1000000; i++ {
            res,err := db.Get([]byte("key"+helper.IntToStr(i)))
            helper.NoErr(err)
            if i% 10000 == 0{
                fmt.Println(string(res))
            }
        }
    }

    _ = write
    helper.PertyPrint(db.Metrics())
    write()
    helper.PertyPrint(db.Metrics())
    read()
    helper.PertyPrint(db.Metrics())
    defer db.Close()
}