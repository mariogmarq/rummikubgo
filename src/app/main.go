package main

import (
	"./game"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

//bytes for moving
var left = []byte{'a'}
var right = []byte{'d'}
var jumptotable = []byte{'w'}
var jumptoplayer = []byte{'s'}
var draw = []byte{'q'}
var enter = []byte{10}
var space = []byte{32}
var commit = []byte{'e'}

func lookError(err error) {
	if err != nil {
		panic(err)
	}
}

//Clears screen
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//Compare two bytes
func equal(a, b []byte) bool {
	res := bytes.Compare(a, b)
	if res == 0 {
		return true
	}
	return false
}

func main() {
	clear()
	//Create game assets
	g := game.CreateGame(false)
	turn := game.CreateGame(false)
	bag := game.CreateBag(rand.New(rand.NewSource(time.Now().UnixNano())))
	mov := game.Move{}

	//Fill player hand
	Tokens, err := bag.Extract(14)
	lookError(err)
	g.Players[0].AddToken(Tokens)

	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	//Allow us to checked pressed key
	var b []byte = make([]byte, 1)

	copygame(g, turn, false)
	//game loop
	for {
		turn.Print()
		fmt.Printf("%d  %d", len(turn.Players[0].Tokens), len(bag.Bag))
		os.Stdin.Read(b)
		clear()

		if equal(b, right) {
			turn.ChangeSelected(true, 0)

		} else if equal(b, left) {
			turn.ChangeSelected(false, 0)

		} else if equal(b, space) {
			turn.SelectToken(0)
		} else if equal(b, enter) {
			if turn.Table.CheckTable() {
				if len(turn.Players[0].Tokens) == 0 {
					clear()
					fmt.Println("Jugador 0 gana")
					break
				}
				copygame(g, turn, true)
			} else {
				turn.Table.Matrix = nil
				copygame(g, turn, false)
			}
		} else if equal(b, draw) {
			Tokens, err := bag.Extract(1)
			lookError(err)
			g.Players[0].AddToken(Tokens)
			turn.Table.Matrix = nil
			copygame(g, turn, false)
			turn.ResetSelected()
		} else if equal(b, jumptotable) {
			turn.ChangeSection(true, 0)
		} else if equal(b, jumptoplayer) {
			turn.ChangeSection(false, 0)
		} else if equal(b, commit) {
			mov = turn.CreateMove(0)
			turn.Table.AddMove(mov)
			turn.ResetSelected()
		}

	}
	fmt.Println("Fin")
}

func copygame(g1, turn game.Game, valid bool) {
	if valid {
		g1.Table.Matrix = nil
		for i := 0; i < len(turn.Table.Matrix); i++ {
			g1.Table.Matrix = append(g1.Table.Matrix, turn.Table.Matrix[i])
		}
		for i := 0; i < len(turn.Players); i++ {
			g1.Players[i].Tokens = nil
			for j := 0; j < len(turn.Players[i].Tokens); j++ {
				g1.Players[i].Tokens = append(g1.Players[i].Tokens, turn.Players[i].Tokens[j])
			}
		}
	} else {

		turn.Table.Matrix = nil
		for i := 0; i < len(g1.Table.Matrix); i++ {
			turn.Table.Matrix = append(turn.Table.Matrix, g1.Table.Matrix[i])
		}
		for i := 0; i < len(g1.Players); i++ {
			turn.Players[i].Tokens = nil
			for j := 0; j < len(g1.Players[i].Tokens); j++ {
				turn.Players[i].Tokens = append(turn.Players[i].Tokens, g1.Players[i].Tokens[j])
			}
		}
	}
}
