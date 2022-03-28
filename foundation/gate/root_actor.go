package gate

import (
	"foundation/framework/base"
	"foundation/framework/bif"
	"foundation/framework/g"
	"foundation/framework/network"
	component2 "foundation/gate/global_component"
	ifs2 "foundation/gate/global_component/ifs"
	message2 "foundation/message"
	"github.com/golang/protobuf/proto"
	"gitlab-ee.funplus.io/watcher/watcher/misc/wlog"
	"reflect"
	"time"
)

var _ bif.IActor = &RootActor{}

type RootActor struct {
	base.Actor
}

func NewActor(boxSize int32, maxRunningGoSize int32) *RootActor {
	actor := &RootActor{}
	actor.Constructor(boxSize, maxRunningGoSize)
	actor.RegisterComponent()
	return actor
}

func (actor *RootActor) RegisterComponent() {
}

func (actor *RootActor) OnRecv(message any) {
	switch msg := message.(type) {
	case message2.NatsNotify: //请求的回复
		{
			sess := msg.ActorRef.Uid
			c := g.Root.GetComponent(component2.KcpCom).(ifs2.IKcpComponent)
			conn := c.Get(sess)
			if conn != nil {
				a := &message2.Reply{
					Sn:   msg.Sn,
					Cmd:  msg.Cmd,
					Data: msg.Data,
				}
				data, err := proto.Marshal(a)
				if err == nil {
					wlog.Errorf("[RootActor.OnRecv] proto marshal err:[%v]", err)
				} else {
					conn.AsyncWritePacket(network.NewDefaultPacket(data), time.Second)
				}
			}
		}
	default:
		wlog.Errorf("[RootActor.OnRecv] unknown type:[%v]", reflect.TypeOf(msg).Name())
	}
}
