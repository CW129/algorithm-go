package main

func main() {
	// 변수
	M := 4
	N := 10
	Solution(N, M)
}

func Solution(N int, M int) int {
	arr := make([]int, N)
	i := 0
	count := 1
	arr[0] = 1
	for {
		i += M
		if i > N {
			i -= N
		}
		if i == N || arr[i] == 1 {
			break
		}
		arr[i] = 1
		count += 1
	}
	return count
}
