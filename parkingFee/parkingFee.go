package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 입출력 예
// fees	records	result
// [180, 5000, 10, 600]	["05:34 5961 IN", "06:00 0000 IN", "06:34 0000 OUT", "07:59 5961 OUT", "07:59 0148 IN", "18:59 0000 IN", "19:09 0148 OUT", "22:59 5961 IN", "23:00 5961 OUT"]	[14600, 34400, 5000]
// [120, 0, 60, 591]	["16:00 3961 IN","16:00 0202 IN","18:00 3961 OUT","18:00 0202 OUT","23:58 3961 IN"]	[0, 591]
// [1, 461, 1, 10]	["00:00 1234 IN"]	[14841]

func main() {
	fees := [4]int{180, 5000, 10, 600}
	records := []string{"05:34 5961 IN", "06:00 0000 IN", "06:34 0000 OUT", "07:59 5961 OUT", "07:59 0148 IN", "18:59 0000 IN", "19:09 0148 OUT", "22:59 5961 IN", "23:00 5961 OUT"}
	fmt.Println(solution(fees[:], records))
}

const (
	endtime = 23*60 + 59
)

func solution(fees []int, records []string) []int {
	useTime := make(map[string]int)
	answer := []int{}
	// 16:00 3961 IN","16:00 0202 IN"

	strToInt := func(s string) int {
		a, _ := strconv.Atoi(s)
		return a
	}

	splitRecords := func(s string) (string, string, int) {
		tmpdata := strings.Split(s, " ")
		timeCon := strings.Split(tmpdata[0], ":")
		minute := strToInt(timeCon[0])*60 + strToInt(timeCon[1])
		return tmpdata[1], tmpdata[2], minute
	}

	calcPrice := func(s int) int {
		price := fees[1]
		if 0 >= s-fees[0] {
			return price
		}
		s = s - fees[0]
		// test1 := float64(s) / float64(fees[2])
		// test2 := s % fees[2]
		// test3 := math.Ceil(test1)
		price += int(math.Ceil(float64(s)/float64(fees[2]))) * fees[3]
		// price += (s/fees[2] + s%fees[2]) * fees[3]
		return price
	}

	for i, s := range records {
		carNum, stat, startTime := splitRecords(s)

		if stat == "IN" {
			found := false
			for j := i + 1; j < len(records); j++ {
				if strings.Contains(records[j], carNum) {
					_, _, endTime := splitRecords(records[j])
					useTime[carNum] = useTime[carNum] + endTime - startTime
					found = true
					break
				}
			}
			if found == false {
				useTime[carNum] = useTime[carNum] + endtime - startTime
			}
		}
	}

	var names []string
	for k, _ := range useTime {
		names = append(names, k)
	}

	sort.Strings(names)

	for _, name := range names {
		answer = append(answer, calcPrice(useTime[name]))
	}

	return answer
}
