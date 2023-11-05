package http

import (
	"io"
	"net"

	"github.com/pkg/errors"
)

func HandleConnection(conn net.Conn) error {
	defer conn.Close()

	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return errors.Wrap(err, "read error")
			}
			break
		}
		buf = append(buf, tmp[:n]...)
	}

	resp := NewResponse(200, "GET", "OK")
	_, err := conn.Write(resp.Bytes())
	if err != nil {
		return errors.Wrap(err, "writing response")
	}

	return nil

}
