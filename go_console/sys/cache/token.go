package cache

import (
	"fmt"
	"lego_console/comm"

	"github.com/liwei1dao/lego/core"
)

const ( //Cache
	Cache_Token core.Redis_Key = "Token:%s" //用户数据缓存
)

//查询用户数据
func (this *Cache) QueryToken(token string) (uId uint32, code core.ErrorCode) {
	Id := fmt.Sprintf(string(Cache_Token), token)
	pool := this.redis.GetPool()
	if err := pool.GetKeyForValue(Id, &uId); err != nil {
		code = comm.ErrorCode_NoLOgin
	}
	return
}

//写入Token
func (this *Cache) WriteToken(token string, uId uint32) {
	Id := fmt.Sprintf(string(Cache_Token), token)
	pool := this.redis.GetPool()
	pool.SetExKeyForValue(Id, uId, this.options.TokenCacheExpirationDate)
}

//清理Token
func (this *Cache) CleanToken(token string) {
	Id := fmt.Sprintf(string(Cache_Token), token)
	pool := this.redis.GetPool()
	pool.Delete(Id)
}
