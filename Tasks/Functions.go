package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

// don't touch below this line

// func main1() {
// 	fmt.Println(concat("Lane,", "Happy Birthday!"))
// 	fmt.Println(concat("Elon", "hope that your Tesla sales go up"))
// 	fmt.Println(concat("What is", "Go and Golang?"))
// }


//MULTIPLE PARAMETERS
// When Multiple arguments are of the same type, the type onlt needs to be declared after the last one, assuming they are in order.
// Ex: func createUser(firstname, lastName string, age int)


//_______________________________________________________

//Assignment

func main() {

	sendsSoFar := 430
	const sendsToAdd =25
	sendsSoFar= incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
	

}

func incrementSends(sendsSoFar, sendsToAdd int) int {

	sendsSoFar = sendsSoFar + sendsToAdd

	return sendsSoFar

}




//Ignoring return values
//We can explicitly use _ to ignore variables
