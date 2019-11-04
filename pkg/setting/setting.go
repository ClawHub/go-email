package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

//服务相关
type Server struct {
	ProjectName  string
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("app.ini")
	if err != nil {
		log.Fatal("setting.Setup, fail to parse 'app.ini' ", err)
	}

	MapTo("server", ServerSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}

func MapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatal("Cfg.MapTo Setting err' ", err)
	}
}
