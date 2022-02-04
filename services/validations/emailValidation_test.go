package validations

import "testing"

func TestEmailValid(t *testing.T) {
	testCases := []struct {
		desc           string
		input          string
		expectedOutput bool
	}{
		{
			desc:           "success case",
			input:          "Nayani@gmail.com",
			expectedOutput: true,
		},
		{
			desc:           "failure case",
			input:          "Nayani",
			expectedOutput: false,
		},
		{
			desc:           "failure case",
			input:          "na",
			expectedOutput: false,
		},
	}
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			result := ValidateEmail(ts.input)
			if result != ts.expectedOutput {
				t.Errorf("Expected %v obtained %v", ts.expectedOutput, result)
			}
		})
	}
}
