package testifypackage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, Add(2, 2), 4)
}
func TestSub(t *testing.T) {
	assert.Equal(t, Sub(2, 2), 0)
}
func TestMul(t *testing.T) {
	assert.Equal(t, Mul(2, 3), 6)
}
func TestDiv(t *testing.T) {
	assert.Equal(t, Div(2, 2), 1)
}
