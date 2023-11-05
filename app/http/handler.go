package http

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"errors"
)

func readData(conn net.Conn) ([]byte, error) {
	fmt.Println("reading data from connection")
	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
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
			return nil, fmt.Errorf("reading response chunk: %w", err)
		}
		fmt.Println("appending data chunk to buffer")
		buf = append(buf, tmp[:n]...)
	}

	return buf, nil
}

func HandleConnection(conn net.Conn) error {
	defer conn.Close()

	fmt.Println("configuring connection read deadline")
	err := conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		return fmt.Errorf("setting connection read deadline: %w", err)
	}

	fmt.Println("reading data from connection")
	reqBuf, err := readData(conn)
	if err != nil {
		return fmt.Errorf("reading data from tcp connection: %w", err)
	}

	fmt.Println("parsing request data")
	req, err := ParseRequestBuffer(reqBuf)
	if err != nil {
		return fmt.Errorf("parsing request buffer: %w", err)
	}

	fmt.Println("validating request & building response")
	resp := req.BuildResponse()

	fmt.Println("writing data to connection")
	_, err = conn.Write(resp.Bytes())
	if err != nil {
		return fmt.Errorf("writing response: %w", err)
	}

	return nil

}
