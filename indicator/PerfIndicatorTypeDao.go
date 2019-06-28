package indicator

var perfIndicatorTypeDao = inMemoryPerfIndicatorTypeDao{}

// PerfIndicatorTypeDao perform operations
type PerfIndicatorTypeDao interface {
	FindByName(string) PerfIndicatorCategory
	Save(PerfIndicatorCategory) error
}

// GetPerfIndicatorTypeDao return the implementation of PerfIndicatorTypeDao
func GetPerfIndicatorTypeDao() PerfIndicatorTypeDao { return &perfIndicatorTypeDao }
