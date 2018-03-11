package model_service

func init() {
	go loadPostKeysLoop()
}
