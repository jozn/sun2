package websocket_service

import (
	"github.com/golang/protobuf/proto"
	"ms/sun/helper"
	"ms/sun2/shared/x"
)

func NewPB_CommandToClient(cmd string) x.PB_CommandToClient {
	p := x.PB_CommandToClient{
		ServerCallId: int64(helper.RandomSeqUid()), //int64(time.Now().UnixNano()),
		Command:      cmd,
	}
	return p
}

func NewPB_CommandToClient_WithData(cmd string, protoMsg proto.Message) x.PB_CommandToClient {
	m := NewPB_CommandToClient(cmd)
	bytes, err := proto.Marshal(protoMsg)
	if err == nil {
		m.Data = bytes
	} else {
		helper.DebugPrintln("ERROR : proto.Marshal NewPB_CommandToClient_WithData, ", err)
	}
	return m
}
