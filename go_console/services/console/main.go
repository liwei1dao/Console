package main

import (
	"flag"
	"lego_console/modules/console"
	"lego_console/services"

	"github.com/liwei1dao/lego"
	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
)

var (
	sID = flag.String("id", "console", "获取需要启动的服务id,id不同,读取的配置文件也不同") //启动服务的Id
)

func main() {
	flag.Parse()
	s := NewService(
		cluster.SetId(*sID),
	)
	lego.Run(s, //运行模块
		console.NewModule(),
	)
}

func NewService(ops ...cluster.Option) core.IService {
	s := new(ConsoleService)
	s.Configure(ops...)
	return s
}

type ConsoleService struct {
	services.BaseService
}

func (this *ConsoleService) InitSys() {
	this.BaseService.InitSys()
}
