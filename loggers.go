package cmn

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	LogInfo    *log.Logger
	LogWarning *log.Logger
	LogError   *log.Logger
)

const logsDir = "logs"

func InitLoggers(logFileName string) {
	if len(logFileName) > 64 {
		panic("Too long log file name")
	}

	if len(logFileName) == 0 {
		LogInfo = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
		LogWarning = log.New(os.Stdout, "[WARNING]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
		LogError = log.New(os.Stdout, "[ERROR]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
		return
	}

	if _, err := os.Stat(logsDir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(logsDir, 0666)
			if err != nil {
				panic(err.Error())
			}
		} else {
			panic(err.Error())
		}
	}

	fStream, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logsDir, logFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	logWriter := io.MultiWriter(os.Stdout, fStream)
	LogInfo = log.New(logWriter, "[INFO]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
	LogWarning = log.New(logWriter, "[WARNING]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
	LogError = log.New(logWriter, "[ERROR]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
}
