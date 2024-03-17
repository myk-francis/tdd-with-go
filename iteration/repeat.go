package main

import "fmt"

func Repeat(character string, numberOfIterations int) string {
	var repeated string
	
	for i := 0; i < numberOfIterations; i++ {
		repeated += character
	}
	return repeated
}

func main() {
	var numOfIteration int

	fmt.Print("How many times should a be repeate?:")
	_,err := fmt.Scan(&numOfIteration)
	if err != nil {
		fmt.Println("Something went wrong!!")
		return 
	}

	fmt.Println(Repeat("a", 5))
}
