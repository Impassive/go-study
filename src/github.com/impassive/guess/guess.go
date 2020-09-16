package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please think of a number between 1 and 100")
	fmt.Println("Press Enter when ready")
	scanner.Scan()

	// input size 100
	// 100/2 guesses before answer
	// n = 1 to 100 O(n)

	// guess := 50
	// for {
	// 	fmt.Println("I guess the number is", guess)
	// 	fmt.Println("Is that:")
	// 	fmt.Println("(a) too high?")
	// 	fmt.Println("(b) too low?")
	// 	fmt.Println("(c) correct?")
	// 	scanner.Scan()
	// 	response := scanner.Text()

	// 	if response == "a" {
	// 		guess--
	// 	} else if response == "b" {
	// 		guess++
	// 	} else if response == "c" {
	// 		fmt.Println("I won!")
	// 		break
	// 	} else {
	// 		fmt.Println("Invelid response, try again.")
	// 	}
	// }

	// binary search

	low := 1
	high := 100

	for {
		guess := (low + high) / 2
		fmt.Println("I guess the number is", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I won!")
			break
		} else {
			fmt.Println("Invelid response, try again.")
		}
	}
}
