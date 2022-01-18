package tailor_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	tailor "github.com/yitsushi/file-tailor"
)

func Test_Tail_oneLine(t *testing.T) {
	file, err := os.Open("fixtures/normal")
	if !assert.NoError(t, err) {
		return
	}

	defer file.Close()

	output, err := tailor.Tail(file, 1)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "line 4\n", string(output))
}

func Test_Tail_twoLines(t *testing.T) {
	file, err := os.Open("fixtures/normal")
	if !assert.NoError(t, err) {
		return
	}

	defer file.Close()

	output, err := tailor.Tail(file, 2)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "line 3\nline 4\n", string(output))
}

func Test_Tail_overRead(t *testing.T) {
	file, err := os.Open("fixtures/normal")
	if !assert.NoError(t, err) {
		return
	}

	defer file.Close()

	output, err := tailor.Tail(file, 10)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "line 1\nline 2\nline 3\nline 4\n", string(output))
}

func Test_Tail_noNewLineAtTheEnd(t *testing.T) {
	file, err := os.Open("fixtures/nonl")
	if !assert.NoError(t, err) {
		return
	}

	defer file.Close()

	output, err := tailor.Tail(file, 1)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "line 3", string(output))
}

func Test_Tail_emptyFile(t *testing.T) {
	file, err := os.Open("fixtures/empty")
	if !assert.NoError(t, err) {
		return
	}

	defer file.Close()

	output, err := tailor.Tail(file, 1)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "", string(output))
}
