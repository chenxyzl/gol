package kcpcomponent

import (
	"foundation/framework/alg/skiplist"
	"foundation/framework/bif"
	"foundation/framework/component"
	"foundation/framework/g"
	"foundation/framework/network"
	component2 "foundation/gate/global_component"
	"foundation/gate/global_component/ifs"
	"foundation/message"
	"github.com/golang/protobuf/proto"
	"github.com/xtaci/kcp-go"
	"gitlab-ee.funplus.io/watcher/watcher/misc/wlog"
	"net"
	"sync"
	"time"
)

var _ ifs.IKcpComponent = &KcpComponent{}

type comparerConn struct {
}

func (compare *comparerConn) CmpScore(v1 interface{}, v2 interface{}) int {
	s1 := v1.(uint64)
	s2 := v2.(uint64)
	switch {
	case s1 < s2:
		return -1
	case s1 == s2:
		return 0
	default:
		return 1
	}
}

func (compare *comparerConn) CmpKey(v1 interface{}, v2 interface{}) int {
	s1 := v1.(uint64)
	s2 := v2.(uint64)
	switch {
	case s1 < s2:
		return -1
	case s1 == s2:
		return 0
	default:
		return 1
	}
}

const (
	authWaitTime = time.Second * 15
)

type KcpComponent struct {
	Node        bif.IActor
	server      *network.Server
	addr        string //":10086"
	mtx         sync.Mutex
	cons        map[uint64]*network.Conn //session,conn
	connMapping map[uint64]uint64        //uid,session
	authCons    *skiplist.SkipList
}

func (c *KcpComponent) Constructor(params ...interface{}) {
	c.addr = params[0].(string)
	c.authCons = skiplist.NewSkipList(&comparerConn{})
	c.cons = make(map[uint64]*network.Conn)
	c.connMapping = make(map[uint64]uint64)
}
func (c *KcpComponent) Name() component.ComType {
	return component2.KcpCom
}
func (c *KcpComponent) Load() {
	server, err := network.StartKcpServer(c.addr, &network.DefaultProtocol{}, nil, func(conn net.Conn, server *network.Server) network.IConn {
		kcpConn := conn.(*kcp.UDPSession)
		kcpConn.SetNoDelay(1, 10, 2, 1)
		kcpConn.SetStreamMode(true)
		kcpConn.SetWindowSize(32, 32)
		kcpConn.SetReadBuffer(1024)
		kcpConn.SetWriteBuffer(1024)
		kcpConn.SetMtu(1400) //最大不要超过1472
		//kcpConn.SetWindowSize(8, 8)
		//kcpConn.SetReadBuffer(8)
		//kcpConn.SetWriteBuffer(8)
		kcpConn.SetACKNoDelay(true)

		return network.NewDefaultConn(conn, server)
	})
	if err != nil {
		panic(err)
	}
	c.server = server
}
func (c *KcpComponent) Start() {

}
func (c *KcpComponent) Tick(now int64) {
	for {
		if c.authCons.Length() == 0 {
			break
		}
		f := c.authCons.First()
		seid := f.Value().(uint64)
		t := g.UUID.GetTime(seid) * 1000
		if now-t < int64(authWaitTime) {
			break
		}
		c.authCons.Delete(seid)

		v, _ := c.cons[seid]
		if v != nil {

		}
	}
}
func (c *KcpComponent) Stop() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for _, v := range c.cons {
		v.Close()
		delete(c.cons, v.GetSessionId())
	}
	c.server.Stop(true)
}
func (c *KcpComponent) Destroy() {

}

func (c *KcpComponent) OnConnect(conn *network.Conn) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if v, ok := c.cons[conn.GetSessionId()]; ok {
		//删除老的
		delete(c.cons, conn.GetSessionId())
		//删除uid映射
		if uid, ok := conn.GetExtraData("uid").(uint64); ok {
			delete(c.connMapping, uid)
		}
		//关闭老的
		v.Close()
	}
	//
	c.cons[conn.GetSessionId()] = conn
	//需要超时断开的
	c.authCons.Insert(conn.GetSessionId())
	return true
}

// OnMessage is called when the connection receives a packet,
// If the return value of false is closed
func (c *KcpComponent) OnMessage(conn *network.Conn, packet network.Packet) bool {
	c.mtx.Lock()
	var conn1 *network.Conn
	ok := false
	if conn1, ok = c.cons[conn.GetSessionId()]; !ok {
		c.mtx.Unlock()
		wlog.Errorf("[KcpComponent.OnMessage] conn:[%d] not found", conn.GetSessionId())
		return false
	}
	conn = conn1
	c.mtx.Unlock()
	pk, ok := packet.(*network.DefaultPacket)
	if !ok {
		wlog.Errorf("[KcpComponent.OnMessage] conn:[%d] packet type error", conn.GetSessionId())
		return false
	}
	req := &message.Request{}
	err := proto.Unmarshal(pk.GetBody(), req)
	if err != nil {
		wlog.Errorf("[KcpComponent.OnMessage] conn:[%d] unmarshal error", conn.GetSessionId())
		return false
	}

	if req.Cmd < uint32(message.CMD_OMin) || req.Cmd > uint32(message.CMD_OMax) {
		wlog.Errorf("[KcpComponent.OnMessage] conn:[%d] range:[%d] error", conn.GetSessionId(), req.Cmd)
		return false
	}

	switch message.CMD(req.Cmd) {
	case message.CMD_Login:
		if !c.doLogin(conn, req) {
			return false
		}
	case message.CMD_Offline:
		if !c.doOffline(conn, req) {
			return false
		}
	default:
		if !c.doOther(conn, req) {
			return false
		}
	}
	return c.sendToHome(conn, req)
}

// OnClose is called when the connection closed
func (c *KcpComponent) OnClose(conn *network.Conn) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	//删除链接
	delete(c.cons, conn.GetSessionId())
	//删除uid映射
	if uid, ok := conn.GetExtraData("uid").(uint64); ok {
		delete(c.connMapping, uid)
	}
}

func (c *KcpComponent) Get(session uint64) *network.Conn {
	conn, ok := c.cons[session]
	if !ok {
		return nil
	}
	return conn
}
