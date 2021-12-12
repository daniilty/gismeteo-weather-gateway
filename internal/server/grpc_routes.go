package server

import (
	"context"

	schema "github.com/daniilty/weather-gateway-schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g GRPC) GetWeather(ctx context.Context, r *schema.Empty) (*schema.Weather, error) {
	weather, err := g.service.GetWeather()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.Weather{
		Info: weather,
	}, nil
}
