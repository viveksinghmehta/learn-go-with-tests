package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	lengthOfSlice := len(numbersToSum)
	sum := make([]int, lengthOfSlice)

	for i, slice := range numbersToSum {
		sum[i] = Sum(slice)
	}
	return sum
}
