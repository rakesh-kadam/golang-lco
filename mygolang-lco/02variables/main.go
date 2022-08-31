package main

import "fmt"

const LoginToken = "JFSFSDFSD" //Public

func main()  {
	var username string = "Rakesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n",isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T \n",smallVal)

	var smallFloat float64 = 255.45532423422534545353423
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n",smallVal)

	// Some default values and some alias
	var someValue int 
	fmt.Println(someValue)
	fmt.Printf("Variable is of type %T \n",someValue)
	
	var someValueString string 
	fmt.Println(someValueString)
	fmt.Printf("Variable is of type %T \n",someValueString)

	// implicite type
	var website = "rakeshkadam.me"
	fmt.Println(website)

	// no var
	numOfUsers := 3000
	fmt.Println(numOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n",LoginToken)

}