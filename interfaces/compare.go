package interfaces

type Compare[C comparable] func(C, C) bool
