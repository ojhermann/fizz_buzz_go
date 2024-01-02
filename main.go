package main

import (
	"fizz_buzz/implementation"
	"os"
)

func main() {
	err := implementation.FizzBuzz(-1, 1, os.Stdout)
	if err != nil {
		panic(err)
	}
}
