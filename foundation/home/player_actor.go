package home

import (
	"foundation/framework/actor"
	"foundation/framework/message"
)

var _ actor.IActor = &PlayerActor{}

type PlayerActor struct {
	actor.Actor
}

func (a *PlayerActor) OnRecv(message message.IMessage) {

}
