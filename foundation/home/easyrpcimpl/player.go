package easyrpcimpl

import (
	"foundation/home"
	"foundation/message"
)

//PlayerRPCService 玩家rpc
type PlayerRPCService struct {
}

//RequestUserData 发送所有数据
func (service *PlayerRPCService) Login(sender *home.PlayerActor, in *message.CS_Login) *message.SC_Login {
	out := &message.SC_Login{}

	return out
}

func (service *PlayerRPCService) GetSenderDelegate(uid uint64) *home.PlayerActor {
	return nil
}
