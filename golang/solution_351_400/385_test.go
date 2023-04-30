package solution_351_400

import (
	"log"
)

//385. 迷你语法分析器
//给定一个用字符串表示的整数的嵌套列表，实现一个解析它的语法分析器。
//
//列表中的每个元素只可能是整数或整数嵌套列表
//
//提示：你可以假定这些字符串都是格式良好的：
//
//字符串非空
//字符串不包含空格
//字符串只包含数字0-9、[、-、,、]


/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
	v int
	ls []*NestedInteger
	isInt bool
}

func (n NestedInteger) IsInteger() bool {
	return n.isInt
}

func (n NestedInteger) GetInteger() int {
	if !n.isInt {
		return 0
	}
	return n.v
}

func (n *NestedInteger) SetInteger(value int) {
	n.isInt = true
	n.v = value
}

func (n *NestedInteger) Add(elem NestedInteger) {
	if !n.isInt {
		if n.ls == nil {
			n.ls = []*NestedInteger{&elem}
		} else {
			n.ls = append(n.ls, &elem)
		}
	}
}

func (n NestedInteger) GetList() []*NestedInteger {
	if n.isInt {
		return nil
	}
	if n.ls == nil {
		n.ls = make([]*NestedInteger, 0)
	}
	return n.ls
}


type NIState385 struct {
	tpFlag uint8  // unknown: 0, int: 1, arr: 2
	intSigned uint8  // unsigned: 0, signed +: 1, signed -: 2
	arrElemProcessing bool // notProcessing: 0, isProcessing: 1 (use comma to judge)
	shouldEnd bool // should end?
}


func solve385(s string, i int) (*NestedInteger, int) {
	var n *NestedInteger
	st := NIState385{
		tpFlag:            0,
		intSigned:         0,
		arrElemProcessing: false,
		shouldEnd:         false,
	}

	// check idx
	l := len(s)
	if i >= l {
		return nil, i
	}

	// check first char to decide array or int
	if s[i] == '[' {
		n = &NestedInteger{}
		st.tpFlag = 2
		i++
	} else if s[i] == '-' {
		n = &NestedInteger{}
		st.tpFlag = 1
		st.intSigned = 2
		i++
	} else if s[i] >= '0' && s[i] <= '9' {
		n = &NestedInteger{}
		st.tpFlag = 1
		st.intSigned = 1
	} else {
		log.Panicf("invalid start char at %d, not array or int", i)
	}

	for i < len(s) {
		c := s[i]
		if st.tpFlag == 1 {  // int
			switch {
			case c >= '0' && c <= '9':
				cur := n.GetInteger()
				if st.intSigned == 1 {
					n.SetInteger(10 * cur + int(c - '0'))
				} else {
					if cur == 0 {
						n.SetInteger(-int(c - '0'))
					} else {
						n.SetInteger(10 * cur - int(c - '0'))
					}
				}
				break
			default:
				i--
				st.shouldEnd = true
				break
			}
		} else {  // arr
			switch c {
			case ']':
				st.arrElemProcessing = false
				st.shouldEnd = true
				break
			case ',':
				if !st.arrElemProcessing {
					log.Panicf("invalid comma at %d, currently not processing elem", i)
				}
				st.arrElemProcessing = false
				break
			default:
				if st.arrElemProcessing {
					log.Panicf("invalid elem start char at %d, currently processing prev elem", i)
				}
				st.arrElemProcessing = true
				elem, endIdx := solve385(s, i)
				n.Add(*elem)
				i = endIdx
				break
			}
		}
		if st.shouldEnd {
			break
		}
		i++
	}


	if !st.shouldEnd {
		if st.tpFlag == 1 {
			i--
		} else {
			log.Panicf("end early at %d", i)
		}
	}
	return n, i
}


func deserialize(s string) *NestedInteger {
	n, i := solve385(s, 0)
	if i != len(s) - 1 {
		log.Panicf("left unprocessed chars start from %d -> %v", i + 1, n)
	}
	return n
}
