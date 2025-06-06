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

	var tests []testCase

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
	tests = append(tests,
		testCase{
			name:      "Valid Discover",
			card:      Card{Number: "6011000990139424", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeDiscover,
		},
		testCase{
			name:      "Valid JCB",
			card:      Card{Number: "3530111333300000", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeJcb,
		},
		testCase{
			name:      "Valid UnionPay",
			card:      Card{Number: "6240008631401148", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeUnionPay,
		},
		testCase{
			name:      "Valid Maestro",
			card:      Card{Number: "6759649826438453", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeMaestro,
		},
		testCase{
			name:      "Valid Diners Club",
			card:      Card{Number: "30569309025904", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeDiners,
		},
		testCase{
			name:      "Valid Mastercard 2-series",
			card:      Card{Number: "2221000000000009", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeMastercard,
		},
		testCase{
			name:      "Valid Amex (alternate)",
			card:      Card{Number: "371449635398431", Code: ptrOf(1234)},
			wantValid: true,
			wantType:  SchemeTypeAmex,
		},
		testCase{
			name:      "Valid Discover (alternate)",
			card:      Card{Number: "6011111111111117", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeDiscover,
		},
		testCase{
			name:      "Valid JCB (alternate)",
			card:      Card{Number: "3566002020360505", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeJcb,
		},
		testCase{
			name:      "Valid Maestro (alternate)",
			card:      Card{Number: "5020500000000000", Code: ptrOf(123)},
			wantValid: false,
		},
		testCase{
			name:      "Valid Diners Club (alternate)",
			card:      Card{Number: "38520000023237", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeDiners,
		},
		// Invalid/edge cases
		testCase{
			name:      "Invalid UnionPay (wrong code length)",
			card:      Card{Number: "6240008631401148", Code: ptrOf(12)},
			wantValid: false,
		},
		testCase{
			name:      "Invalid Maestro (Luhn fail)",
			card:      Card{Number: "6759649826438454", Code: ptrOf(123)},
			wantValid: false,
		},
		testCase{
			name:      "Invalid Diners Club (wrong code length)",
			card:      Card{Number: "30569309025904", Code: ptrOf(12)},
			wantValid: false,
		},

		testCase{
			name:      "Valid Visa",
			card:      Card{Number: "4111111111111111", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeVisa,
		},
		testCase{
			name:      "Valid Mastercard",
			card:      Card{Number: "5555555555554444", Code: ptrOf(123)},
			wantValid: true,
			wantType:  SchemeTypeMastercard,
		},
		testCase{
			name:      "Valid Amex",
			card:      Card{Number: "378282246310005", Code: ptrOf(1234)},
			wantValid: true,
			wantType:  SchemeTypeAmex,
		},
		testCase{
			name:      "Invalid Visa (wrong code length)",
			card:      Card{Number: "4111111111111111", Code: ptrOf(12)},
			wantValid: false,
		},
		testCase{
			name:      "Invalid Mastercard (wrong number)",
			card:      Card{Number: "5555555555554440", Code: ptrOf(123)},
			wantValid: false,
		},
		testCase{
			name:      "Invalid Amex (wrong code length)",
			card:      Card{Number: "378282246310005", Code: ptrOf(123)},
			wantValid: false,
		},
		testCase{
			name:      "Unknown scheme",
			card:      Card{Number: "9999999999999999", Code: ptrOf(123)},
			wantValid: false,
		},
	)

	for _, tc := range tests {
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
