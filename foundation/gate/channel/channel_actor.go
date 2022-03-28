package channel

import (
	"foundation/framework/base"
	"foundation/framework/bif"
	"foundation/framework/network"
	"net"
)

var _ bif.IActor = &ChannelActor{}

type ChannelActor struct {
	*network.Conn
	base.Actor
}

func NewChannelActor(conn net.Conn, server *network.Server) network.IDoConn {
	a := &ChannelActor{
		Conn:  network.NewConn(conn, server),
		Actor: base.Actor{},
	}
	//每个1个线程就可以了
	a.Constructor(10, 1)
	v := network.IDoConn(a)
	return v
}

func (a *ChannelActor) Constructor(boxSize int32, maxRunningGoSize int32) {
	a.Actor.Constructor(boxSize, maxRunningGoSize)
	a.Init()
}

func (a *ChannelActor) Init() {

}

func (a *ChannelActor) OnRecv(message any) {

}
