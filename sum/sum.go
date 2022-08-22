package sum

func Sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}

	return result
}

func SumAll(numbersArr ...[]int) (sums []int) {
	for _, numbers := range numbersArr {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersArr ...[]int) (sums []int) {
	for _, numbers := range numbersArr {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
