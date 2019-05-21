package string_processing

import "testing"
import "reflect"

func Test(t *testing.T) {
	var tests = []struct {
		s string
		z []int
	}{
		{"aabcaabxaaaz", []int{0, 1, 0, 0, 3, 1, 0, 0, 2, 2, 1, 0}},
		{"abab#ababab", []int{0, 0, 2, 0, 0, 4, 0, 4, 0, 2, 0}},
		{"", []int{}},
	}
	for _, c := range tests {
		got := Z_function(c.s)
		if reflect.DeepEqual(got, c.z) == false {
			t.Errorf("Z(%q) == %v, want %v", c.s, got, c.z)
		}
	}
}
