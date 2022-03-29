package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	fmt.Println("-------------")
	for i, card := range d {

		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}

	stringCards := string(byteSlice)

	stringArray := strings.Split(stringCards, ",")

	return deck(stringArray)

}

func (d deck) shuffleCards() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	size := len(d)
	for i := range d {
		randomNumber := r.Intn(size - 1)
		// aux := d[i]
		// d[i] = d[randomNumber]
		// d[randomNumber] = aux

		d[i], d[randomNumber] = d[randomNumber], d[i]
	}
}
