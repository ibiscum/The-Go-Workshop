package swap

import "github.com/stretchr/testify/mock"

type MockSwap struct {
	mock.Mock
}

func (m *MockSwap) Swap(a *int, b *int) {
	m.Called(a, b)
}
