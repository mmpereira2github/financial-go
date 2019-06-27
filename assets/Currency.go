package assets

import (
	"fmt"
)

const PRECISION = 100

type Currency interface {
	Float64() float64
	CurrencyType() *CurrencyType
	String() string
}

type currency struct {
	amount       int64
	currencyType *CurrencyType
}

func NewCurrency(amount float64, currencyType *CurrencyType) Currency {
	return currency{int64(amount * PRECISION), currencyType}
}
func (c currency) Float64() float64            { return float64(c.amount) / PRECISION }
func (c currency) CurrencyType() *CurrencyType { return c.currencyType }
func (c currency) String() string {
	return fmt.Sprintf("Currency(amount=%d.%d, currencyType=%s)", c.amount/100, c.amount%100, (*c.currencyType).String())
}
