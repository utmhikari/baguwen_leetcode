package solution_151_200




//171. Excel表列序号
//给定一个Excel表格中的列名称，返回其相应的列序号。


func titleToNumber171(columnTitle string) int {
	sm := 0
	for _, r := range []rune(columnTitle) {
		sm = sm * 26 + int(r - 'A') + 1
	}
	return sm
}
