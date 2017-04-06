package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {																				// Create the Card struct so we know what a Card looks like.
	Rank  int	
	Suit	int
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())												// Seed the RNG with the current time in nanoseconds.
	return rand.Intn(max - min) + min												// Return a value between the min and max sent to the function.
}

func shuffle(deck []Card) {
	  var tempCard Card
		for position := 0; position < 52; position++ {
			randomNumber := random(0, 51)
			tempCard = deck[position]														// Copy the current card to a placeholder.
			deck[position] = deck[randomNumber]									// Copy a random card to the current card position.
			deck[randomNumber] = tempCard												// Copy the card in the placeholder to the position of the random card.
		}
}

func printCard(Rank int, Suit int) {
	var printableRank string
	var printableSuit string
	switch Rank {
		case 0: 
			printableRank = "Ace"
		case 1:
			printableRank = "Two"
		case 2:
			printableRank = "Three"
		case 3:
			printableRank = "Four"
		case 4:
			printableRank = "Five"
		case 5:
			printableRank = "Six"
		case 6:
			printableRank = "Seven"
		case 7:
			printableRank = "Eight"
		case 8:
			printableRank = "Nine"
		case 9:
			printableRank = "Ten"
		case 10:
			printableRank = "Jack"
		case 11:
			printableRank = "Queen"
		case 12:
			printableRank = "King"
	}
	switch Suit {
		case 0:
			printableSuit = "Clubs"
		case 1:
			printableSuit = "Diamonds"
		case 2:
			printableSuit = "Hearts"
		case 3:
			printableSuit = "Spades"
	}
	fmt.Println(printableRank, "\tof ", printableSuit)
}

func main() {	
	var deck []Card																					// Create the deck that's going to hold all the Cards (a var array of type Card).
	for suit := 0; suit < 4; suit++ {												// Count through the cards (4 suits and 13 ranks).
		for rank := 0; rank < 13; rank++ {	
			card := Card{rank, suit}														// Make a card of the current suit and rank and assign that to a placeholder Card.
			deck = append(deck, card)														// Append the placeholder card to the end (?) of the deck.
		}
	}
	shuffle(deck)																						// Shuffle the cards.
	for position := 0; position < 52; position++ {					// Print each card in the stack, in the newly shuffled order.
		printCard(deck[position].Rank, deck[position].Suit)
	}
}
