package main

import (
	"com.rasen/scraper/common/collector"
	"com.rasen/scraper/config"
	"com.rasen/scraper/database"
)

func main() {
	cfg := config.LoadCfg()
	mgodb := database.NewMongoDB(cfg.MongoDBAddr)
	collector.SearchWeb(cfg,mgodb)
}
