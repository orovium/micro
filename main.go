package main

import (
	"fmt"

	"github.com/micro/go-micro"

	"github.com/orovium/micro/microserver"
)

func main() {
	microserver.Init(nil)

	service := microserver.StartDefaultService(
		micro.Name("ping"),
		micro.Version("1.0.0"),
		micro.Address(":8091"),
	)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
