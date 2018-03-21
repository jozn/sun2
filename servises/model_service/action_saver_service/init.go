package action_saver_service

func init() {
	go saveNewActions()
	go listernAndSaverActions()
}
