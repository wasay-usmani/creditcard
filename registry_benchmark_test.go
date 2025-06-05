package creditcard

import (
	"testing"
)

func BenchmarkValidateCard_10000(b *testing.B) {
	reg, err := NewSchemeRegistry()
	if err != nil {
		b.Fatalf("failed to create registry: %v", err)
	}

	cards := make([]*Card, 10000)
	for i := 0; i < 10000; i++ {
		switch i % 8 {
		case 0:
			cards[i] = &Card{Number: "4111111111111111", Code: strPtr("123")} // valid Visa
		case 1:
			cards[i] = &Card{Number: "5555555555554444", Code: strPtr("123")} // valid Mastercard
		case 2:
			cards[i] = &Card{Number: "378282246310005", Code: strPtr("1234")} // valid Amex
		case 3:
			cards[i] = &Card{Number: "4111111111111112", Code: strPtr("123")} // invalid Visa (Luhn fail)
		case 4:
			cards[i] = &Card{Number: "5555555555554440", Code: strPtr("123")} // invalid Mastercard (Luhn fail)
		case 5:
			cards[i] = &Card{Number: "378282246310006", Code: strPtr("1234")} // invalid Amex (Luhn fail)
		case 6:
			cards[i] = &Card{Number: "6011000990139424", Code: strPtr("123")} // Discover (not in registry)
		case 7:
			cards[i] = &Card{Number: "3530111333300000", Code: strPtr("123")} // JCB (not in registry)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		card := cards[i%len(cards)]
		_, _ = reg.ValidateCard(card)
	}
}

func BenchmarkValidateCard_100000(b *testing.B) {
	reg, err := NewSchemeRegistry()
	if err != nil {
		b.Fatalf("failed to create registry: %v", err)
	}

	cards := make([]*Card, 100000)
	for i := 0; i < 100000; i++ {
		switch i % 8 {
		case 0:
			cards[i] = &Card{Number: "4111111111111111", Code: strPtr("123")} // valid Visa
		case 1:
			cards[i] = &Card{Number: "5555555555554444", Code: strPtr("123")} // valid Mastercard
		case 2:
			cards[i] = &Card{Number: "378282246310005", Code: strPtr("1234")} // valid Amex
		case 3:
			cards[i] = &Card{Number: "4111111111111112", Code: strPtr("123")} // invalid Visa (Luhn fail)
		case 4:
			cards[i] = &Card{Number: "5555555555554440", Code: strPtr("123")} // invalid Mastercard (Luhn fail)
		case 5:
			cards[i] = &Card{Number: "378282246310006", Code: strPtr("1234")} // invalid Amex (Luhn fail)
		case 6:
			cards[i] = &Card{Number: "6011000990139424", Code: strPtr("123")} // Discover (not in registry)
		case 7:
			cards[i] = &Card{Number: "3530111333300000", Code: strPtr("123")} // JCB (not in registry)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		card := cards[i%len(cards)]
		_, _ = reg.ValidateCard(card)
	}
}
