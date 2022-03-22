package skiplist

import (
	"bytes"
	"fmt"
	"math/rand"
)

var P float32 = 0.25

const level int = 32

type Level struct {
	forward *Node
	span    uint32
}

type Node struct {
	value    interface{}
	backward *Node
	level    []Level
}

func NewSkipListNode(level int, value interface{}) *Node {
	sln := &Node{
		value: value,
		level: make([]Level, level),
	}
	return sln
}

func (this *Node) Next(i int) *Node {
	return this.level[i].forward
}

func (this *Node) SetNext(i int, next *Node) {
	this.level[i].forward = next
}

func (this *Node) Span(i int) uint32 {
	return this.level[i].span
}

func (this *Node) SetSpan(i int, span uint32) {
	this.level[i].span = span
}

func (this *Node) Value() interface{} {
	return this.value
}

func (this *Node) Prev() *Node {
	return this.backward
}

type Comparatorer interface {
	CmpScore(interface{}, interface{}) int
	CmpKey(interface{}, interface{}) int
}

type SkipList struct {
	head, tail    *Node
	length, level uint32
	Comparatorer
}

func NewSkipList(cmp Comparatorer) *SkipList {
	sl := &SkipList{
		level:        1,
		length:       0,
		tail:         nil,
		Comparatorer: cmp,
	}
	sl.head = NewSkipListNode(level, nil)
	for i := 0; i < level; i++ {
		sl.head.SetNext(i, nil)
		sl.head.SetSpan(i, 0)
	}
	sl.head.backward = nil
	return sl
}

func (sl *SkipList) Level() uint32 { return sl.level }

func (sl *SkipList) Length() uint32 { return sl.length }

func (sl *SkipList) Head() *Node { return sl.head }

func (sl *SkipList) Tail() *Node { return sl.tail }

func (sl *SkipList) First() *Node { return sl.head.Next(0) }

func (sl *SkipList) randomLevel() int {
	level := 1
	for (rand.Uint32()&0xFFFF) < uint32(P*0xFFFF) && level < level {
		level++
	}
	return level
}

func (sl *SkipList) Insert(value interface{}) *Node {
	var update [level]*Node
	var rank [level]uint32
	x := sl.head
	for i := int(sl.level - 1); i >= 0; i-- {
		if i == int(sl.level-1) {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for next := x.Next(i); next != nil &&
			(sl.CmpScore(next.value, value) < 0 ||
				(sl.CmpScore(next.value, value) == 0 &&
					sl.CmpKey(next.value, value) < 0)); next = x.Next(i) {
			rank[i] += x.Span(i)
			x = next
		}
		update[i] = x
	}

	level := uint32(sl.randomLevel())

	if level > sl.level {
		for i := sl.level; i < level; i++ {
			rank[i] = 0
			update[i] = sl.head
			update[i].SetSpan(int(i), sl.length)
		}
		sl.level = level
	}

	x = NewSkipListNode(int(level), value)
	for i := 0; i < int(level); i++ {
		x.SetNext(i, update[i].Next(i))
		update[i].SetNext(i, x)

		x.SetSpan(i, update[i].Span(i)-(rank[0]-rank[i]))
		update[i].SetSpan(i, rank[0]-rank[i]+1)
	}

	for i := level; i < sl.level; i++ {
		update[i].SetSpan(int(i), update[i].Span(int(i))+1)
	}

	if update[0] == sl.head {
		x.backward = nil
	} else {
		x.backward = update[0]
	}

	if x.Next(0) != nil {
		x.Next(0).backward = x
	} else {
		sl.tail = x
	}
	sl.length++
	return x
}

func (sl *SkipList) DeleteNode(x *Node, update []*Node) {
	for i := 0; i < int(sl.level); i++ {
		if update[i].Next(i) == x {
			update[i].SetSpan(i, update[i].Span(i)+x.Span(i)-1)
			update[i].SetNext(i, x.Next(i))
		} else {
			update[i].SetSpan(i, update[i].Span(i)-1)
		}
	}

	if x.Next(0) != nil {
		x.Next(0).backward = x.backward
	} else {
		sl.tail = x.backward
	}

	for sl.level > 1 && sl.head.Next(int(sl.level-1)) == nil {
		sl.level--
	}
	sl.length--
}

func (sl *SkipList) Delete(value interface{}) int {
	update := make([]*Node, int(sl.level))
	var x *Node = sl.head
	for i := int(sl.level - 1); i >= 0; i-- {
		for next := x.Next(i); next != nil &&
			(sl.CmpScore(next.value, value) < 0 ||
				(sl.CmpScore(next.value, value) == 0 &&
					sl.CmpKey(next.value, value) < 0)); next = x.Next(i) {
			x = next
		}
		update[i] = x
	}

	x = x.Next(0)
	if x != nil &&
		sl.CmpKey(x.value, value) == 0 &&
		sl.CmpScore(x.value, value) == 0 {
		sl.DeleteNode(x, update)
		return 1
	}
	return 0
}

//TODO: 1-based rank
func (sl *SkipList) GetRank(value interface{}) uint32 {
	var rank uint32 = 0
	x := sl.head
	for i := int(sl.level - 1); i >= 0; i-- {
		for next := x.Next(i); next != nil &&
			(sl.CmpScore(next.value, value) < 0 ||
				(sl.CmpScore(next.value, value) == 0 &&
					sl.CmpKey(next.value, value) <= 0)); next = x.Next(i) {
			rank += x.Span(i)
			x = next
		}
		if x != sl.head && sl.CmpKey(x.value, value) == 0 {
			return rank
		}
	}
	return 0
}

func (sl *SkipList) GetNodeByRank(rank uint32) *Node {
	x := sl.head
	var traversed uint32 = 0
	for i := int(sl.level - 1); i >= 0; i-- {
		for next := x.Next(i); next != nil &&
			traversed+x.Span(i) <= rank; next = x.Next(i) {
			traversed += x.Span(i)
			x = next
		}
		if traversed == rank {
			return x
		}
	}
	return nil
}

func (sl *SkipList) Dump() {
	fmt.Println("*************SKIP LIST DUMP START*************")
	for i := int(sl.level - 1); i >= 0; i-- {
		fmt.Printf("level:--------%v--------\n", i)
		x := sl.head
		for x != nil {
			if x == sl.head {
				fmt.Printf("Head span: %v\n", x.Span(i))
			} else {
				fmt.Printf("span: %v value : %v\n", x.Span(i), x.Value())
			}
			x = x.Next(i)
		}
	}
	fmt.Println("*************SKIP LIST DUMP END*************")
}

func (sl *SkipList) DumpString() string {
	var buffer bytes.Buffer
	for i := int(sl.level - 1); i >= 0; i-- {
		buffer.WriteString(fmt.Sprintf("level:--------%v--------\n", i))
		x := sl.head
		for x != nil {
			if x == sl.head {
				buffer.WriteString(fmt.Sprintf("Head span: %v\n", x.Span(i)))
			} else {
				buffer.WriteString(fmt.Sprintf("span: %v value : %+v\n", x.Span(i), x.Value()))
			}
			x = x.Next(i)
		}
	}
	return buffer.String()
}
