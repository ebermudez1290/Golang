package main

func main() {
	// var card string = "ace of spaces"
	// card := newCard()
	// fmt.Println(card)
	//This is the way to init a slice.
	// Cards := deck{"ace of spades", "ace of diamonds"}
	// Cards = append(Cards, "six of spades")
	// print(Cards)
	//cards:=newDeck()
	// hand, remainingDeck := deal(cards,5)
	// hand.print()
	// remainingDeck.print()
	// fmt.Println(cards.toString())
	cards:=newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "ace of spades"
// }
