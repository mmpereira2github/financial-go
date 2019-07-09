package index

import (
	"fmt"
	"time"
)

// UpdateFactorCalculator calculates the update factor using an index and a time
type UpdateFactorCalculator interface {
	GetUpdateFactor(date time.Time) float64
}

type percentageBasedUpdateFactorCalculator struct {
	index *Index
}

// NewUpdateFactorCalculator creates a new UpdateFactorCalculator instance for given index
func NewUpdateFactorCalculator(index *Index) UpdateFactorCalculator {
	switch index.ValueType {
	case PercentageValueType:
		return &percentageBasedUpdateFactorCalculator{index}
	default:
		return nil
	}
}

func (c *percentageBasedUpdateFactorCalculator) GetUpdateFactor(date time.Time) float64 {
	var key string
	switch c.index.IntervalType {
	case Monthly:
		key = fmt.Sprintf("%04d%02d", date.Year(), date.Month())
	default:
		panic(fmt.Errorf("Do know how to build key to get value using c.index.IntervalType=%d", c.index.IntervalType))
	}
	factor := c.index.Values[key]
	return (factor / 100) + 1.0
}
