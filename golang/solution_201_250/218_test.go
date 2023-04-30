package solution_201_250



//218. 天际线问题
//城市的天际线是从远处观看该城市中所有建筑物形成的轮廓。
//给你所有建筑物的位置和高度，请返回由这些建筑物形成的 天际线 。
//
//每个建筑物的几何信息由数组 buildings 表示，
//其中三元组 buildings[i] = [lefti, righti, heighti] 表示：
//
//lefti 是第 i 座建筑物左边缘的 x 坐标。
//righti 是第 i 座建筑物右边缘的 x 坐标。
//heighti 是第 i 座建筑物的高度。
//天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，
//并按 x 坐标 进行 排序 。关键点是水平线段的左端点。列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0 ，仅用于标记天际线的终点。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。
//
//注意：输出天际线中不得有连续的相同高度的水平线。
//例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答案；
//三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]


var ans218 [][]int


func updateOutput218(output *[][]int, x int, y int) {
	if len(*output) == 0 || (*output)[len(*output)-1][0] != x {
		*output = append(*output, []int{x, y})
	} else {
		(*output)[len(*output)-1][1] = y
	}
}

func appendSkyline218(output *[][]int, skyline *[][]int, p int, n int, curY int) {
	for p < n {
		point := (*skyline)[p]
		x, y := point[0], point[1]
		p++
		if curY != y {
			updateOutput218(output, x, y)
			curY = y
		}
	}
}


func mergeSkylines218(leftSkylines [][]int, rightSkylines [][]int) [][]int {
	nl, nr := len(leftSkylines), len(rightSkylines)
	pl, pr := 0, 0
	curY, leftY, rightY := 0, 0, 0
	output := make([][]int, 0)
	var x int

	for pl < nl && pr < nr {
		pointl, pointr := leftSkylines[pl], rightSkylines[pr]
		if pointl[0] < pointr[0] {
			x = pointl[0]
			leftY = pointl[1]
			pl++
		} else {
			x = pointr[0]
			rightY = pointr[1]
			pr++
		}
		maxY := leftY
		if rightY > maxY {
			maxY = rightY
		}
		if curY != maxY {
			updateOutput218(&output, x, maxY)
			curY = maxY
		}
	}

	appendSkyline218(&output, &leftSkylines, pl, nl, curY)
	appendSkyline218(&output, &rightSkylines, pr, nr, curY)

	return output
}

func getSkyline218(buildings [][]int) [][]int {
	l := len(buildings)
	if l == 0 {
		return  make([][]int, 0)
	}
	if l == 1 {
		return [][]int{
			{buildings[0][0], buildings[0][2]},
			{buildings[0][1], 0},
		}
	}

	left := getSkyline218(buildings[:l / 2])
	right := getSkyline218(buildings[l / 2:])
	return mergeSkylines218(left, right)
}