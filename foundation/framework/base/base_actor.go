package base

import (
	"context"
	"fmt"
	"foundation/framework/bif"
	"foundation/framework/component"
	"foundation/framework/component/ifs"
	"foundation/framework/g"
	"foundation/message"
	"github.com/golang/protobuf/proto"
	"gitlab-ee.funplus.io/watcher/watcher/misc/wlog"
	"golang.org/x/sync/semaphore"
	"runtime"
	"sync"
	"sync/atomic"
)

var _ = (bif.IActor)(&Actor{})

type Actor struct {
	bif.IActor
	//线程调度相关
	mux              *semaphore.Weighted
	goNumLock        sync.Mutex
	maxRunningGoSize int32 //size等于1就等同于单线程了
	runningGoNum     int32
	//邮箱
	Boxs chan any
	//组件相关
	components        []bif.IComponent
	componentsMapping map[component.ComType]bif.IComponent //go的泛型太辣鸡了。暂时不用
	forceSync         bool                                 //加载和停止中强制走同步
}

func (actor *Actor) Constructor(boxSize int32, maxRunningGoSize int32) {
	if boxSize <= 0 {
		panic("boxSize must bigger than 0")
	}
	actor.mux = semaphore.NewWeighted(int64(1))
	actor.Boxs = make(chan any, boxSize)
	actor.maxRunningGoSize = maxRunningGoSize
	//
	actor.forceSync = true
	//
	actor.components = make([]bif.IComponent, 0)
	actor.componentsMapping = make(map[component.ComType]bif.IComponent)
}

//GetComponent 获取组件
func (actor *Actor) GetComponent(comType component.ComType) bif.IComponent {
	return actor.componentsMapping[comType]
}

func (actor *Actor) AddComponent(iComponent bif.IComponent, params ...interface{}) {
	//重复检查
	if _, ok := actor.componentsMapping[iComponent.Name()]; ok {
		panic(fmt.Sprintf("component name:%s repeated", iComponent.Name()))
	}
	//构造
	iComponent.Constructor(params...)
	//
	actor.components = append(actor.components, iComponent)
	actor.componentsMapping[iComponent.Name()] = iComponent

}

func (actor *Actor) lock() {
	ctx := context.Background()
	actor.mux.Acquire(ctx, 1)
}

func (actor *Actor) release() {
	actor.mux.Release(1)
}

func (actor *Actor) lockGoNum() {
	actor.goNumLock.Lock()
}

func (actor *Actor) unlockGoNum() {
	actor.goNumLock.Unlock()
}

func (actor *Actor) checkGoNum() {
	actor.lockGoNum()
	for {
		//如果已达到上线则切换到别的go程
		if actor.runningGoNum >= actor.maxRunningGoSize ||
			actor.forceSync && actor.runningGoNum > 0 {
			runtime.Gosched()
		} else {
			atomic.AddInt32(&actor.runningGoNum, 1)
			break
		}
	}
	actor.unlockGoNum()
}

func (actor *Actor) asyncDo(message any) {
	//检查go程序
	actor.checkGoNum()
	//执行异步代码
	actor.lock()
	go func() {
		defer func() {
			atomic.AddInt32(&actor.runningGoNum, -1)
			actor.release()
		}()
		actor.OnRecv(message)
	}()
}

//SafeAsyncDo 同步执行一些事情～ 注意这里不要执行长耗时和异步操作
func (actor *Actor) SafeAsyncDo(f func()) {
	//检查go程序
	actor.checkGoNum()
	//执行异步代码
	actor.lock()
	go func() {
		defer func() {
			atomic.AddInt32(&actor.runningGoNum, -1)
			actor.release()
		}()
		//
		f()
	}()
}

func (actor *Actor) AddMessage(message any) {
	actor.Boxs <- message
}

func (actor *Actor) BeginRecv() {
	go func() {
		for message := range actor.Boxs {
			actor.asyncDo(message)
		}
	}()
}

//Load 生命周期函数
func (actor *Actor) Load() {
	for _, com := range actor.components {
		com.Load()
	}
	//
	actor.forceSync = false
}

//Start 生命周期函数
func (actor *Actor) Start() {
	for _, com := range actor.components {
		com.Start()
	}
}

//Tick 生命周期函数
func (actor *Actor) Tick(time int64) {
	for _, com := range actor.components {
		com.Tick(time)
	}
}

//Stop 生命周期函数
func (actor *Actor) Stop() {
	actor.forceSync = true
	for _, com := range actor.components {
		com.Stop()
	}
}

//Destroy 生命周期函数
func (actor *Actor) Destroy() {
	for _, com := range actor.components {
		com.Destroy()
	}
}

func (actor *Actor) To() *message.ActorRef {
	return &message.ActorRef{
		Uid:  actor.GetUid(),
		Type: message.ActorType(message.ActorType_value[actor.GetType()]),
	}
}

func (actor *Actor) AsyncAsk(target message.IActorRef, cmd message.CMD, msg proto.Message, out proto.Message) message.Code {
	nats := g.Root.GetComponent(component.NatsCom).(ifs.INatsComponent)
	data, err := proto.Marshal(msg)
	if err != nil {
		return message.Code_MarshalError
	}
	req := &message.NatsRequest{
		ActorRef: target.To(),
		Cmd:      uint32(cmd),
		Data:     data,
	}
	rep := nats.Ask(req)
	if rep.Code == message.Code_OK {
		err := proto.Unmarshal(rep.Data, out)
		if err != nil {
			wlog.Error("[Actor].AsyncAsk unmarshal error:[%v]", err)
			rep.Code = message.Code_NatsReplyUnmarshalError
		}
	}
	return rep.Code
}

func (actor *Actor) AsyncTell(target message.IActorRef, cmd message.CMD, msg proto.Message) message.Code {
	nats := g.Root.GetComponent(component.NatsCom).(ifs.INatsComponent)
	data, err := proto.Marshal(msg)
	if err != nil {
		return message.Code_MarshalError
	}
	req := &message.NatsRequest{
		ActorRef: target.To(),
		Cmd:      uint32(cmd),
		Data:     data,
	}
	code := nats.Tell(req)
	return code
}
