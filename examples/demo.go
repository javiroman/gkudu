package main

import (
	"fmt"
	kudu "github.com/javiroman/gkudu/pkg"
)

func main() {
	fmt.Println("Go Kudu Library")

	client := kudu.NewClientBuilder()
	client.AddMasterServerAddr("kudu-master1.node.keedio.cloud:7051")
	c := client.Build()

	fmt.Println(c)
	client.TestCoon()
}
