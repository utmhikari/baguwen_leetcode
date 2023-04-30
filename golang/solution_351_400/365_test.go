package solution_351_400

import "fmt"

//365. 水壶问题

//有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，
//从而可以得到恰好 z升 的水？
//
//如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。


var states365 map[string]bool
var target365 int
var c1365 int
var c2365 int


func key365(x1, x2 int) string {
	return fmt.Sprintf("%d_%d", x1, x2)
}


func solve365(x1, x2 int) bool {
	// reach target?
	if x1 == target365 || x2 == target365 || x1 + x2 == target365 {
		return true
	}

	// visited?
	key := key365(x1, x2)
	_, visited := states365[key]
	if visited {
		return false
	}
	states365[key] = true

	// fill
	if solve365(c1365, x2) {
		return true
	}
	if solve365(x1, c2365) {
		return true
	}

	// empty
	if solve365(0, x2) {
		return true
	}
	if solve365(x1, 0) {
		return true
	}

	// x1 -> x2
	if x1 != 0 && x2 != c2365 {
		x1ToX2 := x1
		x2ToC2 := c2365 - x2
		if x1ToX2 > x2ToC2 {
			x1ToX2 = x2ToC2
		}
		if solve365(x1 - x1ToX2, x2 + x1ToX2) {
			return true
		}
	}

	// x2 -> x1
	if x2 != 0 && x1 != c1365 {
		x2ToX1 := x2
		x1ToC1 := c1365 - x1
		if x2ToX1 > x1ToC1 {
			x2ToX1 = x1ToC1
		}
		if solve365(x1 + x2ToX1, x2 - x2ToX1) {
			return true
		}
	}

	return false
}


func canMeasureWater365Dfs(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if targetCapacity == 0 ||
		jug1Capacity == targetCapacity ||
		jug2Capacity == targetCapacity ||
		jug1Capacity + jug2Capacity == targetCapacity {
		return true
	}

	if jug1Capacity + jug2Capacity < targetCapacity {
		return false
	}

	if jug1Capacity % 2 == 0 && jug2Capacity % 2 == 0 && targetCapacity % 2 != 0 {
		return false
	}

	target365 = targetCapacity
	states365 = make(map[string]bool)
	c1365 = jug1Capacity
	c2365 = jug2Capacity
	return solve365(0, 0)
}

func gcd365(a, b int) int {
	if a == 1 || b == 1 {
		return 1
	}
	if a < b {
		return gcd365(b, a)
	}
	m := a % b
	if m == 0 {
		return b
	}
	return gcd365(m, b)
}

func canMeasureWater365(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if jug1Capacity + jug2Capacity < targetCapacity {
		return false
	}
	if jug1Capacity == 0 || jug2Capacity == 0 {
		return targetCapacity == 0 || jug1Capacity + jug2Capacity == targetCapacity
	}
	return targetCapacity % gcd365(jug1Capacity, jug2Capacity) == 0
}