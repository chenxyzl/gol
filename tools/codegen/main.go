package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"CodeGenerator/function"
	"CodeGenerator/util"

	"github.com/iancoleman/strcase"
	"github.com/urfave/cli"
)

func main() {
	//parse := function.Parse("/Users/andy/GoLang/src/youzu/mhserver/static/protos/rpc/chat.h")
	////function.GenerateServerInterfaceImpl("", parse)

	//function.GeneratePlayerClientProxy("clientProxyOutputFile1",  parse)
	//生成机器人
	//function.GenerateRobotInterface("dungeon", parse)
	//function.GenerateRobotInterfaceImpl("",parse)
	//function.GenerateRobotSessionSender("",parse)
	//function.GenerateRobotPacketRpcDef("",parse)
	//
	//生成WebClient
	//function.GenerateClientSessionPack("", parse)

	//function.GenerateProtoLua("/Users/andy/GoLang/src/youzu/lua/proto","/Users/andy/GoLang/src/youzu/lua/lua")

	CLI()

}

func CLI() {
	app := cli.NewApp()
	rpcCommand := cli.Command{
		Name:      "rpc",
		Usage:     `geek new postgres  go`,
		ArgsUsage: "generate a gin api application 第一个参数:postgres第二个参数:项目名 第三个参数: 数据库连接 第四个参数 数据库名",

		Action: RpcCreate,
	}
	robotCommand := cli.Command{
		Name:      "robot",
		Usage:     `geek new postgres  go-start "root:123456@tcp(127.0.0.1:3306)" dbname`,
		ArgsUsage: "generate a gin api application 第一个参数:postgres第二个参数:项目名 第三个参数: 数据库连接 第四个参数 数据库名",

		Action: RobotCreate,
	}
	webCommand := cli.Command{
		Name:      "web_client",
		Usage:     `geek new postgres  go-start "root:123456@tcp(127.0.0.1:3306)" dbname`,
		ArgsUsage: "generate a gin api application 第一个参数:postgres第二个参数:项目名 第三个参数: 数据库连接 第四个参数 数据库名",

		Action: WebCreate,
	}

	luaCommand := cli.Command{
		Name:      "lua",
		Usage:     `geek new postgres  go-start "root:123456@tcp(127.0.0.1:3306)" dbname`,
		ArgsUsage: "generate a gin api application 第一个参数:postgres第二个参数:项目名 第三个参数: 数据库连接 第四个参数 数据库名",

		Action: LuaCreate,
	}
	app.Commands = []cli.Command{
		rpcCommand,
		webCommand,
		robotCommand,
		luaCommand,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func RpcCreate(c *cli.Context) error {

	protoDir := c.Args()[0]
	_, err := os.Stat(protoDir)
	if err != nil {
		return errors.New("path not exist")
	}
	fileNames, err := util.DirNames(protoDir)
	sort.Strings(fileNames)

	if err != nil {
		return err
	}
	funcMap := make(map[string][]*function.Function)
	serverInterfaceBuffer := &bytes.Buffer{}
	serverInterfaceImplBuffer := &bytes.Buffer{}
	serverUserSenderBuffer := &bytes.Buffer{}

	serverInterfacePath := filepath.Join(c.Args()[2], "rpc_interface.go")
	serverDispatchPath := filepath.Join(c.Args()[3], "rpc_register.go")
	userSenderOutputPath := c.Args()[4]
	playerOutputPath := c.Args()[5]
	function.GenerateServerInterfaceHeader(serverInterfaceBuffer)
	function.GenerateServerInterfaceImplHeader(serverInterfaceImplBuffer)
	function.GenerateUserSenderHeader(serverUserSenderBuffer)

	funcs := make([]*function.Function, 0)
	for _, name := range fileNames {
		if !strings.HasSuffix(name, ".h") {
			continue
		}
		parse := function.Parse(filepath.Join(protoDir, name))
		t := strcase.ToCamel(strings.TrimRight(name, ".h"))
		funcMap[t] = parse
		funcs = append(funcs, parse...)
		function.GenerateServerInterface(serverInterfaceBuffer, strcase.ToCamel(strings.TrimRight(name, ".h")), parse)
		function.GenerateServerInterfaceImpl(serverInterfaceImplBuffer, t, parse)
		function.GenerateUserSender(serverUserSenderBuffer, parse)

	}

	checkErr := checkFuncs(funcs)
	if checkErr != nil {
		return checkErr
	}

	ioutil.WriteFile(serverInterfacePath, serverInterfaceBuffer.Bytes(), os.ModeExclusive|os.FileMode(0664))
	ioutil.WriteFile(serverDispatchPath, serverInterfaceImplBuffer.Bytes(), os.ModeAppend|os.FileMode(0664))
	ioutil.WriteFile(userSenderOutputPath, serverUserSenderBuffer.Bytes(), os.ModeAppend|os.FileMode(0664))
	function.GeneratePlayerClientProxy(playerOutputPath, funcs)

	return nil
}

func RobotCreate(c *cli.Context) error {

	protoDir := c.Args()[0]
	robotInterfacePath := c.Args()[2]
	robotDispatchPath := c.Args()[3]
	robotSenderOutputPath := c.Args()[4]
	robotRpcDefPath := c.Args()[5]
	_, err := os.Stat(protoDir)
	if err != nil {
		return errors.New("path not exist")
	}
	fileNames, err := util.DirNames(protoDir)
	sort.Strings(fileNames)

	if err != nil {
		return err
	}
	funcs := make([]*function.Function, 0)
	for _, name := range fileNames {
		if !strings.HasSuffix(name, ".h") {
			continue
		}
		parse := function.Parse(filepath.Join(protoDir, name))

		funcs = append(funcs, parse...)

	}

	checkErr := checkFuncs(funcs)
	if checkErr != nil {
		return checkErr
	}

	function.GenerateRobotInterface(robotInterfacePath, funcs)
	function.GenerateRobotInterfaceImpl(robotDispatchPath, funcs)
	function.GenerateRobotSessionSender(robotSenderOutputPath, funcs)
	function.GenerateRobotPacketRpcDef(robotRpcDefPath, funcs)

	return nil
}

func WebCreate(c *cli.Context) error {

	protoDir := c.Args()[0]
	clientSessionPackPath := c.Args()[2]
	_, err := os.Stat(protoDir)
	if err != nil {

		return errors.New("path not exist")
	}
	fileNames, err := util.DirNames(protoDir)
	sort.Strings(fileNames)
	if err != nil {
		return err
	}
	funcs := make([]*function.Function, 0)
	for _, name := range fileNames {
		if !strings.HasSuffix(name, ".h") {
			continue
		}
		parse := function.Parse(filepath.Join(protoDir, name))

		funcs = append(funcs, parse...)

	}

	checkErr := checkFuncs(funcs)
	if checkErr != nil {
		return checkErr
	}

	function.GenerateClientSessionPack(clientSessionPackPath, funcs)

	return nil
}

func LuaCreate(c *cli.Context) error {

	hFileDir := c.Args()[0]
	protoDir := c.Args()[1]
	outLuaDir := c.Args()[2]
	netWorkDataRulesFile := c.Args()[3]

	_, err := os.Stat(hFileDir)
	if err != nil {

		return errors.New(err.Error())
	}
	_, err = os.Stat(protoDir)
	if err != nil {

		return errors.New(err.Error())
	}

	function.GenerateProtoLua(protoDir, outLuaDir)
	function.GenerateLuaTable(hFileDir, filepath.Join(outLuaDir, netWorkDataRulesFile))

	return nil
}

func checkFuncs(funcs []*function.Function) error {
	funcSigs := make(map[string]string, len(funcs))
	funcNames := make(map[string]string, len(funcs))
	for _, v := range funcs {
		if name, ok := funcSigs[v.Sig]; ok {
			return errors.New(fmt.Sprintf("消息ID[%s]重复！！！！！！[%s]<->[%s]", v.Sig, v.FuncName, name))
		} else {
			funcSigs[v.Sig] = v.FuncName
		}
		if sig, ok := funcNames[v.FuncName]; ok {
			return errors.New(fmt.Sprintf("消息名字[%s]重复！！！！！！[%s]<->[%s]", v.FuncName, v.Sig, sig))
		} else {
			funcNames[v.FuncName] = v.Sig
		}
	}
	return nil
}
