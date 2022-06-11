package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const MONITORAMENTOS = 2
const DELAY =  5

func main() {

	for { // infinite loop
		menu()
		command := readCommand()

		// if command == 1 {
		// 	fmt.Println("Starting monitoring...")
		// } else if command == 2 {
		// 	fmt.Println("Showing logs...")
		// } else if command == 3 {
		// 	fmt.Println("Exiting...")
		// } else {
		// 	fmt.Printf("%d: Invalid command!", &command)
		// }

		switch command {
		case 1:
			monitoringWebsite()
		case 2:
			fmt.Println("Showing logs...")
		case 3:
			greetings(true)
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0) // expected exit
		default:
			fmt.Println("Invalid command!")
			os.Exit(-1) // unexpected exit
		}
	}
}

func greetings(show_type bool) {
	label := "Start" // same thing that var label string = "Start"
	var name string
	var age int
	var height float32

	fmt.Println("######", label, "######")
	fmt.Println("Hi! Please, what's your name:")
	fmt.Scan(&name)

	fmt.Println("Hi", name)
	fmt.Println("Please, how old are you?")

	fmt.Scan(&age)
	fmt.Println("Oh really?")

	fmt.Println("And what's your height, please?")
	fmt.Scan(&height)

	if show_type {
		fmt.Println("\nType Name: ", reflect.TypeOf(name))
		fmt.Println("Type Age: ", reflect.TypeOf(age))
		fmt.Println("Type Height: ", reflect.TypeOf(height))
	}
}

func menu() {
	fmt.Println("\n#### MENU ####")
	fmt.Println("1) Start monitoring.")
	fmt.Println("2) Show logs.")
	fmt.Println("3) Greetings")
	fmt.Println("0) Exit")

}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func monitoringWebsite() {
	fmt.Println("Starting monitoring...")

	// websites := [3]string
	websites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
	}

	websites = append(websites, "https://www.udemy.com")

	for i := 0; i < MONITORAMENTOS; i++{
		// for url := 0; url < len(websites); url ++{
		for i, url := range websites {
			fmt.Println("\nChecking website with ID number", i)
			testingWebSite(url)
		}
		time.Sleep(DELAY * time.Second)
	}
	

}

func testingWebSite(site string) {
	req, _ := http.Get(site)

	if req.StatusCode == 200 {
		fmt.Println("The website", site, "is ok")
	} else {
		fmt.Println("The website", site, "is out. Status Code:", req.StatusCode)
	}
}
