package pusher_service

import (
	"ms/sun/shared/helper"
	"time"
)

var startTime = time.Now()

type MonitorOut struct {
	Monitor MonitorAllTime
}

type MonitorAllTime struct {
	TotalMysqlRowsLoaded               int
	TotalChatRowsLoaded                int
	TotalSocialRowsLoaded              int
	TotalClosedPipes                   int
	TotalWebSocketsClosed              int
	TotalWebSocketsReceivedText        int
	TotalWebSocketsReceivedBinary      int
	TotalWebSocketsReceivedPanics      int
	TotalWebSocketsSendPanics          int
	TotalRowsProceedWiltActiveUser     int
	TotalChatRowsProceedWiltActiveUser int
	TotalUsersAdded                    int
	CurrentActiveUsers                 int
	CurrentActivePipes                 int
	ActiveTimeSeconds                  int
}

func (m *MonitorAllTime) Finalize() {
	allPipesMap.m.RLock()
	m.CurrentActiveUsers = len(allPipesMap.mp)
	allPipesMap.m.RUnlock()

	m.ActiveTimeSeconds = int(time.Now().Unix() - startTime.Unix())
}

func (m *MonitorAllTime) PrintLoop1Min() {
	go func() {
		for {
			m.Finalize()

			/*out := struct {
				Monitor MonitorAllTime
			}{
				Monitor: *m,
			}*/
			//out := MonitorOut{
			//	Monitor: *m,
			//}
			helper.PertyPrintNew(m)

			time.Sleep(time.Second * 5)
		}
	}()
}

var Monitor = MonitorAllTime{}
