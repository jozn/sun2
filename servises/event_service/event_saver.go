package event_service

import (
	"ms/sun_old/base"
	"ms/sun_old/helper"
	"ms/sun/shared/x"
	"time"
)

var saverBuffer = make(chan x.Event, 100000)

func SaveEvent(eventType EventType, event x.Event) {
	if event.EventId == 0 {
		event.EventId = helper.NextRowsSeqId()
	}
	event.EventType = int(eventType)
	saverBuffer <- event
}

func delaySaveBuffer() {
	ticker := time.Tick(time.Millisecond * 500)
	var events []x.Event
	for {
		select {
		case ev := <-saverBuffer:
			events = append(events, ev)
		case <-ticker:
			if len(events) > 0 {
				evs := events
				events = []x.Event{}
				x.MassReplace_Event(evs, base.DB)
			}
		}
	}
}
