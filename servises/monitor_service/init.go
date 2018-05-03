package monitor_service

func init()  {
    saveMetericLoop()
    go addRpChanHanle_go()
}
