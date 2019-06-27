package assets

import (
	"fmt"
)

const precision = 100

// Currency represents an currency amount of some type
type Currency interface {
	Float64() float64
	CurrencyType() *CurrencyType
	String() string
}

type currency struct {
	amount       int64
	currencyType *CurrencyType
}

// NewCurrency creates a new Currency instance of given amount and type
func NewCurrency(amount float64, currencyType *CurrencyType) Currency {
	return &currency{int64(amount * precision), currencyType}
}
func (c *currency) Float64() float64            { return float64(c.amount) / precision }
func (c *currency) CurrencyType() *CurrencyType { return c.currencyType }
func (c *currency) String() string {
	return fmt.Sprintf("Currency(amount=%d.%d, currencyType=%s)", c.amount/100, c.amount%100, (*c.currencyType).String())
}
