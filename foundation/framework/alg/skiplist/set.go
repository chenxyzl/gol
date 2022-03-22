package skiplist

type Valuer interface {
	Key() uint64
	Score() uint64
	ReCalcScore()
}

type Set struct {
	sl    *SkipList
	index map[uint64]Valuer
}

func NewSet(cmp Comparatorer) *Set {
	return &Set{
		sl:    NewSkipList(cmp),
		index: make(map[uint64]Valuer),
	}
}

func (s *Set) Length() uint32 { return s.sl.Length() }

func (s *Set) Head() *Node { return s.sl.Head() }

func (s *Set) Tail() *Node { return s.sl.Tail() }

func (s *Set) First() *Node { return s.sl.First() }

//Insert (先调用删除,在调用insert)
//必须确保删除和添加时候的key和score是和之前的一致
func (s *Set) Insert(value Valuer) {
	s.Delete(value)
	value.ReCalcScore()
	s.sl.Insert(value)
	s.index[value.Key()] = value
}

func (s *Set) GetElement(key uint64) Valuer {
	if value, exist := s.index[key]; exist {
		return value
	}
	return nil
}

func (s *Set) Delete(value Valuer) {
	if value, exist := s.index[value.Key()]; exist {
		delete(s.index, value.Key())
		s.sl.Delete(value)
	}
}

func (s *Set) DeleteElement(key uint64) {
	if value, exist := s.index[key]; exist {
		s.Delete(value)
	}
}

func (s *Set) GetByRank(rank uint32) interface{} {
	v := s.GetNodeByRank(rank)
	if v == nil {
		return nil
	}
	return v.Value()
}

func (s *Set) GetRank(key uint64) uint32 {
	if value, exist := s.index[key]; exist {
		return s.sl.GetRank(value)
	}
	return 0
}

func (s *Set) GetNodeByRank(rank uint32) *Node {
	return s.sl.GetNodeByRank(rank)
}

func (s *Set) DeleteRangeByRank(start, end uint32) uint32 {
	level := int(s.sl.Level())
	update := make([]*Node, level)
	var removed uint32 = 0
	var traversed uint32 = 0
	x := s.sl.Head()
	for i := level - 1; i >= 0; i-- {
		for next := x.Next(i); next != nil &&
			x.Span(i)+traversed < start; next = x.Next(i) {
			traversed += x.Span(i)
			x = next
		}
		update[i] = x
	}
	x = x.Next(0)
	traversed++
	for x != nil && traversed <= end {
		next := x.Next(0)
		s.sl.DeleteNode(x, update)
		delete(s.index, x.Value().(Valuer).Key())
		removed++
		traversed++
		x = next
	}
	return removed
}

func (s *Set) Dump() {
	s.sl.Dump()
}

//1-based rank
func (s *Set) GetRightRange(start, end uint32, reversal bool) (uint32, uint32) {
	length := s.sl.Length()
	if length == 0 || start == 0 || end < start || start > length {
		return 0, 0
	}
	if reversal {
		start = length + 1 - start
		if end > length {
			end = 1
		} else {
			end = length + 1 - end
		}
	} else {
		if end > length {
			end = length
		}
	}
	return start, end
}

// GetRange return 1-based elements in [start, end]
func (s *Set) GetRange(start uint32, end uint32, reverse bool) []interface{} {
	// var retKey []uint64
	// var retScore []uint64
	var out []interface{}
	if start == 0 {
		start = 1
	}
	if end == 0 {
		end = s.sl.length
	}
	if start > end || start > s.sl.length {
		return out
	}
	if end > s.sl.length {
		end = s.sl.length
	}
	rangeLen := end - start + 1
	if reverse {
		node := s.sl.GetNodeByRank(s.sl.length - start + 1)
		for i := uint32(0); i < rangeLen; i++ {
			// retKey = append(retKey, node.Value().(Valuer).Key())
			// retScore = append(retScore, node.Value().(Valuer).Score())
			out = append(out, node.Value())
			node = node.backward
		}
	} else {
		node := s.sl.GetNodeByRank(start)
		for i := uint32(0); i < rangeLen; i++ {
			// retKey = append(retKey, node.Value().(Valuer).Key())
			// retScore = append(retScore, node.Value().(Valuer).Score())
			out = append(out, node.Value())
			node = node.level[0].forward
		}
	}
	// return retKey, retScore
	return out
}

