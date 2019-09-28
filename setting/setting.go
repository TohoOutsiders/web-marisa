/**
 * @Author: Tomonori
 * @Date: 2019/6/18 14:49
 * @File: setting
 * @Desc:
 */
package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

var (
	Config *Conf
)

type Conf struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
}

type Server struct {
	Port         string `yaml:"port"`
	ReadTimeout  string `yaml:"read-timeout"`
	WriteTimeout string `yaml:"write-timeout"`
}

type Database struct {
	Type        string `yaml:"type"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"table-prefix"`
}

type Redis struct {
	Addr        string        `yaml:"addr"`
	Pass        string        `yaml:"pass"`
	DB          int           `yaml:"db"`
	Timeout     time.Duration `yaml:"timeout"`
	ExpiredTime int           `yaml:"expired-time"`
}

func init() {
	Config = getConf()
	log.Println("[Setting] Config init success")
}

func getConf() *Conf {
	var c *Conf
	file, err := ioutil.ReadFile("Config/config.yml")
	if err != nil {
		log.Println("[Setting] config error: ", err)
	}
	err = yaml.UnmarshalStrict(file, &c)
	if err != nil {
		log.Println("[Setting] yaml unmarshal error: ", err)
	}
	return c
}
