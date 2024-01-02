package implementations

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"
)

const DivisorAlwaysTrueValue string = "AlwaysTrue"

func TestDivisionGame_ToString(t *testing.T) {
	t.Run("returns int string when no divisors", func(t *testing.T) {
		var expected string
		var actual string
		divisionGame := NewDivisionGame(os.Stdout)
		for dividend := -10; dividend <= 10; dividend++ {
			expected = strconv.Itoa(dividend)
			actual = divisionGame.ToString(dividend)
			if actual != expected {
				t.Errorf(
					"should be an intString of dividend b/c there are no divisors: expected %s, but received %s",
					expected, actual,
				)
			}
		}
	})

	t.Run("returns int string when no divisor divides the dividend", func(t *testing.T) {
		divisionGame := NewDivisionGame(os.Stdout)
		divisionGame.Add(&DivisorAlwaysFalse{})
		divisionGame.Add(&DivisorAlwaysFalse{})
		divisionGame.Add(&DivisorAlwaysFalse{})

		var expected string
		var actual string
		for dividend := -10; dividend <= 10; dividend++ {
			expected = strconv.Itoa(dividend)
			actual = divisionGame.ToString(dividend)
			if actual != expected {
				t.Errorf(
					"should be an intString of dividend b/c each fnc returns false: expected %s, but received %s",
					expected, actual,
				)
			}
		}
	})

	t.Run("returns special value when at least one divisor divides the dividend", func(t *testing.T) {
		divisionGame := NewDivisionGame(os.Stdout)
		divisionGame.Add(&DivisorAlwaysFalse{})
		divisionGame.Add(&DivisorAlwaysTrue{})
		divisionGame.Add(&DivisorAlwaysFalse{})

		var expected string
		var actual string
		for dividend := -10; dividend <= 10; dividend++ {
			expected = DivisorAlwaysTrueValue
			actual = divisionGame.ToString(dividend)
			if actual != expected {
				t.Errorf(
					"should be expected b/c one and only one fnc returns true: expected %s, but received %s",
					expected, actual,
				)
			}
		}
	})
}

func TestDivisionGame_Display(t *testing.T) {
	t.Run("will display", func(t *testing.T) {
		var output bytes.Buffer

		divisionGame := NewDivisionGame(&output)
		divisionGame.Add(&DivisorAlwaysTrue{})

		err := divisionGame.Display(1, 1)
		if err != nil {
			t.Error(err)
		}

		expected := fmt.Sprintf("%s\n", DivisorAlwaysTrueValue)
		if output.String() != expected {
			t.Errorf("Expected %s, but received %s", expected, output.String())
		}
	})

	t.Run("will return error", func(t *testing.T) {
		divisionGame := NewDivisionGame(&BogusWriter{})

		err := divisionGame.Display(1, 1)
		if err == nil {
			t.Error("err should be BogusWriter")
		}

		if err.Error() != "BogusWriter" {
			t.Errorf("Expected BogusWriter, but received: %s", err.Error())
		}
	})
}

type DivisorAlwaysFalse struct{}

func (d *DivisorAlwaysFalse) ToString(dividend int) (string, bool) {
	return strconv.Itoa(dividend), false
}

type DivisorAlwaysTrue struct{}

func (d *DivisorAlwaysTrue) ToString(dividend int) (string, bool) {
	return DivisorAlwaysTrueValue, true
}

type BogusWriter struct{}

func (w *BogusWriter) Write(p []byte) (n int, err error) {
	return -1, errors.New("BogusWriter")
}
