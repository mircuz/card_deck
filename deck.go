package main

import "fmt"

type deck []string

func NewDeck() deck {
	semi := []string{"Quadri", "Cuori", "Picche", "Fiori"}
	val := []string{"Asso", "Due", "Tre", "Quattro"}

	d := deck{}
	for _, i := range semi {
		for _, j := range val {
			d = append(d, j+" di "+i)
		}
	}
	return d
}

func (d deck) printCard() {
	for i, card := range d {
		fmt.Println("Carta ", i, " è ", card)
	}
}

func (d deck) newCard(s string) deck {
	d = append(d, s)
	return d
}

func (d deck) deal(hand int) (deck, deck) {
	return d[:hand], d[hand:]
}
