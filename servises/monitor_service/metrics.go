package monitor_service

type AllMetricsType2 struct {
}

type SqlMetricsType struct {
}

type SocialMetricsType struct {
}

type ChatMetricsType struct {
}

type InternalMetricsType struct {
}

type HttpRpcMetricsType struct {
	TotalCall    int
	MethodsCalls map[string]int
}
