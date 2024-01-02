package implementations

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

const emptyString string = ""

func FizzBuzz(start int, end int, writer io.Writer) error {
	var s string
	var sb strings.Builder

	for dividend := start; dividend <= end; dividend++ {
		if dividend%3 == 0 {
			sb.WriteString("Fizz")
		}

		if dividend%5 == 0 {
			sb.WriteString("Buzz")
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
