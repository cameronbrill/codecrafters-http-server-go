package http

import "fmt"

const HTTP1_1 = "HTTP/1.1"
const CRLF = "\r\n"

func NewResponse(statusCode int, httpVersion, message string) Response {
	return Response{
		StatusCode:  statusCode,
		HttpVersion: httpVersion,
		Message:     message,
		body:        "",
		headers:     map[string]string{},
	}
}

type Response struct {
	StatusCode  int
	HttpVersion string
	Message     string
	body        string
	headers     map[string]string
}

func (r *Response) SetBody(body string) {
	r.body = body
	r.headers["Content-Length"] = fmt.Sprintf("%d", len(body))
	r.headers["Content-Type"] = "text/plain"
}

func (r *Response) Bytes() []byte {
	return []byte(r.String())
}

func (r *Response) String() string {
	var respStr string

	// HTTP status line
	respStr += fmt.Sprintf("%s %d %s%s", HTTP1_1, r.StatusCode, r.Message, CRLF)

	// Response headers section
	for k, v := range r.headers {
		respStr += fmt.Sprintf("%s: %s%s", k, v, CRLF)
	}

	// Response Body
	respStr += fmt.Sprintf("%s%s", CRLF, r.body)

	return respStr

}
