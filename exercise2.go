package main

func getOrderedArray(numbers []int) []int {
	arrLen := len(numbers)
	var preArr, postArr []int

	for x := 0; x < arrLen; x++ {
		if numbers[x] == 0 {
			preArr = append(preArr, 0)
		} else {
			postArr = append(postArr, numbers[x])
		}
	}

	return append(preArr, postArr...)
}
