package user

import (
	"testing"
)

func Test_IsUniqueEmail(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			desc:     "Test Case 1",
			input:    "vicky54@gmail.com",
			expected: true,
		},
		{
			desc:     "Test Case 2",
			input:    "vicky54.com",
			expected: false,
		},
		{
			desc:     "Test Case 3",
			input:    "vicky54gmail.com",
			expected: false,
		},
		{
			desc:     "Test Case 4",
			input:    "vicky54@gmail",
			expected: false,
		},
	}
	for _, tcs := range testCases {
		t.Run(tcs.desc, func(t *testing.T) {
			output := IsUniqueEmail(tcs.input)
			if output != tcs.expected {
				t.Errorf("Output: %v Expected: %v", output, tcs.expected)
			}
		})
	}
}

func Test_IsNumberValid(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			desc:     "Test Case 1",
			input:    "7828789845",
			expected: true,
		},
		{
			desc:     "Test Case 2",
			input:    "432342",
			expected: false,
		},
		{
			desc:     "Test Case 3",
			input:    "vickyfadsfa",
			expected: false,
		},
		{
			desc:     "Test Case 4",
			input:    "534534534534534535345",
			expected: false,
		},
	}
	for _, tcs := range testCases {
		t.Run(tcs.desc, func(t *testing.T) {
			output := IsNumberValid(tcs.input)
			if output != tcs.expected {
				t.Errorf("Output: %v Expected: %v", output, tcs.expected)
			}
		})
	}
}

func Test_IsValidId(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			desc:     "Test Case 1",
			input:    "12",
			expected: true,
		},
		{
			desc:     "Test Case 2",
			input:    "ff",
			expected: false,
		},
	}
	for _, tcs := range testCases {
		t.Run(tcs.desc, func(t *testing.T) {
			output := IsValidId(tcs.input)
			if output != tcs.expected {
				t.Errorf("Output: %v Expected: %v", output, tcs.expected)
			}
		})
	}
}
