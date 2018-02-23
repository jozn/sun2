package event_service

import (
	"ms/sun/base"
	"ms/sun2/shared/x"
	"time"
)

var saverBuffer = make(chan x.Event, 100000)

func SaveEvent(eventType EventType, event x.Event) {
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
