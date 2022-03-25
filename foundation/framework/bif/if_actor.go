package bif

import (
	"foundation/framework/component"
	"foundation/message"
)

func init() { _ = func(a IActor) { _ = message.IActorRef(a) } }

type IActor interface {
	GetUid() uint64
	GetType() string
	To() *message.ActorRef
	Load()
	Start()
	Tick(int642 int64)
	Stop()
	Destroy()
	AddMessage(message any)
	OnRecv(message any)

	//GetComponent 获取组件
	GetComponent(comType component.ComType) IComponent
	AddComponent(iComponent IComponent, params ...interface{})

	//SafeAsyncDo 安全的异步执行一些事情～ 注意这里不要执行长耗时和异步操作
	SafeAsyncDo(f func())
}
