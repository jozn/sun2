package main

import (
    "github.com/google/uuid"
    "fmt"
    "time"
    "encoding/binary"
)
func main()  {
    for  {
        bs := make([]byte, 4)
        ti:=time.Now().Unix()
        ti32 := uint32(ti)
        binary.LittleEndian.PutUint32(bs, ti32)
        fmt.Println(bs, ti, ti32)
        time.Sleep(time.Millisecond *1000)
        break
    }
    for  {
        fmt.Println()
        u:=uuid.NewSeq()


        fmt.Println(u, u.ClockSequence() , u.ID() , u.NodeID() , u.Time(), u.URN() )
        fmt.Println(uuid.GetTime())
        //u.ClockSequence()
        time.Sleep(time.Millisecond *200)
    }


}
