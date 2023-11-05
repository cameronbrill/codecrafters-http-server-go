package http

import (
	"fmt"
	"io"
	"net"

	"errors"
)

func HandleConnection(conn net.Conn) error {
	defer conn.Close()

	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return fmt.Errorf("reading response chunk: %w", err)
			}
			break
		}
		buf = append(buf, tmp[:n]...)
	}

	resp := NewResponse(200, "GET", "OK")
	_, err := conn.Write(resp.Bytes())
	if err != nil {
		return fmt.Errorf("writing response: %w", err)
	}

	return nil

}
