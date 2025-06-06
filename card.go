package creditcard

import (
	"strings"
)

type CardOptions func(*Card)
type Card struct {
	Number string
	Code   *int
}

func NewCard(number string, opts ...CardOptions) *Card {
	c := &Card{Number: strings.ReplaceAll(number, " ", "")}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithCode(code int) CardOptions {
	return func(c *Card) {
		c.Code = &code
	}
}

func (c *Card) String() string {
	return c.Number
}

func (c *Card) Bin() string {
	return c.Number[:6]
}

func (c *Card) Last4() string {
	return c.Number[len(c.Number)-4:]
}

func (c *Card) MaskedCard() string {
	maskedLen := len(c.Number) - 6 - 4
	if maskedLen < 0 {
		return c.Number
	}

	return c.Bin() + strings.Repeat("*", maskedLen) + c.Last4()
}
