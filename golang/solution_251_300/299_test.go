package solution_251_300

import "fmt"

//299. 猜数字游戏


func getHint(secret string, guess string) string {
	l := len(secret)
	if l != len(guess) {
		panic("invalid guess")
	}

	idxCache := make([]bool, l)
	numA := 0

	// A
	for i := 0; i < l; i++ {
		bs, bg := secret[i], guess[i]
		if bs == bg {
			numA++
			idxCache[i] = true
		}
	}

	// B
	numB := 0
	sCtr, gCtr := make(map[byte]int), make(map[byte]int)  // counters
	for i := 0; i < l; i++ {
		if idxCache[i] {
			continue
		}
		sCnt, ok := sCtr[secret[i]]
		if ok {
			sCtr[secret[i]] = sCnt + 1
		} else {
			sCtr[secret[i]] = 1
		}
		gCnt, ok := gCtr[guess[i]]
		if ok {
			gCtr[guess[i]] = gCnt + 1
		} else {
			gCtr[guess[i]] = 1
		}
	}

	for b, sCnt := range sCtr {
		gCnt, ok := gCtr[b]
		if ok {
			m := sCnt
			if gCnt < sCnt {
				m = gCnt
			}
			numB += m
		}
	}

	return fmt.Sprintf("%dA%dB", numA, numB)
}