type RangeSpec struct {
	MinEx, MaxEx bool
	Min, Max     uint64
}

func (s *Set) ValueGteMin(value uint64, spec *RangeSpec) bool {
	if spec.MinEx {
		return value > spec.Min
	}
	return value >= spec.Min
}

func (s *Set) ValueLteMax(value uint64, spec *RangeSpec) bool {
	if spec.MaxEx {
		return value < spec.Max
	}
	return value <= spec.Max
}

func (s *Set) IsInRange(rg *RangeSpec) bool {
	if rg.Min > rg.Max ||
		(rg.Min == rg.Max && (rg.MinEx || rg.MaxEx)) {
		return false
	}

	x := s.sl.Tail()
	if x == nil || !s.ValueGteMin(x.Value().(Valuer).Score(), rg) {
		return false
	}

	x = s.sl.First()
	if x == nil || !s.ValueLteMax(x.Value().(Valuer).Score(), rg) {
		return false
	}
	return true
}

func (s *Set) FirstInRange(rg *RangeSpec) *Node {
	if !s.IsInRange(rg) {
		return nil
	}

	x := s.sl.Head()
	for i := int(s.sl.Level() - 1); i >= 0; i-- {
		for next := x.Next(i); next != nil &&
			!s.ValueGteMin(next.Value().(Valuer).Score(), rg); next = x.Next(i) {
			x = next
		}
	}
	x = x.Next(0)
	if !s.ValueLteMax(x.Value().(Valuer).Score(), rg) {
		return nil
	}
	return x
}

func (s *Set) LastInRange(rg *RangeSpec) *Node {
	if !s.IsInRange(rg) {
		return nil
	}

	x := s.sl.Head()
	for i := int(s.sl.Level() - 1); i >= 0; i-- {
		for next := x.Next(i); next != nil &&
			s.ValueLteMax(next.Value().(Valuer).Score(), rg); next = x.Next(i) {
			x = next
		}
	}
	if !s.ValueGteMin(x.Value().(Valuer).Score(), rg) {
		return nil
	}
	return x
}

func (s *Set) DeleteRangeByScore(rg *RangeSpec) uint32 {
	update := make([]*Node, int(s.sl.Level()))
	var removed uint32 = 0
	x := s.sl.Head()
	for i := int(s.sl.Level() - 1); i >= 0; i-- {
		for next := x.Next(i); next != nil &&
			((rg.MinEx && next.Value().(Valuer).Score() <= rg.Min) ||
				(!rg.MinEx && next.Value().(Valuer).Score() < rg.Min)); next = x.Next(i) {
			x = next
		}
		update[i] = x
	}
	x = x.Next(0)
	for x != nil &&
		((rg.MaxEx && x.Value().(Valuer).Score() < rg.Max) ||
			(!rg.MaxEx && x.Value().(Valuer).Score() <= rg.Max)) {
		next := x.Next(0)
		s.sl.DeleteNode(x, update)
		delete(s.index, x.Value().(Valuer).Key())
		removed++
		x = next
	}
	return removed
}

func (s *Set) GetRangeByScore(rg *RangeSpec) []interface{} {
	var values []interface{}
	x := s.FirstInRange(rg)
	for x != nil {
		if !s.ValueLteMax(x.Value().(Valuer).Score(), rg) {
			break
		}
		values = append(values, x.value)
		x = x.Next(0)
	}
	return values
}

func (s *Set) Range(f func(interface{})) {
	for tmp := s.First(); tmp != nil; tmp = tmp.Next(0) {
		f(tmp.Value())
	}
}
