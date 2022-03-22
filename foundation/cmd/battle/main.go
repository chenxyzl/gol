package main

import (
	"fmt"
	"reflect"
)

type M[K string, V any] map[K]V

type A interface {
	GetName() string
}

var m = M[string, any]{}

//todo T应该增加类型限制
func Add[T any]() {
	var a = new(T)
	name := reflect.TypeOf(a).Name()
	if _, ok := m[name]; ok {
		panic(fmt.Sprintf("repeated add:%v", name))
	}
	m[reflect.TypeOf(a).Name()] = a
}

func Get[T any]() T {
	//todo 按照其他语言这里应该是可以直接使用typeof(T)的
	var a T
	var name = reflect.TypeOf(a).Name()
	v, ok := m[name]
	//todo 不ok需要返回nil
	_ = ok
	//if !ok {
	//	return nil
	//}
	//todo 逻辑上应该不用再转一次，实际不转编译不过
	v1, ok1 := v.(T)
	_ = ok1
	//if !ok1 {
	//	return nil
	//}
	return v1
}

type A1 struct {
}

func (a *A1) GetName() string {
	return "A1"
}

type A2 struct {
}

func (a *A2) GetName() string {
	return "A2"
}
func main() {
	//Add[int]()
	Add[A1]()
	Add[A2]()
	a1 := Get[A1]()
	println(a1.GetName())
}
