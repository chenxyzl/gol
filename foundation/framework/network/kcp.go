package network

import (
	"foundation/framework/g"
	"foundation/framework/util"
	"github.com/xtaci/kcp-go"
	"net"
	"time"
)

func init() {
	g.UUID, _ = util.NewUUID(1)
}

var DefaultConfig = &Config{
	PacketReceiveChanLimit: 1024,
	PacketSendChanLimit:    1024,
	ConnReadTimeout:        time.Second * 180,
	ConnWriteTimeout:       time.Second * 180,
}

//StartKcpServer 启动kcp服务器 通常使用默认配置就可以了(也就是dupConfig=nil)
func StartKcpServer(addr string, protocol Protocol, defaultConfig *Config, create ConnectionCreator) (*Server, error) {
	//默认配置
	if defaultConfig == nil {
		defaultConfig = DefaultConfig
	}

	l, err := kcp.Listen(addr)
	if nil != err {
		return nil, err
	}

	server := NewServer(defaultConfig, protocol)

	if create == nil {
		create = func(conn net.Conn, i *Server) IConn {
			//普通模式 setKCPConfig(32, 32, 0, 40, 0, 0, 100, 1400)
			//极速模式 setKCPConfig(32, 32, 1, 10, 2, 1, 30, 1400)
			kcpConn := conn.(*kcp.UDPSession)
			kcpConn.SetNoDelay(1, 10, 2, 1)
			kcpConn.SetStreamMode(true)
			kcpConn.SetWindowSize(32, 32)
			kcpConn.SetReadBuffer(1024)
			kcpConn.SetWriteBuffer(1024)
			kcpConn.SetMtu(1400) //最大不要超过1472
			kcpConn.SetACKNoDelay(true)

			return NewDefaultConn(conn, server)
		}
	}

	go server.Start(l, create)

	return server, nil
}
