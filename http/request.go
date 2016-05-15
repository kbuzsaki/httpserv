package http

import (
	"bufio"
	"io"
	"strings"
)

type QueryParam struct {
	Key string
	Val string
}

type Request struct {
	lines   []string
	Method  string
	Uri     string
	Path    string
	Query   []QueryParam
	Version string
}

func (request *Request) ParseRequestLine(line string) {
	segments := strings.Split(line, " ")
	request.Method = segments[0]
	request.Uri = segments[1]
	request.Version = segments[2]

	uriSegments := strings.Split(request.Uri, "?")
	request.Path = uriSegments[0]

	if len(uriSegments) > 1 {
		querySegments := strings.Split(uriSegments[1], "&")
		for _, querySegment := range querySegments {
			// TODO: url decode the parameters
			paramSegments := strings.Split(querySegment, "=")
			param := QueryParam{paramSegments[0], paramSegments[1]}
			request.Query = append(request.Query, param)
		}
	}
}

func (request *Request) String() string {
	return strings.Join(request.lines, "\n")
}

func ReadRequest(reader io.Reader) (Request, error) {
	request := Request{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// read until the first empty line, which signals the end of the request
		line := scanner.Text()
		if line == "" {
			break
		}

		request.lines = append(request.lines, line)

		// hacky way to check if we've set the method yet
		if request.Method == "" {
			request.ParseRequestLine(line)
		}
	}

	err := scanner.Err()
	if err != nil && err != io.EOF {
		return request, err
	}

	return request, nil
}
