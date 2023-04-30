package solution_101_150

import (
	"fmt"
)

// 149. 直线上最多的点数
// 给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。



var kbMap149 map[string][]int


func max149(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func getKB149(p1 []int, p2 []int) string {
	if p1[0] == p2[0] {
		return fmt.Sprintf("X=%d", p1[0])
	} else if p1[1] == p2[1] {
		return fmt.Sprintf("Y=%d", p1[1])
	}

	dx := float64(p1[0] - p2[0])
	dy := float64(p1[1] - p2[1])

	k := dy / dx
	b := float64(p1[1]) - float64(p1[0]) * k
	return fmt.Sprintf("%.3f_%.3f", k, b)
}


func maxPoints(points [][]int) int {
	l := len(points)
	if l <= 2 {
		return l
	}

	mx := 0
	kbMap149 = make(map[string][]int)
	for i := 0; i < l; i++ {
		firstVisited := make(map[string]bool)
		for j := i + 1; j < l; j++ {
			kb := getKB149(points[i], points[j])
			arr, ok := kbMap149[kb]
			if !ok {
				kbMap149[kb] = []int{i, j}
				firstVisited[kb] = true
			} else {
				_, fv := firstVisited[kb]
				if fv {
					kbMap149[kb] = append(arr, j)
				}
			}
			arr, _ = kbMap149[kb]
			mx = max149(mx, len(arr))
		}
	}
	fmt.Printf("%+v\n", kbMap149)
	return mx
}



