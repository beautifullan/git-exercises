package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFabonacci(t *testing.T) {
	assert.Equal(t, 1, Fibonacci(0))
	assert.Equal(t, 1, Fibonacci(1))
	assert.Equal(t, 2, Fibonacci(2))
	assert.Equal(t, 3, Fibonacci(3))
	assert.Equal(t, 5, Fibonacci(4))
	assert.Equal(t, 89, Fibonacci(10))
}
