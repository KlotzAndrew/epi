package hashtable_test

import (
	"testing"

	"github.com/klotzandrew/epi/hashtable"
)

func TestIsLetterConstructable(t *testing.T) {
	for _, tt := range []struct {
		letter, mag string
		expected    bool
	}{
		{"def", "fed", true},
		{"deff", "fed", false},
		{"def", "ffed", true},
		{"zz", "fed", false},
	} {
		if result := hashtable.IsLetterConstructable(tt.letter, tt.mag); result != tt.expected {
			t.Errorf("IsLetterConstructable(%v, %v) = %v, expected %v", tt.letter, tt.expected, result, tt.expected)
		}
	}
}
