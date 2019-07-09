package assets

import "time"

// PurchaseActivity means an activity to purchase some units of an investment
type PurchaseActivity interface {
	Date() time.Time
	//Costs() []Cost
	Price() float64
	Units() float64
}

type purchaseActivity struct {
	date  time.Time
	price float64
	units float64
}

// NewPurchaseActivity creates a purchase activity without costs
func NewPurchaseActivity(date *time.Time, price float64, units float64) PurchaseActivity {
	return &purchaseActivity{*date, price, units}
}
func (p *purchaseActivity) Date() time.Time { return p.date }
func (p *purchaseActivity) Price() float64  { return p.price }
func (p *purchaseActivity) Units() float64  { return p.units }
