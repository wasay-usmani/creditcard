package creditcard

import (
	"fmt"
)

type RegistryOption func(r *registry) error
type SchemeRegistry interface {
	ValidateCard(card *Card) (*Scheme, error)
	ShowRegisteredSchemes() []*Scheme
}

type registry struct {
	schemes []*Scheme
}

// NewSchemeRegistry returns a new SchemeRegistry instance, which can be used to validate a card.
//
// The returned registry is configured to support Visa, Mastercard, and American Express cards.
// Other schemes can be registered using the opts argument.
//
// The opts argument is a variadic list of RegistryOption functions, which are called in order
// after the registry is initialized. If any of the options return an error, the error is
// returned from NewSchemeRegistry.
func NewSchemeRegistry(opts ...RegistryOption) (SchemeRegistry, error) {
	r := &registry{schemes: make([]*Scheme, 0)}
	r.schemes = append(r.schemes, NewScheme(SchemeNameVisa, SchemeTypeVisa, Code{Name: "CVV", Size: 3}, []int{13, 16, 19}, visaRegexp))
	r.schemes = append(r.schemes, NewScheme(SchemeNameMastercard, SchemeTypeMastercard, Code{Name: "CVC", Size: 3}, []int{16}, mastercardRegexp))
	r.schemes = append(r.schemes, NewScheme(SchemeNameAmex, SchemeTypeAmex, Code{Name: "CID", Size: 4}, []int{15}, amexRegexp))

	for _, opt := range opts {
		if err := opt(r); err != nil {
			panic(err)
		}
	}

	return r, nil
}

// ValidateCard takes a Card and checks if it is valid according to the registered
// schemes. If the card is valid, the function returns the Scheme and nil. If the
// card is not valid, the function returns an error.
func (r *registry) ValidateCard(card *Card) (*Scheme, error) {
	if !LuhnCheck(card.Number) {
		return nil, fmt.Errorf("card number is not valid")
	}

	for _, scheme := range r.schemes {
		if scheme.validator != nil && scheme.validator(card) {
			if card.Code != nil && len(*card.Code) != scheme.Code.Size {
				return scheme.clone(), fmt.Errorf("Code is not valid for %s", scheme.Type)
			}

			return scheme.clone(), nil
		}
	}

	return nil, fmt.Errorf("card is not valid")
}

// ShowRegisteredSchemes returns a copy of the registered schemes.
func (r *registry) ShowRegisteredSchemes() []*Scheme {
	copied := make([]*Scheme, len(r.schemes))
	copy(copied, r.schemes)
	return copied
}
