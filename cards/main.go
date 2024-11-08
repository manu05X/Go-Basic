package main

// func main() {
// 	//var card string = "ACE of Spades"
// 	card := "Ace of Spade"
// 	card = "Five of Dimond"
// 	fmt.Println(card)
// }

func main() {

	//card := newCardd()
	//cards := []string{"Ace of Dimonds", newCard()}
	//cards := deck{"Ace of Dimonds", newCard()}
	//cards = append(cards, "six of spade") // does not have same cards but create new cards and return it while using append

	//fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()
	// //cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// //Type conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// //deckType->[]string Type-> string -> []byte Type
	// cards := newDeck()
	// //fmt.Println(cards.toString()) //return one single string
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

// func newCard() string {
// 	return "Five of Dimonds"
// }
