package cmd

import "fmt"

func output(message interface{}) {
	if debug {
		fmt.Printf("%v", message)
	}
}
