package chat_service

import (
	"github.com/jozn/protobuf/proto"
	"ms/sun_old/base"
	"ms/sun/servises/view_service"
	"ms/sun/shared/x"
)

//todo: support last msgView
func GetChatList(UserId int) (res []*x.PB_ChatView) {
	chats, err := x.NewChat_Selector().UserId_Eq(UserId).OrderBy_UpdatedMs_Desc().GetRows(base.DB)
	if err != nil {
		return
	}

	for _, chat := range chats {
		v := &x.PB_ChatView{
			ChatKey:      chat.ChatKey,
			RoomKey:      chat.RoomKey,
			RoomTypeEnum: int32(chat.RoomTypeEnum),
			UserId:       int32(chat.UserId),
			PeerUserId:   int32(chat.PeerUserId),
			GroupId:      int64(chat.GroupId),
			CreatedTime:  int32(chat.CreatedTime),
			Seq:          int32(chat.Seq),
			SeenSeq:      int32(chat.SeenSeq),
			UpdatedMs:    int64(chat.UpdatedMs),
		}
		v.UserView = view_service.UserViewAndMe(chat.PeerUserId, UserId)

		res = append(res, v)
	}

	return
}

func ChatViewsForChatList(UserId int, RoomKeys []string) (res []*x.PB_ChatView) {
    if len(RoomKeys) == 0 {
        return
    }
    chats, err := x.NewChat_Selector().UserId_Eq(UserId).RoomKey_In(RoomKeys).OrderBy_UpdatedMs_Desc().GetRows(base.DB)
    if err != nil {
        return
    }

    for _, chat := range chats {
        v := &x.PB_ChatView{
            ChatKey:      chat.ChatKey,
            RoomKey:      chat.RoomKey,
            RoomTypeEnum: int32(chat.RoomTypeEnum),
            UserId:       int32(chat.UserId),
            PeerUserId:   int32(chat.PeerUserId),
            GroupId:      int64(chat.GroupId),
            CreatedTime:  int32(chat.CreatedTime),
            Seq:          int32(chat.Seq),
            SeenSeq:      int32(chat.SeenSeq),
            UpdatedMs:    int64(chat.UpdatedMs),
        }
        v.UserView = view_service.UserViewAndMe(chat.PeerUserId, UserId)

        res = append(res, v)
    }

    return
}

func GetMessageList(UserId int, ChatKey string, FromTopMessageId, Limit int) (msgs []*x.PB_MessageView, more bool) {
	rows, err := x.NewDirectMessage_Selector().ChatKey_Eq(ChatKey).OrderBy_MessageId_Desc().
		Limit(Limit + 1).
		GetRows(base.DB)
	if err != nil {
		return
	}

	if len(rows) > Limit {
		more = true
	}

	for i := 0; i < len(rows)-1; i++ {
		m := rows[i]
		v := &x.PB_MessageView{
			ChatKey:            m.ChatKey,
			MessageId:          int64(m.MessageId),
			RoomKey:            m.RoomKey,
			UserId:             int32(m.UserId),
			MessageFileId:      int64(m.MessageFileId),
			MessageTypeEnum:    int32(m.MessageTypeEnum),
			Text:               m.Text,
			CreatedTime:        int32(m.CreatedTime),
			Seq:                int32(m.Seq),
			DeliviryStatusEnum: int32(m.DeliviryStatusEnum),
		}
		extra := &x.PB_MessageTableExtra{}
		err := proto.Unmarshal(m.ExtraPB, extra)
		if err == nil {
			v.MessageFileView = extra.MessageFileView
		}
		//v.UserView = view_service.UserViewAndMe(row.PeerUserId, UserId)

		msgs = append(msgs, v)
	}

	return
}
