package main

import "fmt"

func main() {

	cards := newDeck()
	cards = append(cards, "Six of Spades")

	fmt.Println("Cards: ", cards[0:2])

	// for i, card := range cards {
	// 	fmt.Print(i, card, "\n")
	// }

	// cards.print()
	hand, remainingCards := deal(cards, 5)

	// fmt.Println("Hand: ", hand)
	// fmt.Println("RemainingCards: ", remainingCards)
	hand.print()
	remainingCards.print()
}
