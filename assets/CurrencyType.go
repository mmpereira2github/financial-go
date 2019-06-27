package assets

import "fmt"

type currencyType struct {
	id   int
	name string
}

type CurrencyType interface {
	Id() int
	Name() string
	String() string
}

func NewCurrencyType(id int, name string) CurrencyType { return currencyType{id, name} }
func (c currencyType) Id() int                         { return c.id }
func (c currencyType) Name() string                    { return c.name }
func (c currencyType) String() string {
	return fmt.Sprintf("CurrencyType(id=%d, name='%s')", c.id, c.name)
}
