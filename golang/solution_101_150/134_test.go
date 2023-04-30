package solution_101_150


//
//134. 加油站
//在一条环路上有N个加油站，其中第i个加油站有汽油gas[i]升。
//
//你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1个加油站需要消耗汽油cost[i]升。
//你从其中的一个加油站出发，开始时油箱为空。
//
//如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。


func canCompleteCircuit134(gas []int, cost []int) int {
	l := len(gas)
	if l == 0 {
		return -1
	}

	starts := make([]int, 0)
	sumGas, sumCost := 0, 0
	for i := 0; i < l; i++ {
		sumGas += gas[i]
		sumCost += cost[i]
		if gas[i] >= cost[i] {
			starts = append(starts, i)
		}
	}
	if sumGas < sumCost {
		return -1
	}

	ls := len(starts)
	if ls == 0 {
		return -1
	}

	i := 0
	for i < ls {
		start := starts[i]
		fuel := 0
		x := 0
		for x < l {
			fuel += gas[(start + x) % l] - cost[(start + x) % l]
			if fuel < 0 {
				break
			}
			x++
		}
		if x == l {
			return start
		}
		if ((start + x) % l) < start {
			return -1
		}
 		for i < ls && starts[i] <= start + x {
			i++
		}
	}

	return -1
}