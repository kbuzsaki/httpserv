package server

import (
	"encoding/json"
	"github.com/kbuzsaki/httpserv/http"
	"io/ioutil"
	"path"
)

// sample hello world server and handler
type HelloWorldHandler struct {
}

func (h HelloWorldHandler) Handle(request http.Request) http.Response {
	return http.MakeSimpleResponse("Hello World")
}

// sample echoing server and handler
type EchoHandler struct {
}

func (h EchoHandler) Handle(request http.Request) http.Response {
	body := "<h1>Echo Response</h1>\n"
	body += "<p>You requested path: " + request.Path + "</p>\n"

	requestJsonBytes, _ := json.MarshalIndent(&request, "", "    ")
	body += "<pre>" + string(requestJsonBytes) + "</pre>"

	body += "<table><thead><th>Key</th><th>Value</th></thead><tbody>"
	for _, param := range request.Query {
		body += "<tr><td>" + param.Key + "</td><td>" + param.Val + "</td></tr>\n"
	}
	body += "</tbody></table>\n"

	return http.MakeSimpleResponse(body)
}

// sample file serving handler
type StaticFileHandler struct {
	BasePath string
}

func (h StaticFileHandler) Handle(request http.Request) http.Response {
	var response http.Response
	response.Protocol = http.HttpOneDotOne

	filePath := path.Join(h.BasePath, request.Path)
	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		response.Status = http.StatusNotFound
		response.Body = err.Error()
	} else {
		response.Status = http.StatusOk
		response.Body = string(bytes)
	}

	return response
}
