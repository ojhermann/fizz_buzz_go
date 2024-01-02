package implementations

import (
	"fizz_buzz/interfaces"
	"strconv"
)

func NewDivisor(divisor int, value string, fnc interfaces.Compare[int]) *Divisor {
	return &Divisor{divisor: divisor, value: value, fnc: fnc}
}

type Divisor struct {
	divisor int
	value   string
	fnc     interfaces.Compare[int]
}

func (d *Divisor) ToString(dividend int) (string, bool) {
	if d.fnc(dividend, d.divisor) {
		return d.value, true
	}

	return strconv.Itoa(dividend), false
}
