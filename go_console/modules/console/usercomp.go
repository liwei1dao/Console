package console

import (
	"lego_console/comm"
	"lego_console/pb"
	"lego_console/sys/cache"
	"lego_console/sys/captcha"
	"lego_console/sys/db"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/lib/modules/http"
	"github.com/liwei1dao/lego/sys/log"
)

type UserComp struct {
	cbase.ModuleCompBase
	module *Console
}

func (this *UserComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	err = this.ModuleCompBase.Init(service, module, comp, options)
	this.module = module.(*Console)
	group := this.module.Group("/lego/user")
	group.POST("/loginbycaptcha", this.LoginByCaptchaReq)
	group.POST("/loginbypassword", this.LoginByPasswordReq)
	group.POST("/getuserinfo", CheckToken, this.GetUserinfoReq)
	group.POST("/loginout", CheckToken, this.LoginOutReq)
	return
}

//登录请求 验证码
func (this *UserComp) LoginByCaptchaReq(c *http.Context) {
	req := &pb.LoginByCaptchaReq{}
	c.ShouldBindJSON(req)
	defer log.Debugf("LoginByCaptchaReq:%+v", req)
	if sgin := this.module.ParamSign(map[string]interface{}{"PhonOrEmail": req.PhonOrEmail, "Captcha": req.Captcha}); sgin != req.Sign {
		log.Errorf("LoginByCaptchaReq SignError sgin:%s", sgin)
		HttpStatusOK(c, comm.ErrorCode_SignError, nil)
		return
	}
	var (
		captchecode string
		token       string
		code        core.ErrorCode
		udata       *pb.DB_UserData
	)
	if captchecode, code = captcha.QueryCaptcha(req.PhonOrEmail, pb.CaptchaType_LoginCaptcha); code == comm.ErrorCode_Success && captchecode == req.Captcha {
		if udata, code = db.LoginUserDataByPhonOrEmail(&pb.DB_UserData{
			PhonOrEmail: req.PhonOrEmail,
			Password:    this.module.options.UserInitialPassword,
			NickName:    req.PhonOrEmail,
			UserRole:    pb.UserRole_Guester,
		}); code == comm.ErrorCode_Success {
			cache.CleanToken(udata.Token)
			token = this.module.createToken(udata.Id)
			udata.Token = token
			db.UpDataUserData(udata)
			cache.WriteToken(token, udata.Id)
			cache.WriteUserData(&pb.Cache_UserData{
				Db_UserData: udata,
				IsOnLine:    true,
			})
			HttpStatusOK(c, code, token)
			return
		}
	}
	HttpStatusOK(c, code, nil)
	return
}

//登录请求 密码
func (this *UserComp) LoginByPasswordReq(c *http.Context) {
	req := &pb.LoginByPasswordReq{}
	c.ShouldBindJSON(req)
	defer log.Debugf("LoginByPasswordReq:%+v", req)
	if sgin := this.module.ParamSign(map[string]interface{}{"PhonOrEmail": req.PhonOrEmail, "Password": req.Password}); sgin != req.Sign {
		log.Errorf("LoginByPasswordReq SignError sgin:%s", sgin)
		HttpStatusOK(c, comm.ErrorCode_SignError, nil)
		return
	}
	var (
		code  core.ErrorCode
		token string
		udata *pb.DB_UserData
	)
	if udata, code = db.QueryUserDataByPhonOrEmail(req.PhonOrEmail); code == comm.ErrorCode_Success {
		if udata.Password == req.Password {
			cache.CleanToken(udata.Token)
			token = this.module.createToken(udata.Id)
			udata.Token = token
			db.UpDataUserData(udata)
			cache.WriteToken(token, udata.Id)
			cache.WriteUserData(&pb.Cache_UserData{
				Db_UserData: udata,
				IsOnLine:    true,
			})
			HttpStatusOK(c, code, token)
		}
	}
	HttpStatusOK(c, code, nil)
	return
}

//登录请求 Token
func (this *UserComp) GetUserinfoReq(c *http.Context) {
	uId := c.GetUInt32(UserKey)
	defer log.Debugf("GetUserinfoReq:%d", uId)
	var (
		code  core.ErrorCode
		udata *pb.DB_UserData
	)
	if udata, code = db.QueryUserDataById(uId); code == comm.ErrorCode_Success {
		cache.WriteUserData(&pb.Cache_UserData{
			Db_UserData: udata,
			IsOnLine:    true,
		})
		HttpStatusOK(c, code, udata)
		return
	}
	HttpStatusOK(c, code, nil)
	return
}

//登出
func (this *UserComp) LoginOutReq(c *http.Context) {
	req := &pb.LoginOutReq{}
	c.ShouldBindJSON(req)
	uId := c.GetUInt32(UserKey)
	defer log.Debugf("LoginOutReq:%+v UserId:%d", req, uId)
	var (
		result *pb.DB_UserData
		code   core.ErrorCode
	)
	if result, code = db.QueryUserDataById(uId); code == comm.ErrorCode_Success {
		cache.CleanUserData(uId)
		cache.CleanToken(result.Token)
	}
	HttpStatusOK(c, code, nil)
}
