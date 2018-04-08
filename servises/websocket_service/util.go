package websocket_service

import (
	"github.com/golang/protobuf/proto"
	"ms/sun_old/helper"
	"ms/sun/shared/x"
)

func newPB_CommandToClient_WithData(cmd string, protoMsg proto.Message) x.PB_CommandToClient {
	m := x.PB_CommandToClient{
		ServerCallId: int64(helper.RandomSeqUid()), //int64(time.Now().UnixNano()),
		Command:      cmd,
	}
	bytes, err := proto.Marshal(protoMsg)
	if err == nil {
		m.Data = bytes
	} else {
		helper.DebugPrintln("ERROR : proto.Marshal newPB_CommandToClient_WithData, ", err)
	}
	return m
}
