package main

import "fmt"

func main() {
	fmt.Println("Welcome  to functions in golang")
	greeter()

	result := addder(3, 6)
	proResult, str := proAdder(2, 459, 484, 8)
	fmt.Println(str)
	fmt.Println("Result is ", result)
	fmt.Println("ProResult is ", proResult)
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi pro result function"
}

func addder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Namstey from golang")
}
