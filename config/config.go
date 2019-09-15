package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	GoUrl     string
	VueUrl    string
	DbName    string
	SQLDriver string
	LogFile   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").String(),
		GoUrl:     cfg.Section("web").Key("go_url").String(),
		VueUrl:    cfg.Section("web").Key("vue_url").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		LogFile:   cfg.Section("log").Key("file_name").String(),
	}
}
