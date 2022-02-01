package user

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_CheckMail(t *testing.T) {
	Tests := []struct {
		desc   string
		input  string
		output bool
		err    error
	}{
		{"Case 1", "yes@gmail.com", true, nil},
		{"Case 1", "Nogmail", false, fmt.Errorf("enter valid email")},
	}
	for _, tes := range Tests {
		t.Run(tes.desc, func(t *testing.T) {
			output, err := CheckMail(tes.input)
			if output != tes.output {
				t.Errorf("expected %t got %t", tes.output, output)
			}
			if !reflect.DeepEqual(tes.err, err) {
				t.Errorf("expected %s got %s", tes.err, err)
			}
		})
	}
}
