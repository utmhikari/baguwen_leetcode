package solution_351_400

import (
	"fmt"
	"math"
)

//391. 完美矩形
//我们有 N 个与坐标轴对齐的矩形, 其中 N > 0,
//判断它们是否能精确地覆盖一个矩形区域。
//
//每个矩形用左下角的点和右上角的点的坐标来表示。
//例如， 一个单位正方形可以表示为 [1,1,2,2]。
// ( 左下角的点的坐标为 (1, 1) 以及右上角的点的坐标为 (2, 2) )。


func isRectangleCover(rectangles [][]int) bool {
	sz, minX, minY, maxX, maxY := 0, math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32
	ranges := [][]int{
		{math.MaxInt32, math.MinInt32},  // minX -> maxX of minY
		{math.MaxInt32, math.MinInt32},  // minY -> maxY of minX
		{math.MaxInt32, math.MinInt32},  // minX -> maxX of maxY
		{math.MaxInt32, math.MinInt32},  // minY -> maxY of maxX
	}
	lengths := []int{0, 0, 0, 0}
	pointSet := make(map[string]bool)
	for _, ra := range rectangles {
		// update min/max x/y, ranges & lengths
		if ra[0] < minX {
			minX = ra[0]
			ranges[1] = []int{ra[1], ra[3]}
			lengths[1] = ra[3] - ra[1]
		} else if ra[0] == minX {
			if ra[1] < ranges[1][0] {
				ranges[1][0] = ra[1]
			}
			if ra[3] > ranges[1][1] {
				ranges[1][1] = ra[3]
			}
			lengths[1] += ra[3] - ra[1]
		}
		if ra[1] < minY {
			minY = ra[1]
			ranges[0] = []int{ra[0], ra[2]}
			lengths[0] = ra[2] - ra[0]
		} else if ra[1] == minY {
			if ra[0] < ranges[0][0] {
				ranges[0][0] = ra[0]
			}
			if ra[2] > ranges[0][1] {
				ranges[0][1] = ra[2]
			}
			lengths[0] += ra[2] - ra[0]
		}
		if ra[2] > maxX {
			maxX = ra[2]
			ranges[3] = []int{ra[1], ra[3]}
			lengths[3] = ra[3] - ra[1]
		} else if ra[2] == maxX {
			if ra[1] < ranges[3][0] {
				ranges[3][0] = ra[1]
			}
			if ra[3] > ranges[3][1] {
				ranges[3][1] = ra[3]
			}
			lengths[3] += ra[3] - ra[1]
		}
		if ra[3] > maxY {
			maxY = ra[3]
			ranges[2] = []int{ra[0], ra[2]}
			lengths[2] = ra[2] - ra[0]
		} else if ra[3] == maxY {
			if ra[0] < ranges[2][0] {
				ranges[2][0] = ra[0]
			}
			if ra[2] > ranges[2][1] {
				ranges[2][1] = ra[2]
			}
			lengths[2] += ra[2] - ra[0]
		}

		// update pointSet
		pointStrs := []string{
			fmt.Sprintf("%d_%d", ra[0], ra[1]),
			fmt.Sprintf("%d_%d", ra[0], ra[3]),
			fmt.Sprintf("%d_%d", ra[2], ra[1]),
			fmt.Sprintf("%d_%d", ra[2], ra[3]),
		}
		for _, pointStr := range pointStrs {
			_, ok := pointSet[pointStr]
			if ok {
				delete(pointSet, pointStr)
			} else {
				pointSet[pointStr] = true
			}
		}

		// add size
		sz += (ra[2] - ra[0]) * (ra[3] - ra[1])
	}
	if sz != (maxX - minX) * (maxY - minY) || len(pointSet) != 4 {
		return false
	}
	dx, dy := ranges[0][1] - ranges[0][0], ranges[1][1] - ranges[1][0]
	if dx != ranges[2][1] - ranges[2][0] || dy != ranges[3][1] - ranges[3][0] {
		return false
	}
	return dx == lengths[0] && dx == lengths[2] && dy == lengths[1] && dy == lengths[3]
}