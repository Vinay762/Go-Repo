package main

import "fmt"

//jwtToken := 3000; :- not allowed outside of the method

//constants
const LoginToken string = "sljasllslslfajl"

func main() {
	var username string = "Vinay"
	fmt.Println(username)
	fmt.Printf("Varibles is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type: %T \n", isLoggedIn)

	var smallVal int16 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type: %T \n", smallVal)

	var smallFloat float32 = 539.583
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variables is of type: %T \n", anotherVariable)

	//implicit type
	var website = "google.com"
	fmt.Println(website)

	//no var style
	numberofUser := 3000
	fmt.Println(numberofUser)

}
