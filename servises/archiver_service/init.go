package archiver_service

func init() {
	go cleanEvents()
	go gc_action_fanout()
	go gc_home_fanout()
	go gc_trigger()
}
