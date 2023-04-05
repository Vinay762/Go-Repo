package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Array in golang")
	var fruitList [4]string
	fruitList[0] = "Mango"
	fruitList[1] = "Apple"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is ", fruitList)
	fmt.Println("Fruit list is ", len(fruitList))

	var vegList = [3]string{"Potato", "beans", "mushrooms"}
	fmt.Println(vegList)
}
