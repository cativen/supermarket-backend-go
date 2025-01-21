package common

import (
	"log"
	"os"
)

func InitLog() {
	logFile, err := os.OpenFile("./log/nohup.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
