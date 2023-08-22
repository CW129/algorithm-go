package main

import (
	"strings"
)

// 입출력 예
// maps	result
// ["SOOOL","XXXXO","OOOOO","OXXXX","OOOOE"]	16
// ["LOOXS","OOOOX","OOOOO","OOOOO","EOOOO"]	-1

func main() {
	// maps := []string{"SOOOL", "XXXXO", "OOOOO", "OXXXX", "OOOOE"}
	// maps := []string{"XXSXX", "LOOOO", "XXXXX", "XXXXX", "XXXXE"}
	maps := []string{"LOOXS", "OOOOX", "OOOOO", "OOOOO", "EOOOO"}
	solution(maps)
}

type Pos struct {
	y int
	x int
}

func Bfs(sP Pos, dP Pos, maze [][]string) int {
	queue := make(chan Pos, 12)
	queue <- sP

	visited := make([][]int, len(maze))
	for i := range visited {
		visited[i] = make([]int, len(maze[0]))
	}

	visited[sP.y][sP.x] = 1
	defer close(queue)

	for {
		current := <-queue
		if current.y == dP.y && current.x == dP.x {
			return visited[current.y][current.x] - 1
		}
		if 0 < current.y && maze[current.y-1][current.x] != "X" && visited[current.y-1][current.x] == 0 {
			pos := &Pos{current.y - 1, current.x}
			queue <- *pos
			visited[current.y-1][current.x] = visited[current.y][current.x] + 1
		}

		if current.y < len(maze)-1 && maze[current.y+1][current.x] != "X" && visited[current.y+1][current.x] == 0 {
			pos := &Pos{current.y + 1, current.x}
			queue <- *pos
			visited[current.y+1][current.x] = visited[current.y][current.x] + 1
		}

		if 0 < current.x && maze[current.y][current.x-1] != "X" && visited[current.y][current.x-1] == 0 {
			// left
			pos := &Pos{current.y, current.x - 1}
			queue <- *pos
			visited[current.y][current.x-1] = visited[current.y][current.x] + 1
		}
		if current.x < len(maze[0])-1 && maze[current.y][current.x+1] != "X" && visited[current.y][current.x+1] == 0 {
			// right
			pos := &Pos{current.y, current.x + 1}
			queue <- *pos
			visited[current.y][current.x+1] = visited[current.y][current.x] + 1
		}

		if len(queue) == 0 { // 길 없을 시
			return -1
		}
	}
}

func solution(maps []string) int {
	maze := [][]string{}
	sP := &Pos{}
	lP := &Pos{}
	eP := &Pos{}

	for i, s := range maps {
		split := strings.Split(s, "")
		maze = append(maze, split)

		if strings.Contains(s, "S") {
			sP = &Pos{y: i, x: strings.Index(s, "S")}
		}
		if strings.Contains(s, "L") {
			lP = &Pos{y: i, x: strings.Index(s, "L")}
		}
		if strings.Contains(s, "E") {
			eP = &Pos{y: i, x: strings.Index(s, "E")}
		}
	}

	toLever := Bfs(*sP, *lP, maze)
	if toLever == -1 {
		return -1
	}
	// fmt.Println(sP)
	toEnd := Bfs(*lP, *eP, maze)
	if toEnd == -1 {
		return -1
	}
	answer := toLever + toEnd
	return answer
}
