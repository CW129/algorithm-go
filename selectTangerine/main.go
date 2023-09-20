package main

import (
	"sort"
)

// 입출력 예
// k	tangerine	result
// 6	[1, 3, 2, 5, 4, 5, 2, 3]	3
// 4	[1, 3, 2, 5, 4, 5, 2, 3]	2
// 2	[1, 1, 1, 1, 2, 2, 2, 3]	1
func main() {
	k := 6
	tangerine := []int{1, 3, 2, 5, 4, 5, 2, 3}
	solution(k, tangerine)
}

func solution(k int, tangerine []int) int {
	table := make(map[int]int)
	sortSlice := []int{}
	answer := 0
	for _, s := range tangerine {
		table[s]++
	}
	for _, s := range table {
		sortSlice = append(sortSlice, s)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sortSlice)))

	for _, s := range sortSlice {
		k -= s
		answer++
		if k <= 0 {
			break
		}
	}
	return answer
}
