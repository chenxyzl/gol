package nats_component

import (
	"fmt"
	"foundation/framework/bif"
	"foundation/framework/component"
	"foundation/framework/component/ifs"
	"foundation/framework/g"
	message2 "foundation/message"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/protobuf"
	"gitlab-ee.funplus.io/watcher/watcher/misc/wlog"
	"runtime/debug"
	"time"
)

var _ ifs.INatsComponent = &NatsComponent{}

//rpc单次执行时间
const (
	slowTime = time.Millisecond * 20
)

type NatsComponent struct {
	Node            bif.IActor
	addr            string
	enc             *nats.EncodedConn // NATS的Conn
	dispatcherMap   map[uint32]ifs.RPCFunc
	subscriberTopic string // 订阅路径
}

func (c *NatsComponent) Constructor(args ...interface{}) {
	if len(args) != 1 {
		panic("[NatsComponent] params must eq 1。\r\n first:nats url")
	}
	addr, ok := args[0].(string)
	if !ok {
		panic("[NatsComponent] params[0] must string")
	}
	c.addr = addr
	c.dispatcherMap = make(map[uint32]ifs.RPCFunc)
}
func (c *NatsComponent) Name() component.ComType {
	return component.NatsCom
}
func (c *NatsComponent) Load() {
	name := fmt.Sprintf("%v-%v", g.RoleTyp(), g.Id())
	// 设置参数
	opts := make([]nats.Option, 0)
	opts = append(opts, nats.Name(name))
	user := g.GlobalConfig.GetString("Nats.User")
	pwd := g.GlobalConfig.GetString("Nats.Pwd")
	if len(user) > 0 {
		opts = append(opts, nats.UserInfo(user, pwd))
	}
	opts = append(opts, nats.ReconnectWait(time.Second*time.Duration(g.GlobalConfig.GetInt("Nats.ReconnectWait"))))
	opts = append(opts, nats.MaxReconnects(g.GlobalConfig.GetInt("Nats.MaxReconnects")))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		wlog.Warn("[NatsComponent] nats.Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.DiscoveredServersHandler(func(nc *nats.Conn) {
		wlog.Info("[NatsComponent] nats.DiscoveredServersHandler", nc.DiscoveredServers())
	}))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		if nil == err {
			wlog.Info("[NatsComponent] nats.DisconnectErrHandler")
		} else {
			wlog.Warn("[NatsComponent] nats.DisconnectErrHandler,error=[%v]", err)
		}
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		wlog.Warn("[NatsComponent] nats.ClosedHandler")
	}))
	opts = append(opts, nats.ErrorHandler(func(nc *nats.Conn, subs *nats.Subscription, err error) {
		wlog.Warn("[NatsComponent] nats.ErrorHandler subs=[%s] error=[%s]", subs.Subject, err.Error())
	}))

	// 创建nats client
	nc, err := nats.Connect(c.addr, opts...)
	if err != nil {
		panic(err)
	}
	enc, err1 := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if nil != err1 {
		panic(err1)
	}
	c.enc = enc
}
func (c *NatsComponent) Start() {
	c.RegisterSubscriberSelf()
}
func (c *NatsComponent) Tick(int64 int64) {

}
func (c *NatsComponent) Stop() {
	c.enc.FlushTimeout(3 * time.Second)
	c.enc.Close()
}
func (c *NatsComponent) Destroy() {

}

func (c *NatsComponent) RegisterSubscriberSelf() {
	c.enc.Subscribe(fmt.Sprintf("%s.%d", g.RoleTyp(), g.Id()), func(msg *nats.Msg) {
		m := &message2.NatsRequest{}
		err := proto.Unmarshal(msg.Data, m)
		if err != nil {
			wlog.Error("proto unmarshal error")
		} else {
			//分发到对应的
			c.Node.AddMessage(m)
		}
	})
}

func (c *NatsComponent) RegisterEvent(cmd uint32, handler ifs.RPCFunc) {
	if _, ok := c.dispatcherMap[cmd]; ok {
		panic(fmt.Sprintf("[NatsComponent] t register duplicate cmd[%d]", cmd))
	}
	c.dispatcherMap[cmd] = handler
}

func (c *NatsComponent) Reply(url string, rep *message2.NatsReply) {
	c.enc.Publish(url, rep)
}

//Dispatch 消息分发
func (c *NatsComponent) Dispatch(req *message2.NatsRequest) {
	start := time.Now()
	wlog.Debugf("[NatsComponent] Dispatch uid=[%d], cmd=[%d]", req.Uid, req.Cmd)

	// 是否recover
	defer func() {
		err := recover()
		if nil != err {
			trace := string(debug.Stack())
			fmt.Println("panic:", req.Uid, req.Cmd, err, trace)
			wlog.Errorf("[NatsComponent] Dispatch uid=[%d] cmd=[%d] panic(%v) stack:%s", req.Uid, req.Cmd, err, trace)
		}
	}()

	f, ok := c.dispatcherMap[req.Cmd]
	if !ok {
		wlog.Errorf("[NatsComponent] no handler[%d]", req.Cmd)
	}

	err := f(req)
	span := time.Since(start)
	if span > slowTime {
		wlog.Errorf("[NatsComponent] Dispatch slow uid[%d] cmd[%d] execute_time[%d(ms)]", req.Uid, req.Cmd, span.Milliseconds())
	}
	if err != nil {
		wlog.Error(err)
	}
}
