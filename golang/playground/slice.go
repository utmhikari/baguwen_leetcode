package playground

import (
	"fmt"
)


func TestSlice() {
	s := make([]string, 3)
	s[0] = "1"
	s[1] = "haha"
	s[2] = "i9"
	fmt.Println("emp:", s)

	s = append(s, "d")
	s = append(s, "e", "f")

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l1 := s[2:5]
	fmt.Println("sl1:", l1)

	l2 := s[:5]
	fmt.Println("sl2:", l2)

	l3 := s[2:]
	fmt.Println("sl3:", l3)

	l3[0] = "99999"
	fmt.Println("now_s:", s)  // sublist


}


