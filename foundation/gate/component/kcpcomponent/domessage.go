package kcpcomponent

import (
	"foundation/framework/component"
	"foundation/framework/component/ifs"
	"foundation/framework/g"
	"foundation/framework/network"
	"foundation/message"
	"github.com/golang/protobuf/proto"
	"gitlab-ee.funplus.io/watcher/watcher/misc/wlog"
	"time"
)

func (c *KcpComponent) doLogin(conn *network.Conn, packet *message.Request) bool {
	if packet.Cmd != uint32(message.CMD_Login) {
		wlog.Error("[KcpComponent.doLogin] first must login, cmd:[%v]", packet.Cmd)
		return false
	}

	login := &message.CS_Login{}
	err := proto.Unmarshal(packet.Data, login)
	if err != nil {
		wlog.Error("[KcpComponent.doLogin] unmarshal err:[%v]", err)
		return false
	}
	conn.PutExtraData("uid", login.Uid)
	return true
}

func (c *KcpComponent) checkConnOk(conn *network.Conn) bool {
	uid := conn.GetExtraData("uid").(uint64)
	_, ok := c.connMapping[uid]
	if !ok {
		return false
	}
	return true
}

func (c *KcpComponent) doOffline(conn *network.Conn, packet *message.Request) bool {
	if !c.checkConnOk(conn) {
		return false
	}
	return true
}

func (c *KcpComponent) doOther(conn *network.Conn, packet *message.Request) bool {
	if !c.checkConnOk(conn) {
		return false
	}
	return true
}

func (c *KcpComponent) sendToHome(conn *network.Conn, packet *message.Request) bool {
	nats := g.Root.GetComponent(component.NatsCom).(ifs.INatsComponent)
	uid := conn.GetExtraData("uid").(uint64)
	if 0 == uid {
		return false
	}
	v := nats.Ask(&message.NatsRequest{
		ActorRef: &message.ActorRef{
			Uid:  uid,
			Type: message.ActorType_PlayerActor,
		},
		Sn:   packet.Sn,
		Cmd:  packet.Cmd,
		Data: packet.Data,
	})
	if v == nil {
		return false
	}

	a := &message.Reply{
		Sn:   v.Sn,
		Cmd:  v.Cmd,
		Data: v.Data,
		Code: v.Code,
	}
	data, err := proto.Marshal(a)
	if err == nil {
		wlog.Errorf("[KcpComponent.sendToHome] proto marshal err:[%v]", err)
		return false
	} else {
		conn.AsyncWritePacket(network.NewDefaultPacket(data), time.Second)
	}
	return true
}
