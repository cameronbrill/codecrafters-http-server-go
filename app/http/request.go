package http

import (
	"errors"
	"fmt"
	"strings"
)

type Request struct {
	Method      string
	Path        string
	HttpVersion string
	Headers     map[string]string
}

func (r Request) BuildResponse() Response {
	resp := NewResponse(
		200,
		r.HttpVersion,
		"OK",
	)

	fmt.Println("identifying path")
	pathIdentifier, pathItems, err := identifyPath(r.Path)
	if err != nil {
		if errors.Is(err, PathNotFoundError) || errors.Is(err, EmptyPathError) {
			resp.StatusCode = 404
			resp.Message = "Not Found"
		}
	}

	if pathIdentifier == EchoPath {
		resp.SetBody(pathItems)
	}

	return resp
}

func (r *Request) parseStartLine(line string) error {
	parts := strings.Split(line, " ")

	if len(parts) != 3 {
		return errors.New("invalid start line format")
	}

	r.Method = parts[0]
	r.Path = parts[1]
	r.HttpVersion = parts[2]
	return nil
}

func ParseRequestBuffer(reqBuf []byte) (Request, error) {
	var r *Request = &Request{}

	reqStr := string(reqBuf)
	fmt.Println("parsing request string: ", reqStr)
	reqLines := strings.Split(reqStr, CRLF)
	for i, reqLine := range reqLines {
		if i == 0 {
			err := r.parseStartLine(reqLine)
			if err != nil {
				return *r, fmt.Errorf("parsing start line: %w", err)
			}
		}
	}

	return *r, nil
}
