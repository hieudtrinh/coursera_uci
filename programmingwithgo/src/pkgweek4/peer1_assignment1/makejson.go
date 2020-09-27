// main prompts the user to first enter a name, and then enter an address
// creates a map from this and the converts the map into a json object and prints it
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	person := make(map[string]string)
	var pname, paddr string
	fmt.Printf("Please enter Name of person: ")
	fmt.Scan(&pname)
	fmt.Printf("Please enter address of person: ")
	fmt.Scan(&paddr)
	person["name"] = pname
	person["address"] = paddr
	jobj, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("JSON Object: %v\n", string(jobj))

}
