package message_adder

import (
	"github.com/jozn/protobuf/proto"
	"ms/sun_old/base"
	"ms/sun/shared/helper"
	"ms/sun/servises/log_service"
	"ms/sun/servises/sun_utils"
	"ms/sun/shared/config"
	"ms/sun/shared/x"
)
//todo save MessageFiles to its tables for later retrival
var chatLogger = log_service.NewSimpleLogger("chat")

type chatDirect struct {
	User1Id        int
	User2Id        int
	User1ChatKey   string
	User2ChatKey   string
	User1Chat      *x.Chat
	User2Chat      *x.Chat
	AddChan        chan *x.PB_MessageView
	Err            error
	lastActiveTime int
}

func newChaatDirectByRoomKey(RoomKey string) *chatDirect {
	u1, u2 := sun_utils.RoomKeyToUsers(RoomKey)
	res := &chatDirect{
		User1Id:        u1,
		User2Id:        u2,
		User1ChatKey:   sun_utils.UsersToChatKey(u1, u2),
		User2ChatKey:   sun_utils.UsersToChatKey(u2, u1),
		AddChan:        make(chan *x.PB_MessageView, 100),
		lastActiveTime: helper.TimeNow(),
	}

	res.loadOrCreateRooms()
	go res.listenForAdding()

	return res
}

func (s *chatDirect) loadOrCreateRooms() error {
	if s.User1Chat != nil && s.User2Chat != nil {
		return nil
	}
	var e1, e2 error
	s.User1Chat, e1 = getOrCreateDirectChatForPeers(s.User1Id, s.User2Id)
	s.User2Chat, e2 = getOrCreateDirectChatForPeers(s.User2Id, s.User1Id)
	if e1 != nil {
		s.Err = e1
	}

	if e2 != nil {
		s.Err = e2
	}

	if config.IS_DEBUG && s.Err != nil {
		chatLogger.Printf("%s - Err: %s - %v", "loadOrCreateRooms() has error: ", s.Err, s)
	}
	return s.Err
}

func (s *chatDirect) listenForAdding() {
	defer helper.JustRecover()
	for msgPb := range s.AddChan {
		if s.Err != nil {
			continue
		}
		s.lastActiveTime = helper.TimeNow()
		adderUserId := int(msgPb.UserId)
		peerUserId := s.User1Id
		if adderUserId == s.User1Id {
			peerUserId = s.User2Id
		}
		extraPb := &x.PB_MessageTableExtra{
			MessageFileView: msgPb.MessageFileView,
		}
		extraBlob, _ := proto.Marshal(extraPb)
		delivery := 4 //int(x.RoomMessageDeliviryStatusEnum_SENT)
		m1 := x.DirectMessage{
			ChatKey:            s.User1ChatKey,
			MessageId:          int(msgPb.MessageId),
			RoomKey:            msgPb.RoomKey,
			UserId:             int(msgPb.UserId),
			MessageFileId:      int(msgPb.MessageFileId),
			MessageTypeEnum:    int(msgPb.MessageTypeEnum),
			Text:               msgPb.Text,
			CreatedTime:        int(msgPb.CreatedTime),
			Seq:                s.User1Chat.Seq + 1,
			DeliviryStatusEnum: int(delivery),
			ExtraPB:            extraBlob,
		}
		m1.Save(base.DB)

		m2 := x.DirectMessage{
			ChatKey:            s.User2ChatKey,
			MessageId:          int(msgPb.MessageId),
			RoomKey:            msgPb.RoomKey,
			UserId:             int(msgPb.UserId),
			MessageFileId:      int(msgPb.MessageFileId),
			MessageTypeEnum:    int(msgPb.MessageTypeEnum),
			Text:               msgPb.Text,
			CreatedTime:        int(msgPb.CreatedTime),
			Seq:                int(msgPb.Seq),
			DeliviryStatusEnum: int(delivery),
			ExtraPB:            extraBlob,
		}
		m2.Save(base.DB)

		s.User1Chat.Seq += 1
		s.User1Chat.UpdatedMs += helper.TimeNowMs()

		s.User2Chat.Seq += 1
		s.User2Chat.UpdatedMs += helper.TimeNowMs()

		s.User1Chat.Update(base.DB)
		s.User2Chat.Update(base.DB)

		upToMe := &x.ChatSync{
			SyncId:         helper.NextRowsSeqId(),
			ToUserId:       adderUserId,
			ChatSyncTypeId: 3, //chat_service.CHAT_SYNC_MSG_RECIVED_TO_SERVER,
			RoomKey:        msgPb.RoomKey,
			ChatKey:        "",
			MessageId:      int(msgPb.MessageId),
			MessagePb:      []byte{},
			MessageJson:    "",
			CreatedTime:    helper.TimeNow(),
		}

		msgBlob, _ := proto.Marshal(msgPb)
		upToPeer := &x.ChatSync{
			SyncId:         helper.NextRowsSeqId(),
			ToUserId:       peerUserId,
			ChatSyncTypeId: 2, //chat_service.CHAT_SYNC_NEW_MESSAGE,
			RoomKey:        msgPb.RoomKey,
			ChatKey:        "",
			MessageId:      int(msgPb.MessageId),
			MessagePb:      msgBlob,
			MessageJson:    "",
			CreatedTime:    helper.TimeNow(),
		}

		upToMe.Save(base.DB)
		upToPeer.Save(base.DB)
	}
}

func getOrCreateDirectChatForPeers(me int, peer int) (*x.Chat, error) {
	chatMe, err := x.NewChat_Selector().UserId_Eq(me).PeerUserId_Eq(peer).GetRow(base.DB)
	if err != nil {
		chatMe = &x.Chat{
			ChatKey:      sun_utils.UsersToChatKey(me, peer),
			RoomKey:      sun_utils.UsersToRoomKey(me, peer),
			RoomTypeEnum: 1,
			UserId:       me,
			PeerUserId:   peer,
			GroupId:      0,
			CreatedTime:  helper.TimeNow(),
			Seq:          0,
			SeenSeq:      0,
			UpdatedMs:    helper.TimeNowMs(),
		}
		err = chatMe.Save(base.DB)
		if err != nil {
			return nil, err
		}
	}

	return chatMe, nil
}
