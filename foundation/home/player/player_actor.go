package player

import (
	"foundation/framework/actor"
	"foundation/framework/message"
)

var _ actor.IActor = &PlayerActor{}

type PlayerActor struct {
	actor.Actor
}

func (a *PlayerActor) Constructor(boxSize int32, maxRunningGoSize int32) {
	a.Actor.Constructor(boxSize, maxRunningGoSize)
	a.Init()
}

func (a *PlayerActor) Init() {

}

func (a *PlayerActor) OnRecv(message message.IMessage) {

}
