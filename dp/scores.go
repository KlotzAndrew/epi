package dp

import "fmt"

// NumCombinationsForScore .
func NumCombinationsForScore(score int, playScores []int) int {
	numCombos := [][]int{}
	for i := 0; i < len(playScores); i++ {
		numCombos = append(numCombos, make([]int, score+1))
	}

	for i, v := range playScores {
		numCombos[i][0] = 1

		for j := 1; j <= score; j++ {
			without, with := 0, 0
			if i-1 >= 0 {
				without = numCombos[i-1][j]
			}

			if j >= v {
				with = numCombos[i][j-v]
				fmt.Println("j > v", j, v, with)
			}

			numCombos[i][j] = without + with
		}
	}
	return numCombos[len(playScores)-1][score]
}
