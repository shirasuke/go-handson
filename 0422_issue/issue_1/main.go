package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 14, 21, 28}
	displayNumbers(numbers)
}

func displayNumbers(nums []int) {
	for _, num := range nums {
		switch {
		case num == 0:
			fmt.Println("Zero")
			return
		case num%7 == 0:
			fmt.Println("Lucky Number")
			return
		case num%2 == 0:
			fmt.Println("Even Number")
			return
		case num%2 == 1:
			fmt.Println("Odd Number")
			return
		default:
			fmt.Println("Error Value")
			return
		}
	}
}