package main

import "fmt"

// slice is passed as reference
func tryModifySlice(s []int) {
	s[0] = 10
}

func main() {
	s := []int{1, 2, 3}
	tryModifySlice(s)

	fmt.Printf("s[0] = %v\n", s[0])
}
