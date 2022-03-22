package easyrpcimpl

import (
	"foundation/home/player"
	"foundation/message"
)

//PlayerRPCService 玩家rpc
type PlayerRPCService struct {
}

//Login 发送所有数据
func (service *PlayerRPCService) Login(player *player.PlayerActor, in *message.CS_Login) *message.SC_Login {
	out := &message.SC_Login{}

	return out
}

func (service *PlayerRPCService) GetSenderDelegate(uid uint64) *player.PlayerActor {
	return nil
}
