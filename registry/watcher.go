package registry

import "time"

// Watcher 是一个接口,用于返回 registry 中服务的更新信息
type Watcher interface {
	// Next is a blocking call
	Next() (*Result, error)
	Stop()
}

// Result is returned by a call to next on the watcher.
// Actions can be: create, update, delete.
type Result struct {
	Action  string
	Service *Service
}

// EventType defines registry event type
type EventType int

const (
	Create EventType = 0
	Delete EventType = 1
	Update EventType = 2
)

func (t EventType) String() string {
	switch t {
	case Create:
		return "create"
	case Delete:
		return "delete"
	case Update:
		return "update"
	default:
		return "unknown"
	}
}

// Event is registry event
type Event struct {
	// Id is registry id
	Id string
	// Type defines type of event
	Type EventType
	// Timestamp is event timestamp
	Timestamp time.Time
	// Service is registry service
	Service *Service
}
