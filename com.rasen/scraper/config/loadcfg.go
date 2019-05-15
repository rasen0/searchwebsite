package config

import (
	"os"
	"path/filepath"
	"strings"

	"com.rasen/common/config"
)

const CONFIG_FILE = "config.conf"

type DefaultConfig struct{
	MongoDBAddr    string `json:"mongodb_addr"`
	UserAgent  string `json:"user_agent"`
	GuideWeb   []string `json:"guide_web"`
}

func LoadCfg() *DefaultConfig{
	absPath,_ :=filepath.Abs(filepath.Dir(os.Args[0]))
	absPath = strings.Replace(absPath,"\\","/",-1)
	path := filepath.Join(absPath,CONFIG_FILE)
	cfg := new(DefaultConfig)
	err := config.LoadCfg(path,cfg)
	if err!= nil && len(cfg.GuideWeb) < 1{
		os.Exit(0)
	}
	return cfg
}
