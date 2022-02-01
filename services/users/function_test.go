package users

import "testing"

func TestEmailValidation(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expectedOut bool
	}{
		{
			desc:     "Success case",
			input:    "gopi@gmail.com",
			expectedOut: true,
		},
		{
			desc:     "Success case",
			input:    "gopi@gmail",
			expectedOut: false,
		},
	}

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			out := emailValidation(ts.input)
			if out != ts.expectedOut {
				t.Errorf("Expected %v Obtained %v", ts.expectedOut, out)
			}
		})
	}
}

func TestCheckId(t *testing.T) {
	testCases := []struct {
		desc     string
		input    int
		expecOut bool
	}{
		{
			desc:     "Success case",
			input:    1,
			expecOut: true,
		},
		{
			desc:     "Success case",
			input:    0,
			expecOut: false,
		},
	}
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			out := idCheck(ts.input)
			if out != ts.expecOut {
				t.Errorf("Expected %v Obtained %v", ts.expecOut, out)
			}
		})
	}
}
