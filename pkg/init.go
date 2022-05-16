// Package gkudu provides primitives for sorting slices and user-defined
// collections.
package gkudu

import (
	"github.com/javiroman/gkudu/pkg/proto/kudu/client"
	"github.com/javiroman/gkudu/pkg/proto/kudu/common"
)

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

func (c *Connection) AddMasterServerAddr(h string) {
	c.Host = h
}

// Build comment here
func (c *Connection) Build() (string, int) {
	return c.Host, c.Port
}

func TestPB() *client.ServerMetadataPB {
	location := "location"
	host := "localhost"
	var port uint32 = 8080
	rpc := []*common.HostPortPB{{
		Host: &host,
		Port: &port,
	}}

	return &client.ServerMetadataPB{
		RpcAddresses: rpc,
		Location:     &location,
	}
}
