package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
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
	
	websites := getWebSitesFromAFile()

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
	req, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro:", err)
	}
	
	if req.StatusCode == 200 {
		fmt.Println("The website", site, "is ok")
		createLog(site, true)
	} else {
		fmt.Println("The website", site, "is out. Status Code:", req.StatusCode)
		createLog(site, false)
	}
}

func getWebSitesFromAFile()[]string{
	var urls []string

	// file, err := ioutil.ReadFile("websites.txt") // reads and shows the whole file
	file, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)
	
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		
		urls = append(urls, line)

		if err == io.EOF{
			break 
		}
	}

	file.Close()

	return urls
}


func createLog(website string, status bool){
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	
	file.WriteString(time.Now().Format("02/01/2006 15:04:04") + "-" + website + "- online:" + strconv.FormatBool(status)+"\n")
	file.Close()
}

// Date/Time documentation
// https://go.dev/src/time/format.go 