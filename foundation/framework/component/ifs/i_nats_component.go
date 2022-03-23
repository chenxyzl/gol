package ifs

import (
	"foundation/framework/bif"
	"foundation/framework/component"
	"foundation/message"
)

func init() { _ = func(a INatsComponent) { _ = bif.IComponent(a) } }

//RPCFunc 注册用的回掉函数
type RPCFunc func(req *message.NatsRequest) message.Code
type INatsComponent interface {
	Constructor(...interface{})
	Name() component.ComType
	Load()
	Start()
	Tick(int64 int64)
	Stop()
	Destroy()
	RegisterEvent(cmd uint32, handler RPCFunc, hasReply bool)
	Dispatch(req *message.NatsRequest)
	Reply(url string, reply *message.NatsReply)
}
