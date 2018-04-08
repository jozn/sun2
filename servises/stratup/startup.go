package stratup

import (
	_ "ms/sun/servises/archiver_service"
	_ "ms/sun/servises/model_service/action_fanout_service"
	_ "ms/sun/servises/model_service/action_saver_service"
	_ "ms/sun/servises/model_service/home_fanout_service"
	_ "ms/sun/servises/model_service/notify_saver_service"
	_ "ms/sun/servises/model_service/trigger_service"
	_ "ms/sun/servises/post_counter_service"
)
