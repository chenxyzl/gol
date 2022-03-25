package message

//外部rpc范围 100~999
const (
	CMD_LoginWorld CMD = 10000
)

func init() {
	CMD_name[int32(CMD_LoginWorld)] = "LoginWorld"
	CMD_value["LoginWorld"] = int32(CMD_LoginWorld)
}
