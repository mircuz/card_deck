package main

func main() {
	cards := NewDeck()
	hd, rest := cards.deal(3)
	hd.printCard()
	rest.printCard()
}
