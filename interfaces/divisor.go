package interfaces

type Divisor interface {
	ToString(dividend int) (string, bool)
}
