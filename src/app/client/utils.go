package client

import (
	"../game"
)

//tokenToString converts a token to a normalized string
func tokenToString(token game.Token) string {
	rv := ""
	if token.Value > 10 {
		rv = rv + string(token.Value)
	} else {
		rv = rv + "0" + string(token.Value)
	}

	switch token.Color {
	case game.BLACK:
		rv = rv + "D"
	case game.RED:
		rv = rv + "R"
	case game.BLUE:
		rv = rv + "B"
	case game.YELLOW:
		rv = rv + "Y"
	default:
		rv = rv + "W"
	}

	return rv
}

//moveToString converts a move into a normalized string
func moveToString(move game.Move) string {
	rv := ""
	for _, v := range move.Tokens {
		rv = rv + tokenToString(v)
	}

	return rv
}
