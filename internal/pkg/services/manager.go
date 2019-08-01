package services

import (
	"fmt"
	"log"
)

// ServiceEntry means a service that can be executed through manager
type ServiceEntry struct {
	ID           string
	InputFactory func() interface{}
	Invoke       func(interface{}) (interface{}, *Status)
}

// Manager supports the registry, unregistry and retrival of services
var Manager interface {
	GetServiceEntryByID(id string) (*ServiceEntry, *Status)
	Register(entry *ServiceEntry) error
}

type manager struct {
	registry map[string]*ServiceEntry
}

func init() {
	Manager = &manager{
		registry: make(map[string]*ServiceEntry, 10),
	}
}

func (r *manager) GetServiceEntryByID(id string) (*ServiceEntry, *Status) {
	if s := r.registry[id]; s != nil {
		return s, nil
	}
	return nil, &Status{
		Code:  ServiceNotFound,
		Error: fmt.Errorf("Service %s not found", id),
	}
}

func (r *manager) Register(entry *ServiceEntry) error {
	log.Println("Registering=>", entry)
	r.registry[entry.ID] = entry
	return nil
}
