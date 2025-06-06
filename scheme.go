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

type CodeName string

// Code names
const (
	CVV CodeName = "CVV"
	CVC CodeName = "CVC"
	CVN CodeName = "CVN"
	CID CodeName = "CID"
)

type CodeSize int

func (c CodeSize) IsEqual(other int) bool {
	return int(c) == int(other)
}

// Code sizes
const (
	CodeSize3 CodeSize = 3
	CodeSize4 CodeSize = 4
)

// Regular expressions
var (
	visaRegexp       = regexp.MustCompile(`^(?:4\d{12}(?:\d{3})?)$`)
	mastercardRegexp = regexp.MustCompile(
		`^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12}))$`)
	amexRegexp     = regexp.MustCompile(`^3[47]\d{13}$`)
	dinersRegexp   = regexp.MustCompile(`^3(?:0[0-5]|[68]\d)\d{11}$`)
	discoverRegexp = regexp.MustCompile(
		`^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$`)
	jcbRegexp      = regexp.MustCompile(`^(?:2131|1800|35\d{3})\d{11}$`)
	unionPayRegexp = regexp.MustCompile(`^(62\d{14,17})$`)
	maestroRegexp  = regexp.MustCompile(`^(5018|5020|5038|6304|6759|6761|6763)\d{8,15}$`)
)

type CardLength []int

var (
	VisaCardLength       = CardLength{16, 18, 19}
	MastercardCardLength = CardLength{16}
	AmexCardLength       = CardLength{15}
	DinersCardLength     = CardLength{14, 16, 19}
	DiscoverCardLength   = CardLength{16, 19}
	JcbCardLength        = CardLength{16, 17, 18, 19}
	UnionPayCardLength   = CardLength{14, 15, 16, 17, 18, 19}
	MaestroCardLength    = CardLength{12, 13, 14, 15, 16, 17, 18, 19}
)

type SchemeValidator func(card *Card) bool
type Scheme struct {
	Lengths   CardLength      `json:"lengths"`
	Code      Code            `json:"code"`
	Type      SchemeType      `json:"scheme_type"`
	Name      SchemeName      `json:"scheme_name"`
	regex     *regexp.Regexp  `json:"-"`
	validator SchemeValidator `json:"-"`
}

type Code struct {
	Name CodeName `json:"name"`
	Size CodeSize `json:"size"`
}

// NewScheme returns a new Scheme, which can be used to validate cards.
//
// The lengths argument is a list of valid lengths for the card number.
// The code argument is the information about the code (CVC, CID, etc.)
// to be validated. The regex argument is a regular expression that the
// card number must match.
func NewScheme(name SchemeName, schemeType SchemeType, code Code,
	len CardLength, regex *regexp.Regexp) *Scheme {
	return &Scheme{
		Lengths:   len,
		Code:      code,
		Type:      schemeType,
		Name:      name,
		regex:     regex,
		validator: RegexpValidator(regex),
	}
}

func RegexpValidator(rex *regexp.Regexp) SchemeValidator {
	return func(card *Card) bool {
		return rex.MatchString(card.Number)
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
