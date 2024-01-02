package implementations

import (
	"fizz_buzz/interfaces"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const EmptyString string = ""

func NewDivisionGame(writer io.Writer) *DivisionGame {
	return &DivisionGame{writer: writer}
}

type DivisionGame struct {
	divisors []interfaces.Divisor
	sb       strings.Builder
	writer   io.Writer
}

func (d *DivisionGame) Add(divisor interfaces.Divisor) {
	d.divisors = append(d.divisors, divisor)
}

func (d *DivisionGame) ToString(dividend int) string {
	for _, divisor := range d.divisors {
		s, isDivisible := divisor.ToString(dividend)
		if isDivisible {
			d.sb.WriteString(s)
		}
	}

	s := d.sb.String()
	d.sb.Reset()
	if s == EmptyString {
		return strconv.Itoa(dividend)
	}
	return s
}

func (d *DivisionGame) Display(start int, end int) error {
	for dividend := start; dividend <= end; dividend++ {
		_, err := fmt.Fprintf(d.writer, "%s\n", d.ToString(dividend))
		if err != nil {
			return err
		}
	}
	return nil
}
