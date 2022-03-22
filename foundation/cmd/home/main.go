package main

import (
	"foundation/framework"
	"foundation/framework/g"
	"math/rand"
	"time"

	"gitlab-ee.funplus.io/watcher/watcher/misc"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// For the core dump
	err := misc.ResetWorkingDirectory()
	if err != nil {
		panic(err)
	}

	//设置角色类型
	g.SetRoleType(g.Home)

	//runCmdCobra.Execute()
	framework.Execute(runCmdCobra)
}
