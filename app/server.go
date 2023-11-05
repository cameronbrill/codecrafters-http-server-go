package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/http"
)

func main() {
	fmt.Println("starting server")
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer func() {
		err = l.Close()
		if err != nil {
			fmt.Println("closing listener: ", err.Error())
			os.Exit(1)
		}
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		err = http.HandleConnection(conn)
		if err != nil {
			fmt.Println("handling connection: ", err.Error())
			os.Exit(1)
		}
	}

}
