package main

import "fmt"

func main() {
	// var = variables, informs go that you are creating a variable
	// card = name of the variable
	// string = setting type of the variable
	// "Ace of Spades!" actual value
	var card string = "Explicity typing variable"
	card2 := "Ace of Spades, the type of variable is being inferred from the value"

	fmt.Println(card, card2)
}
