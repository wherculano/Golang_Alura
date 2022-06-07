package main

import "fmt"
import "reflect"



func main(){
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
	fmt.Println(age, "!! Oh really?")

	fmt.Println("And what's your height, please?")
	fmt.Scan(&height)
	
	fmt.Println("Type Name: ", reflect.TypeOf(name))
	fmt.Println("Type Age: ", reflect.TypeOf(age))
	fmt.Println("Type Height: ", reflect.TypeOf(height))
}