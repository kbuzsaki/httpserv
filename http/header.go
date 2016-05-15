package http

import (
	"bufio"
	"io"
	"strings"
)

type Header struct {
	lines   []string
	Method  string
	Uri     string
	Path    string
	Query   map[string]string
	Version string
}

func (header *Header) ParseRequestLine(line string) {
	segments := strings.Split(line, " ")
	header.Method = segments[0]
	header.Uri = segments[1]
	header.Version = segments[2]

	uriSegments := strings.Split(header.Uri, "?")
	header.Path = uriSegments[0]

	if len(uriSegments) > 1 {
		header.Query = make(map[string]string)
		querySegments := strings.Split(uriSegments[1], "&")
		for _, param := range querySegments {
			paramSegments := strings.Split(param, "=")
			header.Query[paramSegments[0]] = paramSegments[1]
		}
	}
}

func (header *Header) String() string {
	return strings.Join(header.lines, "\n")
}

func ReadHeader(reader io.Reader) (Header, error) {
	header := Header{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// read until the first empty line, which signals the end of the header
		line := scanner.Text()
		if line == "" {
			break
		}

		header.lines = append(header.lines, line)

		// hacky way to check if we've set the method yet
		if header.Method == "" {
			header.ParseRequestLine(line)
		}
	}

	err := scanner.Err()
	if err != nil && err != io.EOF {
		return header, err
	}

	return header, nil
}
