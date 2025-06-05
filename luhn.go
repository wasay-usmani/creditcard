package creditcard

import (
	"strconv"
)

func LuhnCheck(number string) bool {
	digits := make([]int, len(number))
	double := false

	var sum int

	for i, r := range number {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			return false // invalid character
		}

		digits[i] = n
	}

	// Process digits from right to left
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i]
		if double {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		sum += n
		double = !double
	}

	return sum%10 == 0
}
