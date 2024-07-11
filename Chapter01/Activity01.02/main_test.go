package main

import (
	"testing"

	"github.com/ibiscum/The-Go-Workshop/Chapter01/Activity01.02/swap"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSwap(t *testing.T) {
	// Initialize the mock object
	mockSwap := new(swap.MockSwap)

	// Set up the expectations
	mockSwap.On("Swap", mock.AnythingOfType("*int"), mock.AnythingOfType("*int")).Run(func(args mock.Arguments) {
		// Obtain the actual pointer to int from arguments and set the result
		aPointer := args.Get(0).(*int)
		*aPointer = 15
		bPointer := args.Get(1).(*int)
		*bPointer = 5
	}).Once()

	// Call the method with a real pointer
	a, b := 5, 15
	mockSwap.Swap(&a, &b)

	// Assertions
	assert.Equal(t, 15, a, "Expected a to be 15")
	assert.Equal(t, 5, b, "Expected b to be 5")

	mockSwap.AssertExpectations(t) // Assert that expectations were met
}
