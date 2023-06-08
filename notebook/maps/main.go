package main

import "fmt"

// Maps are Reference Type, that is, pass by reference

func main() {
	fruits := map[string]string{
		"blueberry": "blue",
		"raspberry": "red",
		"pineapple": "yellow",
		"apple":     "red",
	}

	fruits["dragon fruit"] = "pink"

	fmt.Println(fruits)

	// declaring an empty map or initializing a map
	vegetables := make(map[string]string)
	vegetables["potato"] = "white"
	fmt.Println(vegetables)

	iteratingOverMapAndPrint(fruits)
}

func iteratingOverMapAndPrint(m map[string]string) {
	for key, val := range m {
		fmt.Println(key, val)
	}
}
