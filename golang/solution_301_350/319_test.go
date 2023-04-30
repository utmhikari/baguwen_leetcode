package solution_301_350


//319. 灯泡开关
//
//初始时有 n 个灯泡处于关闭状态。
//
//对某个灯泡切换开关意味着：如果灯泡状态为关闭，那该灯泡就会被开启；而灯泡状态为开启，那该灯泡就会被关闭。
//
//第 1 轮，每个灯泡切换一次开关。即，打开所有的灯泡。
//
//第 2 轮，每两个灯泡切换一次开关。 即，每两个灯泡关闭一个。
//
//第 3 轮，每三个灯泡切换一次开关。
//
//第 i 轮，每 i 个灯泡切换一次开关。 而第 n 轮，你只切换最后一个灯泡的开关。
//
//找出 n 轮后有多少个亮着的灯泡。


func bulbSwitch319(n int) int {
	cnt := 0
	i := 1
	for {
		p := i * i
		if p > n {
			break
		}
		cnt++
		i++
	}
	return cnt
}