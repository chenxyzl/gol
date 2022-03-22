package function

//rpc方向
type Direction string

const (
	Direction_Client2GameServer Direction = "Client2GameServer"
	Direction_GameServer2Client Direction = "GameServer2Client"
	Direction_Server2Server     Direction = "Server2Server"
	Direction_Client2Server     Direction = "Client2Server"
	Direction_Server2Client     Direction = "Server2Client"
	Direction_Unknow            Direction = ""
)

type Argument struct {
	Name string
	kind string
}

type Function struct {
	Comment          string
	Async            bool
	AsyncOutFuncName string
	RetType          string
	FuncName         string
	Direction        Direction
	Agrs             string
	Kind             string
	Sig              string
}

func (f Function) GetOutFuncName() string {
	if f.Async {
		return f.AsyncOutFuncName
	}
	return f.FuncName
}

func (f Function) GetRealRetType() string {
	if f.Direction == Direction_GameServer2Client {
		return f.Agrs
	}
	return f.RetType
}

func (f Function) GetIsNotifyStr() string {
	if f.Direction == Direction_GameServer2Client {
		return "true"
	}
	return "false"
}

func NewFunction() *Function {
	return &Function{
		Comment:          "",
		RetType:          "",
		Async:            false,
		AsyncOutFuncName: "",
		FuncName:         "",
		Direction:        Direction_Unknow,
		Agrs:             "",
		Kind:             "",
		Sig:              "",
	}
}
