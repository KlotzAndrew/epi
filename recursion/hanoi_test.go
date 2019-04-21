package recursion_test

import (
	"reflect"
	"testing"

	"github.com/klotzandrew/epi/recursion"
)

func TestHanoiSteps(t *testing.T) {
	for _, tt := range []struct {
		pegs  int
		steps []recursion.HanoiStep
	}{
		{3, []recursion.HanoiStep{
			recursion.HanoiStep{From: 0, To: 1},
			recursion.HanoiStep{From: 0, To: 2},
			recursion.HanoiStep{From: 1, To: 2},
			recursion.HanoiStep{From: 0, To: 1},
			recursion.HanoiStep{From: 2, To: 0},
			recursion.HanoiStep{From: 2, To: 1},
			recursion.HanoiStep{From: 0, To: 1},
		}},
	} {
		if result := recursion.HanoiSteps(tt.pegs); !reflect.DeepEqual(result, tt.steps) {
			t.Errorf("HanoiSteps(%v) = %v, expected %v", tt.pegs, result, tt.steps)
		}
	}
}
