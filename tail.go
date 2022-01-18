package tailor

import (
	"io"
	"os"
)

// Tail reads the last `limit` lines of a file.
func Tail(file *os.File, limit int) ([]byte, error) {
	_, _ = file.Seek(-1, io.SeekEnd)

	content := []byte{}

	for {
		char := make([]byte, 1)

		_, err := file.Read(char)
		if err != nil && err != io.EOF {
			return content, err
		}

		if isTermination(char[0]) {
			if limit == 0 {
				break
			}

			limit--
		}

		if len(content) == 0 && !isTermination(char[0]) {
			limit--
		}

		content = append(char, content...)

		if head, err := file.Seek(-2, io.SeekCurrent); head == 0 && err != nil {
			break
		}
	}

	// If the file is empty, it will yield with a single null byte. In this
	// case we want an empty response.
	if len(content) == 1 && content[0] == 0x0 {
		content = []byte{}
	}

	return content, nil
}

func isTermination(char byte) bool {
	return char == '\n' || char == '\r' || char == 0
}
