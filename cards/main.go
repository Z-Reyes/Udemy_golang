package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)
	hand.print()
	fmt.Println(("End of hand values"))
	remainingDeck.print()
	fmt.Println(hand.toString())
	hand.saveToFile("card_test.txt")

	filehand, err := newDeckFromFile("card_test.txt")

	if err == nil {
		fmt.Println("READ DECK FROM FILE:")
		filehand.print()
	} else {
		fmt.Println("Error OCCURRED!!!!! :((((((")
	}

	filehand.shuffle()
	fmt.Println("After shuffle")
	filehand.print()

}
