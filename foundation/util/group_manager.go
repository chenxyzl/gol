package util

import (
	"context"
	"sync"
)

type GroupManager struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
	once   sync.Once
}

func NewGroupManager() *GroupManager {
	ret := new(GroupManager)
	ret.ctx, ret.cancel = context.WithCancel(context.Background())
	return ret
}

func (this *GroupManager) Close() {
	this.once.Do(this.cancel)
}

func (this *GroupManager) Wait() {
	this.wg.Wait()
}

func (this *GroupManager) Add(delta int) {
	this.wg.Add(delta)
}

func (this *GroupManager) Done() {
	this.wg.Done()
}

func (this *GroupManager) Chan() <-chan struct{} {
	return this.ctx.Done()
}
