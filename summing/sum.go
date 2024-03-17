package main

func Sum(nums []int) int {
	var sum int = 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for index, numbers := range numbersToSum {
		sums[index] = Sum(numbers)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}