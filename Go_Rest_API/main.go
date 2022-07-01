package main

import (
	"Go_Rest_API/routes"
	"fmt"
)

func main() {
	fmt.Println("Starting Server ...")
	routes.HandleRequest()
}
