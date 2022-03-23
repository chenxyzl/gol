package playerrpcimpl

import (
	"foundation/home/player"
	"foundation/home/playerrpc"
)

func init() {
	playerrpc.GetPlayer = func(uid uint64) *player.PlayerActor {
		//todo 存在则直接返回玩家
		//不存在则按照 foundation/readme.md来创建玩家 或者返回空
		return nil
	}
}
