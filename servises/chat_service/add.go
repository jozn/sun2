package chat_service

import (
	"ms/sun/helper"
	"ms/sun2/servises/chat_service/message_adder"
	"ms/sun2/servises/file_service"
	"ms/sun2/shared/x"
)

func ChatAddNewMsg(UserId int, msg *x.PB_MessageView, blob []byte) {
	if len(blob) > 0 && msg.MessageFileView != nil {
		cassRow := file_service.Row{
			Id:        helper.NextRowsSeqId(),
			Extension: msg.MessageFileView.Extension,
			Data:      blob,
		}

		file_service.SaveMsgFile(cassRow)
		msg.MessageFileView.MessageFileId = int64(cassRow.Id)
	}
	message_adder.MessageAdderChanel <- msg
}
