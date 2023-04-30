package solution_151_200

import "math"

func maximumGap164(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}

	minVal, maxVal := math.MaxInt32, math.MinInt32
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	d := (maxVal - minVal) / (l - 1)
	if d < 1 {
		d = 1
	}
	bucketSize := (maxVal - minVal) / d + 1

	buckets := make([][]int, bucketSize)
	for i := 0; i < bucketSize; i++ {
		buckets[i] = make([]int, 2)
		buckets[i][0] = -1
	}
	for _, num := range nums {
		bid := (num - minVal) / d
		if buckets[bid][0] == -1 {
			buckets[bid][0] = num
			buckets[bid][1] = num
		} else {
			if num < buckets[bid][0] {
				buckets[bid][0] = num
			}
			if num > buckets[bid][1] {
				buckets[bid][1] = num
			}
		}
 	}

 	prev := -1
 	ans := 0
	for i, bucket := range buckets {
		if bucket[0] == -1 {
			continue
		}

		if prev != -1 {
			mxOffset := bucket[0] - buckets[prev][1]
			if mxOffset > ans {
				ans = mxOffset
			}
		}
		prev = i
	}
	return ans
}