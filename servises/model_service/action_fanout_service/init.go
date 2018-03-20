package action_fanout_service

func init() {
	go listernAndSaverEvents()
	go saveActionFansLooper()
}
