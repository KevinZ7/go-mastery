package main

import "fmt"

type compositeKey struct {
	id   int
	name string
}

func main() {
	// Creating map of key(string) : value(int)
	intMap := make(map[string]int)

	// insert
	intMap["one"] = 1

	// get element
	fmt.Printf("printing one in map: %d \n", intMap["one"])

	// delete
	delete(intMap, "one")
	fmt.Printf("Deleting map[one] \n")
	_, ok := intMap["one"]
	if ok {
		fmt.Printf("map['one'] was not deleted, value is \n")
	} else {
		fmt.Printf("map['one'] was deleted \n")
	}

	// Using structs as a map key
	compositeKeyMap := make(map[compositeKey]int)
	compositeKeyMap[compositeKey{id: 2, name: "kz"}] = 1
	fmt.Printf("the value for the key {id:2, name:'kz'} in compositeKeyMap is %d", compositeKeyMap[compositeKey{id: 2, name: "kz"}])

}
