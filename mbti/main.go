package main

// ["AN", "CF", "MJ", "RT", "NA"]	[5, 3, 2, 7, 5]	"TCMA"
// ["TR", "RT", "TR"]	[7, 1, 3]	"RCJA"

func main() {
	survey := [...]string{"TR", "RT", "TR"}
	choices := [...]int{7, 1, 3}
	solution(survey[:], choices[:])
}

func solution(survey []string, choices []int) string {
	var count = [...]int{3, 2, 1, 0, 1, 2, 3}
	personalCode := map[string]int{
		"R": 0, "T": 0,
		"C": 0, "F": 0,
		"J": 0, "M": 0,
		"A": 0, "N": 0,
	}
	for i, s := range survey {
		if choices[i] < 4 {
			personalCode[string(s[0])] += count[choices[i]-1]
		} else {
			personalCode[string(s[1])] += count[choices[i]-1]
		}
	}
	a := ""
	if personalCode["R"] >= personalCode["T"] {
		a += "R"
	} else {
		a += "T"
	}
	if personalCode["C"] >= personalCode["F"] {
		a += "C"
	} else {
		a += "F"
	}
	if personalCode["J"] >= personalCode["M"] {
		a += "J"
	} else {
		a += "M"
	}
	if personalCode["A"] >= personalCode["N"] {
		a += "A"
	} else {
		a += "N"
	}
	return a
}
