package db

import (
	"lego_console/comm"
	"lego_console/pb"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/sys/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Sql_SysUserDataIdTable core.SqlTable = "usersid" //话数据
	Sql_SysUserDataTable   core.SqlTable = "users"
)

type UserId struct {
	Id uint32 `bson:"_id"`
}

//查询用户 Id
func (this *DB) QueryUserDataById(uid uint32) (result *pb.DB_UserData, code core.ErrorCode) {
	result = &pb.DB_UserData{}
	if err := this.mgo.FindOne(Sql_SysUserDataTable, bson.M{"_id": uid}).Decode(result); err != nil {
		log.Errorf("QueryUserData err:%v", err)
		code = comm.ErrorCode_SqlExecutionError
	}
	return
}

//查询用户 Id
func (this *DB) QueryUserDataByPhonOrEmail(phonoremail string) (result *pb.DB_UserData, code core.ErrorCode) {
	result = &pb.DB_UserData{}
	if err := this.mgo.FindOne(Sql_SysUserDataTable, bson.M{"phonoremail": phonoremail}).Decode(result); err != nil {
		code = comm.ErrorCode_SqlExecutionError
	}
	return
}

//查询用户 账号
func (this *DB) LoginUserDataByPhonOrEmail(data *pb.DB_UserData) (result *pb.DB_UserData, code core.ErrorCode) {
	result = &pb.DB_UserData{}
	if err := this.mgo.FindOne(Sql_SysUserDataTable, bson.M{"phonoremail": data.PhonOrEmail}).Decode(result); err != nil {
		return this.registeredUser(data)
	}
	return
}

//更新用户信息
func (this *DB) UpDataUserData(data *pb.DB_UserData) (code core.ErrorCode) {
	if _, err := this.mgo.UpdateOne(Sql_SysUserDataTable,
		bson.M{"_id": data.Id},
		bson.M{"$set": bson.M{
			"nickname": data.NickName,
			"headurl":  data.HeadUrl,
			"userrole": data.UserRole,
			"token":    data.Token,
		}},
		new(options.UpdateOptions).SetUpsert(true)); err != nil {
		log.Errorf("UpDataUserData err:%v", err)
		code = comm.ErrorCode_SqlExecutionError
	}
	return comm.ErrorCode_Success
}

//注册账号
func (this *DB) registeredUser(data *pb.DB_UserData) (result *pb.DB_UserData, code core.ErrorCode) {
	var newuId UserId
	if err := this.mgo.FindOneAndDelete(Sql_SysUserDataIdTable, bson.D{}).Decode(&newuId); err != nil {
		log.Warnf("registeredUser err=%v", err)
		code = comm.ErrorCode_SqlExecutionError
		return
	}
	result = &pb.DB_UserData{
		Id:          newuId.Id,
		PhonOrEmail: data.PhonOrEmail,
		Password:    data.Password,
		NickName:    data.NickName,
		HeadUrl:     data.HeadUrl,
	}
	if _, err := this.mgo.InsertOne(Sql_SysUserDataTable, result); err != nil {
		log.Errorf("registeredUser err=%v", err)
		code = comm.ErrorCode_SqlExecutionError
	}
	return
}
