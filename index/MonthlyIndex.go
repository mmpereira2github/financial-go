package index

const precision = 1000000000

type monthlyIndex struct {
	refIndexId int
	year       int
	month      int
	value      int64
}

type MonthlyIndex interface {
	RefIndex() RefIndex
	Year() int
	Month() int
	Float64() float64
}

func NewMonthlyIndex(year int, month int, value float64, refIndex RefIndex) MonthlyIndex {
	return monthlyIndex{year: year, month: month, value: int64(value * precision), refIndexId: refIndex.Id()}
}
func (i monthlyIndex) RefIndex() RefIndex {
	if ref, error := GetRefIndexDao().GetById(i.refIndexId); error == nil {
		return ref
	} else {
		panic(error)
	}

}
func (i monthlyIndex) Year() int        { return i.year }
func (i monthlyIndex) Month() int       { return i.month }
func (i monthlyIndex) Float64() float64 { return float64(i.value) / precision }
