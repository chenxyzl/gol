package home

import (
	"foundation/framework/base"
	"foundation/framework/bif"
	"foundation/framework/component/nats_component"
	"foundation/framework/g"
	"foundation/framework/message"
	"foundation/home/easyrpc"
	"foundation/home/easyrpcimpl"
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
	easyrpc.RegisterPlayerService(&easyrpcimpl.PlayerRPCService{})
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
func (actor *RootActor) OnRecv(message message.IMessage) {
	//消息有多种类型
	//nats消息
}
