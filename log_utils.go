package raselper

import (
	"log"
	"os"
)

var file, logFileErr = os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

func Log(info string) error {
	if logFileErr != nil {
		return logFileErr
	}
	log.SetOutput(file)
	log.Println(info)
	return nil
}
