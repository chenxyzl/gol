package network

import "net"

type DefaultConn struct {
	Conn
}

// NewDefaultConn returns a wrapper of raw conn
func NewDefaultConn(conn net.Conn, srv *Server) *DefaultConn {
	v := &DefaultConn{}
	v.Constructor(conn, srv)
	return v
}

// OnConnect is called when the connection was accepted,
// If the return value of false is closed
func (conn *DefaultConn) OnConnect() bool {
	return true
}

// OnMessage is called when the connection receives a packet,
// If the return value of false is closed
func (conn *DefaultConn) OnMessage(packet Packet) bool {
	return true
}

// OnClose is called when the connection closed
func (conn *DefaultConn) OnClose() {

}
