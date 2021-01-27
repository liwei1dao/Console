package cache

import (
	"lego_console/pb"
	"reflect"

	"github.com/liwei1dao/lego/core"
)

const ( //Cache
	Cache_HostMonitor core.Redis_Key = "HostMonitor" //主机监听数据
)

//添加新的HostMonitor
func (this *Cache) AddNewHostMonitor(data *pb.HostMonitor) {
	pool := this.redis.GetPool()
	pool.SetListByRPush(string(Cache_HostMonitor), []interface{}{data})
	if len, err := pool.GetListCount(string(Cache_HostMonitor)); err == nil && len > this.options.MonitorTotalTime {
		pool.GetListByLPop(string(Cache_HostMonitor), data)
	}
}

//添加新的HostMonitor
func (this *Cache) GetHostMonitor(timeleng int32) (result []*pb.HostMonitor) {
	result = make([]*pb.HostMonitor, 0)
	pool := this.redis.GetPool()
	value := pool.GetListByLrange(string(Cache_HostMonitor), 0, timeleng, reflect.TypeOf(&pb.HostMonitor{}))
	if value != nil && len(value) > 0 {
		result = make([]*pb.HostMonitor, len(value))
		for i, v := range value {
			result[i] = v.(*pb.HostMonitor)
		}
	}
	return
}
