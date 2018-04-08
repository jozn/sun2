package main

import (
	"math/rand"
	"ms/sun_old/helper"
	"ms/sun/servises/chat_service"
	"ms/sun/servises/sun_utils"
	"ms/sun/shared/x"
    "ms/sun/scripts/facts/fact_utils"
)

func main() {
	for i := 0; i < 1000; i++ {
		if rand.Intn(10) <= 6 {
            factMsgText()
		} else {
            factMsgPhoto()
		}
	}
}

func factMsgText() {
	u1 := fact_utils.GetRandUserId()
	u2 := fact_utils.GetRandUserId()
	msg := &x.PB_MessageView{
		ChatKey:            sun_utils.UsersToChatKey(u1, u2),
		MessageId:          int64(helper.NextRowsSeqId()),
		RoomKey:            sun_utils.UsersToRoomKey(u1, u2),
		UserId:             int32(u1),
		MessageFileId:      0,
		MessageTypeEnum:    int32(1),
		Text:               helper.FactRandStrEmoji(100, true),
		CreatedTime:        int32(helper.TimeNow()),
		Seq:                int32(0),
		DeliviryStatusEnum: int32(1),
	}
	chat_service.ChatAddNewMsg(u1, msg, []byte{})
}

func factMsgPhoto() {
	u1 := fact_utils.GetRandUserId()
	u2 := fact_utils.GetRandUserId()
	msg := &x.PB_MessageView{
		ChatKey:            sun_utils.UsersToChatKey(u1, u2),
		MessageId:          int64(helper.NextRowsSeqId()),
		RoomKey:            sun_utils.UsersToRoomKey(u1, u2),
		UserId:             int32(u1),
		MessageFileId:      0,
		MessageTypeEnum:    int32(1),
		Text:               helper.FactRandStrEmoji(100, true),
		CreatedTime:        int32(helper.TimeNow()),
		Seq:                int32(0),
		DeliviryStatusEnum: int32(1),
	}
	_,bt,_ := fact_utils.RandImage()
	msgFile := &x.PB_MessageFileView{
        MessageFileId: int64(helper.NextRowsSeqId()),
        FileTypeEnum:  int32(2),
        Size:          int32(len(bt)),
        Width:         int32(500),
        Height:        int32(350),
        Duration:      int32(0),
        Extension:     ".jpg",
    }
    msg.MessageFileId = msgFile.MessageFileId
    msg.MessageFileView = msgFile
	chat_service.ChatAddNewMsg(u1, msg, bt)
}
