package bif

import (
	"foundation/framework/component"
	"foundation/framework/message"
)

func init() { _ = func(a IActor) { _ = a.(IActorRef) } }

type IActor interface {
	GetUid() uint64
	GetActorType() string
	Load()
	Start()
	Tick(int642 int64)
	Stop()
	Destroy()
	OnRecv(message message.IMessage)

	//GetComponent 获取组件
	GetComponent(comType component.ComType) IComponent
	AddComponent(iComponent IComponent, params ...interface{})
}
