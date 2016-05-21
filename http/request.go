package http

import (
	"bufio"
	"errors"
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
			paramSegments := strings.Split(querySegment, "=")

			// TODO: url decode the parameters
			param := QueryParam{paramSegments[0], ""}
			if len(paramSegments) > 1 {
				param.Val = paramSegments[1]
			}

			request.Query = append(request.Query, param)
		}
	}
}

func (request *Request) Param(key string) QueryParam {
	for _, param := range request.Query {
		if param.Key == key {
			return param
		}
	}

	return QueryParam{}
}

func (request *Request) String() string {
	return strings.Join(request.lines, "\n")
}

func ReadRequest(reader io.Reader) (Request, error) {
	request := Request{}

	scanner := bufio.NewScanner(reader)

	// first read the request line
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return request, err
		} else {
			return request, errors.New("Empty request")
		}
	}
	line := scanner.Text()
	request.lines = append(request.lines, line)
	request.ParseRequestLine(line)

	// then read each of the headers
	for scanner.Scan() {
		// read until the first empty line, which signals the end of the headers
		line = scanner.Text()
		if line == "" {
			break
		}

		request.lines = append(request.lines, line)
	}

	err := scanner.Err()
	if err != nil && err != io.EOF {
		return request, err
	}

	return request, nil
}
