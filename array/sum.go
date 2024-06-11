package main

func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) []int {
	// create a array of size "args"
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	// iterate and substitute the index with sum
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
func SumAllTails(numbersToSum ...[]int) []int {
	// create a array of size "args"

	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func main() {
	SumAllTails([]int{1, 2, 3}, []int{2, 3, 4})
}
