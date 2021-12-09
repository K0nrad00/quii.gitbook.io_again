package main

func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum = sum + i
	}
	return sum
}
