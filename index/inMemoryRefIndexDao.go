package index

import "fmt"

var inMemoryRefIndexRepo = map[int]RefIndex{
	1: &refIndex{1, "IPCA"},
	2: &refIndex{2, "CDI"},
}

type inMemoryRefIndexDao struct{}

//	getById(int) (RefIndex, error)
//	findByName(string) RefIndex
//	save(RefIndex) error

func (d inMemoryRefIndexDao) GetById(id int) (RefIndex, error) {
	if refIndex, ok := inMemoryRefIndexRepo[id]; ok {
		return refIndex, nil
	} else {
		return nil, fmt.Errorf("No RefIndex found for id=%d", id)
	}
}

func (d inMemoryRefIndexDao) FindByName(name string) RefIndex {
	for _, ref := range inMemoryRefIndexRepo {
		if ref.Name() == name {
			return ref
		}
	}
	return nil
}

func (d inMemoryRefIndexDao) Save(refIndex RefIndex) error {
	inMemoryRefIndexRepo[refIndex.Id()] = refIndex
	return nil
}
