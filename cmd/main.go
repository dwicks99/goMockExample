package main

import (
	"fmt"
	"scratch/mockExample/userservice"
)

func main() {

	fmt.Println("...listening for request...")

	// Handle Subsequent requests
	userservice.HandleRequests()
}
