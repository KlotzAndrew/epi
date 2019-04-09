package ptypes_test

import (
	"testing"

	"github.com/klotzandrew/epi/ptypes"
)

func TestReverseDigits(t *testing.T) {
	for _, tt := range []struct {
		x, expected int64
	}{
		{12, 21},
		{123, 321},
		{101, 101},
		{100, 1},
		{-13, -31},
	} {
		if result := ptypes.ReverseDigits(tt.x); result != tt.expected {
			t.Errorf("ReverseDigits(%v) = %v, expected %v", tt.x, result, tt.expected)
		}
	}
}
