package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//The port will be passed as an argument
	if len(os.Args) != 2 {
		fmt.Println("No port passed as argument for the server")
		return
	}
	PORT := os.Args[1]

	//Create the listener to the port, with tcp4 protocol
	listener, err := net.Listen("tcp4", PORT)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println(">> Listening to ", PORT)

	//Will get two connections and handle them, the purpose of the server is being concurrent
	for {
		con1, _ := listener.Accept()
		w1 := bufio.NewWriter(con1)
		fmt.Fprintf(w1, "Connected, waiting for second player")
		fmt.Println(">> User connected")
		con2, _ := listener.Accept()
		fmt.Fprintf(w1, "Match found!")
		w2 := bufio.NewWriter(con2)
		fmt.Fprintf(w2, "Host found!")
		fmt.Println(">> Second user connected")
		go HandleMatch(con1, con2)
	}
}

func HandleMatch(con1, con2 net.Conn) {
	fmt.Println(">> match started")
	//Firts create the reader for the con1 and the writer for con2
	reader := bufio.NewReader(con1)
	writer := bufio.NewWriter(con2)

	//Now the server will be a simple bridge between the both clients
	for {
		netdata, _ := reader.ReadString('\n')
		fmt.Fprintf(writer, netdata)
		if netdata == "STOP\n" {
			return
		}
	}
}
