package dp_test

import (
	"reflect"
	"testing"

	"github.com/klotzandrew/epi/dp"
)

func TestNumCombinationsForScore(t *testing.T) {
	for _, tt := range []struct {
		score      int
		playScores []int
		expected   int
	}{
		{12, []int{2, 3, 7}, 4},
	} {
		if result := dp.NumCombinationsForScore(tt.score, tt.playScores); !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("NumCombinationsForScore(%v, %v) = %v, expected %v", tt.score, tt.playScores, result, tt.expected)
		}
	}
}
