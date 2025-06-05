# creditcard

A Go library for credit card scheme detection and validation. Supports Visa, Mastercard, Amex, and easily extensible to other schemes (Discover, JCB, UnionPay, Maestro, Diners Club, etc).

## Installation

```
go get github.com/wasay-usmani/creditcard
```

## Getting Started

### Creating a New Scheme Registry

By default, the registry supports Visa, Mastercard, and American Express. You can add more schemes using options:

```go
import "github.com/wasay-usmani/creditcard"

// Create a registry with default schemes
reg, err := creditcard.NewSchemeRegistry()
if err != nil {
    // handle error
}

// Create a registry with additional schemes (Discover, JCB, etc)
reg, err := creditcard.NewSchemeRegistry(
    creditcard.RegisterDiscover(),
    creditcard.RegisterJCB(),
    creditcard.RegisterUnionPay(),
    creditcard.RegisterMaestro(),
    creditcard.RegisterDiners(),
)
if err != nil {
    // handle error
}
```

### Registering Your Own Scheme

You can register a custom scheme using a regular expression or your own validator:

```go
myScheme := creditcard.NewScheme(
    creditcard.SchemeName("MyCard"),
    creditcard.SchemeType("mycard"),
    creditcard.Code{Name: "CVC", Size: 3},
    []int{16},
    myRegexp, // *regexp.Regexp for your card pattern
)
reg, err := creditcard.NewSchemeRegistry(creditcard.RegisterScheme(myScheme))
```

#### Custom Validator Example

If you want to use your own validation logic:

```go
myValidator := func(card *creditcard.Card) bool {
    // custom logic
    return strings.HasPrefix(card.Number, "777") && len(card.Number) == 16
}
myScheme := &creditcard.Scheme{
    Lengths:   []int{16},
    Code:      creditcard.Code{Name: "CVC", Size: 3},
    Type:      creditcard.SchemeType("mycard"),
    Name:      creditcard.SchemeName("MyCard"),
    validator: myValidator,
}
reg, err := creditcard.NewSchemeRegistry(creditcard.RegisterScheme(myScheme))
```

### Unregistering Schemes

You can remove a scheme from the registry:

```go
reg, err := creditcard.NewSchemeRegistry(
    creditcard.UnregisterScheme(creditcard.SchemeTypeDiscover),
)
```

### Validating Cards

```go
card := &creditcard.Card{Number: "4111111111111111", Code: creditcard.StrPtr("123")}
matchedScheme, err := reg.ValidateCard(card)
if err != nil {
    // invalid card
} else {
    fmt.Println("Card is valid for scheme:", matchedScheme.Name)
}
```

## Utilities

- `ShowRegisteredSchemes()` returns a copy of all registered schemes.
- `StrPtr(s string) *string` is a helper to create a pointer to a string.

## License

MIT
