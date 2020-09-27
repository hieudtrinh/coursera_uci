package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var slice = make([]int, 3)
	var index = 0

	for {
		fmt.Println("Insert a Number (or X to leave):")
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		} else {
			// removes '\n'
			input = strings.TrimSuffix(input, "\n")
			if input == "X" {
				return
			} else {
				value, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("You should insert an integer or X!")
				} else {
					if index < 3 {
						slice[index] = value
						var sorted = make([]int, len(slice))
						copy(sorted, slice)
						sort.Ints(sorted)
						fmt.Println(sorted)
						index++
					} else {
						slice = append(slice, value)
						sort.Ints(slice)
						fmt.Println(slice)
					}
				}
			}
		}
	}
}
