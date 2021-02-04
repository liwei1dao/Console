package console

import (
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib/modules/console"
)

func NewModule() core.IModule {
	m := new(Console)
	return m
}

type Console struct {
	console.Console
}
