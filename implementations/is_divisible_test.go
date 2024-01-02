package implementations

import "testing"

func TestIsDivisible(t *testing.T) {
	t.Run("returns false when divisor is zero", func(t *testing.T) {
		divisor := 0
		for dividend := -100; dividend <= 100; dividend++ {
			if IsDivisible(dividend, divisor) {
				t.Errorf(
					"Expect false when divisor is zero: dividend %d, divisor %d",
					dividend, divisor,
				)
			}
		}
	})
}
