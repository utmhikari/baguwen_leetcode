package playground

import (
	"fmt"
	"strconv"
)

func testStringToDigit() {
	fmt.Println("=============== test string to digit ==================")

	var s string = "1234567"
	var i int
	var i32 int32
	var ui64 uint64

	i, _ = strconv.Atoi(s)
	i64, _ := strconv.ParseInt(s, 10, 32)
	i32 = int32(i64)
	ui64, _ = strconv.ParseUint(s, 10, 64)
	fmt.Printf("string: %s, int: %d, int32:%d, uint64: %d\n",
		s, i, i32, ui64)

	var sf string = "1234512345.12345678901234"
	var f32 float32
	var f64 float64

	f64, _ = strconv.ParseFloat(sf, 64)
	f32In64, _ := strconv.ParseFloat(sf, 32)
	f32 = float32(f32In64)
	fmt.Printf("string: %s, float32: %f, float64: %f\n",
		sf, f32, f64)

	fmt.Println("====================================================")
}

func testDigitToString() {
	fmt.Println("=============== test digit to string ==================")

	var i int = 12345
	var i32 int32 = -678
	var ui64 uint64 = 901234567890
	var strInt string = fmt.Sprintf("%d", i)
	var strInt32 string = strconv.FormatInt(int64(i32), 10)
	var	strUInt64 string = strconv.FormatUint(ui64, 10)
	fmt.Printf("strInt: %s, strInt32: %s, strUInt64: %s\n",
		strInt, strInt32, strUInt64)

	var f32 float32 = 3.1415926
	var f64 float64 = -3.1415926535897
	var strFloat32 string = fmt.Sprintf("%.5f", f32)
	var strFloat64 string = strconv.FormatFloat(f64, 'f', 10, 64)
	fmt.Printf("strFloat32: %s, strFloat64: %s\n", strFloat32, strFloat64)

	fmt.Println("====================================================")
}

func testStringAndRuneAndByte() {
	fmt.Println("=============== test string & rune & byte ==================")

	var s string = "helloworld"
	var runes []rune = []rune(s)
	var bytes []byte = []byte(s)
	fmt.Printf("string: %s, runes: %v, bytes: %v\n", s, runes, bytes)

	for i, char := range s {
		if i >= 0 {
			fmt.Printf("type of index: %T\n", i) // int
			fmt.Printf("type of char: %T\n", char)  // int32
			fmt.Printf("digit value of char: %d\n", char) // 104
			fmt.Printf("output char: %c\n", char) // h
			fmt.Printf("is equal with rune value: %v\n", 'h' == rune(char)) // true
			fmt.Printf("is equal with byte value: %v\n", 'h' == byte(char)) // true
			// invalid expression -> rune: int32, byte: uint8
			// fmt.Printf("is rune & byte equals: %v\n", rune(char) == byte(char))
			break
		}
	}

	fmt.Println("====================================================")
}

func testIntAndRuneAndByte() {
	fmt.Println("=============== test int & rune & byte ==================")

	var i int
	var r rune
	var b byte
	var ui8 uint8

	i, r = 123456, 123456
	fmt.Printf("int is equal with rune int: %v\n", int32(i) == r) // true
	b, ui8 = 'a', 'a'
	fmt.Printf("byte char is equal with uint8 char: %v\n", b == ui8) // true
	i = int('a')
	r, b, ui8 = rune(i), byte(i), uint8(i)
	fmt.Printf("int: %c, rune: %c, byte: %c, uint8: %c\n", i, r, b, ui8) // 4 * 'a'

	fmt.Println("====================================================")
}


func testInterfaceToInt() {
	fmt.Println("=============== test interface to digit ==================")

	var itf interface{}
	itf = 1234567890123456789
	i := itf.(int)

	// cannot convert to int32/int64/float32/float64

	fmt.Printf("interface: %v, int: %d\n", itf, i)

	fmt.Println("====================================================")

}

func TestDataTypes() {
	testStringToDigit()
	testDigitToString()
	testStringAndRuneAndByte()
	testIntAndRuneAndByte()
	testInterfaceToInt()
}
