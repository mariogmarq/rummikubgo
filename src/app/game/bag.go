package game

import (
	"errors"
	"fmt"
	"github.com/logrusorgru/aurora"
	"math/rand"
)

//Token colors

//BLACK color
const BLACK = 1

//RED color
const RED = 2

//BLUE color
const BLUE = 3

//YELLOW color
const YELLOW = 4

//WILDCARD color
const WILDCARD = 0

//Token struct for representing the tuple Value and Color
type Token struct {

	//Value of the token
	Value int

	//Color of the token
	Color int

	//Selected in the player
	Selected int //0 no selected, 1 selected, 2 for move, 3 for selected and move

}

//Bag struct for representing an array of 106 Tokens
type Bag struct {
	Bag []Token
}

//CreateBag from a random
func CreateBag() Bag {
	bag := Bag{Bag: nil}
	var token Token

	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			for k := 0; k < 2; k++ {
				token = Token{Value: j + 1, Color: i + 1, Selected: 0}
				bag.Bag = append(bag.Bag, token)
			}
		}
	}

	for i := 0; i < 2; i++ {
		token = Token{Value: 0, Color: 0}
		bag.Bag = append(bag.Bag, token)
	}

	rand.Shuffle(len(bag.Bag), func(i, j int) {
		bag.Bag[i], bag.Bag[j] = bag.Bag[j], bag.Bag[i]
	})

	return bag
}

//Extract n Token from bag b, error if n < 0
func (b *Bag) Extract(n int) ([]Token, error) {
	var arr []Token
	if n < 0 {
		return arr, errors.New("Negative integer for extracting from bag")
	} else if n > len(b.Bag) {
		n = len(b.Bag)
	}

	if n == 0 {
		return arr, nil
	}

	for i := 0; i < n; i++ {
		arr = append(arr, b.Bag[i])
	}

	b.Bag = b.Bag[n+1 : len(b.Bag)]

	return arr, nil
}

//Print a token
func (t Token) Print() {
	switch t.Selected {
	case 0:
		switch t.Color {
		case RED:
			fmt.Printf("%d", aurora.Red(t.Value))
		case BLUE:
			fmt.Printf("%d", aurora.Blue(t.Value))
		case BLACK:
			fmt.Printf("%d", aurora.Black(t.Value))
		case YELLOW:
			fmt.Printf("%d", aurora.BrightYellow(t.Value))
		default:
			fmt.Printf("%d", aurora.BrightMagenta(t.Value))
		}
	case 1:
		switch t.Color {
		case RED:
			fmt.Printf("%d", aurora.SlowBlink(aurora.Red(t.Value)))
		case BLUE:
			fmt.Printf("%d", aurora.SlowBlink(aurora.Blue(t.Value)))
		case BLACK:
			fmt.Printf("%d", aurora.SlowBlink(aurora.BrightBlack(t.Value)))
		case YELLOW:
			fmt.Printf("%d", aurora.SlowBlink(aurora.BrightYellow(t.Value)))
		default:
			fmt.Printf("%d", aurora.SlowBlink(aurora.BrightMagenta(t.Value)))
		}
	case 2:
		switch t.Color {
		case RED:
			fmt.Printf("%d", aurora.BgRed(t.Value))
		case BLUE:
			fmt.Printf("%d", aurora.BgBlue(t.Value))
		case BLACK:
			fmt.Printf("%d", aurora.BgBrightBlack(t.Value))
		case YELLOW:
			fmt.Printf("%d", aurora.BgBrightYellow(t.Value))
		default:
			fmt.Printf("%d", aurora.BgBrightMagenta(t.Value))
		}
	case 3:
		switch t.Color {
		case RED:
			fmt.Printf("%d", aurora.SlowBlink(aurora.BgRed(t.Value)))
		case BLUE:
			fmt.Printf("%d", aurora.SlowBlink(aurora.BgBlue(t.Value)))
		case BLACK:
			fmt.Printf("%d", aurora.SlowBlink(aurora.BgBrightBlack(t.Value)))
		case YELLOW:
			fmt.Printf("%d", aurora.SlowBlink(aurora.BgBrightYellow(t.Value)))
		default:
			fmt.Printf("%d", aurora.SlowBlink(aurora.BgBrightMagenta(t.Value)))
		}
	}
}
