package index

import (
	"fmt"
	"time"
)

// TimeIterator calculates the next time based on the interval type
type TimeIterator interface {
	Next() time.Time
}

type monthlyTimeIterator struct {
	currentTime time.Time
}

func (b *monthlyTimeIterator) Next() time.Time {
	b.currentTime = b.currentTime.AddDate(0, 1, 0)
	return b.currentTime
}

// NewTimeIterator creates a new TimeInterator for given interval type and start time
func NewTimeIterator(startTime time.Time, intervalType IntervalType) (TimeIterator, error) {
	switch intervalType {
	case Monthly:
		return &monthlyTimeIterator{startTime}, nil
	default:
		return nil, fmt.Errorf("No TimeIterator found for intervalType=%d", intervalType)
	}
}
