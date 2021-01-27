package db

import (
	"lego_console/pb"

	"github.com/liwei1dao/lego/core"
)

type (
	IDB interface {
		//查询用户 Id
		QueryUserDataById(uid uint32) (result *pb.DB_UserData, code core.ErrorCode)
		//查询用户 PhonOrEmail
		QueryUserDataByPhonOrEmail(PhonOrEmail string) (result *pb.DB_UserData, code core.ErrorCode)
		//查询用户 email
		LoginUserDataByPhonOrEmail(data *pb.DB_UserData) (result *pb.DB_UserData, code core.ErrorCode)
		//更新用户数据
		UpDataUserData(data *pb.DB_UserData) (code core.ErrorCode)
	}
)

var (
	defsys IDB
)

func OnInit(config map[string]interface{}, option ...Option) (err error) {
	defsys, err = newsys(newOptions(config, option...))
	return
}

func NewSys(option ...Option) (sys IDB, err error) {
	sys, err = newsys(newOptionsByOption(option...))
	return
}

//查询用户 Id
func QueryUserDataById(uid uint32) (result *pb.DB_UserData, code core.ErrorCode) {
	return defsys.QueryUserDataById(uid)
}

//查询用户 Id
func QueryUserDataByPhonOrEmail(PhonOrEmail string) (result *pb.DB_UserData, code core.ErrorCode) {
	return defsys.QueryUserDataByPhonOrEmail(PhonOrEmail)
}

//查询用户 email
func LoginUserDataByPhonOrEmail(data *pb.DB_UserData) (result *pb.DB_UserData, code core.ErrorCode) {
	return defsys.LoginUserDataByPhonOrEmail(data)
}

//更新用户数据
func UpDataUserData(data *pb.DB_UserData) (code core.ErrorCode) {
	return defsys.UpDataUserData(data)
}
