package main
// Write a program which reads information from a file and represents it in a slice of structs. 
// Assume that there is a text file which contains a series of names. Each line of the text file 
// has a first name and a last name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields, fname for the first name, 
// and lname for the last name. Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file. Your program will 
// successively read each line of the text file and create a struct which contains the first 
// and last names found in the file. Each struct created will be added to a slice, and after 
// all lines have been read from the file, your program will have a slice containing one 
// struct for each line in the file. After reading all lines from the file, your program should 
// iterate through your slice of structs and print the first and last names found in each struct.

// Submit your source code for the program, “read.go”.

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"regexp"
)

type Person struct {
	fname string
	lname string
}

func getInputDataFileName(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	filename, _ := reader.ReadString('\n')
	return strings.Trim(filename, "\n")
}

func main() {
	filename := getInputDataFileName("Enter full path name for data file: ")
	//file, _ := os.Open("/Users/hieutrinh/git/coursera_uci/programmingwithgo/src/pkgweek4/assignment2/data.txt")
	file, _ := os.Open(filename)
	defer file.Close()

	// make a slice of window size of 3 and capacity of 3
	persons := make([]Person, 0)

	// 1 or more spaces
	pattern := regexp.MustCompile(` +`)   
	
	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// processing each line
		tokens := pattern.Split(line,-1)
		person := Person{fname: tokens[0], lname: tokens[1]}
		persons = append(persons, person)
	}

	// Iterate through the 'persons' slice and print each entry "fname/lastname" in pair
	for _, person := range persons {
		fmt.Printf("%s/%s", person.fname, person.lname)
	}
}