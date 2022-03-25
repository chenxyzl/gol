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
	Node                 bif.IActor
	addr                 string
	enc                  *nats.EncodedConn // NATS的Conn
	dispatcherReplayMap  map[uint32]ifs.RPCFunc
	dispatcherNoReplyMap map[uint32]ifs.RPCFunc
	subscriberTopic      map[string]*nats.Subscription // 订阅路径
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
	c.dispatcherReplayMap = make(map[uint32]ifs.RPCFunc)
	c.dispatcherNoReplyMap = make(map[uint32]ifs.RPCFunc)
	c.subscriberTopic = make(map[string]*nats.Subscription)
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
		wlog.Warn("[NatsComponent.Load] nats.Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.DiscoveredServersHandler(func(nc *nats.Conn) {
		wlog.Info("[NatsComponent.Load] nats.DiscoveredServersHandler", nc.DiscoveredServers())
	}))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		if nil == err {
			wlog.Info("[NatsComponent.Load] nats.DisconnectErrHandler")
		} else {
			wlog.Warn("[NatsComponent.Load] nats.DisconnectErrHandler,error=[%v]", err)
		}
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		wlog.Warn("[NatsComponent.Load] nats.ClosedHandler")
	}))
	opts = append(opts, nats.ErrorHandler(func(nc *nats.Conn, subs *nats.Subscription, err error) {
		wlog.Warn("[NatsComponent.Load] nats.ErrorHandler subs=[%s] error=[%s]", subs.Subject, err.Error())
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
	c.RegisterSubscriber(fmt.Sprintf("%s.%d", g.RoleTyp(), g.Id()))
}
func (c *NatsComponent) Tick(int64 int64) {

}
func (c *NatsComponent) Stop() {
	c.enc.FlushTimeout(3 * time.Second)
	c.enc.Close()
}
func (c *NatsComponent) Destroy() {

}

func (c *NatsComponent) RegisterSubscriber(topic string) error {
	if _, ok := c.subscriberTopic[topic]; ok {
		return fmt.Errorf("[NatsComponent.RegisterSubscriberSelf] repeated topic:[%s]", topic)
	}
	sub, err := c.enc.Subscribe(topic, func(msg *nats.Msg) {
		m := &message2.NatsRequest{}
		err := proto.Unmarshal(msg.Data, m)
		//逻辑上肯定不存在这种情况 ～ 除非被外部攻击-此时不需要处理这种错误
		if err != nil {
			wlog.Error("[NatsComponent.RegisterSubscriberSelf] proto unmarshal error")
		} else {
			//分发到对应的
			c.Node.AddMessage(m)
		}
	})
	if err != nil {
		wlog.Errorf("[NatsComponent.RegisterSubscriberSelf] sub error: %v", err)
		return err
	}
	c.subscriberTopic[topic] = sub
	return nil
}

func (c *NatsComponent) RegisterEvent(cmd uint32, handler ifs.RPCFunc, hasReply bool) {
	if _, ok := c.dispatcherReplayMap[cmd]; ok {
		panic(fmt.Sprintf("[NatsComponent.RegisterEvent] t register duplicate cmd[%d]", cmd))
	}
	if _, ok := c.dispatcherNoReplyMap[cmd]; ok {
		panic(fmt.Sprintf("[NatsComponent.RegisterEvent] t register duplicate cmd[%d]", cmd))
	}
	if hasReply {
		c.dispatcherReplayMap[cmd] = handler
	} else {
		c.dispatcherNoReplyMap[cmd] = handler
	}
}

func (c *NatsComponent) Reply(url string, rep *message2.NatsReply) {
	c.enc.Publish(url, rep)
}

//Dispatch 消息分发
func (c *NatsComponent) Dispatch(req *message2.NatsRequest) {
	//错误兜底--主要在于reply时候
	defer func() {
		err := recover()
		if nil != err {
			trace := string(debug.Stack())
			println("panic:", req.ActorRef, req.Cmd, err, trace)
			wlog.Errorf("[NatsComponent.Dispatch] Dispatch uid=[%v] cmd=[%d] panic(%v) stack:%s", req.ActorRef, req.Cmd, err, trace)
		}
	}()

	start := time.Now()
	wlog.Debugf("[NatsComponent.Dispatch] Dispatch uid=[%v], cmd=[%d]", req.ActorRef, req.Cmd)
	span := time.Since(start)
	if f, ok := c.dispatcherReplayMap[req.Cmd]; ok {
		c.dispatchReplyHandler(f, req)
	} else if f, ok := c.dispatcherNoReplyMap[req.Cmd]; ok {
		c.dispatchNoReplyHandler(f, req)
	} else {
		wlog.Errorf("[NatsComponent.Dispatch] no handler[%d]", req.Cmd)
	}
	if span > slowTime {
		wlog.Warnf("[NatsComponent.Dispatch] Dispatch slow uid[%v] cmd[%d] execute_time[%d(ms)]", req.ActorRef, req.Cmd, span.Milliseconds())
	}
}

func (c *NatsComponent) dispatchReplyHandler(f ifs.RPCFunc, req *message2.NatsRequest) {
	code := message2.Code_OK
	// 是否recover
	defer func() {
		err := recover()
		if nil != err {
			trace := string(debug.Stack())
			println("panic:", req.ActorRef, req.Cmd, err, trace)
			wlog.Errorf("[NatsComponent.dispatchReplyHandler] Dispatch uid=[%v] cmd=[%d] panic(%v) stack:%s", req.ActorRef, req.Cmd, err, trace)
			code = message2.Code_UnknownError
		}
		//如果有错误则返回给目标url--这个url是临时的，不存在重复消费的情况
		if code != message2.Code_OK {
			wlog.Debugf("[NatsComponent.dispatchReplyHandler] uid:[%d] cmd:[%d] reply error:[%d]", code)
			rep := &message2.NatsReply{
				ActorRef:  req.ActorRef,
				Cmd:  req.Cmd,
				Data: nil,
				Code: code,
			}
			c.Reply(req.ReplayUrl, rep)
		}
	}()
	//do
	code = f(req)
	//log
}

func (c *NatsComponent) dispatchNoReplyHandler(f ifs.RPCFunc, req *message2.NatsRequest) {
	code := message2.Code_OK
	// 是否recover
	defer func() {
		err := recover()
		if nil != err {
			trace := string(debug.Stack())
			println("panic:", req.ActorRef, req.Cmd, err, trace)
			wlog.Errorf("[NatsComponent.dispatchNoReplyHandler] Dispatch uid=[%v] cmd=[%d] panic(%v) stack:%s", req.ActorRef, req.Cmd, err, trace)
			code = message2.Code_UnknownError
		}
		if code != message2.Code_OK {
			wlog.Debugf("[NatsComponent.dispatchNoReplyHandler] uid:[%d] cmd:[%d] reply error:[%d]", code)
		}
	}()

	//do
	code = f(req)
	//log
}

func (c *NatsComponent) Ask(req *message2.NatsRequest) (*message2.NatsReply, message2.Code) {
	//todo 更具一致性hash算法 找到对应的nodeId
	//todo AsyncCall调用nats的request/reply来发起请求
	c.Node.SafeAsyncDo(func() {
		//c.enc.Request()
	})
	return nil, message2.Code_OK
}
func (c *NatsComponent) Tell(req *message2.NatsRequest) message2.Code {
	//todo 更具一致性hash算法 找到对应的nodeId
	//todo AsyncCall调用用nats的notify发送消息
	c.Node.SafeAsyncDo(func() {
		//c.enc.Publish()
	})
	return message2.Code_OK
}
