package main

type Pos struct {
	y int
	x int
}

type Queue []interface{}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(data interface{}) {
	*q = append(*q, data)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}

func numIslands(grid [][]byte) int {
	num_islands := 0
	m := len(grid)
	n := len(grid[0])
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}

	bfs := func(y int, x int) {
		queue := &Queue{}
		pos := &Pos{y, x}
		visited[y][x] = 1
		queue.Enqueue(pos)
		for {
			if queue.IsEmpty() == true {
				return
			}
			tmp := queue.Dequeue().(*Pos)

			if tmp.y > 0 && grid[tmp.y-1][tmp.x] == '1' && visited[tmp.y-1][tmp.x] != 1 { //up
				pos = &Pos{tmp.y - 1, tmp.x}
				queue.Enqueue(pos)
				visited[tmp.y-1][tmp.x] = 1
			}
			if tmp.y < m-1 && grid[tmp.y+1][tmp.x] == '1' && visited[tmp.y+1][tmp.x] != 1 { //down
				pos = &Pos{tmp.y + 1, tmp.x}
				queue.Enqueue(pos)
				visited[tmp.y+1][tmp.x] = 1
			}
			if tmp.x > 0 && grid[tmp.y][tmp.x-1] == '1' && visited[tmp.y][tmp.x-1] != 1 { //left
				pos = &Pos{tmp.y, tmp.x - 1}
				queue.Enqueue(pos)
				visited[tmp.y][tmp.x-1] = 1
			}
			if tmp.x < n-1 && grid[tmp.y][tmp.x+1] == '1' && visited[tmp.y][tmp.x+1] != 1 { //right
				pos = &Pos{tmp.y, tmp.x + 1}
				queue.Enqueue(pos)
				visited[tmp.y][tmp.x+1] = 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && visited[i][j] != 1 {
				bfs(i, j)
				num_islands++
			}
		}
	}

	return num_islands
}
