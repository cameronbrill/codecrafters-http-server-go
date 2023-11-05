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
		fmt.Println("ERROR Failed to bind to port 4221")
		os.Exit(1)
	}

	defer func() {
		fmt.Println("closing tcp listener")
		err = l.Close()
		if err != nil {
			fmt.Println("ERROR closing listener: ", err.Error())
			os.Exit(1)
		}
	}()

	for {
		fmt.Println("waiting for connection")
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR accepting connection: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("handling connection")
		err = http.HandleConnection(conn)
		if err != nil {
			fmt.Println("ERROR handling connection: ", err.Error())
		}
	}

}
