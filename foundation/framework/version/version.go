package version

import (
	"foundation/framework/version_impl"
	"runtime"
)

//-trimpath -tags -ldflags "-X foundation/framework/version.Version=20200508065715-24e910024bf8 -X foundation/framework/version.VersionMsg=6LCD5pW05LiA5LiL6YWN572uCg## -X foundation/framework/version.DebugMode=true "
var (
	//编译时注入
	Version    = "" //20200508065715-24e910024bf8
	VersionMsg = "" //6LCD5pW05LiA5LiL6YWN572uCg##
	BuildTime  = "dev-null"
	CommitID   = "dev-null"
	GoVer      = runtime.Version()
	//版本管理
	VersionMng *version_impl.VersionMng
)
