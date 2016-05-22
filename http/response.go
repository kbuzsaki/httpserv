package http

import (
	"io"
	"strconv"
)

type Protocol struct {
	Name    string
	Version string
}

var HttpOneDotOne = Protocol{"HTTP", "1.1"}

func (protocol *Protocol) String() string {
	return protocol.Name + "/" + protocol.Version
}

type Status struct {
	Code int
	Name string
}

func (status *Status) String() string {
	return strconv.Itoa(status.Code) + " " + status.Name
}

func (status *Status) WriteTo(writer io.Writer) (int64, error) {
	n, err := writer.Write([]byte(status.String() + "\n"))
	return int64(n), err
}

var StatusOk Status = Status{200, "OK"}
var StatusMovedTemporarily = Status{302, "Found"}
var StatusNotFound Status = Status{404, "Not Found"}
var StatusInternalError = Status{500, "Internal Server Error"}

type ResponseHeader struct {
	Key string
	Val string
}

type Response struct {
	Protocol Protocol
	Status   Status
	Headers  []ResponseHeader
	Body     string
}

func MakeSimpleResponse(body string) Response {
	return Response{HttpOneDotOne, StatusOk, nil, body}
}

func MakeErrorResponse(status Status, err error) Response {
	body := "<h1>" + status.String() + "</h1>" + err.Error()
	return Response{HttpOneDotOne, status, nil, body}
}

func MakeRedirectResponse(path string) Response {
	headers := []ResponseHeader{{"Location", path}}
	return Response{HttpOneDotOne, StatusMovedTemporarily, headers, ""}
}
