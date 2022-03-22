package actor

type IActorRef interface {
	GetUid() uint64
	GetActorType() string
}
