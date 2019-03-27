package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, Calculate(2), 4)
}

func TestCalculateComplex(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}

// smsServiceMock
type smsServiceMock struct {
	mock.Mock
}

func (m *smsServiceMock) SendChargeNotification(value int) bool {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)
	args := m.Called(value)
	return args.Bool(0)
}

func (m *smsServiceMock) DummyFunc() {
	fmt.Println("Dummy")
}

func TestChargeCustomer(t *testing.T) {
	smsService := new(smsServiceMock)
	smsService.On("SendChargeNotification", 100).Return(true)

	myService := MyService{smsService}
	myService.ChargeCustomer(100)

	smsService.AssertExpectations(t)
}
