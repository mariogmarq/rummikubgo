package game

import "fmt"

//Player struct for representing a person playing the game
type Player struct {
	Tokens []Token
	Score  int
}

//CreatePlayer strutc empty
func CreatePlayer() Player {
	player := Player{Tokens: nil, Score: 0}
	return player
}

//AddToken from bag
func (p *Player) AddToken(arr []Token) {
	p.Tokens = append(p.Tokens, arr...)
	p.OrderHand()
	p.Tokens[0].Selected = true
}

//OrderHand player array
func (p *Player) OrderHand() {
	for i := 0; i < len(p.Tokens); i++ {
		for j := 1; j < len(p.Tokens); j++ {
			if p.Tokens[j].Value < p.Tokens[j-1].Value {
				p.Tokens[j], p.Tokens[j-1] = p.Tokens[j-1], p.Tokens[j]
			}
		}
	}
}

//Print player hand
func (p Player) Print() {
	for _, v := range p.Tokens {
		v.Print()
		fmt.Printf(" ")
	}
}

//FindSelected token position
func (p Player) FindSelected() int {
	for i, v := range p.Tokens {
		if v.Selected {
			return i
		}
	}
	return -1
}

//ChangeSelected takes a bool, true right, false left
func (p *Player) ChangeSelected(right bool) {
	pos := p.FindSelected()
	if pos != -1 {
		if right {
			if pos < len(p.Tokens)-1 {
				p.Tokens[pos].Selected, p.Tokens[pos+1].Selected = false, true
			}
		} else {
			if pos > 0 {
				p.Tokens[pos].Selected, p.Tokens[pos-1].Selected = false, true
			}
		}
	}
}
