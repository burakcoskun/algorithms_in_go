package tree

import "testing"

func Test(t *testing.T) {
	var bit [11]int
	for idx := 0; idx < 10; idx++ {
		update(idx, idx, bit[:])
	}
	for idx := 0; idx < 10; idx++ {
		res := Query_range(idx, idx, bit[:])
		if res != idx {
			t.Errorf("%v == %v", res, idx)
		}
	}

	res := Query_range(0, 9, bit[:])
	if res != 45 {
		t.Errorf("%v == %v", res, 45)
	}
}
