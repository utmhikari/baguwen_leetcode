package solution_301_350

//312. 戳气球
//有 n 个气球，编号为0 到 n - 1，
//每个气球上都标有一个数字，这些数字存在数组  nums  中。
//
//现在要求你戳破所有的气球。
//戳破第 i 个气球，你可以获得  nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。
//  这里的 i - 1 和 i + 1 代表和  i  相邻的两个气球的序号。
//如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
//
//求所能获得硬币的最大数量。


//var ansCache312 map[string]int
//
//func hashNums312(nums []int) string {
//	var b bytes.Buffer
//	for _, n := range nums {
//		b.WriteRune(rune(n))
//		b.WriteRune(',')
//	}
//	hs := md5.New()
//	hs.Write([]byte(b.String()))
//	return hex.EncodeToString(hs.Sum(nil))
//}
//
//
//func getProduct312(nums []int, i int) int {
//	left, right := 1, 1
//	if i > 0 {
//		left = nums[i - 1]
//	}
//	if i < len(nums) - 1 {
//		right = nums[i + 1]
//	}
//	return left * right * nums[i]
//}
//
//
//func solve312(nums []int) int {
//	key := hashNums312(nums)
//	ans, ok := ansCache312[key]
//	if ok {
//		return ans
//	}
//
//	l := len(nums)
//	if l == 1 {
//		// no need to cache ans with len 1
//		return getProduct312(nums, 0)
//	}
//
//	mx := 0
//	for i := 0; i < l; i++ {
//		product := getProduct312(nums, i)
//
//		// next arr
//		next := make([]int, l - 1)
//		j, k := 0, 0
//		for j < l - 1 {
//			if k == i {
//				k++
//				continue
//			}
//			next[j] = nums[k]
//			j++
//			k++
//		}
//
//		// get next ans, cal max val
//		nextAns := solve312(next)
//		sum := product + nextAns
//		if sum > mx {
//			mx = sum
//		}
//	}
//
//	ansCache312[key] = mx
//	return mx
//}


func maxCoins312(nums []int) int {
	//ansCache312 = make(map[string]int)
	//return solve312(nums)



	// SOLUTION:
	// 定义算法solve(i, j)
	// 表示在开区间i, j编号内部填满气球能够得到的最多硬币数
	// 计算方法，需要在开区间i, j取一个mid，然后如下求和
	// sum = val[i] * val[j] * val[mid] + solve(i, mid) + solve(mid, j)
	// 如果i >= j - 1，即开区间里没有编号，就取0
	// 这个算法可以转换成dp的形式

	l := len(nums)
	dp := make([][]int, l + 2)
	for i := 0; i < l + 2; i++ {
		dp[i] = make([]int, l + 2)
	}
	val := make([]int, l + 2)

	val[0] = 1
	val[l + 1] = 1
	for i := 1; i <= l; i++ {
		val[i] = nums[i - 1]
	}

	for i := l - 1; i >= 0; i-- {
		for j := i + 2; j <= l + 1; j++ {
			for k := i + 1; k < j; k++ {
				sm := val[i] * val[k] * val[j]
				sm += dp[i][k] + dp[k][j]
				if sm > dp[i][j] {
					dp[i][j] = sm
				}
			}
		}
	}

	return dp[0][l + 1]
}