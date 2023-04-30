package solution_101_150

import (
	"fmt"
	"reflect"
	"testing"
)

//131. 分割回文串
//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。


var cache131 map[int]map[int]bool

func isPalindrome131(s string, left int, right int) bool {
	mp, ok := cache131[left]
	if ok {
		_, ok = mp[right]
		if ok {
			return true
		}
	}

	if right - left > 0 {
		for right > left {
			if s[left] != s[right] {
				return false
			}
			right--
			left++
		}
	}

	mp, ok = cache131[left]
	if !ok {
		cache131[left] = make(map[int]bool)
	}
	mp, ok = cache131[left]
	if ok {
		mp[right] = true
	}

	return true
}


func partition131Internal(s string, left int, right int) [][]string {
	ans := make([][]string, 0)

	if right < left {
		return ans
	}

	for i := left; i <= right; i++ {
		if isPalindrome131(s, left, i) {
			if i == right {
				ans = append(ans, []string{s[left:right+1]})
			} else {
				otherAns := partition131Internal(s, i + 1, right)
				fmt.Printf("%d, %d, %s, %v+\n", left, i, s[left:i+1], otherAns)
				for _, otherItem := range otherAns {
					item := make([]string, len(otherItem) + 1)
					item[0] = s[left:i + 1]
					for j := 0; j < len(otherItem); j++ {
						item[j + 1] = otherItem[j]
					}
					ans = append(ans, item)
				}
			}
		}
	}

	return ans
}


func partition131(s string) [][]string {
	cache131 = make(map[int]map[int]bool)
	if len(s) == 0 {
		return make([][]string, 0)
	}
	return partition131Internal(s, 0, len(s) - 1)
}


type test131 struct {
	Input string
	Expected [][]string
}

func Test_131(t *testing.T) {
	cases := append(
		make([]test131, 0),
		test131{
			Input: "aab",
			Expected: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		test131{
			Input: "a",
			Expected: [][]string{
				{"a"},
			},
		},
		test131{
			Input: "aa",
			Expected: [][]string{
				{"a", "a"},
				{"aa"},
			},
		},
	)

	for i, c := range cases {
		ans := partition131(c.Input)
		if !reflect.DeepEqual(ans, c.Expected) {
			t.Errorf("failed at %d -> %v+, %v+\n", i, c, ans)
		} else {
			fmt.Printf("successful at %d -> %v+\n", i, c)
		}
	}
}


