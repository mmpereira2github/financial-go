package indicator

var perfIndicatorTypeDao = inMemoryPerfIndicatorTypeDao{}

// PerfIndicatorTypeDao perform operations
type PerfIndicatorTypeDao interface {
	GetByID(int) (PerfIndicatorType, error)
	FindByName(string) PerfIndicatorType
	Save(PerfIndicatorType) error
}

// GetPerfIndicatorTypeDao return the implementation of PerfIndicatorTypeDao
func GetPerfIndicatorTypeDao() PerfIndicatorTypeDao { return &perfIndicatorTypeDao }
