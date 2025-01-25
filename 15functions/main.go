package main

import "fmt"

func main() {
	fmt.Println("Welcom to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro message: ", myMessage)
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Namastey from golang")
}
