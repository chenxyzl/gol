package nats_component

import (
	"fmt"
	"foundation/framework/bif"
	"foundation/framework/component"
	"foundation/framework/component/ifs"
	"foundation/framework/g"
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
	Node          bif.IActor
	addr          string
	enc           *nats.EncodedConn // NATS的Conn
	dispatcherMap map[uint32]ifs.RPCFunc
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

}
func (c *NatsComponent) Start() {
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
	c.GoLoop()
}
func (c *NatsComponent) Tick(int64 int64) {

}
func (c *NatsComponent) Stop() {
	c.enc.FlushTimeout(3 * time.Second)
	c.enc.Close()
}
func (c *NatsComponent) Destroy() {

}

func (c *NatsComponent) GoLoop() {

}

func (c *NatsComponent) RegisterEvent(cmd uint32, handler ifs.RPCFunc) {
	if _, ok := c.dispatcherMap[cmd]; ok {
		panic(fmt.Sprintf("[NatsComponent] t register duplicate cmd[%d]", cmd))
	}
	c.dispatcherMap[cmd] = handler
}

//Dispatch 消息分发
func (c *NatsComponent) Dispatch(uid uint64, cmd uint32, b []byte) error {
	start := time.Now()
	wlog.Debug("[NatsComponent] Dispatch uid=[%d], cmd=[%d]", uid, cmd)

	// 是否recover
	defer func() {
		err := recover()
		if nil != err {
			trace := string(debug.Stack())
			fmt.Println("panic:", uid, cmd, err, trace)
			wlog.Error("[NatsComponent] Dispatch uid=[%d] cmd=[%d] panic(%v) stack:%s", uid, cmd, err, trace)
		}
	}()

	f, ok := c.dispatcherMap[cmd]
	if !ok {
		return fmt.Errorf("[NatsComponent] no handler[%d]", cmd)
	}

	err := f(uid, cmd, b)
	span := time.Since(start)
	if span > slowTime {
		wlog.Error("[NatsComponent] Dispatch slow uid[%d] cmd[%d] execute_time[%d(ms)]", uid, cmd, span.Milliseconds())
	}
	return err
}
