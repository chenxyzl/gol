package network

import (
	"golang.org/x/net/websocket"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/xtaci/kcp-go"
)

type testCallback struct {
	numConn   uint32
	numMsg    uint32
	numDiscon uint32
}

func (t *testCallback) OnMessage(conn *Conn, msg Packet) bool {

	atomic.AddUint32(&t.numMsg, 1)
	a, _ := msg.(*DefaultPacket)
	r := append([]byte("pong:"), a.GetBody()...)
	//fmt.Println("OnMessage", conn.GetExtraData(), string(msg.(*DefaultPacket).GetBody()))
	conn.AsyncWritePacket(NewDefaultPacket(r), time.Second*1)
	return true
}

func (t *testCallback) OnConnect(conn *Conn) bool {
	atomic.AddUint32(&t.numConn, 1)
	//fmt.Println("OnConnect", conn.GetExtraData())
	return true
}

func (t *testCallback) OnClose(conn *Conn) {
	atomic.AddUint32(&t.numDiscon, 1)

	//fmt.Println("OnDisconnect", conn.GetExtraData())
}

func Test_KCPServer(t *testing.T) {
	const latency = time.Millisecond * 50 * 10000
	callback := &testCallback{}
	config := &Config{
		PacketReceiveChanLimit: 1024,
		PacketSendChanLimit:    1024,
		ConnReadTimeout:        latency,
		ConnWriteTimeout:       latency,
	}
	server, err := StartKcpServer(":10086", callback, &DefaultProtocol{}, config, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Stop(true)

	time.Sleep(time.Second)

	wg := sync.WaitGroup{}
	const max_con = 1
	for i := 0; i < max_con; i++ {
		wg.Add(1)
		time.Sleep(time.Nanosecond)
		go func() {
			defer wg.Done()

			c, e := kcp.Dial("127.0.0.1:10086")
			if nil != e {
				t.FailNow()
			}
			defer c.Close()

			c.Write(NewDefaultPacket([]byte("abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz" +
				"abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz,abcdefghijklmnopqrstuvwxyz")).Serialize())
			b := make([]byte, 10240)
			c.SetReadDeadline(time.Now().Add(latency))
			count, e := c.Read(b)
			if nil != e {
				t.Fatalf("ret:%d,error:%s", count, e.Error())
			}
			v := string(b[4:count])
			println(v)
			//time.Sleep(time.Second)
		}()
	}

	wg.Wait()
	time.Sleep(time.Second * 2)

	n := atomic.LoadUint32(&callback.numConn)
	if n != max_con {
		t.Errorf("numConn[%d] should be [%d]", n, max_con)
	}

	n = atomic.LoadUint32(&callback.numMsg)
	if n != max_con {
		t.Errorf("numMsg[%d] should be [%d]", n, max_con)
	}

	n = atomic.LoadUint32(&callback.numDiscon)
	if n != max_con {
		t.Errorf("numDiscon[%d] should be [%d]", n, max_con)
	}
}

func Benchmark_KCPServer(b *testing.B) {

	config := &Config{
		PacketReceiveChanLimit: 1024,
		PacketSendChanLimit:    1024,
	}

	createConn := func(conn net.Conn, i *Server) *Conn {
		kcpConn := conn.(*kcp.UDPSession)
		kcpConn.SetNoDelay(1, 10, 2, 1)
		kcpConn.SetStreamMode(true)
		kcpConn.SetWindowSize(4096, 4096)
		kcpConn.SetReadBuffer(4 * 1024 * 1024)
		kcpConn.SetWriteBuffer(4 * 1024 * 1024)
		kcpConn.SetACKNoDelay(true)
		return NewConn(conn, i)
	}

	callback := &testCallback{}
	server, err := StartKcpServer(":10086", callback, &DefaultProtocol{}, config, createConn)
	if err != nil {
		b.Fatal(err)
	}

	time.Sleep(time.Millisecond * 100)

	wg := sync.WaitGroup{}
	var max_con uint32 = 0
	c, e := kcp.Dial("127.0.0.1:10086")
	if nil != e {
		b.FailNow()
	}

	go func() {
		for {
			buf := make([]byte, 1024)
			c.SetReadDeadline(time.Now().Add(time.Second * 2))
			_, er := c.Read(buf)
			if nil != er {
				//b.FailNow()
				return
			}
			wg.Done()
		}

	}()

	for i := 0; i < b.N; i++ {
		max_con++

		wg.Add(1)
		go func() {

			c.Write(NewDefaultPacket([]byte("ping")).Serialize())

			//time.Sleep(time.Second)
		}()
	}

	wg.Wait()
	//time.Sleep(time.Second * 2)
	server.Stop(true)

	n := atomic.LoadUint32(&callback.numMsg)
	b.Logf("numMsg[%d]", n)
	if n != callback.numMsg {
		b.Errorf("numMsg[%d] should be [%d]", n, max_con)
	}
	/*
		n = atomic.LoadUint32(&numConn)
		b.Logf("numConn[%d]", n)
		if n != max_con {
			b.Errorf("numConn[%d] should be [%d]", n, max_con)
		}



		n = atomic.LoadUint32(&numDiscon)
		b.Logf("numDiscon[%d]", n)
		if n != numDiscon {
			b.Errorf("numDiscon[%d] should be [%d]", n, max_con)
		}
	*/
}

func Test_TCPServer(t *testing.T) {

	l, err := net.Listen("tcp", ":10086")
	if nil != err {
		panic(err)
	}

	config := &Config{
		PacketReceiveChanLimit: 1024,
		PacketSendChanLimit:    1024,
		ConnReadTimeout:        time.Millisecond * 50,
		ConnWriteTimeout:       time.Millisecond * 50,
	}

	callback := &testCallback{}
	server := NewServer(config, callback, &DefaultProtocol{})

	go server.Start(l, func(conn net.Conn, i *Server) *Conn {
		return NewConn(conn, server)
	})

	time.Sleep(time.Second)

	wg := sync.WaitGroup{}
	const max_con = 2000
	for i := 0; i < max_con; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c, e := net.Dial("tcp", "127.0.0.1:10086")
			if nil != e {
				t.FailNow()
			}
			defer c.Close()
			c.Write(NewDefaultPacket([]byte("ping")).Serialize())
			b := make([]byte, 1024)
			c.SetReadDeadline(time.Now().Add(time.Second * 2))
			c.Read(b)
			//time.Sleep(time.Second)
		}()
	}

	wg.Wait()
	//time.Sleep(max_sleep)
	server.Stop(true)

	n := atomic.LoadUint32(&callback.numConn)
	if n != max_con {
		t.Errorf("numConn[%d] should be [%d]", n, max_con)
	}

	n = atomic.LoadUint32(&callback.numMsg)
	if n != max_con {
		t.Errorf("numMsg[%d] should be [%d]", n, max_con)
	}

	n = atomic.LoadUint32(&callback.numDiscon)
	if n != max_con {
		t.Errorf("numDiscon[%d] should be [%d]", n, max_con)
	}
}

func Test_WsServer(t *testing.T) {

	config := &Config{
		PacketReceiveChanLimit: 1024,
		PacketSendChanLimit:    1024,
		ConnReadTimeout:        time.Millisecond * 50,
		ConnWriteTimeout:       time.Millisecond * 50,
	}

	callback := &testCallback{}
	server := NewServer(config, callback, &DefaultProtocol{})

	mux := http.NewServeMux()
	mux.Handle("/ws", websocket.Handler(func(conn *websocket.Conn) {
		conn.PayloadType = websocket.BinaryFrame
		NewConn(conn, server)
	}))
	httpServer := &http.Server{Addr: ":10086", Handler: mux}
	go func() {
		time.Sleep(time.Second * 1)
		httpServer.Close()
	}()
	httpServer.ListenAndServe()
}
