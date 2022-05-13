// Package gkudu provides primitives for sorting slices and user-defined
// collections.
package gkudu

type Connection struct {
	Host string
	Port int
}

// NewClientBuilder comment here
func NewClientBuilder() *Connection {
	return &Connection{
		Host: "localhost",
		Port: 8080,
	}
}

// Build comment here
func (k *Connection) Build() (string, int) {
	return k.Host, k.Port
}
