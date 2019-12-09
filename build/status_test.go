package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFaild(t *testing.T) {
	tests := []struct {
		j      job
		expect error
	}{
		{
			j: job{
				User:   "",
				Name:   "",
				Image:  "",
				API:    "",
				Rate:   0,
				Status: 0,
			},

			expect: errors.New("User / Name / Image Empty! "),
		},
		{
			j: job{
				User:   "xxxx",
				Name:   "",
				Image:  "",
				API:    "",
				Rate:   0,
				Status: 0,
			},
			expect: errors.New("User / Name / Image Empty! "),
		},
		{
			j: job{
				User:   "xxxx",
				Name:   "xxxx",
				Image:  "",
				API:    "",
				Rate:   0,
				Status: 0,
			},
			expect: errors.New("User / Name / Image Empty! "),
		},
	}
	for _, test := range tests {
		err := building("xx", &test.j)
		assert.EqualValues(t, err.Error(), test.expect.Error())
	}
}
