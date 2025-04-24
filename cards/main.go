package main

func main() {
	// var card string = "Ace of Spades"
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Clubs")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
