package main

// Write a Bubble Sort program in Go. The program should prompt the user to 
// type in a sequence of up to 10 integers. The program should print the 
// integers out on one line, in sorted order, from least to greatest. Use 
// your favorite search tool to find a description of how the bubble sort 
// algorithm works.

// As part of this program, you should write a function called BubbleSort() 
// which takes a slice of integers as an argument and returns nothing. The 
// BubbleSort() function should modify the slice so that the elements are 
// in sorted order.

// A recurring operation in the bubble sort algorithm is the Swap operation 
// which swaps the position of two adjacent elements in the slice. You 
// should write a Swap() function which performs this operation. Your Swap() 
// function should take two arguments, a slice of integers and an index 
// value i which indicates a position in the slice. The Swap() function 
// should return nothing, but it should swap the contents of the slice in 
// position i with the contents in position i+1.


import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserInput(prompt string) []int {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	split := strings.Split(text, ",")
	sli := []int{}
	for _, s := range split {
		num, _ := strconv.Atoi(s)
		sli = append(sli, num)
	}
	return sli
}

func BubbleSort(sli []int) {
	var (
		n       = len(sli)
		swapped = true
	)

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
				swapped = true
			}
		}
		// The last element is the largest number in the slice
		// so, we only need to bubble sort the remaining n-1
		n = n - 1
	}
}

func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}

func PrintSlice(sli []int) {
	for i := range sli {
		fmt.Printf("%d ", sli[i])
	}
}

func main() {
	prompt := "Enter up to 10 integers with comma separated: "
	sli := getUserInput(prompt)
	if (len(sli) <= 10) {
		BubbleSort(sli)
		PrintSlice(sli)
	} else {
		fmt.Println("Error: You entered more than 10 entries")
	}
}