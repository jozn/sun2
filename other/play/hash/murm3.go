package main

import (
    "github.com/spaolacci/murmur3"
    "fmt"
    "ms/sun/helper"
    "log"
)
func main()  {
    //mp := make(map[string]int)
    mp := make(map[int]string)

    for i:=0; i< 1000000 ; i++  {
        s := fmt.Sprintf("%d_liked_%d", helper.NextRowsSeqId(), helper.NextRowsSeqId())
        i:=int(murmur3.Sum64([]byte(s)))
        v,ok := mp[i]
        if ok {
            log.Panic(v, s, i)
        }
        mp[i] = s
        //fmt.Println(i)
    }

}
