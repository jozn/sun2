package x

import (
	"ms/sun_old/base"
	"strings"
	"time"
)

func init() {
	go triggerLoader()
}

type TriggerStringModel interface {
	OnInsert(ins []string)
	OnUpdate(ins []string)
	OnDelete(ins []string)
}

type TriggerIntModel interface {
	OnInsert(ins []int)
	OnUpdate(ins []int)
	OnDelete(ins []int)
}

type TriggerModelListener struct {
	Action  TriggerIntModel
	Comment TriggerIntModel
	Post    TriggerIntModel
	User    TriggerIntModel
	Chat    TriggerStringModel
}

var lastLoaded int
var ArrTriggerListeners = make([]TriggerModelListener, 10)
var ActivateTrigger = false

func triggerLoader() {
	time.Sleep(time.Second * 10)
	for {
		time.Sleep(time.Second)
		if !ActivateTrigger {
			continue
		}

		selector := NewTriggerLog_Selector()
		if lastLoaded > 0 {
			selector.Id_GT(lastLoaded)
		}
		triggers, err := selector.OrderBy_Id_Asc().GetRows(base.DB)
		if err != nil || len(triggers) == 0 {
			continue
		}
		lastLoaded = triggers[len(triggers)-1].Id

		for _, listener := range ArrTriggerListeners {
			collect := triggerModelWalk{}
			for _, trig := range triggers {

				switch strings.ToUpper(trig.ChangeType) {
				case "INSERT":
					switch strings.ToUpper(trig.ModelName) {
					case "ACTION":
						if listener.Action != nil {
							collect.Action.OnInsert = append(collect.Action.OnInsert, trig.TargetId)
						}
					case "COMMENT":
						if listener.Comment != nil {
							collect.Comment.OnInsert = append(collect.Comment.OnInsert, trig.TargetId)
						}
					case "POST":
						if listener.Post != nil {
							collect.Post.OnInsert = append(collect.Post.OnInsert, trig.TargetId)
						}
					case "USER":
						if listener.User != nil {
							collect.User.OnInsert = append(collect.User.OnInsert, trig.TargetId)
						}
					case "CHAT":
						if listener.Chat != nil {
							collect.Chat.OnInsert = append(collect.Chat.OnInsert, trig.TargetStr)
						}
					}
				case "UPDATE":
					switch strings.ToUpper(trig.ModelName) {
					case "ACTION":
						if listener.Action != nil {
							collect.Action.OnUpdate = append(collect.Action.OnUpdate, trig.TargetId)
						}
					case "COMMENT":
						if listener.Comment != nil {
							collect.Comment.OnUpdate = append(collect.Comment.OnUpdate, trig.TargetId)
						}
					case "POST":
						if listener.Post != nil {
							collect.Post.OnUpdate = append(collect.Post.OnUpdate, trig.TargetId)
						}
					case "USER":
						if listener.User != nil {
							collect.User.OnUpdate = append(collect.User.OnUpdate, trig.TargetId)
						}
					case "CHAT":
						if listener.Chat != nil {
							collect.Chat.OnUpdate = append(collect.Chat.OnUpdate, trig.TargetStr)
						}
					}
				case "DELETE":
					switch strings.ToUpper(trig.ModelName) {
					case "ACTION":
						if listener.Action != nil {
							collect.Action.OnDelete = append(collect.Action.OnDelete, trig.TargetId)
						}
					case "COMMENT":
						if listener.Comment != nil {
							collect.Comment.OnDelete = append(collect.Comment.OnDelete, trig.TargetId)
						}
					case "POST":
						if listener.Post != nil {
							collect.Post.OnDelete = append(collect.Post.OnDelete, trig.TargetId)
						}
					case "USER":
						if listener.User != nil {
							collect.User.OnDelete = append(collect.User.OnDelete, trig.TargetId)
						}
					case "CHAT":
						if listener.Chat != nil {
							collect.Chat.OnDelete = append(collect.Chat.OnDelete, trig.TargetStr)
						}
					}
				}
			}

			//each

			if listener.Action != nil {
				if len(collect.Action.OnInsert) != 0 {
					listener.Action.OnInsert(collect.Action.OnInsert)
				}
				if len(collect.Action.OnUpdate) != 0 {
					listener.Action.OnUpdate(collect.Action.OnUpdate)
				}
				if len(collect.Action.OnDelete) != 0 {
					listener.Action.OnDelete(collect.Action.OnDelete)
				}
			}

			if listener.Comment != nil {
				if len(collect.Comment.OnInsert) != 0 {
					listener.Comment.OnInsert(collect.Comment.OnInsert)
				}
				if len(collect.Comment.OnUpdate) != 0 {
					listener.Comment.OnUpdate(collect.Comment.OnUpdate)
				}
				if len(collect.Comment.OnDelete) != 0 {
					listener.Comment.OnDelete(collect.Comment.OnDelete)
				}
			}

			if listener.Post != nil {
				if len(collect.Post.OnInsert) != 0 {
					listener.Post.OnInsert(collect.Post.OnInsert)
				}
				if len(collect.Post.OnUpdate) != 0 {
					listener.Post.OnUpdate(collect.Post.OnUpdate)
				}
				if len(collect.Post.OnDelete) != 0 {
					listener.Post.OnDelete(collect.Post.OnDelete)
				}
			}

			if listener.User != nil {
				if len(collect.User.OnInsert) != 0 {
					listener.User.OnInsert(collect.User.OnInsert)
				}
				if len(collect.User.OnUpdate) != 0 {
					listener.User.OnUpdate(collect.User.OnUpdate)
				}
				if len(collect.User.OnDelete) != 0 {
					listener.User.OnDelete(collect.User.OnDelete)
				}
			}

			if listener.Chat != nil {
				if len(collect.Chat.OnInsert) != 0 {
					listener.Chat.OnInsert(collect.Chat.OnInsert)
				}
				if len(collect.Chat.OnUpdate) != 0 {
					listener.Chat.OnUpdate(collect.Chat.OnUpdate)
				}
				if len(collect.Chat.OnDelete) != 0 {
					listener.Chat.OnDelete(collect.Chat.OnDelete)
				}
			}

		}
	}
}

type triggerStringCollection struct {
	OnInsert []string
	OnUpdate []string
	OnDelete []string
}

type triggerIntCollection struct {
	OnInsert []int
	OnUpdate []int
	OnDelete []int
}

type triggerModelWalk struct {
	Action  triggerIntCollection
	Comment triggerIntCollection
	Post    triggerIntCollection
	User    triggerIntCollection
	Chat    triggerStringCollection
}
