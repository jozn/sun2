package monitor_service

import (
	"github.com/hokaccha/go-prettyjson"
	"math/rand"
	"ms/sun/shared/base"
	"ms/sun/shared/x"
	"time"
    "ms/sun/shared/helper"
)

var allMetrics = []interface{}{SqlMetrics, SocialMetrics, ChatMetrics, InternalMetrics, HttpRpcMetrics}

var SqlMetrics = &SqlMetricsType{}
var SocialMetrics = &SocialMetricsType{}
var ChatMetrics = &ChatMetricsType{}
var InternalMetrics = &InternalMetricsType{}
var HttpRpcMetrics = &HttpRpcMetricsType{
    MethodsCalls: make(map[string]int),
}

func saveMetericLoop() {
	now := time.Now()
	ml := &x.MetricLog{
		InstanceId:   rand.Intn(1000),
		StartFrom:    now.Format(time.Stamp),
		EndTo:        "",
		StartTime:    int(now.Unix()),
		Duration:     "",
		MetericsJson: "",
	}

	go func() {
		for {
			time.Sleep(time.Second * 1)
			ml.EndTo = time.Now().Format(time.Stamp)
			ml.Duration = time.Now().Sub(now).String()

			bts, _ := prettyjson.Marshal(allMetrics)
			ml.MetericsJson = string(bts)
			err := ml.Save(base.DB)
			helper.NoErrDebug(err)
		}
	}()
}
