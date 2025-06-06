package creditcard

import (
	"fmt"
	"slices"
)

// RegisterScheme returns a RegistryOption that, when applied to a registry,
// registers the provided scheme. If a scheme with the same type already
// exists in the registry, an error is returned.
func RegisterScheme(scheme *Scheme) RegistryOption {
	return func(r *registry) error {
		for _, s := range r.schemes {
			if s.Type == scheme.Type {
				return fmt.Errorf("scheme type %s already exists", scheme.Type)
			}
		}

		r.schemes = append(r.schemes, scheme)

		return nil
	}
}

// UnregisterScheme returns a RegistryOption that, when applied to a registry,
// unregisters the scheme with the provided type. If no scheme with the
// provided type exists in the registry, the returned function is a no-op.
func UnregisterScheme(t SchemeType) RegistryOption {
	return func(r *registry) error {
		for i, s := range r.schemes {
			if s.Type == t {
				r.schemes = slices.Delete(r.schemes, i, 1)
			}
		}

		return nil
	}
}

// RegisterVisa returns a RegistryOption that, when applied to a registry,
// registers the Visa scheme.
func RegisterVisa() RegistryOption {
	return RegisterScheme(NewScheme(SchemeNameVisa, SchemeTypeVisa, Code{Name: CVV, Size: CodeSize3}, VisaCardLength, visaRegexp))
}

// RegisterMastercard returns a RegistryOption that, when applied to a registry,
// registers the Mastercard scheme.
func RegisterMastercard() RegistryOption {
	return RegisterScheme(NewScheme(SchemeNameMastercard, SchemeTypeMastercard, Code{Name: CVC, Size: CodeSize3}, MastercardCardLength, mastercardRegexp))
}

// RegisterAmex returns a RegistryOption that, when applied to a registry,
// registers the American Express scheme.
func RegisterAmex() RegistryOption {
	return RegisterScheme(NewScheme(SchemeNameAmex, SchemeTypeAmex, Code{Name: CID, Size: CodeSize4}, AmexCardLength, amexRegexp))
}

// RegisterJCB returns a RegistryOption that, when applied to a registry,
// registers the JCB scheme.
func RegisterJCB() RegistryOption {
	return RegisterScheme(NewScheme(SchemeNameJcb, SchemeTypeJcb, Code{Name: CVV, Size: CodeSize3}, JcbCardLength, jcbRegexp))
}

// RegisterDiscover returns a RegistryOption that, when applied to a registry,
// registers the Discover scheme.
func RegisterDiscover() RegistryOption {
	return RegisterScheme(NewScheme(SchemeNameDiscover, SchemeTypeDiscover, Code{Name: CID, Size: CodeSize3}, DinersCardLength, discoverRegexp))
}

// RegisterUnionPay returns a RegistryOption that, when applied to a registry,
// registers the UnionPay scheme.
func RegisterUnionPay() RegistryOption {
	return RegisterScheme(NewScheme(SchemeNameUnionPay, SchemeTypeUnionPay, Code{Name: CVN, Size: CodeSize3}, UnionPayCardLength, unionPayRegexp))
}

// RegisterMaestro returns a RegistryOption that, when applied to a registry,
// registers the Maestro scheme.
func RegisterMaestro() RegistryOption {
	return RegisterScheme(NewScheme(SchemeNameMaestro, SchemeTypeMaestro, Code{Name: CVC, Size: CodeSize3}, MaestroCardLength, maestroRegexp))
}

// RegisterDiners returns a RegistryOption that, when applied to a registry,
// registers the Diners Club scheme.
func RegisterDiners() RegistryOption {
	return RegisterScheme(NewScheme(SchemeNameDiners, SchemeTypeDiners, Code{Name: CVV, Size: CodeSize3}, DinersCardLength, dinersRegexp))
}
