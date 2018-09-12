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
		{[]int{}, true},
		{[]int{5}, true},
		{[]int{10, 5}, true},
		{[]int{5, 10}, false},
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

func TestHeapify(t *testing.T) {
	tables := [][]int{
		[]int{},
		[]int{5},
		[]int{10, 20},
		[]int{10, 20, 15},
		[]int{4, 7, 8, 3, 2, 6, 5},
		[]int{1, 4, 3, 7, 8, 9, 10},
	}

	for _, test := range tables {
		original := make([]int, len(test))
		copy(test, original)
		heap.Heapify(test)

		t.Logf("Result %v", test)

		if !heap.VerifyHeap(test) {
			t.Errorf("Unable to heapify %v", original)
		}
	}
}
