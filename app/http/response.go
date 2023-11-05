package http

import "fmt"

const HTTP1_1 = "HTTP/1.1"
const CRLF = "\r\n"

func NewResponse(statusCode int, httpVersion, message string) Response {
	return Response{
		StatusCode:  statusCode,
		HttpVersion: httpVersion,
		Message:     message,
	}
}

type Response struct {
	StatusCode  int
	HttpVersion string
	Message     string
}

func (r *Response) Bytes() []byte {
	return []byte(r.String())
}

func (r *Response) String() string {
	var respStr string

	// HTTP status line
	respStr += fmt.Sprintf("%s %d %s%s", HTTP1_1, r.StatusCode, r.Message, CRLF)

	// Response headers section
	respStr += CRLF

	return respStr

}
