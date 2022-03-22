package ifs

import (
	"foundation/framework/bif"
	"foundation/framework/component"
)

func init() { _ = func(a INatsComponent) { _ = bif.IComponent(a) } }

//RPCFunc 注册用的回掉函数
type RPCFunc func(uid uint64, cmd uint32, b []byte) error
type INatsComponent interface {
	Constructor(...interface{})
	Name() component.ComType
	Load()
	Start()
	Tick(int64 int64)
	Stop()
	Destroy()
	RegisterEvent(cmd uint32, handler RPCFunc)
}
