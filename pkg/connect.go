// Package gkudu provides primitives for sorting slices and user-defined
// collections.
package gkudu

import (
	"github.com/javiroman/gkudu/pkg/proto/kudu/client"
	"github.com/javiroman/gkudu/pkg/proto/kudu/common"
	pb "github.com/javiroman/gkudu/pkg/proto/kudu/master"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
	"strings"
)

// Connection comment ihere
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

// AddMasterServerAddr comment here
func (c *Connection) AddMasterServerAddr(a string) {
	c.Host = strings.Split(a, ":")[0]
	port := strings.Split(a, ":")[1]
	c.Port, _ = strconv.Atoi(port)
}

// Build comment here
func (c *Connection) Build() *client.ServerMetadataPB {
	location := c.Host
	host := c.Host
	var port uint32 = uint32(c.Port)
	rpc := []*common.HostPortPB{{
		Host: &host,
		Port: &port,
	}}

	return &client.ServerMetadataPB{
		RpcAddresses: rpc,
		Location:     &location,
	}
}

func (c *Connection) TestCoon() {
	// Set up a connection to the server.
	addr := "kudu-master1.node.keedio.cloud:7051"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.ListMasters()
	/*
		c := pb.NewGreeterClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	*/
}
