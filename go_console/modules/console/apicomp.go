package console

import (
	"fmt"
	"lego_console/comm"
	"lego_console/pb"
	"lego_console/sys/captcha"
	"math/rand"
	"time"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/lib/modules/http"
	"github.com/liwei1dao/lego/sys/log"
)

type APIComp struct {
	cbase.ModuleCompBase
	module *Console
}

func (this *APIComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	err = this.ModuleCompBase.Init(service, module, comp, options)
	this.module = module.(*Console)
	group := this.module.Group("/lego/api")
	group.POST("/sendemailcaptcha", this.SendEmailCaptchaReq)
	return
}

//获取验证码
func (this *APIComp) SendEmailCaptchaReq(c *http.Context) {
	req := &pb.SendEmailCaptchaReq{}
	c.ShouldBindJSON(req)
	defer log.Debugf("SendEmailCaptchaReq:%+v", req)
	if sgin := this.module.ParamSign(map[string]interface{}{"Mailbox": req.Mailbox, "CaptchaType": req.CaptchaType}); sgin != req.Sign {
		log.Errorf("LoginByCaptchaReq SignError sgin:%s", sgin)
		HttpStatusOK(c, comm.ErrorCode_SignError, nil)
		return
	}
	if req.Mailbox == "" {
		HttpStatusOK(c, comm.ErrorCode_ReqParameterError, nil)
		return
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	captchacode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	if err := captcha.SendEmailCaptcha(req.Mailbox, captchacode); err == nil {
		captcha.WriteCaptcha(req.Mailbox, captchacode, req.CaptchaType)
		HttpStatusOK(c, comm.ErrorCode_Success, nil)
	} else {
		HttpStatusOK(c, comm.ErrorCode_ReqParameterError, nil)
	}
}
