package services

import (
	"fmt"
	"lego_console/sys/cache"
	"lego_console/sys/captcha"
	"lego_console/sys/db"

	"github.com/liwei1dao/lego/sys/log"

	"github.com/liwei1dao/lego/base/cluster"
)

type BaseService struct {
	cluster.ClusterService
}

func (this *BaseService) InitSys() {
	this.ClusterService.InitSys()
	if err := cache.OnInit(this.Service.GetSettings().Sys["cache"]); err != nil {
		panic(fmt.Sprintf("statr cache err:%v", err))
	} else {
		log.Infof("Sys cache Init success !")
	}
	if err := db.OnInit(this.Service.GetSettings().Sys["db"]); err != nil {
		panic(fmt.Sprintf("statr db err:%v %s", err))
	} else {
		log.Infof("Sys db Init success !")
	}
	if err := captcha.OnInit(this.Service.GetSettings().Sys["captcha"]); err != nil {
		panic(fmt.Sprintf("statr captcha err:%v %s", err.Error()))
	} else {
		log.Infof("Sys captcha Init success !")
	}
}
