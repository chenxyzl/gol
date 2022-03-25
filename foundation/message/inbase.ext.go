package message

func (x *NatsRequest) GetUid() uint64 {
	return x.ActorRef.GetUid()
}

func (x *NatsRequest) GetType() ActorType {
	return x.ActorRef.GetType()
}
