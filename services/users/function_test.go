package users

import "testing"

func TestEmailValid(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expecOut bool
	}{
		{
			desc:     "success case",
			input:    "sudheer@gmail.com",
			expecOut: true,
		},
		{
			desc:     "failure case",
			input:    "sudheer",
			expecOut: false,
		},
		{
			desc:     "failure case",
			input:    "su",
			expecOut: false,
		},
	}
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			result := isEmailValid(ts.input)
			if result != ts.expecOut {
				t.Errorf("Expected %v obtained %v", ts.expecOut, result)
			}
		})
	}
}
