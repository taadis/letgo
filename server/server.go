// Package server
package server

import "fmt"

type Server interface {
	fmt.Stringer
	// Start the server
	Start() error
	// Stop the server
	Stop() error
}