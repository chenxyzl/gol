package framework

import (
	"foundation/framework/bif"
	"foundation/framework/g"
	"gitlab-ee.funplus.io/watcher/watcher/misc/wlog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(actor bif.IActor) {
	g.Root = actor
	//加载db
	wlog.Warn("load begin...")
	actor.Load()
	wlog.Warn("load complete...")
	//开始
	wlog.Warn("start begin...")
	actor.Start()
	wlog.Warn("start complete...")
	//退出信号
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	//tick
	ticker := time.NewTicker(time.Second)
LOOP:
	for {
		select {
		case sig := <-sigterm:
			wlog.Warnf("[main] os sig=[%v]", sig)
			break LOOP
		case now := <-ticker.C:
			actor.Tick(now.Unix())
		}
	}

	//停止
	wlog.Warn("stop begin...")
	actor.Stop()
	wlog.Warn("stop complete...")
	//销毁
	wlog.Warn("destroy begin...")
	actor.Destroy()
	wlog.Warn("destroy complete...")
	//退出
	wlog.Warn("exist...")
}
