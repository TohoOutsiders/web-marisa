package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
)

func init() {
	var err error
	Cfg, err = ini.Load("Config/config.ini")
	if err != nil {
		log.Fatalln("Fial to parse 'Config/config.ini: ", err)
	}
	LoadServer()
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalln("Fail to get config section 'server': ", err)
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(3000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
