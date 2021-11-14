package redis_driver

import "github.com/go-redis/redis"

type ConfigRedis struct {
	REDIS_ENDPOINT string 
	REDIS_PASSWORD string
}

func (config *ConfigRedis) InitialRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.REDIS_ENDPOINT,
		Password: config.REDIS_PASSWORD,
		DB: 0,
	})

	return rdb
}
