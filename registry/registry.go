// Package registry is an interface for service registry and discovery.
package registry

type Registry interface {
	Register(*Service) error
	Deregister(*Service) error
	GetService(name string) ([]*Service, error)
}

type Service struct {
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Metadata  map[string]string `json:"metadata"`
	Nodes     []*Node           `json:"nodes"`
	Endpoints []*Endpoint       `json:"endpoints"`
}

type Node struct {
	Id       string            `json:"id"`
	Address  string            `json:"address"`
	Port     int               `json:"port"`
	Metadata map[string]string `json:"metadata"`
}

type Endpoint struct {
	Name     string            `json:"name"`
	Request  *Value            `json:"request"`
	Response *Value            `json:"response"`
	Metadata map[string]string `json:"metadata"`
}

type Value struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []*Value `json:"values"`
}
