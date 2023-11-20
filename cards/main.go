package main
import "fmt"

func main() {
	//cards := newDeck()
	//	cards.print()

	//	hand, new_deck := deal(cards, 13)

	//	hand.print()
	// new_deck.print()
	// fmt.Println(cards.toByteSlice())
	// cards.saveToFile("mycards.deck")

	loadCard := newDeckFromFile("mycards.deck")
	loadCard.print()
	loadCard.shuffle()

	fmt.Println("************************")
	loadCard.print()
}

// function with a receiver
