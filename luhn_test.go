package creditcard

import "testing"

func TestLuhnCheck(t *testing.T) {
	tests := []struct {
		name     string
		number   string
		expected bool
	}{
		// Valid card numbers
		{"Visa valid", "4539578763621486", true},
		{"Discover valid", "6011000990139424", true},
		{"American Express valid", "378282246310005", true},
		{"Mastercard valid", "5555555555554444", true},
		{"Diners Club valid", "30569309025904", true},
		{"JCB valid", "3530111333300000", true},

		// Invalid card numbers
		{"Visa invalid", "4539578763621487", false},
		{"Discover invalid", "6011000990139425", false},
		{"American Express invalid", "378282246310006", false},
		{"Mastercard invalid", "5555555555554440", false},
		{"Diners Club invalid", "30569309025900", false},
		{"JCB invalid", "3530111333300001", false},

		// Edge cases
		{"Non-digit string", "abcdefg", false},
		{"All zeros", "0000000000000000", true}, // valid Luhn
	}

	for _, tt := range tests {
		result := LuhnCheck(tt.number)
		if result != tt.expected {
			t.Errorf("Test case failed: %s | LuhnCheck(%q) = %v; want %v", tt.name, tt.number, result, tt.expected)
		}
	}
}
