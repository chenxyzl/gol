package player

import (
	"foundation/framework/base"
	"foundation/framework/bif"
	"foundation/framework/message"
)

var _ bif.IActor = &PlayerActor{}

type PlayerActor struct {
	base.Actor
}

func (a *PlayerActor) Constructor(boxSize int32, maxRunningGoSize int32) {
	a.Actor.Constructor(boxSize, maxRunningGoSize)
	a.Init()
}

func (a *PlayerActor) Init() {

}

func (a *PlayerActor) OnRecv(message message.IMessage) {

}
