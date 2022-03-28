package ifs

import (
	"foundation/framework/bif"
	"foundation/framework/component"
	"foundation/framework/network"
)

func init() { _ = func(a IKcpComponent) { _ = bif.IComponent(a) } }

//RPCFunc 注册用的回掉函数
type RPCFunc func(uid uint64, cmd uint32, b []byte) error
type IKcpComponent interface {
	Constructor(...interface{})
	Name() component.ComType
	Load()
	Start()
	Tick(in int64)
	Stop()
	Destroy()
	Get(in uint64) *network.Conn

	OnConnect(conn *network.Conn) bool
	OnMessage(conn *network.Conn, packet network.Packet) bool
	OnClose(conn *network.Conn)
}
