package memory_registry

import (
	"sync"

	"github.com/taadis/letgo/errors"
	"github.com/taadis/letgo/registry"
)

type Registry struct {
	sync.RWMutex
	services map[string][]*registry.Service
	//Watchers map[string]*Watcher
}

func (r *Registry) String() string {
	return "memory_registry"
}

func (r *Registry) Register(service *registry.Service) error {
	r.Lock()
	defer r.Unlock()

	services := addServices(r.services[service.Name], []*registry.Service{service})
	r.services[service.Name] = services
	return nil
}

func (r *Registry) Deregister(service *registry.Service) error {
	r.Lock()
	defer r.Unlock()

	services := delServices(r.services[service.Name], []*registry.Service{service})
	r.services[service.Name] = services
	return nil
}

func (r *Registry) GetService(name string) ([]*registry.Service, error) {
	r.RLock()
	defer r.RUnlock()

	services, ok := r.services[name]
	if !ok || len(services) == 0 {
		return nil, errors.NotFound("", "not found service %s", name)
	}
	return services, nil
}

func NewRegistry() registry.Registry {
	services := make(map[string][]*registry.Service)

	return &Registry{
		services: services,
	}
}
