package main

import "fmt"

func main() {
	// topFactor := factorial(10)
	// fmt.Println(topFactor)

	numbers := []int{1, 10, 15}
	sum := sumup(3, 21, 19, 45, 25)
	anotherSum := sumup(1, numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val
	}

	return sum
}

// func factorial(number int) int {
// 	if number <= 0 {
// 		return 1
// 	}
// 	return number * factorial(number-1)

// 	// result := 1

// 	// for i := 1; i <= number; i++ {
// 	// 	result *= i
// 	// }

// 	// return result
// }
