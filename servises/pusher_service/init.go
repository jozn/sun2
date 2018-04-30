package pusher_service

import "ms/sun/servises/log_service"

var logPushLoader = log_service.NewSimpleLoggerWithExtension("push_loader", ".gol")

var logPush = log_service.NewSimpleLoggerWithExtension("push", ".gol")
var logPushPipes = log_service.NewSimpleLoggerWithExtension("push_pipe", ".gol")
//var logRpc = log_service.NewSimpleLoggerWithExtension("rpc", ".gol")
//var logHttpRpc = log_service.NewSimpleLoggerWithExtension("rpc_http", ".gol")
//var logChatSync = log_service.NewSimpleLoggerWithExtension("chat_sync", ".gol")

func init() {
	go loadChatPushersLoop()
	go routePushChatLoop()
	go allPipesMap.cleanUnConnectedUsersLoop()
}
