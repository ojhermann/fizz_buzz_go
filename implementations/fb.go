package implementations

import (
	"io"
)

func FizzBuzz(start int, end int, writer io.Writer) error {
	divisionGame := NewDivisionGame(writer)
	divisionGame.Add(NewDivisor(3, "Fizz", IsDivisible))
	divisionGame.Add(NewDivisor(5, "Buzz", IsDivisible))

	err := divisionGame.Display(start, end)
	if err != nil {
		return err
	}
	return nil
}
