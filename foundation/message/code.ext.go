package message

//系统错误 100~999
const (
	//系统错误 100~999
	Code_UnknownError              Code = 100
	Code_NasRpcNotRegister         Code = 101
	Code_ActorNoFound              Code = 102
	Code_NatsRpcError              Code = 103
	Code_NatsRequestUnmarshalError Code = 104
	Code_NatsReplyUnmarshalError   Code = 105
	Code_MarshalError              Code = 106
	Code_UnmarshalError            Code = 107
)

func init() {

	Code_name[int32(Code_NasRpcNotRegister)] = "NasRpcNotRegister "
	Code_value["NasRpcNotRegister "] = int32(Code_NasRpcNotRegister)

	Code_name[int32(Code_ActorNoFound)] = "ActorNoFound"
	Code_value["ActorNoFound"] = int32(Code_ActorNoFound)


	Code_name[int32(Code_NatsRpcError)] = "NatsRpcError"
	Code_value["NatsRpcError"] = int32(Code_NatsRpcError)

	Code_name[int32(Code_NatsRequestUnmarshalError)] = "NatsRequestUnmarshalError"
	Code_value["NatsRequestUnmarshalError"] = int32(Code_NatsRequestUnmarshalError)

	Code_name[int32(Code_NatsReplyUnmarshalError)] = "NatsReplyUnmarshalError"
	Code_value["NatsReplyUnmarshalError"] = int32(Code_NatsReplyUnmarshalError)

	Code_name[int32(Code_MarshalError)] = "MarshalError"
	Code_value["MarshalError"] = int32(Code_MarshalError)

	Code_name[int32(Code_UnmarshalError)] = "UnmarshalError"
	Code_value["UnmarshalError"] = int32(Code_UnmarshalError)
}
