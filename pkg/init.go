// Package gkudu provides primitives for sorting slices and user-defined
// collections.
package gkudu

import (
	pb "github.com/javiroman/gkudu/pkg/proto"
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

func TestPB() *pb.ServerMetadataPB {
	location := "location"
	host := "localhost"
	var port uint32 = 8080
	rpc := []*pb.HostPortPB{{
		Host: &host,
		Port: &port,
	}}

	return &pb.ServerMetadataPB{
		RpcAddresses: rpc,
		Location:     &location,
	}
}
