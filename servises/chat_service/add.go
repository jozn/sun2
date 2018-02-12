package chat_service

import (
	"ms/sun2/servises/chat_service/message_adder"
	"ms/sun2/shared/x"
)

func ChatAddNewMsg(UserId int, msg *x.PB_MessageView, blob []byte) {
	message_adder.MessageAdderChanel <- msg
}
