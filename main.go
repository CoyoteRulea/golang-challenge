package main

import "fmt"

func main() {
	fmt.Println("Tests for Excercise 1")
	fmt.Println("=====================")
	testArray := []int{0, 5, 6, 0, 7, 6, 5}
	fmt.Printf("showUniquevalue Input: %v - Output: %d\n", testArray, showUniquevalue(testArray))
	testArray = []int{0, 5, 6, 7, 7, 6, 5}
	fmt.Printf("showUniquevalue Input: %v - Output: %d\n", testArray, showUniquevalue(testArray))
	testArray = []int{0, 7, 6, 0, 7, 6, 5}
	fmt.Printf("showUniquevalue Input: %v - Output: %d\n", testArray, showUniquevalue(testArray))
	testArray = []int{0, 5, 7, 0, 7, 6, 5}
	fmt.Printf("showUniquevalue Input: %v - Output: %d\n", testArray, showUniquevalue(testArray))
	testArray = []int{0, 5, 8, 0, 1, 9, 0}
	fmt.Println()
	fmt.Println("Tests for Excercise 1")
	fmt.Println("=====================")
	fmt.Printf("getOrderedArray Input: %v - Output: %v\n", testArray, getOrderedArray(testArray))
	testArray = []int{2, 7, 0, 5, 0, 0, 1, 9, 0, 4}
	fmt.Printf("getOrderedArray Input: %v - Output: %v\n", testArray, getOrderedArray(testArray))
	testArray = []int{0, 0, 0, 1, 2, 5, 6}
	fmt.Printf("getOrderedArray Input: %v - Output: %v\n", testArray, getOrderedArray(testArray))
	testArray = []int{9, 5, 6, 7, 0, 0, 0, 0}
	fmt.Printf("getOrderedArray Input: %v - Output: %v\n", testArray, getOrderedArray(testArray))
}
