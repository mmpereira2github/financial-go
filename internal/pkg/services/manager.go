package services

import "log"

// ServiceEntry means a service that can be executed through manager
type ServiceEntry struct {
	ID           string
	InputFactory func() interface{}
	Invoke       func(interface{}) (Status, interface{})
}

var Manager interface {
	GetServiceEntryById(id string) *ServiceEntry
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

func (r *manager) GetServiceEntryById(id string) *ServiceEntry {
	return r.registry[id]
}

func (r *manager) Register(entry *ServiceEntry) error {
	log.Println("Registering=>", entry)
	r.registry[entry.ID] = entry
	return nil
}
