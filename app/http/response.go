package http

import "fmt"

const HTTP1_1 = "HTTP/1.1"
const CRLF = "\r\n"

func NewResponse(statusCode int, method, message string) Response {
	return Response{
		StatusCode: statusCode,
		Method:     method,
		Message:    message,
	}
}

type Response struct {
	StatusCode int
	Method     string
	Message    string
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
