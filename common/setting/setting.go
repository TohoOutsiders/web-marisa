/**
 * @Author: Tomonori
 * @Date: 2019/6/18 14:49
 * @File: setting
 * @Desc:
 */
package setting

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

var (
	Config *Conf
)

type (
	Conf struct {
		Server   Server   `yaml:"server"`
		Database Database `yaml:"database"`
		Redis    Redis    `yaml:"redis"`
		RabbitMQ RabbitMQ `yaml:"rabbitmq"`
	}

	Server struct {
		Port         string `yaml:"port"`
		ReadTimeout  string `yaml:"read-timeout"`
		WriteTimeout string `yaml:"write-timeout"`
	}

	Database struct {
		Type        string `yaml:"type"`
		User        string `yaml:"user"`
		Password    string `yaml:"password"`
		Host        string `yaml:"host"`
		Name        string `yaml:"name"`
		TablePrefix string `yaml:"table-prefix"`
	}

	Redis struct {
		Addr        string        `yaml:"addr"`
		Pass        string        `yaml:"pass"`
		DB          int           `yaml:"db"`
		Timeout     time.Duration `yaml:"timeout"`
		ExpiredTime int           `yaml:"expired-time"`
	}

	RabbitMQ struct {
		Addr string `yaml:"addr"`
		Port int    `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	}
)

func init() {
	config, err := getConf()
	if err != nil {
		log.Fatal(err)
	}
	Config = config
	log.Println("[Setting] Config init success")
}

func getConf() (*Conf, error) {
	err := errors.New("[Config] get config file error")

	var c *Conf
	file, err := ioutil.ReadFile("Config/config.yml")
	if err != nil {
		log.Println("[Setting] config error: ", err)
		return nil, err
	}
	err = yaml.UnmarshalStrict(file, &c)
	if err != nil {
		log.Println("[Setting] yaml unmarshal error: ", err)
		return nil, err
	}
	return c, nil
}
