package game

import (
	"errors"
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
}

//Bag struct for representing an array of 106 Tokens
type Bag struct {
	Bag []Token
}

//CreateBag from a random
func CreateBag() Bag {
	var bag Bag
	var token Token

	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			for k := 0; k < 2; k++ {
				token = Token{Value: j + 1, Color: i + 1}
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

//Extract n Toekn from bag b
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
