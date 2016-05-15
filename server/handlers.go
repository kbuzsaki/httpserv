package server

import (
	"github.com/kbuzsaki/httpserv/http"
)

// sample hello world server and handler
type HelloWorldHandler struct {
}

func (h HelloWorldHandler) Handle(request http.Request) http.Response {
	return http.Response{http.HttpOneDotOne, http.StatusOk, "Hello World"}
}

// sample echoing server and handler
type EchoHandler struct {
}

func (h EchoHandler) Handle(request http.Request) http.Response {
	var response http.Response

	response.Protocol = http.HttpOneDotOne
	response.Status = http.StatusOk

	body := "<h1>Echo Response</h1>\n"
	body += "<p>You requested path: " + request.Path + "</p>\n"

	body += "<table><thead><th>Key</th><th>Value</th></thead><tbody>"
	for _, param := range request.Query {
		body += "<tr><td>" + param.Key + "</td><td>" + param.Val + "</td></tr>\n"
	}
	body += "</tbody></table>\n"

	response.Body = body

	return response
}
