package main

import (
	"./client"
	"./game"
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"net"
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
	if len(os.Args) != 2 {
		fmt.Print("Please pass an ip and port in the call")
		return
	}

	//Connect to the server
	clear()
	fmt.Printf("Connecting to %s\n", os.Args[1])
	con, host, err := client.Connect(os.Args[1])
	lookError(err)
	reader := bufio.NewReader(con)

	clear()
	//Create game assets
	g := game.CreateGame(false)
	turn := game.CreateGame(false)
	bag := game.CreateBag(rand.New(rand.NewSource(time.Now().UnixNano())))
	mov := game.Move{}
	yourTurn := false

	//If host create bag and send the other player
	if host {
		//Fill player hand
		Tokens, err := bag.Extract(14)
		lookError(err)
		g.Players[0].AddToken(Tokens)

		//Write to the server
		Tokens, err = bag.Extract(14)
		s := ""
		for _, v := range Tokens {
			s = s + client.TokenToString(v)
		}
		s += "\n"
		fmt.Fprintf(con, s)
		_, _ = reader.ReadString('\n')
		yourTurn = true
	} else {
		var tokens []game.Token
		data, _ := reader.ReadString('\n')
		for i := 0; i+2 < len(data); i = i + 3 {
			tokens = append(tokens, client.StringToToken(data[i:i+3]))
		}
		g.Players[0].AddToken(tokens)
		fmt.Fprintf(con, "Read\n")
	}

	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	//Allow us to checked pressed key
	var b []byte = make([]byte, 1)

	copygame(g, turn, false, con)
	//game loop
	for {
		if yourTurn {
			turn.Print()
			fmt.Printf("%d", len(turn.Players[0].Tokens))
			os.Stdin.Read(b)
			//clear()

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
					copygame(g, turn, true, con)
				} else {
					turn.Table.Matrix = nil
					copygame(g, turn, false, con)
				}
				yourTurn = false
			} else if equal(b, draw) {
				var Tokens []game.Token
				if host {
					Tokens, err = bag.Extract(1)
					lookError(err)
				} else {
					fmt.Fprintf(con, "Draw\n")
					fmt.Println("Waiting for string")
					data, _ := reader.ReadString('\n')
					fmt.Println("Got ", client.StringToToken(data[0:3]))
					token := client.StringToToken(data[0:3])
					Tokens = append(Tokens, token)
					fmt.Fprintf(con, "Done\n")
				}
				g.Players[0].AddToken(Tokens)
				turn.Table.Matrix = nil
				copygame(g, turn, false, con)
				turn.ResetSelected()
				yourTurn = false
			} else if equal(b, jumptotable) {
				turn.ChangeSection(true, 0)
			} else if equal(b, jumptoplayer) {
				turn.ChangeSection(false, 0)
			} else if equal(b, commit) {
				mov = turn.CreateMove(0)
				turn.Table.AddMove(mov)
				turn.ResetSelected()
			}

		} else {
			//clear()
			fmt.Println("Waiting turn")
			data, _ := reader.ReadString('\n')
			fmt.Println(data)
			fmt.Println("Data printed")
			if data == "Draw\n" {
				fmt.Println("Drawing")
				t, _ := bag.Extract(1)
				fmt.Fprintf(con, client.TokenToString(t[0])+"\n")
				_, _ = reader.ReadString('\n')
				yourTurn = true
			} else {
				data, _ := reader.ReadString('\n')
				g.Table = client.StringToTable(data)
				yourTurn = true
			}
			//clear()
		}
	}
	fmt.Fprintf(con, "STOP\n")
}

func copygame(g1, turn game.Game, valid bool, con net.Conn) {
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

	fmt.Fprintf(con, client.TableToString(g1.Table))
}
