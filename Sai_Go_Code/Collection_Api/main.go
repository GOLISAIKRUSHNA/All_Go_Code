package main

import (
	routes "collection/routes"
	"fmt"
)

func main() {

	fmt.Println("Api Start")

	router := routes.SetupRouter()
	// fmt.Println("Program executed")

	// Start the server on port 8080
	router.Run()

}
