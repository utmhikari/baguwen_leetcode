package solution_251_300


//278. 第一个错误的版本
//
//你是产品经理，目前正在带领一个团队开发新的产品。
//不幸的是，你的产品的最新版本没有通过质量检测。
//由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
//
//假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
//
//你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。
//实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。


func isBadVersion(n int) bool {
	return true
}


func solve278(left int, right int) int {
	if left == right {
		if isBadVersion(left) {
			return left
		} else {
			return 0
		}
	}
	if left > right {
		return 0
	}

	mid := left + (right - left) / 2
	if isBadVersion(mid) {
		leftAns := solve278(left, mid - 1)
		if leftAns > 0 {
			return leftAns
		} else {
			return mid
		}
	} else {
		return solve278(mid + 1, right)
	}
}


func firstBadVersion278(n int) int {
	return solve278(1, n)
}