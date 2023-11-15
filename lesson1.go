//Declaring a Variable in Go

package main

import "fmt"


func main(){

	//initialize variable here

	var smsSendingLimit int 
	var costPerSMS float64
	var hasPermission bool
	var username string


	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission,username)
}