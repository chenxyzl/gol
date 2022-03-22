package version_impl

import (
	"fmt"
	"strconv"
	"strings"
)

var prefixs = []string{"v", "V", "Ver", "ver", "Version", "version"}

// 支持若干级的版本号对比
// 默认对比3位， SetCheckNum()可以修改对比位数
// 支持特定前缀， 支持任意后缀（用-分开）
type Version struct {
	all      string
	prefix   string
	suffix   string
	number   []int64
	validNum int
}

func (v *Version) String() string {
	return v.all
}

func (v *Version) SetCheckNum(num int) {
	v.validNum = num
}

func NewVersion(c string) (*Version, error) {
	v := &Version{
		all:      c,
		number:   []int64{},
		validNum: 3,
	}
	for i := 0; i < len(prefixs); i++ {
		prefix := prefixs[len(prefixs)-1-i]
		if strings.HasPrefix(c, prefix) {
			v.prefix = prefix
			break
		}
	}
	strArray1 := strings.Split(c, "-")

	if len(strArray1) > 1 {
		for i := 0; i < len(strArray1)-1; i++ {
			if strings.HasSuffix(strArray1[i], ".") {
				return nil, fmt.Errorf("format err, number > 0 : %s", c)
			}
		}
		s := strArray1[1]
		for i := 2; i < len(strArray1); i++ {
			s += strArray1[i]
		}
		v.suffix = s
	}
	strArray2 := strings.TrimPrefix(strArray1[0], v.prefix)
	numbers := strings.Split(strArray2, ".")
	for index, number := range numbers {
		i, err := strconv.Atoi(number)
		if err != nil {
			return nil, fmt.Errorf("number err, index : %d, value : %s", index, number)
		}
		if i < 0 {
			return nil, fmt.Errorf("number need >0, index : %d, value : %s", index, number)
		}
		v.number = append(v.number, int64(i))
	}
	return v, nil
}

// Equals checks if v is equal to o.
func (v Version) Equals(o Version) bool {
	return (v.Compare(o) == 0)
}

// EQ checks if v is equal to o.
func (v Version) EQ(o Version) bool {
	return (v.Compare(o) == 0)
}

// NE checks if v is not equal to o.
func (v Version) NE(o Version) bool {
	return (v.Compare(o) != 0)
}

// GT checks if v is greater than o.
func (v Version) GT(o Version) bool {
	return (v.Compare(o) == 1)
}

// GTE checks if v is greater than or equal to o.
func (v Version) GTE(o Version) bool {
	return (v.Compare(o) >= 0)
}

// GE checks if v is greater than or equal to o.
func (v Version) GE(o Version) bool {
	return (v.Compare(o) >= 0)
}

// LT checks if v is less than o.
func (v Version) LT(o Version) bool {
	return (v.Compare(o) == -1)
}

// LTE checks if v is less than or equal to o.
func (v Version) LTE(o Version) bool {
	return (v.Compare(o) <= 0)
}

// LE checks if v is less than or equal to o.
func (v Version) LE(o Version) bool {
	return (v.Compare(o) <= 0)
}

// return -1, v < o ; return +1, v > o ; return 0, v == o
func (v Version) Compare(o Version) int {

	for i := 0; i < len(v.number); i++ {
		if i >= v.validNum {
			return 0
		}
		if len(o.number) > i {
			if v.number[i] > o.number[i] {
				return -1
			} else if v.number[i] < o.number[i] {
				return 1
			}
			continue
		} else {
			if v.number[i] > 0 {
				return -1
			}
			continue
		}
	}
	return 0
}
