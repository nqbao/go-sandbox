package heap_test

import (
	"testing"

	"github.com/nqbao/learn-go/heap"
)

func TestVerifyHeap(t *testing.T) {
	tables := []struct {
		input  []int
		result bool
	}{
		{[]int{}, true},
		{[]int{5}, true},
		{[]int{3, 2, 1}, true},
		{[]int{3, 2, 1, 0}, true},
		{[]int{6, 4, 5, 3, 2, 1}, true},
		{[]int{8, 10, 5, 3, 2, 1}, false},
	}

	for _, test := range tables {
		if heap.VerifyHeap(test.input) != test.result {
			t.Errorf("Expect heap %v to be %v", test.input, test.result)
		}
	}
}
