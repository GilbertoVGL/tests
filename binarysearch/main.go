package main

import "fmt"

var iterations int

func main() {
	a := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	target := 1
	fmt.Println("target:", target)
	ok := search(a, target)

	fmt.Println("found my number?", ok)
	fmt.Printf("took %d iterations to find it\n", iterations)
}

func search(source []int, target int) bool {
	iterations++
	fmt.Println("iterations:", iterations)

	if len(source) == 0 {
		return false
	}

	mid := (len(source) - 1) / 2
	fmt.Println("len(source):", len(source))
	fmt.Println("mid:", mid)
	fmt.Println("source[mid]:", source[mid])
	if source[mid] == target {
		return true
	}

	if target > source[mid] {
		return search(source[mid+1:], target)
	}
	return search(source[:mid], target)
}
