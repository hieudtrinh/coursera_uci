package main

// Let us assume the following formula for displacement s as a function of 
// time t, acceleration a, initial velocity vo, and initial displacement so.

// S = 1/2 a T^2 + VoT + So

// Write a program which first prompts the user to enter values for 
// acceleration, initial velocity, and initial displacement. Then the program 
// should prompt the user to enter a value for time and the program should 
// compute the displacement after the entered time.

// You will need to define and use a function called GenDisplaceFn() which 
// takes three float64 arguments, acceleration a, initial velocity vo, and 
// initial displacement so. GenDisplaceFn() should return a function which 
// computes displacement as a function of time, assuming the given values 
// acceleration, initial velocity, and initial displacement. The function 
// returned by GenDisplaceFn() should take one float64 argument t, 
// representing time, and return one float64 argument which is the displacement 
// travelled after time t.

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

/*
 * prompt user for input and return input text to caller
 */
func GetLine(s string) string {
	fmt.Printf(s)
        in := bufio.NewReader(os.Stdin)
        line, _ := in.ReadString('\n')
        line = strings.TrimSpace(line)
	return line
}

/*
 * Return the linear kinematic function to the caller
 */
func GenDisplaceFn(a float64, v float64, s float64) func(float64) float64 {
	return func(t float64) float64 {
		displ := 0.5 * a * t * t + v * t + s;
		return displ
	}
} 

func main() {
	fmt.Println("Week2 - peer assignment1 - create a routine that solves a linear kinematics")
	a,_ := strconv.ParseFloat(GetLine("Enter acceleration: "), 64)
	v,_ := strconv.ParseFloat(GetLine("Enter initial velocity: "), 64)
	s,_ := strconv.ParseFloat(GetLine("Enter initial displacement: "), 64)
	t,_ := strconv.ParseFloat(GetLine("Enter initial displacement: "), 64)
	
	fn := GenDisplaceFn(a,v,s)
	fmt.Printf("Linear kinenatic value: %f\n", fn(t))
	fmt.Println("Thank you for using")
}