package ifs_base

import "foundation/framework/component"

type IComponent interface {
	Constructor(...interface{}) //初始化函数
	Name() component.ComType
	Load()
	Start()
	Tick(int64 int64)
	Stop()
	Destroy()
}
