package main

import "fmt"

// 입출력 예
// queue1	queue2	result
// [3, 2, 7, 2]	[4, 6, 5, 1]	2
// [1, 2, 1, 2]	[1, 10, 1, 2]	7
// [1, 1]	[1, 5]	-1

func SumQueue(queue []int) int {
	sum := 0
	for _, s := range queue {
		sum += s
	}
	return sum
}

func Dequeue(queue []int) (int, []int) {
	data := queue[0]
	queue = queue[1:]
	return data, queue
}

func Enqueue(queue []int, data int) []int {
	queue = append(queue, data)
	return queue
}

func main() {
	queue1 := []int{3, 2, 7, 2}
	queue2 := []int{4, 6, 5, 1}
	Test := solution(queue1, queue2)
	fmt.Println(Test)
}

func solution(queue1 []int, queue2 []int) int {
	count := 0
	flag := len(queue1)
	var data int
	Sum1 := SumQueue(queue1)
	Sum2 := SumQueue(queue2)
	total := Sum1 + Sum2
	if total%2 == 1 {
		return -1
	}

	for {
		if Sum1 == Sum2 {
			return count
		}
		if Sum1 < Sum2 {
			data, queue2 = Dequeue(queue2)
			Sum2 = Sum2 - data
			Sum1 = Sum1 + data
			queue1 = append(queue1, data)
		} else {
			data, queue1 = Dequeue(queue1)
			Sum1 = Sum1 - data
			Sum2 = Sum2 + data
			queue2 = append(queue2, data)
		}
		count++
		if count > flag*3 {
			return -1
		}
	}
}
