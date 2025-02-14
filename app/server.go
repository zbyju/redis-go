package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/app/logic"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	fmt.Println("Handling new connection")
	for {
		buffer := make([]byte, 256)
		len, err := c.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection: ", err.Error())
			return
		}

		input := string(buffer[:len])
		cmd, err := logic.ParseCommand(input)
		if err != nil {
			fmt.Println("Error parsing the command: ", err.Error())
			return
		}

		res, err := cmd.Execute()
		if err != nil {
			fmt.Println("Error executing the command: ", err.Error())
			return
		}

		c.Write([]byte(fmt.Sprintf("+%s\r\n", *res)))
	}
}
