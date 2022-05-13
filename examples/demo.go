package main

import (
	"fmt"
	kudu "github.com/javiroman/gkudu/pkg"
)

func main() {
	fmt.Println("Go Kudu Library")

	client := kudu.NewClientBuilder()
	//client.AddMasterServerAddr("localhost")
	c, i := client.Build()
	fmt.Println(c, i)
}
