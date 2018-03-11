package home_fanout_service

func init() {
	go listernAndSaverNotifys()
	go saveHomeFansLooper()
}
