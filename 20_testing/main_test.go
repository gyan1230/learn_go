package main

import "testing"

func TestMysum(t *testing.T) {
	type test struct {
		data   []int
		result int
	}

	tests := []test{
		test{[]int{1, 2, 3, 4, 5}, 15},
		test{[]int{12, 30}, 42},
		test{[]int{-1, 0, 12}, 11},
	}

	for _, v := range tests {
		x := mysum(v.data...)
		if x != v.result {
			t.Error("Expected ", v.result, "Got", x)
		}
	}
}
