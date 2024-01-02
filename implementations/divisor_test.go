package implementations

import (
	"strconv"
	"testing"
)

func TestDivisor_ToString(t *testing.T) {
	t.Run("returns the special string if fnc returns true", func(t *testing.T) {
		specialString := "specialString"
		d := Divisor{
			divisor: 0,
			value:   specialString,
			fnc:     func(i int, i2 int) bool { return true },
		}
		res, isDivisible := d.ToString(1)
		if res != specialString || !isDivisible {
			t.Errorf(
				"fnc always returns true, so expected res == %s and isDivisible == true, but received %s, %t",
				specialString, res, isDivisible,
			)
		}
	})

	t.Run("returns string representation of dividend if fnc returns false", func(t *testing.T) {
		dividend := 1
		intString := strconv.Itoa(dividend)
		d := Divisor{
			divisor: 0,
			value:   "someString",
			fnc:     func(i int, i2 int) bool { return false },
		}
		res, isDivisible := d.ToString(1)
		if res != intString || isDivisible {
			t.Errorf(
				"fnc always returns false, so expected res == %s and isDivisible == false, but received %s, %t",
				intString, res, isDivisible,
			)
		}
	})
}
