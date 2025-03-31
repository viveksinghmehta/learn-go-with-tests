package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	var sum []int

	for _, slice := range numbersToSum {
		sum = append(sum, Sum(slice))
	}
	return sum
}

func SumAllTails(tailsToSum ...[]int) []int {
	var sum []int
	for _, slice := range tailsToSum {
		if len(slice) == 0 {
			sum = append(sum, 0)
		} else {
			tails := slice[1:]
			sum = append(sum, Sum(tails))
		}
	}
	return sum
}
