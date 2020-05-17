package main

import (
	"./game"
	"bytes"
	"os"
	"os/exec"
)

//bytes for moving
var left = []byte{'a'}
var right = []byte{'d'}

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
	bag := game.CreateBag()
	player := game.CreatePlayer()

	//Fill player hand
	Tokens, err := bag.Extract(14)
	lookError(err)
	player.AddToken(Tokens)

	//NOT MINE
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	//NOT MINE

	var b []byte = make([]byte, 1)
	//game loop
	for {
		player.Print()
		os.Stdin.Read(b)
		clear()
		if equal(b, right) {
			player.ChangeSelected(true)
		} else if equal(b, left) {
			player.ChangeSelected(false)
		}
	}
}
