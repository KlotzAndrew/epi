package strings_test

import (
	"testing"

	"github.com/klotzandrew/epi/strings"
)

func TestIntToString(t *testing.T) {
	for _, tt := range []struct {
		s string
		v int64
	}{
		{"123", 123},
		{"-123", -123},
		{"9223372036854775807", 1<<63 - 1},
		{"1", 1},
		{"0", 0},
	} {
		if result := strings.StringToInt(tt.s); result != tt.v {
			t.Errorf("StringToInt(%s) = %d, %d", tt.s, result, tt.v)
		}

		if result := strings.IntToString(tt.v); result != tt.s {
			t.Errorf("IntToString(%d) = %s, %s", tt.v, result, tt.s)
		}
	}
}
