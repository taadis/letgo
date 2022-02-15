package memory_registry

import (
	"github.com/taadis/letgo/registry"
)

var (
	store = map[string][]*registry.Registry{}
	//Store = map[string][]*registry.Service{
	//	"foo":{
	//		Name:    "foo",
	//		Version: "0.0.1",
	//		Nodes: []*registry.Node{
	//			{
	//				Id:      "foo-0.0.1-1",
	//				Address: "localhost",
	//				Port:    11111,
	//			},
	//			{
	//				Id:     "foo-0.0.1-2",
	//				Addres: "localhost",
	//				Port:   11111,
	//			},
	//		},
	//	},
	//	{
	//		Name:    "foo",
	//		Version: "0.0.2",
	//		Nodes: []*registry.Node{
	//			{
	//				Id:      "foo-0.0.2-1",
	//				Address: "localhost",
	//				Port:    11112,
	//			},
	//		},
	//	},
	//	{
	//		Name:    "foo",
	//		Version: "0.0.3",
	//		Nodes: []*registry.Node{
	//			{
	//				Id:      "foo-0.0.3-1",
	//				Address: "localhost",
	//				Port:    11113,
	//			},
	//		},
	//	},
	//}
)

func addNodes(oldNodes []*registry.Node, addNodes []*registry.Node) []*registry.Node {
	for _, addNode := range addNodes {
		var seen bool
		for i, oldNode := range oldNodes {
			if oldNode.Id == addNode.Id {
				seen = true
				oldNodes[i] = addNode
				break
			}
		}
		if !seen {
			oldNodes = append(oldNodes, addNode)
		}
	}
	return oldNodes
}

func delNodes(oldNodes []*registry.Node, delNodes []*registry.Node) []*registry.Node {
	var nodes []*registry.Node
	for _, oldNode := range oldNodes {
		var rem bool
		for _, delNode := range delNodes {
			if delNode.Id == oldNode.Id {
				rem = true
				break
			}
		}
		if !rem {
			nodes = append(nodes, oldNode)
		}
	}
	return nodes
}

func addServices(oldServices, addServices []*registry.Service) []*registry.Service {
	for _, addService := range addServices {
		var seen bool
		for i, oldService := range oldServices {
			if oldService.Version == addService.Version {
				addService.Nodes = addNodes(oldService.Nodes, addService.Nodes)
				seen = true
				oldServices[i] = addService
				break
			}
			if seen {
				oldServices = append(oldServices, addService)
			}
		}
	}
	return oldServices
}

func delServices(oldServices []*registry.Service, delServices []*registry.Service) []*registry.Service {
	var services []*registry.Service
	for i, oldService := range oldServices {
		var rem bool
		for _, delService := range delServices {
			if oldService.Version == delService.Version {
				if len(oldServices[i].Nodes) == 0 {
					rem = true
				}
			}
		}
		if !rem {
			services = append(services, oldService)
		}
	}
	return services
}
