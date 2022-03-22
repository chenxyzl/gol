package home

import (
	"foundation/framework/actor"
	"foundation/framework/component/nats_component"
	"foundation/framework/g"
	"foundation/framework/message"
)

var _ actor.IActor = &RootActor{}

type RootActor struct {
	actor.Actor
}

func NewActor(boxSize int32, maxRunningGoSize int32) *RootActor {
	actor := &RootActor{}
	actor.Constructor(boxSize, maxRunningGoSize)
	actor.RegisterComponent()
	return actor
}

func (actor *RootActor) RegisterComponent() {
	actor.AddComponent(&nats_component.NatsComponent{}, g.GlobalConfig.GetString("Nats.Url"))
}

func (actor *RootActor) OnRecv(message message.IMessage) {
	//消息有多种类型
	//nats消息
}
