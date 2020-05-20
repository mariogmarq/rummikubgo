package main

import (
	"./game"
	"bytes"
	"fmt"
	"os"
	"os/exec"
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
	mov := game.Move{}

	//Fill player hand
	Tokens, err := g.Bag.Extract(14)
	lookError(err)
	g.Players[0].AddToken(Tokens)

	//NOT MINE
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	//NOT MINE

	var b []byte = make([]byte, 1)

	//game loop
	for {
		g.Print()
		fmt.Println(mov.Score)
		os.Stdin.Read(b)
		clear()

		if equal(b, right) {
			g.ChangeSelected(true, 0)

		} else if equal(b, left) {
			g.ChangeSelected(false, 0)

		} else if equal(b, space) {
			g.SelectToken(0) //Cambiar

		} else if equal(b, enter) {
			g.EndTurn()

		} else if equal(b, draw) {
			Tokens, err := g.Bag.Extract(1)
			lookError(err)
			g.Players[0].AddToken(Tokens)
			g.Players[0].ResetSelected()
		} else if equal(b, jumptotable) {
			g.ChangeSection(true, 0)
		} else if equal(b, jumptoplayer) {
			g.ChangeSection(false, 0)
		} else if equal(b, commit) {
			mov = g.CreateMove(0)
			g.Turn = append(g.Turn, mov)
		}

	}
}
