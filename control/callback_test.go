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
			fileName: "1234_first-name",
			expect: struct {
				id   int
				name string
			}{id: 1234, name: "first-name"},
		},
		{
			fileName: "1234-first-name",
			expect: struct {
				id   int
				name string
			}{id: 0, name: ""},
		},
		{
			fileName: "1234_first_name",
			expect: struct {
				id   int
				name string
			}{id: 0, name: ""},
		},
		{
			fileName: "a1234_first_name",
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
