package message

//系统错误 100~999
const (
	//系统错误 100~999
	Code_UnknownError              Code = 100
	Code_NasRpcNotRegister         Code = 101
	Code_ActorNoFound              Code = 102
	Code_NatsRequestUnmarshalError Code = 103
	Code_NatsReplyUnmarshalError   Code = 104
)

func init() {

	Code_name[int32(Code_NasRpcNotRegister)] = "NasRpcNotRegister "
	Code_value["NasRpcNotRegister "] = int32(Code_NasRpcNotRegister)

	Code_name[int32(Code_ActorNoFound)] = "ActorNoFound"
	Code_value["ActorNoFound"] = int32(Code_ActorNoFound)

	Code_name[int32(Code_NatsRequestUnmarshalError)] = "NatsRequestUnmarshalError"
	Code_value["NatsRequestUnmarshalError"] = int32(Code_NatsRequestUnmarshalError)

	Code_name[int32(Code_NatsReplyUnmarshalError)] = "NatsReplyUnmarshalError"
	Code_value["NatsReplyUnmarshalError"] = int32(Code_NatsReplyUnmarshalError)
}
