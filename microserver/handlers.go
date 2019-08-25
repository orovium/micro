package microserver

import (
	"context"
	"database/sql"
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
	GetLogger().Trace(ctx)
	service, _ := GetService()
	rsp.Status = getHealthStats(service)
	return nil
}

func getHealthStats(service *Service) *proto.ServiceStatus {
	return &proto.ServiceStatus{
		DBStats: getDBStats(service),
	}
}

func getDBStats(service *Service) *proto.DBStats {
	db, err := service.GetDB()
	if err != nil {
		return &proto.DBStats{}
	}

	return dbStats2proto(db.Stats())
}

func dbStats2proto(stats sql.DBStats) *proto.DBStats {
	return &proto.DBStats{
		MaxOpenConnections: int64(stats.MaxOpenConnections),
		OpenConnections:    int64(stats.OpenConnections),
		InUse:              int64(stats.InUse),
		Idle:               int64(stats.Idle),
		WaitCount:          int64(stats.WaitCount),
		WaitDuration:       int64(stats.WaitDuration),
		MaxIdleClosed:      int64(stats.MaxIdleClosed),
		MaxLifetimeClosed:  int64(stats.MaxLifetimeClosed),
	}
}
