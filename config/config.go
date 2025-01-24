package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Zhihu       string
	Avatar      string
	UserName    string
	UserDesc    string
}

type SystemConfig struct {
	AppName         string
	Version         string
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	Cfg = new(tomlConfig)
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	dir, _ := os.Getwd()
	Cfg.System.CurrentDir = dir
	if err != nil {
		log.Fatal(err)
	}
}
