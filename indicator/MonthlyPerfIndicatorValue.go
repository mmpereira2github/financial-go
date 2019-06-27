package indicator

const precision = 1000000000

type monthlyPerfIndicatorValue struct {
	refIndexID int
	year       int
	month      int
	value      int64
}

// MonthlyPerfIndicatorValue means a performance indicator value for a given month
type MonthlyPerfIndicatorValue interface {
	PerfIndicatorType() PerfIndicatorType
	Year() int
	Month() int
	Float64() float64
}

// NewMonthlyPerfIndicatorValue creates a new MonthlyPerfIndicatorValue instance for gibev value for year/month and performance indicator type
func NewMonthlyPerfIndicatorValue(year int, month int, value float64, refIndex PerfIndicatorType) MonthlyPerfIndicatorValue {
	return &monthlyPerfIndicatorValue{year: year, month: month, value: int64(value * precision), refIndexID: refIndex.ID()}
}
func (i *monthlyPerfIndicatorValue) PerfIndicatorType() PerfIndicatorType {
	ref, error := GetPerfIndicatorTypeDao().GetByID(i.refIndexID)
	if error == nil {
		return ref
	}
	panic(error)
}

func (i *monthlyPerfIndicatorValue) Year() int        { return i.year }
func (i *monthlyPerfIndicatorValue) Month() int       { return i.month }
func (i *monthlyPerfIndicatorValue) Float64() float64 { return float64(i.value) / precision }
