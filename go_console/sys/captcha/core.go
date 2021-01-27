package captcha

import (
	"lego_console/pb"

	"github.com/liwei1dao/lego/core"
)

type (
	ICaptcha interface {
		//发送验证码到邮件
		SendEmailCaptcha(email, captcha string) (err error)
		//查询验证码
		QueryCaptcha(cId string, ctype pb.CaptchaType) (captcha string, code core.ErrorCode)
		//写入验证码
		WriteCaptcha(cId, captcha string, ctype pb.CaptchaType)
	}
)

var (
	defsys ICaptcha
)

func OnInit(config map[string]interface{}, option ...Option) (err error) {
	defsys, err = newsys(newOptions(config, option...))
	return
}

func NewSys(option ...Option) (sys ICaptcha, err error) {
	sys, err = newsys(newOptionsByOption(option...))
	return
}

//发送验证码到邮件
func SendEmailCaptcha(email, captcha string) (err error) {
	return defsys.SendEmailCaptcha(email, captcha)
}

//查询验证码
func QueryCaptcha(cId string, ctype pb.CaptchaType) (captcha string, code core.ErrorCode) {
	return defsys.QueryCaptcha(cId, ctype)
}

//写入验证码
func WriteCaptcha(cId, captcha string, ctype pb.CaptchaType) {
	defsys.WriteCaptcha(cId, captcha, ctype)
}
