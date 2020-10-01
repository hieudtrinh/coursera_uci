package main
//	Animal		Food eaten		Locomtion method	Spoken sound
//  ============================================================
//	cow			grass			walk				moo
//	bird		worms			fly					peep
//	snake		mice			slither				hsss
//
// Your program should present the user with a prompt, “>”, to indicate that
// the user can type a request. Your program should accept one command at a
// time from the user, print out a response, and print out a new prompt on a
// new line. Your program should continue in this loop forever. Every command
// from the user must be either a “newanimal” command or a “query” command.
//
// Each “newanimal” command must be a single line containing three strings.
// The first string is “newanimal”. The second string is an arbitrary string
// which will be the name of the new animal. The third string is the type of
// the new animal, either “cow”, “bird”, or “snake”. Your program should
// process each newanimal command by creating the new animal and printing
// “Created it!” on the screen.
//
// Each “query” command must be a single line containing 3 strings. The first
// string is “query”. The second string is the name of the animal. The third
// string is the name of the information requested about the animal, either
// “eat”, “move”, or “speak”. Your program should process each query command
// by printing out the requested data.
//
// Define an interface type called Animal which describes the methods of an
// animal. Specifically, the Animal interface should contain the methods
// Eat(), Move(), and Speak(), which take no arguments and return no values.
// The Eat() method should print the animal’s food, the Move() method should
// print the animal’s locomotion, and the Speak() method should print the
// animal’s spoken sound. Define three types Cow, Bird, and Snake. For each
// of these three types, define methods Eat(), Move(), and Speak() so that
// the types Cow, Bird, and Snake all satisfy the Animal interface. When the
// user creates an animal, create an object of the appropriate type. Your
// program should call the appropriate method when the user issues a query
// command.

import (
	"bufio"
	"fmt"
	"go/types"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

// COW type definition
type Cow types.Nil

func (t Cow) Eat() {
	fmt.Println("grass")
}

func (t Cow) Move() {
	fmt.Println("walk")
}

func (t Cow) Speak() {
	fmt.Println("moo")
}

// BIRD type definition
type Bird types.Nil

func (t Bird) Eat() {
	fmt.Println("worms")
}

func (t Bird) Move() {
	fmt.Println("fly")
}

func (t Bird) Speak() {
	fmt.Println("peep")
}

// SNAKE type definition
type Snake types.Nil

func (t Snake) Eat() {
	fmt.Println("mice")
}
func (t Snake) Move() {
	fmt.Println("slither")
}
func (t Snake) Speak() {
	fmt.Println("hsss")
}

func getUserInput() []string {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	return strings.Fields(text)
}

func getAnimalType(animalType string) Animal {
	switch animalType {
	case "cow":
		return Cow{}
	case "bird":
		return Bird{}
	case "snake":
		return Snake{}
	default:
		return nil
	}
}

func main() {
	animals := make(map[string]Animal)
	// request example:
	// 1) newanimal sammy cow
	// 2) newanimal miso bird
	// 3) newanimal kenny snake
	// 4) sammy eat | sammy move | sammy speak
	// 4) miso eat | miso move | miso speak
	// 4) kenny eat | kenny move | kenny speak
	fmt.Println("Welcome to my virtual animal farm!")
	fmt.Println("You can create animals or query information about the animals you have created before")
	fmt.Println("Example: <newanimal> <animal_name> [cow|bird|snake]")
	fmt.Println("         <animal_name> [eat|move|speak]")
	for {
		request := getUserInput()
		switch {
		case len(request) == 2:
			// <name_of_animal> <request_information>
			animal := animals[request[0]]
			if animal == nil {
				fmt.Println("Animal not exist: ", animals[request[0]])
				continue
			}
			switch request[1] {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				{
					fmt.Println("Requested information is not supported: ", request[1])
					continue
				}
			}
		case len(request) == 3:
			// <newanimal> <name_of_animal> <animal_type>
			if request[0] != "newanimal" {
				fmt.Println("Error - unkonw request ", request[0])
				continue
			}
			animal := animals[request[1]]
			if animal != nil {
				fmt.Println("Requested animal name exist: ", request[1])
				continue
			}
			animals[request[1]] = getAnimalType(request[2])
			fmt.Println("Created it!")

		default:
			fmt.Println("Not supported request: ", request)
		}
	}
}
