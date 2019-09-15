package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"

	proto "orovio/micro/proto"
)

func main() {
	service := micro.NewService(micro.Name("ping.client"))
	service.Init()

	ping := proto.NewPingService("ping", service.Client())

	rsp, err := ping.Ping(context.TODO(), &proto.PingRequest{
		Request: &proto.RequestEnvelope{
			ServiceMethod: "GET",
			Seq:           1,
			Headers: []*proto.Header{
				&proto.Header{
					Key:   "Authorization",
					Value: "Bearer myToken",
				},
			},
		},
		Message: "ping",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp)
}
