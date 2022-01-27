package services

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_IsUniqueEmail(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockService := NewMockServices(controller)

	testCase := []struct {
		desc     string
		input    string
		expected bool
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case 1",
			input:    "vicky542011@gmail.com",
			expected: true,
			mock:     []*gomock.Call{mockService.EXPECT().IsUniqueEmail("vicky542011@gmail.com").Return(true)},
		},
		{
			desc:     "Test Case 2",
			input:    "vicky542011gmail.com",
			expected: false,
			mock:     []*gomock.Call{mockService.EXPECT().IsUniqueEmail("vicky542011gmail.com").Return(false)},
		},
		{
			desc:     "Test Case 3",
			input:    "vicky.com",
			expected: false,
			mock:     []*gomock.Call{mockService.EXPECT().IsUniqueEmail("vicky.com").Return(false)},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			result := mockService.IsUniqueEmail(tcs.input)
			if result != tcs.expected {
				t.Errorf("output: %v, expected: %v", result, tcs.expected)
			}
		})
	}
}

func Test_IsNumberValid(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockService := NewMockServices(controller)

	testCase := []struct {
		desc     string
		input    string
		expected bool
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case 1",
			input:    "7828789845",
			expected: true,
			mock:     []*gomock.Call{mockService.EXPECT().IsNumberValid("7828789845").Return(true)},
		},
		{
			desc:     "Test Case 2",
			input:    "64832648263423427346823",
			expected: false,
			mock:     []*gomock.Call{mockService.EXPECT().IsNumberValid("64832648263423427346823").Return(false)},
		},
		{
			desc:     "Test Case 3",
			input:    "2323232",
			expected: false,
			mock:     []*gomock.Call{mockService.EXPECT().IsNumberValid("2323232").Return(false)},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			result := mockService.IsNumberValid(tcs.input)
			if result != tcs.expected {
				t.Errorf("output: %v, expected: %v", result, tcs.expected)
			}
		})
	}
}
