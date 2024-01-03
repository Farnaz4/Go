package main

import "fmt"

//Structucts are like pyhton dictionaries
//a collection of key value pairs



// type messageToSend struct {
// 	phoneNumber int
// 	message string

	
// }

// //do not edit below this line

// func test(m messageToSend) {

// 	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
// 	fmt.Println("=========================")
// }

// func main(){

// 	test(messageToSend {
// 		phoneNumber: 104255510981
// 		message: "Thanks for signing up",
// 	})
// 	test(messageToSend {
// 		phoneNumber: 10151616681
// 		message: "Love to have you abroad!",
// 	})
// 	test(messageToSend {
// 		phoneNumber: 104444410981
// 		message: "We are so excited to have you here",
// 	})



// }

//=============================Lesson 2==================

type messageToSend struct {

	message string
	sender user
	recipeint user

}

type user struct {
	name string
	number int
}

func canSendMessage(mToSend, messageToSend) bool{
	//?
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipeint.name == "" {
		return false
	}
	if mToSend.recipeint.number == 0 {
		return false
	}
	return true

}