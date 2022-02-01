package driver

import (
	"errors"
	"testing"
)

func TestConnectToMySql(t *testing.T) {
	tests := []struct {
		desc       string
		driverName string
		expected   error
	}{
		{
			desc:       "Case1",
			driverName: "postgres",
			expected:   errors.New("Could not able to connet to the given database driver"),
		},
		{
			desc:       "Case2",
			driverName: "mysql",
			expected:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db, err := ConnectToMySql(test.driverName)

			if errors.Is(err, test.expected) && db == nil {
				t.Errorf("Expected: %v, Got: %v", test.expected, err)
			}
		})
	}
}
