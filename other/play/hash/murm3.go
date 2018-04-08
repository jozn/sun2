package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"log"
	"ms/sun/shared/helper"
)

func main() {
	f1()
}

func f2()  {
    //mp := make(map[string]int)
    mp := make(map[int]string)

    for i := 0; i < 10000000; i++ {
        s := fmt.Sprintf("%d", helper.NextRowsSeqId())
        i := int(murmur3.Sum64([]byte(s)))
        v, ok := mp[i]
        if ok {
            log.Panic(v, s, i)
        }
        mp[i] = ""//s
        //fmt.Println(i)
    }

    //fmt.Println(mp)
}

func f1()  {
    //mp := make(map[string]int)
    mp := make(map[int]string)

    for i := 0; i < 10000000; i++ {
        s := fmt.Sprintf("%d_liked_%d", helper.NextRowsSeqId(), helper.NextRowsSeqId())
        j := int(murmur3.Sum64([]byte(s)))
        v, ok := mp[j]
        if ok {
            log.Panic(v, s, j)
        }
        mp[j] = "" //s
        //j = j >>1
        fmt.Println(j)
    }

    fmt.Println(len(mp))
    //fmt.Println(mp)
}
