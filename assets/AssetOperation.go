package assets

import "time"

// ActivityVisitor visits an activity to find out the implementation
type ActivityVisitor interface {
	VisitSaleActivity(op *SaleActivity)
	VisitPurchaseActivity(op *PurchaseActivity)
	VisitIncomeActivity(op *IncomeActivity)
}

// Activity means an investment activity
type Activity interface {
	ID() int64
	When() time.Time
	Memo() string
}

// CostCategory means a type of cost
type CostCategory interface {
	ID() int
	Name() string
	Description() string
}

// Cost means some amount paid to buy/keep something
type Cost interface {
	CostType() CostCategory
	CostTypeID() int
	Amount() float64
}

// PurchaseActivity means an activity to purchase some units of an investment
type PurchaseActivity interface {
	ID() int64
	When() time.Time
	Costs() []Cost
	Price() float64
	Units() float64
}

// SaleActivity means an activity to sale some units of an investment
type SaleActivity interface {
	ID() int64
	When() time.Time
	Costs() []Cost
	Price() Currency
	Units() int
}

// IncomeActivity means an investment income
type IncomeActivity interface {
	ID() int64
	When() time.Time
	Costs() []Cost
	Amount() float64
}
