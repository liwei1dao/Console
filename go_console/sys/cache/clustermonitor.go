package cache

import (
	"fmt"
	"lego_console/pb"
	"reflect"

	"github.com/liwei1dao/lego/core"
)

const ( //Cache
	Cache_ClusterMonitor core.Redis_Key = "ClusterMonitor:%s" //集群监听数据
)

//添加新的ClusterMonitor
func (this *Cache) AddNewClusterMonitor(data map[string]*pb.ClusterMonitor) {
	pool := this.redis.GetPool()
	for k, v := range data {
		id := fmt.Sprintf(string(Cache_ClusterMonitor), k)
		pool.SetListByRPush(id, []interface{}{v})
		if len, err := pool.GetListCount(id); err == nil && len > this.options.MonitorTotalTime {
			pool.GetListByLPop(string(Cache_ClusterMonitor), v)
		}
	}
}

//添加新的ClusterMonitor
func (this *Cache) GetClusterMonitor(sIs string, timeleng int32) (result []*pb.ClusterMonitor) {
	result = make([]*pb.ClusterMonitor, 0)
	id := fmt.Sprintf(string(Cache_ClusterMonitor), sIs)
	pool := this.redis.GetPool()
	value := pool.GetListByLrange(id, 0, timeleng, reflect.TypeOf(&pb.ClusterMonitor{}))
	if value != nil && len(value) > 0 {
		result = make([]*pb.ClusterMonitor, len(value))
		for i, v := range value {
			result[i] = v.(*pb.ClusterMonitor)
		}
	}
	return
}
