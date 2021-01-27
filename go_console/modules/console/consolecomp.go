package console

import (
	"lego_console/comm"
	"lego_console/pb"
	"lego_console/sys/cache"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/lib/modules/http"
	"github.com/liwei1dao/lego/sys/log"
)

type ConsoleComp struct {
	cbase.ModuleCompBase
	module *Console
}

func (this *ConsoleComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	err = this.ModuleCompBase.Init(service, module, comp, options)
	this.module = module.(*Console)
	group := this.module.Group("/lego/console")
	group.POST("/getprojectinfo", CheckToken, this.GetProjectInfo)
	group.POST("/gethostinfo", CheckToken, this.GetHostInfo)
	group.POST("/getcpuinfo", CheckToken, this.GetCpuInfo)
	group.POST("/getmemoryinfo", CheckToken, this.GetMemoryInfo)
	group.POST("/gethostmonitordata", CheckToken, this.GetHostMonitorData)
	group.POST("/getconsolecluster", CheckToken, this.GetClusterMonitorData)
	return
}

//获取集群信息
func (this *ConsoleComp) GetProjectInfo(c *http.Context) {
	uId := c.GetUInt32(UserKey)
	defer log.Debugf("GetClusterInfo:%d", uId)
	var (
		code  core.ErrorCode
		udata *pb.Cache_UserData
	)
	if udata, code = cache.QueryUserData(uId); code == comm.ErrorCode_Success && udata.Db_UserData.UserRole >= pb.UserRole_Developer {
		HttpStatusOK(c, code, this.module.options)
		return
	} else {
		HttpStatusOK(c, comm.ErrorCode_AuthorityLow, nil)
		return
	}
}

//获取控制台信息
func (this *ConsoleComp) GetHostInfo(c *http.Context) {
	uId := c.GetUInt32(UserKey)
	defer log.Debugf("GetHostInfo:%d", uId)
	var (
		code  core.ErrorCode
		udata *pb.Cache_UserData
	)
	if udata, code = cache.QueryUserData(uId); code == comm.ErrorCode_Success && udata.Db_UserData.UserRole >= pb.UserRole_Developer {
		hostInfo := this.module.Hostmonitorcomp.hostInfo
		HttpStatusOK(c, code, hostInfo)
		return
	} else {
		HttpStatusOK(c, comm.ErrorCode_AuthorityLow, nil)
		return
	}
}

//获取控制台信息
func (this *ConsoleComp) GetCpuInfo(c *http.Context) {
	uId := c.GetUInt32(UserKey)
	defer log.Debugf("GetHostInfo:%d", uId)
	var (
		code  core.ErrorCode
		udata *pb.Cache_UserData
	)
	if udata, code = cache.QueryUserData(uId); code == comm.ErrorCode_Success && udata.Db_UserData.UserRole >= pb.UserRole_Developer {
		cpuinfo := this.module.Hostmonitorcomp.cpuInfo
		HttpStatusOK(c, code, cpuinfo)
		return
	} else {
		HttpStatusOK(c, comm.ErrorCode_AuthorityLow, nil)
		return
	}
}

//获取控制台信息
func (this *ConsoleComp) GetMemoryInfo(c *http.Context) {
	uId := c.GetUInt32(UserKey)
	defer log.Debugf("GetconSoleinfo:%d", uId)
	var (
		code  core.ErrorCode
		udata *pb.Cache_UserData
	)
	if udata, code = cache.QueryUserData(uId); code == comm.ErrorCode_Success && udata.Db_UserData.UserRole >= pb.UserRole_Developer {
		memoryInfo := this.module.Hostmonitorcomp.memoryInfo
		HttpStatusOK(c, code, memoryInfo)
		return
	} else {
		HttpStatusOK(c, comm.ErrorCode_AuthorityLow, nil)
		return
	}
}

//获取主机监控信息
func (this *ConsoleComp) GetHostMonitorData(c *http.Context) {
	req := &pb.QueryHostMonitorDataReq{}
	c.ShouldBindJSON(req)
	uId := c.GetUInt32(UserKey)
	defer log.Debugf("GetHostMonitorData:%v", req)
	var (
		code  core.ErrorCode
		udata *pb.Cache_UserData
	)
	if udata, code = cache.QueryUserData(uId); code == comm.ErrorCode_Success && udata.Db_UserData.UserRole >= pb.UserRole_Developer {
		resp := this.module.Hostmonitorcomp.GetHostMonitorData(req.QueryTime)
		HttpStatusOK(c, code, resp)
		return
	} else {
		HttpStatusOK(c, comm.ErrorCode_AuthorityLow, nil)
		return
	}
}

//获取控制台信息
func (this *ConsoleComp) GetClusterMonitorData(c *http.Context) {
	req := &pb.QueryClusterMonitorDataReq{}
	c.ShouldBindJSON(req)
	uId := c.GetUInt32(UserKey)
	defer log.Debugf("GetConsoleCluster:%d", uId)
	var (
		code  core.ErrorCode
		udata *pb.Cache_UserData
	)
	if udata, code = cache.QueryUserData(uId); code == comm.ErrorCode_Success && udata.Db_UserData.UserRole >= pb.UserRole_Developer {
		resp := this.module.Clustermonitorcomp.GetClusterMonitorDataResp(req.QueryTime)
		HttpStatusOK(c, code, resp)
		return
	} else {
		HttpStatusOK(c, comm.ErrorCode_AuthorityLow, nil)
		return
	}
}
