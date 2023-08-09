package main

// N	stages	result
// 5	[2, 1, 2, 6, 2, 4, 3, 3]	[3,4,2,1,5]
// 4	[4,4,4,4,4]	[4,1,2,3]

func main() {
	N := 5
	stages := []int{2, 1, 2, 6, 2, 4, 3, 3}
	solution(N, stages)
}

func solution(N int, stages []int) []int {
	answer := []int{}
	stg := make(map[int]float32)
	for i := 1; i <= N+1; i++ {
		stg[i] = 0
		if i != N+1 {
			answer = append(answer, i)
		}
	}

	// []int len: 8, cap: 8, [2,1,2,6,2,4,3,3]
	for _, s := range stages {
		if s != N+1 {
			stg[s] = stg[s] + 1
		}
	}

	// map[int]float32 [1: 0.125, 2: 0.42857143, 3: 0.5, 4: 0.5, 5: 0, 6: 0, ]
	user := float32(len(stages))
	for i := 1; i < len(stg); i++ {
		tmp := stg[i]
		stg[i] = stg[i] / float32(user)
		user = user - tmp
		if i == N+1 {
			stg[i] = 0
		}
	}

	// insertion sort
	for i := 1; i < len(answer); i++ {
		j := i
		for j > 0 {
			if stg[answer[j-1]] < stg[answer[j]] {
				answer[j-1], answer[j] = answer[j], answer[j-1]
			}
			j--
		}
	}

	return answer

}
