package cache

import (
	"lego_console/pb"

	"github.com/liwei1dao/lego/core"
)

type (
	ICache interface {
		//查询token
		QueryToken(token string) (uId uint32, code core.ErrorCode)
		//写入Token
		WriteToken(token string, uId uint32)
		//清理Token缓存
		CleanToken(token string)
		//查询用户缓存数据
		QueryUserData(uId uint32) (data *pb.Cache_UserData, code core.ErrorCode)
		//写入用户缓存数据
		WriteUserData(data *pb.Cache_UserData)
		//清理用户缓存
		CleanUserData(uid uint32)
		//添加新的host监控数据
		AddNewHostMonitor(data *pb.HostMonitor)
		//添加新的Cluster监控数据
		AddNewClusterMonitor(data map[string]*pb.ClusterMonitor)
		//获取host监控数据
		GetHostMonitor(timeleng int32) (result []*pb.HostMonitor)
		//获取Cluster监控数据
		GetClusterMonitor(sIs string, timeleng int32) (result []*pb.ClusterMonitor)
	}
)

var (
	defsys ICache
)

func OnInit(config map[string]interface{}, option ...Option) (err error) {
	defsys, err = newsys(newOptions(config, option...))
	return
}

func NewSys(option ...Option) (sys ICache, err error) {
	sys, err = newsys(newOptionsByOption(option...))
	return
}

//查询token
func QueryToken(token string) (uId uint32, code core.ErrorCode) {
	return defsys.QueryToken(token)
}

//写入Token
func WriteToken(token string, uId uint32) {
	defsys.WriteToken(token, uId)
}

//清理Token缓存
func CleanToken(token string) {
	defsys.CleanToken(token)
}

//查询用户缓存数据
func QueryUserData(uId uint32) (data *pb.Cache_UserData, code core.ErrorCode) {
	return defsys.QueryUserData(uId)
}

//写入用户缓存数据
func WriteUserData(data *pb.Cache_UserData) {
	defsys.WriteUserData(data)
}

//清理用户缓存
func CleanUserData(uid uint32) {
	defsys.CleanUserData(uid)
}

//添加新的host监控数据
func AddNewHostMonitor(data *pb.HostMonitor) {
	defsys.AddNewHostMonitor(data)
}

//添加新的Cluster监控数据
func AddNewClusterMonitor(data map[string]*pb.ClusterMonitor) {
	defsys.AddNewClusterMonitor(data)
}

//获取host监控数据
func GetHostMonitor(timeleng int32) (result []*pb.HostMonitor) {
	return defsys.GetHostMonitor(timeleng)
}

//获取Cluster监控数据
func GetClusterMonitor(sIs string, timeleng int32) (result []*pb.ClusterMonitor) {
	return defsys.GetClusterMonitor(sIs, timeleng)
}
