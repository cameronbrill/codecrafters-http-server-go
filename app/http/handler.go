package http

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"errors"
)

func HandleConnection(conn net.Conn) error {
	defer conn.Close()

	fmt.Println("reading data from connection")
	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	err := conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		return fmt.Errorf("setting connection read deadline: %w", err)
	}

	for {

		fmt.Println("reading data chunk")
		n, err := conn.Read(tmp)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Printf("got EOF, end of reading data from connection")
				break
			}
			if errors.Is(err, os.ErrDeadlineExceeded) {
				fmt.Printf("hit read deadline, end of reading data from connection")
				break
			}
			return fmt.Errorf("reading response chunk: %w", err)
		}
		fmt.Println("appending data chunk to buffer")
		buf = append(buf, tmp[:n]...)
	}

	resp := NewResponse(200, "GET", "OK")
	fmt.Println("writing data to connection")
	_, err = conn.Write(resp.Bytes())
	if err != nil {
		return fmt.Errorf("writing response: %w", err)
	}

	return nil

}
