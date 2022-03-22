package actor

import (
	"foundation/framework/component"
	"foundation/framework/component/ifs/ifs.base"
	"foundation/framework/message"
)

func init() { _ = func(a IActor) { _ = a.(IActorRef) } }

type TComponent interface {
	ifs_base.IComponent
}

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
	GetComponent(comType component.ComType) ifs_base.IComponent
	AddComponent(iComponent ifs_base.IComponent, params ...interface{})
}
