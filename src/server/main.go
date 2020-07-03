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
		fmt.Fprintf(con1, "HOST\n")
		fmt.Println(">> User connected")
		con2, _ := listener.Accept()
		fmt.Fprintf(con1, "PLAYERFOUND\n")
		fmt.Fprintf(con2, "GUEST\n")
		fmt.Println(">> Second user connected")
		go handleMatch(con1, con2)
	}
}

func handleMatch(con1, con2 net.Conn) {
	fmt.Println(">> match started")
	//Firts create the readers and writters for the connections
	reader1 := bufio.NewReader(con1)
	reader2 := bufio.NewReader(con2)

	//Now the server will be a simple bridge between the both clients
	for {
		netdata, _ := reader1.ReadString('\n')
		fmt.Fprintf(con2, netdata)
		if netdata == "STOP\n" {
			return
		}
		netdata, _ = reader2.ReadString('\n')
		fmt.Fprintf(con1, netdata)
		if netdata == "STOP\n" {
			return
		}

	}
}
