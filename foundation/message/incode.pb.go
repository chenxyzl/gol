package message

//系统错误 100~999
const (
	//系统错误 100~999
	UnknownError              = 100
	NasRpcNotRegister         = 101
	ActorNoFound              = 102
	NatsRequestUnmarshalError = 103
	NatsReplyUnmarshalError   = 104
)

func init() {
	Code_name[ActorNoFound] = "ActorNoFound"
	Code_value["ActorNoFound"] = ActorNoFound

	Code_name[NatsRequestUnmarshalError] = "NatsRequestUnmarshalError"
	Code_value["NatsRequestUnmarshalError"] = NatsRequestUnmarshalError

	Code_name[NatsReplyUnmarshalError] = "NatsReplyUnmarshalError"
	Code_value["NatsReplyUnmarshalError"] = NatsReplyUnmarshalError
}
