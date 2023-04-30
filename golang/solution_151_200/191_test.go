package solution_151_200






func hammingWeight191(num uint32) int {
	var b uint32 = 1
	cnt := 0
	for i := 0; i < 31; i++ {
		if (num & b) == b {
			cnt++
		}
		b <<= 1
	}
	if (num & b) == b {
		cnt++
	}
	return cnt
}