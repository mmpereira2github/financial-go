package indicator

var indicatorDao = inMemoryIndicatorDao{}

// Dao provides access to repository operations related to Indicator
type Dao interface {
	FindByRange(cat PerfIndicatorCategory, startYear int, startMonth int, endYear int, endMonth int) []MonthlyPerfIndicatorValue
	Save(PerfIndicatorCategory, MonthlyPerfIndicatorValue) error
}

// GetIndicatorDao return the implementation of PerfIndicatorTypeDao
func GetIndicatorDao() Dao { return &indicatorDao }
