package users

import (
	"strconv"
	"testing"
)

func Test_validateId(t *testing.T) {
	testCases := []struct {
		caseId      int
		input       int
		expectedOut bool
	}{
		// Success case
		{
			caseId:      1,
			input:       1,
			expectedOut: true,
		},
		// Error cases
		{
			caseId:      2,
			input:       -1,
			expectedOut: false,
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			out := validateId(tc.input)
			if out != tc.expectedOut {
				t.Errorf("TestCase[%v] Expected: \\t%v\\nGot: \\t%v\\n\\", tc.caseId, tc.expectedOut, out)
			}
		})
	}
}

func Test_validateEmail(t *testing.T) {
	testCases := []struct {
		caseId      int
		input       string
		expectedOut bool
	}{
		// Success case
		{
			caseId:      1,
			input:       "abc22@test.com",
			expectedOut: true,
		},
		// Error cases
		{
			caseId:      2,
			input:       "xyz-test.com",
			expectedOut: false,
		},
		{
			caseId:      3,
			input:       "hello",
			expectedOut: false,
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			out := validateEmail(tc.input)
			if out != tc.expectedOut {
				t.Errorf("TestCase[%v] Expected: \t%v \nGot: \t%v\n", tc.caseId, tc.expectedOut, out)
			}
		})
	}
}

func Test_validatePhone(t *testing.T) {
	testCases := []struct {
		caseId      int
		input       string
		expectedOut bool
	}{
		// Success case
		{
			caseId:      1,
			input:       "3278928923",
			expectedOut: true,
		},
		// Error cases
		{
			caseId:      2,
			input:       "abcdefghih",
			expectedOut: false,
		},
		{
			caseId:      3,
			input:       "28903091283102983",
			expectedOut: false,
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			out := validatePhone(tc.input)
			if out != tc.expectedOut {
				t.Errorf("TestCase[%v] Expected: \t%v \nGot: \t%v\n", tc.caseId, tc.expectedOut, out)
			}
		})
	}
}
