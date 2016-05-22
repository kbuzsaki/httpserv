package server

import (
	"encoding/json"
	"github.com/kbuzsaki/httpserv/http"
	"io/ioutil"
	"os"
	"path"
	"strings"
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
	filePath := path.Join(h.BasePath, request.Path)

	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return http.MakeErrorResponse(http.StatusNotFound, err)
		} else {
			return http.MakeErrorResponse(http.StatusInternalError, err)
		}
	} else if info.IsDir() {
		return h.serveDirectory(request, filePath)
	} else {
		return h.serveFile(request, filePath)
	}
}

func (h StaticFileHandler) serveFile(request http.Request, filePath string) http.Response {
	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		return http.MakeErrorResponse(http.StatusInternalError, err)
	} else {
		return http.MakeSimpleResponse(string(bytes))
	}
}

func (h StaticFileHandler) serveDirectory(request http.Request, filePath string) http.Response {
	if !strings.HasSuffix(request.Path, "/") {
		return http.MakeRedirectResponse("/" + filePath + "/")
	}

	indexPath := path.Join(filePath, "index.html")
	if indexInfo, err := os.Stat(indexPath); err == nil && !indexInfo.IsDir() {
		return h.serveFile(request, indexPath)
	}

	entries, err := ioutil.ReadDir(filePath)
	if err != nil {
		return http.MakeErrorResponse(http.StatusInternalError, err)
	}

	body := "<h1>Contents of " + request.Path + "</h1>"

	body += "<ul>"
	for _, entry := range entries {
		href := request.Path + entry.Name()
		body += "<li><a href=\"" + href + "\">" + entry.Name() + "</a></li>"
	}
	body += "</ul>"

	return http.MakeSimpleResponse(body)
}
