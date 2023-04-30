package solution_51_100

import "testing"


type state65 struct {
	signed bool
	hasE bool
	hasDot bool
	hasNum bool
}

func isNumber65(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}

	state := state65{}

	for i := 0; i < l; i++ {
		c := s[i:i+1]
		switch {
		case c == "+" || c == "-":
			if state.signed {
				return false
			}
			state.signed = true
			break
		case c >= "0" && c <= "9":
			state.signed = true
			state.hasNum = true
			break
		case c == ".":
			if state.hasE || state.hasDot || l == 1 {
				return false
			}
			state.signed = true
			state.hasDot = true
			break
		case c == "E" || c == "e":
			if state.hasE || i == 0 || i == l - 1 || !state.hasNum {
				return false
			}
			state.hasE = true
			state.signed = false
			state.hasNum = false
			break
		default:
			return false
		}
	}

	return state.hasNum
}


type test65 struct {
	s string
	expected bool
}


func Test_65(t *testing.T) {
	positives := []string{
		"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10",
		"-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789",
	}
	negatives := []string{
		"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53",
		".", ".e1", "4e+", "+e4",
	}

	inputs := make([]test65, 0)
	for _, pos := range positives {
		inputs = append(inputs, test65{pos, true})
	}
	for _, neg := range negatives {
		inputs = append(inputs, test65{neg, false})
	}

	for i, input := range inputs {
		ans := isNumber65(input.s)
		if ans != input.expected {
			t.Errorf("failed at %d -> %+v, got %v\n", i, input, ans)
		} else {
			t.Logf("success at %d -> %+v\n", i, input)
		}
	}

}