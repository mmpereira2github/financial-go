package main

import (
	"financial/model"
	"fmt"
)

func main() {
	currencyType := model.NewCurrencyType(1, "Real")
	currency := model.NewCurrency(23.45, currencyType)
	refIndex := model.NewRefIndex(1, "IPCA")
	fmt.Println(currency)
	fmt.Println(refIndex)
}
