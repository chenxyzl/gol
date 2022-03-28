package channel

import (
	"foundation/framework/base"
	"foundation/framework/network"
	"net"
)

//ChannelActor 对外链接的actor，不用实例化
type ChannelActor struct {
	base.Actor
	network.Conn
}

// NewChannelActor returns a wrapper of raw conn
func NewChannelActor(conn net.Conn, srv *network.Server) *ChannelActor {
	v := &ChannelActor{}
	v.Conn.Constructor(conn, srv)
	v.Actor.Constructor(10, 1)
	return v
}

// OnConnect is called when the connection was accepted,
// If the return value of false is closed
func (conn *ChannelActor) OnConnect() bool {
	return true
}

// OnMessage is called when the connection receives a packet,
// If the return value of false is closed
//gate message
func (conn *ChannelActor) OnMessage(packet network.Packet) bool {
	return true
}

//OnRecv rpc message
func (conn *ChannelActor) OnRecv(message any) {

}

// OnClose is called when the connection closed
func (conn *ChannelActor) OnClose() {

}
