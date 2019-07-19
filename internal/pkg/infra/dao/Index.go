package dao

import "financial/internal/pkg/index"

var indexDaoImpl *IndexDao

// IndexDao is the interface to access indexes in the application
type IndexDao interface {
	FindByID(id string) (*index.Index, error)
}

// SetIndexDao set the implementation of IndexDao to be used
func SetIndexDao(impl IndexDao) {
	indexDaoImpl = &impl
}

// GetIndexDao returns the IndexDao that shall be used
func GetIndexDao() *IndexDao { return indexDaoImpl }
