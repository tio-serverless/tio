package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFaild(t *testing.T) {
	tests := []struct {
		address string
		user    string
		name    string
		expect  error
	}{
		{
			address: "127.0.0.1:80",
			user:    "",
			name:    "xxx",
			expect:  errors.New("User / Name Empty! "),
		},
		{
			address: "127.0.0.1:80",
			user:    "xxx",
			name:    "",
			expect:  errors.New("User / Name Empty! "),
		},
	}
	for _, test := range tests {
		err := faild(test.address, test.user, test.name)
		assert.EqualValues(t, err.Error(), test.expect.Error())
	}
}
