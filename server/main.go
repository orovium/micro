package main

import (
	"context"
	"fmt"
	"net/http"

	micro "github.com/micro/go-micro"

	proto "github.com/orovium/micro/proto"
)

// Ping is the main entrypoint for this service
type Ping struct{}

// Ping is the handler for the Ping call. Ping returns pong if service is alive.
func (p *Ping) Ping(ctx context.Context, req *proto.PingRequest, rsp *proto.PingResponse) error {
	rsp.Response = &proto.ResponseEnvelope{
		ServiceMethod: http.MethodGet,
		Seq:           1,
		Error:         "",
		HttpCode:      http.StatusOK,
	}

	rsp.Message = "pong"
	fmt.Printf("Context: %v", ctx)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("ping"),
		micro.Version("1.0.0"),
		micro.Address(":8091"),
	)

	service.Init()

	proto.RegisterPingHandler(service.Server(), new(Ping))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
