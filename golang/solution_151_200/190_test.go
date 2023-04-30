package solution_151_200

func reverseBits190(num uint32) uint32 {
	var n uint32 = 0
	var b uint32 =1
	for i := 0; i < 31; i++ {
		if (num & b) == b {
			n += 1 << (31 - i)
		}
		b <<= 1
	}
	if (num & b) == b {
		n += 1
	}
	return n
}
