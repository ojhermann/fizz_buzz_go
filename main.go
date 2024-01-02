package main

import (
	"fizz_buzz/implementations"
	"os"
)

func main() {
	err := implementations.FizzBuzz(-100, 100, os.Stdout)
	if err != nil {
		panic(err)
	}
}
