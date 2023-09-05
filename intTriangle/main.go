package main

import "fmt"

// [[7], [3, 8], [8, 1, 0], [2, 7, 4, 4], [4, 5, 2, 6, 5]]	30
func main() {
	triangle := [][]int{{7}, {3, 8}, {8, 1, 0}, {2, 7, 4, 4}, {4, 5, 2, 6, 5}}
	answer := solution(triangle)
	fmt.Println(answer)
}

func solution(triangle [][]int) int {
	// num := make(map[int]int)
	i := len(triangle) - 1
	for {
		for j, s := range triangle[i] {
			if j < len(triangle[i])-1 {
				if s >= triangle[i][j+1] {
					triangle[i-1][j] = triangle[i-1][j] + s
				} else if s < triangle[i][j+1] {
					triangle[i-1][j] = triangle[i-1][j] + triangle[i][j+1]
				}
			}
		}
		i--
		if i == 0 {
			return triangle[0][0]
		}
	}
}
