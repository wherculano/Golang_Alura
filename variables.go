package main

import ( 
	"fmt"
	"reflect"
	"os"
)


func main(){
	greetings(true)
	
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

	switch command{
	case 1:
		fmt.Println("Starting monitoring...")
	case 2:
		fmt.Println("Showing logs...")
	case 3:
		fmt.Println("Exiting...")
		os.Exit(0) // expected exit
	default:
		fmt.Println("Invalid command!")
		os.Exit(-1)  // unexpected exit
	}
}

func greetings(show_type bool){
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

	if show_type{
		fmt.Println("\nType Name: ", reflect.TypeOf(name))
		fmt.Println("Type Age: ", reflect.TypeOf(age))
		fmt.Println("Type Height: ", reflect.TypeOf(height))
	}
}

func menu(){
	fmt.Println("\n#### MENU ####")
	fmt.Println("1) Start monitoring.")
	fmt.Println("2) Show logs.")
	fmt.Println("3) Exit")
}

func readCommand() int{
	var command int
	fmt.Scan(&command)
	return command
}