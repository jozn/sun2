package action_saver_service

func init() {
	go saveNewActiones()
	go listernAndSaverActions()
}
