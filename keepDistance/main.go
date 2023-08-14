package main

import (
	"fmt"
	"math"
	"strings"
)

// 입출력 예
// places
// [["POOOP", "OXXOX", "OPXPX", "OOXOX", "POXXP"],
//  ["POOPX", "OXPXP", "PXXXO", "OXXXO", "OOOPP"],
//  ["PXOPX", "OXOXP", "OXPOX", "OXXOP", "PXPOX"],
//  ["OOOXX", "XOOOX", "OOOXX", "OXOOX", "OOOOO"],
//  ["PXPXP", "XPXPX", "PXPXP", "XPXPX", "PXPXP"]]

// result
//  [1, 0, 1, 1, 1]

/*
"POOOP",
"OXXOX",
"OPXPX",
"OOXOX",
"POXXP"

"POOPX",
"OXPXP",
"PXXXO",
"OXXXO",
"OOOPP"
*/

func main() {
	places := [][]string{
		{"POOOP", "OXXOX", "OPXPX", "OOXOX", "POXXP"},
		{"POOPX", "OXPXP", "PXXXO", "OXXXO", "OOOPP"},
		{"PXOPX", "OXOXP", "OXPOX", "OXXOP", "PXPOX"},
		{"OOOXX", "XOOOX", "OOOXX", "OXOOX", "OOOOO"},
		{"PXPXP", "XPXPX", "PXPXP", "XPXPX", "PXPXP"}}

	fmt.Println(solution(places))
}

type Coordinate struct {
	x []int
	y []int
}

func solution(places [][]string) []int {
	answer := []int{1, 1, 1, 1, 1}
	// 최초 for문 -> 각 대기실별로 확인
	for a, room := range places {
		roomSlice := [][]string{}
		coordinate := Coordinate{}
		// 현재 대기실의 상황 2차원 배열로 재정리 및 P의 좌표값 추출
		for j, line := range room {
			s := strings.Split(line, "")
			for i, table := range s {
				if table == "P" {
					coordinate.x = append(coordinate.x, i)
					//[]int len: 6, cap: 8, [0,4,1,3,0,4]
					coordinate.y = append(coordinate.y, j)
					//[]int len: 6, cap: 8, [0,0,2,2,4,4]
				}
			}
			roomSlice = append(roomSlice, s)
		}
		//좌표값으로 맨허튼 거리 계산
	Loop:
		for i := range coordinate.x {
			for j := i + 1; j < len(coordinate.x); j++ {
				if (math.Abs(float64(coordinate.x[i])-float64(coordinate.x[j])) + math.Abs(float64(coordinate.y[i])-float64(coordinate.y[j]))) == 1 {
					answer[a] = 0
					break Loop
				}
				if (math.Abs(float64(coordinate.x[i])-float64(coordinate.x[j])) + math.Abs(float64(coordinate.y[i])-float64(coordinate.y[j]))) <= 2 {
					switch {
					case coordinate.y[i] == coordinate.y[j] && roomSlice[coordinate.y[i]][coordinate.x[i]+1] != "X":
						answer[a] = 0
						break Loop

					case coordinate.x[i] == coordinate.x[j] && roomSlice[coordinate.y[i]+1][coordinate.x[i]] != "X":
						answer[a] = 0
						break Loop

					case coordinate.x[i] < coordinate.x[j]:
						if roomSlice[coordinate.y[i]][coordinate.x[i]+1] != "X" {
							answer[a] = 0
							break Loop
						}
						if roomSlice[coordinate.y[j]][coordinate.x[j]-1] != "X" {
							answer[a] = 0
							break Loop
						}

					case coordinate.x[i] > coordinate.x[j]:
						if roomSlice[coordinate.y[i]][coordinate.x[i]-1] != "X" {
							answer[a] = 0
							break Loop
						}
						if roomSlice[coordinate.y[j]][coordinate.x[j]+1] != "X" {
							answer[a] = 0
							break Loop
						}
					}
				}
			}
		}
	}
	return answer
}
