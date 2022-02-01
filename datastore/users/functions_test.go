package users

import (
	"reflect"
	"strconv"
	"testing"
	"user-curd/entities"
)

func Test_formQuery(t *testing.T) {
	testCases := []struct {
		caseId         int
		input          entities.User
		expectedQuery  string
		expectedValues []interface{}
	}{
		// All fields
		{
			caseId: 1,
			input: entities.User{
				Id:    1,
				Name:  "jessi",
				Email: "jess88@example.com",
				Phone: "7892311212",
				Age:   22,
			},
			expectedQuery:  " name = ?, email = ?, phone = ?, age = ?",
			expectedValues: []interface{}{"jessi", "jess88@example.com", "7892311212", 22, 1},
		},
		// invalid id
		{
			caseId: 1,
			input: entities.User{
				Id:    -1,
				Name:  "jessi",
				Email: "jess88@example.com",
				Phone: "7892311212",
				Age:   22,
			},
			expectedQuery:  "",
			expectedValues: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			outQ, outV := formQuery(tc.input)
			if outQ != tc.expectedQuery {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedQuery, outQ)
			}
			if !reflect.DeepEqual(outV, tc.expectedValues) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedValues, outV)
			}
		})
	}
}
