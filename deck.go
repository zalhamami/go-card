package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// d is receiver from deck type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// This function is used to turn deck into string
func (d deck) toString() string {
	return strings.Join(d, ", ")
}

// This function is used to save the deck into file
func (d deck) saveToFile(filename string) error {
	byteConversion := []byte(d.toString())
	return ioutil.WriteFile(filename, byteConversion, 0666)
}

// This function is used to create a new deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// The underscore is to tell Go those variable are not used
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// This function is used to return separate deck
func deal(d deck, handSize int) (deck, deck) {
	dealCards := d[:handSize]
	remainingCards := d[handSize:]
	return dealCards, remainingCards
}
