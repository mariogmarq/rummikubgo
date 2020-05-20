package game

import (
	"math/rand"
	"time"
)

//Game struct for zipping all other classes
type Game struct {
	Players []Player
	Table   Table
	Bag     Bag
	Turn    []Move
}

//Print game to screen
func (g Game) Print() {
	g.Table.Print()
	g.Players[0].Print()
}

//CreateGame function
func CreateGame(online bool) Game {
	g := Game{}
	g.Bag = CreateBag(rand.New(rand.NewSource(time.Now().UnixNano())))
	if !online {
		g.Players = append(g.Players, CreatePlayer())
	}
	return g
}

//EndTurn of a player
func (g *Game) EndTurn() {
	isValid := g.Table.CheckTable()

	if isValid {
		for _, v := range g.Turn {
			if !v.IsValid() {
				isValid = false
			}
		}
	}

	if isValid {
		for _, v := range g.Turn {
			g.Table.AddMove(v)
		}
	}
	g.Turn = nil
}
