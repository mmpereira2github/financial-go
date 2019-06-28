package indicator

const precision = 1000000000

type monthlyPerfIndicatorValue struct {
	perfIndicatorCategory PerfIndicatorCategory
	year                  int
	month                 int
	value                 int64
}

// MonthlyPerfIndicatorValue means a performance indicator value for a given month
type MonthlyPerfIndicatorValue interface {
	PerfIndicatorCategory() PerfIndicatorCategory
	Year() int
	Month() int
	Float64() float64
}

// NewMonthlyPerfIndicatorValue creates a new MonthlyPerfIndicatorValue instance for gibev value for year/month and performance indicator type
func NewMonthlyPerfIndicatorValue(year int, month int, value float64, perfIndicatorCategory PerfIndicatorCategory) MonthlyPerfIndicatorValue {
	return &monthlyPerfIndicatorValue{year: year, month: month, value: int64(value * precision), perfIndicatorCategory: perfIndicatorCategory}
}

func (i *monthlyPerfIndicatorValue) PerfIndicatorCategory() PerfIndicatorCategory {
	return i.perfIndicatorCategory
}
func (i *monthlyPerfIndicatorValue) Year() int        { return i.year }
func (i *monthlyPerfIndicatorValue) Month() int       { return i.month }
func (i *monthlyPerfIndicatorValue) Float64() float64 { return float64(i.value) / precision }
