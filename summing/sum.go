package main

func Sum(nums [5]int) int {
	var sumOfNumbers int = 0

	for i := 0; i < len(nums); i++ {
		sumOfNumbers += nums[i]
	}

	return sumOfNumbers
}