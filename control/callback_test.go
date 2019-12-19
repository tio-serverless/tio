package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitUidAndSrvName(t *testing.T) {
	cases := []struct {
		fileName string
		expect   struct {
			id    int
			name  string
			stype string
		}
	}{
		{
			fileName: "1234-firstname-grpc",
			expect: struct {
				id    int
				name  string
				stype string
			}{id: 0, name: "", stype: ""},
		},
		{
			fileName: "1234-firstname-grpc.zip",
			expect: struct {
				id    int
				name  string
				stype string
			}{id: 1234, name: "firstname", stype: "grpc"},
		},
		{
			fileName: "1234_firstname_grpc",
			expect: struct {
				id    int
				name  string
				stype string
			}{id: 0, name: "", stype: ""},
		},
		{
			fileName: "1234_firstname_grpc.zip",
			expect: struct {
				id    int
				name  string
				stype string
			}{id: 0, name: "", stype: ""},
		},
		{
			fileName: "1234_firstname_grpc",
			expect: struct {
				id    int
				name  string
				stype string
			}{id: 0, name: "", stype: ""},
		},
		{
			fileName: "a1234_firstname_grpc",
			expect: struct {
				id    int
				name  string
				stype string
			}{id: 0, name: "", stype: ""},
		},
	}

	for _, c := range cases {
		uid, name, stype := splitUidAndSrvName(c.fileName)
		assert.EqualValues(t, c.expect.id, uid)
		assert.EqualValues(t, c.expect.name, name)
		assert.EqualValues(t, c.expect.stype, stype)
	}
}

func TestTrimTimestamp(t *testing.T) {
	cases := []struct {
		fileName string
		expect   string
	}{
		{
			fileName: "2-xt-grpc-123456.zip",
			expect:   "2-xt-grpc",
		},
		{
			fileName: "2-xt-grpc-123456",
			expect:   "2-xt-grpc",
		},
	}

	for _, c := range cases {
		actual := trimTimestamp(c.fileName)
		assert.EqualValuesf(t, c.expect, actual, actual)
	}
}
