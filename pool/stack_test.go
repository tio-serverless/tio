package pool

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	s := NewStack()

	assert.EqualValues(t, s.Len(), 0)

	s.Push(func(s string) error {
		fmt.Println("1" + s)
		return nil
	})

	s.Push(func(s string) error {
		fmt.Println("2" + s)
		return nil
	})

	s.Push(func(s string) error {
		fmt.Println("3" + s)
		return nil
	})

	s.Push(func(s string) error {
		fmt.Println("4" + s)
		return nil
	})

	assert.EqualValues(t, s.Len(), 4)

	f := s.Pull()
	f("1")
	f = s.Pull()
	f("2")
	f = s.Pull()
	f("3")
	f = s.Pull()
	f("4")

	assert.EqualValues(t, s.Len(), 0)
}
