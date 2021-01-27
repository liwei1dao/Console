package cache

import (
	"fmt"
	"lego_console/comm"
	"lego_console/pb"
	"lego_console/sys/db"

	"github.com/liwei1dao/lego/core"
)

const ( //Cache
	Cache_Users core.Redis_Key = "Users:%d" //用户数据缓存
)

//查询用户数据
func (this *Cache) QueryUserData(uId uint32) (result *pb.Cache_UserData, code core.ErrorCode) {
	Id := fmt.Sprintf(string(Cache_Users), uId)
	pool := this.redis.GetPool()
	result = &pb.Cache_UserData{}
	if err := pool.GetKeyForValue(Id, result); err != nil {
		result, code = this.synchronizeUserToCache(uId)
	}
	return
}

//同步用户数据到缓存
func (this *Cache) synchronizeUserToCache(uId uint32) (result *pb.Cache_UserData, code core.ErrorCode) {
	var user *pb.DB_UserData
	if user, code = db.QueryUserDataById(uId); code == comm.ErrorCode_Success {
		result = &pb.Cache_UserData{
			Db_UserData: user,
			IsOnLine:    false,
		}
		this.writeUserDataByEx(result)
	}
	return
}

//离线用户缓存读取之后保存10分钟
func (this *Cache) writeUserDataByEx(result *pb.Cache_UserData) {
	Id := fmt.Sprintf(string(Cache_Users), result.Db_UserData.Id)
	pool := this.redis.GetPool()
	pool.SetExKeyForValue(Id, result, this.options.UserCacheExpirationDate)
}

//登录用户缓存信息长期驻留
func (this *Cache) WriteUserData(data *pb.Cache_UserData) {
	Id := fmt.Sprintf(string(Cache_Users), data.Db_UserData.Id)
	pool := this.redis.GetPool()
	pool.SetKeyForValue(Id, data)
}

//清理用户缓存
func (this *Cache) CleanUserData(uid uint32) {
	Id := fmt.Sprintf(string(Cache_Users), uid)
	pool := this.redis.GetPool()
	pool.Delete(Id)
}
