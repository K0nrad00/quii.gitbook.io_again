package main

func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum = sum + i
	}
	return sum
}

// I was close to the below function but not getting in entirely
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
