package main

import "fmt"

func main() {

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	fmt.Println("Cards: ", cards)

	for i, card := range cards {
		fmt.Print(i, card, "\n")
	}

}

func newCard() string {
	return "Five of Diamonds"
}
