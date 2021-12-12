package server

import (
	"github.com/daniilty/gismeteo-weather-gateway/internal/core"
	schema "github.com/daniilty/weather-gateway-schema"
)

type GRPC struct {
	schema.UnimplementedGismeteoWeatherGatewayServer

	service core.Service
}

func NewGRPC(service core.Service) GRPC {
	return GRPC{
		service: service,
	}
}
