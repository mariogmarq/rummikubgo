package game

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
