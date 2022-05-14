package main

import (
	"fmt"
	kudu "github.com/javiroman/gkudu/pkg"
)

func main() {
	fmt.Println("Go Kudu Library")

	client := kudu.NewClientBuilder()
	client.AddMasterServerAddr("127.0.0.1")
	c, i := client.Build()

	fmt.Println(c, i)

	fmt.Println(kudu.TestPB())
}
