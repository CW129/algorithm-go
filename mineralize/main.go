package main

import (
	"fmt"
	"sort"
)

// import (
// 	"math"
// 	"sort"
// )

func main() {
	picks := []int{1, 3, 2}
	// picks := []int{0, 1, 1}
	minerals := []string{"diamond", "diamond", "diamond", "iron", "iron", "diamond", "iron", "stone"}
	// minerals := []string{"diamond", "diamond", "diamond", "diamond", "diamond", "iron", "iron", "iron", "iron", "iron", "diamond"}
	answer := solution(picks, minerals)
	fmt.Println(answer)
}

// // 입출력 예
// // picks	minerals	result
// // [1, 3, 2]	["diamond", "diamond", "diamond", "iron", "iron", "diamond", "iron", "stone"]	12
// // [0, 1, 1]	["diamond", "diamond", "diamond", "diamond", "diamond", "iron", "iron", "iron", "iron", "iron", "diamond"]	50

// // 사용할 곡괭이가 없거나 광산에 있는 모든 광물을 캘 때까지

// func mining() int {
// 	return 0
// }

// var dia = 25
// var iron = 4
// var stone = 0

// type SortM struct {
// 	weight int
// 	pos    int
// }

// func solution(picks []int, minerals []string) int {
// 	// 곡괭이 개수
// 	pickCount := 0
// 	tmp := [][]string{}
// 	for _, s := range picks {
// 		pickCount += s
// 	}
// 	if len(minerals) > pickCount*5 {
// 		minerals = minerals[:(pickCount * 5)]
// 	}
// 	sortMineral := make(map[int]int)
// 	j := 1
// 	conv := 0
// 	for i, s := range minerals {
// 		if i == j*5 {
// 			j++
// 		}
// 		if s == "diamond" {
// 			conv = dia
// 		} else if s == "iron" {
// 			conv = iron
// 		} else if s == "stone" {
// 			conv = stone
// 		}
// 		sortMineral[j] = sortMineral[j] + conv
// 	}

// 	var sortM []SortM
// 	for k, v := range sortMineral {
// 		sortM = append(sortM, SortM{v, k})
// 	}
// 	sort.Slice(sortM, func(i, j int) bool {
// 		return sortM[i].weight > sortM[j].weight
// 	})

// 	for i := 0; i < int(math.Ceil(float64(len(minerals)/5))); i++ {
// 		tmp = append(tmp, minerals[i*5:i*5+5])
// 	}

// 	return 0
// }

// 정답
type Fatigue struct {
	diamond int
	iron    int
	stone   int
}

func NewFatigue(mineral string) *Fatigue {
	switch mineral {
	case "diamond":
		return &Fatigue{
			diamond: 1,
			iron:    5,
			stone:   25,
		}
	case "iron":
		return &Fatigue{
			diamond: 1,
			iron:    1,
			stone:   5,
		}
	case "stone":
		return &Fatigue{
			diamond: 1,
			iron:    1,
			stone:   1,
		}
	default:
		panic(mineral)
	}
}

func (f *Fatigue) Add(other *Fatigue) *Fatigue {
	return &Fatigue{
		diamond: f.diamond + other.diamond,
		iron:    f.iron + other.iron,
		stone:   f.stone + other.stone,
	}
}

func GetFatigues(minerals []string) []*Fatigue {
	if len(minerals) == 0 {
		return nil
	}

	var ret []*Fatigue
	for i, mineral := range minerals {
		if i%5 == 0 {
			ret = append(ret, &Fatigue{})
		}
		window := ret[len(ret)-1]
		fatigue := NewFatigue(mineral)
		ret[len(ret)-1] = window.Add(fatigue)
	}
	if ret[len(ret)-1].diamond == 0 {
		ret = ret[:len(ret)-1]
	}
	return ret
}

const (
	DiaPick = iota + 1
	IronPick
	StonePick
)

func solution(pickCounts []int, minerals []string) int {
	fatigues := GetFatigues(minerals)
	picks := GetPicks(pickCounts)
	min := Min(len(fatigues), len(picks))
	fatigues = fatigues[:min]
	picks = picks[:min]

	var ret int
	for len(picks) > 0 {
		pick := picks[len(picks)-1]
		picks = picks[:len(picks)-1]
		switch pick {
		case DiaPick:
			sort.Slice(fatigues, func(i, j int) bool {
				return fatigues[i].diamond < fatigues[j].diamond
			})
			ret += fatigues[0].diamond
		case IronPick:
			sort.Slice(fatigues, func(i, j int) bool {
				return fatigues[i].iron < fatigues[j].iron
			})
			ret += fatigues[0].iron
		case StonePick:
			sort.Slice(fatigues, func(i, j int) bool {
				return fatigues[i].stone < fatigues[j].stone
			})
			ret += fatigues[0].stone
		}
		fatigues = fatigues[1:]
	}

	return ret
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetPicks(counts []int) []int {
	var ret []int
	for pick, count := range counts {
		ret = AppendPick(ret, pick+1, count)
	}
	return ret
}

func AppendPick(picks []int, pick int, count int) []int {
	for i := 0; i < count; i++ {
		picks = append(picks, pick)
	}
	return picks
}
