package network

import (
	"errors"
	"foundation/framework/g"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// IConn is an interface of methods that are used as callbacks on a connection
type IConn interface {
	//Do 启动函数
	Do()

	// OnConnect is called when the connection was accepted,
	// If the return value of false is closed
	OnConnect() bool

	// OnMessage is called when the connection receives a packet,
	// If the return value of false is closed
	OnMessage(Packet) bool

	// OnClose is called when the connection closed
	OnClose()
}

// Error type
var (
	ErrConnClosing   = errors.New("use of closed network connection")
	ErrWriteBlocking = errors.New("write packet was blocking")
	ErrReadBlocking  = errors.New("read packet was blocking")
)

// Conn exposes a set of callbacks for the various events that occur on a connection
type Conn struct {
	IConn
	srv               *Server
	conn              net.Conn               // the raw connection
	extraData         map[string]interface{} // to save extra data
	closeOnce         sync.Once              // close the conn, once, per instance
	closeFlag         int32                  // close flag
	closeChan         chan struct{}          // close chanel
	packetSendChan    chan Packet            // packet send chanel
	packetReceiveChan chan Packet            // packeet receive chanel
	sessionId         uint64
}

// newConn returns a wrapper of raw conn
func (c *Conn) Constructor(conn net.Conn, srv *Server) {
	c.srv = srv
	c.conn = conn
	c.closeChan = make(chan struct{})
	c.packetSendChan = make(chan Packet, srv.Config.PacketSendChanLimit)
	c.packetReceiveChan = make(chan Packet, srv.Config.PacketReceiveChanLimit)
	c.sessionId = g.UUID.Generate()
}

// GetExtraData gets the extra data from the Conn
func (c *Conn) GetExtraData(str string) interface{} {
	return c.extraData[str]
}

// PutExtraData puts the extra data with the Conn
func (c *Conn) PutExtraData(str string, data interface{}) {
	c.extraData[str] = data
}

// GetRawConn returns the raw net.TCPConn from the Conn
func (c *Conn) GetRawConn() net.Conn {
	return c.conn
}

func (c *Conn) GetSessionId() uint64 {
	return c.sessionId
}

func (c *Conn) GetRemoteIp() string {
	return c.conn.RemoteAddr().String()
}

// Close closes the connection
func (c *Conn) Close() {
	c.closeOnce.Do(func() {
		atomic.StoreInt32(&c.closeFlag, 1)
		close(c.closeChan)
		close(c.packetSendChan)
		close(c.packetReceiveChan)
		c.conn.Close()
		c.OnClose()
	})
}

// IsClosed indicates whether or not the connection is closed
func (c *Conn) IsClosed() bool {
	return atomic.LoadInt32(&c.closeFlag) == 1
}

// AsyncWritePacket async writes a packet, this method will never block
func (c *Conn) AsyncWritePacket(p Packet, timeout time.Duration) (err error) {
	if c.IsClosed() {
		return ErrConnClosing
	}

	defer func() {
		if e := recover(); e != nil {
			err = ErrConnClosing
		}
	}()

	if timeout == 0 {
		select {
		case c.packetSendChan <- p:
			return nil

		default:
			return ErrWriteBlocking
		}

	} else {
		select {
		case c.packetSendChan <- p:
			return nil

		case <-c.closeChan:
			return ErrConnClosing

		case <-time.After(timeout):
			return ErrWriteBlocking
		}
	}
}

// Do it
func (c *Conn) Do() {
	if !c.OnConnect() {
		return
	}
	asyncDo(c.handleLoop, c.srv.WaitGroup)
	asyncDo(c.readLoop, c.srv.WaitGroup)
	asyncDo(c.writeLoop, c.srv.WaitGroup)
}

func (c *Conn) readLoop() {
	defer func() {
		recover()
		c.Close()
	}()

	for {
		select {
		case <-c.srv.exitChan:
			return

		case <-c.closeChan:
			return

		default:
		}

		c.conn.SetReadDeadline(time.Now().Add(c.srv.Config.ConnReadTimeout))
		p, err := c.srv.protocol.ReadPacket(c.conn)
		if err != nil {
			return
		}

		c.packetReceiveChan <- p
	}
}

func (c *Conn) writeLoop() {
	defer func() {
		recover()
		c.Close()
	}()

	for {
		select {
		case <-c.srv.exitChan:
			return

		case <-c.closeChan:
			return

		case p := <-c.packetSendChan:
			if c.IsClosed() {
				return
			}
			c.conn.SetWriteDeadline(time.Now().Add(c.srv.Config.ConnWriteTimeout))
			if _, err := c.conn.Write(p.Serialize()); err != nil {
				return
			}
		}
	}
}

func (c *Conn) handleLoop() {
	defer func() {
		recover()
		c.Close()
	}()

	for {
		select {
		case <-c.srv.exitChan:
			return

		case <-c.closeChan:
			return

		case p := <-c.packetReceiveChan:
			if c.IsClosed() {
				return
			}
			if !c.OnMessage(p) {
				c.Close()
				return
			}
		}
	}
}

func asyncDo(fn func(), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		fn()
		wg.Done()
	}()
}
