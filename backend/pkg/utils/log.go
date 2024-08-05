package utils

import (
	"log"
	"os"
	"time"
)

var (
	LoggerInfo  *log.Logger
	LoggerError *log.Logger
	logFile     *os.File
	Info        = "\033[34m" // Blue
	Error       = "\033[31m" // Red
	Warn        = "\033[33m" // Yellow
	Reset       = "\033[0m"
)

const maxFileSize = 24 * 1024

func InitLogger() {
	var err error
	logFile, err = os.OpenFile("log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error opening file: " + err.Error())
		return
	}

	LoggerInfo = log.New(logFile, "INFO: ", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	LoggerError = log.New(logFile, "ERROR: ", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
}

func RotateLogFile() {
	if stat, err := logFile.Stat(); err == nil {
		if stat.Size() >= maxFileSize {
			err := logFile.Close()
			if err != nil {
				log.Println("Error closing file: " + err.Error())
				return
			}
			newName := "log/app-" + time.Now().Format("2006-01-02_15-04-05") + ".log"
			err = os.Rename("log/app.log", newName)
			if err != nil {
				log.Println("Error renaming file: " + err.Error())
				LoggerError.Println(Error + "Error renaming file: " + err.Error() + Reset)
				return
			}
			LoggerInfo.Println(Info + "Log file rotated and file renamed to: " + newName + Reset)
			InitLogger()
		}
	}
}

func CleanUp() {
	if logFile != nil {
		Close()
		err := logFile.Close()
		if err != nil {
			LoggerError.Println("Error closing file: " + err.Error() + Reset)
			return
		}
	}
}

func Welcome() {
	LoggerInfo.Println(Warn + "       __                 __   __                               " + Reset)
	LoggerInfo.Println(Warn + "      / /___ _____  ___  / /  / /   ____  ____ _____ ____  _____" + Reset)
	LoggerInfo.Println(Warn + " __  / / __ `/ __ \\/ _ \\/ /  / /   / __ \\/ __ `/ __ `/ _ \\/ ___/" + Reset)
	LoggerInfo.Println(Warn + "/ /_/ / /_/ / / / /  __/ /  / /___/ /_/ / /_/ / /_/ /  __/ /    " + Reset)
	LoggerInfo.Println(Warn + "\\____/\\__,_/_/ /_/\\___/_/  /_____/\\____/\\__, /\\__, /\\___/_/     " + Reset)
	LoggerInfo.Println(Warn + "                                       /____//____/             " + Reset)
}

func Close() {
	LoggerInfo.Println(Warn + "   ________                    __" + Reset)
	LoggerInfo.Println(Warn + "  / ____/ /___  ________  ____/ /" + Reset)
	LoggerInfo.Println(Warn + " / /   / / __ \\/ ___/ _ \\/ __  / " + Reset)
	LoggerInfo.Println(Warn + "/ /___/ / /_/ (__  )  __/ /_/ /  " + Reset)
	LoggerInfo.Println(Warn + "\\____/_/\\____/____/\\___/\\__,_/   " + Reset)
	LoggerInfo.Println("Database closed")
	LoggerInfo.Println("Server closed")
	LoggerInfo.Println("Logger closed")
	LoggerInfo.Println("Goodbye!")
}
