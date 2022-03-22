package bif

type IActorRef interface {
	GetUid() uint64
	GetActorType() string
}
