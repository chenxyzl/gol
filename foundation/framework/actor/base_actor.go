package actor

import (
	"context"
	"fmt"
	"foundation/framework/component"
	"foundation/framework/component/ifs/ifs.base"
	"foundation/framework/message"
	"golang.org/x/sync/semaphore"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	AsyncTimeOut = time.Duration(1000) * time.Millisecond
)

var _ = (IActor)(&Actor{})

type Actor struct {
	IActor
	//线程调度相关
	lock             *semaphore.Weighted
	goNumLock        sync.Mutex
	maxRunningGoSize int32 //size等于1就等同于单线程了
	runningGoNum     int32
	//邮箱
	Boxs chan message.IMessage
	//组件相关
	components        []ifs_base.IComponent
	componentsMapping map[component.ComType]ifs_base.IComponent //go的泛型太辣鸡了。暂时不用
}

func (actor *Actor) Constructor(boxSize int32, maxRunningGoSize int32) {
	if boxSize <= 0 {
		panic("boxSize must bigger than 0")
	}
	actor.lock = semaphore.NewWeighted(int64(1))
	actor.Boxs = make(chan message.IMessage, boxSize)
	actor.maxRunningGoSize = maxRunningGoSize
	//
	actor.components = make([]ifs_base.IComponent, 0)
	actor.componentsMapping = make(map[component.ComType]ifs_base.IComponent)
}

//GetComponent 获取组件
func (actor *Actor) GetComponent(comType component.ComType) ifs_base.IComponent {
	return actor.componentsMapping[comType]
}

func (actor *Actor) AddComponent(iComponent ifs_base.IComponent, params ...interface{}) {
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

func (actor *Actor) Lock() {
	ctx := context.Background()
	actor.lock.Acquire(ctx, 1)
}

func (actor *Actor) Unlock() {
	actor.lock.Release(1)
}

func (actor *Actor) LockGoNum() {
	actor.goNumLock.Lock()
}

func (actor *Actor) UnlockGoNum() {
	actor.goNumLock.Unlock()
}

func (actor *Actor) BeginRecv() {
	go func() {
		for message := range actor.Boxs {
			actor.LockGoNum()
			for {
				//如果已达到上线则切换到别的go程
				if actor.runningGoNum >= actor.maxRunningGoSize {
					runtime.Gosched()
				} else {
					atomic.AddInt32(&actor.runningGoNum, 1)
					break
				}
			}
			actor.UnlockGoNum()
			actor.Lock()
			go func() {
				defer func() {
					atomic.AddInt32(&actor.runningGoNum, -1)
					actor.Unlock()
				}()
				actor.OnRecv(message)
			}()
		}
	}()
}

// AsyncCall f内的代码不是线程安全的，通常这个函数为基础函数，不需要业务开发人员调用/
func (actor *Actor) AsyncCall(f func()) {
	actor.Unlock()
	defer func() {
		actor.Lock()
	}()

	c := make(chan bool)

	timeoutTimer := time.After(AsyncTimeOut)
	go func() {
		f()
		c <- true
	}()
	select {
	case <-timeoutTimer:
		return
	case <-c:
		return
	}
}

//Load 生命周期函数
func (actor *Actor) Load() {
	for _, com := range actor.components {
		com.Load()
	}
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

func (actor *Actor) AsyncAsk(target IActorRef, msg *message.IMessage) *message.IMessage {
	//todo 更具一致性hash算法 找到对应的nodeId
	//todo AsyncCall调用nats的request/reply来发起请求
	return nil
}

func (actor *Actor) AsyncTell(target IActorRef, msg *message.IMessage) *message.IMessage {
	//todo 更具一致性hash算法 找到对应的nodeId
	//todo AsyncCall调用用nats的notify发送消息
	return nil
}
