package notify_saver_service

func init()  {
    go saveNewNotifyes()
    go listernAndSaverNotifys()
}