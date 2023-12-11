package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFabonacci(t *testing.T) {
	assert.Equal(t, 0, Fibonacci(0))
	assert.Equal(t, 1, Fibonacci(1))
	assert.Equal(t, 1, Fibonacci(2))
	assert.Equal(t, 2, Fibonacci(3))
	assert.Equal(t, 3, Fibonacci(4))
	assert.Equal(t, 55, Fibonacci(10))
	//this will be very slow if we use recursive call, is there any way to improve it ?
	assert.Equal(t, 4181, Fibonacci(19))
	assert.Equal(t, 12586269025, Fibonacci(50))
}
