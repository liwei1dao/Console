package console

import (
	"lego_console/comm"
	"lego_console/sys/cache"
	nethttp "net/http"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib/modules/http"
	"github.com/liwei1dao/lego/sys/log"
)

const (
	UserKey string = "UserId"
)

func HttpStatusOK(c *http.Context, code core.ErrorCode, data interface{}) {
	defer log.Debugf("%s resp: code:%d data:%+v", c.Request.RequestURI, code, data)
	c.JSON(nethttp.StatusOK, http.OutJson{ErrorCode: code, Message: comm.GetErrorCodeMsg(code), Data: data})
}

// JWTAuthMiddleware 基于JWT的认证中间件
func CheckToken(c *http.Context) {
	token := c.Request.Header.Get("X-Token")
	if token == "" {
		HttpStatusOK(c, comm.ErrorCode_NoLOgin, nil)
		c.Abort()
		return
	}
	uId, code := cache.QueryToken(token)
	if code != comm.ErrorCode_Success {
		HttpStatusOK(c, code, nil)
		c.Abort()
		return
	}
	c.Set(UserKey, uId)
	c.Next()
}
