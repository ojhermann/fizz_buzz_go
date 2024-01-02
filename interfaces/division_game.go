package interfaces

type DivisionGame interface {
	ToString(dividend int) string
	Display(start int, end int) error
}
