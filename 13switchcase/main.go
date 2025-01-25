package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and case in golang")

	// rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Dice value is a and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough
	case 4:
		fmt.Println("You can move to 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("What was that")
	}
}
