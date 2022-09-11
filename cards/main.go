package main

import "fmt"

func main() {

	cards := newDeck()

	//! Calling method on deck type
	cards.print()

	//! Receiving multiple value from function
	hand, remainDeck := deal(cards, 5)

	fmt.Println("hand", hand)
	fmt.Println("remainDeck", remainDeck)

}
