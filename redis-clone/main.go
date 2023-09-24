// this will work on windows with WSL however its more of a headache than i expected.
// i will finish the project and figure out how to get this to work even better
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Listening for port :6379")

	// Create new server on port 6379 return err if unsuccessful
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for port 6379 return err if unsuccessful
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	// closing connection
	defer conn.Close()

	// infinite loop to read client msg, then respond with OK
	for {
		buf := make([]byte, 1024)

		// reading client message
		_, err := conn.Read(buf)
		if err != nil {
			// if error is true then we reach end of file and terminate loop
			if err == io.EOF {
				break
			}
			// otherwise print error and exit with status 1
			fmt.Println("error reading command: ", err.Error())
			os.Exit(1)
		}

		// ignore request and return PONG
		conn.Write([]byte("+OK\r\n"))
	}
}
