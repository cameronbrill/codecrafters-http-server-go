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

var PathNotFoundError = errors.New("path not found")
var EmptyPathError = errors.New("path cannot be empty")

func validatePath(path string) error {
	if path == "" {
		return fmt.Errorf("%w: %s", EmptyPathError, path)
	}
	if path != "/" {
		return fmt.Errorf("%w: %s", PathNotFoundError, path)
	}

	return nil
}

func (r Request) BuildResponse() Response {
	resp := Response{
		HttpVersion: r.HttpVersion,
		StatusCode:  200,
		Message:     "OK",
	}
	err := validatePath(r.Path)
	if err != nil {
		if errors.Is(err, PathNotFoundError) {
			resp.StatusCode = 404
			resp.Message = "Not Found"
		}
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
