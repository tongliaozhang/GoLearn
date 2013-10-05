package collection

import (
	"fmt"
)

func Slice_c1(s []int) {
	fmt.Println(s)
	s[1] = 100
}

func Slice_Create() {
	s := []int{1, 2, 3}
	Slice_c1(s)

	fmt.Println(s)

	n := make([]int, 5, 10)

	fmt.Println(n, len(n), cap(n))
}

func Slice_append() {
	s := make([]int, 5, 10)
	fmt.Println(s, len(s), cap(s))
	s1 := append(s, 1, 2, 3, 4, 5, 6)
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s1)
}
