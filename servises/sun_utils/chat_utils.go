package sun_utils

import (
	"errors"
	"fmt"
	"ms/sun_old/config"
	"ms/sun/shared/helper"
	"strings"
)

//RoomKey; d24_6
//ChatKey: d6:24
func KeyNewMessageKey(UserId int, RoomKey string) string {
	//return fmt.Sprintf("%d_%d_%s", UserId, helper.TimeNow(), helper.RandString(4)) //todo extrac this to client
	return fmt.Sprintf("%s|%d|%s|%d", RoomKey, helper.TimeNowMs(), helper.RandString(3), UserId)
}

//"d25_56" "d56:24" ////not yet "g6_15646548_cdc" "g6:6_15646548_cdc"
func RoomKeyToOtherUser(roomKey string, userId int) int {
	key := strings.Replace(roomKey, "d", "", -1)
	key = strings.Replace(key, "g", "", -1)
	key = strings.Replace(key, ":", "_", -1)

	parts := strings.Split(key, "_")
	if len(parts) != 2 {
		return 0
	}

	i1 := helper.StrToInt(parts[0], 0)
	i2 := helper.StrToInt(parts[1], 0)

	if i1 == userId {
		return i2
	}

	if i2 == userId {
		return i1
	}

	return 0
}

func RoomKeyToUsers(roomKey string) (User1Id, User2Id int) {
	key := strings.Replace(roomKey, "d", "", -1)
	key = strings.Replace(key, "g", "", -1)
	key = strings.Replace(key, ":", "_", -1)

	parts := strings.Split(key, "_")
	if len(parts) != 2 {
		return 0, 0
	}

	User1Id = helper.StrToInt(parts[0], 0)
	User2Id = helper.StrToInt(parts[1], 0)

	return
}

//d(lower userId)_(HigherUserId)
func UsersToRoomKey(me int, peer int) string {
	if me > peer {
		me, peer = peer, me
	}
	return fmt.Sprintf("d%d_%d", me, peer)
}

func UsersToChatKey(me int, peer int) string {
	return fmt.Sprintf("d%d:%d", me, peer)
}

func RoomKeyToChatKey(RoomKey string) (ChatKey string, err error) {
	arr := strings.Split(RoomKey, "_")
	if len(arr) == 2 && (RoomKey[0] == 'd' || RoomKey[0] == 'g') {
		u1 := helper.StrToInt(arr[0][1:], 0)
		u2 := helper.StrToInt(arr[1], 0)
		if u1 > 0 && u2 > 0 {
			return UsersToChatKey(u1, u2), nil
		}
	}
	if config.IS_DEBUG {
		panic("WRONG Keys_RoomKeyToChatKey param for RoomKey: " + RoomKey)
	}
	return "", errors.New("WRONG RoomKey ")
}

func IsMyChatKey(ChatKey string, MyUserId int) bool {
	if len(ChatKey) > 0 && ChatKey[0] == 'd' {
		str := ChatKey[1:strings.Index(ChatKey, ":")]
		uid := helper.StrToInt(str, 0)
		return uid == MyUserId
	}
	return false
}

func ChatKeyToRoomKey(ChatKey string) (RoomKey string, err error) {
	arr := strings.Split(ChatKey, ":")
	if len(arr) == 2 && (ChatKey[0] == 'd' || ChatKey[0] == 'g') {
		u1 := helper.StrToInt(arr[0][1:], 0)
		u2 := helper.StrToInt(arr[1], 0)
		if u1 > 0 && u2 > 0 {
			return UsersToRoomKey(u1, u2), nil
		}
	}
	if config.IS_DEBUG {
		panic("WRONG Keys_ChatKeyToRoomKey param for ChatKey: " + ChatKey)
	}
	return "", errors.New("WRONG ChatKey ")
}

/*func GetOrCreateDirectChatForPeers(me int, peer int) (*x.Chat, error) {
    keyChach := fmt.Sprintf("chat_%d_%d", me, peer)
    ch, ok := RowCache.GetChatByChatId(keyChach)
    if ok {
        return ch, nil
    }

    chatMe, err := x.NewChat_Selector().UserId_Eq(me).PeerUserId_Eq(peer).GetRow(base.DB)
    if err != nil {
        chatMe = &x.Chat{
            ChatKey:              UsersToChatKey(me, peer),
            RoomKey:              UsersToRoomKey(me, peer),
            RoomTypeEnumId:       int(x.RoomTypeEnum_DIRECT),
            UserId:               me,
            PeerUserId:           peer,
            GroupId:              0,
            CreatedSe:            helper.TimeNow(),
            StartMessageIdFrom:   helper.NextRowsSeqIdBase(),
            LastDeletedMessageId: 0,
            LastSeenMessageId:    0,
            LastMessageId:        0,
            UpdatedMs:            helper.TimeNowMs(),
        }
        err = chatMe.Save(base.DB)
        if err != nil {
            return nil, err
        }
    }
    RowCache.SetChatByKey(keyChach, chatMe)
    return chatMe, nil
}*/
