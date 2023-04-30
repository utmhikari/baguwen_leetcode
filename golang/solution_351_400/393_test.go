package solution_351_400


// 393. UTF-8 编码验证
//UTF-8 中的一个字符可能的长度为 1 到 4 字节，遵循以下的规则：
//
//对于 1 字节的字符，字节的第一位设为 0 ，后面 7 位为这个符号的 unicode 码。
//对于 n 字节的字符 (n > 1)，第一个字节的前 n 位都设为1，第 n+1 位设为 0 ，
//后面字节的前两位一律设为 10 。剩下的没有提及的二进制位，全部为这个符号的 unicode 码。
//这是 UTF-8 编码的工作方式：
//
//Char. number range  |        UTF-8 octet sequence
//(hexadecimal)    |              (binary)
//--------------------+---------------------------------------------
//0000 0000-0000 007F | 0xxxxxxx
//0000 0080-0000 07FF | 110xxxxx 10xxxxxx
//0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
//0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
//给定一个表示数据的整数数组，返回它是否为有效的 utf-8 编码。


func validUtf8(data []int) bool {
	i := 0
	for i < len(data) {
		// header
		switch {
		case data[i] < 128:
			i++
			break
		case data[i] >= 192 && data[i] < 224:
			if i + 1 >= len(data) {
				return false
			}
			if data[i + 1] < 128 || data[i + 1] >= 192 {
				return false
			}
			i += 2
			break
		case data[i] >= 224 && data[i] < 240:
			if i + 2 >= len(data) {
				return false
			}
			if data[i + 1] < 128 || data[i + 1] >= 192 {
				return false
			}
			if data[i + 2] < 128 || data[i + 2] >= 192 {
				return false
			}
			i += 3
		case data[i] >= 240 && data[i] < 248:
			if i + 3 >= len(data) {
				return false
			}
			if data[i + 1] < 128 || data[i + 1] >= 192 {
				return false
			}
			if data[i + 2] < 128 || data[i + 2] >= 192 {
				return false
			}
			if data[i + 3] < 128 || data[i + 3] >= 192 {
				return false
			}
			i += 4
		default:
			return false
		}
	}
	return true
}