package websocket_service

import (
	"ms/sun2/servises/log_service"
)

const PB_ResponseToClient = "PB_ResponseToClient"

var logPipes = log_service.NewSimpleLoggerWithExtension("pipe", ".gol")
var logRpc = log_service.NewSimpleLoggerWithExtension("rpc", ".gol")
var logHttpRpc = log_service.NewSimpleLoggerWithExtension("rpc_http", ".gol")

var AllPipesMap *pipesMap
var commandsToClientMapRegistery *commandsToClientMap

func init() {
	AllPipesMap = &pipesMap{
		mp: make(map[int]*UserDevicePipe, 100),
	}

	commandsToClientMapRegistery = &commandsToClientMap{
		mp: make(map[int64]commandToClientCallback, 100),
	}
	go intervalRunCommandsCallbacksTimeOutChecker(commandsToClientMapRegistery)
}
