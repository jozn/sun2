package file_service

import (
	"github.com/syncthing/syncthing/lib/rand"
	"ms/sun/servises/file_service/file_common"
	"ms/sun/shared/base"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"net/http"
	"time"
)

var httpMonitorChan = make(chan *httpWirterWrapper, 10000)
var inoLogDbChan = make(chan x.XfileServiceInfoLog, 1000)

type httpWirterWrapper struct {
	w    http.ResponseWriter
	r    *http.Request
	req  file_common.RowReq
	code int
}

func (w *httpWirterWrapper) Header() http.Header {
	return w.w.Header()
}

func (w *httpWirterWrapper) Write(b []byte) (int, error) {
	return w.w.Write(b)
}

func (w *httpWirterWrapper) WriteHeader(code int) {
	w.code = code
	w.w.WriteHeader(code)
}

var seqLog = 0
var instanceLogId = rand.Intn(10000) + 1

func init() {
	go saveMonitor_go()
	go saveFileMetics_go()
}

func saveMonitor_go() {
	defer func() {
		err := recover()
		if err != nil {
			go saveMonitor_go()
		}
	}()
	var arrInfo []x.XfileServiceInfoLog
	var arrReqs []x.XfileServiceRequestLog
	tic := time.NewTicker(time.Second)
	for {
		select {
		case wrap := <-httpMonitorChan:
			seqLog++
			//if wrap.req != nil {
			m := x.XfileServiceRequestLog{
				Id:          helper.NextRowsSeqId(),
				LocalSeq:    seqLog,
				InstanceId:  instanceLogId,
				Url:         wrap.r.RequestURI,
				HttpCode:    wrap.code,
				CreatedTime: time.Now().Format("15:04:05"),
			}
			metricsLog.Exts[wrap.req.FileExtensionWithDot] += 1
			metricsLog.Sizes[wrap.req.RequestedImageSize] += 1
			arrReqs = append(arrReqs, m)
			//}
		case info := <-inoLogDbChan:
			arrInfo = append(arrInfo, info)
		case _ = <-tic.C:
			x.MassReplace_XfileServiceInfoLog(arrInfo, base.DB)
			arrInfo = arrInfo[:0]
			x.MassReplace_XfileServiceRequestLog(arrReqs, base.DB)
			arrReqs = arrReqs[:0]
		}

	}
}

type metricsLogType struct {
	Exts  map[string]int
	Sizes map[int]int
}

var metricsLog = &metricsLogType{
	make(map[string]int, 50), make(map[int]int, 50),
}

func saveFileMetics_go() {
	defer helper.JustRecover()
	m := x.XfileServiceMetricLog{
		Id:         helper.NextRowsSeqId(),
		InstanceId: instanceLogId,
		MetricJson: "",
	}
	for {
		time.Sleep(time.Second)
		m.MetricJson = helper.ToJson(metricsLog)
		m.Save(base.DB)
	}
}

func cleanOldLogs_go() {
	for {
		time.Sleep(time.Minute)
		x.NewXfileServiceInfoLog_Deleter().Id_LT(helper.NanoTimeBeforeNowSeconds(86400)).Delete(base.DB)
		x.NewXfileServiceMetricLog_Deleter().Id_LT(helper.NanoTimeBeforeNowSeconds(86400)).Delete(base.DB)
		x.NewXfileServiceRequestLog_Deleter().Id_LT(helper.NanoTimeBeforeNowSeconds(86400)).Delete(base.DB)
	}
}
