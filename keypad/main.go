package main

import (
	"fmt"
	"math"
)

// [1, 3, 4, 5, 8, 2, 1, 4, 5, 9, 5]	"right"	"LRLLLRLLRRL"
// [7, 0, 8, 2, 8, 3, 1, 5, 7, 6, 2]	"left"	"LRLLRRLLLRR"
// [1, 2, 3, 4, 5, 6, 7, 8, 9, 0]	"right"	"LLRLLRLLRL"

func main() {
	hand := "left"
	number := [...]int{7, 0, 8, 2, 8, 3, 1, 5, 7, 6, 2}

	fmt.Println(solution(number[:], hand))

}

func leftPos(pos int, s int) int {
	count := (((s - pos) % 3) + ((s - pos) / 3))

	return int(math.Abs(float64(count)))
}

func rightPos(pos int, s int) int {
	count := ((pos - s) % 3) + ((pos - s) / 3)
	return int(math.Abs(float64(count)))
}

func solution(number []int, hand string) string {
	a := ""

	lp := 11
	rp := 12
	if hand == "left" {
		hand = "L"
	} else {
		hand = "R"
	}

	for _, s := range number {
		if s == 0 {
			s = 11
		}
		if lp == s || s == 1 || s == 4 || s == 7 {
			a += "L"
			lp = s
		} else if rp == s || s == 3 || s == 6 || s == 9 {
			a += "R"
			rp = s
		} else {
			lCount := leftPos(lp, s)
			rCount := rightPos(rp, s)

			if lCount == rCount {
				a += hand
			}
			if lCount < rCount {
				a += "L"
			} else {
				a += "R"
			}

		}
	}
	return a
}
