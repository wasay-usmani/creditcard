package creditcard

import (
	"testing"
)

func ptrOf[T any](v T) *T { return &v }

func TestNewSchemeRegistry_RegisterAndUnregister(t *testing.T) {
	reg, err := NewSchemeRegistry(RegisterDiscover())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	schemes := reg.ShowRegisteredSchemes()
	found := false

	for _, s := range schemes {
		if s.Type == SchemeTypeDiscover {
			found = true

			break
		}
	}

	if !found {
		t.Errorf("Discover scheme was not registered")
	}

	// Unregister the scheme
	reg, err = NewSchemeRegistry(UnregisterScheme(SchemeTypeDiscover))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	schemes = reg.ShowRegisteredSchemes()
	for _, s := range schemes {
		if s.Type == SchemeTypeDiscover {
			t.Errorf("Discover scheme was not unregistered")
		}
	}
}

func TestValidateCard(t *testing.T) {
	type testCase struct {
		name      string
		card      Card
		wantValid bool
		wantType  SchemeType
	}

	reg, err := NewSchemeRegistry(
		RegisterDiscover(),
		RegisterJCB(),
		RegisterUnionPay(),
		RegisterMaestro(),
		RegisterDiners(),
	)

	if err != nil {
		t.Fatalf("failed to create registry: %v", err)
	}

	// Add new test cases for the additional schemes
	testCases := []testCase{
		{
			name:      "Valid Visa",
			card:      Card{Number: "4242424242424242", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeVisa,
		},
		{
			name:      "Valid Visa Debit",
			card:      Card{Number: "4000056655665556", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeVisa,
		},
		{
			name:      "Valid Mastercard",
			card:      Card{Number: "5555555555554444", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeMastercard,
		},
		{
			name:      "Valid Mastercard 2-series",
			card:      Card{Number: "2223003122003222", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeMastercard,
		},
		{
			name:      "Valid Mastercard Debit",
			card:      Card{Number: "5200828282828210", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeMastercard,
		},
		{
			name:      "Valid Mastercard Prepaid",
			card:      Card{Number: "5105105105105100", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeMastercard,
		},
		{
			name:      "Valid American Express",
			card:      Card{Number: "378282246310005", Code: ptrOf(1234)},
			wantValid: true,
			wantType:  SchemeTypeAmex,
		},
		{
			name:      "Valid American Express",
			card:      Card{Number: "371449635398431", Code: ptrOf(1234)},
			wantValid: true,
			wantType:  SchemeTypeAmex,
		},
		{
			name:      "Valid Discover",
			card:      Card{Number: "6011111111111117", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeDiscover,
		},
		{
			name:      "Valid Discover",
			card:      Card{Number: "6011000990139424", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeDiscover,
		},
		{
			name:      "Valid Discover Debit",
			card:      Card{Number: "6011981111111113", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeDiscover,
		},
		{
			name:      "Valid Diners Club",
			card:      Card{Number: "3056930009020004", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeDiners,
		},
		{
			name:      "Valid Diners Club 14-digit",
			card:      Card{Number: "36227206271667", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeDiners,
		},
		{
			name:      "Valid JCB",
			card:      Card{Number: "3566002020360505", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeJcb,
		},
		{
			name:      "Valid UnionPay",
			card:      Card{Number: "6200000000000005", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeUnionPay,
		},
		{
			name:      "Valid UnionPay Debit",
			card:      Card{Number: "6200000000000047", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeUnionPay,
		},
		{
			name:      "Valid UnionPay 19-digit",
			card:      Card{Number: "6205500000000000004", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeUnionPay,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			scheme, err := reg.ValidateCard(&tc.card)
			if tc.wantValid {
				if err != nil {
					t.Errorf("expected valid, got error: %v", err)
				}

				if scheme != nil && scheme.Type != tc.wantType {
					t.Errorf("expected scheme type %v, got %v", tc.wantType, scheme.Type)
				}
			} else if err == nil {
				t.Errorf("expected error, got valid scheme: %+v", scheme)
			}
		})
	}
}
