package rpc_service

func pan(string string) {

}

func toLimit(paramLimit int32, defualtLimit int) int {
	limit := defualtLimit
	if paramLimit != 0 {
		limit = int(paramLimit)
	}
	return limit
}
