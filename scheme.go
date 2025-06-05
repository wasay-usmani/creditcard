package creditcard

import "regexp"

type SchemeType string
type SchemeName string

// Scheme types
const (
	SchemeTypeVisa       SchemeType = "visa"
	SchemeTypeMastercard SchemeType = "mastercard"
	SchemeTypeAmex       SchemeType = "american-express"
	SchemeTypeDiners     SchemeType = "diners-club"
	SchemeTypeDiscover   SchemeType = "discover"
	SchemeTypeJcb        SchemeType = "jcb"
	SchemeTypeUnionPay   SchemeType = "unionpay"
	SchemeTypeMaestro    SchemeType = "maestro"
)

// Scheme names
const (
	SchemeNameVisa       SchemeName = "Visa"
	SchemeNameMastercard SchemeName = "Mastercard"
	SchemeNameAmex       SchemeName = "American Express"
	SchemeNameDiners     SchemeName = "Diners Club"
	SchemeNameDiscover   SchemeName = "Discover"
	SchemeNameJcb        SchemeName = "JCB"
	SchemeNameUnionPay   SchemeName = "UnionPay"
	SchemeNameMaestro    SchemeName = "Maestro"
)

var (
	visaRegexp       = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?)$`)
	mastercardRegexp = regexp.MustCompile(`^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12}))$`)
	amexRegexp       = regexp.MustCompile(`^3[47][0-9]{13}$`)
	dinersRegexp     = regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`)
	discoverRegexp   = regexp.MustCompile(`^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$`)
	jcbRegexp        = regexp.MustCompile(`^(?:2131|1800|35\d{3})\d{11}$`)
	unionPayRegexp   = regexp.MustCompile(`^(62[0-9]{14,17})$`)
	maestroRegexp    = regexp.MustCompile(`^(5018|5020|5038|6304|6759|6761|6763)[0-9]{8,15}$`)
)

type SchemeValidator func(card *Card) bool
type Scheme struct {
	Lengths   []int           `json:"lengths"`
	Code      Code            `json:"code"`
	Type      SchemeType      `json:"scheme_type"`
	Name      SchemeName      `json:"scheme_name"`
	regex     *regexp.Regexp  `json:"-"`
	validator SchemeValidator `json:"-"`
}

type Code struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

// NewScheme returns a new Scheme, which can be used to validate cards.
//
// The lengths argument is a list of valid lengths for the card number.
// The code argument is the information about the code (CVC, CID, etc.)
// to be validated. The regex argument is a regular expression that the
// card number must match.
func NewScheme(name SchemeName, schemeType SchemeType, code Code,
	lengths []int, regex *regexp.Regexp) *Scheme {
	return &Scheme{
		Lengths:   lengths,
		Code:      code,
		Type:      schemeType,
		Name:      name,
		regex:     regex,
		validator: RegexpValidator(regex),
	}
}

func RegexpValidator(regexp *regexp.Regexp) SchemeValidator {
	return func(card *Card) bool {
		if regexp.MatchString(card.Number) {
			return true
		}

		return false
	}
}

// Is returns true if the scheme type matches the target SchemeType.
func (s *Scheme) Is(target SchemeType) bool {
	return s.Type == target
}

func (s *Scheme) clone() *Scheme {
	return &Scheme{
		Lengths: s.Lengths,
		Code:    s.Code,
		Type:    s.Type,
		Name:    s.Name,
	}
}
