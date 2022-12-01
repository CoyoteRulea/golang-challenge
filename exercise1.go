package main

func showUniquevalue(numbers []int) int {
	arrLen := len(numbers)
	for x := 0; x < arrLen; x++ {
		duplicated := false

		for y := 0; y < arrLen; y++ {
			if y == x {
				continue
			}

			if numbers[x] == numbers[y] {
				duplicated = true
				break
			}
		}

		if !duplicated {
			return numbers[x]
		}
	}

	return -1
}
