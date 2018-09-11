package heap

// VerifyHeap verify if an array of integer satisfies
// the heap property, which parent node is larger or equal to its child
func VerifyHeap(heap []int) bool {
	return verifyHeapNode(heap, 0)
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

}
