package console

import (
	"github.com/liwei1dao/lego/lib/modules/http"
	"github.com/liwei1dao/lego/utils/mapstructure"
)

type Options struct {
	http.Options
	SignKey             string                 `json:"-"` //签名密钥
	UserInitialPassword string                 `json:"-"` //用户初始密码
	ProjectName         string                 //项目名称
	ProjectDes          string                 //项目描述
	ProjectVersion      float64                //版本
	ProjectTime         string                 //立项时间
	ProjectMember       map[string]interface{} //项目经理 成员-职务
}

func (this *Options) LoadConfig(settings map[string]interface{}) (err error) {
	if err = this.Options.LoadConfig(settings); err == nil {
		if settings != nil {
			err = mapstructure.Decode(settings, this)
		}
	}
	return
}
