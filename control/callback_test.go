package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitUidAndSrvName(t *testing.T) {
	cases := []struct {
		fileName string
		expect   struct {
			id   int
			name string
		}
	}{
		{
			fileName: "1234-firstname",
			expect: struct {
				id   int
				name string
			}{id: 1234, name: "firstname"},
		},
		{
			fileName: "1234_firstname",
			expect: struct {
				id   int
				name string
			}{id: 0, name: ""},
		},
		{
			fileName: "1234_firstname",
			expect: struct {
				id   int
				name string
			}{id: 0, name: ""},
		},
		{
			fileName: "a1234_firstname",
			expect: struct {
				id   int
				name string
			}{id: 0, name: ""},
		},
	}

	for _, c := range cases {
		uid, name := splitUidAndSrvName(c.fileName)
		assert.EqualValues(t, c.expect.id, uid)
		assert.EqualValues(t, c.expect.name, name)
	}
}
