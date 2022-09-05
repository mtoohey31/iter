package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	next, ok := Empty[bool]()()
	assert.Zero(t, next)
	assert.False(t, ok)
}
