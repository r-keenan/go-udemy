package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.printDeck()
}