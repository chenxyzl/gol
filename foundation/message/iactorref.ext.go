package message

type IActorRef interface {
	GetUid() uint64
	GetType() string
	To() *ActorRef
}
