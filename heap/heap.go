package heap

// VerifyHeap verify if an array of integer satisfies
// the heap property, which parent node is larger or equal to its child
func VerifyHeap(heap []int) bool {
	return verifyHeapNode(heap, 0)
}

// Sort sorts an array using HeapSort algorithm
func Sort(heap []int) {
	for size := len(heap); size > 1; size-- {
		heapify(heap, size)

		// move the first to the last
		heap[size-1], heap[0] = heap[0], heap[size-1]
	}
}

func verifyHeapNode(heap []int, node int) bool {
	leftChild := node * 2
	rightChild := node*2 + 1
	size := len(heap)

	if leftChild < size && heap[node] < heap[leftChild] {
		return false
	}

	if rightChild < size && heap[node] < heap[rightChild] {
		return false
	}

	return true
}

// Heapify converts an array to ensure it satisfies the heap property
// this function performs inplace edit of the given array
func Heapify(heap []int) {
	heapify(heap, len(heap))
}

func heapify(heap []int, size int) {
	// convert all subtrees of the heaps
	for i := len(heap) / 2; i >= 0; i-- {
		heapifySubtree(heap, i, size)
	}
}

func heapifySubtree(heap []int, node int, size int) {
	left := 2*node + 1
	right := 2*node + 2

	largest := node

	if left < size && heap[left] > heap[largest] {
		largest = left
	}

	if right < size && heap[right] > heap[largest] {
		largest = right
	}

	// move down
	if largest != node {
		heap[largest], heap[node] = heap[node], heap[largest]
		heapifySubtree(heap, largest, size)
	}
}
