package implementations

import (
	"fizz_buzz/interfaces"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const emptyString string = ""

func FizzBuzz(start int, end int, writer io.Writer) error {
	var isDivisible bool
	var s string
	var sb strings.Builder

	divisors := []interfaces.Divisor{
		NewDivisor(3, "Fizz", IsDivisible),
		NewDivisor(5, "Buzz", IsDivisible),
	}

	for dividend := start; dividend <= end; dividend++ {
		for _, divisor := range divisors {
			s, isDivisible = divisor.ToString(dividend)
			if isDivisible {
				sb.WriteString(s)
			}
		}

		s = sb.String()
		sb.Reset()
		if s == emptyString {
			s = strconv.Itoa(dividend)
		}

		_, err := fmt.Fprintf(writer, "%s\n", s)
		if err != nil {
			return err
		}
	}

	return nil
}
