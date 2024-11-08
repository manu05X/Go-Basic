package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new Type of 'deck'
// which is a slice of string
type deck []string

// create newDeck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Dimonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Reciver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
here d is instance of deck or d is equvalent to cards
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

or we can say that d is similar to "this" in JS
*/

// returning 2 seperate value using one function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// deckType->[]string Type-> string -> []byte Type
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save a file to hard drive
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// read a file of hard drive
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",") //Ace of Spades,Two of Spades,Three of Spades,.,.,.,.,Four of Clubs

	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
