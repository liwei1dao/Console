package console

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib"
	"github.com/liwei1dao/lego/lib/modules/http"
	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/utils/crypto"
)

func NewModule() core.IModule {
	m := new(Console)
	return m
}

type Console struct {
	http.Http
	options            *Options
	apicomp            *APIComp
	usercomp           *UserComp
	consolecomp        *ConsoleComp
	Hostmonitorcomp    *HostMonitorComp
	Clustermonitorcomp *ClusterMonitorComp
}

func (this *Console) GetType() core.M_Modules {
	return lib.SM_ConsoleModule
}

func (this *Console) NewOptions() (options core.IModuleOptions) {
	return new(Options)
}

func (this *Console) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	err = this.Http.Init(service, module, options)
	this.options = options.(*Options)
	return
}

func (this *Console) OnInstallComp() {
	this.ModuleBase.OnInstallComp()
	this.apicomp = this.RegisterComp(new(APIComp)).(*APIComp)
	this.usercomp = this.RegisterComp(new(UserComp)).(*UserComp)
	this.consolecomp = this.RegisterComp(new(ConsoleComp)).(*ConsoleComp)
	this.Hostmonitorcomp = this.RegisterComp(new(HostMonitorComp)).(*HostMonitorComp)
	this.Clustermonitorcomp = this.RegisterComp(new(ClusterMonitorComp)).(*ClusterMonitorComp)
}

func (this *Console) ParamSign(param map[string]interface{}) (sign string) {
	var keys []string
	for k, _ := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	builder := strings.Builder{}
	for _, v := range keys {
		builder.WriteString(v)
		builder.WriteString("=")
		switch reflect.TypeOf(param[v]).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
			builder.WriteString(fmt.Sprintf("%d", param[v]))
			break
		default:
			builder.WriteString(fmt.Sprintf("%s", param[v]))
			break
		}
		builder.WriteString("&")
	}
	builder.WriteString("key=" + this.options.SignKey)
	log.Infof("orsign:%s", builder.String())
	sign = crypto.MD5EncToLower(builder.String())
	return
}

//创建token
func (this *Console) createToken(uId uint32) (token string) {
	orsign := fmt.Sprintf("%d|%d", uId, time.Now().Unix())
	token = crypto.AesEncryptCBC(orsign, this.options.SignKey)
	return
}

//校验token
func (this *Console) checkToken(token string, uId uint32) (check bool) {
	orsign := crypto.AesDecryptCBC(token, this.options.SignKey)
	tempstr := strings.Split(orsign, "|")
	if len(tempstr) == 2 && tempstr[0] == fmt.Sprintf("%d", uId) {
		return true
	}
	return false
}
