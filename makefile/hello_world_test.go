package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getText(t *testing.T) {
	assert.Equal(t, getText(), "Hello world!")

}
