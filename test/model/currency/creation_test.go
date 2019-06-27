package main

import (
	"financial/assets"
	"testing"
)

func TestCurrencyCreation(t *testing.T) {
	currencyType := assets.NewCurrencyType(1, "Real")
	currency := assets.NewCurrency(23.45, &currencyType)
	if value := currency.Float64(); value != 23.45 {
		t.Errorf("expected 23.45 != received = %f", currency.Float64())
	}
	expected2String := "Currency(amount=23.45, currencyType=CurrencyType(id=1, name='Real'))"
	if currency.String() != expected2String {
		t.Errorf("expected=%s received=%s", expected2String, currency.String())
	}
}
