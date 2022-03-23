package playerrpc

import "foundation/home/player"

type GetPlayerDelegate func(uid uint64) *player.PlayerActor

var GetPlayer GetPlayerDelegate
