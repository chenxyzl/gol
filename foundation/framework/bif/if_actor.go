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
	AddMessage(message message.IMessage)
	OnRecv(message message.IMessage)

	//GetComponent 获取组件
	GetComponent(comType component.ComType) IComponent
	AddComponent(iComponent IComponent, params ...interface{})

	//SafeAsyncDo 安全的异步执行一些事情～ 注意这里不要执行长耗时和异步操作
	SafeAsyncDo(f func())
}
