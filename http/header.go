package http

import (
	"bufio"
	"io"
	"strings"
)

type Header struct {
	lines []string
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
	}

	err := scanner.Err()
	if err != nil && err != io.EOF {
		return header, err
	}

	return header, nil
}
