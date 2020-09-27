package main

// Write a program which prompts the user to first enter a name, and then enter 
// an address. Your program should create a map and add the name and address to 
// the map using the keys “name” and “address”, respectively. Your program should 
// use Marshal() to create a JSON object from the map, and then your program 
// should print the JSON object.

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	s,_ := reader.ReadString('\n')
	return strings.Trim(s, "\n")
}

func main() {
	name := getUserInput("Enter name: ")
	address := getUserInput("Enter address: ")
	person := make(map[string]string)
	person["name"] = name
	person["address"] = address
	barr,_ := json.Marshal(person)
	fmt.Print(string(barr))
}