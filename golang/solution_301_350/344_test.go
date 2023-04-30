package solution_301_350





func reverseString344(s []byte)  {
	i, j := 0, len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
