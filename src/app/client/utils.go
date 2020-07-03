package client

import (
	"../game"
	"strconv"
	"strings"
)

//TokenToString converts a token to a normalized string
func TokenToString(token game.Token) string {
	rv := ""
	if token.Value >= 10 {
		rv = rv + strconv.Itoa(token.Value)
	} else {
		rv = rv + "0" + strconv.Itoa(token.Value)
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

//MoveToString converts a move into a normalized string
func MoveToString(move game.Move) string {
	rv := ""
	for _, v := range move.Tokens {
		rv = rv + TokenToString(v)
	}

	rv = rv + ";"
	return rv
}

//TableToString converts the table into a normalized string
func TableToString(table game.Table) string {
	rv := ""
	for _, v := range table.Matrix {
		rv = rv + MoveToString(v)
	}

	rv = rv + "\n"
	return rv
}

//StringToToken transforms a normalized string into a token
func StringToToken(s string) game.Token {
	var token game.Token
	value, _ := strconv.Atoi(s[0:2])
	if s[2] == 'D' {
		token = game.Token{Value: value, Color: game.BLACK, Selected: 0}
	} else if s[2] == 'R' {
		token = game.Token{Value: value, Color: game.RED, Selected: 0}
	} else if s[2] == 'B' {
		token = game.Token{Value: value, Color: game.BLUE, Selected: 0}
	} else if s[2] == 'Y' {
		token = game.Token{Value: value, Color: game.YELLOW, Selected: 0}
	} else {
		token = game.Token{Value: value, Color: game.WILDCARD, Selected: 0}
	}

	return token
}

//StringToMove transforms a normalized string into a move
func StringToMove(s string) game.Move {
	var mov game.Move
	for i := 0; i+2 < len(s); i = i + 3 {
		mov.Tokens = append(mov.Tokens, StringToToken(s[i:i+3]))
	}

	return mov
}

//StringToTable transforms a normalized string into a table
func StringToTable(s string) game.Table {
	var table game.Table
	parsed := strings.Split(s, ";")

	for _, v := range parsed {
		table.Matrix = append(table.Matrix, StringToMove(v))
	}

	return table
}
