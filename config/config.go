package config

import (
	"log"
	utils "myapp/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	DbUser    string
	DbPass	  string
	DbPort    string
	LogFile   string
}

var Config ConfigList

func init(){
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

// iniファイルのデータを読み込みConfigListに設定する
func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil{
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:cfg.Section("web").Key("port").MustString("8080"),
		// SQLDriver:cfg.Section("db").Key("driver").String(),
		SQLDriver:cfg.Section("newdb").Key("driver").String(),
		// DbName:cfg.Section("db").Key("name").String(),
		DbName:cfg.Section("newdb").Key("name").String(),
		DbUser:cfg.Section("newdb").Key("user").String(),
		DbPass:cfg.Section("newdb").Key("password").String(),
		DbPort:cfg.Section("newdb").Key("port").String(),
		LogFile:cfg.Section("web").Key("logfile").String(),
	}
}

