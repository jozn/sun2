package trigger_service

import (
	"ms/sun/base"
	"ms/sun2/shared/x"
	"strings"
	"time"
)

func init() {
	go triggerLoader()
}

type TriggerStringModel interface {
	OnInsert([]string)
	OnUpdate([]string)
	OnDelete([]string)
}

type TriggerIntModel interface {
	OnInsert([]int)
	OnUpdate([]int)
	OnDelete([]int)
}

type TriggerModelListener struct {
	Chat TriggerStringModel
	Post TriggerStringModel
}

var lastLoaded int
var ArrTriggerListeners = make([]TriggerModelListener, 10)
var ActivateTrigger = false

func triggerLoader() {
	for {
		time.Sleep(time.Second)
		if !ActivateTrigger {
			continue
		}

		selector := x.NewTriggerLog_Selector()
		if lastLoaded > 0 {
			selector.Id_GE(lastLoaded)
		}
		triggers, err := selector.OrderBy_Id_Asc().GetRows(base.DB)
		if err != nil || len(triggers) == 0 {
			continue
		}
		lastLoaded = triggers[len(triggers)-1].Id
		collect := triggerStringCollection{}
		for _, litener := range ArrTriggerListeners {
			for _, trig := range triggers {
				//based on model
				modelId := trig.TargetStr
				switch strings.ToUpper(trig.ChangeType) {
				case "INSERT":
					switch strings.ToUpper(trig.ModelName) {
					case "CHAT": //model.ModleNameUpper
						if litener.Chat != nil {
							collect.OnInsert = append(collect.OnInsert, modelId)
						}
					}
				case "UPDATE":
					switch strings.ToUpper(trig.ModelName) {
					case "CHAT":
						if litener.Chat != nil {
							collect.OnUpdate = append(collect.OnUpdate, modelId)
						}
					}
				case "DELETE":
					switch strings.ToUpper(trig.ModelName) {
					case "CHAT":
						if litener.Chat != nil {
							collect.OnDelete = append(collect.OnDelete, modelId)
						}
					}
				}
			}

			//each
			if litener.Chat != nil {
				if len(collect.OnInsert) != 0 {
					litener.Chat.OnInsert(collect.OnInsert)
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

type TriggerModelWalk struct {
	Chat TriggerStringModel
	Post TriggerStringModel
}
