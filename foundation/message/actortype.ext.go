package message

const (
	ActorType_PlayerActor ActorType = 1 //占位
)

func init() {
	Code_name[int32(ActorType_PlayerActor)] = "PlayerActor "
	Code_value["PlayerActor "] = int32(ActorType_PlayerActor)
}

func (x ActorType) Name() string {
	return Code_name[int32(x)]
}
