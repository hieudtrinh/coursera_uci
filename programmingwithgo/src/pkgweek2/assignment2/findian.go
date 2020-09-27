package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

// This function prompts for user input and performs a business logic
// and return true if the user input string meet our search criterias
func findian(prompt string) (bool) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	s,_ := reader.ReadString('\n')
	s = strings.ToLower(strings.Trim(s, "\n"))
	if len(s) < 2 {
		return false
	}
	
	containsA := strings.ContainsAny(s,"a")
	if (s[0] == 'i' && s[len(s)-1] == 'n' && containsA) {
		return true;
	}
	return false;
	
}

func main() {
	var prompt string = "Enter a string: "
	for i:=0; i<2; i++ {
		if findian(prompt) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
