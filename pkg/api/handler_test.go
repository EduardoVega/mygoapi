package api

import (
	"testing"
)

func TestEmptyString(t *testing.T) {

	tests := []struct {
		Value    string
		Expected bool
	}{
		{
			"   ",
			true,
		},
		{
			"ok",
			false,
		},
	}

	for _, test := range tests {
		r := emptyString(test.Value)
		if r != test.Expected {
			t.Errorf("Returned result was incorrect, got: %v want: %v", r, test.Expected)
		}
	}
}
