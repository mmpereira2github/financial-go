package index

// ValueType identifies what represents the indicator value
type ValueType int

const (
	// PercentageValueType says each indicator value is a percentage value related to value before
	PercentageValueType ValueType = 0
	// AjustedValueType says each indicator value is an ajusted value and percentage is obtained using previues value (% = curr/previous -1)
	AjustedValueType ValueType = 1
)

// IntervalType specifies the interval between indicator values
type IntervalType int

const (
	// Monthly means there is a value per month only
	Monthly IntervalType = 0
)

// IndexValue is the value of an index at date
type IndexValue struct {
	Value float64
	Date  int
}

// Index means a referenced index in the financial market used to compare performance
type Index struct {
	ID           string
	ValueType    ValueType
	IntervalType IntervalType
	Values       map[string]float64
}
