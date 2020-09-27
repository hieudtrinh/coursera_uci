// main  reads information from a file and represents it in a slice of structs
// then after reading the file prints the first and last name stored in each struct of the slice
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var fileName string
	nameSlice := make([]name, 0, 3)
	fmt.Printf("Enter name of file (along with extension i.e. .txt): ")
	fmt.Scan(&fileName)
	f, err := os.Open(fileName)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		nameArr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if len(nameArr) == 1 {
			continue
		}

		nameSlice = append(nameSlice, name{fname: nameArr[0], lname: nameArr[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, n := range nameSlice {
		fmt.Printf("First Name: %v, last Name: %v\n", n.fname, n.lname)
	}

}
