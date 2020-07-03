package client

import (
	"bufio"
	"fmt"
	"net"
)

//Connect to the server and waits for the match to start, it will return if you are the host or not
func Connect(port string) (net.Conn, bool, error) {
	con, err := net.Dial("tcp4", port)
	if err != nil {
		return nil, false, err
	}
	fmt.Println("Connected")
	reader := bufio.NewReader(con)
	fmt.Println("Waiting for msg")
	netdata, _ := reader.ReadString('\n')
	fmt.Println("recieved")
	fmt.Println(netdata)
	if netdata == "HOST\n" {
		_, _ = reader.ReadString('\n')
		return con, true, nil
	} else {
		return con, false, nil
	}

}
