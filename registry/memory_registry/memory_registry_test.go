package memory_registry

import (
	"testing"

	"github.com/taadis/letgo/registry"
)

func TestMemoryRegistry(t *testing.T) {
	tests := []*registry.Service{
		{
			Name:    "foo",
			Version: "0.0.1",
			Nodes: []*registry.Node{
				{
					Id:      "foo-0.0.1-1",
					Address: "localhost",
					Port:    11111,
				},
				{
					Id:      "foo-0.0.1-2",
					Address: "localhost",
					Port:    11112,
				},
			},
		},
		{
			Name:    "bar",
			Version: "0.0.1",
			Nodes: []*registry.Node{
				{
					Id:      "bar-0.0.1-1",
					Address: "localhost",
					Port:    11198,
				},
				{
					Id:      "bar-0.0.1-2",
					Address: "localhost",
					Port:    11199,
				},
			},
		},
	}

	r := NewRegistry()
	wantString := "memory_registry"
	if r.String() != wantString {
		t.Fatalf("want %s but got %s", wantString, r.String())
	}

	// register
	for _, service := range tests {
		err := r.Register(service)
		if err != nil {
			t.Fatalf("register error: %v", err)
		}
	}

	// using
	services, err := r.GetService("foo")
	if err != nil {
		t.Fatalf("registry get service error: %v", err)
	}
	t.Logf("got services: %+v", services)

	// deregister
	for _, service := range tests {
		err := r.Deregister(service)
		if err != nil {
			t.Fatalf("deregister error: %v", err)
		}
	}
}
