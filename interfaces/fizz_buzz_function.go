package interfaces

import "io"

type FizzBuzzFunction func(int, int, io.Writer) error
