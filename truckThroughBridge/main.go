package main

import (
	"fmt"
)

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

func (q *Queue) FirstData() int {
	if q.IsEmpty() {
		return 10000
	}
	return (*q)[0].(int)
}

func sumWeight(bridge []int) int {
	sum := 0
	for i := 1; i < len(bridge); i++ {
		sum += bridge[i]
	}
	return sum
}

// 입출력 예
// bridge_length	weight	truck_weights	return
// 2	10	[7,4,5,6]	8
// 100	100	[10]	101
// 100	100	[10,10,10,10,10,10,10,10,10,10]	110

func main() {
	bridge_length := 100
	weight := 100
	truck_weights := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	answer := solution(bridge_length, weight, truck_weights)
	fmt.Println(answer)
}

func solution(bridge_length int, weight int, truck_weights []int) int {
	queue := &Queue{}
	bridge := make([]int, bridge_length)
	flag := true
	count := 0

	for _, s := range truck_weights {
		queue.Enqueue(s)
	}

	for flag {
		if queue.FirstData()+sumWeight(bridge) <= weight {
			bridge = bridge[1:]
			bridge = append(bridge, queue.Dequeue().(int))
			if sumWeight(bridge) != 0 {
				count++
			}
		} else {
			bridge = bridge[1:]
			bridge = append(bridge, 0)
			count++
		}
		if sumWeight(bridge) == 0 && queue.IsEmpty() == true {
			return count + 1
		}
	}
	fmt.Println(count)
	return count
}
