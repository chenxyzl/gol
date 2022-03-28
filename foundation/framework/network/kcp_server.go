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

//StartKcpServer 启动kcp服务器 通常使用默认配置就可以了(也就是dupConfig=nil)
func StartKcpServer(addr string, callback ConnCallback, protocol Protocol, defaultConfig *Config, create ConnectionCreator) (*Server, error) {
	//默认配置
	if defaultConfig == nil {
		defaultConfig = &Config{
			PacketReceiveChanLimit: 1024,
			PacketSendChanLimit:    1024,
			ConnReadTimeout:        time.Second * 180,
			ConnWriteTimeout:       time.Second * 180,
		}
	}

	l, err := kcp.Listen(addr)
	if nil != err {
		return nil, err
	}

	server := NewServer(defaultConfig, callback, protocol)

	if create == nil {
		create = func(conn net.Conn, i *Server) IDoConn {

			/*
				setKCPConfig(int sndwnd, int rcvwnd, int nodelay, int interval, int resend, int nc, int minrto, int mtu)
				sndwnd  最大发送窗口  默认值 32
				rcvwnd  最大接收窗口  默认值 32

				nodelay 是否启用 nodelay模式，0不启用；1启用。
				interval 协议内部工作的 interval，单位毫秒，比如 10ms或者 20ms
				resend  快速重传模式，默认0关闭，可以设置2（2次ACK跨越将会直接重传）
				nc  是否关闭流控，默认是0代表不关闭，1代表关闭。

				普通模式：ikcp_nodelay(kcp, 0, 40, 0, 0); 极速模式： ikcp_nodelay(kcp, 1, 10, 2, 1);

				minrto
				最小RTO
				不管是 TCP还是 KCP计算 RTO时都有最小 RTO的限制，即便计算出来RTO为4`0ms，由于默认的 RTO是100ms，协议只有在100ms后才能检测到丢包，快速模式下该值为30ms，可以手动更改该值：

				mtu  默认值 1400
				最大传输单元
				纯算法协议并不负责探测 MTU，默认 mtu是1400字节，可以使用ikcp_setmtu来设置该值。该值将会影响数据包归并及分片时候的最大传输单元。
			*/

			//普通模式
			//setKCPConfig(32, 32, 0, 40, 0, 0, 100, 1400)

			//极速模式
			//setKCPConfig(32, 32, 1, 10, 2, 1, 30, 1400)

			//普通模式：ikcp_nodelay(kcp, 0, 40, 0, 0); 极速模式： ikcp_nodelay(kcp, 1, 10, 2, 1);

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

			return NewConn(conn, server)
		}
	}

	go server.Start(l, create)

	return server, nil
}
