package users

import (
	"fmt"
	"layer/user/models"
	"reflect"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	testUser := models.User{
		Name:  "Ridhdhish",
		Email: "rid@gmail.com",
		Phone: "8320578360",
		Age:   21,
	}

	tests := []struct {
		desc           string
		user           models.User
		id             int
		expectedFields string
		expectedArgs   []interface{}
	}{
		{
			desc:           "Case1",
			user:           testUser,
			id:             1,
			expectedFields: " name = ?, email = ?, phone = ?, age = ?",
			expectedArgs: []interface{}{
				"Ridhdhish",
				"rid@gmail.com",
				"8320578360",
				21,
				1,
			},
		},
		{
			desc:           "Case2",
			user:           models.User{},
			id:             -1,
			expectedFields: "",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			fields, args := formUpdateQuery(test.user)

			fmt.Println(reflect.DeepEqual(args, test.expectedArgs))

			if fields != test.expectedFields {
				t.Errorf("Expected: %v, Got: %v", test.expectedFields, fields)
			}

			if !reflect.DeepEqual(args, test.expectedArgs) {
				t.Errorf("Expected: %v, Got: %v", test.expectedArgs, args)
			}
		})
	}
}
