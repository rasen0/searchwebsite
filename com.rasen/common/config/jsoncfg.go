package config

import (
	log "com.rasen/common/structlog"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

func LoadCfg(path string,arg interface{}) error{
	fd,err := os.Open(path)
	if err != nil{
		log.Logger.WithFields(logrus.Fields{"err":err}).Error("open cfg file fail")
		return err
	}
	if stat,_ := fd.Stat();stat.IsDir(){
		log.Logger.WithFields(logrus.Fields{"err":stat}).Error("cfg file path is wrong")
	}
	decoder := json.NewDecoder(fd)
	err = decoder.Decode(arg)
	if err != nil{
		log.Logger.WithFields(logrus.Fields{"err":err}).Error("decode cfg file fail")
		return err
	}
	return nil
}
