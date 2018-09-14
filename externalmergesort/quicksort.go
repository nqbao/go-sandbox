package externalmergesort

// perform inline quicksort
// I choose quicksort because it does not require extra memory and it is generally faster than
// heapsort because it does not need to swap the element around.
// see https://medium.com/@k2u4yt/quicksort-vs-heapsort-3b6dc5395083
func quicksort(input []int32, size int) {
	quicksortFrom(input, 0, size-1)
}

func quicksortFrom(input []int32, left int, right int) {
	if left < right {
		p := quicksortParition(input, left, right)
		quicksortFrom(input, left, p-1)
		quicksortFrom(input, p+1, right)
	}
}

func quicksortParition(input []int32, left int, right int) int {
	start := left
	value := input[left]
	left = left + 1

	for {
		// find the left mark
		for {
			if left <= right && input[left] <= value {
				left = left + 1
			} else {
				break
			}
		}

		// find the right mark
		for {
			if right >= left && input[right] >= value {
				right = right - 1
			} else {
				break
			}
		}

		// swap
		if right < left {
			break
		} else {
			input[left], input[right] = input[right], input[left]
		}
	}

	// swap the right mark
	input[start], input[right] = input[right], input[start]

	return right
}
