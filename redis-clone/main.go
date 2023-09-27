// WSL isn't needed. I used the archived Microsoft Redis app
// https://github.com/microsoftarchive/redis
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Listening for port :6379")

	// Create new server on port 6379 return err if unsuccessful
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for port 6379
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	// closing connection
	defer conn.Close()

	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		if value.typ != "array" {
			fmt.Println("Invalid request, expected Array")
			continue
		}

		if len(value.array) == 0 {
			fmt.Println("Invalid request, expected Array length > 0")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		writer := NewWriter(conn)

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}
		result := handler(args)
		writer.Write(result)
	}

}

// // infinite loop to read client msg, then respond with OK
// for {
// 	buf := make([]byte, 1024)

// 	// reading client message
// 	_, err := conn.Read(buf)
// 	if err != nil {
// 		// if error is true then we reach end of file and terminate loop
// 		if err == io.EOF {
// 			break
// 		}
// 		// otherwise print error and exit with status 1
// 		fmt.Println("error reading command: ", err.Error())
// 		os.Exit(1)
// 	}

// 	// ignore request and return OK
// 	conn.Write([]byte("+OK\r\n"))
// }
