package indicator

import "fmt"

const (
	notFoundErrMsg = "No PerfIndicatorCategory found for id=%d"
)

var inMemoryPerfIndicatorTypeRepo = map[string]PerfIndicatorCategory{}

type inMemoryPerfIndicatorTypeDao struct{}

func (d *inMemoryPerfIndicatorTypeDao) FindByName(name string) PerfIndicatorCategory {
	return inMemoryPerfIndicatorTypeRepo[name]
}

func (d *inMemoryPerfIndicatorTypeDao) Save(perfIndicatorCategory PerfIndicatorCategory) error {
	perf := d.FindByName(perfIndicatorCategory.Name())
	if perf == nil {
		inMemoryPerfIndicatorTypeRepo[perfIndicatorCategory.Name()] = perfIndicatorCategory
		return nil
	}
	return fmt.Errorf("Duplicated %v", perfIndicatorCategory)
}
