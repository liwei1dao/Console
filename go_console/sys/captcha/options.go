package captcha

import (
	"github.com/liwei1dao/lego/utils/mapstructure"
)

type Option func(*Options)
type Options struct {
	Serverhost            string
	Fromemail             string
	Fompasswd             string
	Serverport            int
	RedisUrl              string
	CaptchaExpirationdate int
}

func SetServerhost(v string) Option {
	return func(o *Options) {
		o.Serverhost = v
	}
}

func SetFromemail(v string) Option {
	return func(o *Options) {
		o.Fromemail = v
	}
}

func SetFompasswd(v string) Option {
	return func(o *Options) {
		o.Fompasswd = v
	}
}

func SetRedisUrl(v string) Option {
	return func(o *Options) {
		o.RedisUrl = v
	}
}

func SetCaptchaExpirationdate(v int) Option {
	return func(o *Options) {
		o.CaptchaExpirationdate = v
	}
}

func newOptions(config map[string]interface{}, opts ...Option) Options {
	options := Options{
		RedisUrl:              "redis://127.0.0.1:6379/1",
		CaptchaExpirationdate: 60,
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
		RedisUrl:              "redis://127.0.0.1:6379/1",
		CaptchaExpirationdate: 60,
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}
