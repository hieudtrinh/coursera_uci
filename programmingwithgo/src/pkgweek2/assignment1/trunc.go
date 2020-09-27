package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

/*
	promtp for user input and return the input text enter by user
*/
func readUserInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	return text, err
}

func main() {
	var prompt string = "Enter a floating number: "
	
	// Run the test twice in the for loop
	for i:=0; i<2; i++ {
		var text, err = readUserInput(prompt)
		if err != nil {
			fmt.Printf("Encountered input error %s\n", err)
			continue;
		}
		text = strings.Trim(text, "\n")
		var val, err2 = strconv.ParseFloat(text, 64)
		if err2 != nil {
			fmt.Printf("Encountered input error %s\n", err2)
		} else {
			fmt.Println("User entered: ", val)
			fmt.Println("Truncated value: ", math.Round(val))
		}
		
	}
}
