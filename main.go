package main

import (
	"fizz_buzz/implementations"
	"os"
)

func main() {
	err := implementations.FizzBuzz(-1, 1, os.Stdout)
	if err != nil {
		panic(err)
	}
}
