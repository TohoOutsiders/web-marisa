package cache

import (
	"github.com/go-redis/redis"
	"log"
	"server/setting"
	"time"
)

type Redis struct {
	Conn *redis.Client
}

func (r *Redis) Client() *redis.Client {
	return r.Conn
}

func (r *Redis) Connect() error {
	conf := setting.Config.Redis

	options := &redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Pass,
		DB:           conf.DB,
		ReadTimeout:  conf.Timeout * time.Second,
		WriteTimeout: conf.Timeout * time.Second,
	}
	client := redis.NewClient(options)

	if _, err := client.Ping().Result(); err != nil {
		return err
	}

	r.Conn = client

	log.Println("Connect Redis Success")

	return nil
}
