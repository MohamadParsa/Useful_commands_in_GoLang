package main

import (
	"io"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_concurrentPrint(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	text := "expected text"
	go concurrentPrint(text, &waitGroup)
	waitGroup.Wait()

	w.Close()
	output, _ := io.ReadAll(r)

	os.Stdout = stdout

	expected := text + "\n"
	assert.Equal(t, string(output), expected)

}
