package home

import (
	"foundation/framework/base"
	"foundation/framework/bif"
	"foundation/framework/component"
	"foundation/framework/component/ifs"
	"foundation/framework/component/nats_component"
	"foundation/framework/g"
	"foundation/home/playerrpc"
	"foundation/home/playerrpcimpl"
	message2 "foundation/message"
	"gitlab-ee.funplus.io/watcher/watcher/misc/wlog"
	"reflect"
)

var _ bif.IActor = &RootActor{}

type RootActor struct {
	base.Actor
}

func NewActor(boxSize int32, maxRunningGoSize int32) *RootActor {
	actor := &RootActor{}
	actor.Constructor(boxSize, maxRunningGoSize)
	return actor
}

func (actor *RootActor) RegisterComponent() {
	actor.AddComponent(&nats_component.NatsComponent{}, g.GlobalConfig.GetString("Nats.Url"))
}

func (actor *RootActor) RegisterRpc() {
	playerrpc.RegisterPlayerService(&playerrpcimpl.PlayerRPCService{})
}

//Load 生命周期函数
func (actor *RootActor) Load() {
	//先注册component
	actor.RegisterComponent()
	//调用基类的load
	actor.Actor.Load()
	//再调用rpc注册
	actor.RegisterRpc()
}
func (actor *RootActor) OnRecv(msg any) {
	switch v := msg.(type) {
	case *message2.NatsRequest:
		wlog.Debug("nats rpc msg")
		c := g.Root.GetComponent(component.NatsCom).(ifs.INatsComponent)
		c.Dispatch(v)
	case *message2.Request: //只有gateway才会收到这种消息
		wlog.Debug("client rpc msg")
	default:
		wlog.Warnf("unknown msg type:%v", reflect.TypeOf(msg).Name())
	}
}
