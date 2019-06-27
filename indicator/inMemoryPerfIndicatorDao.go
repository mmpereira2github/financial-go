package indicator

import "fmt"

const (
	notFoundErrMsg = "No PerfIndicatorType found for id=%d"
)

var inMemoryPerfIndicatorTypeRepo = map[int]PerfIndicatorType{
	1: &perfIndicatorType{1, "IPCA"},
	2: &perfIndicatorType{2, "CDI"},
}

type inMemoryPerfIndicatorTypeDao struct{}

func (d *inMemoryPerfIndicatorTypeDao) GetByID(id int) (PerfIndicatorType, error) {
	refIndex, ok := inMemoryPerfIndicatorTypeRepo[id]
	if ok {
		return refIndex, nil
	}
	return nil, fmt.Errorf(notFoundErrMsg, id)
}

func (d *inMemoryPerfIndicatorTypeDao) FindByName(name string) PerfIndicatorType {
	for _, ref := range inMemoryPerfIndicatorTypeRepo {
		if ref.Name() == name {
			return ref
		}
	}
	return nil
}

func (d *inMemoryPerfIndicatorTypeDao) Save(refIndex PerfIndicatorType) error {
	inMemoryPerfIndicatorTypeRepo[refIndex.ID()] = refIndex
	return nil
}
