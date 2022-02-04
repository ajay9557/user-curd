package Users

import (
	"Icrud/TModels"
	"fmt"
	"reflect"
	"testing"
)

func TestFromUpdateQuery(t *testing.T) {
	testUser := TModels.User{
		Name:  "Ridhdhish",
		Email: "rid@gmail.com",
		Phone: "8320578360",
		Age:   21,
	}

	tests := []struct {
		desc           string
		user           TModels.User
		expectedFields string
		expectedArgs   []interface{}
	}{
		{
			desc:           "Case1",
			user:           testUser,
			expectedFields: "name = ?,email = ?,phone = ?,age = ?,",
			expectedArgs: []interface{}{
				"Ridhdhish",
				"rid@gmail.com",
				"8320578360",
				21,
			},
		},
		// {
		// 	desc:           "Case2",
		// 	user:           TModels.User{},
		// 	id:             -1,
		// 	expectedFields: "",
		// },
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
