package network

import (
	"net"
	"sync"
	"time"
)

type Config struct {
	PacketSendChanLimit    uint32        // the limit of packet send channel
	PacketReceiveChanLimit uint32        // the limit of packet receive channel
	ConnReadTimeout        time.Duration // read timeout
	ConnWriteTimeout       time.Duration // write timeout
}

type Server struct {
	Config    *Config         // server configuration
	protocol  Protocol        // customize packet protocol
	exitChan  chan struct{}   // notify all goroutines to shutdown
	WaitGroup *sync.WaitGroup // wait for all goroutines
	closeOnce sync.Once
	listener  net.Listener
}

// NewServer creates a server
func NewServer(config *Config, protocol Protocol) *Server {
	return &Server{
		Config:    config,
		protocol:  protocol,
		exitChan:  make(chan struct{}),
		WaitGroup: &sync.WaitGroup{},
	}
}

type ConnectionCreator func(net.Conn, *Server) IConn

// Start starts service
func (s *Server) Start(listener net.Listener, create ConnectionCreator) {
	s.listener = listener
	s.WaitGroup.Add(1)
	defer func() {
		s.WaitGroup.Done()
	}()

	for {
		select {
		case <-s.exitChan:
			return

		default:
		}

		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		s.WaitGroup.Add(1)
		go func() {
			create(conn, s).Do()
			s.WaitGroup.Done()
		}()
	}
}

// Stop stops service
func (s *Server) Stop(wait bool) {
	s.closeOnce.Do(func() {
		close(s.exitChan)
		s.listener.Close()
	})
	if wait {
		s.WaitGroup.Wait()
	}
}
