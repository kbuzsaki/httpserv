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
var StatusNotFound Status = Status{404, "Not Found"}

type Response struct {
	Protocol Protocol
	Status   Status
	Body     string
}
