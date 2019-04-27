package graphs_test

import (
	"reflect"
	"testing"

	"github.com/klotzandrew/epi/graphs"
)

func TestSearchMaze(t *testing.T) {
	for _, tt := range []struct {
		maze       [][]bool
		start, end graphs.Cordinate
		path       []graphs.Cordinate
	}{
		{
			[][]bool{
				{true, false, false},
				{true, true, true},
				{false, false, true},
			},
			graphs.Cordinate{X: 0, Y: 0},
			graphs.Cordinate{X: 2, Y: 2},
			[]graphs.Cordinate{
				graphs.Cordinate{X: 0, Y: 0},
				graphs.Cordinate{X: 1, Y: 0},
				graphs.Cordinate{X: 1, Y: 1},
				graphs.Cordinate{X: 1, Y: 2},
			},
		},
	} {
		if result := graphs.SearchMaze(tt.maze, tt.start, tt.end); !reflect.DeepEqual(result, tt.path) {
			t.Errorf("SearchMaze(%v) = %v, expected %v", tt.maze, result, tt.path)
		}
	}
}
