package message

import "fmt"

func (x *NatsRequest) GetUid() uint64 {
	return x.ActorRef.GetUid()
}

func (x *NatsRequest) GetType() ActorType {
	return x.ActorRef.GetType()
}

func (x *NatsRequest) GetTopic() string {
	return fmt.Sprintf("%s.%d", x.GetType().Name(), x.GetUid())
}
