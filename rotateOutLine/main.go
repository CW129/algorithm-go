package main

// rows	columns	queries	result
// 6	6	[[2,2,5,4],[3,3,6,6],[5,1,6,3]]	[8, 10, 25]
// 3	3	[[1,1,2,2],[1,2,2,3],[2,1,3,2],[2,2,3,3]]	[1, 1, 5, 3]
// 100	97	[[1,1,100,97]]	[1]
func main() {
	rows := 3
	columns := 3
	queries := [][]int{{1, 1, 2, 2}, {1, 2, 2, 3}, {2, 1, 3, 2}, {2, 2, 3, 3}}
	// queries := [][]int{{1, 1, 2, 2}, {1, 2, 2, 3}}

	solution(rows, columns, queries)
}

func minValue(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func solution(rows int, columns int, queries [][]int) []int {
	count := 1
	table := make([][]int, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			table[i] = append(table[i], count)
			count++
		}
	}
	count = 0
	answer := []int{}

	for i := range queries {
		x1, y1 := queries[i][0]-1, queries[i][1]-1
		x2, y2 := queries[i][2]-1, queries[i][3]-1

		tmp := []int{table[x1][y1], table[x1][y2], table[x2][y2]}
		min := tmp[0]
		min = minValue(min, tmp[1])
		min = minValue(min, tmp[2])

		for x := x1 + 1; x <= x2; x++ {
			table[x-1][y1] = table[x][y1]
			min = minValue(min, table[x-1][y1])
		}
		for y := y2 - 1; y >= y1; y-- {
			table[x1][y+1] = table[x1][y]
			min = minValue(min, table[x1][y+1])
		}
		table[x1][y1+1] = tmp[0]
		for x := x2 - 1; x >= x1; x-- {
			table[x+1][y2] = table[x][y2]
			min = minValue(min, table[x+1][y2])
		}
		table[x1+1][y2] = tmp[1]
		for y := y1 + 1; y <= y2; y++ {
			table[x2][y-1] = table[x2][y]
			min = minValue(min, table[x2][y-1])
		}
		table[x2][y2-1] = tmp[2]
		answer = append(answer, min)
	}

	return answer

}
