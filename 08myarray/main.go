package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitlist)
	fmt.Println("Fruit list is: ", len(fruitlist))

	var vegList = [3]string{"potatoe", "beans", "mushroom"}

	fmt.Println("Vegy list is: ", vegList)
	fmt.Println("Vegy list is: ", len(vegList))
}
