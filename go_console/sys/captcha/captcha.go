package captcha

import (
	"fmt"
	"lego_console/comm"
	"lego_console/pb"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/sys/redis"
	"github.com/liwei1dao/lego/sys/sdks/email"
)

func newsys(options Options) (sys *Captcha, err error) {
	sys = &Captcha{options: options}
	if sys.email, err = email.NewSys(
		email.SetServerhost(options.Serverhost),
		email.SetFromemail(options.Fromemail),
		email.SetFompasswd(options.Fompasswd),
		email.SetServerport(options.Serverport),
	); err != nil {
		return
	}
	if sys.redis, err = redis.NewSys(redis.SetRedisUrl(options.RedisUrl)); err != nil {
		return
	}
	return
}

const ( //Cache
	Cache_Captcha core.Redis_Key = "Captcha:%s-%d" //验证码缓存
)

type Captcha struct {
	options Options
	email   email.IEmail
	redis   redis.IRedisFactory
}

//发送邮箱验证码
func (this *Captcha) SendEmailCaptcha(email, captcha string) (err error) {
	return this.email.SendMail(email, "Verification Code", fmt.Sprintf("Your CCPC verification code:%s, please do not forward others", captcha))
}

//获取验证码
func (this *Captcha) QueryCaptcha(cId string, ctype pb.CaptchaType) (captcha string, code core.ErrorCode) {
	Id := fmt.Sprintf(string(Cache_Captcha), cId, ctype)
	pool := this.redis.GetPool()
	if err := pool.GetKeyForValue(Id, &captcha); err == nil {
		return captcha, comm.ErrorCode_Success
	} else {
		return "", comm.ErrorCode_CacheReadError
	}
}

//写入验证码
func (this *Captcha) WriteCaptcha(cId, captcha string, ctype pb.CaptchaType) {
	Id := fmt.Sprintf(string(Cache_Captcha), cId, ctype)
	pool := this.redis.GetPool()
	pool.SetExKeyForValue(Id, captcha, this.options.CaptchaExpirationdate)
}
