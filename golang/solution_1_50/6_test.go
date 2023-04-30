package solution_1_50

import "bytes"

//6. Z 字形变换
//将一个给定字符串 s 根据给定的行数 numRows ，
//以从上往下、从左到右进行Z 字形排列。
//
//比如输入字符串为 "PAYPALISHIRING"行数为 3 时，排列如下：
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，
//比如："PAHNAPLSIIGYIR"。
//
//请你实现这个将字符串进行指定行数变换的函数：


func convert6(s string, numRows int) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	step := 2 * (numRows - 1)

	var b bytes.Buffer
	for i := 0; i < numRows; i++ {
		switch i {
		case 0:
			for j := i; j < l; j += step {
				b.WriteString(s[j:j+1])
			}
			break
		case numRows - 1:
			for j := i; j < l; j += step {
				b.WriteString(s[j:j+1])
			}
			break
		default:
			x, y := i, step - i
			for x < l || y < l {
				if x < l {
					b.WriteString(s[x:x+1])
				}
				if y < l {
					b.WriteString(s[y:y+1])
				}
				x += step
				y += step
			}
			break
		}
	}

	return b.String()
}
