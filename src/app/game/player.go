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
	p.Tokens[0].Selected = 1
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

//delete a position
func (p *Player) delete(pos int) {
	var arr []Token
	for i, v := range p.Tokens {
		if i != pos {
			arr = append(arr, v)
		}
	}
	p.Tokens = arr
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
		if v.Selected == 1 || v.Selected == 3 {
			return i
		}
	}
	return -1
}

//ChangeSelected takes a bool, true right, false left
func (p *Player) ChangeSelected(right bool) {
	pos := p.FindSelected()
	val1 := 0
	val2 := 1
	if pos != -1 {
		if right {
			if pos < len(p.Tokens)-1 {
				if p.Tokens[pos+1].Selected == 2 {
					val2 = 3
				}
				if p.Tokens[pos].Selected == 3 {
					val1 = 2
				}
				p.Tokens[pos].Selected, p.Tokens[pos+1].Selected = val1, val2
			}
		} else {
			if pos > 0 {
				if p.Tokens[pos-1].Selected == 2 {
					val2 = 3
				}
				if p.Tokens[pos].Selected == 3 {
					val1 = 2
				}
				p.Tokens[pos].Selected, p.Tokens[pos-1].Selected = val1, val2
			}
		}
	}
}

//SelectToken from hand
func (p *Player) SelectToken() {
	pos := p.FindSelected()
	if p.Tokens[pos].Selected == 3 {
		p.Tokens[pos].Selected = 1
	} else {
		p.Tokens[pos].Selected = 3
	}
}

//CreateMove from hand
func (p *Player) CreateMove() Move {
	var pos []int
	mov := Move{Score: 0}
	for i := range p.Tokens {
		if p.Tokens[i].Selected == 2 || p.Tokens[i].Selected == 3 {
			pos = append(pos, i)
		}
	}
	for _, v := range pos {
		mov.Tokens = append(mov.Tokens, p.Tokens[v])
	}

	return mov
}

//SearchToken in hand and return position
func (p Player) SearchToken(t Token) int {
	for i, v := range p.Tokens {
		if v == t {
			return i
		}
	}

	return -1
}

//RemoveMove from hand
func (p *Player) RemoveMove(m Move) {
	for _, v := range m.Tokens {
		p.delete(p.SearchToken(v))
	}
}

//ResetSelected config
func (p *Player) ResetSelected() {
	for i := range p.Tokens {
		p.Tokens[i].Selected = 0
	}
	if len(p.Tokens) > 0 {
		p.Tokens[0].Selected = 1
	}
}
