package solution_301_350


//315. 计算右侧小于当前元素的个数
//
//
//给定一个整数数组 nums，
//按要求返回一个新数组counts。
//数组 counts 有该性质： counts[i] 的值是nums[i] 右侧小于nums[i] 的元素的数量。


// [5,2,6,1,3,2,4,0]

// e.g. 6
// 1: 1
// 3: 2 + 1
// 4: 0 + 1


// ========================= MergeSort ===============================

var MSIndices []int
var MSTmp []int
var MSTmpIndices []int
var MSAns []int

func merge315(a *[]int, l int, mid int, r int) {
	i, j, p := l, mid + 1, l
	for i <= mid && j <= r {
		if (*a)[i] <= (*a)[j] {
			MSTmp[p] = (*a)[i]
			MSTmpIndices[p] = MSIndices[i]
			MSAns[MSIndices[i]] += j - mid - 1
			i++
		} else {
			MSTmp[p] = (*a)[j]
			MSTmpIndices[p] = MSIndices[j]
			j++
		}
		p++
	}

	for i <= mid {
		MSTmp[p] = (*a)[i]
		MSTmpIndices[p] = MSIndices[i]
		MSAns[MSIndices[i]] += j - mid - 1
		i++
		p++
	}

	for j <= r {
		MSTmp[p] = (*a)[j]
		MSTmpIndices[p] = MSIndices[j]
		j++
		p++
	}

	for k := l; k <= r; k++ {
		MSIndices[k] = MSTmpIndices[k]
		(*a)[k] = MSTmp[k]
	}
}


func MS315(a *[]int, l int, r int) {
	if l >= r {
		return
	}

	mid := (l + r) >> 1
	MS315(a, l, mid)
	MS315(a, mid + 1, r)
	merge315(a, l, mid, r)
}


func solveMS315(nums []int) []int {
	length := len(nums)
	MSIndices = make([]int, length)
	MSTmp = make([]int, length)
	MSTmpIndices = make([]int, length)
	MSAns = make([]int, length)

	for i := 0; i < length; i++ {
		MSIndices[i] = i
	}

	l, r := 0, length - 1
	MS315(&nums, l, r)
	return MSAns
}


func countSmaller(nums []int) []int {
	return solveMS315(nums)
}