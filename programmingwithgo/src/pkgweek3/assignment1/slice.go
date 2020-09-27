package main

import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)

func getUserInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	s,err := reader.ReadString('\n')
	return strings.Trim(s, "\n"), err
}

func main() {
	var prompt string = "Entry a number: "

	// make a slice of window size of 3 and capacity of 3
	sli := make([]int, 3)
	var text,_ = getUserInput(prompt)

	for i:=0; text[0] != 'X'; i++ {
		fmt.Println(text)
		num,_ := strconv.Atoi(text)
		if i == len(sli) {
			// grow the slice size
			sli = append(sli, num)
		} else {
			sli[i] = num
		}
		text,_ = getUserInput(prompt)
	}
	
	fmt.Println("Program exit - user requested")
}