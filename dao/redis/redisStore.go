package redis

import (
	"encoding/json"
	"time"
)

type cacheDb struct{}

func (c cacheDb) Set(key string, value interface{}, expiration int) {

	v, err := json.Marshal(value)
	if err == nil {
		Client.Set(Ctx, key, string(v), time.Second*time.Duration(expiration))
	}
}
func (c cacheDb) Get(key string, obj interface{}) bool {
	valueStr, err1 := Client.Get(Ctx, key).Result()
	if err1 == nil && valueStr != "" {
		err2 := json.Unmarshal([]byte(valueStr), obj)
		return err2 == nil
	}
	return false
}

var CacheDb = &cacheDb{}
