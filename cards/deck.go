package main

import "fmt"

//!  Create a new type of deck
//! which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{
		"Spades", "Diamonds", "Hearts", "Clubs",
	}
	cardValue := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
	}

	//! If you don't want to use any variable
	//! Go will throw an error to over it
	//! Fix it with underscore

	for _, suit := range cardSuits {
		// Here first variable references to index
		// second actual value of slice
		for _, value := range cardValue {
			// Append func don't modify the original array its return a new one
			cards = append(cards, suit+" of "+value)

		}
	}

	return cards
}

//! Receiver function
//! (d deck) is the receiver
func (d deck) print() {
	// For every iteration new variables are declared
	//ie new variable assignment operator is used
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
