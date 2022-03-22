package g

import (
	"foundation/framework/actor"
	"foundation/framework/util"
	"github.com/spf13/viper"
)

type RoleType string

const (
	RoleTypeKnown RoleType = "unknown"
	Login         RoleType = "login"
	Gate          RoleType = "gate"
	Home          RoleType = "home"
	World         RoleType = "world"
)

var (
	DebugModel   = true          //是否调试模式
	_clusterName string          //集群名字
	_roleType    = RoleTypeKnown //进程的类型
	_Id          uint64          //进程在集群的id 最大值为2^12-1

	GlobalConfig       *viper.Viper //全局配置
	EnableWdsyncMD5Map bool         //wdsync是否开启md5
	PProf              string       //pprof性能检测
	EtcdAddr           string       //etcd地址

	Root actor.IActor
	UUID *util.UUID
)

func Init(Id1 uint64) {
	if _roleType == "" {
		panic("must set roleType")
	}
	//初始化类型和id
	_clusterName = GlobalConfig.GetString("Global.Name")
	_Id = Id1
	//初始化uuid
	u, err := util.NewUUID(_Id)
	if err != nil {
		panic(err)
	}
	UUID = u
	//初始化其他
}

func SetRoleType(r RoleType) {
	_roleType = r
}

func RoleTyp() RoleType {
	return _roleType
}

func Id() uint64 {
	return _Id
}
