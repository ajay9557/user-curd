package users

import "testing"

func TestEmailValid(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expectedOut bool
	}{
		{
			desc:     "success case",
			input:    "sudheer@gmail.com",
			expectedOut: true,
		},
		{
			desc:     "failure case",
			input:    "sudheer",
			expectedOut: false,
		},
		{
			desc:     "failure case",
			input:    "su",
			expectedOut: false,
		},
	}
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			result := isEmailValid(ts.input)
			if result != ts.expectedOut {
				t.Errorf("Expected %v obtained %v", ts.expectedOut, result)
			}
		})
	}
}
