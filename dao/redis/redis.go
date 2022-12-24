package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

var (
	Client *redis.Client
	Nil    = redis.Nil
)
var ctx = context.Background()

func RedisInit(conf *ini.File) (err error) {
	port, _ := conf.Section("redis").Key("port").Int()
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Section("redis").Key("host").String(), port),
		Password: fmt.Sprintf("%s", conf.Section("redis").Key("password").String()),
		DB:       0,
	})
	_, err = Client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
func Close() {
	_ = Client.Close()
}
