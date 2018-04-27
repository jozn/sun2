package pusher_service

import "ms/sun/servises/log_service"

var logPipes = log_service.NewSimpleLoggerWithExtension("pipe", ".gol")
var logRpc = log_service.NewSimpleLoggerWithExtension("rpc", ".gol")
var logHttpRpc = log_service.NewSimpleLoggerWithExtension("rpc_http", ".gol")
var logChatSync = log_service.NewSimpleLoggerWithExtension("chat_sync", ".gol")

func init() {
	go loadChatPushersLoop()
	go routePushChatLoop()
	go allPipesMap.cleanUnConnectedUsersLoop()
}
