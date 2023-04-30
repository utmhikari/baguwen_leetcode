package solution_1_50

import (
	"bytes"
	"fmt"
)

//38. 外观数列
//给定一个正整数 n ，输出外观数列的第 n 项。
//
//「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。
//
//你可以将其视作是由递归公式定义的数字字符串序列：
//
//countAndSay(1) = "1"
//countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。
//前五项如下：
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//第一项是数字 1
//描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
//描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
//描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
//描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"



func countAndSay38(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay38(n - 1)
	cc, cnt := prev[0], 1
	var b bytes.Buffer
	for i := 1; i < len(prev); i++ {
		if prev[i] == cc {
			cnt++
		} else {
			b.WriteString(fmt.Sprintf("%d", cnt))
			b.WriteRune(rune(cc))
			cc, cnt = prev[i], 1
		}
	}
	b.WriteString(fmt.Sprintf("%d", cnt))
	b.WriteRune(rune(cc))
	return b.String()
}