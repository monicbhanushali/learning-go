package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {

	baburao := person{"Baburao", "Apte"}
	fmt.Println(baburao)

	raju := person{firstname: "raju", lastname: "chacha"}
	fmt.Println(raju)

	var shyam person
	shyam.firstname = "Shyam"
	fmt.Println(shyam)
	fmt.Printf("%+v", shyam)
}
