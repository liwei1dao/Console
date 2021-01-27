package db

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/sys/mgo"
	"go.mongodb.org/mongo-driver/bson"
)

func newsys(options Options) (sys *DB, err error) {
	sys = &DB{options: options}
	if sys.mgo, err = mgo.NewSys(
		mgo.SetMongodbUrl(options.MongodbUrl),
		mgo.SetMongodbDatabase(options.MongodbDatabase),
		mgo.SetTimeOut(options.TimeOut),
	); err == nil {
		err = sys.checkDbInit()
	}
	return
}

type DB struct {
	options Options
	mgo     mgo.IMongodb
}

//校验数据库初始化工作是否完成
func (this DB) checkDbInit() (err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*60)
	count, err := this.mgo.CountDocuments(Sql_SysUserDataIdTable, bson.M{})
	if err != nil || count == 0 {
		//批量插入数据
		leng := 1000000
		cIds := make([]interface{}, leng)
		for i, _ := range cIds {
			cIds[i] = 1000000 + i
		}
		data := make([]interface{}, leng)
		r := rand.New(rand.NewSource(time.Now().Unix()))
		n := 0
		for _, i := range r.Perm(leng) {
			data[n] = bson.M{"_id": i}
			n++
		}
		var (
			err error
		)
		begin := time.Now()
		if _, err = this.mgo.InsertManyByCtx(Sql_SysUserDataIdTable, ctx, data); err != nil {
			return fmt.Errorf("CheckDbInit  err=%s", err.Error())
		}
		log.Debugf("CheckDbInit succ time consuming:%v", time.Now().Sub(begin))
	}
	return
}
