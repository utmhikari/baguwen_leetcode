package solution_301_350

import (
	"fmt"
	"testing"
)

//335. 路径交叉
//
//给定一个含有 n 个正数的数组 x。
//从点 (0,0) 开始，先向北移动 x[0] 米，然后向西移动 x[1] 米，
//向南移动 x[2] 米，向东移动 x[3] 米，持续移动。也就是说，每次移动后你的方位会发生逆时针变化。
//
//编写一个 O(1) 空间复杂度的一趟扫描算法，判断你所经过的路径是否相交。



func isSelfCrossing(distance []int) bool {
	l := len(distance)
	if l < 4 {
		return false
	}

	borders := [][]int{
		{0,0,0}, // up line (y, x1, x2)
		{0,0,0}, // left line (x, y1, y2)
		{0,0,0}, // down line (y, x1, x2)
		{0,0,0}, // right line (x, y1, y2)
	}

	point := []int{0,0}

	for i := 0; i < l; i++ {
		dis := distance[i]
		mod := i % 4
		nextPoint := []int{point[0], point[1]}

		switch mod {
		case 0: // up
			nextPoint[1] += dis
			if i > 2 {
				// cross with up line?
				if nextPoint[1] >= borders[0][0] &&
					nextPoint[0] >= borders[0][1] &&
					nextPoint[0] <= borders[0][2] {
					return true
				}
			}
			// cross with right line?
			if point[0] == borders[3][0] && point[1] < borders[3][1] && nextPoint[1] >= borders[3][1] {
				return true
			}
			// update right line
			if i >= 3 && point[0] > borders[3][0] && nextPoint[1] >= borders[3][1] {
				// prev right line should be left line
				borders[1] = borders[3]
			}
			borders[3] = []int{point[0], point[1], nextPoint[1]}
			break
		case 1: // left
			nextPoint[0] -= dis
			if i > 2 {
				// cross with left line?
				if nextPoint[0] <= borders[1][0] &&
					nextPoint[1] >= borders[1][1] &&
					nextPoint[1] <= borders[1][2] {
					return true
				}
			}
			// update up line
			if i >= 3 && point[1] > borders[0][0] && nextPoint[0] <= borders[0][2] {
				// prev up line should be down line
				borders[2] = borders[0]
			}
			borders[0] = []int{point[1], nextPoint[0], point[0]}
			break
		case 2: // down
			nextPoint[1] -= dis
			// cross with down line?
			if i > 2 {
				if nextPoint[1] <= borders[2][0] &&
					nextPoint[0] >= borders[2][1] &&
					nextPoint[0] <= borders[2][2] {
					return true
				}
			}
			// update left line
			borders[1] = []int{point[0], nextPoint[1], point[1]}
			break
		case 3: // right
			nextPoint[0] += dis
			// cross with right line?
			if i > 2 {
				if nextPoint[0] >= borders[3][0] &&
					nextPoint[1] >= borders[3][1] &&
					nextPoint[1] <= borders[3][2] {
					return true
				}
			}
			// update down line
			borders[2] = []int{point[1], point[0], nextPoint[0]}
			break
		default:
			break
		}

		point = nextPoint
	}

	return false
}


func Test_335(t *testing.T) {
	input := []int{3,3,3,2,1,1}
	expected := false
	ans := isSelfCrossing(input)
	if ans != expected {
		t.Error(fmt.Sprintf("failed at ans %v", input))
	} else {
		fmt.Println("success")
	}
}