package Users

import "testing"

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		desc     string
		email    string
		expected bool
	}{
		{
			desc:     "Case1",
			email:    "ridhdhish@gmail.com",
			expected: true,
		},
		{
			desc:     "Case2",
			email:    "ridhdhish@.com",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			isValid := validateEmail(test.email)
			if isValid != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, isValid)
			}
		})
	}
}

func TestValidatePhone(t *testing.T) {
	tests := []struct {
		desc     string
		phone    string
		expected bool
	}{
		{
			desc:     "Case1",
			phone:    "8320578360",
			expected: true,
		},
		{
			desc:     "Case2",
			phone:    "67896",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			isValid := validatePhone(test.phone)
			if isValid != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, isValid)
			}
		})
	}
}
