package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"log"
	"ms/sun/helper"
)

func main() {
	f2()
}

func f2()  {
    //mp := make(map[string]int)
    mp := make(map[int]string)

    for i := 0; i < 100; i++ {
        s := fmt.Sprintf("%d", helper.NextRowsSeqId())
        i := int(murmur3.Sum64([]byte(s)))
        v, ok := mp[i]
        if ok {
            log.Panic(v, s, i)
        }
        mp[i] = ""//s
        //fmt.Println(i)
    }

    fmt.Println(mp)
}

func f1()  {
    //mp := make(map[string]int)
    mp := make(map[int]string)

    for i := 0; i < 1000; i++ {
        s := fmt.Sprintf("%d_liked_%d", helper.NextRowsSeqId(), helper.NextRowsSeqId())
        i := int(murmur3.Sum64([]byte(s)))
        v, ok := mp[i]
        if ok {
            log.Panic(v, s, i)
        }
        mp[i] = ""//s
        //fmt.Println(i)
    }

    fmt.Println(mp)
}
