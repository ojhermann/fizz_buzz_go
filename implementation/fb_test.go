package implementation

import (
	"bytes"
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("implements interface", func(t *testing.T) {
		expected := "func(int, int, io.Writer) error"
		actual := reflect.TypeOf(FizzBuzz)
		if actual.String() != expected {
			t.Errorf("expected %s, but received %s", expected, actual.String())
		}
	})

	t.Run("can output correctly", func(t *testing.T) {
		start := -1
		end := 1
		output := &bytes.Buffer{}

		err := FizzBuzz(start, end, output)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		expected := "-1\nFizzBuzz\n1\n"
		actual := output.String()
		if actual != expected {
			t.Errorf("expected %s, but received: %v", expected, actual)
		}
	})
}
