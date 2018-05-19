package helper

import (
	"log"
	"ms/sun/shared/config"
)

func NoErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func NoErrDebug(err error) {
	if config.IS_DEBUG {
		if err != nil {
			log.Panic(err)
		}
	}
}

func XCLogErr(err error) {
	if config.IS_DEBUG {
		if err != nil {
			log.Println(err)
		}
	}
}

func XCLog(str ...interface{}) {
	if config.IS_DEBUG {
		if len(str) > 0 {
			log.Println("CQL: ", str)
		}
	}
}
