package main

import (
    "time"
    "ms/sun/shared/helper"
    "fmt"
    //"github.com/tonnerre/golang-pretty"
    "github.com/kr/pretty"
    "github.com/davecgh/go-spew/spew"
)

type MonitorAllTime struct {
    TotalMysqlRowsLoaded               int
    TotalChatRowsLoaded                int
    TotalSocialRowsLoaded              int
    TotalClosedPipes                   int
    TotalWebSocketsClosed              int
    TotalWebSocketsReceivedText        int
    TotalWebSocketsReceivedBinary      int
    TotalWebSocketsReceivedPanics      int
    TotalWebSocketsSendPanics          int
    TotalRowsProceedWiltActiveUser     int
    TotalChatRowsProceedWiltActiveUser int
    TotalUsersAdded                    int
    CurrentActiveUsers                 int
    CurrentActivePipes                 int
    ActiveTimeSeconds                  int
}

func (k MonitorAllTime)Ajlk()  {

}

func main() {

    go func() {
        for i := 1; i < 100; i += 5 {
            m := MonitorAllTime{
                TotalSocialRowsLoaded: 5000000000000,

            }
            helper.PertyPrintNew(m)
            //fmt.Printf("%+%v \n", pretty.Formatter(m))
            fmt.Printf("8 %#+ v 22\n", m)
            pretty.Print(m)
            pretty.Print("\n==========\n")
            //fmt.Printf("%#+ v", pretty.Formatter(m))
            //pretty.Logln(m)
            //c:= spew.C
            //spew.Config.DisableMethods = true
            spew.Config.DisablePointerAddresses = true
            spew.Config.DisablePointerMethods = true
            spew.Dump(m)
            time.Sleep(time.Minute)
        }
    }()

    time.Sleep(time.Hour)
}
