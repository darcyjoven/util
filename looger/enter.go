package looger

import (
	"log"
	"os"
	"path/filepath"

	"github.com/darcyjoven/util"
)

func Execute(path, name, ext string) error {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	if path != "" {
		path = filepath.Join(path, name+util.GetShortDate()+"."+ext)
		logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Println(err)
			return err
		}
		log.SetOutput(logFile)
	}
	log.Println("log init successed !")
	return nil
}
