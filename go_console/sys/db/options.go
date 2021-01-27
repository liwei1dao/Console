package db

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

type Option func(*Options)
type Options struct {
	MongodbUrl      string
	MongodbDatabase string
	TimeOut         time.Duration
}

func SetTMongodbUrl(v string) Option {
	return func(o *Options) {
		o.MongodbUrl = v
	}
}

func SetTMongodbDatabase(v string) Option {
	return func(o *Options) {
		o.MongodbDatabase = v
	}
}

func SetTimeOut(v time.Duration) Option {
	return func(o *Options) {
		o.TimeOut = v
	}
}

func newOptions(config map[string]interface{}, opts ...Option) Options {
	options := Options{
		TimeOut: time.Second * 3,
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
		TimeOut: time.Second * 3,
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}
