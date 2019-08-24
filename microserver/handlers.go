package microserver

import (
	"context"
	"net/http"

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
	return nil
}
