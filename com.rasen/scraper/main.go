package main

import (
	"com.rasen/scraper/common/collector"
	"com.rasen/scraper/config"
	"com.rasen/scraper/database"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.LoadCfg()
	mgodb := database.NewMongoDB(cfg.MongoDBAddr)
	//collector.SearchWeb(cfg,mgodb)
	collector.SearchWebWithMap(cfg,mgodb)

	go func() {
		killSig := make(chan os.Signal)
		for{
			signal.Notify(killSig,os.Kill,os.Interrupt,syscall.SIGTERM,syscall.SIGQUIT)
			switch <-killSig {
			case os.Kill,os.Interrupt,syscall.SIGTERM,syscall.SIGQUIT:
				database.MemFd.Close()
			}
		}
	}()
}
