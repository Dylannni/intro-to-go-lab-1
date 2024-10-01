package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	return append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i, num := range slice {
		slice[i] = f(num)
	}
}

func mapArray(f func(a int) int, array *[3]int) {
	for i, num := range array {
		array[i] = f(num)
	}
}

func main() {
	// Slice is a reference type
	// modify the slice inside the function will affect the original slice.
	intsSlice := []int{1, 2, 3}
	mapSlice(addOne, intsSlice)
	fmt.Println("Modified slice:", intsSlice)

	// Array is a value type
	// Creates a copy when passed into a function.
	intsArray := [3]int{1, 2, 3}
	mapArray(addOne, &intsArray)
	fmt.Println("Modified slice:", intsArray)

	oriSlice := []int{2, 3, 4, 5, 6}
	newSlice := intsSlice[1:3]
	mapSlice(square, newSlice)

	fmt.Println("Original slice:", oriSlice)
	fmt.Println("New slice:", newSlice)

	intsSlice = double(intsSlice)
	fmt.Println(intsSlice)
}
