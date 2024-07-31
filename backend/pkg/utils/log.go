package utils

import (
	"log"
	"os"
	"time"
)

var (
	Logger  *log.Logger
	logFile *os.File
)

const maxFileSize = 900 * 1024

func InitLogger() {
	var err error
	logFile, err = os.OpenFile("log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error opening file: " + err.Error())
		return
	}

	Logger = log.New(logFile, "INFO: ", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile)
}

func RotateLogFile() {
	if stat, err := logFile.Stat(); err == nil {
		if stat.Size() > maxFileSize {
			err := logFile.Close()
			if err != nil {
				log.Println("Error closing file: " + err.Error())
				return
			}
			newName := "log/app-" + time.Now().Format("2006-01-02_15-04-05") + ".log"
			err = os.Rename("log/app.log", newName)
			if err != nil {
				log.Println("Error renaming file: " + err.Error())
				Logger.Println("Error renaming file: " + err.Error())
				return
			}
			InitLogger()
		}
	}
}
