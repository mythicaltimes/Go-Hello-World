package deck

import (
	"testing"
	"fmt"
)

func ExampleCard(){
	fmt.Println(Card{Rank: Ace,Suit: Heart})
	fmt.Println(Card{Rank: Two,Suit: Spade})
	fmt.Println(Card{Rank: Nine,Suit: Diamond})
	fmt.Println(Card{Rank: Jack,Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T){
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4{
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestJokers(t *testing.T){
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards{
		if c.Suit == Joker{
			count++
		}
	if count != 3{
		t.Error("Expected 3 Jokers, received:", count)
	}

	}
}