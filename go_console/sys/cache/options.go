package cache

import "github.com/liwei1dao/lego/utils/mapstructure"

type Option func(*Options)
type Options struct {
	RedisUrl                 string
	UserCacheExpirationDate  int
	TokenCacheExpirationDate int
	MonitorTotalTime         int //监控总时长 小时为单位
}

func SetRedisUrl(v string) Option {
	return func(o *Options) {
		o.RedisUrl = v
	}
}

func SetUserCacheExpirationDate(v int) Option {
	return func(o *Options) {
		o.UserCacheExpirationDate = v
	}
}

func SetTokenCacheExpirationDate(v int) Option {
	return func(o *Options) {
		o.TokenCacheExpirationDate = v
	}
}

func newOptions(config map[string]interface{}, opts ...Option) Options {
	options := Options{
		RedisUrl:                 "redis://127.0.0.1:6379/1",
		UserCacheExpirationDate:  60,
		TokenCacheExpirationDate: 3600,
	}
	if config != nil {
		mapstructure.Decode(config, &options)
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func newOptionsByOption(opts ...Option) Options {
	options := Options{
		RedisUrl:                 "redis://127.0.0.1:6379/1",
		UserCacheExpirationDate:  60,
		TokenCacheExpirationDate: 3600,
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}
