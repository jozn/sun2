package chat_service

import (
	"ms/sun/shared/helper"
	"ms/sun/servises/chat_service/message_adder"
	"ms/sun/servises/file_service_old"
	"ms/sun/shared/x"
)

func ChatAddNewMsg(UserId int, msg *x.PB_MessageView, blob []byte) {
	if len(blob) > 0 && msg.MessageFileView != nil {
		cassRow := file_service_old.Row{
			Id:        helper.NextRowsSeqId(),
			Extension: msg.MessageFileView.Extension,
			Data:      blob,
		}

		file_service_old.SaveMsgFile(cassRow)
		msg.MessageFileView.MessageFileId = int64(cassRow.Id)
	}
	message_adder.MessageAdderChanel <- msg
}
