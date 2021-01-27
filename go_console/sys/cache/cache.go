package cache

import "github.com/liwei1dao/lego/sys/redis"

func newsys(options Options) (sys *Cache, err error) {
	sys = &Cache{options: options}
	sys.redis, err = redis.NewSys(redis.SetRedisUrl(options.RedisUrl))
	return
}

type Cache struct {
	options Options
	redis   redis.IRedisFactory
}
