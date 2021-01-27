package comm

import (
	"fmt"

	"github.com/liwei1dao/lego/core"
)

const (
	ErrorCode_Success core.ErrorCode = 0 //成功

	ErrorCode_NoFindService         core.ErrorCode = 10 //没有找到远程服务器
	ErrorCode_RpcFuncExecutionError core.ErrorCode = 11 //Rpc方法执行错误
	ErrorCode_CacheReadError        core.ErrorCode = 12 //缓存读取失败
	ErrorCode_SqlExecutionError     core.ErrorCode = 13 //数据库执行错误
	ErrorCode_ReqParameterError     core.ErrorCode = 14 //请求参数错误
	ErrorCode_SignError             core.ErrorCode = 15 //签名错误
	ErrorCode_NoLOgin               core.ErrorCode = 16 //未登录
	ErrorCode_AuthorityLow          core.ErrorCode = 17 //权限不足

	//Login登录错误码
	ErrorCode_AlreadyRegister             core.ErrorCode = 2001 //帐号已被注册
	ErrorCode_LoginAccountOrPasswordError core.ErrorCode = 2002 //登录帐号密码错误
	ErrorCode_UserCookieLose              core.ErrorCode = 2004 //用户Cookie数据丢失
	ErrorCode_UserDataError               core.ErrorCode = 2005 //用户数据错误
)

var ErrorCodeMsg = map[core.ErrorCode]string{
	ErrorCode_Success:                     "成功",
	ErrorCode_NoFindService:               "没有找到远程服务器",
	ErrorCode_RpcFuncExecutionError:       "Rpc方法执行错误",
	ErrorCode_CacheReadError:              "缓存读取失败",
	ErrorCode_SqlExecutionError:           "数据库执行错误",
	ErrorCode_ReqParameterError:           "请求参数错误",
	ErrorCode_SignError:                   "签名错误",
	ErrorCode_AlreadyRegister:             "帐号已被注册",
	ErrorCode_LoginAccountOrPasswordError: "登录帐号密码错误",
	ErrorCode_UserCookieLose:              "用户Cookie数据丢失",
	ErrorCode_AuthorityLow:                "您的权限不足",
}

func GetErrorCodeMsg(code core.ErrorCode) string {
	if v, ok := ErrorCodeMsg[code]; ok {
		return v
	} else {
		return fmt.Sprintf("ErrorCode:%d", code)
	}
}
