package tool

import (
	"com.rasen/common/structlog"
	"github.com/sirupsen/logrus"
	"time"
)

func LoopRun(succInterval time.Duration,failInterval time.Duration,fn func()bool){
	defer func() {
		if err := recover();err != nil{
			structlog.Logger.WithFields(logrus.Fields{"err":err}).Error("loopRun function fail")
			//fmt.Fprintln(os.Stderr,"loopRun function fail")
			LoopRun(succInterval,failInterval,fn)
		}
	}()
	time.Sleep(failInterval)
	for{
		if fn(){
			time.Sleep(succInterval)
		}else{
			time.Sleep(failInterval)
		}
	}
}
